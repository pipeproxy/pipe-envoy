package adsc

import (
	"context"
	"reflect"
	"sort"
	"sync"
	"time"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/wzshiming/logger"
)

type ADSC struct {
	mutex sync.Mutex
	ctx   context.Context
	node  *envoy_config_core_v3.Node

	XDSClient *XDSClient
	SDSClient *SDSClient

	watchTime   time.Time
	initialLoad time.Duration

	watchCDS bool
	watchLDS bool

	// httpLDS contains received listeners with a http_connection_manager filter.
	httpLDS map[string]*envoy_config_listener_v3.Listener

	// tcpLDS contains all listeners of type TCP (not-HTTP)
	tcpLDS map[string]*envoy_config_listener_v3.Listener

	// All received cds of type eds, keyed by name
	edsCDS map[string]*envoy_config_cluster_v3.Cluster

	// All received cds of no-eds type, keyed by name
	cds map[string]*envoy_config_cluster_v3.Cluster

	// All received rds, keyed by route name
	rds map[string]*envoy_config_route_v3.RouteConfiguration

	// All received eds, keyed by cluster name
	eds map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment

	// All received sds, keyed by secret name
	sds map[string]*envoy_extensions_transport_sockets_tls_v3.Secret

	HandleEDSCDS  func(clusters map[string]*envoy_config_cluster_v3.Cluster)
	HandleCDS     func(clusters map[string]*envoy_config_cluster_v3.Cluster)
	HandleRDS     func(clusters map[string]*envoy_config_route_v3.RouteConfiguration)
	HandleHttpLDS func(listeners map[string]*envoy_config_listener_v3.Listener)
	HandleTcpLDS  func(listeners map[string]*envoy_config_listener_v3.Listener)
	HandleEDS     func(clusters map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment)
	HandleSDS     func(secrets map[string]*envoy_extensions_transport_sockets_tls_v3.Secret)
}

func NewADSC(xdsAddress, sdsAddress string, node *envoy_config_core_v3.Node) *ADSC {
	a := &ADSC{
		node:     node,
		watchCDS: true,
		watchLDS: true,
	}

	a.XDSClient = NewXDSClient(xdsAddress, nil, a.xdsConfig())
	if sdsAddress != "" {
		a.SDSClient = NewSDSClient(sdsAddress, nil, a.sdsConfig())
	}
	return a
}

func (a *ADSC) Run(ctx context.Context) error {
	a.ctx = ctx
	log := logger.FromContext(ctx)
	if a.SDSClient != nil {
		go func() {
			log := log.WithName("sds")
			ctx = logger.WithContext(ctx, log)
			for ctx.Err() == nil {
				err := a.SDSClient.Run(ctx)
				if err != nil {
					log.Error(err, "run sds")
				}
				time.Sleep(time.Second)
			}
		}()
	}

	log = log.WithName("xds")
	ctx = logger.WithContext(ctx, log)
	for ctx.Err() == nil {
		err := a.XDSClient.Run(ctx)
		if err != nil {
			log.Error(err, "run xds")
		}
		time.Sleep(time.Second)
	}
	return nil
}

func (a *ADSC) xdsConfig() *XDSConfig {
	return &XDSConfig{
		Node:      a.node,
		OnConnect: a.watchXDS,
		HandleCDS: a.handleCDS,
		HandleRDS: a.handleRDS,
		HandleLDS: a.handleLDS,
		HandleEDS: a.handleEDS,
	}
}
func (a *ADSC) sdsConfig() *SDSConfig {
	return &SDSConfig{
		Node:      a.node,
		HandleSDS: a.handleSDS,
	}
}

func (a *ADSC) watchXDS(cli *XDSClient) error {
	a.watchTime = time.Now()
	if a.watchCDS {
		err := cli.SendRsc(ClusterType, nil)
		if err != nil {
			return err
		}
	} else {
		if a.watchLDS {
			err := cli.SendRsc(ListenerType, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *ADSC) handleCDS(cli *XDSClient, clusters []*envoy_config_cluster_v3.Cluster) {
	logger.FromContext(a.ctx).Info("CDS", "size", len(clusters))
	cn := make([]string, 0, len(clusters))
	secrets := []string{}
	edscds := map[string]*envoy_config_cluster_v3.Cluster{}
	cds := map[string]*envoy_config_cluster_v3.Cluster{}
	for _, c := range clusters {
		secrets = append(secrets, GetSDSName(c.TransportSocket)...)
		switch v := c.ClusterDiscoveryType.(type) {
		case *envoy_config_cluster_v3.Cluster_Type:
			if v.Type != envoy_config_cluster_v3.Cluster_EDS {
				cds[c.Name] = c
				continue
			}
		}
		cn = append(cn, c.Name)
		edscds[c.Name] = c
	}

	cn = removeDuplicates(cn)
	if len(cn) != 0 {
		cli.SendRsc(EndpointType, cn)
	}

	if a.SDSClient != nil {
		secrets = removeDuplicates(secrets)
		for _, secret := range secrets {
			a.SDSClient.SendRsc(SecretType, []string{secret})
		}
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.edsCDS, edscds) {
		a.edsCDS = edscds
		if a.HandleEDSCDS != nil {
			a.HandleEDSCDS(edscds)
		}
	}
	if !reflect.DeepEqual(a.cds, cds) {
		a.cds = cds
		if a.HandleCDS != nil {
			a.HandleCDS(cds)
		}
	}
}

func (a *ADSC) handleEDS(cli *XDSClient, eds []*envoy_config_endpoint_v3.ClusterLoadAssignment) {
	logger.FromContext(a.ctx).Info("EDS", "size", len(eds))
	la := map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment{}

	for _, cla := range eds {
		la[cla.ClusterName] = cla
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.eds, la) {
		a.eds = la
		if a.HandleEDS != nil {
			a.HandleEDS(la)
		}
	}

	if a.initialLoad == 0 {
		if a.watchLDS {
			cli.SendRsc(ListenerType, nil)
		} else {
			a.initialDone()
		}
	}
}

func (a *ADSC) handleLDS(cli *XDSClient, listeners []*envoy_config_listener_v3.Listener) {
	logger.FromContext(a.ctx).Info("LDS", "size", len(listeners))
	lh := map[string]*envoy_config_listener_v3.Listener{}
	lt := map[string]*envoy_config_listener_v3.Listener{}

	routes := []string{}
	secrets := []string{}
	for _, l := range listeners {

		// The last filter is the actual destination for inbound listener
		if l.ApiListener != nil || len(l.FilterChains) == 0 {
			// This is an API Listener
			// TODO: extract VIP and rds or cluster
			continue
		}
		filterChain := SelectFilterChain(l.FilterChains)
		secrets = append(secrets, GetSDSName(filterChain.TransportSocket)...)
		filter := filterChain.Filters[0]
		if filter.Name == wellknown.TCPProxy {
			lt[l.Name] = l
		} else if filter.Name == wellknown.HTTPConnectionManager {
			lh[l.Name] = l

			config := GetHTTPConnectionManager(filter)
			if config == nil {
				continue
			}
			if rds, ok := config.RouteSpecifier.(*envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager_Rds); ok && rds != nil && rds.Rds != nil {
				routes = append(routes, rds.Rds.RouteConfigName)
			}
			//// Getting from config is too painful..
			//port := l.Address.GetSocketAddress().GetPortValue()
			//if port == 15002 {
			//	rds = append(rds, "http_proxy")
			//} else {
			//	rds = append(rds, fmt.Sprintf("%d", port))
			//}
		}
	}

	routes = removeDuplicates(routes)
	if len(routes) != 0 {
		cli.SendRsc(RouteType, routes)
	}

	if a.SDSClient != nil {
		secrets = removeDuplicates(secrets)
		for _, secret := range secrets {
			a.SDSClient.SendRsc(SecretType, []string{secret})
		}
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.httpLDS, lh) {
		a.httpLDS = lh
		if a.HandleHttpLDS != nil {
			a.HandleHttpLDS(lh)
		}
	}
	if !reflect.DeepEqual(a.tcpLDS, lt) {
		a.tcpLDS = lt
		if a.HandleTcpLDS != nil {
			a.HandleTcpLDS(lt)
		}
	}
}

func (a *ADSC) handleRDS(xds *XDSClient, configurations []*envoy_config_route_v3.RouteConfiguration) {
	logger.FromContext(a.ctx).Info("RDS", "size", len(configurations))
	rds := map[string]*envoy_config_route_v3.RouteConfiguration{}
	for _, r := range configurations {
		rds[r.Name] = r
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.rds, rds) {
		a.rds = rds
		if a.HandleRDS != nil {
			a.HandleRDS(rds)
		}
	}

	if a.initialLoad == 0 {
		a.initialDone()
	}
}

func (a *ADSC) initialDone() {
	a.initialLoad = time.Since(a.watchTime)
}

func (a *ADSC) handleSDS(cli *SDSClient, secrets []*envoy_extensions_transport_sockets_tls_v3.Secret) {
	logger.FromContext(a.ctx).Info("SDS", "size", len(secrets))
	sds := map[string]*envoy_extensions_transport_sockets_tls_v3.Secret{}
	for _, r := range secrets {
		sds[r.Name] = r
	}
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.sds, sds) {
		a.sds = sds
		if a.HandleSDS != nil {
			a.HandleSDS(sds)
		}
	}
}

func removeDuplicates(a []string) (ret []string) {
	if len(a) <= 1 {
		return a
	}
	sort.Strings(a)
	ret = a[:1]
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return ret
}

type cache struct {
	VersionInfo string
	Nonce       string
	Names       []string
}

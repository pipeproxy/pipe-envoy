package adsc

import (
	"context"
	"crypto/tls"
	"reflect"
	"sort"
	"sync"
	"time"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/wzshiming/xds/utils"
	xds_v3 "github.com/wzshiming/xds/v3"
)

type ADSC struct {
	mutex sync.Mutex

	*xds_v3.Client

	watchTime   time.Time
	initialLoad time.Duration

	watchCDS bool
	watchLDS bool

	// httpLDS contains received listeners with a http_connection_manager filter.
	httpLDS map[string]*envoy_config_listener_v3.Listener

	// tcpLDS contains all listeners of type TCP (not-HTTP)
	tcpLDS map[string]*envoy_config_listener_v3.Listener

	// All received CDS of type eds, keyed by name
	edsCDS map[string]*envoy_config_cluster_v3.Cluster

	// All received CDS of no-eds type, keyed by name
	CDS map[string]*envoy_config_cluster_v3.Cluster

	// All received RDS, keyed by route name
	RDS map[string]*envoy_config_route_v3.RouteConfiguration

	// All received endpoints, keyed by cluster name
	eds map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment

	// All received endpoints, keyed by secret name
	sds map[string]*envoy_extensions_transport_sockets_tls_v3.Secret

	HandleEDSCDS  func(clusters map[string]*envoy_config_cluster_v3.Cluster)
	HandleCDS     func(clusters map[string]*envoy_config_cluster_v3.Cluster)
	HandleRDS     func(clusters map[string]*envoy_config_route_v3.RouteConfiguration)
	HandleHttpLDS func(listeners map[string]*envoy_config_listener_v3.Listener)
	HandleTcpLDS  func(listeners map[string]*envoy_config_listener_v3.Listener)
	HandleEDS     func(clusters map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment)
	HandleSDS     func(secrets map[string]*envoy_extensions_transport_sockets_tls_v3.Secret)
}

func NewADSC(address string, t *tls.Config, node utils.NodeConfig) *ADSC {
	a := &ADSC{}
	c := a.config()
	c.NodeConfig = node
	a.Client = xds_v3.NewClient(address, t, c)
	return a
}

func (a *ADSC) Run(ctx context.Context) error {
	return a.Client.Run(ctx)
}

func (a *ADSC) config() *xds_v3.Config {
	return &xds_v3.Config{
		OnConnect: a.watch,
		HandleCDS: a.handleCDS,
		HandleRDS: a.handleRDS,
		HandleLDS: a.handleLDS,
		HandleEDS: a.handleEDS,
		HandleSDS: a.handleSDS,
	}
}

func (a *ADSC) watch(cli *xds_v3.Client) error {
	a.watchTime = time.Now()
	if a.watchCDS {
		err := cli.SendRsc(xds_v3.ClusterType, nil)
		if err != nil {
			return err
		}
	} else {
		if a.watchLDS {
			err := cli.SendRsc(xds_v3.ListenerType, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *ADSC) handleCDS(xds *xds_v3.Client, ll []*envoy_config_cluster_v3.Cluster) {
	cn := make([]string, 0, len(ll))
	secrets := []string{}
	edscds := map[string]*envoy_config_cluster_v3.Cluster{}
	cds := map[string]*envoy_config_cluster_v3.Cluster{}
	for _, c := range ll {
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

	if len(cn) > 0 {
		xds.SendRsc(xds_v3.EndpointType, removeDuplicates(cn))
	}
	if len(secrets) > 0 {
		xds.SendRsc(xds_v3.SecretType, removeDuplicates(secrets))
	}
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.edsCDS, edscds) {
		a.edsCDS = edscds
		if a.HandleEDSCDS != nil {
			a.HandleEDSCDS(edscds)
		}
	}
	if !reflect.DeepEqual(a.CDS, cds) {
		a.CDS = cds
		if a.HandleCDS != nil {
			a.HandleCDS(cds)
		}
	}
}

func (a *ADSC) handleEDS(xds *xds_v3.Client, eds []*envoy_config_endpoint_v3.ClusterLoadAssignment) {
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
			xds.SendRsc(xds_v3.ListenerType, nil)
		} else {
			a.initialLoad = time.Since(a.watchTime)
		}
	}
}

func (a *ADSC) handleLDS(xds *xds_v3.Client, ll []*envoy_config_listener_v3.Listener) {
	lh := map[string]*envoy_config_listener_v3.Listener{}
	lt := map[string]*envoy_config_listener_v3.Listener{}

	routes := []string{}
	secrets := []string{}
	for _, l := range ll {

		// The last filter is the actual destination for inbound listener
		if l.ApiListener != nil || len(l.FilterChains) == 0 {
			// This is an API Listener
			// TODO: extract VIP and RDS or cluster
			continue
		}
		filterChain := l.FilterChains[0]
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
			//	RDS = append(RDS, "http_proxy")
			//} else {
			//	RDS = append(RDS, fmt.Sprintf("%d", port))
			//}
		}
	}

	if len(routes) > 0 {
		xds.SendRsc(xds_v3.RouteType, removeDuplicates(routes))
	}
	if len(secrets) > 0 {
		xds.SendRsc(xds_v3.SecretType, removeDuplicates(secrets))
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

func (a *ADSC) handleRDS(xds *xds_v3.Client, configurations []*envoy_config_route_v3.RouteConfiguration) {
	rds := map[string]*envoy_config_route_v3.RouteConfiguration{}
	for _, r := range configurations {
		rds[r.Name] = r
	}

	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !reflect.DeepEqual(a.RDS, rds) {
		a.RDS = rds
		if a.HandleRDS != nil {
			a.HandleRDS(rds)
		}
	}

	if a.initialLoad == 0 {
		a.initialLoad = time.Since(a.watchTime)
	}
}

func (a *ADSC) handleSDS(xds *xds_v3.Client, secrets []*envoy_extensions_transport_sockets_tls_v3.Secret) {
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

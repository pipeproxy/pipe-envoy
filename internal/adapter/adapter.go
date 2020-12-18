package adapter

import (
	"context"
	"io/ioutil"
	"sync/atomic"
	"syscall"
	"time"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/adapter/adsc"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/convert"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe-xds/internal/proxy"
	"github.com/wzshiming/lockfile"
	"github.com/wzshiming/logger"
	meshconfig "istio.io/api/mesh/v1alpha1"
)

const (
	XDS = `unix:./etc/istio/proxy/XDS`
	SDS = `unix:./etc/istio/proxy/SDS`
)

// -c etc/istio/proxy/envoy-rev0.json
// --restart-epoch 0
// --drain-time-s 45
// --parent-shutdown-time-s 60
// --service-cluster istio-ingressgateway
// --service-node router~172.17.0.4~istio-ingressgateway-74df9684c8-mbldh.istio-system~istio-system.svc.cluster.local
// --local-address-ip-version v4 --bootstrap-version 3
// --log-format-prefix-with-location 0 --log-format %Y-%m-%dT%T.%fZ.%l.envoy %n.%v
// -l warning
// --component-log-level misc:error
type Config struct {
	ConfigFile            string
	RestartEpoch          uint32
	DrainTimeS            uint32
	ParentShutdownTimeS   uint32
	ServiceCluster        string
	ServiceNone           string
	LocalAddressIPVersion string
	BootstrapVersion      string
	Bootstrap             *envoy_config_bootstrap_v3.Bootstrap

	BasePath string
}

type Adapter struct {
	ads        *adsc.ADSC
	conf       *config.ConfigCtx
	config     *Config
	meshConfig *meshconfig.MeshConfig
	pid        lockfile.Lockfile
	ready      uint32
}

func NewAdapter(c *Config) (*Adapter, error) {
	return &Adapter{
		config: c,
	}, nil
}

func (a *Adapter) Start(ctx context.Context) error {
	err := a.init(ctx)
	if err != nil {
		return err
	}
	return a.ads.Start(ctx)
}

func (a *Adapter) init(ctx context.Context) error {
	log := logger.FromContext(ctx).WithValues("file", a.config.ConfigFile)

	data, err := ioutil.ReadFile(a.config.ConfigFile)
	if err != nil {
		return err
	}

	bootstrap, err := encoding.UnmarshalBootstrap(data)
	if err != nil {
		return err
	}
	a.config.Bootstrap = bootstrap

	//meshConfig, err := mesh.ApplyMeshConfigDefaults(string(data))
	//if err != nil {
	//	return err
	//}
	//a.meshConfig = meshConfig

	a.conf = config.NewConfigCtx(ctx, a.config.BasePath, time.Second)
	_, err = convert.Convert_config_bootstrap_v3_Bootstrap(a.conf, bootstrap)
	if err != nil {
		return err
	}

	a.conf.Save()

	a.ads = adsc.NewADSC(XDS, SDS, bootstrap.Node)

	a.ads.HandleCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
		for _, cluster := range clusters {
			_, err := convert.Convert_config_cluster_v3_Cluster(a.conf, cluster)
			if err != nil {
				log.Error(err, "cds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleEDSCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
		for _, cluster := range clusters {
			_, err := convert.Convert_config_cluster_v3_Cluster(a.conf, cluster)
			if err != nil {
				log.Error(err, "cds.eds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleEDS = func(endpoints map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment) {
		for _, endpoint := range endpoints {
			_, err := convert.Convert_config_endpoint_v3_ClusterLoadAssignment(a.conf, endpoint)
			if err != nil {
				log.Error(err, "eds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleTcpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
		for _, listener := range listeners {
			_, err := convert.Convert_config_listener_v3_Listener(a.conf, listener, "")
			if err != nil {
				log.Error(err, "tcp.lds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleHttpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
		for _, listener := range listeners {
			_, err := convert.Convert_config_listener_v3_Listener(a.conf, listener, "")
			if err != nil {
				log.Error(err, "http.lds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleRDS = func(routes map[string]*envoy_config_route_v3.RouteConfiguration) {
		for _, route := range routes {
			_, err := convert.Convert_config_route_v3_RouteConfiguration(a.conf, route)
			if err != nil {
				log.Error(err, "rds")
				continue
			}
		}
		a.conf.Update()
	}
	a.ads.HandleSDS = func(secrets map[string]*envoy_extensions_transport_sockets_tls_v3.Secret) {
		for _, secret := range secrets {
			_, err := convert.Convert_extensions_transport_sockets_tls_v3_Secret(a.conf, secret)
			if err != nil {
				log.Error(err, "sds")
				continue
			}
		}
		a.conf.Update()
	}

	a.StartAndUpdatePipe(ctx)
	return nil
}

func (a *Adapter) StartAndUpdatePipe(ctx context.Context) {
	log := logger.FromContext(ctx)
	pipe := proxy.NewPipe(a.config.BasePath)
	go func() {
		for {
			err := pipe.Run(ctx)
			if err != nil {
				log.Error(err, "run pipe")
			}
			time.Sleep(time.Second)
		}
	}()

	a.conf.Watch(
		ctx,
		func() {
			atomic.CompareAndSwapUint32(&a.ready, 0, 1)
			log.Info("load")
			err := pipe.Signal(syscall.SIGHUP)
			if err != nil {
				log.Error(err, "SendSignal")
			}
		},
	)
}

func (a *Adapter) Ready() bool {
	return atomic.LoadUint32(&a.ready) != 0
}

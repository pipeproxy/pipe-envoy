package proxy

import (
	"context"
	"sync/atomic"
	"syscall"
	"time"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/adsc"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/convert"
	"github.com/pipeproxy/pipe-xds/internal/pipe"
	"github.com/wzshiming/lockfile"
	"github.com/wzshiming/logger"
)

type Config struct {
	Node     *envoy_config_core_v3.Node
	XDS      string
	SDS      string
	BasePath string
	Interval time.Duration
}

type Proxy struct {
	ads    *adsc.ADSC
	conf   *config.ConfigCtx
	config *Config
	pid    lockfile.Lockfile
	ready  uint32
}

func NewProxy(c *Config) (*Proxy, error) {
	if c.Interval == 0 {
		c.Interval = time.Second
	}
	return &Proxy{
		conf:   config.NewConfigCtx(c.BasePath, c.Interval),
		ads:    adsc.NewADSC(c.XDS, c.SDS, c.Node),
		config: c,
	}, nil
}

func (a *Proxy) ConfigCtx() *config.ConfigCtx {
	return a.conf
}

func (a *Proxy) ADSC() *adsc.ADSC {
	return a.ads
}

func (a *Proxy) Run(ctx context.Context) error {
	err := a.init(ctx)
	if err != nil {
		return err
	}
	return a.ads.Run(ctx)
}

func (a *Proxy) init(ctx context.Context) error {
	log := logger.FromContext(ctx)

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
			_, err := convert.Convert_config_endpoint_v3_ClusterLoadAssignment(a.conf, endpoint, false)
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

	a.startAndUpdatePipe(ctx)
	return nil
}

func (a *Proxy) startAndUpdatePipe(ctx context.Context) {
	log := logger.FromContext(ctx)
	p := pipe.NewPipe(a.config.BasePath)
	go func() {
		for {
			err := p.Run(ctx)
			if err != nil {
				log.Error(err, "run p")
			}
			time.Sleep(time.Second)
		}
	}()

	a.conf.Watch(
		ctx,
		func() {
			atomic.CompareAndSwapUint32(&a.ready, 0, 1)
			log.Info("load")
			err := p.Signal(syscall.SIGHUP)
			if err != nil {
				log.Error(err, "SendSignal")
			}
		},
	)
}

func (a *Proxy) Ready() bool {
	return atomic.LoadUint32(&a.ready) != 0
}

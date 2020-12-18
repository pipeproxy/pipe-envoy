package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy, tls bind.TLS) (bind.StreamHandler, error) {
	var d bind.StreamDialer

	switch s := c.ClusterSpecifier.(type) {
	case *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_Cluster:
		d = conf.CDS(s.Cluster)
	case *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedClusters:
		d0, err := Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_WeightedCluster(conf, s.WeightedClusters)
		if err != nil {
			return nil, err
		}
		d = d0
	}

	var s bind.StreamHandler
	s = bind.ForwardStreamHandlerConfig{
		Dialer: d,
	}
	if tls != nil {
		s = bind.TLSDownStreamHandlerConfig{
			Handler: s,
			TLS:     tls,
		}
	}
	return s, nil
}

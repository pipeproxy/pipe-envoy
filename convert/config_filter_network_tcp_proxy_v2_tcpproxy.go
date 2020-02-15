package convert

import (
	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_tcp_proxy_v2_TcpProxy(conf *config.ConfigCtx, c *envoy_config_filter_network_tcp_proxy_v2.TcpProxy, tls bind.TLS) (bind.StreamHandler, error) {
	switch c.StatPrefix {
	case "tcp":
		switch s := c.ClusterSpecifier.(type) {
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_Cluster:

			var d bind.StreamHandler = bind.StreamHandlerForwardConfig{
				Dialer: bind.RefStreamDialer(config.XdsName(s.Cluster)),
			}

			if tls != nil {
				d = bind.StreamHandlerTLSDownConfig{
					Handler: d,
					TLS:     tls,
				}
			}

			ref, err := conf.RegisterComponents("", d)
			if err != nil {
				return nil, err
			}

			return bind.RefStreamHandler(ref), nil
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_WeightedClusters:
		}
	}
	logger.Todof("%#v", c)
	return nil, nil
}

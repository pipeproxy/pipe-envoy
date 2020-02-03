package convert

import (
	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_tcp_proxy_v2_TcpProxy(conf *config.ConfigCtx, c *envoy_config_filter_network_tcp_proxy_v2.TcpProxy, tlsName string) (string, error) {
	switch c.StatPrefix {
	case "tcp":
		switch s := c.ClusterSpecifier.(type) {
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_Cluster:
			if tlsName == "" {
				clusterRef, err := config.MarshalRef(config.XdsName(s.Cluster))
				if err != nil {
					return "", err
				}

				d, err := config.MarshalKindStreamHandlerForward(clusterRef)
				if err != nil {
					return "", err
				}

				return conf.RegisterComponents("", d)
			}

			tlsRef, err := config.MarshalRef(tlsName)
			if err != nil {
				return "", err
			}

			clusterRef, err := config.MarshalRef(config.XdsName(s.Cluster))
			if err != nil {
				return "", err
			}

			d, err := config.MarshalKindStreamHandlerForward(clusterRef)
			if err != nil {
				return "", err
			}
			d, err = config.MarshalKindStreamHandlerTlsDown(tlsRef, d)
			if err != nil {
				return "", err
			}

			return conf.RegisterComponents("", d)
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy_WeightedClusters:
		}
	}
	logger.Todof("%#v", c)
	return "", nil
}

package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy, tls bind.TLS) (bind.StreamHandler, error) {
	var d bind.StreamDialer

	switch s := c.ClusterSpecifier.(type) {
	case *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_Cluster:
		d = bind.RefStreamDialerConfig{
			Name: s.Cluster,
		}
	case *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedClusters:
		dialers := []bind.LbStreamDialerWeight{}
		for _, cluster := range s.WeightedClusters.Clusters {
			dialers = append(dialers, bind.LbStreamDialerWeight{
				Weight: uint(cluster.Weight),
				Dialer: bind.RefStreamDialerConfig{
					Name: cluster.Name,
				},
			})
		}

		switch len(dialers) {
		case 0:
			d = bind.NoneStreamDialer{}
		case 1:
			d = dialers[0].Dialer
		default:
			d = bind.LbStreamDialerConfig{
				Policy:  bind.LbStreamDialerLoadBalancePolicyEnumEnumRoundRobin,
				Dialers: dialers,
			}
		}
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

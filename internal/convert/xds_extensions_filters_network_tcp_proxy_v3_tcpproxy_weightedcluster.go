package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_WeightedCluster(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedCluster) (bind.StreamDialer, error) {
	dialers := []bind.LbStreamDialerWeight{}
	for _, cluster := range c.Clusters {
		dialer, err := Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_WeightedCluster_ClusterWeight(conf, cluster)
		if err != nil {
			return nil, err
		}
		dialers = append(dialers, dialer)
	}

	switch len(dialers) {
	case 0:
		return bind.NoneStreamDialer{}, nil
	case 1:
		return dialers[0].Dialer, nil
	default:
		return bind.LbStreamDialerConfig{
			Policy:  bind.RoundRobinBalancePolicy{},
			Dialers: dialers,
		}, nil
	}
}

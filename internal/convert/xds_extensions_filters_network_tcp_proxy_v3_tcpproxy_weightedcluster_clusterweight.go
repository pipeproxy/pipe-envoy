package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_WeightedCluster_ClusterWeight(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedCluster_ClusterWeight) (bind.LbStreamDialerWeight, error) {
	return bind.LbStreamDialerWeight{
		Weight: uint(c.Weight),
		Dialer: conf.CDS(c.Name),
	}, nil
}

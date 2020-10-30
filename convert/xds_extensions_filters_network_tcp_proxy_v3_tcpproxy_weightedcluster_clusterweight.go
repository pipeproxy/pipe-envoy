package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_WeightedCluster_ClusterWeight(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedCluster_ClusterWeight) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_WeightedCluster_ClusterWeight %s\n", string(data))
	return "", nil
}

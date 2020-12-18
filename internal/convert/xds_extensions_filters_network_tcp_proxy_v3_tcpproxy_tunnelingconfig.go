package convert

import (
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy_TunnelingConfig(conf *config.ConfigCtx, c *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_TunnelingConfig) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy_TunnelingConfig %s\n", string(data))
	return "", nil
}

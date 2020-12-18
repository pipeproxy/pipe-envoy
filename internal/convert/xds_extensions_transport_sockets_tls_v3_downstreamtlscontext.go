package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_transport_sockets_tls_v3_DownstreamTlsContext(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.DownstreamTlsContext) (bind.TLS, error) {
	return Convert_extensions_transport_sockets_tls_v3_CommonTlsContext(conf, c.CommonTlsContext)
}

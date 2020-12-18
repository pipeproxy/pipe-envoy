package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_extensions_transport_sockets_tls_v3_CommonTlsContext_CertificateProvider(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_CertificateProvider) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_CertificateProvider %s\n", string(data))
	return "", nil
}

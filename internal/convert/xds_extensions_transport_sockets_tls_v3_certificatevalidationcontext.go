package convert

import (
	"log"

	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_transport_sockets_tls_v3_CertificateValidationContext(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.CertificateValidationContext) (bind.TLS, error) {
	if c.TrustedCa != nil {
		input, err := Convert_config_core_v3_DataSource(conf, c.TrustedCa)
		if err != nil {
			return nil, err
		}
		return bind.ValidationTLSConfig{
			Ca: input,
		}, nil
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_transport_sockets_tls_v3.CertificateValidationContext %s\n", string(data))
	return nil, nil
}

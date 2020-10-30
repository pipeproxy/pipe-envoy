package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_transport_sockets_tls_v3_CertificateValidationContext(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.CertificateValidationContext) (bind.TLS, error) {
	input, err := Convert_config_core_v3_DataSource(conf, c.TrustedCa)
	if err != nil {
		return nil, err
	}

	return bind.ValidationTLSConfig{
		Ca: input,
	}, nil
}

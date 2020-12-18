package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_transport_sockets_tls_v3_CommonTlsContext_CombinedCertificateValidationContext(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_CombinedCertificateValidationContext) (bind.TLS, error) {
	var merge []bind.TLS
	tls, err := Convert_extensions_transport_sockets_tls_v3_SdsSecretConfig(conf, c.ValidationContextSdsSecretConfig)
	if err != nil {
		return nil, err
	}
	if tls != nil {
		merge = append(merge, tls)
	}
	tls, err = Convert_extensions_transport_sockets_tls_v3_CertificateValidationContext(conf, c.DefaultValidationContext)
	if err != nil {
		return nil, err
	}
	if tls != nil {
		merge = append(merge, tls)
	}
	d := bind.MergeTLSConfig{
		Merge: merge,
	}
	return d, nil
}

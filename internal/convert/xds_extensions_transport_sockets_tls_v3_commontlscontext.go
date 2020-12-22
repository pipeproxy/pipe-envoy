package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"

	"log"
)

func Convert_extensions_transport_sockets_tls_v3_CommonTlsContext(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext) (bind.TLS, error) {

	merge := []bind.TLS{}
	for _, tlsCertificateSdsSecretConfig := range c.TlsCertificateSdsSecretConfigs {
		tls, err := Convert_extensions_transport_sockets_tls_v3_SdsSecretConfig(conf, tlsCertificateSdsSecretConfig)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	}

	switch t := c.ValidationContextType.(type) {
	case *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_ValidationContext:
		tls, err := Convert_extensions_transport_sockets_tls_v3_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	case *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_ValidationContextSdsSecretConfig:
		tls, err := Convert_extensions_transport_sockets_tls_v3_SdsSecretConfig(conf, t.ValidationContextSdsSecretConfig)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	case *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_CombinedValidationContext:
		tls, err := Convert_extensions_transport_sockets_tls_v3_CommonTlsContext_CombinedCertificateValidationContext(conf, t.CombinedValidationContext)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
		return nil, nil
	default:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_extensions_transport_sockets_tls_v3.CommonTlsContext %s\n", string(data))
	}

	if len(merge) == 0 {
		return nil, nil
	}
	d := bind.MergeTLSConfig{
		Merge: merge,
	}
	return d, nil
}

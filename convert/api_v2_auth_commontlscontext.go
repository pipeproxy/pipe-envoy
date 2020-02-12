package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CommonTlsContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CommonTlsContext) (bind.TLS, error) {

	merge := []bind.TLS{}
	for _, tlsCertificateSdsSecretConfig := range c.TlsCertificateSdsSecretConfigs {
		tls, err := Convert_api_v2_auth_SdsSecretConfig(conf, tlsCertificateSdsSecretConfig)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	}

	switch t := c.ValidationContextType.(type) {
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContext:
		tls, err := Convert_api_v2_auth_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContextSdsSecretConfig:
		tls, err := Convert_api_v2_auth_SdsSecretConfig(conf, t.ValidationContextSdsSecretConfig)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
	case *envoy_api_v2_auth.CommonTlsContext_CombinedValidationContext:
		tls, err := Convert_api_v2_auth_CommonTlsContext_CombinedCertificateValidationContext(conf, t.CombinedValidationContext)
		if err != nil {
			return nil, err
		}
		if tls != nil {
			merge = append(merge, tls)
		}
		return nil, nil
	}

	ref, err := conf.RegisterComponents("", bind.TLSMergeConfig{Merge: merge})
	if err != nil {
		return nil, err
	}

	return bind.RefTLS(ref), nil

}

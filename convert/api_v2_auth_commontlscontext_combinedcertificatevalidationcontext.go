package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CommonTlsContext_CombinedCertificateValidationContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CommonTlsContext_CombinedCertificateValidationContext) (bind.TLS, error) {
	merge := []bind.TLS{}
	tls, err := Convert_api_v2_auth_SdsSecretConfig(conf, c.ValidationContextSdsSecretConfig)
	if err != nil {
		return nil, err
	}
	if tls != nil {
		merge = append(merge, tls)
	}
	tls, err = Convert_api_v2_auth_CertificateValidationContext(conf, c.DefaultValidationContext)
	if err != nil {
		return nil, err
	}
	if tls != nil {
		merge = append(merge, tls)
	}

	ref, err := conf.RegisterComponents("", bind.TLSMergeConfig{Merge: merge})
	if err != nil {
		return nil, err
	}

	return bind.RefTLS(ref), nil
}

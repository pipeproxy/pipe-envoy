package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CertificateValidationContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CertificateValidationContext) (bind.TLS, error) {
	input, err := Convert_api_v2_core_DataSource(conf, c.TrustedCa)
	if err != nil {
		return nil, err
	}

	ref, err := conf.RegisterComponents("", bind.TLSValidationConfig{
		Ca: input,
	})
	if err != nil {
		return nil, err
	}

	return bind.RefTLS(ref), nil
}

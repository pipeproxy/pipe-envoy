package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CertificateValidationContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CertificateValidationContext) (string, error) {
	name, err := Convert_api_v2_core_DataSource(conf, c.TrustedCa)
	if err != nil {
		return "", err
	}
	caRef, err := config.MarshalRef(name)
	if err != nil {
		return "", err
	}
	d, err := config.MarshalKindTlsValidation(caRef)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)
}

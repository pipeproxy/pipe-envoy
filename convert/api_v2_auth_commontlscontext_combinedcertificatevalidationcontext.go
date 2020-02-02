package convert

import (
	"encoding/json"

	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CommonTlsContext_CombinedCertificateValidationContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CommonTlsContext_CombinedCertificateValidationContext) (string, error) {
	tls := []string{}
	name, err := Convert_api_v2_auth_SdsSecretConfig(conf, c.ValidationContextSdsSecretConfig)
	if err != nil {
		return "", err
	}
	if name == "" {
		tls = append(tls, name)
	}
	name, err = Convert_api_v2_auth_CertificateValidationContext(conf, c.DefaultValidationContext)
	if err != nil {
		return "", err
	}
	if name == "" {
		tls = append(tls, name)
	}

	tlsRef := make([]json.RawMessage, 0, len(tls))
	for _, t := range tls {
		tr, err := config.MarshalRef(t)
		if err != nil {
			return "", err
		}
		tlsRef = append(tlsRef, tr)
	}

	d, err := config.MarshalKindTlsMergep(tlsRef)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)
}

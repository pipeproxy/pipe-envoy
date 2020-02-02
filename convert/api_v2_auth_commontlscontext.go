package convert

import (
	"encoding/json"

	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_CommonTlsContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CommonTlsContext) (string, error) {

	tls := []string{}
	for _, tlsCertificateSdsSecretConfig := range c.TlsCertificateSdsSecretConfigs {
		name, err := Convert_api_v2_auth_SdsSecretConfig(conf, tlsCertificateSdsSecretConfig)
		if err != nil {
			return "", err
		}
		if name != "" {
			tls = append(tls, name)
		}
	}

	switch t := c.ValidationContextType.(type) {
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContext:
		name, err := Convert_api_v2_auth_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return "", err
		}
		if name != "" {
			tls = append(tls, name)
		}
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContextSdsSecretConfig:
		name, err := Convert_api_v2_auth_SdsSecretConfig(conf, t.ValidationContextSdsSecretConfig)
		if err != nil {
			return "", err
		}
		if name != "" {
			tls = append(tls, name)
		}
	case *envoy_api_v2_auth.CommonTlsContext_CombinedValidationContext:
		name, err := Convert_api_v2_auth_CommonTlsContext_CombinedCertificateValidationContext(conf, t.CombinedValidationContext)
		if err != nil {
			return "", err
		}
		if name != "" {
			tls = append(tls, name)
		}
		return "", nil
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

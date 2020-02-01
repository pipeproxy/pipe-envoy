package convert

import (
	"encoding/json"

	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_auth_CommonTlsContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.CommonTlsContext) (string, error) {

	tls := []json.RawMessage{}
	for _, tlsCertificateSdsSecretConfig := range c.TlsCertificateSdsSecretConfigs {
		name, err := Convert_api_v2_core_ConfigSource(conf, tlsCertificateSdsSecretConfig.SdsConfig)
		if err != nil {
			return "", err
		}
		if name != "" && tlsCertificateSdsSecretConfig.Name != "" {
			conf.AppendSDS(tlsCertificateSdsSecretConfig.Name)
			d, err := config.MarshalRef(config.XdsName(tlsCertificateSdsSecretConfig.Name))
			if err != nil {
				return "", err
			}
			tls = append(tls, d)
		}
	}

	switch t := c.ValidationContextType.(type) {
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContext:
		logger.Todof("%#v", c)
		return "", nil
	case *envoy_api_v2_auth.CommonTlsContext_ValidationContextSdsSecretConfig:
		name, err := Convert_api_v2_core_ConfigSource(conf, t.ValidationContextSdsSecretConfig.SdsConfig)
		if err != nil {
			return "", err
		}

		if name != "" && t.ValidationContextSdsSecretConfig.Name != "" {
			conf.AppendSDS(t.ValidationContextSdsSecretConfig.Name)
			d, err := config.MarshalRef(config.XdsName(t.ValidationContextSdsSecretConfig.Name))
			if err != nil {
				return "", err
			}
			tls = append(tls, d)
		}
	case *envoy_api_v2_auth.CommonTlsContext_CombinedValidationContext:
		logger.Todof("%#v", c)
		return "", nil
	}

	d, err := config.MarshalKindTlsMergep(tls)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)

}

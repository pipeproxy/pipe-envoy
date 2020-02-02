package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_SdsSecretConfig(conf *config.ConfigCtx, c *envoy_api_v2_auth.SdsSecretConfig) (string, error) {
	name, err := Convert_api_v2_core_ConfigSource(conf, c.SdsConfig)
	if err != nil {
		return "", err
	}

	if name != "" && c.Name != "" {
		conf.AppendSDS(c.Name)
		d, err := config.MarshalRef(config.XdsName(c.Name))
		if err != nil {
			return "", err
		}
		return conf.RegisterComponents("", d)
	}

	return "", nil
}

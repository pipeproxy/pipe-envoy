package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_SdsSecretConfig(conf *config.ConfigCtx, c *envoy_api_v2_auth.SdsSecretConfig) (bind.TLS, error) {
	name, err := Convert_api_v2_core_ConfigSource(conf, c.SdsConfig)
	if err != nil {
		return nil, err
	}

	if name != "" && c.Name != "" {
		conf.AppendSDS(c.Name)
		ref, err := conf.RegisterComponents("", bind.RefTLS(config.XdsName(c.Name)))
		if err != nil {
			return nil, err
		}
		return bind.RefTLS(ref), nil
	}

	return nil, nil
}

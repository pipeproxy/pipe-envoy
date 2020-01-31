package convert

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_bootstrap_v2_Bootstrap(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap) (string, error) {

	if c.Node != nil {
		_, err := Convert_api_v2_core_Node(conf, c.Node)
		if err != nil {
			return "", err
		}
	}

	if c.Admin != nil {
		_, err := Convert_config_bootstrap_v2_Admin(conf, c.Admin)
		if err != nil {
			return "", err
		}
	}

	if c.StaticResources != nil {
		_, err := Convert_config_bootstrap_v2_Bootstrap_StaticResources(conf, c.StaticResources)
		if err != nil {
			return "", err
		}
	}

	if c.DynamicResources != nil {
		_, err := Convert_config_bootstrap_v2_Bootstrap_DynamicResources(conf, c.DynamicResources)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

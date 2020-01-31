package convert_config_bootstrap_v2

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_core "github.com/wzshiming/envoy/convert/api/v2/core"
)

func Convert_Bootstrap(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap) (string, error) {

	if c.Node != nil {
		_, err := convert_api_v2_core.Convert_Node(conf, c.Node)
		if err != nil {
			return "", err
		}
	}

	if c.Admin != nil {
		_, err := Convert_Admin(conf, c.Admin)
		if err != nil {
			return "", err
		}
	}

	if c.StaticResources != nil {
		_, err := Convert_Bootstrap_StaticResources(conf, c.StaticResources)
		if err != nil {
			return "", err
		}
	}

	if c.DynamicResources != nil {
		_, err := Convert_Bootstrap_DynamicResources(conf, c.DynamicResources)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

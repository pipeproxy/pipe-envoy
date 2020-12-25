package convert

import (
	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
)

func Convert_config_bootstrap_v3_Bootstrap(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.Bootstrap) (string, error) {

	if c.Admin != nil {
		_, err := Convert_config_bootstrap_v3_Admin(conf, c.Admin)
		if err != nil {
			return "", err
		}
	}

	if c.StaticResources != nil {
		_, err := Convert_config_bootstrap_v3_Bootstrap_StaticResources(conf, c.StaticResources)
		if err != nil {
			return "", err
		}
	}

	//if c.DynamicResources != nil {
	//	_, err := Convert_config_bootstrap_v3_Bootstrap_DynamicResources(conf, c.DynamicResources)
	//	if err != nil {
	//		return "", err
	//	}
	//}

	return "", nil
}

package convert

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_bootstrap_v2_Bootstrap_StaticResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap_StaticResources) (string, error) {
	for _, cluster := range c.Clusters {
		_, err := Convert_api_v2_Cluster(conf, cluster)
		if err != nil {
			return "", err
		}
	}

	for _, listener := range c.Listeners {
		_, err := Convert_api_v2_Listener(conf, listener)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

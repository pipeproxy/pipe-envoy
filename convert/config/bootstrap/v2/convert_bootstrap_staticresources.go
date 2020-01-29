package convert_config_bootstrap_v2

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2 "github.com/wzshiming/envoy/convert/api/v2"
)

func Convert_Bootstrap_StaticResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap_StaticResources) (string, error) {
	for _, cluster := range c.Clusters {
		_, err := convert_api_v2.Convert_Cluster(conf, cluster)
		if err != nil {
			return "", err
		}
	}

	for _, listener := range c.Listeners {
		_, err := convert_api_v2.Convert_Listener(conf, listener)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

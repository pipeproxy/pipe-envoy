package convert

import (
	"fmt"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
)

func Convert_config_bootstrap_v3_Bootstrap_StaticResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.Bootstrap_StaticResources) (string, error) {
	for _, cluster := range c.Clusters {
		_, err := Convert_config_cluster_v3_Cluster(conf, cluster)
		if err != nil {
			return "", err
		}
	}

	for i, listener := range c.Listeners {
		_, err := Convert_config_listener_v3_Listener(conf, listener, fmt.Sprintf("%s.%d", "static", i))
		if err != nil {
			return "", err
		}
	}

	return "", nil
}

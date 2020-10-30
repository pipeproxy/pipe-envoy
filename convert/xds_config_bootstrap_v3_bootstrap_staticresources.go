package convert

import (
	"fmt"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/config"
)

func Convert_config_bootstrap_v3_Bootstrap_StaticResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.Bootstrap_StaticResources) (string, error) {
	for i, cluster := range c.Clusters {
		cds, err := Convert_config_cluster_v3_Cluster(conf, cluster)
		if err != nil {
			return "", err
		}
		conf.RegisterCDS(fmt.Sprintf("static.cluster.%d", i), cds, cluster)
	}

	for i, listener := range c.Listeners {
		lds, err := Convert_config_listener_v3_Listener(conf, listener)
		if err != nil {
			return "", err
		}
		conf.RegisterLDS(fmt.Sprintf("static.listener.%d", i), lds, listener)
	}

	return "", nil
}

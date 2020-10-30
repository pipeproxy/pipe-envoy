package convert

import (
	"log"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_config_cluster_v3_Cluster_CommonLbConfig_LocalityWeightedLbConfig(conf *config.ConfigCtx, c *envoy_config_cluster_v3.Cluster_CommonLbConfig_LocalityWeightedLbConfig) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_cluster_v3.Cluster_CommonLbConfig_LocalityWeightedLbConfig %s\n", string(data))
	return "", nil
}

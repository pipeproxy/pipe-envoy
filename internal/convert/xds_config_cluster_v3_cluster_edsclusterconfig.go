package convert

import (
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_cluster_v3_Cluster_EdsClusterConfig(conf *config.ConfigCtx, c *envoy_config_cluster_v3.Cluster_EdsClusterConfig) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_cluster_v3.Cluster_EdsClusterConfig %s\n", string(data))
	return "", nil
}

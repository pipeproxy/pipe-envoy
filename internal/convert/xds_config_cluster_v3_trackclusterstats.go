package convert

import (
	"log"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_cluster_v3_TrackClusterStats(conf *config.ConfigCtx, c *envoy_config_cluster_v3.TrackClusterStats) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_cluster_v3.TrackClusterStats %s\n", string(data))
	return "", nil
}

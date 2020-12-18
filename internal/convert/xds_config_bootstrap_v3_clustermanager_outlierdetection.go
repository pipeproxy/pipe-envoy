package convert

import (
	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_bootstrap_v3_ClusterManager_OutlierDetection(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.ClusterManager_OutlierDetection) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_bootstrap_v3.ClusterManager_OutlierDetection %s\n", string(data))
	return "", nil
}

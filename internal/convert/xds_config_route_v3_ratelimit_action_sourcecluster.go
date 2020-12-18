package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_route_v3_RateLimit_Action_SourceCluster(conf *config.ConfigCtx, c *envoy_config_route_v3.RateLimit_Action_SourceCluster) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RateLimit_Action_SourceCluster %s\n", string(data))
	return "", nil
}

package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_route_v3_RateLimit_Action_HeaderValueMatch(conf *config.ConfigCtx, c *envoy_config_route_v3.RateLimit_Action_HeaderValueMatch) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RateLimit_Action_HeaderValueMatch %s\n", string(data))
	return "", nil
}

package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RedirectAction(conf *config.ConfigCtx, c *envoy_config_route_v3.RedirectAction) (bind.HTTPHandler, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RedirectAction %s\n", string(data))
	return nil, nil
}

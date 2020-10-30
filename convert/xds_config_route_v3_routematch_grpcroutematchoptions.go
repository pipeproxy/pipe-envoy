package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_config_route_v3_RouteMatch_GrpcRouteMatchOptions(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteMatch_GrpcRouteMatchOptions) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RouteMatch_GrpcRouteMatchOptions %s\n", string(data))
	return "", nil
}

package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_Route(conf *config.ConfigCtx, c *envoy_config_route_v3.Route) (bind.HTTPHandler, error) {
	switch a := c.Action.(type) {
	case *envoy_config_route_v3.Route_Route:
		handler, err := Convert_config_route_v3_RouteAction(conf, a.Route)
		if err != nil {
			return nil, err
		}
		return handler, nil
	case *envoy_config_route_v3.Route_Redirect:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return nil, nil
	case *envoy_config_route_v3.Route_DirectResponse:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return nil, nil
	case *envoy_config_route_v3.Route_FilterAction:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return nil, nil
	}
	return nil, nil
}

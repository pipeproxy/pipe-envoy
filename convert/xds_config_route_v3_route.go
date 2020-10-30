package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_Route(conf *config.ConfigCtx, c *envoy_config_route_v3.Route) (bind.MuxNetHTTPHandlerRoute, error) {
	r := bind.MuxNetHTTPHandlerRoute{}
	switch p := c.Match.PathSpecifier.(type) {
	case *envoy_config_route_v3.RouteMatch_Prefix:
		r.Prefix = p.Prefix
	case *envoy_config_route_v3.RouteMatch_Path:
		r.Path = p.Path
	case *envoy_config_route_v3.RouteMatch_SafeRegex:
		r.Regexp = p.SafeRegex.Regex
	}

	var handler bind.HTTPHandler
	switch a := c.Action.(type) {
	case *envoy_config_route_v3.Route_Route:
		handle, err := Convert_config_route_v3_RouteAction(conf, a.Route)
		if err != nil {
			return r, err
		}
		handler = handle
	case *envoy_config_route_v3.Route_Redirect:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return r, nil
	case *envoy_config_route_v3.Route_DirectResponse:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return r, nil
	case *envoy_config_route_v3.Route_FilterAction:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.Route %s\n", string(data))
		return r, nil
	}

	r.Handler = handler

	return r, nil
}

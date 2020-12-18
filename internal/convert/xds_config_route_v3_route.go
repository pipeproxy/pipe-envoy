package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_Route(conf *config.ConfigCtx, c *envoy_config_route_v3.Route) (bind.HTTPHandler, error) {
	switch a := c.Action.(type) {
	case *envoy_config_route_v3.Route_Route:
		return Convert_config_route_v3_RouteAction(conf, a.Route)
	case *envoy_config_route_v3.Route_Redirect:
		return Convert_config_route_v3_RedirectAction(conf, a.Redirect)
	case *envoy_config_route_v3.Route_DirectResponse:
		return Convert_config_route_v3_DirectResponseAction(conf, a.DirectResponse)
	case *envoy_config_route_v3.Route_FilterAction:
		return Convert_config_route_v3_FilterAction(conf, a.FilterAction)
	}
	return nil, nil
}

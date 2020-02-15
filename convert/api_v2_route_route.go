package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_route_Route(conf *config.ConfigCtx, c *envoy_api_v2_route.Route) (bind.HTTPHandlerMuxRoute, string, error) {
	r := bind.HTTPHandlerMuxRoute{}
	switch p := c.Match.PathSpecifier.(type) {
	case *envoy_api_v2_route.RouteMatch_Prefix:
		r.Prefix = p.Prefix
	case *envoy_api_v2_route.RouteMatch_Path:
		r.Path = p.Path
	case *envoy_api_v2_route.RouteMatch_Regex:
		r.Regexp = p.Regex
	case *envoy_api_v2_route.RouteMatch_SafeRegex:
		logger.Todof("%#v", c)
		return r, "", nil
	}

	var handler bind.HTTPHandler
	switch a := c.Action.(type) {
	case *envoy_api_v2_route.Route_Route:
		handle, err := Convert_api_v2_route_RouteAction(conf, a.Route)
		if err != nil {
			return r, "", err
		}
		handler = handle
	case *envoy_api_v2_route.Route_Redirect:
		logger.Todof("%#v", c)
		return r, "", nil
	case *envoy_api_v2_route.Route_DirectResponse:
		handle, err := Convert_api_v2_route_DirectResponseAction(conf, a.DirectResponse)
		if err != nil {
			return r, "", err
		}
		handler = handle
	case *envoy_api_v2_route.Route_FilterAction:
		logger.Todof("%#v", c)
		return r, "", nil
	}

	r.Handler = handler
	return r, "", nil
}

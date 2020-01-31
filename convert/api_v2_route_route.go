package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_route_Route(conf *config.ConfigCtx, c *envoy_api_v2_route.Route) (*config.Route, string, error) {
	r := &config.Route{}
	switch p := c.Match.PathSpecifier.(type) {
	case *envoy_api_v2_route.RouteMatch_Prefix:
		r.Prefix = p.Prefix
	case *envoy_api_v2_route.RouteMatch_Path:
		r.Path = p.Path
	case *envoy_api_v2_route.RouteMatch_Regex:
		r.Regexp = p.Regex
	case *envoy_api_v2_route.RouteMatch_SafeRegex:
		logger.Todof("%#v", c)
		return nil, "", nil
	}

	name := ""
	switch a := c.Action.(type) {
	case *envoy_api_v2_route.Route_Route:
		name0, err := Convert_api_v2_route_RouteAction(conf, a.Route)
		if err != nil {
			return nil, "", err
		}
		name = name0
	case *envoy_api_v2_route.Route_Redirect:
		logger.Todof("%#v", c)
		return nil, "", nil
	case *envoy_api_v2_route.Route_DirectResponse:
		logger.Todof("%#v", c)
		return nil, "", nil
	case *envoy_api_v2_route.Route_FilterAction:
		logger.Todof("%#v", c)
		return nil, "", nil
	}

	d, err := config.MarshalRef(name)
	if err != nil {
		return nil, "", err
	}
	r.Handler = d
	return r, "", nil
}

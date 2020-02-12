package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_route_VirtualHost(conf *config.ConfigCtx, c *envoy_api_v2_route.VirtualHost) (bind.HttpHandler, error) {
	rs := []bind.HttpHandlerMuxRoute{}
	for _, route := range c.Routes {
		r, _, err := Convert_api_v2_route_Route(conf, route)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}

	d := bind.HttpHandlerMuxConfig{
		Routes:   rs,
		NotFound: nil,
	}

	name := config.XdsName(c.Name + ".virtual-host")
	ref, err := conf.RegisterComponents(name, d)
	if err != nil {
		return nil, err
	}

	return bind.RefHttpHandler(ref), nil
}

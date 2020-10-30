package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_VirtualHost(conf *config.ConfigCtx, c *envoy_config_route_v3.VirtualHost) (bind.HTTPHandler, error) {
	rs := []bind.MuxNetHTTPHandlerRoute{}
	for _, route := range c.Routes {
		r, err := Convert_config_route_v3_Route(conf, route)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}

	var d bind.HTTPHandler
	d = bind.MuxNetHTTPHandlerConfig{
		Routes: rs,
	}

	//if c.Name != "" {
	//	d = bind.DefNetHTTPHandlerConfig{
	//		Name: c.Name,
	//		Def:  d,
	//	}
	//}
	return d, nil
}

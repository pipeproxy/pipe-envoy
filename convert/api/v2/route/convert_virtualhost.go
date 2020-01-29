package convert_api_v2_route

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
)

func Convert_VirtualHost(conf *config.ConfigCtx, c *envoy_api_v2_route.VirtualHost) (string, error) {
	rs := []*config.Route{}
	for _, route := range c.Routes {
		r, _, err := Convert_Route(conf, route)
		if err != nil {
			return "", err
		}
		rs = append(rs, r)
	}
	d, err := config.MarshalKindHttpHandlerMux(rs, nil)
	if err != nil {
		return "", err
	}
	name := config.XdsName(c.Name + ".virtual-host")
	return conf.RegisterComponents(name, d)
}

package convert

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_RouteConfiguration(conf *config.ConfigCtx, c *envoy_api_v2.RouteConfiguration) (bind.HTTPHandler, error) {
	handlers := []bind.HTTPHandler{}
	for _, virtualHost := range c.VirtualHosts {
		handler, err := Convert_api_v2_route_VirtualHost(conf, virtualHost)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	d := bind.HTTPHandlerPollerConfig{
		Poller:   "round_robin",
		Handlers: handlers,
	}

	name := config.XdsName(c.Name)

	ref, err := conf.RegisterComponents(name, d)
	if err != nil {
		return nil, err
	}

	return bind.RefHTTPHandler(ref), nil
}

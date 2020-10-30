package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteConfiguration(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteConfiguration) (bind.HTTPHandler, error) {
	hosts := []bind.HostsNetHTTPHandlerRoute{}
	for _, virtualHost := range c.VirtualHosts {
		handler, err := Convert_config_route_v3_VirtualHost(conf, virtualHost)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, bind.HostsNetHTTPHandlerRoute{
			Domains: virtualHost.Domains,
			Handler: handler,
		})
	}
	var d bind.HTTPHandler
	if len(hosts) == 0 {
		d = bind.DirectNetHTTPHandlerConfig{
			Code: 404,
			Body: bind.InlineIoReaderConfig{
				Data: "404 not found",
			},
		}
	} else {
		d = bind.HostsNetHTTPHandlerConfig{
			Hosts: hosts,
		}
	}

	if c.Name != "" {
		d = bind.DefNetHTTPHandlerConfig{
			Name: c.Name,
			Def:  d,
		}
	}
	return d, nil
}

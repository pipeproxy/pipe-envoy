package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteConfiguration(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteConfiguration) (bind.HTTPHandler, error) {
	hosts := []bind.HostNetHTTPHandlerRoute{}
	for _, virtualHost := range c.VirtualHosts {
		handler, err := Convert_config_route_v3_VirtualHost(conf, c.Name, virtualHost)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, bind.HostNetHTTPHandlerRoute{
			Domains: virtualHost.Domains,
			Handler: handler,
		})
	}
	var handers []bind.HTTPHandler

	reqHeader := bind.EditRequestHeaderNetHTTPHandlerConfig{
		Del: c.RequestHeadersToRemove,
	}
	for _, add := range c.RequestHeadersToAdd {
		if add.Append.GetValue() {
			reqHeader.Add = append(reqHeader.Add, bind.EditRequestHeaderNetHTTPHandlerPair{
				Key:   add.Header.GetKey(),
				Value: add.Header.GetValue(),
			})
		} else {
			reqHeader.Set = append(reqHeader.Set, bind.EditRequestHeaderNetHTTPHandlerPair{
				Key:   add.Header.GetKey(),
				Value: add.Header.GetValue(),
			})
		}
	}
	if len(reqHeader.Del) != 0 || len(reqHeader.Set) != 0 || len(reqHeader.Add) != 0 {
		handers = append(handers, reqHeader)
	}

	if len(hosts) == 0 {
		handers = append(handers, bind.DirectNetHTTPHandlerConfig{
			Code: 404,
			Body: bind.InlineIoReaderConfig{
				Data: "404 not found",
			},
		})
	} else {
		handers = append(handers, bind.HostNetHTTPHandlerConfig{
			Hosts: hosts,
		})
	}

	respHeader := bind.EditResponseHeaderNetHTTPHandlerConfig{
		Del: c.ResponseHeadersToRemove,
	}
	for _, add := range c.ResponseHeadersToAdd {
		if add.Append.GetValue() {
			respHeader.Add = append(respHeader.Add, bind.EditResponseHeaderNetHTTPHandlerPair{
				Key:   add.Header.GetKey(),
				Value: add.Header.GetValue(),
			})
		} else {
			respHeader.Set = append(respHeader.Set, bind.EditResponseHeaderNetHTTPHandlerPair{
				Key:   add.Header.GetKey(),
				Value: add.Header.GetValue(),
			})
		}
	}
	if len(respHeader.Del) != 0 || len(respHeader.Set) != 0 || len(respHeader.Add) != 0 {
		handers = append(handers, respHeader)
	}

	var d bind.HTTPHandler

	if len(handers) == 1 {
		d = handers[0]
	} else {
		d = bind.MultiNetHTTPHandlerConfig{
			Multi: handers,
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

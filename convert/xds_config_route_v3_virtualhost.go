package convert

import (
	"fmt"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_VirtualHost(conf *config.ConfigCtx, name string, c *envoy_config_route_v3.VirtualHost) (bind.HTTPHandler, error) {

	var h bind.HTTPHandler
	for i := range c.Routes {
		index := len(c.Routes) - 1 - i
		route := c.Routes[index]
		handler, err := Convert_config_route_v3_Route(conf, route)
		if err != nil {
			return nil, err
		}
		name := fmt.Sprintf("%s|%s|%s|%d", name, c.Name, route.Name, index)
		routeName := name + ".route"
		conf.RegisterRDS("rds.http."+routeName, bind.DefNetHTTPHandlerConfig{
			Name: routeName,
			Def:  handler,
		}, nil)
		handler = bind.RefNetHTTPHandlerConfig{
			Name: routeName,
		}

		query := Convert_config_route_v3_RouteMatch_Query(conf, route.Match)
		if query != nil {
			routeName = name + ".match.query"
			conf.RegisterRDS("rds.http."+routeName, bind.DefNetHTTPHandlerConfig{
				Name: routeName,
				Def: bind.QueryNetHTTPHandlerConfig{
					Queries: []bind.QueryNetHTTPHandlerRoute{
						{
							Matches: query,
							Handler: handler,
						},
					},
					NotFound: h,
				},
			}, nil)
			handler = bind.RefNetHTTPHandlerConfig{
				Name: routeName,
			}
		}
		header := Convert_config_route_v3_RouteMatch_Header(conf, route.Match)
		if header != nil {
			routeName = name + ".match.header"
			conf.RegisterRDS("rds.http."+routeName, bind.DefNetHTTPHandlerConfig{
				Name: routeName,
				Def: bind.HeaderNetHTTPHandlerConfig{
					Headers: []bind.HeaderNetHTTPHandlerRoute{
						{
							Matches: header,
							Handler: handler,
						},
					},
					NotFound: h,
				},
			}, nil)
			handler = bind.RefNetHTTPHandlerConfig{
				Name: routeName,
			}
		}

		routeName = name + ".match.path"
		path := Convert_config_route_v3_RouteMatch_Path(conf, route.Match, handler)
		paths := bind.PathNetHTTPHandlerConfig{
			Paths: []bind.PathNetHTTPHandlerRoute{
				path,
			},
			NotFound: h,
		}
		conf.RegisterRDS("rds.http."+routeName, bind.DefNetHTTPHandlerConfig{
			Name: routeName,
			Def:  paths,
		}, nil)
		handler = bind.RefNetHTTPHandlerConfig{
			Name: routeName,
		}
		h = handler
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

	handers = append(handers, h)

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

	//if c.Name != "" {
	//	d = bind.DefNetHTTPHandlerConfig{
	//		Name: c.Name,
	//		Def:  d,
	//	}
	//}
	return d, nil
}

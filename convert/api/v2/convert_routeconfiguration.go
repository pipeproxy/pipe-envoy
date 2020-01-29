package convert_api_v2

import (
	"encoding/json"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_route "github.com/wzshiming/envoy/convert/api/v2/route"
)

func Convert_RouteConfiguration(conf *config.ConfigCtx, c *envoy_api_v2.RouteConfiguration) (string, error) {
	list := []json.RawMessage{}
	for _, virtualHost := range c.VirtualHosts {
		name, err := convert_api_v2_route.Convert_VirtualHost(conf, virtualHost)
		if err != nil {
			return "", err
		}
		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		list = append(list, ref)
	}

	d, err := config.MarshalKindHttpHandlerPoller("round_robin", list)
	if err != nil {
		return "", err
	}

	name := config.XdsName(c.Name)

	return conf.RegisterComponents(name, d)
}

package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_route_DirectResponseAction(conf *config.ConfigCtx, c *envoy_api_v2_route.DirectResponseAction) (string, error) {
	body, err := Convert_api_v2_core_DataSource(conf, c.Body)
	if err != nil {
		return "", err
	}

	bodyRef, err := config.MarshalRef(body)
	if err != nil {
		return "", err
	}

	d, err := config.MarshalKindHttpHandlerDirect(int(c.Status), bodyRef)
	if err != nil {
		return "", err
	}

	return conf.RegisterComponents("", d)
}

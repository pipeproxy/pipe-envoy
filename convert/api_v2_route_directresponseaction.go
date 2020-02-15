package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_route_DirectResponseAction(conf *config.ConfigCtx, c *envoy_api_v2_route.DirectResponseAction) (bind.HTTPHandler, error) {
	body, err := Convert_api_v2_core_DataSource(conf, c.Body)
	if err != nil {
		return nil, err
	}

	d := bind.HTTPHandlerDirectConfig{
		Code: int(c.Status),
		Body: body,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	return bind.RefHTTPHandler(ref), nil
}

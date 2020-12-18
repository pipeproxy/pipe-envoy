package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_DirectResponseAction(conf *config.ConfigCtx, c *envoy_config_route_v3.DirectResponseAction) (bind.HTTPHandler, error) {
	b, err := Convert_config_core_v3_DataSource(conf, c.Body)
	if err != nil {
		return nil, err
	}
	return bind.DirectNetHTTPHandlerConfig{
		Code: int(c.Status),
		Body: b,
	}, nil
}

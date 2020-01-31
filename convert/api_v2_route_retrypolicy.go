package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_route_RetryPolicy(conf *config.ConfigCtx, c *envoy_api_v2_route.RetryPolicy) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

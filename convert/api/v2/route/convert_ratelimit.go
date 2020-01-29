package convert_api_v2_route

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_RateLimit(conf *config.ConfigCtx, c *envoy_api_v2_route.RateLimit) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

package convert_api_v2

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_ScopedRouteConfiguration_Key_Fragment_StringKey(conf *config.ConfigCtx, c *envoy_api_v2.ScopedRouteConfiguration_Key_Fragment_StringKey) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

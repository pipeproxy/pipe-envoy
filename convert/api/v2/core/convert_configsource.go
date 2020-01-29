package convert_api_v2_core

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_ConfigSource(conf *config.ConfigCtx, c *envoy_api_v2_core.ConfigSource) (string, error) {

	logger.Todof("%#v", c)
	return "", nil
}

package convert

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_Listener_ConnectionBalanceConfig(conf *config.ConfigCtx, c *envoy_api_v2.Listener_ConnectionBalanceConfig) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

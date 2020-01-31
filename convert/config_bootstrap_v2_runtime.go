package convert

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_bootstrap_v2_Runtime(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Runtime) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

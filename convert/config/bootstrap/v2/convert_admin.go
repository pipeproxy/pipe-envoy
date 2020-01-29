package convert_config_bootstrap_v2

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_Admin(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Admin) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

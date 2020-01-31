package convert

import (
	envoy_config_filter_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_accesslog_v2_HeaderFilter(conf *config.ConfigCtx, c *envoy_config_filter_accesslog_v2.HeaderFilter) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

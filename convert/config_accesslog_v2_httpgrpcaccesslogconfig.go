package convert

import (
	envoy_config_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_accesslog_v2_HttpGrpcAccessLogConfig(conf *config.ConfigCtx, c *envoy_config_accesslog_v2.HttpGrpcAccessLogConfig) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

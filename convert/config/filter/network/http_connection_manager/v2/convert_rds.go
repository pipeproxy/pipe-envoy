package convert_config_filter_network_http_connection_manager_v2

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_Rds(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.Rds) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

package convert

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_filter_network_http_connection_manager_v2_Rds(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.Rds) (string, error) {
	name := c.RouteConfigName
	if name != "" {
		conf.AppendRDS(name)
		return config.XdsName(name), nil
	}
	return Convert_api_v2_core_ConfigSource(conf, c.ConfigSource)
}

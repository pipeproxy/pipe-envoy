package convert_config_filter_network_http_connection_manager_v2

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_HttpConnectionManager(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager) (string, error) {
	name := ""
	switch r := c.RouteSpecifier.(type) {
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_Rds:
		name = r.Rds.RouteConfigName
		if name != "" {
			conf.AppendRDS(name)
		}

	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_RouteConfig:
		logger.Todof("%#v", c)
		return "", nil
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_ScopedRoutes:
		logger.Todof("%#v", c)
		return "", nil
	}

	routeName := config.XdsName(name)

	ref, err := config.MarshalRef(routeName)
	if err != nil {
		return "", err
	}

	d, err := config.MarshalKindStreamHandlerHTTP(ref, nil)
	if err != nil {
		return "", err
	}

	name = config.XdsName(name + ".route")
	return conf.RegisterComponents(name, d)
}

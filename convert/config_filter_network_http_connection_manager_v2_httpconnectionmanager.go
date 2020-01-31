package convert

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_http_connection_manager_v2_HttpConnectionManager(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager) (string, error) {
	routeName := ""
	switch r := c.RouteSpecifier.(type) {
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_Rds:
		name, err := Convert_config_filter_network_http_connection_manager_v2_Rds(conf, r.Rds)
		if err != nil {
			return "", err
		}
		routeName = name
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_RouteConfig:
		name, err := Convert_api_v2_RouteConfiguration(conf, r.RouteConfig)
		if err != nil {
			return "", err
		}
		routeName = name
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_ScopedRoutes:
		logger.Todof("%#v", c)
		return "", nil
	}

	ref, err := config.MarshalRef(routeName)
	if err != nil {
		return "", err
	}

	for _, accessLog := range c.AccessLog {
		name, err := Convert_config_filter_accesslog_v2_AccessLog(conf, accessLog, ref)
		if err != nil {
			return "", err
		}
		ref, err = config.MarshalRef(name)
		if err != nil {
			return "", err
		}
	}

	d, err := config.MarshalKindStreamHandlerHTTP(ref, nil)
	if err != nil {
		return "", err
	}

	name := routeName + ".route"
	return conf.RegisterComponents(name, d)
}

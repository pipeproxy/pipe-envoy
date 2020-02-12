package convert

import (
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_config_filter_network_http_connection_manager_v2_HttpConnectionManager(conf *config.ConfigCtx, c *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager, tls bind.TLS) (bind.StreamHandler, error) {
	var route bind.HttpHandler
	switch r := c.RouteSpecifier.(type) {
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_Rds:
		handler, err := Convert_config_filter_network_http_connection_manager_v2_Rds(conf, r.Rds)
		if err != nil {
			return nil, err
		}
		route = handler
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_RouteConfig:
		handler, err := Convert_api_v2_RouteConfiguration(conf, r.RouteConfig)
		if err != nil {
			return nil, err
		}
		route = handler
	case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager_ScopedRoutes:
		logger.Todof("%#v", c)
		return nil, nil
	}

	for _, accessLog := range c.AccessLog {
		n, err := Convert_config_filter_accesslog_v2_AccessLog(conf, accessLog, route)
		if err != nil {
			return nil, err
		}
		route = n
	}

	d := bind.StreamHandlerHttpConfig{
		Handler: route,
		TLS:     tls,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}
	return bind.RefStreamHandler(ref), nil
}

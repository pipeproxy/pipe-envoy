package convert

import (
	"fmt"

	"github.com/wzshiming/envoy/internal/logger"

	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"

	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"

	envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/golang/protobuf/proto"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_listener_Filter(conf *config.ConfigCtx, c *envoy_api_v2_listener.Filter) (string, error) {

	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_api_v2_listener.Filter_TypedConfig:
		msg, err := config.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return "", err
		}
		filterConfig = msg
	case *envoy_api_v2_listener.Filter_Config:
		return "", fmt.Errorf("not suppert envoy_api_v2_listener.Filter_Config")
	}

	switch c.Name {
	case "envoy.http_connection_manager":
		switch p := filterConfig.(type) {
		case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager:
			return Convert_config_filter_network_http_connection_manager_v2_HttpConnectionManager(conf, p)
		}
	case "envoy.tcp_proxy":
		switch p := filterConfig.(type) {
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy:
			return Convert_config_filter_network_tcp_proxy_v2_TcpProxy(conf, p)
		}

	}
	logger.Todof("%#v", c)
	return "", nil
}

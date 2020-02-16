package convert

import (
	"fmt"

	envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	envoy_config_filter_network_http_connection_manager_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	envoy_config_filter_network_tcp_proxy_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/tcp_proxy/v2"
	"github.com/golang/protobuf/proto"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
	"github.com/wzshiming/envoy/wellknown"
)

func Convert_api_v2_listener_Filter(conf *config.ConfigCtx, c *envoy_api_v2_listener.Filter, tls bind.TLS) (bind.StreamHandler, error) {

	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_api_v2_listener.Filter_TypedConfig:
		msg, err := config.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	case *envoy_api_v2_listener.Filter_Config:
		return nil, fmt.Errorf("not suppert envoy_api_v2_listener.Filter_Config")
	}

	switch c.Name {
	case wellknown.HTTPConnectionManager, wellknown.HTTPConnectionManagerAlias:
		switch p := filterConfig.(type) {
		case *envoy_config_filter_network_http_connection_manager_v2.HttpConnectionManager:
			return Convert_config_filter_network_http_connection_manager_v2_HttpConnectionManager(conf, p, tls)
		}
	case wellknown.TCPProxy, wellknown.TCPProxyAlias:
		switch p := filterConfig.(type) {
		case *envoy_config_filter_network_tcp_proxy_v2.TcpProxy:
			return Convert_config_filter_network_tcp_proxy_v2_TcpProxy(conf, p, tls)
		}
	}
	logger.Todof("%#v", c)
	return nil, nil
}

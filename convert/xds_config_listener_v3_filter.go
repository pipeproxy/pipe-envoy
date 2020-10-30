package convert

import (
	"log"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_filters_network_tcp_proxy_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/proto"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_Filter(conf *config.ConfigCtx, c *envoy_config_listener_v3.Filter, tls bind.TLS) (bind.StreamHandler, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_listener_v3.Filter_TypedConfig:
		msg, err := encoding.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	}

	switch c.Name {
	case wellknown.HTTPConnectionManager:
		switch p := filterConfig.(type) {
		case *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager:
			return Convert_extensions_filters_network_http_connection_manager_v3_HttpConnectionManager(conf, p, tls)
		}
	case wellknown.TCPProxy:
		switch p := filterConfig.(type) {
		case *envoy_extensions_filters_network_tcp_proxy_v3.TcpProxy:
			return Convert_extensions_filters_network_tcp_proxy_v3_TcpProxy(conf, p, tls)
		}
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_listener_v3.Filter %s\n", string(data))
	return nil, nil
}

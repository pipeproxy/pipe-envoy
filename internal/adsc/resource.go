package adsc

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_extensions_filters_network_http_connection_manager_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	resource "github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

const (
	ClusterType  = resource.ClusterType
	EndpointType = resource.EndpointType
	ListenerType = resource.ListenerType
	RouteType    = resource.RouteType
	SecretType   = resource.SecretType
	RuntimeType  = resource.RuntimeType
	AnyType      = resource.AnyType
)

// GetHTTPConnectionManager creates a HttpConnectionManager from filter
func GetHTTPConnectionManager(filter *envoy_config_listener_v3.Filter) *envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager {
	config := &envoy_extensions_filters_network_http_connection_manager_v3.HttpConnectionManager{}

	// use typed config if available
	if typedConfig := filter.GetTypedConfig(); typedConfig != nil {
		ptypes.UnmarshalAny(typedConfig, config)
	}
	return config
}

func GetSDSName(c *envoy_config_core_v3.TransportSocket) []string {
	if c == nil {
		return nil
	}

	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_core_v3.TransportSocket_TypedConfig:
		msg, err := encoding.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil
		}
		filterConfig = msg
	}

	switch c.Name {
	case wellknown.TransportSocketTls:
		switch p := filterConfig.(type) {
		case *envoy_extensions_transport_sockets_tls_v3.DownstreamTlsContext:
			return Convert_extensions_transport_sockets_tls_v3_CommonTlsContext(p.CommonTlsContext)
		case *envoy_extensions_transport_sockets_tls_v3.UpstreamTlsContext:
			return Convert_extensions_transport_sockets_tls_v3_CommonTlsContext(p.CommonTlsContext)
		}
	}
	return nil
}

func Convert_extensions_transport_sockets_tls_v3_CommonTlsContext(c *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext) []string {
	out := []string{}
	for _, t := range c.TlsCertificateSdsSecretConfigs {
		out = append(out, t.Name)
	}

	switch t := c.ValidationContextType.(type) {
	case *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_ValidationContextSdsSecretConfig:
		name := t.ValidationContextSdsSecretConfig.Name
		if name != "" {
			out = append(out, name)
		}
	//case *envoy_extensions_transport_sockets_tls_v3.CommonTlsContext_CombinedValidationContext:
	//	name := t.CombinedValidationContext.ValidationContextSdsSecretConfig.Name
	//	if name != "" {
	//		out = append(out, name)
	//	}
	default:
	}
	return out
}

func SelectFilterChain(fc []*envoy_config_listener_v3.FilterChain) *envoy_config_listener_v3.FilterChain {
	for _, filter := range fc {
		if filter.FilterChainMatch == nil {
			return filter
		}
		if filter.FilterChainMatch.DestinationPort.GetValue() != 0 {
			continue
		}
		if filter.FilterChainMatch.ApplicationProtocols == nil {
			continue
		}
		for _, proto := range filter.FilterChainMatch.ApplicationProtocols {
			if proto == "h2c" {
				return filter
			}
		}
	}
	return fc[0]
}

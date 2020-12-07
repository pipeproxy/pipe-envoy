package adsc

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
)

var (
	plaintextHTTPALPNs = []string{"http/1.0", "http/1.1", "h2c"}
	istioHTTPPlaintext = []string{"istio", "istio-http/1.0", "istio-http/1.1", "istio-h2"}
	httpTLS            = []string{"http/1.0", "http/1.1", "h2c", "istio-http/1.0", "istio-http/1.1", "istio-h2"}
	tcpTLS             = []string{"istio-peer-exchange", "istio"}

	protDescrs = map[string][]string{
		"App: HTTP TLS":         httpTLS,
		"App: Istio HTTP Plain": istioHTTPPlaintext,
		"App: TCP TLS":          tcpTLS,
		"App: HTTP":             plaintextHTTPALPNs,
	}
)

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

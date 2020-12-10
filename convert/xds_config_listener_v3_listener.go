package convert

import (
	"reflect"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_Listener(conf *config.ConfigCtx, c *envoy_config_listener_v3.Listener) (bind.Service, error) {
	if len(c.FilterChains) == 0 {
		return bind.NoneService{}, nil
	}

	network, address, err := Convert_config_core_v3_Address(conf, c.Address)
	if err != nil {
		return nil, err
	}

	//TODO: Support dynamic selection of filter,
	s, err := SelectFilterChain(conf, c.FilterChains)
	if err != nil {
		return nil, err
	}

	s = bind.LogStreamHandlerConfig{
		Handler: s,
		Output: bind.FileIoWriterConfig{
			Path: "/dev/stderr",
		},
	}

	var d bind.Service
	d = bind.StreamServiceConfig{
		DisconnectOnClose: true,
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigNetworkEnum(network),
			Address: address,
			Virtual: c.DeprecatedV1 != nil && !c.DeprecatedV1.BindToPort.GetValue(),
		},
		Handler: s,
	}

	if c.Name != "" {
		d = bind.DefServiceConfig{
			Name: c.Name,
			Def: bind.TagsServiceConfig{
				Service: d,
				Tag:     c.Name,
			},
		}
	}
	return d, nil
}

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

func SelectFilterChain(conf *config.ConfigCtx, fc []*envoy_config_listener_v3.FilterChain) (bind.StreamHandler, error) {
	if len(fc) == 1 {
		return Convert_config_listener_v3_FilterChain(conf, fc[0])
	}
	var tcpStream bind.StreamHandler
	var httpStream bind.StreamHandler
	var tcpTlsStream bind.StreamHandler
	var httpTlsStream bind.StreamHandler
	for _, filter := range fc {
		if filter.FilterChainMatch == nil {
			h, err := Convert_config_listener_v3_FilterChain(conf, filter)
			if err != nil {
				return nil, err
			}
			tcpStream = h
			continue
		}
		if filter.FilterChainMatch.DestinationPort.GetValue() != 0 {
			continue
		}
		if filter.FilterChainMatch.ApplicationProtocols == nil {
			if filter.FilterChainMatch.TransportProtocol == "raw_buffer" {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				tcpStream = h
			}
			continue
		}
		switch len(filter.FilterChainMatch.ApplicationProtocols) {
		case len(plaintextHTTPALPNs):
			if reflect.DeepEqual(plaintextHTTPALPNs, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				httpStream = h
			}
		case len(httpTLS):
			if reflect.DeepEqual(httpTLS, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				httpTlsStream = h
			}
		case len(tcpTLS):
			if reflect.DeepEqual(httpTLS, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				tcpTlsStream = h
			}
		}
	}

	var lastStream bind.StreamHandler
	isHttp := httpStream != nil
	isTcp := tcpStream != nil
	isTlsTcp := tcpTlsStream != nil
	isTlsHttp := httpTlsStream != nil
	isTls := isTlsTcp || isTlsHttp

	var tlsStream bind.StreamHandler
	if isTls {
		switch {
		case isTlsHttp && isTlsTcp:
			tlsStream = bind.MuxStreamHandlerConfig{
				Routes: []bind.MuxStreamHandlerRoute{
					{
						Pattern: "http",
						Handler: httpTlsStream,
					},
				},
				NotFound: tcpTlsStream,
			}
		case isTlsHttp && !isTlsTcp:
			tlsStream = httpTlsStream
		case isTlsTcp && !isTlsHttp:
			tlsStream = tcpTlsStream
		}
	}

	switch {
	case isHttp && isTcp && isTls:
		lastStream = bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "http",
					Handler: httpStream,
				},
				{
					Pattern: "tls",
					Handler: tlsStream,
				},
			},
			NotFound: tcpStream,
		}
	case isHttp && !isTcp && isTls:
		lastStream = bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "http",
					Handler: httpStream,
				},
				{
					Pattern: "tls",
					Handler: tlsStream,
				},
			},
		}
	case isTcp && !isHttp && isTls:
		lastStream = bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "tls",
					Handler: tlsStream,
				},
			},
			NotFound: tcpStream,
		}
	case isHttp && isTcp && !isTls:
		lastStream = bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "http",
					Handler: httpStream,
				},
			},
			NotFound: tcpStream,
		}
	case isHttp && !isTcp && !isTls:
		lastStream = httpStream
	case isTcp && !isHttp && !isTls:
		lastStream = tcpStream
	}
	return lastStream, nil
}

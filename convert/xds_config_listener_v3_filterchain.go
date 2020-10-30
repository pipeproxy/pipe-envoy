package convert

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_FilterChain(conf *config.ConfigCtx, c *envoy_config_listener_v3.FilterChain) (bind.StreamHandler, error) {

	var tls bind.TLS
	if c.TransportSocket != nil {
		t, err := Convert_config_core_v3_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return nil, err
		}
		tls = t
	}

	var httpStream bind.StreamHandler
	var tcpStream bind.StreamHandler
	var lastStream bind.StreamHandler = bind.NoneStreamHandler{}

	for _, filter := range c.Filters {
		s, err := Convert_config_listener_v3_Filter(conf, filter, tls)
		if err != nil {
			return nil, err
		}
		switch c.Name {
		case wellknown.HTTPConnectionManager:
			httpStream = s
		case wellknown.TCPProxy:
			tcpStream = s
		}
		lastStream = s
	}

	isHttp := httpStream != nil
	isTcp := tcpStream != nil
	switch {
	case isHttp && isTcp:
		lastStream = bind.MuxStreamHandlerConfig{
			Routes: []bind.MuxStreamHandlerRoute{
				{
					Pattern: "http",
					Handler: httpStream,
				},
			},
			NotFound: tcpStream,
		}
	case isHttp && !isTcp:
		lastStream = httpStream
	case isTcp && !isHttp:
		lastStream = tcpStream
	}
	return lastStream, nil
}

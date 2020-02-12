package convert

import (
	envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_listener_FilterChain(conf *config.ConfigCtx, c *envoy_api_v2_listener.FilterChain) (bind.StreamHandler, error) {
	var tls bind.TLS
	if c.TransportSocket != nil {
		t, err := Convert_api_v2_core_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return nil, err
		}
		tls = t
	} else if c.TlsContext != nil {
		t, err := Convert_api_v2_auth_DownstreamTlsContext(conf, c.TlsContext)
		if err != nil {
			return nil, err
		}
		tls = t
	}

	multi := []bind.StreamHandler{}
	for _, filter := range c.Filters {
		handler, err := Convert_api_v2_listener_Filter(conf, filter, tls)
		if err != nil {
			return nil, err
		}

		multi = append(multi, handler)
	}

	d := bind.StreamHandlerMultiConfig{Multi: multi}
	name := config.XdsName(c.Name)
	ref, err := conf.RegisterComponents(name, d)
	if err != nil {
		return nil, err
	}

	return bind.RefStreamHandler(ref), nil
}

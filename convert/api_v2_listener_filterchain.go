package convert

import (
	"encoding/json"

	envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_listener_FilterChain(conf *config.ConfigCtx, c *envoy_api_v2_listener.FilterChain) (string, error) {
	tlsName := ""
	if c.TransportSocket != nil {
		name, err := Convert_api_v2_core_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return "", err
		}
		tlsName = name
	} else if c.TlsContext != nil {
		name, err := Convert_api_v2_auth_DownstreamTlsContext(conf, c.TlsContext)
		if err != nil {
			return "", err
		}
		tlsName = name
	}

	switch len(c.Filters) {
	case 1:
		name, err := Convert_api_v2_listener_Filter(conf, c.Filters[0], tlsName)
		if err != nil {
			return "", err
		}
		if c.Name == "" {
			return name, nil
		}
		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}
		name = config.XdsName(c.Name)
		return conf.RegisterComponents(name, ref)
	default:
		list := []json.RawMessage{}
		for _, filter := range c.Filters {
			name, err := Convert_api_v2_listener_Filter(conf, filter, tlsName)
			if err != nil {
				return "", err
			}
			ref, err := config.MarshalRef(name)
			if err != nil {
				return "", err
			}

			list = append(list, ref)
		}

		d, err := config.MarshalKindStreamHandlerMulti(list)
		if err != nil {
			return "", err
		}

		name := config.XdsName(c.Name)
		return conf.RegisterComponents(name, d)
	}

}

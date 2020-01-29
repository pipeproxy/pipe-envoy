package convert_api_v2_endpoint

import (
	"encoding/json"

	envoy_api_v2_endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/wzshiming/envoy/config"
)

func Convert_LocalityLbEndpoints(conf *config.ConfigCtx, c *envoy_api_v2_endpoint.LocalityLbEndpoints) (string, error) {
	switch len(c.LbEndpoints) {
	case 1:
		return Convert_LbEndpoint(conf, c.LbEndpoints[0])
	default:
		list := []json.RawMessage{}
		for _, lbEndpoint := range c.LbEndpoints {
			name, err := Convert_LbEndpoint(conf, lbEndpoint)
			if err != nil {
				return "", err
			}
			ref, err := config.MarshalRef(name)
			if err != nil {
				return "", err
			}

			list = append(list, ref)
		}

		d, err := config.MarshalKindStreamHandlerPoller("round_robin", list)
		if err != nil {
			return "", err
		}

		return conf.RegisterComponents("", d)
	}
}

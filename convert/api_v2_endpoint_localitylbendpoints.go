package convert

import (
	"encoding/json"

	envoy_api_v2_endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_endpoint_LocalityLbEndpoints(conf *config.ConfigCtx, c *envoy_api_v2_endpoint.LocalityLbEndpoints) (string, error) {
	list := []json.RawMessage{}
	for _, lbEndpoint := range c.LbEndpoints {
		name, err := Convert_api_v2_endpoint_LbEndpoint(conf, lbEndpoint)
		if err != nil {
			return "", err
		}
		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		list = append(list, ref)
	}

	d, err := config.MarshalKindDialerPoller("round_robin", list)
	if err != nil {
		return "", err
	}

	return conf.RegisterComponents("", d)
}

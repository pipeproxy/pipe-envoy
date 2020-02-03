package convert

import (
	"encoding/json"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_ClusterLoadAssignment(conf *config.ConfigCtx, c *envoy_api_v2.ClusterLoadAssignment) (string, error) {
	list := []json.RawMessage{}
	for _, endpoint := range c.Endpoints {
		name, err := Convert_api_v2_endpoint_LocalityLbEndpoints(conf, endpoint)
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

	name := config.XdsName(c.ClusterName)

	return conf.RegisterComponents(name, d)
}

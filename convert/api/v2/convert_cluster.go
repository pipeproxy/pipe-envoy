package convert_api_v2

import (
	"encoding/json"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_core "github.com/wzshiming/envoy/convert/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_Cluster(conf *config.ConfigCtx, c *envoy_api_v2.Cluster) (string, error) {

	if c.ClusterDiscoveryType == nil {
		list := []json.RawMessage{}
		for _, host := range c.Hosts {
			name, err := convert_api_v2_core.Convert_AddressForward(conf, host)
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

		name := config.XdsName(c.Name)
		return conf.RegisterComponents(name, d)
	}

	switch d := c.ClusterDiscoveryType.(type) {
	case *envoy_api_v2.Cluster_Type:
		switch d.Type {
		case envoy_api_v2.Cluster_EDS:
			name := c.Name
			if name != "" {
				conf.AppendEDS(name)
				return config.XdsName(name), nil
			}
		}
	case *envoy_api_v2.Cluster_ClusterType:
	}

	logger.Todof("%#v", c)
	return "", nil
}

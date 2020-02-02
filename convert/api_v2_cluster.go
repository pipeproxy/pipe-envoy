package convert

import (
	"encoding/json"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_Cluster(conf *config.ConfigCtx, c *envoy_api_v2.Cluster) (string, error) {

	tlsName := ""
	if c.TransportSocket != nil {
		name, err := Convert_api_v2_core_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return "", err
		}
		tlsName = name
	} else if c.TlsContext != nil {
		name, err := Convert_api_v2_auth_UpstreamTlsContext(conf, c.TlsContext)
		if err != nil {
			return "", err
		}
		tlsName = name
	}

	if c.ClusterDiscoveryType == nil {
		list := []json.RawMessage{}
		for _, host := range c.Hosts {
			name, err := Convert_api_v2_core_AddressForward(conf, host)
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

		if tlsName != "" {
			tlsRef, err := config.MarshalRef(tlsName)
			if err != nil {
				return "", err
			}

			d, err = config.MarshalKindStreamHandlerTlsUp(tlsRef, d)
			if err != nil {
				return "", err
			}
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
				name := config.XdsName(name)
				if tlsName != "" {
					tlsRef, err := config.MarshalRef(tlsName)
					if err != nil {
						return "", err
					}

					edsRef, err := config.MarshalRef(name)
					if err != nil {
						return "", err
					}

					d, err := config.MarshalKindStreamHandlerTlsUp(tlsRef, edsRef)
					if err != nil {
						return "", err
					}
					return conf.RegisterComponents("", d)
				}
				return config.XdsName(name), nil
			}
		}
	case *envoy_api_v2.Cluster_ClusterType:
	}

	logger.Todof("%#v", c)
	return "", nil
}

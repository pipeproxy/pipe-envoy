package convert

import (
	"log"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_cluster_v3_Cluster(conf *config.ConfigCtx, c *envoy_config_cluster_v3.Cluster) (bind.StreamDialer, error) {
	var tls bind.TLS
	if c.TransportSocket != nil {
		t, err := Convert_config_core_v3_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return nil, err
		}
		tls = t
	}

	var d bind.StreamDialer
	switch {
	case c.LoadAssignment != nil:
		dialer, err := Convert_config_endpoint_v3_ClusterLoadAssignment(conf, c.LoadAssignment)
		if err != nil {
			return nil, err
		}
		d = dialer
	case c.ClusterDiscoveryType != nil:
		switch t := c.ClusterDiscoveryType.(type) {
		case *envoy_config_cluster_v3.Cluster_Type:
			switch t.Type {
			case envoy_config_cluster_v3.Cluster_EDS:
				d = conf.EDS(c.Name)
			case envoy_config_cluster_v3.Cluster_ORIGINAL_DST:
				d = bind.DialerStreamDialerConfig{
					Original: true,
					Virtual:  c.UpstreamBindConfig == nil,
				}
			case envoy_config_cluster_v3.Cluster_STATIC:
				d = bind.NoneStreamDialer{}
			default:
				data, _ := encoding.Marshal(c)
				log.Printf("[TODO] envoy_config_core_v3.Cluster_Type %s\n", string(data))
			}
		case *envoy_config_cluster_v3.Cluster_ClusterType:
			data, _ := encoding.Marshal(c)
			log.Printf("[TODO] envoy_config_core_v3.Cluster_ClusterType %s\n", string(data))
			return nil, nil
		}
	default:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_core_v3.Cluster %s\n", string(data))
	}

	if tls != nil {
		d = bind.TLSStreamDialerConfig{
			Dialer: d,
			TLS:    tls,
		}
	}

	if c.Name != "" {
		d = conf.RegisterCDS(c.Name, d, c)
	}

	return d, nil
}

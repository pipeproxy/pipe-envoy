package convert

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_Cluster(conf *config.ConfigCtx, c *envoy_api_v2.Cluster) (bind.StreamDialer, error) {

	var tls bind.TLS
	if c.TransportSocket != nil {
		t, err := Convert_api_v2_core_TransportSocket(conf, c.TransportSocket)
		if err != nil {
			return nil, err
		}
		tls = t
	} else if c.TlsContext != nil {
		t, err := Convert_api_v2_auth_UpstreamTlsContext(conf, c.TlsContext)
		if err != nil {
			return nil, err
		}
		tls = t
	}

	var d bind.StreamDialer

	switch {
	case c.LoadAssignment != nil:
		dialer, err := Convert_api_v2_ClusterLoadAssignment(conf, c.LoadAssignment)
		if err != nil {
			return nil, err
		}
		d = dialer
	case c.ClusterDiscoveryType != nil:
		switch t := c.ClusterDiscoveryType.(type) {
		case *envoy_api_v2.Cluster_Type:
			switch t.Type {
			case envoy_api_v2.Cluster_EDS:
				name := c.Name
				if name != "" {
					conf.AppendEDS(name)
					d = bind.RefStreamDialer(config.XdsName(name))
				}
			}
		case *envoy_api_v2.Cluster_ClusterType:
			logger.Todof("%#v", c)
			return nil, nil
		}
	default:
		dialers := []bind.StreamDialer{}
		for _, host := range c.Hosts {
			dialer, err := Convert_api_v2_core_AddressDialer(conf, host)
			if err != nil {
				return nil, err
			}
			dialers = append(dialers, dialer)
		}
		d = bind.StreamDialerPollerConfig{
			Poller:  "round_robin",
			Dialers: dialers,
		}
	}

	if tls != nil {
		d = bind.StreamDialerTLSConfig{
			Dialer: d,
			TLS:    tls,
		}
	}

	if c.Name != "" {
		ref, err := conf.RegisterComponents(config.XdsName(c.Name), d)
		if err != nil {
			return nil, err
		}
		d = bind.RefStreamDialer(ref)
	}

	return d, nil
}

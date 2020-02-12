package convert

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_Cluster(conf *config.ConfigCtx, c *envoy_api_v2.Cluster) (bind.Dialer, error) {

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

	if c.ClusterDiscoveryType == nil {
		dialers := []bind.Dialer{}
		for _, host := range c.Hosts {
			dialer, err := Convert_api_v2_core_AddressDialer(conf, host)
			if err != nil {
				return nil, err
			}
			dialers = append(dialers, dialer)
		}

		var d bind.Dialer = bind.DialerPollerConfig{
			Poller:  "round_robin",
			Dialers: dialers,
		}
		if tls != nil {
			d = bind.DialerTlsConfig{
				Dialer: d,
				TLS:    tls,
			}
		}

		ref, err := conf.RegisterComponents(config.XdsName(c.Name), d)
		if err != nil {
			return nil, err
		}

		return bind.RefDialer(ref), nil
	}

	switch d := c.ClusterDiscoveryType.(type) {
	case *envoy_api_v2.Cluster_Type:
		switch d.Type {
		case envoy_api_v2.Cluster_EDS:
			name := c.Name
			if name != "" {
				conf.AppendEDS(name)

				var d bind.Dialer = bind.RefDialer(config.XdsName(name))
				if tls != nil {
					d = bind.DialerTlsConfig{
						Dialer: d,
						TLS:    tls,
					}
				}

				ref, err := conf.RegisterComponents("", d)
				if err != nil {
					return nil, err
				}

				return bind.RefDialer(ref), nil
			}
		}
	case *envoy_api_v2.Cluster_ClusterType:
	}

	logger.Todof("%#v", c)
	return nil, nil
}

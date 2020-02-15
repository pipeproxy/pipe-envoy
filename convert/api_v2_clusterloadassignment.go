package convert

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_ClusterLoadAssignment(conf *config.ConfigCtx, c *envoy_api_v2.ClusterLoadAssignment) (bind.StreamDialer, error) {
	dialers := []bind.StreamDialer{}
	for _, endpoint := range c.Endpoints {
		dialer, err := Convert_api_v2_endpoint_LocalityLbEndpoints(conf, endpoint)
		if err != nil {
			return nil, err
		}

		dialers = append(dialers, dialer)
	}

	d := bind.StreamDialerPollerConfig{
		Poller:  "round_robin",
		Dialers: dialers,
	}

	ref, err := conf.RegisterComponents(config.XdsName(c.ClusterName), d)
	if err != nil {
		return nil, err
	}

	return bind.RefStreamDialer(ref), nil
}

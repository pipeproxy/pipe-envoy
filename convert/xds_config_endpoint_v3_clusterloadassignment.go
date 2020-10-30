package convert

import (
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_endpoint_v3_ClusterLoadAssignment(conf *config.ConfigCtx, c *envoy_config_endpoint_v3.ClusterLoadAssignment) (bind.StreamDialer, error) {
	dialers := []bind.LbStreamDialerWeight{}
	for _, endpoint := range c.Endpoints {
		dialer, err := Convert_config_endpoint_v3_LocalityLbEndpoints(conf, endpoint)
		if err != nil {
			return nil, err
		}

		dialers = append(dialers, bind.LbStreamDialerWeight{
			Weight: uint(endpoint.LoadBalancingWeight.GetValue()),
			Dialer: dialer,
		})
	}

	var d bind.StreamDialer
	switch len(dialers) {
	case 0:
		d = bind.NoneStreamDialer{}
	case 1:
		d = dialers[0].Dialer
	default:
		d = bind.LbStreamDialerConfig{
			Policy:  bind.LbStreamDialerLoadBalancePolicyEnumEnumRoundRobin,
			Dialers: dialers,
		}
	}

	if c.ClusterName != "" {
		d = bind.DefStreamDialerConfig{
			Name: "eds." + c.ClusterName,
			Def:  d,
		}
	}
	return d, nil
}

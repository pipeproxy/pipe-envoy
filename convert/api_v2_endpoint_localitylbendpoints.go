package convert

import (
	envoy_api_v2_endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_endpoint_LocalityLbEndpoints(conf *config.ConfigCtx, c *envoy_api_v2_endpoint.LocalityLbEndpoints) (bind.StreamDialer, error) {
	dialers := []bind.StreamDialer{}
	for _, lbEndpoint := range c.LbEndpoints {
		dialer, err := Convert_api_v2_endpoint_LbEndpoint(conf, lbEndpoint)
		if err != nil {
			return nil, err
		}

		dialers = append(dialers, dialer)
	}

	d := bind.StreamDialerPollerConfig{
		Poller:  "round_robin",
		Dialers: dialers,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	return bind.RefStreamDialer(ref), nil
}

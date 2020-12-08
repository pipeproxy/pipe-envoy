package convert

import (
	"log"

	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_endpoint_v3_LbEndpoint(conf *config.ConfigCtx, c *envoy_config_endpoint_v3.LbEndpoint) (bind.StreamDialer, error) {
	switch h := c.HostIdentifier.(type) {
	case *envoy_config_endpoint_v3.LbEndpoint_Endpoint:
		network, address, err := Convert_config_core_v3_Address(conf, h.Endpoint.Address)
		if err != nil {
			return nil, err
		}
		return bind.DialerStreamDialerConfig{
			Network: bind.DialerStreamDialerNetworkEnum(network),
			Address: address,
		}, nil
	case *envoy_config_endpoint_v3.LbEndpoint_EndpointName:
		return bind.RefStreamDialerConfig{
			Name: h.EndpointName,
		}, nil
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_endpoint_v3.LbEndpoint %s\n", string(data))
	return nil, nil
}

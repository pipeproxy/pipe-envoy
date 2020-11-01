package convert

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_Listener(conf *config.ConfigCtx, c *envoy_config_listener_v3.Listener) (bind.Service, error) {
	if c.DeprecatedV1 != nil && !c.DeprecatedV1.BindToPort.GetValue() {
		return bind.NoneService{}, nil
	}
	if len(c.FilterChains) == 0 || len(c.FilterChains[0].Filters) == 0 {
		return bind.NoneService{}, nil
	}

	network, address, err := Convert_config_core_v3_Address(conf, c.Address)
	if err != nil {
		return nil, err
	}

	filterChain := c.FilterChains[0]
	s, err := Convert_config_listener_v3_FilterChain(conf, filterChain)
	if err != nil {
		return nil, err
	}

	var d bind.Service
	d = bind.StreamServiceConfig{
		DisconnectOnClose: true,
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigListenerNetworkEnum(network),
			Address: address,
		},
		Handler: s,
	}

	if c.Name != "" {
		d = bind.DefServiceConfig{
			Name: c.Name,
			Def:  d,
		}
	}
	return d, nil
}

package convert

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/internal/adsc"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_Listener(conf *config.ConfigCtx, c *envoy_config_listener_v3.Listener) (bind.Service, error) {
	if len(c.FilterChains) == 0 {
		return bind.NoneService{}, nil
	}

	network, address, err := Convert_config_core_v3_Address(conf, c.Address)
	if err != nil {
		return nil, err
	}

	//TODO: Support dynamic selection of filter,
	filterChain := adsc.SelectFilterChain(c.FilterChains)
	s, err := Convert_config_listener_v3_FilterChain(conf, filterChain)
	if err != nil {
		return nil, err
	}

	s = bind.LogStreamHandlerConfig{
		Handler: s,
		Output: bind.FileIoWriterConfig{
			Path: "/dev/stderr",
		},
	}

	var d bind.Service
	d = bind.StreamServiceConfig{
		DisconnectOnClose: true,
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigListenerNetworkEnum(network),
			Address: address,
			Virtual: c.DeprecatedV1 != nil && !c.DeprecatedV1.BindToPort.GetValue(),
		},
		Handler: s,
	}

	if c.Name != "" {
		d = bind.DefServiceConfig{
			Name: c.Name,
			Def: bind.TagsServiceConfig{
				Service: d,
				Tag:     c.Name,
			},
		}
	}
	return d, nil
}

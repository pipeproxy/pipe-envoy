package convert

import (
	"fmt"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_Listener(conf *config.ConfigCtx, c *envoy_api_v2.Listener) (bind.Service, error) {
	listener, err := Convert_api_v2_core_AddressListener(conf, c.Address)
	if err != nil {
		return nil, err
	}

	if len(c.FilterChains) == 0 || len(c.FilterChains[0].Filters) == 0 {
		return nil, fmt.Errorf("not filter chains")
	}

	multi := []bind.StreamHandler{}
	for _, filterChain := range c.FilterChains {
		handler, err := Convert_api_v2_listener_FilterChain(conf, filterChain)
		if err != nil {
			return nil, err
		}

		multi = append(multi, handler)
	}

	d := bind.ServiceStreamConfig{
		Listener: listener,
		Handler: bind.StreamHandlerMultiConfig{
			Multi: multi,
		},
	}

	ref, err := conf.RegisterComponents(config.XdsName(c.Name), d)
	if err != nil {
		return nil, err
	}

	err = conf.RegisterService(ref)
	if err != nil {
		return nil, err
	}

	return bind.RefService(ref), nil

}

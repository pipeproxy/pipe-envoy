package convert

import (
	"fmt"
	"strings"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_AddressDialer(conf *config.ConfigCtx, c *envoy_api_v2_core.Address) (bind.Dialer, error) {
	network, address, err := convertAddress(c)
	if err != nil {
		return nil, err
	}

	d := bind.DialerNetworkConfig{
		Network: network,
		Address: address,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	return bind.RefDialer(ref), nil
}

func Convert_api_v2_core_AddressListener(conf *config.ConfigCtx, c *envoy_api_v2_core.Address) (bind.ListenerListenConfig, error) {
	network, address, err := convertAddress(c)
	if err != nil {
		return nil, err
	}

	d := bind.ListenerListenConfigNetworkConfig{
		Network: network,
		Address: address,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	return bind.RefListenerListenConfig(ref), nil
}

func convertAddress(c *envoy_api_v2_core.Address) (string, string, error) {
	switch a := c.Address.(type) {
	case *envoy_api_v2_core.Address_SocketAddress:
		addr := a.SocketAddress
		network := strings.ToLower(addr.Protocol.String())
		if addr.Ipv4Compat {
			network = network + "4"
		}
		switch p := addr.PortSpecifier.(type) {
		case *envoy_api_v2_core.SocketAddress_PortValue:
			address := fmt.Sprintf("%s:%d", addr.Address, p.PortValue)
			return network, address, nil
		case *envoy_api_v2_core.SocketAddress_NamedPort:

		}

	case *envoy_api_v2_core.Address_Pipe:
	}
	logger.Todof("%#v", c)
	return "", "", fmt.Errorf("todo envoy_api_v2_core.Address")
}

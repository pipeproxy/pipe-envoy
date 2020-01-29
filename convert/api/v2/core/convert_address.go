package convert_api_v2_core

import (
	"fmt"
	"strings"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_AddressForward(conf *config.ConfigCtx, c *envoy_api_v2_core.Address) (string, error) {
	network, address, err := convertAddress(c)
	if err != nil {
		return "", err
	}
	d, err := config.MarshalKindStreamHandlerForward(network, address)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)
}

func Convert_AddressListener(conf *config.ConfigCtx, c *envoy_api_v2_core.Address) (string, error) {
	network, address, err := convertAddress(c)
	if err != nil {
		return "", err
	}
	d, err := config.MarshalKindListenConfigNetwork(network, address)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)
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

package convert

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
)

func Convert_config_core_v3_Address(conf *config.ConfigCtx, c *envoy_config_core_v3.Address) (string, string, error) {
	switch a := c.Address.(type) {
	case *envoy_config_core_v3.Address_SocketAddress:
		addr := a.SocketAddress
		network := strings.ToLower(addr.Protocol.String())
		if addr.Ipv4Compat {
			network = network + "4"
		}
		switch p := addr.PortSpecifier.(type) {
		case *envoy_config_core_v3.SocketAddress_PortValue:
			address := net.JoinHostPort(addr.Address, strconv.Itoa(int(p.PortValue)))
			return network, address, nil
		case *envoy_config_core_v3.SocketAddress_NamedPort:
			address := net.JoinHostPort(addr.Address, p.NamedPort)
			return network, address, nil
		}

	case *envoy_config_core_v3.Address_Pipe:
		return "unix", a.Pipe.Path, nil
	}

	return "", "", fmt.Errorf("todo envoy_config_core_v3.Address")
}

package address

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

func newSocketAddress(config *envoy_api_v2_core.SocketAddress) (*socketAddress, error) {
	network := strings.ToLower(config.Protocol.String())
	if config.Ipv4Compat {
		network = network + "4"
	}

	if config.ResolverName != "" {
		logger.Todoln("ResolverName", config)
	}

	switch p := config.PortSpecifier.(type) {
	case *envoy_api_v2_core.SocketAddress_PortValue:
		return &socketAddress{
			network: network,
			address: config.Address,
			port:    p.PortValue,
		}, nil
	case *envoy_api_v2_core.SocketAddress_NamedPort:
		logger.Todoln("NamedPort", p)
	}
	return nil, errors.New("todo socket address")
}

type socketAddress struct {
	network string
	address string
	port    uint32
}

func (s *socketAddress) Listen(ctx context.Context) (net.Listener, error) {
	var listenConfig net.ListenConfig
	return listenConfig.Listen(ctx, s.network, fmt.Sprintf("%s:%d", s.address, s.port))
}

func (s *socketAddress) String() string {
	return fmt.Sprintf("%s://%s:%d", s.network, s.address, s.port)
}

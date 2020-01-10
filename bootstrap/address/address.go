package address

import (
	"context"
	"errors"
	"net"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

type Address interface {
	Listen(ctx context.Context) (net.Listener, error)
}

func NewAddress(config *envoy_api_v2_core.Address) (Address, error) {
	switch a := config.Address.(type) {
	case *envoy_api_v2_core.Address_SocketAddress:
		return newSocketAddress(a.SocketAddress)
	case *envoy_api_v2_core.Address_Pipe:
		logger.Todoln("address", a)
	}
	return nil, errors.New("todo address")
}

package xds

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/envoy/convert"

	"github.com/wzshiming/envoy/ads"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/stream"
)

const name = "xds"

func init() {
	configure.Register(name, NewXDSWithConfig)
}

type Config struct {
	NodeID  string
	Forward stream.Handler
}

var xds *XDS

func NewXDSWithConfig(ctx context.Context, conf *Config) (once.Once, error) {
	if xds != nil {
		return xds, nil
	}

	adsConfig := ads.Config{
		NodeID: conf.NodeID,
		ContextDialer: func(ctx context.Context, s string) (conn net.Conn, err error) {
			p1, p2 := net.Pipe()
			go conf.Forward.ServeStream(ctx, p1)
			return p2, nil
		},
	}

	config, ok := convert.GetXdsWithContext(ctx)
	if !ok || config == nil {
		return nil, fmt.Errorf("xds content is not configured")
	}

	x, err := NewXDS(config, &adsConfig)
	if err != nil {
		return nil, err
	}
	xds = x
	return x, nil
}

package ads

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/envoy/ads"
	"github.com/wzshiming/envoy/convert"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/stream"
)

const (
	name = "ads"
)

func init() {
	configure.Register(name, NewADSWithConfig)
}

type Config struct {
	Name    string `json:"@Name"`
	NodeID  string
	Forward stream.Handler
}

var adsMap = map[string]*ADS{}

func NewADSWithConfig(ctx context.Context, conf *Config) (once.Once, error) {
	if a, ok := adsMap[conf.Name]; ok {
		return a, nil
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

	a, err := NewADS(config, &adsConfig)
	if err != nil {
		return nil, err
	}

	adsMap[conf.Name] = a
	return a, nil
}

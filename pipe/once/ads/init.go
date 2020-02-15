package ads

import (
	"context"
	"fmt"
	"net"

	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/client/ads"
	"github.com/wzshiming/envoy/internal/node"
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

const (
	name = "ads"
)

func init() {
	decode.Register(name, NewADSWithConfig)
}

type Config struct {
	Name   string `json:"@Name"`
	NodeID string
	Dialer dialer.Dialer
}

var adsMap = map[string]*ADS{}

func NewADSWithConfig(ctx context.Context, conf *Config) (once.Once, error) {
	if a, ok := adsMap[conf.Name]; ok {
		return a, nil
	}

	adsConfig := ads.Config{
		NodeConfig: &node.Config{
			NodeID: conf.NodeID,
		},
		ContextDialer: func(ctx context.Context, s string) (conn net.Conn, err error) {
			return conf.Dialer.DialStream(ctx)
		},
	}

	config, ok := config.GetXdsWithContext(ctx)
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

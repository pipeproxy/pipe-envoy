package direct

import (
	"time"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/proxy"
)

type Config struct {
	XDSAddr  string
	NodeID   string
	Metadata map[string]interface{}
	BasePath string
}

func NewDirect(c *Config) (*proxy.Proxy, error) {
	meta, err := MapToProtoStruct(c.Metadata)
	if err != nil {
		return nil, err
	}
	return proxy.NewProxy(&proxy.Config{
		Node: &envoy_config_core_v3.Node{
			Id:       c.NodeID,
			Metadata: meta,
		},
		BasePath: c.BasePath,
		XDS:      c.XDSAddr,
		Interval: time.Second,
	})
}

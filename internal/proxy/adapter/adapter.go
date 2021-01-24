package adapter

import (
	"io/ioutil"
	"time"

	"github.com/pipeproxy/pipe-xds/internal/convert"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe-xds/internal/proxy"
)

const (
	XDS = `unix:./etc/istio/proxy/XDS`
	SDS = `unix:./etc/istio/proxy/SDS`
)

type Config struct {
	ConfigFile string
	BasePath   string
}

func NewAdapter(c *Config) (*proxy.Proxy, error) {
	data, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		return nil, err
	}
	bootstrap, err := encoding.UnmarshalBootstrap(data)
	if err != nil {
		return nil, err
	}
	p, err := proxy.NewProxy(&proxy.Config{
		IsDynamic: bootstrap.DynamicResources != nil,
		Node:      bootstrap.GetNode(),
		BasePath:  c.BasePath,
		XDS:       XDS,
		SDS:       SDS,
		Interval:  time.Second,
	})
	if err != nil {
		return nil, err
	}
	_, err = convert.Convert_config_bootstrap_v3_Bootstrap(p.ConfigCtx(), bootstrap)
	if err != nil {
		return nil, err
	}
	return p, nil
}

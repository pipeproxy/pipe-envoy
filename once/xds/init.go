package xds

import (
	"fmt"

	"github.com/wzshiming/envoy/once/ads"

	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
)

const (
	name = "xds"
)

func init() {
	configure.Register(name, NewXDSWithConfig)
}

type Config struct {
	XDS string
	ADS once.Once
}

func NewXDSWithConfig(conf *Config) (once.Once, error) {
	a, ok := conf.ADS.(*ads.ADS)
	if !ok || a == nil {
		return nil, fmt.Errorf("need ads")
	}

	switch conf.XDS {
	default:
		return nil, fmt.Errorf("%q is not define in XDS", conf.XDS)
	case "cds", "lds":
	}
	return NewXDS(a, conf.XDS), nil
}

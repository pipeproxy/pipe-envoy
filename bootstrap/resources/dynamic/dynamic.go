package dynamic

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bootstrap/resources/dynamic/xds"
	"github.com/wzshiming/envoy/internal/logger"
)

type DynamicResources struct {
	lds *xds.ConfigSource
	cds *xds.ConfigSource
	ads *xds.ApiConfigSource
}

func NewDynamicResources(config *envoy_config_bootstrap_v2.Bootstrap_DynamicResources) (*DynamicResources, error) {
	s := &DynamicResources{}

	if config.LdsConfig != nil {
		lds, err := xds.NewConfigSource(config.LdsConfig)
		if err != nil {
			return nil, err
		}
		s.lds = lds
	}

	if config.CdsConfig != nil {
		cds, err := xds.NewConfigSource(config.CdsConfig)
		if err != nil {
			return nil, err
		}
		s.cds = cds
	}

	if config.AdsConfig != nil {
		ads, err := xds.NewApiConfigSource(config.AdsConfig)
		if err != nil {
			return nil, err
		}
		s.ads = ads
	}
	return s, nil
}

func (d *DynamicResources) Start() error {
	logger.Todoln("DynamicResources", "start")
	return nil
}

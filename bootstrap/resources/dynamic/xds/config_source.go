package xds

import (
	"errors"
	"time"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/bootstrap/utils"
)

type ConfigSource struct {
	configSourceSpecifier ConfigSourceSpecifier
	initialFetchTimeout   time.Duration
	resourceApiVersion    envoy_api_v2_core.ApiVersion
}

func NewConfigSource(config *envoy_api_v2_core.ConfigSource) (*ConfigSource, error) {
	c := &ConfigSource{}

	c.resourceApiVersion = config.ResourceApiVersion

	initialFetchTimeout, err := utils.Duration(config.InitialFetchTimeout)
	if err != nil {
		return nil, err
	}
	c.initialFetchTimeout = initialFetchTimeout

	switch s := config.ConfigSourceSpecifier.(type) {
	case *envoy_api_v2_core.ConfigSource_Path:
		conf, err := newConfigSourcePath(s)
		if err != nil {
			return nil, err
		}
		c.configSourceSpecifier = conf
	case *envoy_api_v2_core.ConfigSource_ApiConfigSource:
		conf, err := NewApiConfigSource(s.ApiConfigSource)
		if err != nil {
			return nil, err
		}
		c.configSourceSpecifier = conf
	case *envoy_api_v2_core.ConfigSource_Ads:
		conf, err := newConfigSourceAds(s)
		if err != nil {
			return nil, err
		}
		c.configSourceSpecifier = conf
	case *envoy_api_v2_core.ConfigSource_Self:
		conf, err := newConfigSourceSelf(s)
		if err != nil {
			return nil, err
		}
		c.configSourceSpecifier = conf
	default:
		return nil, errors.New("todo socket address")
	}

	return c, nil
}

type ConfigSourceSpecifier interface {
}

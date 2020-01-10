package xds

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

type configSourceAds struct {
}

func newConfigSourceAds(config *envoy_api_v2_core.ConfigSource_Ads) (*configSourceAds, error) {
	logger.Todoln("ConfigSource_Ads", config)
	return &configSourceAds{}, nil
}

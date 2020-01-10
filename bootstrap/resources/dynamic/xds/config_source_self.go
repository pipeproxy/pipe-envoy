package xds

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

type configSourceSelf struct {
}

func newConfigSourceSelf(config *envoy_api_v2_core.ConfigSource_Self) (*configSourceSelf, error) {
	logger.Todoln("ConfigSource_Self", config)
	return &configSourceSelf{}, nil
}

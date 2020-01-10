package xds

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

type configSourcePath struct {
	path string
}

func newConfigSourcePath(config *envoy_api_v2_core.ConfigSource_Path) (*configSourcePath, error) {
	return &configSourcePath{config.Path}, nil
}

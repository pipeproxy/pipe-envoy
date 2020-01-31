package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_ConfigSource(conf *config.ConfigCtx, c *envoy_api_v2_core.ConfigSource) (string, error) {
	switch s := c.ConfigSourceSpecifier.(type) {
	case *envoy_api_v2_core.ConfigSource_Path:
	case *envoy_api_v2_core.ConfigSource_ApiConfigSource:
		return Convert_api_v2_core_ApiConfigSource(conf, s.ApiConfigSource)
	case *envoy_api_v2_core.ConfigSource_Ads:
		return config.KindOnceADS, nil
	case *envoy_api_v2_core.ConfigSource_Self:
	}

	logger.Todof("%#v", c)
	return "", nil
}

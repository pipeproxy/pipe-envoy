package convert_api_v2_listener

import (
	envoy_api_v2_listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_ActiveRawUdpListenerConfig(conf *config.ConfigCtx, c *envoy_api_v2_listener.ActiveRawUdpListenerConfig) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

package convert

import (
	"log"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_config_listener_v3_ActiveRawUdpListenerConfig(conf *config.ConfigCtx, c *envoy_config_listener_v3.ActiveRawUdpListenerConfig) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_listener_v3.ActiveRawUdpListenerConfig %s\n", string(data))
	return "", nil
}

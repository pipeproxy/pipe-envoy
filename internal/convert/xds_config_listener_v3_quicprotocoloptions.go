package convert

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_listener_v3_QuicProtocolOptions(conf *config.ConfigCtx, c *envoy_config_listener_v3.QuicProtocolOptions) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_listener_v3.QuicProtocolOptions %s\n", string(data))
	return "", nil
}

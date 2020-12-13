package convert

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"log"
)

func Convert_config_listener_v3_UdpGsoBatchWriterOptions(conf *config.ConfigCtx, c *envoy_config_listener_v3.UdpGsoBatchWriterOptions) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_listener_v3.UdpGsoBatchWriterOptions %s\n", string(data))
	return "", nil
}

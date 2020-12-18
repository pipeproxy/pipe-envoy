package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_core_v3_KeepaliveSettings(conf *config.ConfigCtx, c *envoy_config_core_v3.KeepaliveSettings) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.KeepaliveSettings %s\n", string(data))
	return "", nil
}

package convert

import (
	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_bootstrap_v3_RuntimeLayer_DiskLayer(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.RuntimeLayer_DiskLayer) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_bootstrap_v3.RuntimeLayer_DiskLayer %s\n", string(data))
	return "", nil
}

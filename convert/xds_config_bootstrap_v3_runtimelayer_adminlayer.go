package convert

import (
	"log"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_config_bootstrap_v3_RuntimeLayer_AdminLayer(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.RuntimeLayer_AdminLayer) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_bootstrap_v3.RuntimeLayer_AdminLayer %s\n", string(data))
	return "", nil
}

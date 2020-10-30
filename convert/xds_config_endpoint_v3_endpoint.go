package convert

import (
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_endpoint_v3_Endpoint(conf *config.ConfigCtx, c *envoy_config_endpoint_v3.Endpoint) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_endpoint_v3.Endpoint %s\n", string(data))
	return "", nil
}

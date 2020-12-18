package convert

import (
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_endpoint_v3_ClusterLoadAssignment_Policy(conf *config.ConfigCtx, c *envoy_config_endpoint_v3.ClusterLoadAssignment_Policy) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_endpoint_v3.ClusterLoadAssignment_Policy %s\n", string(data))
	return "", nil
}

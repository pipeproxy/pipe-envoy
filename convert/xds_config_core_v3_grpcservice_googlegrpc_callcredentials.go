package convert

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_core_v3_GrpcService_GoogleGrpc_CallCredentials(conf *config.ConfigCtx, c *envoy_config_core_v3.GrpcService_GoogleGrpc_CallCredentials) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.GrpcService_GoogleGrpc_CallCredentials %s\n", string(data))
	return "", nil
}

package convert

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"

	"log"
)

func Convert_config_core_v3_GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials(conf *config.ConfigCtx, c *envoy_config_core_v3.GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.GrpcService_GoogleGrpc_CallCredentials_ServiceAccountJWTAccessCredentials %s\n", string(data))
	return "", nil
}

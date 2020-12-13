package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_core_v3_GrpcService(conf *config.ConfigCtx, c *envoy_config_core_v3.GrpcService) (bind.StreamDialer, error) {
	switch t := c.TargetSpecifier.(type) {
	case *envoy_config_core_v3.GrpcService_EnvoyGrpc_:
		return conf.CDS(t.EnvoyGrpc.ClusterName), nil
	case *envoy_config_core_v3.GrpcService_GoogleGrpc_:
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.GrpcService %s\n", string(data))
	return nil, nil
}

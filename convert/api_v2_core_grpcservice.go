package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_GrpcService(conf *config.ConfigCtx, c *envoy_api_v2_core.GrpcService) (string, error) {
	switch t := c.TargetSpecifier.(type) {
	case *envoy_api_v2_core.GrpcService_EnvoyGrpc_:
		clusterName := t.EnvoyGrpc.ClusterName
		return config.XdsName(clusterName), nil
	case *envoy_api_v2_core.GrpcService_GoogleGrpc_:
	}
	logger.Todof("%#v", c)
	return "", nil
}

package xds

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

type grpcServiceEnvoyGrpc struct {
	clusterName string
}

func newGrpcServiceEnvoyGrpc(config *envoy_api_v2_core.GrpcService_EnvoyGrpc) (*grpcServiceEnvoyGrpc, error) {
	return &grpcServiceEnvoyGrpc{
		clusterName: config.ClusterName,
	}, nil
}

package xds

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

type grpcServiceGoogleGrpc struct {
}

func newGrpcServiceGoogleGrpc(config *envoy_api_v2_core.GrpcService_GoogleGrpc) (*grpcServiceGoogleGrpc, error) {
	logger.Todoln("GrpcService_GoogleGrpc", config)
	return &grpcServiceGoogleGrpc{}, nil
}

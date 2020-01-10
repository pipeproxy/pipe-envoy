package xds

import (
	"net/http"
	"time"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/bootstrap/utils"
)

type GrpcService struct {
	targetSpecifier targetSpecifier
	timeout         time.Duration
	initialMetadata http.Header
}

func NewGrpcService(config *envoy_api_v2_core.GrpcService) (*GrpcService, error) {
	g := &GrpcService{}
	timeout, err := utils.Duration(config.Timeout)
	if err != nil {
		return nil, err
	}
	g.timeout = timeout

	switch t := config.TargetSpecifier.(type) {
	case *envoy_api_v2_core.GrpcService_EnvoyGrpc_:
		s, err := newGrpcServiceEnvoyGrpc(t.EnvoyGrpc)
		if err != nil {
			return nil, err
		}
		g.targetSpecifier = s
	case *envoy_api_v2_core.GrpcService_GoogleGrpc_:
		s, err := newGrpcServiceGoogleGrpc(t.GoogleGrpc)
		if err != nil {
			return nil, err
		}
		g.targetSpecifier = s
	}

	initialMetadata := http.Header{}
	for _, header := range config.InitialMetadata {
		initialMetadata[header.Key] = append(initialMetadata[header.Key], header.Value)
	}
	g.initialMetadata = initialMetadata
	return g, nil
}

type targetSpecifier interface {
}

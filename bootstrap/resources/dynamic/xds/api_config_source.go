package xds

import (
	"time"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/internal/logger"
)

type ApiConfigSource struct {
	apiType             envoy_api_v2_core.ApiConfigSource_ApiType
	transportApiVersion envoy_api_v2_core.ApiVersion
	clusterNames        []string
	grpcServices        []*GrpcService
	refreshDelay        time.Duration
	requestTimeout      time.Duration
}

func NewApiConfigSource(config *envoy_api_v2_core.ApiConfigSource) (*ApiConfigSource, error) {
	c := &ApiConfigSource{}
	c.apiType = config.ApiType
	c.transportApiVersion = config.TransportApiVersion
	c.clusterNames = config.ClusterNames

	for _, s := range config.GrpcServices {
		svc, err := NewGrpcService(s)
		if err != nil {
			return nil, err
		}
		c.grpcServices = append(c.grpcServices, svc)
	}

	if config.RateLimitSettings != nil {
		logger.Todoln("RateLimitSettings", config.RateLimitSettings)
	}

	if config.SetNodeOnFirstMessageOnly {
		logger.Todoln("SetNodeOnFirstMessageOnly", config.SetNodeOnFirstMessageOnly)
	}

	return c, nil
}

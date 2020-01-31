package convert

import (
	envoy_api_v2_endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_endpoint_ClusterStats_DroppedRequests(conf *config.ConfigCtx, c *envoy_api_v2_endpoint.ClusterStats_DroppedRequests) (string, error) {
	logger.Todof("%#v", c)
	return "", nil
}

package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

const AdsName = "ads@ads"

func Convert_api_v2_core_ApiConfigSource(conf *config.ConfigCtx, c *envoy_api_v2_core.ApiConfigSource) (string, error) {
	switch c.ApiType {
	case envoy_api_v2_core.ApiConfigSource_GRPC:
		if len(c.GrpcServices) == 1 {
			svc := c.GrpcServices[0]
			dialer, err := Convert_api_v2_core_GrpcService(conf, svc)
			if err != nil {
				return "", err
			}
			nodeID := ""
			node, ok := GetNodeWithContext(conf.Ctx())
			if ok {
				nodeID = node.Id
			}

			ref, err := conf.RegisterComponents(AdsName, bind.OnceAdsConfig{
				NodeID: nodeID,
				Dialer: dialer,
			})
			if err != nil {
				return "", err
			}

			return ref, nil
		}
	}
	logger.Todof("%#v", c)
	return "", nil
}

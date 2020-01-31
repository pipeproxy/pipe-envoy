package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_ApiConfigSource(conf *config.ConfigCtx, c *envoy_api_v2_core.ApiConfigSource) (string, error) {
	switch c.ApiType {
	case envoy_api_v2_core.ApiConfigSource_GRPC:
		if len(c.GrpcServices) == 1 {
			svc := c.GrpcServices[0]
			switch t := svc.TargetSpecifier.(type) {
			case *envoy_api_v2_core.GrpcService_EnvoyGrpc_:
				clusterName := t.EnvoyGrpc.ClusterName
				ref := config.XdsName(clusterName)
				r, err := config.MarshalRef(ref)
				if err != nil {
					return "", err
				}

				nodeID := ""
				node, ok := GetNodeWithContext(conf.Ctx())
				if ok {
					nodeID = node.Id
				}

				r, err = config.MarshalKindOnceADS(nodeID, r)
				if err != nil {
					return "", err
				}
				return conf.RegisterComponents("", r)
			case *envoy_api_v2_core.GrpcService_GoogleGrpc_:
			}
		}
	}
	logger.Todof("%#v", c)
	return "", nil
}

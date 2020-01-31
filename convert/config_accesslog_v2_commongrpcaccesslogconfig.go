package convert

import (
	"encoding/json"

	envoy_config_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_accesslog_v2_CommonGrpcAccessLogConfig(conf *config.ConfigCtx, c *envoy_config_accesslog_v2.CommonGrpcAccessLogConfig, handler json.RawMessage) (string, error) {

	ref, err := Convert_api_v2_core_GrpcService(conf, c.GrpcService)
	if err != nil {
		return "", err
	}
	r, err := config.MarshalRef(ref)
	if err != nil {
		return "", err
	}
	nodeID := ""
	node, ok := GetNodeWithContext(conf.Ctx())
	if ok {
		nodeID = node.Id
	}
	r, err = config.MarshalKindOnceAccessLog(nodeID, c.LogName, r)
	if err != nil {
		return "", err
	}
	name, err := conf.RegisterComponents("", r)
	if err != nil {
		return "", err
	}

	r, err = config.MarshalRef(name)
	if err != nil {
		return "", err
	}

	r, err = config.MarshalKindHttpHandlerAccessLog(r, handler)
	if err != nil {
		return "", err
	}

	return conf.RegisterComponents("", r)
}

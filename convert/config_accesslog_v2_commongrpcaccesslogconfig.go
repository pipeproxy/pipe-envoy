package convert

import (
	envoy_config_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_accesslog_v2_CommonGrpcAccessLogConfig(conf *config.ConfigCtx, c *envoy_config_accesslog_v2.CommonGrpcAccessLogConfig, handler bind.HTTPHandler) (bind.HTTPHandler, error) {

	dialer, err := Convert_api_v2_core_GrpcService(conf, c.GrpcService)
	if err != nil {
		return nil, err
	}

	nodeID := ""
	node, ok := GetNodeWithContext(conf.Ctx())
	if ok {
		nodeID = node.Id
	}

	name, err := conf.RegisterComponents("", bind.OnceAccessLogConfig{
		NodeID:  nodeID,
		LogName: c.LogName,
		Dialer:  dialer,
	})
	if err != nil {
		return nil, err
	}

	r := bind.HTTPHandlerAccessLogConfig{
		AccessLog: bind.RefOnce(name),
		Handler:   handler,
	}

	ref, err := conf.RegisterComponents("", r)
	if err != nil {
		return nil, err
	}

	return bind.RefHTTPHandler(ref), nil
}

package convert

import (
	"encoding/json"
	"fmt"

	"github.com/wzshiming/envoy/internal/logger"

	envoy_config_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"

	"github.com/golang/protobuf/proto"

	envoy_config_filter_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_filter_accesslog_v2_AccessLog(conf *config.ConfigCtx, c *envoy_config_filter_accesslog_v2.AccessLog, handler json.RawMessage) (string, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_filter_accesslog_v2.AccessLog_TypedConfig:
		msg, err := config.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return "", err
		}
		filterConfig = msg
	case *envoy_config_filter_accesslog_v2.AccessLog_Config:
		return "", fmt.Errorf("not suppert envoy_config_filter_accesslog_v2.AccessLog_Config")
	}

	switch f := filterConfig.(type) {
	case *envoy_config_accesslog_v2.HttpGrpcAccessLogConfig:
		return Convert_config_accesslog_v2_CommonGrpcAccessLogConfig(conf, f.CommonConfig, handler)
	case *envoy_config_accesslog_v2.FileAccessLog:
		d, err := config.MarshalKindOutputFile(f.Path)
		if err != nil {
			return "", err
		}
		d, err = config.MarshalKindHttpHandlerLog(d, handler)
		if err != nil {
			return "", err
		}
		return conf.RegisterComponents("", d)
	case *envoy_config_accesslog_v2.CommonGrpcAccessLogConfig:
		return Convert_config_accesslog_v2_CommonGrpcAccessLogConfig(conf, f, handler)
	}
	logger.Todof("%#v", c)
	return "", nil
}

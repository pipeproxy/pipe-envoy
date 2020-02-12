package convert

import (
	"fmt"

	envoy_config_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	envoy_config_filter_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	"github.com/golang/protobuf/proto"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
	"github.com/wzshiming/envoy/wellknown"
)

func Convert_config_filter_accesslog_v2_AccessLog(conf *config.ConfigCtx, c *envoy_config_filter_accesslog_v2.AccessLog, handler bind.HttpHandler) (bind.HttpHandler, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_filter_accesslog_v2.AccessLog_TypedConfig:
		msg, err := config.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	case *envoy_config_filter_accesslog_v2.AccessLog_Config:
		return nil, fmt.Errorf("not suppert envoy_config_filter_accesslog_v2.AccessLog_Config")
	}

	switch c.Name {
	case wellknown.HTTPGRPCAccessLog:
		switch f := filterConfig.(type) {
		case *envoy_config_accesslog_v2.HttpGrpcAccessLogConfig:
			return Convert_config_accesslog_v2_CommonGrpcAccessLogConfig(conf, f.CommonConfig, handler)
		}
	case wellknown.FileAccessLog:
		switch f := filterConfig.(type) {
		case *envoy_config_accesslog_v2.FileAccessLog:
			d := bind.HttpHandlerLogConfig{
				Output:  bind.OutputFileConfig{Path: f.Path},
				Handler: handler,
			}

			ref, err := conf.RegisterComponents("", d)
			if err != nil {
				return nil, err
			}

			return bind.RefHttpHandler(ref), nil
		}
	}
	logger.Todof("%#v", c)
	return nil, nil
}

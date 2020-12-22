package convert

import (
	"log"

	envoy_config_accesslog_v3 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	envoy_extensions_access_loggers_file_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/file/v3"
	"github.com/golang/protobuf/proto"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_accesslog_v3_AccessLog_StreamHandler(conf *config.ConfigCtx, c *envoy_config_accesslog_v3.AccessLog, h bind.StreamHandler) (bind.StreamHandler, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_accesslog_v3.AccessLog_TypedConfig:
		msg, err := encoding.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	}
	switch c.Name {
	case "envoy.access_loggers.file":
		switch p := filterConfig.(type) {
		case *envoy_extensions_access_loggers_file_v3.FileAccessLog:
			return bind.LogStreamHandlerConfig{
				Handler: h,
				Output: bind.FileIoWriterConfig{
					Path: p.Path,
				},
			}, nil
		}
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_accesslog_v3.AccessLog %s\n", string(data))
	return bind.LogStreamHandlerConfig{
		Handler: h,
		Output: bind.FileIoWriterConfig{
			Path: "/dev/stderr",
		},
	}, nil
}

func Convert_config_accesslog_v3_AccessLog_HTTPHandler(conf *config.ConfigCtx, c *envoy_config_accesslog_v3.AccessLog, h bind.HTTPHandler) (bind.HTTPHandler, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_accesslog_v3.AccessLog_TypedConfig:
		msg, err := encoding.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	}
	switch c.Name {
	case "envoy.access_loggers.file":
		switch p := filterConfig.(type) {
		case *envoy_extensions_access_loggers_file_v3.FileAccessLog:
			return bind.LogNetHTTPHandlerConfig{
				Handler: h,
				Output: bind.FileIoWriterConfig{
					Path: p.Path,
				},
			}, nil
		}
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_accesslog_v3.AccessLog %s\n", string(data))
	return bind.LogNetHTTPHandlerConfig{
		Handler: h,
		Output: bind.FileIoWriterConfig{
			Path: "/dev/stderr",
		},
	}, nil
}

package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_DataSource(conf *config.ConfigCtx, c *envoy_api_v2_core.DataSource) (bind.Input, error) {
	switch s := c.Specifier.(type) {
	case *envoy_api_v2_core.DataSource_Filename:
		return bind.InputFileConfig{Path: s.Filename}, nil
	case *envoy_api_v2_core.DataSource_InlineBytes:
		return bind.InputInlineConfig{Data: string(s.InlineBytes)}, nil
	case *envoy_api_v2_core.DataSource_InlineString:
		return bind.InputInlineConfig{Data: s.InlineString}, nil
	}
	logger.Todof("%#v", c)
	return nil, nil
}

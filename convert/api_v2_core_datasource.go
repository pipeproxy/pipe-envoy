package convert

import (
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_DataSource(conf *config.ConfigCtx, c *envoy_api_v2_core.DataSource) (string, error) {
	switch s := c.Specifier.(type) {
	case *envoy_api_v2_core.DataSource_Filename:
		d, err := config.MarshalKindInputFile(s.Filename)
		if err != nil {
			return "", err
		}
		return conf.RegisterComponents("", d)
	case *envoy_api_v2_core.DataSource_InlineBytes:
		d, err := config.MarshalKindInputInline(string(s.InlineBytes))
		if err != nil {
			return "", err
		}
		return conf.RegisterComponents("", d)
	case *envoy_api_v2_core.DataSource_InlineString:
		d, err := config.MarshalKindInputInline(s.InlineString)
		if err != nil {
			return "", err
		}
		return conf.RegisterComponents("", d)
	}
	logger.Todof("%#v", c)
	return "", nil
}

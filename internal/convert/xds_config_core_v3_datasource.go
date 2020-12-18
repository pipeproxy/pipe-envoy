package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_core_v3_DataSource(conf *config.ConfigCtx, c *envoy_config_core_v3.DataSource) (bind.IoReader, error) {
	switch s := c.Specifier.(type) {
	case *envoy_config_core_v3.DataSource_Filename:
		return bind.FileIoReaderConfig{Path: s.Filename}, nil
	case *envoy_config_core_v3.DataSource_InlineBytes:
		return bind.InlineIoReaderConfig{Data: string(s.InlineBytes)}, nil
	case *envoy_config_core_v3.DataSource_InlineString:
		return bind.InlineIoReaderConfig{Data: s.InlineString}, nil
	}
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.DataSource %s\n", string(data))
	return nil, nil
}

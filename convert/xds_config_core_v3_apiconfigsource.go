package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_core_v3_ApiConfigSource(conf *config.ConfigCtx, c *envoy_config_core_v3.ApiConfigSource) (bind.StreamDialer, error) {
	switch c.ApiType {
	case envoy_config_core_v3.ApiConfigSource_GRPC:
		if len(c.GrpcServices) == 1 {
			svc := c.GrpcServices[0]
			return Convert_config_core_v3_GrpcService(conf, svc)
		}
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.ApiConfigSource %s\n", string(data))
	return nil, nil
}

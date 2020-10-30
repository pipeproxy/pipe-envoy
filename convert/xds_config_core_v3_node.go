package convert

import (
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/config"
)

func Convert_config_core_v3_Node(conf *config.ConfigCtx, c *envoy_config_core_v3.Node) (string, error) {
	conf.SetNodeID(c.Id)
	return "", nil
}

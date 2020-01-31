package convert

import (
	"context"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/wzshiming/envoy/config"
)

type nodeCtxKeyType int

func GetNodeWithContext(ctx context.Context) (*envoy_api_v2_core.Node, bool) {
	i := ctx.Value(nodeCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*envoy_api_v2_core.Node)
	return p, ok
}

func Convert_api_v2_core_Node(conf *config.ConfigCtx, c *envoy_api_v2_core.Node) (string, error) {
	if c != nil {
		conf.WithValue(nodeCtxKeyType(0), c)
	}
	return "", nil
}

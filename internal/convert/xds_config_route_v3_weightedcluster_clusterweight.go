package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_WeightedCluster_ClusterWeight(conf *config.ConfigCtx, c *envoy_config_route_v3.WeightedCluster_ClusterWeight) (bind.LbNetHTTPHandlerWeight, error) {
	return bind.LbNetHTTPHandlerWeight{
		Weight: uint(c.Weight.GetValue()),
		Handler: bind.ForwardNetHTTPHandlerConfig{
			H2c:    true,
			Dialer: conf.CDS(c.Name),
		},
	}, nil
}

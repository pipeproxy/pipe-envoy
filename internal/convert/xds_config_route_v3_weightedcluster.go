package convert

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_WeightedCluster(conf *config.ConfigCtx, c *envoy_config_route_v3.WeightedCluster) (bind.HTTPHandler, error) {
	d := bind.LbNetHTTPHandlerConfig{
		Policy: bind.RoundRobinBalancePolicy{},
	}
	for _, weighted := range c.Clusters {
		w, err := Convert_config_route_v3_WeightedCluster_ClusterWeight(conf, weighted)
		if err != nil {
			return nil, err
		}
		d.Handlers = append(d.Handlers, w)
	}
	return d, nil
}

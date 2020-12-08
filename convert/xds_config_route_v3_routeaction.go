package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteAction(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteAction) (bind.HTTPHandler, error) {
	switch s := c.ClusterSpecifier.(type) {
	case *envoy_config_route_v3.RouteAction_Cluster:
		return bind.ForwardNetHTTPHandlerConfig{
			Dialer: bind.RefStreamDialerConfig{
				Name: s.Cluster,
			},
		}, nil
	//case *envoy_config_route_v3.RouteAction_ClusterHeader:

	case *envoy_config_route_v3.RouteAction_WeightedClusters:
		d := bind.LbNetHTTPHandlerConfig{
			Policy: bind.RoundRobinBalancePolicy{},
		}
		for _, weighted := range s.WeightedClusters.Clusters {
			w, err := Convert_config_route_v3_WeightedCluster_ClusterWeight(conf, weighted)
			if err != nil {
				return nil, err
			}
			d.Handlers = append(d.Handlers, w)
		}
		return d, nil
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RouteAction %s\n", string(data))
	return nil, nil
}

package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteAction(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteAction) (bind.HTTPHandler, error) {
	switch s := c.ClusterSpecifier.(type) {
	case *envoy_config_route_v3.RouteAction_Cluster:
		return bind.ForwardNetHTTPHandlerConfig{
			H2c:    true,
			Dialer: conf.CDS(s.Cluster),
		}, nil
	//case *envoy_config_route_v3.RouteAction_ClusterHeader:
	case *envoy_config_route_v3.RouteAction_WeightedClusters:
		return Convert_config_route_v3_WeightedCluster(conf, s.WeightedClusters)
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RouteAction %s\n", string(data))
	return nil, nil
}

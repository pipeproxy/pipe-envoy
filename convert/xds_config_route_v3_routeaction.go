package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteAction(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteAction) (bind.HTTPHandler, error) {
	name := ""
	switch s := c.ClusterSpecifier.(type) {
	case *envoy_config_route_v3.RouteAction_Cluster:
		name = s.Cluster
	case *envoy_config_route_v3.RouteAction_ClusterHeader:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.RouteAction %s\n", string(data))
		return nil, nil
	case *envoy_config_route_v3.RouteAction_WeightedClusters:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_config_route_v3.RouteAction %s\n", string(data))
		return nil, nil
	}

	d := bind.ForwardNetHTTPHandlerConfig{
		Dialer: bind.RefStreamDialerConfig{
			Name: name,
		},
	}
	return d, nil
}

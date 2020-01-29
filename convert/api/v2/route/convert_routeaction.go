package convert_api_v2_route

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_RouteAction(conf *config.ConfigCtx, c *envoy_api_v2_route.RouteAction) (string, error) {
	name := ""
	switch s := c.ClusterSpecifier.(type) {
	case *envoy_api_v2_route.RouteAction_Cluster:
		name = s.Cluster
	case *envoy_api_v2_route.RouteAction_ClusterHeader:
		logger.Todof("%#v", c)
		return "", nil
	case *envoy_api_v2_route.RouteAction_WeightedClusters:
		logger.Todof("%#v", c)
		return "", nil
	}

	refName := config.XdsName(name)
	ref, err := config.MarshalRef(refName)
	if err != nil {
		return "", err
	}

	d, err := config.MarshalKindHttpHandlerForward("http://"+name, ref)
	if err != nil {
		return "", err
	}

	return conf.RegisterComponents("", d)
}

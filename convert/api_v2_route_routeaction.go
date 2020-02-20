package convert

import (
	envoy_api_v2_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_route_RouteAction(conf *config.ConfigCtx, c *envoy_api_v2_route.RouteAction) (bind.HTTPHandler, error) {
	name := ""
	switch s := c.ClusterSpecifier.(type) {
	case *envoy_api_v2_route.RouteAction_Cluster:
		name = s.Cluster
	case *envoy_api_v2_route.RouteAction_ClusterHeader:
		logger.Todof("%#v", c)
		return nil, nil
	case *envoy_api_v2_route.RouteAction_WeightedClusters:
		logger.Todof("%#v", c)
		return nil, nil
	}

	d := bind.HTTPHandlerForwardConfig{
		RoundTripper: bind.HTTPRoundTripperTransportConfig{
			Dialer: bind.RefStreamDialer(config.XdsName(name)),
		},
		URL: "http://" + name,
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	return bind.RefHTTPHandler(ref), nil
}

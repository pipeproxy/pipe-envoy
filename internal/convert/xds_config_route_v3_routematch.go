package convert

import (
	"log"

	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_type_matcher_v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_route_v3_RouteMatch(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteMatch) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_route_v3.RouteMatch %s\n", string(data))
	return "", nil
}

func Convert_config_route_v3_RouteMatch_Path(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteMatch, handler bind.HTTPHandler) bind.PathNetHTTPHandlerRoute {
	path := bind.PathNetHTTPHandlerRoute{
		Handler: handler,
	}
	switch p := c.PathSpecifier.(type) {
	case *envoy_config_route_v3.RouteMatch_Prefix:
		path.Prefix = p.Prefix
	case *envoy_config_route_v3.RouteMatch_Path:
		path.Path = p.Path
	case *envoy_config_route_v3.RouteMatch_SafeRegex:
		path.Regexp = p.SafeRegex.Regex
	}
	return path
}

func Convert_config_route_v3_RouteMatch_Header(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteMatch) []bind.HeaderNetHTTPHandlerRouteMatch {
	if len(c.Headers) == 0 {
		return nil
	}
	headers := []bind.HeaderNetHTTPHandlerRouteMatch{}
	for _, c := range c.Headers {
		switch p := c.HeaderMatchSpecifier.(type) {
		case *envoy_config_route_v3.HeaderMatcher_ExactMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:   c.Name,
				Exact: p.ExactMatch,
			})
		case *envoy_config_route_v3.HeaderMatcher_SafeRegexMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:    c.Name,
				Regexp: p.SafeRegexMatch.Regex,
			})
		case *envoy_config_route_v3.HeaderMatcher_RangeMatch:

		case *envoy_config_route_v3.HeaderMatcher_PresentMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:     c.Name,
				Present: p.PresentMatch,
			})
		case *envoy_config_route_v3.HeaderMatcher_PrefixMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:    c.Name,
				Prefix: p.PrefixMatch,
			})
		case *envoy_config_route_v3.HeaderMatcher_SuffixMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:    c.Name,
				Suffix: p.SuffixMatch,
			})
		case *envoy_config_route_v3.HeaderMatcher_ContainsMatch:
			headers = append(headers, bind.HeaderNetHTTPHandlerRouteMatch{
				Key:      c.Name,
				Contains: p.ContainsMatch,
			})
		}
	}

	return headers
}

func Convert_config_route_v3_RouteMatch_Query(conf *config.ConfigCtx, c *envoy_config_route_v3.RouteMatch) []bind.QueryNetHTTPHandlerRouteMatch {
	if len(c.QueryParameters) == 0 {
		return nil
	}
	queries := []bind.QueryNetHTTPHandlerRouteMatch{}
	for _, c := range c.QueryParameters {
		switch p := c.QueryParameterMatchSpecifier.(type) {
		case *envoy_config_route_v3.QueryParameterMatcher_StringMatch:
			if p.StringMatch.IgnoreCase {
				continue
			}
			switch p := p.StringMatch.MatchPattern.(type) {
			case *envoy_type_matcher_v3.StringMatcher_Exact:
				queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
					Key:   c.Name,
					Exact: p.Exact,
				})
			case *envoy_type_matcher_v3.StringMatcher_Prefix:
				queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
					Key:    c.Name,
					Prefix: p.Prefix,
				})
			case *envoy_type_matcher_v3.StringMatcher_Suffix:
				queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
					Key:    c.Name,
					Suffix: p.Suffix,
				})
			case *envoy_type_matcher_v3.StringMatcher_SafeRegex:
				queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
					Key:    c.Name,
					Regexp: p.SafeRegex.Regex,
				})
			case *envoy_type_matcher_v3.StringMatcher_Contains:
				queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
					Key:      c.Name,
					Contains: p.Contains,
				})
			}
		case *envoy_config_route_v3.QueryParameterMatcher_PresentMatch:
			queries = append(queries, bind.QueryNetHTTPHandlerRouteMatch{
				Key:     c.Name,
				Present: p.PresentMatch,
			})
		}
	}
	return queries
}

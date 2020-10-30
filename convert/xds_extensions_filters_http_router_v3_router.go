package convert

import (
	"log"

	envoy_extensions_filters_http_router_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_extensions_filters_http_router_v3_Router(conf *config.ConfigCtx, c *envoy_extensions_filters_http_router_v3.Router) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_http_router_v3.Router %s\n", string(data))
	return "", nil
}

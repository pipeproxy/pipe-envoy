package convert

import (
	"log"

	envoy_extensions_filters_http_cors_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cors/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_extensions_filters_http_cors_v3_Cors(conf *config.ConfigCtx, c *envoy_extensions_filters_http_cors_v3.Cors) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_http_cors_v3.Cors %s\n", string(data))
	return "", nil
}

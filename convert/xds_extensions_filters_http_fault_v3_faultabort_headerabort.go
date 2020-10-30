package convert

import (
	"log"

	envoy_extensions_filters_http_fault_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_extensions_filters_http_fault_v3_FaultAbort_HeaderAbort(conf *config.ConfigCtx, c *envoy_extensions_filters_http_fault_v3.FaultAbort_HeaderAbort) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_http_fault_v3.FaultAbort_HeaderAbort %s\n", string(data))
	return "", nil
}

package convert

import (
	"log"

	envoy_extensions_filters_listener_original_dst_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_dst/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_extensions_filters_listener_original_dst_v3_OriginalDst(conf *config.ConfigCtx, c *envoy_extensions_filters_listener_original_dst_v3.OriginalDst) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_extensions_filters_listener_original_dst_v3.OriginalDst %s\n", string(data))
	return "", nil
}

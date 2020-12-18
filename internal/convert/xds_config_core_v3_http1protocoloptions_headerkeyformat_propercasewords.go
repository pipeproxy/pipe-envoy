package convert

import (
	"log"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_config_core_v3_Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords(conf *config.ConfigCtx, c *envoy_config_core_v3.Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.Http1ProtocolOptions_HeaderKeyFormat_ProperCaseWords %s\n", string(data))
	return "", nil
}

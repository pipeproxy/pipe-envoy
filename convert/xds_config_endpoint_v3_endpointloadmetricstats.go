package convert

import (
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"

	"log"
)

func Convert_config_endpoint_v3_EndpointLoadMetricStats(conf *config.ConfigCtx, c *envoy_config_endpoint_v3.EndpointLoadMetricStats) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_endpoint_v3.EndpointLoadMetricStats %s\n", string(data))
	return "", nil
}

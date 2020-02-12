package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_DownstreamTlsContext(conf *config.ConfigCtx, c *envoy_api_v2_auth.DownstreamTlsContext) (bind.TLS, error) {
	return Convert_api_v2_auth_CommonTlsContext(conf, c.CommonTlsContext)
}

package convert

import (
	"fmt"

	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/golang/protobuf/proto"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_core_TransportSocket(conf *config.ConfigCtx, c *envoy_api_v2_core.TransportSocket) (string, error) {
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_api_v2_core.TransportSocket_TypedConfig:
		msg, err := config.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return "", err
		}
		filterConfig = msg
	case *envoy_api_v2_core.TransportSocket_Config:
		return "", fmt.Errorf("not suppert envoy_api_v2_core.TransportSocket_Config")
	}

	switch c.Name {
	case "envoy.transport_sockets.tls":
		switch p := filterConfig.(type) {
		case *envoy_api_v2_auth.DownstreamTlsContext:
			return Convert_api_v2_auth_DownstreamTlsContext(conf, p)
		}
	}

	logger.Todof("%#v", c)
	return "", nil
}

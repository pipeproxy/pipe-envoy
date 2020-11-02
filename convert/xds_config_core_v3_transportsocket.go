package convert

import (
	"log"
	"os"
	"strconv"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/proto"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_core_v3_TransportSocket(conf *config.ConfigCtx, c *envoy_config_core_v3.TransportSocket) (bind.TLS, error) {
	if ok, _ := strconv.ParseBool(os.Getenv("TransportSocket")); !ok {
		return nil, nil
	}
	var filterConfig proto.Message
	switch t := c.ConfigType.(type) {
	case *envoy_config_core_v3.TransportSocket_TypedConfig:
		msg, err := encoding.UnmarshalAny(t.TypedConfig)
		if err != nil {
			return nil, err
		}
		filterConfig = msg
	}

	switch c.Name {
	case wellknown.TransportSocketTls:
		switch p := filterConfig.(type) {
		case *envoy_extensions_transport_sockets_tls_v3.DownstreamTlsContext:
			return Convert_extensions_transport_sockets_tls_v3_DownstreamTlsContext(conf, p)
		case *envoy_extensions_transport_sockets_tls_v3.UpstreamTlsContext:
			return Convert_extensions_transport_sockets_tls_v3_UpstreamTlsContext(conf, p)
		}
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_core_v3.TransportSocket %s\n", string(data))
	return nil, nil
}

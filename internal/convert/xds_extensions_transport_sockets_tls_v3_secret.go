package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe/bind"

	"log"
)

func Convert_extensions_transport_sockets_tls_v3_Secret(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.Secret) (bind.TLS, error) {
	var d bind.TLS
	switch t := c.Type.(type) {
	case *envoy_extensions_transport_sockets_tls_v3.Secret_TlsCertificate:
		tls, err := Convert_extensions_transport_sockets_tls_v3_TlsCertificate(conf, t.TlsCertificate)
		if err != nil {
			return nil, err
		}
		d = tls
	case *envoy_extensions_transport_sockets_tls_v3.Secret_SessionTicketKeys:
		tls, err := Convert_extensions_transport_sockets_tls_v3_TlsSessionTicketKeys(conf, t.SessionTicketKeys)
		if err != nil {
			return nil, err
		}
		d = tls
	case *envoy_extensions_transport_sockets_tls_v3.Secret_ValidationContext:
		tls, err := Convert_extensions_transport_sockets_tls_v3_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return nil, err
		}
		d = tls
	default:
		data, _ := encoding.Marshal(c)
		log.Printf("[TODO] envoy_extensions_transport_sockets_tls_v3.Secret %s\n", string(data))
		return nil, nil
	}
	if c.Name != "" {
		d = conf.RegisterSDS(c.Name, d, c)
	}
	return d, nil
}

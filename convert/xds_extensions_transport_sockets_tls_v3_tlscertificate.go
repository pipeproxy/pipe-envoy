package convert

import (
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_extensions_transport_sockets_tls_v3_TlsCertificate(conf *config.ConfigCtx, c *envoy_extensions_transport_sockets_tls_v3.TlsCertificate) (bind.TLS, error) {
	keyRef, err := Convert_config_core_v3_DataSource(conf, c.PrivateKey)
	if err != nil {
		return nil, err
	}

	certRef, err := Convert_config_core_v3_DataSource(conf, c.CertificateChain)
	if err != nil {
		return nil, err
	}

	d := bind.FromTLSConfig{
		Cert: certRef,
		Key:  keyRef,
	}
	return d, nil
}

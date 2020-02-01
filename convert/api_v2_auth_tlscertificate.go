package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_TlsCertificate(conf *config.ConfigCtx, c *envoy_api_v2_auth.TlsCertificate) (string, error) {
	name, err := Convert_api_v2_core_DataSource(conf, c.PrivateKey)
	if err != nil {
		return "", err
	}
	keyRef, err := config.MarshalRef(name)
	if err != nil {
		return "", err
	}

	name, err = Convert_api_v2_core_DataSource(conf, c.CertificateChain)
	if err != nil {
		return "", err
	}
	certRef, err := config.MarshalRef(name)
	if err != nil {
		return "", err
	}

	d, err := config.MarshalKindTlsFrom("", certRef, keyRef)
	if err != nil {
		return "", err
	}
	return conf.RegisterComponents("", d)
}

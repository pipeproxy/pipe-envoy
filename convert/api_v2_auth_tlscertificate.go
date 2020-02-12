package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_api_v2_auth_TlsCertificate(conf *config.ConfigCtx, c *envoy_api_v2_auth.TlsCertificate) (bind.TLS, error) {
	keyRef, err := Convert_api_v2_core_DataSource(conf, c.PrivateKey)
	if err != nil {
		return nil, err
	}

	certRef, err := Convert_api_v2_core_DataSource(conf, c.CertificateChain)
	if err != nil {
		return nil, err
	}

	ref, err := conf.RegisterComponents("", bind.TLSFromConfig{
		Domain: "",
		Cert:   certRef,
		Key:    keyRef,
	})
	if err != nil {
		return nil, err
	}

	return bind.RefTLS(ref), nil
}

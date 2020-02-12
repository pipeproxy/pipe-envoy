package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_auth_Secret(conf *config.ConfigCtx, c *envoy_api_v2_auth.Secret) (bind.TLS, error) {
	switch t := c.Type.(type) {
	case *envoy_api_v2_auth.Secret_TlsCertificate:
		tls, err := Convert_api_v2_auth_TlsCertificate(conf, t.TlsCertificate)
		if err != nil {
			return nil, err
		}
		if tls == nil {
			return nil, nil
		}

		ref, err := conf.RegisterComponents(config.XdsName(c.Name), tls)
		if err != nil {
			return nil, err
		}

		return bind.RefTLS(ref), nil
	case *envoy_api_v2_auth.Secret_SessionTicketKeys:
		tls, err := Convert_api_v2_auth_TlsSessionTicketKeys(conf, t.SessionTicketKeys)
		if err != nil {
			return nil, err
		}
		if tls == nil {
			return nil, nil
		}

		ref, err := conf.RegisterComponents(config.XdsName(c.Name), tls)
		if err != nil {
			return nil, err
		}

		return bind.RefTLS(ref), nil
	case *envoy_api_v2_auth.Secret_ValidationContext:
		tls, err := Convert_api_v2_auth_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return nil, err
		}
		if tls == nil {
			return nil, nil
		}

		ref, err := conf.RegisterComponents(config.XdsName(c.Name), tls)
		if err != nil {
			return nil, err
		}

		return bind.RefTLS(ref), nil
	}
	logger.Todof("%#v", c)
	return nil, nil
}

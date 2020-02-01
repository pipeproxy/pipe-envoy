package convert

import (
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/internal/logger"
)

func Convert_api_v2_auth_Secret(conf *config.ConfigCtx, c *envoy_api_v2_auth.Secret) (string, error) {
	switch t := c.Type.(type) {
	case *envoy_api_v2_auth.Secret_TlsCertificate:
		name, err := Convert_api_v2_auth_TlsCertificate(conf, t.TlsCertificate)
		if err != nil {
			return "", err
		}
		if name == "" {
			return "", nil
		}

		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		name = config.XdsName(c.Name)
		return conf.RegisterComponents(name, ref)
	case *envoy_api_v2_auth.Secret_SessionTicketKeys:
		name, err := Convert_api_v2_auth_TlsSessionTicketKeys(conf, t.SessionTicketKeys)
		if err != nil {
			return "", err
		}
		if name == "" {
			return "", nil
		}

		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		name = config.XdsName(c.Name)
		return conf.RegisterComponents(name, ref)
	case *envoy_api_v2_auth.Secret_ValidationContext:
		name, err := Convert_api_v2_auth_CertificateValidationContext(conf, t.ValidationContext)
		if err != nil {
			return "", err
		}
		if name == "" {
			return "", nil
		}

		ref, err := config.MarshalRef(name)
		if err != nil {
			return "", err
		}

		name = config.XdsName(c.Name)
		return conf.RegisterComponents(name, ref)
	}
	logger.Todof("%#v", c)
	return "", nil
}

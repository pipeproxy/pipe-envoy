package validation

import (
	"crypto/tls"
	"crypto/x509"
)

func NewValidation(ca []byte) (*tls.Config, error) {
	conf := &tls.Config{}
	conf.RootCAs = x509.NewCertPool()
	conf.RootCAs.AppendCertsFromPEM(ca)
	return conf, nil
}

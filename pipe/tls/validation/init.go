package validation

import (
	"io/ioutil"

	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/stdio/input"
	"github.com/wzshiming/pipe/pipe/tls"
)

const name = "validation"

func init() {
	decode.Register(name, NewValidationWithConfig)
}

type Config struct {
	Ca input.Input
}

func NewValidationWithConfig(conf *Config) (tls.TLS, error) {
	ca, err := ioutil.ReadAll(conf.Ca)
	if err != nil {
		return nil, err
	}
	err = conf.Ca.Close()
	if err != nil {
		return nil, err
	}

	tlsConfig, err := NewValidation(ca)
	if err != nil {
		return nil, err
	}
	return tls.WrapTLS(tlsConfig), nil
}

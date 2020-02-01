package merge

import (
	"crypto/tls"
)

func NewMerge(config []*tls.Config) *tls.Config {
	switch len(config) {
	case 0:
		return &tls.Config{}
	case 1:
		return config[0]
	}

	n := config[0].Clone()
	for _, v := range config[1:] {
		if v.RootCAs != nil && n.RootCAs == nil {
			n.RootCAs = v.RootCAs
		}
		if v.ClientCAs != nil && n.ClientCAs == nil {
			n.ClientCAs = v.ClientCAs
		}
		if v.ServerName != "" && n.ServerName == "" {
			n.ServerName = v.ServerName
		}
		if len(v.Certificates) != 0 {
			n.Certificates = append(n.Certificates, v.Certificates...)
		}
	}

	return n
}

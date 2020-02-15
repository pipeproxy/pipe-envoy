package access_log

import (
	"context"
	"net"

	"github.com/wzshiming/envoy/internal/client/access_log"
	"github.com/wzshiming/envoy/internal/node"
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/once"
	"github.com/wzshiming/pipe/pipe/stream/dialer"
)

const (
	name = "access_log"
)

func init() {
	decode.Register(name, NewAccessLogWithConfig)
}

type Config struct {
	Name    string `json:"@Name"`
	NodeID  string
	LogName string
	Dialer  dialer.Dialer
}

var accessLogMap = map[string]*AccessLog{}

func NewAccessLogWithConfig(conf *Config) (once.Once, error) {
	if a, ok := accessLogMap[conf.Name]; ok {
		return a, nil
	}

	accessLogConfig := access_log.Config{
		NodeConfig: &node.Config{
			NodeID: conf.NodeID,
		},
		ContextDialer: func(ctx context.Context, s string) (conn net.Conn, err error) {
			return conf.Dialer.DialStream(ctx)
		},
	}

	a, err := NewAccessLog(conf.LogName, 32, &accessLogConfig)
	if err != nil {
		return nil, err
	}

	accessLogMap[conf.Name] = a
	return a, nil
}

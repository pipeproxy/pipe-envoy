package access_log

import (
	"context"
	"net"

	"github.com/wzshiming/envoy/internal/client/access_log"

	"github.com/wzshiming/envoy/internal/node"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
	"github.com/wzshiming/pipe/stream"
)

const (
	name = "access_log"
)

func init() {
	configure.Register(name, NewAccessLogWithConfig)
}

type Config struct {
	Name    string `json:"@Name"`
	NodeID  string
	LogName string
	Forward stream.Handler
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
			p1, p2 := net.Pipe()
			go conf.Forward.ServeStream(ctx, p1)
			return p2, nil
		},
	}

	a, err := NewAccessLog(conf.LogName, 32, &accessLogConfig)
	if err != nil {
		return nil, err
	}

	accessLogMap[conf.Name] = a
	return a, nil
}

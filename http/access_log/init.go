package access_log

import (
	"fmt"
	"net/http"

	"github.com/wzshiming/envoy/once/access_log"
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/once"
)

const (
	name = "access_log"
)

func init() {
	configure.Register(name, NewAccessLogWithConfig)
}

type Config struct {
	AccessLog once.Once
	Handler   http.Handler
}

func NewAccessLogWithConfig(conf *Config) (http.Handler, error) {
	a, ok := conf.AccessLog.(*access_log.AccessLog)
	if !ok || a == nil {
		return nil, fmt.Errorf("need AccessLog")
	}
	return a.Handler(conf.Handler), nil
}

package none

import (
	"github.com/wzshiming/pipe/configure"
	"github.com/wzshiming/pipe/service"
)

const name = "none"

func init() {
	configure.Register(name, NewNoneWithConfig)
}

func NewNoneWithConfig() service.Service {
	return newNone()
}

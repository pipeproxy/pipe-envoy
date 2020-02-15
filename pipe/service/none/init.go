package none

import (
	"github.com/wzshiming/pipe/configure/decode"
	"github.com/wzshiming/pipe/pipe/service"
)

const name = "none"

func init() {
	decode.Register(name, NewNoneWithConfig)
}

func NewNoneWithConfig() service.Service {
	return newNone()
}

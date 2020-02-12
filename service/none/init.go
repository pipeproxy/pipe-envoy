package none

import (
	"github.com/wzshiming/pipe/configure/manager"
	"github.com/wzshiming/pipe/service"
)

const name = "none"

func init() {
	manager.Register(name, NewNoneWithConfig)
}

func NewNoneWithConfig() service.Service {
	return newNone()
}

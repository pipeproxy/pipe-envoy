package utils

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
)

func Duration(p *duration.Duration) (time.Duration, error) {
	if p == nil {
		return 0, nil
	}
	return ptypes.Duration(p)
}

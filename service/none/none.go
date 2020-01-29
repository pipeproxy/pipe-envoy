package none

import (
	"context"
)

type none struct {
	ch chan struct{}
}

func newNone() *none {
	return &none{
		ch: make(chan struct{}),
	}
}

func (n *none) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
	case <-n.ch:
	}
	return nil
}

func (n *none) Close() error {
	select {
	case n.ch <- struct{}{}:
	default:
	}
	return nil
}

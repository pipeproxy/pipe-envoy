package logger

import (
	"io"
)

// nopCloser
type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

// emptyWriter
type emptyWriter struct{}

func (emptyWriter) Write(p []byte) (n int, err error) { return len(p), nil }

func (emptyWriter) Close() (err error) { return nil }

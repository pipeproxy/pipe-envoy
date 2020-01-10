package admin

import (
	"context"
	"io"
	"net"
	"net/http"

	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bootstrap/address"
	"github.com/wzshiming/envoy/internal/logger"
)

type Admin struct {
	log     io.WriteCloser
	address address.Address
	mux     *http.ServeMux
}

func NewAdmin(config *envoy_config_bootstrap_v2.Admin) (*Admin, error) {
	log, err := logger.NewLogger(config.AccessLogPath)
	if err != nil {
		return nil, err
	}

	if config.ProfilePath != "" {
		logger.Todo("profile", config.ProfilePath)
	}

	address, err := address.NewAddress(config.Address)
	if err != nil {
		return nil, err
	}

	if len(config.SocketOptions) != 0 {
		logger.Todo("SocketOptions", config.SocketOptions)
	}
	return &Admin{
		log:     log,
		address: address,
	}, nil
}

func (a *Admin) initMux() {
	a.mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		logger.Todo("admin", r.URL.String())
	})
}

func (a *Admin) Start() error {
	ctx := context.Background()
	listener, err := a.address.Listen(ctx)
	if err != nil {
		return err
	}

	logger.Infoln("listen to", a.address)
	go func() {
		err := a.serve(listener)
		if err != nil {
			logger.Errorln(err)
		}
	}()
	return nil
}

func (a *Admin) serve(listener net.Listener) error {
	return http.Serve(listener, a.mux)
}

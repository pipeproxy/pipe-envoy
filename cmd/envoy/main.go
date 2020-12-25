package main

import (
	"context"
	"os"
	"syscall"

	_ "github.com/cncf/udpa/go/udpa/type/v1"
	_ "github.com/pipeproxy/pipe-xds/internal/convert"
	"github.com/pipeproxy/pipe-xds/internal/proxy/adapter"
	"github.com/spf13/pflag"
	"github.com/wzshiming/logger"
	"github.com/wzshiming/logger/zap"
	"github.com/wzshiming/notify"
)

var conf adapter.Config

var (
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	pflag.CommandLine.Init(os.Args[0], pflag.ContinueOnError)
	pflag.StringVarP(&conf.ConfigFile, "config", "c", "etc/istio/proxy/envoy-rev0.json", "Path to configuration file")

	pflag.StringVar(&conf.BasePath, "base-path", "etc/istio/proxy/pipe", "Path to pipe configuration dir")
	pflag.Parse()

	logger.SetLogger(zap.New(zap.WriteTo(os.Stderr), zap.UseDevMode(true)).WithName("adapter"))
	notify.OnSlice([]os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}, cancel)
}

func main() {
	log := logger.FromContext(ctx)

	adap, err := adapter.NewAdapter(&conf)
	if err != nil {
		log.Error(err, "new adapter")
		return
	}

	log.Info("start")
	err = adap.Run(ctx)
	if err != nil {
		log.Error(err, "start adapter")
		return
	}
}

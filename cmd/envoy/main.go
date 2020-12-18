package main

import (
	"context"
	"os"
	"syscall"

	_ "github.com/cncf/udpa/go/udpa/type/v1"
	"github.com/pipeproxy/pipe-xds/internal/adapter"
	_ "github.com/pipeproxy/pipe-xds/internal/convert"
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
	pflag.Uint32Var(&conf.RestartEpoch, "restart-epoch", 0, "Hot restart epoch")
	pflag.Uint32Var(&conf.DrainTimeS, "drain-time-s", 0, "Hot restart and LDS removal drain time in seconds")
	pflag.Uint32Var(&conf.ParentShutdownTimeS, "parent-shutdown-time-s", 0, "Hot restart parent shutdown time in seconds")
	pflag.StringVar(&conf.ServiceCluster, "service-cluster", "", "Cluster name")
	pflag.StringVar(&conf.ServiceNone, "service-node", "", "None name")
	pflag.StringVar(&conf.LocalAddressIPVersion, "local-address-ip-version", "", "The local IP address version (v4 or v6).")
	pflag.StringVar(&conf.BootstrapVersion, "bootstrap-version", "", "API version to parse the bootstrap config as (e.g. 3). If unset, all known versions will be attempted")

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

	err = adap.Start(ctx)
	if err != nil {
		log.Error(err, "start adapter")
		return
	}
	<-ctx.Done()
}

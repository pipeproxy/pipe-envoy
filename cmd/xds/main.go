package main

import (
	"context"
	"encoding/json"
	"os"
	"syscall"

	"github.com/pipeproxy/pipe-xds/internal/proxy/direct"
	"github.com/spf13/pflag"
	"github.com/wzshiming/logger"
	"github.com/wzshiming/logger/zap"
	"github.com/wzshiming/notify"
)

var (
	conf        direct.Config
	ctx, cancel = context.WithCancel(context.Background())
)

func init() {
	pflag.StringVarP(&conf.XDSAddr, "xds-address", "u", conf.XDSAddr, "xds server")
	pflag.StringVarP(&conf.NodeID, "node-id", "n", conf.NodeID, "node id")
	metadataJSON := "{}"
	pflag.StringVarP(&metadataJSON, "metadata", "m", metadataJSON, "node metadata")
	pflag.StringVar(&conf.BasePath, "base-path", "tmp", "Path to pipe configuration dir")
	pflag.Parse()

	logger.SetLogger(zap.New(zap.WriteTo(os.Stderr), zap.UseDevMode(true)).WithName("xds"))
	notify.OnSlice([]os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}, cancel)

	err := json.Unmarshal([]byte(metadataJSON), &conf.Metadata)
	if err != nil {
		logger.Log.Error(err, "Unmarshal")
	}
}

func main() {
	log := logger.FromContext(ctx)
	adap, err := direct.NewDirect(&conf)
	if err != nil {
		log.Error(err, "new direct")
		return
	}

	log.Info("start")
	err = adap.Run(ctx)
	if err != nil {
		log.Error(err, "start direct")
		return
	}
}

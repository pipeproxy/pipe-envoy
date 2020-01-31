package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/spf13/pflag"
	"github.com/wzshiming/envoy/config"
	convert_config_bootstrap_v2 "github.com/wzshiming/envoy/convert/config/bootstrap/v2"
	_ "github.com/wzshiming/envoy/init"
	"github.com/wzshiming/envoy/internal/logger"
	"github.com/wzshiming/pipe"
)

var conf string
var debug bool

func init() {
	pflag.StringVarP(&conf, "config", "c", "", "")
	pflag.BoolVarP(&debug, "debug", "d", false, "")
	if debug {
		logger.Debug()
	}
	for _, arg := range pflag.Args() {
		logger.Infoln(arg)
	}
	pflag.Parse()
}

func main() {
	data, err := ioutil.ReadFile(conf)
	if err != nil {
		logger.Fatalln(err)
	}

	ctx, conf, err := convertXDS(context.Background(), data)
	if err != nil {
		logger.Fatalln(err)
	}

	pipe, err := pipe.NewPipeWithConfig(ctx, conf)
	if err != nil {
		logger.Fatalln(err)
	}

	err = pipe.Run()
	if err != nil {
		logger.Fatalln(err)
	}

	return
}

func convertXDS(ctx context.Context, data []byte) (context.Context, []byte, error) {
	conf, err := config.UnmarshalBootstrap(data)
	if err != nil {
		return nil, nil, err
	}

	c := config.NewConfigCtx(ctx)

	_, err = convert_config_bootstrap_v2.Convert_Bootstrap(c, conf)
	if err != nil {
		return nil, nil, err
	}

	pipeConfig, err := json.Marshal(c)
	if err != nil {
		return nil, nil, err
	}

	return c.Ctx(), pipeConfig, nil
}

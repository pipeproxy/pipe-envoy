package main

import (
	"context"
	"io/ioutil"

	_ "github.com/wzshiming/envoy/init"

	"github.com/spf13/pflag"
	"github.com/wzshiming/envoy/convert"
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

	ctx, conf, err := convert.ConvertXDS(context.Background(), data)
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

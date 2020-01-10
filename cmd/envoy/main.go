package main

import (
	"io/ioutil"

	"github.com/spf13/pflag"
	"github.com/wzshiming/envoy/bootstrap"
	"github.com/wzshiming/envoy/internal/logger"
)

var config string
var debug bool

func init() {
	pflag.StringVarP(&config, "config", "c", "", "")
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

	data, err := ioutil.ReadFile(config)
	if err != nil {
		logger.Fatalln(err)
	}

	config, err := bootstrap.UnmarshalBootstrap(data)
	if err != nil {
		logger.Fatalln(err)
	}

	boot, err := bootstrap.NewBootstrap(config)
	if err != nil {
		logger.Fatalln(err)
	}

	err = boot.Start()
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infoln("started")
	<-make(chan struct{})
}

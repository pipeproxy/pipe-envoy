package main

import (
	"flag"
	"io/ioutil"
	"reflect"

	_ "github.com/wzshiming/envoy/init"

	"github.com/wzshiming/pipe/build"
	"github.com/wzshiming/pipe/configure/manager"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		b := build.NewBuild("bind")
		manager.ForEach(func(typ, kind string, out0Type reflect.Type, fun reflect.Value) {
			b.Add(kind, out0Type, fun)
		})
	} else {
		b := build.NewBuild(args[0])
		manager.ForEach(func(typ, kind string, out0Type reflect.Type, fun reflect.Value) {
			b.Add(kind, out0Type, fun)
		})
		ioutil.WriteFile(args[0]+".go", b.Bytes(), 0655)
	}
}

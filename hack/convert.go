//+build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/wzshiming/envoy/internal/logger"
	"github.com/wzshiming/gotype"
)

func init() {
	flag.Parse()
}

func main() {
	for _, arg := range flag.Args() {
		run(arg)
	}
}

func run(imp string) {

	impprter := gotype.NewImporter()

	typ, err := impprter.Import(imp, "")
	if err != nil {
		logger.Fatal(err)
	}

	num := typ.NumChild()
	pkg := typ.Name()
	logger.Info(imp, " ", num)

	for i := 0; i != num; i++ {
		child := typ.Child(i)
		name := child.Name()

		if 'A' > name[0] || 'Z' < name[0] || strings.Contains(name, "DeprecatedV1") || strings.HasSuffix(name, "ValidationError") || child.Kind() != gotype.Struct || child.NumField() <= 3 {
			continue
		}

		path := filepath.Join(strings.TrimPrefix(imp, "github.com/envoyproxy/go-control-plane/envoy/"), strings.ToLower(name)+".go")
		path = strings.ReplaceAll(path, "/", "_")

		generateFile(name, pkg, imp, path)
	}
}

func generateFile(name, pkg, imp, path string) error {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, `

package convert

import (
    %s %q
    "github.com/wzshiming/envoy/config"
    "github.com/wzshiming/envoy/internal/logger"
)

func Convert_%s_%s(conf *config.ConfigCtx, c *%s.%s) (string, error) {
    logger.Todof("%%#v", c)
	return "", nil
}

`, pkg, imp, strings.TrimPrefix(pkg, "envoy_"), name, pkg, name)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	d, err := os.Stat(path)
	if err == nil {
		if d.Size() != 0 {
			return nil
		}
	}
	src := buf.Bytes()
	src0, err := format.Source(src)
	if err == nil {
		src = src0
	}
	err = ioutil.WriteFile(path, src, 0644)
	if err != nil {
		return err
	}
	return nil
}

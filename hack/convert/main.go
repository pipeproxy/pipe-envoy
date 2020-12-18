package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

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
		log.Fatal(err)
	}

	num := typ.NumChild()
	pkg := typ.Name()

	for i := 0; i != num; i++ {
		child := typ.Child(i)
		name := child.Name()

		if 'A' > name[0] || 'Z' < name[0] || strings.Contains(name, "Deprecated") || strings.HasSuffix(name, "ValidationError") || child.Kind() != gotype.Struct {
			continue
		}
		_, ok := child.MethodByName("ProtoMessage")
		if !ok {
			continue
		}

		path := filepath.Join(strings.TrimPrefix(imp, "github.com/envoyproxy/go-control-plane/envoy/"), strings.ToLower(name)+".go")
		path = "xds_" + strings.ReplaceAll(path, "/", "_")

		generateFile(name, pkg, imp, path)
	}
}

func generateFile(name, pkg, imp, path string) error {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, `

package convert

import (
	"log"

	%s %q
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
)

func Convert_%s_%s(conf *config.ConfigCtx, c *%s.%s) (string, error) {
	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] %s.%s %%s\n", string(data))
	return "", nil
}

`, pkg, imp, strings.TrimPrefix(pkg, "envoy_"), name, pkg, name, pkg, name)
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

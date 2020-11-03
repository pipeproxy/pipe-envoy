package config

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"syscall"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pipeproxy/pipe-xds/encoding"
	"github.com/pipeproxy/pipe-xds/internal/adsc"
	"github.com/pipeproxy/pipe/bind"
	"github.com/pipeproxy/pipe/config"
	"sigs.k8s.io/yaml"
)

const (
	configFile = "pipe.yml"
	pidFile    = "pipe.pid"
)

type ConfigCtx struct {
	cli      *adsc.ADSC
	ctx      context.Context
	basePath string
	cds      map[string]bind.StreamDialer
	eds      map[string]bind.StreamDialer
	lds      map[string]bind.Service
	rds      map[string]bind.HTTPHandler
	sds      map[string]bind.TLS
	xds      map[string]proto.Message
	updateCh chan struct{}
}

func NewConfigCtx(ctx context.Context, cli *adsc.ADSC, basePath string) *ConfigCtx {
	return &ConfigCtx{
		cli:      cli,
		ctx:      ctx,
		basePath: basePath,
		cds:      map[string]bind.StreamDialer{},
		eds:      map[string]bind.StreamDialer{},
		lds:      map[string]bind.Service{},
		rds:      map[string]bind.HTTPHandler{},
		sds:      map[string]bind.TLS{},
		xds:      map[string]proto.Message{},
		updateCh: make(chan struct{}, 1),
	}
}

func (c *ConfigCtx) update() {
	select {
	case c.updateCh <- struct{}{}:
	default:
	}
}

func (c *ConfigCtx) RegisterADS(name string, dialer bind.StreamDialer) {
	//c.adsc[name] = ads
	c.update()
}

func (c *ConfigCtx) SetNodeID(nodeid string) {
	c.cli.NodeID = nodeid
}

func (c *ConfigCtx) RegisterCDS(name string, dialer bind.StreamDialer, msg proto.Message) {
	c.cds[name] = dialer
	c.xds[name] = msg
	c.update()
}

func (c *ConfigCtx) RegisterEDS(name string, dialer bind.StreamDialer, msg proto.Message) {
	c.eds[name] = dialer
	c.xds[name] = msg
	c.update()
}

func (c *ConfigCtx) RegisterLDS(name string, service bind.Service, msg proto.Message) {
	c.lds[name] = service
	c.xds[name] = msg
	c.update()
}

func (c *ConfigCtx) RegisterRDS(name string, handler bind.HTTPHandler, msg proto.Message) {
	c.rds[name] = handler
	c.xds[name] = msg
	c.update()
}

func (c *ConfigCtx) RegisterSDS(name string, tls bind.TLS, msg proto.Message) {
	c.sds[name] = tls
	c.xds[name] = msg
	c.update()
}

func (c *ConfigCtx) save() {
	componentSortd := []sortd{}
	serviceSortd := []sortd{}

	for name, d := range c.cds {
		componentSortd = append(componentSortd, sortd{name, d})
	}
	for name, d := range c.eds {
		componentSortd = append(componentSortd, sortd{name, d})
	}
	for name, d := range c.lds {
		serviceSortd = append(serviceSortd, sortd{name, d})
	}
	for name, d := range c.rds {
		componentSortd = append(componentSortd, sortd{name, d})
	}
	for name, d := range c.sds {
		componentSortd = append(componentSortd, sortd{name, d})
	}

	components := make([]bind.Component, 0, len(componentSortd))
	sort.Slice(componentSortd, func(i, j int) bool {
		return componentSortd[i].Name < componentSortd[j].Name
	})
	for _, com := range componentSortd {
		f := com.Name + ".yml"
		c.writeFile(f, com.Component, c.xds[com.Name])

		var d bind.Component
		switch com.Component.(type) {
		case bind.StreamDialer:
			d = bind.LoadStreamDialerConfig{Load: bind.FileIoReaderConfig{Path: f}}
		case bind.HTTPHandler:
			d = bind.LoadNetHTTPHandlerConfig{Load: bind.FileIoReaderConfig{Path: f}}
		case bind.TLS:
			d = bind.LoadTLSConfig{Load: bind.FileIoReaderConfig{Path: f}}
		}
		components = append(components, d)
	}

	services := make([]bind.Service, 0, len(serviceSortd))
	sort.Slice(serviceSortd, func(i, j int) bool {
		return serviceSortd[i].Name < serviceSortd[j].Name
	})
	for _, com := range serviceSortd {
		f := com.Name + ".yml"
		c.writeFile(f, com.Component, c.xds[com.Name])

		if reflect.DeepEqual(com.Component, bind.NoneService{}) {
			continue
		}

		var d bind.Service
		switch com.Component.(type) {
		case bind.Service:

			d = bind.LoadServiceConfig{Load: bind.FileIoReaderConfig{Path: f}}
		}
		services = append(services, d)
	}

	services = append(services, defaultServices...)

	d := bind.MultiOnceConfig{
		Multi: []bind.Once{
			bind.ServiceOnceConfig{
				Service: bind.MultiServiceConfig{
					Multi: services,
				},
			},
			bind.ComponentsOnceConfig{
				Components: components,
			},
		},
	}

	c.writeFile(configFile, d, nil)
}

func (c *ConfigCtx) Run(ctx context.Context) error {
	if !c.existFile(configFile) {
		c.save()
	}
	c.deleteFile(pidFile)
	return c.startPipe(ctx)
}

func (c *ConfigCtx) startPipe(ctx context.Context) error {
	cmd := exec.Command("pipe", "-c", configFile, "-p", pidFile)
	cmd.Dir = c.basePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
	loop:
		for {
			select {
			case <-c.updateCh:
				for {
					select {
					case <-c.updateCh:
					case <-time.After(time.Second / 10):
						c.save()
						cmd.Process.Signal(syscall.SIGHUP)
						continue loop
					}
				}
			case <-ctx.Done():
				cmd.Process.Signal(syscall.SIGQUIT)
				break loop
			}
		}
	}()
	return cmd.Wait()
}

type sortd struct {
	Name      string
	Component bind.Component
}

func (c *ConfigCtx) existFile(name string) bool {
	file := filepath.Join(c.basePath, name)
	_, err := os.Stat(file)
	return err == nil
}

func (c *ConfigCtx) deleteFile(name string) {
	file := filepath.Join(c.basePath, name)
	os.Remove(file)
	return
}

func (c *ConfigCtx) writeFile(name string, com bind.Component, msg proto.Message) {
	data, _ := yaml.Marshal(com)
	file := filepath.Join(c.basePath, name)
	if msg != nil {
		commit, err := encoding.Marshal(msg)
		if err != nil {
			log.Println(err)
		}
		commit, _ = yaml.JSONToYAML(commit)
		commit = bytes.ReplaceAll(commit, []byte{'\n'}, comm)
		data = append(data, comm...)
		data = append(data, commit...)
	}
	ioutil.WriteFile(file, data, 0644)
}

var comm = []byte{'\n', '#', ' '}

var defaultServices = []bind.Service{
	bind.StreamServiceConfig{
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
			Address: ":15021",
		},
		Handler: bind.HTTP1StreamHandlerConfig{
			Handler: BuildHealthWithHTTPHandler(),
		},
	},
	bind.StreamServiceConfig{
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
			Address: ":15090",
		},
		Handler: bind.HTTP1StreamHandlerConfig{
			Handler: BuildPrometheusWithHTTPHandler(),
		},
	},
	bind.StreamServiceConfig{
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigListenerNetworkEnumEnumTCP,
			Address: ":15000",
		},
		Handler: bind.HTTP1StreamHandlerConfig{
			Handler: BuildAdminWithHTTPHandler(),
		},
	},
}

func BuildAdminWithHTTPHandler() bind.HTTPHandler {
	return bind.MuxNetHTTPHandlerConfig{
		Routes: []bind.MuxNetHTTPHandlerRoute{
			{
				Path: "/",
				Handler: bind.MultiNetHTTPHandlerConfig{
					Multi: []bind.HTTPHandler{
						config.BuildContentTypeHTMLWithHTTPHandler(),
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `<pre>
<a href="pprof/">{{.Path}}pprof/</a>
<a href="expvar">{{.Path}}expvar</a>
<a href="must_quit">{{.Path}}must_quit</a>
<a href="healthz/ready">{{.Path}}healthz/ready</a>
<a href="stats/prometheus">{{.Path}}stats/prometheus</a>
<a href="config_dump">{{.Path}}config_dump</a>
<a href="config_dump_edit.sh">{{.Path}}config_dump_edit.sh</a>
</pre>`,
							},
						},
					},
				},
			},
			{
				Prefix:  "/pprof/",
				Handler: bind.PprofNetHTTPHandler{},
			},
			{
				Path:    "/expvar",
				Handler: bind.ExpvarNetHTTPHandler{},
			},
			{
				Path:    "/must_quit",
				Handler: bind.QuitNetHTTPHandler{},
			},
			{
				Path:    "/config_dump",
				Handler: bind.ConfigDumpNetHTTPHandlerConfig{},
			},
			{
				Path: "/config_dump_edit.sh",
				Handler: bind.MultiNetHTTPHandlerConfig{
					Multi: []bind.HTTPHandler{
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `#!/bin/sh
URL="{{.Scheme}}://{{.Host}}"
RESOURCE="$URL/config_dump"
TMP=.pipe_edit_tmp_file.yaml

# Check if editing is allowed
curl -sL -v -X OPTIONS "$RESOURCE" 2>&1 | \
grep "< Allow:" | grep "PUT" > /dev/null || \
{ echo "Editing Not Allowed"; exit 1;}

# Editing
curl -sL "$RESOURCE?yaml" > $TMP && \
vi $TMP && \
curl -sL -X PUT "$RESOURCE" -d "$(cat $TMP)" && \
rm $TMP

# sh -c "$(curl -sL {{.Scheme}}://{{.Host}}{{.Path}})"
`,
							},
						},
					},
				},
			},
		},
	}
}

func BuildHealthWithHTTPHandler() bind.HTTPHandler {
	return bind.MuxNetHTTPHandlerConfig{
		Routes: []bind.MuxNetHTTPHandlerRoute{
			{
				Path: "/healthz/ready",
				Handler: bind.DirectNetHTTPHandlerConfig{
					Code: http.StatusOK,
					Body: bind.InlineIoReaderConfig{
						Data: ``,
					},
				},
			},
		},
	}
}

func BuildPrometheusWithHTTPHandler() bind.HTTPHandler {
	return bind.MuxNetHTTPHandlerConfig{
		Routes: []bind.MuxNetHTTPHandlerRoute{
			{
				Path:    "/stats/prometheus",
				Handler: bind.MetricsNetHTTPHandler{},
			},
		},
	}
}

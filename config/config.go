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
	"sync"
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
	interval time.Duration
	mux      sync.Mutex
}

func NewConfigCtx(ctx context.Context, cli *adsc.ADSC, basePath string, interval time.Duration) *ConfigCtx {
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
		interval: interval,
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

func (c *ConfigCtx) RegisterCDS(name string, dialer bind.StreamDialer, msg proto.Message) bind.StreamDialer {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "cds." + name
	c.cds[name] = bind.DefStreamDialerConfig{
		Name: name,
		Def:  dialer,
	}
	c.xds[name] = msg
	c.update()
	return bind.RefStreamDialerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) CDS(name string) bind.StreamDialer {
	name = "cds." + name
	return bind.RefStreamDialerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) RegisterEDS(name string, dialer bind.StreamDialer, msg proto.Message) bind.StreamDialer {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "eds." + name
	c.eds[name] = bind.DefStreamDialerConfig{
		Name: name,
		Def:  dialer,
	}
	c.xds[name] = msg
	c.update()
	return bind.RefStreamDialerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) EDS(name string) bind.StreamDialer {
	name = "eds." + name
	return bind.RefStreamDialerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) RegisterLDS(name string, service bind.Service, msg proto.Message) bind.Service {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "lds." + name
	c.lds[name] = bind.DefServiceConfig{
		Name: name,
		Def:  service,
	}
	c.xds[name] = msg
	c.update()
	return bind.RefServiceConfig{
		Name: name,
	}
}

func (c *ConfigCtx) LDS(name string) bind.Service {
	name = "lds." + name
	return bind.RefServiceConfig{
		Name: name,
	}
}

func (c *ConfigCtx) RegisterRDS(name string, handler bind.HTTPHandler, msg proto.Message) bind.HTTPHandler {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "rds." + name
	c.rds[name] = bind.DefNetHTTPHandlerConfig{
		Name: name,
		Def:  handler,
	}
	c.xds[name] = msg
	c.update()
	return bind.RefNetHTTPHandlerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) RDS(name string) bind.HTTPHandler {
	name = "rds." + name
	return bind.RefNetHTTPHandlerConfig{
		Name: name,
	}
}

func (c *ConfigCtx) RegisterSDS(name string, tls bind.TLS, msg proto.Message) bind.TLS {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "sds." + name
	c.sds[name] = bind.DefTLSConfig{
		Name: name,
		Def:  tls,
	}
	c.xds[name] = msg
	c.update()
	return bind.RefTLSConfig{
		Name: name,
	}
}

func (c *ConfigCtx) SDS(name string) bind.TLS {
	name = "sds." + name
	return bind.RefTLSConfig{
		Name: name,
	}
}

func (c *ConfigCtx) save() {
	c.mux.Lock()
	defer c.mux.Unlock()
	componentSorted := make([]sorted, 0, len(c.cds)+len(c.eds)+len(c.rds)+len(c.sds))
	serviceSorted := make([]sorted, 0, len(c.lds))

	for name, d := range c.cds {
		componentSorted = append(componentSorted, sorted{name, d})
	}
	for name, d := range c.eds {
		componentSorted = append(componentSorted, sorted{name, d})
	}
	for name, d := range c.rds {
		componentSorted = append(componentSorted, sorted{name, d})
	}
	for name, d := range c.sds {
		componentSorted = append(componentSorted, sorted{name, d})
	}
	for name, d := range c.lds {
		serviceSorted = append(serviceSorted, sorted{name, d})
	}

	components := make([]bind.Component, 0, len(componentSorted))
	sort.Slice(componentSorted, func(i, j int) bool {
		return componentSorted[i].Name < componentSorted[j].Name
	})
	for _, com := range componentSorted {
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

	services := make([]bind.Service, 0, len(serviceSorted))
	sort.Slice(serviceSorted, func(i, j int) bool {
		return serviceSorted[i].Name < serviceSorted[j].Name
	})
	for _, svc := range serviceSorted {
		f := svc.Name + ".yml"
		c.writeFile(f, svc.Component, c.xds[svc.Name])

		if reflect.DeepEqual(svc.Component, bind.NoneService{}) {
			continue
		}

		switch s := svc.Component.(type) {
		default:
			services = append(services, bind.LoadServiceConfig{
				Load: bind.FileIoReaderConfig{
					Path: f,
				},
			})
		case *bind.DefServiceConfig:
			components = append(components, bind.LoadServiceConfig{
				Load: bind.FileIoReaderConfig{
					Path: f,
				},
			})
			services = append(services, bind.RefServiceConfig{
				Name: s.Name,
			})
		}
	}
	for _, svc := range defaultServices {
		components = append(components, svc)
		services = append(services, bind.RefServiceConfig{
			Name: svc.Name,
		})
	}
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
					case <-time.After(c.interval):
						c.save()
						cmd.Process.Signal(syscall.SIGHUP)
						continue loop
					case <-ctx.Done():
						cmd.Process.Signal(syscall.SIGQUIT)
						break loop
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

type sorted struct {
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

var defaultServices = []bind.DefServiceConfig{
	{
		Name: "_health",
		Def: bind.TagsServiceConfig{
			Tag: "health",
			Service: bind.StreamServiceConfig{
				Listener: bind.ListenerStreamListenConfigConfig{
					Network: bind.ListenerStreamListenConfigNetworkEnumEnumNetworkTCP,
					Address: ":15021",
				},
				Handler: bind.HTTP1StreamHandlerConfig{
					Handler: BuildHealthWithHTTPHandler(),
				},
			},
		},
	},
	{
		Name: "_metric",
		Def: bind.TagsServiceConfig{
			Tag: "metric",
			Service: bind.StreamServiceConfig{
				Listener: bind.ListenerStreamListenConfigConfig{
					Network: bind.ListenerStreamListenConfigNetworkEnumEnumNetworkTCP,
					Address: ":15090",
				},
				Handler: bind.HTTP1StreamHandlerConfig{
					Handler: BuildPrometheusWithHTTPHandler(),
				},
			},
		},
	},
	{
		Name: "_admin",
		Def: bind.TagsServiceConfig{
			Tag: "admin",
			Service: bind.StreamServiceConfig{
				Listener: bind.ListenerStreamListenConfigConfig{
					Network: bind.ListenerStreamListenConfigNetworkEnumEnumNetworkTCP,
					Address: ":15000",
				},
				Handler: bind.HTTP1StreamHandlerConfig{
					Handler: bind.LogNetHTTPHandlerConfig{
						Output: bind.FileIoWriterConfig{
							Path: "/dev/stderr",
						},
						Handler: BuildAdminWithHTTPHandler(),
					},
				},
			},
		},
	},
}

func BuildAdminWithHTTPHandler() bind.HTTPHandler {
	return bind.PathNetHTTPHandlerConfig{
		Paths: []bind.PathNetHTTPHandlerRoute{
			{
				Path: "/",
				Handler: bind.MultiNetHTTPHandlerConfig{
					Multi: []bind.HTTPHandler{
						config.BuildContentTypeHTMLWithHTTPHandler(),
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `
http://localhost:15000/pprof/
http://localhost:15000/expvar
http://localhost:15000/quitquitquit
http://localhost:15000/config_dump
http://localhost:15021/healthz/ready
http://localhost:15090/stats/prometheus
`,
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
				Path:    "/quitquitquit",
				Handler: bind.QuitNetHTTPHandler{},
			},
			{
				Path:    "/config_dump",
				Handler: bind.ConfigDumpNetHTTPHandlerConfig{},
			},
		},
	}
}

func BuildHealthWithHTTPHandler() bind.HTTPHandler {
	return bind.PathNetHTTPHandlerConfig{
		Paths: []bind.PathNetHTTPHandlerRoute{
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
	return bind.PathNetHTTPHandlerConfig{
		Paths: []bind.PathNetHTTPHandlerRoute{
			{
				Path:    "/stats/prometheus",
				Handler: bind.MetricsNetHTTPHandler{},
			},
		},
	}
}

package config

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/pipeproxy/pipe-xds/internal/encoding"
	"github.com/pipeproxy/pipe-xds/internal/proxy"
	"github.com/pipeproxy/pipe/bind"
	"github.com/wzshiming/logger"
	"sigs.k8s.io/yaml"
)

type ConfigCtx struct {
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

func NewConfigCtx(ctx context.Context, basePath string, interval time.Duration) *ConfigCtx {
	os.MkdirAll(basePath, 0755)
	return &ConfigCtx{
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

func (c *ConfigCtx) Update() {
	select {
	case c.updateCh <- struct{}{}:
	default:
	}
}

func (c *ConfigCtx) RegisterCDS(name string, dialer bind.StreamDialer, msg proto.Message) bind.StreamDialer {
	c.mux.Lock()
	defer c.mux.Unlock()
	name = "cds." + name
	i := 1
	for n := name; ; i++ {
		_, ok := c.cds[n]
		if !ok {
			name = n
			break
		}
		n = fmt.Sprintf("%s.%d", name, i)
	}
	c.cds[name] = bind.DefStreamDialerConfig{
		Name: name,
		Def:  dialer,
	}
	c.xds[name] = msg
	c.Update()
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
	i := 1
	for n := name; ; i++ {
		_, ok := c.eds[n]
		if !ok {
			name = n
			break
		}
		n = fmt.Sprintf("%s.%d", name, i)
	}
	c.eds[name] = bind.DefStreamDialerConfig{
		Name: name,
		Def:  dialer,
	}
	c.xds[name] = msg
	c.Update()
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
	i := 1
	for n := name; ; i++ {
		_, ok := c.lds[n]
		if !ok {
			name = n
			break
		}
		n = fmt.Sprintf("%s.%d", name, i)
	}
	c.lds[name] = bind.DefServiceConfig{
		Name: name,
		Def:  service,
	}
	c.xds[name] = msg
	c.Update()
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
	i := 1
	for n := name; ; i++ {
		_, ok := c.rds[n]
		if !ok {
			name = n
			break
		}
		n = fmt.Sprintf("%s.%d", name, i)
	}
	c.rds[name] = bind.DefNetHTTPHandlerConfig{
		Name: name,
		Def:  handler,
	}
	c.xds[name] = msg
	c.Update()
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
	i := 1
	for n := name; ; i++ {
		_, ok := c.sds[n]
		if !ok {
			name = n
			break
		}
		n = fmt.Sprintf("%s.%d", name, i)
	}
	c.sds[name] = bind.DefTLSConfig{
		Name: name,
		Def:  tls,
	}
	c.xds[name] = msg
	c.Update()
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

func (c *ConfigCtx) Save() {
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

	c.writeFile(proxy.ConfigFile, d, nil)
}

func (c *ConfigCtx) Watch(ctx context.Context, update func()) {
	go func() {
	loop:
		for {
			select {
			case <-c.updateCh:
				for {
					select {
					case <-c.updateCh:
					case <-time.After(c.interval):
						c.Save()
						update()
						continue loop
					case <-ctx.Done():
						break loop
					}
				}
			case <-ctx.Done():
				break loop
			}
		}
	}()
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
	file := filepath.Join(c.basePath, name)
	log := logger.FromContext(c.ctx)
	data, err := yaml.Marshal(com)
	if err != nil {
		log.Error(err, "Marshal")
	}
	file, err = filepath.Abs(file)
	if err != nil {
		log.Error(err, "Abs")
	}
	if msg != nil {
		commit, err := encoding.Marshal(msg)
		if err != nil {
			log.Error(err, "Marshal")
		}
		commit, err = yaml.JSONToYAML(commit)
		if err != nil {
			log.Error(err, "JSONToYAML")
		}
		commit = bytes.ReplaceAll(commit, []byte{'\n'}, comm)
		data = append(data, comm...)
		data = append(data, commit...)
	}
	err = ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Error(err, "WriteFile")
	}
}

var comm = []byte{'\n', '#', ' '}

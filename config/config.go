package config

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"sort"
	"sync"

	"github.com/wzshiming/envoy/bind"
)

type ConfigCtx struct {
	init         []bind.Once
	componentMap map[string]bind.PipeComponent
	eds          []string
	rds          []string
	sds          []string
	ctx          context.Context
	services     []string
	mut          sync.Mutex
}

func (c *ConfigCtx) MarshalJSON() ([]byte, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	conf := bind.PipeConfig{}

	switch len(c.services) {
	case 0:
		conf.Pipe = bind.ServiceNone{}
	case 1:
		conf.Pipe = bind.RefService(c.services[0])

	default:
		multi := make([]bind.Service, 0, len(c.services))
		for _, service := range c.services {
			multi = append(multi, bind.RefService(service))
		}

		conf.Pipe = bind.ServiceMultiConfig{
			Multi: multi,
		}
	}

	keys := make([]string, 0, len(c.componentMap))
	for key := range c.componentMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	conf.Components = make([]bind.PipeComponent, 0, len(c.componentMap))
	for _, key := range keys {
		conf.Components = append(conf.Components, c.componentMap[key])
	}

	conf.Init = c.init

	return json.Marshal(conf)
}

func (c *ConfigCtx) Ctx() context.Context {
	c.mut.Lock()
	defer c.mut.Unlock()

	return c.ctx
}

func (c *ConfigCtx) WithValue(key, val interface{}) context.Context {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.ctx = context.WithValue(c.ctx, key, val)
	return c.ctx
}

func (c *ConfigCtx) ResetEDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	eds := c.eds
	c.eds = nil
	return eds
}

func (c *ConfigCtx) ResetRDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	rds := c.rds
	c.rds = nil
	return rds
}

func (c *ConfigCtx) ResetSDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	sds := c.sds
	c.sds = nil
	return sds
}

func (c *ConfigCtx) EDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	return c.eds
}

func (c *ConfigCtx) RDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	return c.rds
}

func (c *ConfigCtx) SDS() []string {
	c.mut.Lock()
	defer c.mut.Unlock()

	return c.sds
}

func (c *ConfigCtx) AppendEDS(eds string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for _, v := range c.eds {
		if v == eds {
			return
		}
	}

	c.eds = append(c.eds, eds)
}

func (c *ConfigCtx) AppendRDS(rds string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for _, v := range c.rds {
		if v == rds {
			return
		}
	}

	c.rds = append(c.rds, rds)
}

func (c *ConfigCtx) AppendSDS(sds string) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for _, v := range c.sds {
		if v == sds {
			return
		}
	}

	c.sds = append(c.sds, sds)
}

func (c *ConfigCtx) RegisterComponents(name string, d bind.PipeComponent) (string, error) {

	c.mut.Lock()
	defer c.mut.Unlock()

	if name == "" {
		n, raw, err := marshalName(name, d)
		if err != nil {
			return "", err
		}
		name = n
		d = bind.RawPipeComponent(raw)
	}

	d = bind.NamePipeComponent{
		Name:          name,
		PipeComponent: d,
	}
	c.componentMap[name] = d
	return name, nil
}

func (c *ConfigCtx) RegisterService(name string) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	for _, services := range c.services {
		if services == name {
			return nil
		}
	}
	c.services = append(c.services, name)
	return nil
}

func (c *ConfigCtx) RegisterInit(d bind.Once) (string, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.init = append(c.init, d)
	return "", nil
}

func marshalName(name string, m json.Marshaler) (string, json.RawMessage, error) {
	d, err := m.MarshalJSON()
	if err != nil {
		return "", nil, err
	}
	hash := md5.Sum(d)
	name = "auto@" + hex.EncodeToString(hash[:])

	return name, d, nil
}

type xdsCtxKeyType int

func GetXdsWithContext(ctx context.Context) (*ConfigCtx, bool) {
	i := ctx.Value(xdsCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*ConfigCtx)
	return p, ok
}

func NewConfigCtx(ctx context.Context) *ConfigCtx {
	c := &ConfigCtx{
		ctx:          ctx,
		componentMap: map[string]bind.PipeComponent{},
	}
	ctx = context.WithValue(ctx, xdsCtxKeyType(0), c)
	c.ctx = ctx
	return c
}

func XdsName(name string) string {
	if name == "" {
		return ""
	}
	return "xds@" + name
}

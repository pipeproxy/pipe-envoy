package config

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
)

type Config struct {
	Pipe       json.RawMessage
	Init       []json.RawMessage `json:",omitempty"`
	Components []json.RawMessage `json:",omitempty"`
}

type ConfigCtx struct {
	init         []json.RawMessage
	componentMap map[string]json.RawMessage
	eds          []string
	rds          []string
	ctx          context.Context
	services     []string
	mut          sync.Mutex
}

func (c *ConfigCtx) MarshalJSON() ([]byte, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	conf := Config{}
	switch len(c.services) {
	case 0:
		conf.Pipe = []byte(`{"@Kind":"none"}`)
	case 1:

		pipe, err := MarshalRef(c.services[0])
		if err != nil {
			return nil, err
		}
		conf.Pipe = pipe

	default:
		list := []json.RawMessage{}
		for _, service := range c.services {
			ref, err := MarshalRef(service)
			if err != nil {
				return nil, err
			}
			list = append(list, ref)
		}

		pipe, err := MarshalKindServiceMulti(list)
		if err != nil {
			return nil, err
		}
		conf.Pipe = pipe
	}

	for _, component := range c.componentMap {
		conf.Components = append(conf.Components, component)
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

func (c *ConfigCtx) AppendEDS(eds string) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.eds = append(c.eds, eds)
}

func (c *ConfigCtx) AppendRDS(rds string) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.rds = append(c.rds, rds)
}

func (c *ConfigCtx) RegisterComponents(name string, d json.RawMessage) (string, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	n, d, err := MarshalName(name, d)
	if err != nil {
		return "", err
	}

	if c.componentMap == nil {
		c.componentMap = map[string]json.RawMessage{}
	}
	c.componentMap[n] = d
	return n, nil
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

func (c *ConfigCtx) RegisterInit(d json.RawMessage) (string, error) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.init = append(c.init, d)
	return "", nil
}

func MarshalKind(kind string, i interface{}) (json.RawMessage, error) {
	d, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	if len(d) > 2 && d[0] == '{' && d[1] != '}' {
		d = append([]byte(fmt.Sprintf(`{"@Kind":%q,`, kind)), d[1:]...)
	}

	return d, nil
}

func MarshalName(name string, d json.RawMessage) (string, json.RawMessage, error) {
	if name == "" {
		hash := md5.Sum(d)
		name = "auto@" + hex.EncodeToString(hash[:])
	}

	if len(d) > 2 && d[0] == '{' && d[1] != '}' {
		d = append([]byte(fmt.Sprintf(`{"@Name":%q,`, name)), d[1:]...)
	}

	return name, d, nil
}

func MarshalRef(ref string) (json.RawMessage, error) {
	return []byte(fmt.Sprintf(`{"@Ref":%q}`, ref)), nil
}

func NewConfigCtx(ctx context.Context) *ConfigCtx {
	return &ConfigCtx{
		ctx: ctx,
	}
}

package config

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Config struct {
	Pipe       json.RawMessage
	Init       []json.RawMessage `json:",omitempty"`
	Components []json.RawMessage `json:",omitempty"`
}

type ConfigCtx struct {
	Init         []json.RawMessage
	ComponentMap map[string]json.RawMessage
	EDS          []string
	RDS          []string
	Ctx          context.Context
	Services     []string
}

func (c ConfigCtx) MarshalJSON() ([]byte, error) {
	conf := Config{}
	switch len(c.Services) {
	case 0:
		conf.Pipe = []byte(`{"@Kind":"none"}`)
	case 1:

		pipe, err := MarshalRef(c.Services[0])
		if err != nil {
			return nil, err
		}
		conf.Pipe = pipe

	default:
		list := []json.RawMessage{}
		for _, service := range c.Services {
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

	for _, component := range c.ComponentMap {
		conf.Components = append(conf.Components, component)
	}
	conf.Init = c.Init

	return json.Marshal(conf)
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

func (c *ConfigCtx) RegisterComponents(name string, d json.RawMessage) (string, error) {
	n, d, err := MarshalName(name, d)
	if err != nil {
		return "", err
	}

	if c.ComponentMap == nil {
		c.ComponentMap = map[string]json.RawMessage{}
	}
	c.ComponentMap[n] = d
	return n, nil
}

func (c *ConfigCtx) RegisterService(name string) error {
	for _, services := range c.Services {
		if services == name {
			return nil
		}
	}
	c.Services = append(c.Services, name)
	return nil
}

func (c *ConfigCtx) RegisterInit(d json.RawMessage) (string, error) {
	c.Init = append(c.Init, d)
	return "", nil
}

package config

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type ConfigCtx struct {
	Pipe       json.RawMessage
	Init       []json.RawMessage `json:",omitempty"`
	Components []json.RawMessage `json:",omitempty"`

	EDS                 []string            `json:"-"`
	RDS                 []string            `json:"-"`
	Ctx                 context.Context     `json:"-"`
	Services            []string            `json:"-"`
	DuplicateComponents map[string]struct{} `json:"-"`
}

func (c ConfigCtx) MarshalJSON() ([]byte, error) {
	switch len(c.Services) {
	case 0:
		c.Pipe = []byte(`{"@Kind":"none"}`)
	case 1:

		pipe, err := MarshalRef(c.Services[0])
		if err != nil {
			return nil, err
		}
		c.Pipe = pipe

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
		c.Pipe = pipe
	}
	type tmp ConfigCtx
	return json.Marshal(tmp(c))
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

	if c.DuplicateComponents == nil {
		c.DuplicateComponents = map[string]struct{}{}
	}
	if _, ok := c.DuplicateComponents[n]; ok {
		return n, nil
	}
	c.DuplicateComponents[n] = struct{}{}

	c.Components = append(c.Components, d)
	return n, nil
}

func (c *ConfigCtx) RegisterService(name string) error {
	c.Services = append(c.Services, name)
	return nil
}

func (c *ConfigCtx) RegisterInit(d json.RawMessage) (string, error) {
	c.Init = append(c.Init, d)
	return "", nil
}

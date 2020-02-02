package clean

import (
	"bytes"
	"encoding/json"
)

type node struct {
	Config json.RawMessage
	Name   string
	Refs   []string
}
type dependencies struct {
	m map[string]*node
}

func newDependencies() *dependencies {
	d := &dependencies{
		m: map[string]*node{},
	}
	return d
}

func (d *dependencies) Reset() {
	d.m = map[string]*node{}
}

func (d *dependencies) AppendNode(name string, refs []string, config []byte) {
	d.m[name] = &node{
		Name:   name,
		Refs:   refs,
		Config: config,
	}
}

func (d *dependencies) Nodes() map[string]*node {
	return d.m
}

func (d *dependencies) Decode(config []byte) error {
	rs, err := d.decode(config, 0)
	if err != nil {
		return err
	}

	d.AppendNode("", rs, nil)
	return nil
}

func (d *dependencies) decode(config []byte, deep int) ([]string, error) {
	config = bytes.TrimSpace(config)
	if len(config) == 0 {
		return nil, nil
	}
	switch config[0] {
	case '[':
		return d.decodeSlice(config, deep)
	case '{':
		return d.decodeMap(config, deep+1)
	}
	return nil, nil
}

func (d *dependencies) decodeMap(config []byte, deep int) ([]string, error) {

	var field struct {
		Name string `json:"@Name"`
		Ref  string `json:"@Ref"`
	}

	err := json.Unmarshal(config, &field)
	if err != nil {
		return nil, err
	}

	data := map[string]json.RawMessage{}
	err = json.Unmarshal(config, &data)
	if err != nil {
		return nil, err
	}

	if field.Ref != "" {
		if field.Name != "" {
			if field.Ref != "" {
				d.AppendNode(field.Name, []string{field.Ref}, config)
				return nil, nil
			}
		} else {
			return []string{field.Ref}, nil
		}
	}

	rs := []string{}
	for _, raw := range data {
		r, err := d.decode(raw, deep)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r...)
	}

	if field.Name == "" {
		return rs, nil
	}

	d.AppendNode(field.Name, rs, config)

	return nil, nil
}

func (d *dependencies) decodeSlice(config []byte, deep int) ([]string, error) {
	data := []json.RawMessage{}
	err := json.Unmarshal(config, &data)
	if err != nil {
		return nil, err
	}

	rs := []string{}
	for _, raw := range data {
		r, err := d.decode(raw, deep)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r...)
	}
	return rs, nil

}

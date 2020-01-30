package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"sigs.k8s.io/yaml"
)

var (
	marshaler   = &jsonpb.Marshaler{OrigName: true, Indent: "  "}
	unmarshaler = jsonpb.Unmarshaler{
		AnyResolver: ResolveFunc(defaultResolveAny),
	}
)

func Marshal(pb proto.Message) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := marshaler.Marshal(buf, pb)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(data []byte, pb proto.Message) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return unmarshaler.UnmarshalNext(dec, pb)
}

func UnmarshalAny(a *any.Any) (proto.Message, error) {
	msg, err := defaultResolveAny(a.TypeUrl)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(a.Value, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func UnmarshalBootstrap(config []byte) (*envoy_config_bootstrap_v2.Bootstrap, error) {
	config, err := yaml.YAMLToJSON(config)
	if err != nil {
		return nil, err
	}

	bootstrap := &envoy_config_bootstrap_v2.Bootstrap{}

	err = Unmarshal(config, bootstrap)
	if err != nil {
		return nil, err
	}
	return bootstrap, nil
}

func MarshalBootstrap(bootstrap *envoy_config_bootstrap_v2.Bootstrap) ([]byte, error) {
	config, err := Marshal(bootstrap)
	if err != nil {
		return nil, err
	}

	return yaml.JSONToYAML(config)
}

type ResolveFunc func(typeUrl string) (proto.Message, error)

func (r ResolveFunc) Resolve(typeUrl string) (proto.Message, error) {
	return r(typeUrl)
}

func defaultResolveAny(typeUrl string) (proto.Message, error) {
	// Only the part of typeUrl after the last slash is relevant.
	mname := typeUrl
	if slash := strings.LastIndex(mname, "/"); slash >= 0 {
		mname = mname[slash+1:]
	}

	mt := proto.MessageType(mname)
	if mt == nil {
		return nil, fmt.Errorf("unknown message type %q", mname)
	}
	return reflect.New(mt.Elem()).Interface().(proto.Message), nil
}

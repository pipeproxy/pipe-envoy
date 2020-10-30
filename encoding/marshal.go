package encoding

import (
	"bytes"
	"encoding/json"
	"log"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/reflect/protoregistry"
	"sigs.k8s.io/yaml"
)

var (
	anyResolver = &dynamicAnyResolver{
		filter: map[string]struct{}{},
	}
	marshaler = &jsonpb.Marshaler{
		OrigName:    true,
		AnyResolver: anyResolver,
	}
	unmarshaler = jsonpb.Unmarshaler{
		AnyResolver: anyResolver,
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
	msg, err := anyResolver.Resolve(a.TypeUrl)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(a.Value, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func UnmarshalBootstrap(config []byte) (*envoy_config_bootstrap_v3.Bootstrap, error) {
	config, err := yaml.YAMLToJSON(config)
	if err != nil {
		return nil, err
	}

	bootstrap := &envoy_config_bootstrap_v3.Bootstrap{}

	err = Unmarshal(config, bootstrap)
	if err != nil {
		return nil, err
	}
	return bootstrap, nil
}

func MarshalBootstrap(bootstrap *envoy_config_bootstrap_v3.Bootstrap) ([]byte, error) {
	config, err := Marshal(bootstrap)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(config)
}

type dynamicAnyResolver struct {
	filter map[string]struct{}
}

func (d *dynamicAnyResolver) Resolve(typeURL string) (proto.Message, error) {
	mt, err := protoregistry.GlobalTypes.FindMessageByURL(typeURL)
	if err != nil {
		if _, ok := d.filter[typeURL]; !ok {
			log.Println(err, typeURL)
			d.filter[typeURL] = struct{}{}
		}
		return &empty.Empty{}, nil
	}
	return proto.MessageV1(mt.New().Interface()), nil
}

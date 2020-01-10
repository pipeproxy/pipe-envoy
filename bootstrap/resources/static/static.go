package static

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bootstrap/listener"
)

func NewStaticResources(config *envoy_config_bootstrap_v2.Bootstrap_StaticResources) (*StaticResources, error) {
	s := &StaticResources{}
	for _, l := range config.Listeners {
		listen, err := listener.NewListener(l)
		if err != nil {
			return nil, err
		}
		s.Listeners = append(s.Listeners, listen)
	}
	return s, nil
}

type StaticResources struct {
	Listeners []*listener.Listener
	//Clusters             []*v22.Cluster  `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	//Secrets              []*auth.Secret  `protobuf:"bytes,3,rep,name=secrets,proto3" json:"secrets,omitempty"`
}

func (s *StaticResources) Start() error {
	for _, listener := range s.Listeners {
		err := listener.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

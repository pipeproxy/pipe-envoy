package listener

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/wzshiming/envoy/bootstrap/address"
	"github.com/wzshiming/envoy/internal/logger"
)

func NewListener(config *envoy_api_v2.Listener) (*Listener, error) {
	l := &Listener{}
	l.name = config.Name

	address, err := address.NewAddress(config.Address)
	if err != nil {
		return nil, err
	}
	l.address = address
	logger.Todo("Listener", l)
	return l, nil
}

type Listener struct {
	name    string
	address address.Address
	// FilterChains                     []*listener.FilterChain
	//UseOriginalDst                   *wrappers.BoolValue               `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst,proto3" json:"use_original_dst,omitempty"` // Deprecated: Do not use.
	//PerConnectionBufferLimitBytes    *wrappers.UInt32Value             `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	//Metadata                         *core.Metadata                    `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	//DeprecatedV1                     *Listener_DeprecatedV1            `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	//DrainType                        Listener_DrainType                `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	//ListenerFilters                  []*listener.ListenerFilter        `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	//ListenerFiltersTimeout           *duration.Duration                `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	//ContinueOnListenerFiltersTimeout bool                              `protobuf:"varint,17,opt,name=continue_on_listener_filters_timeout,json=continueOnListenerFiltersTimeout,proto3" json:"continue_on_listener_filters_timeout,omitempty"`
	//Transparent                      *wrappers.BoolValue               `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	//Freebind                         *wrappers.BoolValue               `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	//SocketOptions                    []*core.SocketOption              `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	//TcpFastOpenQueueLength           *wrappers.UInt32Value             `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	//TrafficDirection                 core.TrafficDirection             `protobuf:"varint,16,opt,name=traffic_direction,json=trafficDirection,proto3,enum=envoy.api.v2.core.TrafficDirection" json:"traffic_direction,omitempty"`
	//UdpListenerConfig                *listener.UdpListenerConfig       `protobuf:"bytes,18,opt,name=udp_listener_config,json=udpListenerConfig,proto3" json:"udp_listener_config,omitempty"`
	//ApiListener                      *v2.ApiListener                   `protobuf:"bytes,19,opt,name=api_listener,json=apiListener,proto3" json:"api_listener,omitempty"`
	//ConnectionBalanceConfig          *Listener_ConnectionBalanceConfig `protobuf:"bytes,20,opt,name=connection_balance_config,json=connectionBalanceConfig,proto3" json:"connection_balance_config,omitempty"`
	//ReusePort                        bool                              `protobuf:"varint,21,opt,name=reuse_port,json=reusePort,proto3" json:"reuse_port,omitempty"`
}

func (l *Listener) Start() error {

	return nil

}

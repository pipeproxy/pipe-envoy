package convert

import (
	"fmt"
	"reflect"
	"sort"

	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
)

func Convert_config_listener_v3_Listener(conf *config.ConfigCtx, c *envoy_config_listener_v3.Listener, kind string) (bind.Service, error) {
	if len(c.FilterChains) == 0 {
		return bind.NoneService{}, nil
	}

	network, address, err := Convert_config_core_v3_Address(conf, c.Address)
	if err != nil {
		return nil, err
	}

	//TODO: Support dynamic selection of filter,
	s, err := SelectFilterChain(conf, c.FilterChains)
	if err != nil {
		return nil, err
	}

	s = bind.LogStreamHandlerConfig{
		Handler: s,
		Output: bind.FileIoWriterConfig{
			Path: "/dev/stderr",
		},
	}

	var d bind.Service
	d = bind.StreamServiceConfig{
		DisconnectOnClose: true,
		Listener: bind.ListenerStreamListenConfigConfig{
			Network: bind.ListenerStreamListenConfigNetworkEnum(network),
			Address: address,
			Virtual: c.DeprecatedV1 != nil && !c.DeprecatedV1.BindToPort.GetValue(),
		},
		Handler: s,
	}

	if c.Name != "" {
		name := c.Name
		if kind != "" {
			name = fmt.Sprintf("%s.%s", kind, name)
		}
		d = conf.RegisterLDS(name, bind.TagsServiceConfig{
			Service: d,
			Tag:     c.Name,
		}, c)
	} else if kind != "" {
		d = conf.RegisterLDS(kind, bind.TagsServiceConfig{
			Service: d,
			Tag:     kind,
		}, c)
	}
	return d, nil
}

var (
	plaintextHTTPALPNs = []string{"http/1.0", "http/1.1", "h2c"}
	istioHTTPPlaintext = []string{"istio", "istio-http/1.0", "istio-http/1.1", "istio-h2"}
	httpTLS            = []string{"http/1.0", "http/1.1", "h2c", "istio-http/1.0", "istio-http/1.1", "istio-h2"}
	tcpTLS             = []string{"istio-peer-exchange", "istio"}

	protDescrs = map[string][]string{
		"App: HTTP TLS":         httpTLS,
		"App: Istio HTTP Plain": istioHTTPPlaintext,
		"App: TCP TLS":          tcpTLS,
		"App: HTTP":             plaintextHTTPALPNs,
	}
)

func SelectFilterChain(conf *config.ConfigCtx, fc []*envoy_config_listener_v3.FilterChain) (bind.StreamHandler, error) {
	ports := map[uint32]*prefixMux{}
	ports[0] = &prefixMux{}
	for _, filter := range fc {
		if filter.FilterChainMatch == nil {
			h, err := Convert_config_listener_v3_FilterChain(conf, filter)
			if err != nil {
				return nil, err
			}
			ports[0].tcp.tcpStream = h
			continue
		}
		port := filter.FilterChainMatch.DestinationPort.GetValue()
		if ports[port] == nil {
			ports[port] = &prefixMux{}
		}
		if filter.FilterChainMatch.ApplicationProtocols == nil {
			h, err := Convert_config_listener_v3_FilterChain(conf, filter)
			if err != nil {
				return nil, err
			}
			ports[port].tcp.tcpStream = h
			continue
		}
		switch len(filter.FilterChainMatch.ApplicationProtocols) {
		case len(plaintextHTTPALPNs):
			if reflect.DeepEqual(plaintextHTTPALPNs, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				ports[port].tcp.httpStream = h
			}
		case len(httpTLS):
			if reflect.DeepEqual(httpTLS, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				ports[port].tls.httpStream = h
			}
		case len(tcpTLS):
			if reflect.DeepEqual(httpTLS, filter.FilterChainMatch.ApplicationProtocols) {
				h, err := Convert_config_listener_v3_FilterChain(conf, filter)
				if err != nil {
					return nil, err
				}
				ports[port].tls.tcpStream = h
			}
		}
	}
	m := bind.DestinationStreamHandlerConfig{
		NotFound: ports[0].Handler(),
	}
	ks := []int{}
	for k := range ports {
		if k == 0 {
			continue
		}
		ks = append(ks, int(k))
	}
	if len(ks) != 0 {
		sort.Ints(ks)
		for _, k := range ks {
			port := ports[uint32(k)]
			h := port.Handler()
			if h == nil {
				continue
			}
			m.Destinations = append(m.Destinations, bind.DestinationStreamHandlerRoute{
				Ports:   []uint32{uint32(k)},
				Handler: h,
			})
		}
	}
	if len(m.Destinations) == 0 {
		return m.NotFound, nil
	}
	return m, nil
}

type prefixMux struct {
	tcp prefixMuxProto
	tls prefixMuxProto
}

func (p *prefixMux) Handler() bind.StreamHandler {
	n := p.tcp.Handler()
	t := p.tls.Handler()
	if t == nil {
		return n
	}
	s, ok := n.(bind.PrefixStreamHandlerConfig)
	if ok {
		s.Routes = append(s.Routes, bind.PrefixStreamHandlerRoute{
			Pattern: bind.PrefixStreamHandlerProtocolEnumProtocolTLS,
			Handler: t,
		})
		return s
	}
	return bind.PrefixStreamHandlerConfig{
		Routes: []bind.PrefixStreamHandlerRoute{
			{
				Pattern: bind.PrefixStreamHandlerProtocolEnumProtocolTLS,
				Handler: t,
			},
		},
		NotFound: n,
	}
}

type prefixMuxProto struct {
	tcpStream  bind.StreamHandler
	httpStream bind.StreamHandler
}

func (p *prefixMuxProto) Handler() bind.StreamHandler {
	var lastStream bind.StreamHandler
	isHttp := p.httpStream != nil
	isTcp := p.tcpStream != nil

	switch {
	case isHttp && isTcp:
		lastStream = bind.PrefixStreamHandlerConfig{
			Routes: []bind.PrefixStreamHandlerRoute{
				{
					Pattern: bind.PrefixStreamHandlerProtocolEnumProtocolHTTP1,
					Handler: p.httpStream,
				},
				{
					Pattern: bind.PrefixStreamHandlerProtocolEnumProtocolHTTP2,
					Handler: p.httpStream,
				},
			},
			NotFound: p.tcpStream,
		}
	case isHttp && !isTcp:
		lastStream = p.httpStream
	case isTcp && !isHttp:
		lastStream = p.tcpStream
	}
	return lastStream
}

package config

import (
	"encoding/json"
)

const (
	kindDecoder       = "github.com/wzshiming/pipe/codec.Decoder"
	kindEncoder       = "github.com/wzshiming/pipe/codec.Encoder"
	kindUnmarshaler   = "github.com/wzshiming/pipe/codec.Unmarshaler"
	kindMarshaler     = "github.com/wzshiming/pipe/codec.Marshaler"
	kindHttpHandler   = "net/http.Handler"
	kindListenConfig  = "github.com/wzshiming/pipe/listener.ListenConfig"
	kindService       = "github.com/wzshiming/pipe/service.Service"
	kindStreamHandler = "github.com/wzshiming/pipe/stream.Handler"
	kindTLS           = "github.com/wzshiming/pipe/tls.TLS"
	kindOnce          = "github.com/wzshiming/pipe/once.Once"

	KindStreamHandlerForward = kindStreamHandler + "@forward"
	KindStreamHandlerHTTP    = kindStreamHandler + "@http"
	KindStreamHandlerPoller  = kindStreamHandler + "@poller"
	KindStreamHandlerMulti   = kindStreamHandler + "@multi"
	KindHttpHandlerForward   = kindHttpHandler + "@forward"
	KindHttpHandlerPoller    = kindHttpHandler + "@poller"
	KindHttpHandlerMux       = kindHttpHandler + "@mux"
	KindServiceServer        = kindService + "@server"
	KindServiceMulti         = kindService + "@multi"
	KindListenConfigNetwork  = kindListenConfig + "@network"
	KindOnceXDS              = kindOnce + "@xds"
)

func XdsName(name string) string {
	if name == "" {
		return ""
	}
	return "xds@" + name
}

func MarshalKindHttpHandlerForward(pass string, forward json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerForward, struct {
		Pass    string
		Forward json.RawMessage
	}{
		Pass:    pass,
		Forward: forward,
	})
}

type Route struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler json.RawMessage
}

func MarshalKindHttpHandlerMux(routes []*Route, notFound json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerMux, struct {
		Routes   []*Route
		NotFound json.RawMessage
	}{
		Routes:   routes,
		NotFound: notFound,
	})
}

func MarshalKindHttpHandlerPoller(poller string, handlers []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerPoller, struct {
		Poller   string
		Handlers []json.RawMessage
	}{
		Poller:   poller,
		Handlers: handlers,
	})
}

func MarshalKindStreamHandlerHTTP(handler, tls json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerHTTP, struct {
		Handler json.RawMessage
		TLS     json.RawMessage
	}{
		Handler: handler,
		TLS:     tls,
	})
}

func MarshalKindServiceMulti(multi []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindServiceMulti, struct {
		Multi []json.RawMessage
	}{
		Multi: multi,
	})
}

func MarshalKindServiceServer(listener, handler json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindServiceServer, struct {
		Listener json.RawMessage
		Handler  json.RawMessage
	}{
		Listener: listener,
		Handler:  handler,
	})
}

func MarshalKindListenConfigNetwork(network, address string) (json.RawMessage, error) {
	return MarshalKind(KindListenConfigNetwork, struct {
		Network string
		Address string
	}{
		Network: network,
		Address: address,
	})
}

func MarshalKindStreamHandlerMulti(multi []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerMulti, struct {
		Multi []json.RawMessage
	}{
		Multi: multi,
	})
}

func MarshalKindStreamHandlerForward(network, address string) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerForward, struct {
		Network string
		Address string
	}{
		Network: network,
		Address: address,
	})
}

func MarshalKindStreamHandlerPoller(poller string, handlers []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerPoller, struct {
		Poller   string
		Handlers []json.RawMessage
	}{
		Poller:   poller,
		Handlers: handlers,
	})
}

func MarshalKindOnceXDS(nodeID string, forward json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindOnceXDS, struct {
		NodeID  string
		Forward json.RawMessage
	}{
		NodeID:  nodeID,
		Forward: forward,
	})
}

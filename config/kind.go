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
	kindOutput        = "io.WriteCloser"

	KindStreamHandlerForward            = kindStreamHandler + "@forward"
	KindStreamHandlerHTTP               = kindStreamHandler + "@http"
	KindStreamHandlerPoller             = kindStreamHandler + "@poller"
	KindStreamHandlerMulti              = kindStreamHandler + "@multi"
	KindHttpHandlerForward              = kindHttpHandler + "@forward"
	KindHttpHandlerPoller               = kindHttpHandler + "@poller"
	KindHttpHandlerMux                  = kindHttpHandler + "@mux"
	KindHttpHandlerDirect               = kindHttpHandler + "@direct"
	KindHttpHandlerLog                  = kindHttpHandler + "@log"
	KindHttpHandlerAddResponseHeader    = kindHttpHandler + "@add_response_header"
	KindHttpHandlerAddRequestHeader     = kindHttpHandler + "@add_request_header"
	KindHttpHandlerRemoveResponseHeader = kindHttpHandler + "@remove_response_header"
	KindHttpHandlerRemoveRequestHeader  = kindHttpHandler + "@remove_request_header"
	KindHttpHandlerMulti                = kindHttpHandler + "@multi"
	KindOutputFile                      = kindOutput + "@file"

	KindServiceServer       = kindService + "@server"
	KindServiceMulti        = kindService + "@multi"
	KindListenConfigNetwork = kindListenConfig + "@network"
	KindOnceADS             = kindOnce + "@ads"
	KindOnceXDS             = kindOnce + "@xds"
)

func XdsName(name string) string {
	if name == "" {
		return ""
	}
	return "xds@" + name
}

func MarshalKindHttpHandlerMulti(multi []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerMulti, struct {
		Multi []json.RawMessage
	}{
		Multi: multi,
	})
}

func MarshalKindOutputFile(path string) (json.RawMessage, error) {
	return MarshalKind(KindOutputFile, struct {
		Path string
	}{
		Path: path,
	})
}

func MarshalKindHttpHandlerLog(output, handler json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerLog, struct {
		Output  json.RawMessage
		Handler json.RawMessage
	}{
		Output:  output,
		Handler: handler,
	})
}

func MarshalKindHttpHandlerRemoveRequestHeader(key string) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerRemoveRequestHeader, struct {
		Key string
	}{
		Key: key,
	})
}

func MarshalKindHttpHandlerRemoveResponseHeader(key string) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerRemoveResponseHeader, struct {
		Key string
	}{
		Key: key,
	})
}

func MarshalKindHttpHandlerAddRequestHeader(key, value string) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerAddRequestHeader, struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	})
}

func MarshalKindHttpHandlerAddResponseHeader(key, value string) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerAddResponseHeader, struct {
		Key   string
		Value string
	}{
		Key:   key,
		Value: value,
	})
}

func MarshalKindHttpHandlerDirect(code int, body string) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerDirect, struct {
		Code int `json:",omitempty"`
		Body string
	}{
		Code: code,
		Body: body,
	})
}

func MarshalKindHttpHandlerForward(pass string, forward json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerForward, struct {
		Pass    string
		Forward json.RawMessage `json:",omitempty"`
	}{
		Pass:    pass,
		Forward: forward,
	})
}

type Route struct {
	Prefix  string `json:",omitempty"`
	Path    string `json:",omitempty"`
	Regexp  string `json:",omitempty"`
	Handler json.RawMessage
}

func MarshalKindHttpHandlerMux(routes []*Route, notFound json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerMux, struct {
		Routes   []*Route
		NotFound json.RawMessage `json:",omitempty"`
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
		TLS     json.RawMessage `json:",omitempty"`
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

func MarshalKindOnceADS(nodeID string, forward json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindOnceADS, struct {
		NodeID  string
		Forward json.RawMessage
	}{
		NodeID:  nodeID,
		Forward: forward,
	})
}

func MarshalKindOnceXDS(xds string, ads json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindOnceXDS, struct {
		XDS string
		ADS json.RawMessage
	}{
		XDS: xds,
		ADS: ads,
	})
}

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
	kindDialer        = "github.com/wzshiming/pipe/dialer.Dialer"
	kindService       = "github.com/wzshiming/pipe/service.Service"
	kindStreamHandler = "github.com/wzshiming/pipe/stream.Handler"
	kindTLS           = "github.com/wzshiming/pipe/tls.TLS"
	kindOnce          = "github.com/wzshiming/pipe/once.Once"
	kindOutput        = "io.WriteCloser"
	kindInput         = "io.ReadCloser"

	KindStreamHandlerForward            = kindStreamHandler + "@forward"
	KindStreamHandlerHTTP               = kindStreamHandler + "@http"
	KindStreamHandlerPoller             = kindStreamHandler + "@poller"
	KindStreamHandlerMulti              = kindStreamHandler + "@multi"
	KindStreamHandlerTlsUp              = kindStreamHandler + "@tls_up"
	KindStreamHandlerTlsDown            = kindStreamHandler + "@tls_down"
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
	KindHttpHandlerAccessLog            = kindHttpHandler + "@access_log"
	KindOutputFile                      = kindOutput + "@file"
	KindInputFile                       = kindInput + "@file"
	KindInputInline                     = kindInput + "@inline"
	KindTlsMerge                        = kindTLS + "@merge"
	KindTlsFrom                         = kindTLS + "@from"
	KindTlsValidation                   = kindTLS + "@validation"
	KindServiceServer                   = kindService + "@server"
	KindServiceMulti                    = kindService + "@multi"
	KindListenConfigNetwork             = kindListenConfig + "@network"
	KindDialerNetwork                   = kindDialer + "@network"
	KindDialerPoller                    = kindDialer + "@poller"
	KindOnceADS                         = kindOnce + "@ads"
	KindOnceXDS                         = kindOnce + "@xds"
	KindOnceAccessLog                   = kindOnce + "@access_log"
)

func XdsName(name string) string {
	if name == "" {
		return ""
	}
	return "xds@" + name
}

func MarshalKindTlsMergep(merge []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindTlsMerge, struct {
		Merge []json.RawMessage
	}{
		Merge: merge,
	})
}

func MarshalKindStreamHandlerTlsUp(tls, handler json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerTlsUp, struct {
		TLS     json.RawMessage
		Handler json.RawMessage
	}{
		TLS:     tls,
		Handler: handler,
	})
}

func MarshalKindStreamHandlerTlsDown(tls, handler json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerTlsDown, struct {
		TLS     json.RawMessage
		Handler json.RawMessage
	}{
		TLS:     tls,
		Handler: handler,
	})
}

func MarshalKindTlsValidation(ca json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindTlsValidation, struct {
		Ca json.RawMessage
	}{
		Ca: ca,
	})
}

func MarshalKindTlsFrom(domain string, cert, key json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindTlsFrom, struct {
		Domain string
		Cert   json.RawMessage
		Key    json.RawMessage
	}{
		Domain: domain,
		Cert:   cert,
		Key:    key,
	})
}

func MarshalKindHttpHandlerAccessLog(accessLog, handler json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerAccessLog, struct {
		AccessLog json.RawMessage
		Handler   json.RawMessage
	}{
		AccessLog: accessLog,
		Handler:   handler,
	})
}

func MarshalKindOnceAccessLog(nodeID, logName string, dialer json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindOnceAccessLog, struct {
		NodeID  string
		LogName string
		Dialer  json.RawMessage
	}{
		NodeID:  nodeID,
		LogName: logName,
		Dialer:  dialer,
	})
}

func MarshalKindHttpHandlerMulti(multi []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerMulti, struct {
		Multi []json.RawMessage
	}{
		Multi: multi,
	})
}

func MarshalKindInputInline(data string) (json.RawMessage, error) {
	return MarshalKind(KindInputInline, struct {
		Data string
	}{
		Data: data,
	})
}

func MarshalKindInputFile(path string) (json.RawMessage, error) {
	return MarshalKind(KindInputFile, struct {
		Path string
	}{
		Path: path,
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

func MarshalKindHttpHandlerDirect(code int, body json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerDirect, struct {
		Code int `json:",omitempty"`
		Body json.RawMessage
	}{
		Code: code,
		Body: body,
	})
}

func MarshalKindHttpHandlerForward(pass string, dialer json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindHttpHandlerForward, struct {
		Pass   string
		Dialer json.RawMessage `json:",omitempty"`
	}{
		Pass:   pass,
		Dialer: dialer,
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

func MarshalKindStreamHandlerForward(dialer json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindStreamHandlerForward, struct {
		Dialer json.RawMessage
	}{
		Dialer: dialer,
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

func MarshalKindOnceADS(nodeID string, dialer json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindOnceADS, struct {
		NodeID string
		Dialer json.RawMessage
	}{
		NodeID: nodeID,
		Dialer: dialer,
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

func MarshalKindDialerNetwork(network, address string) (json.RawMessage, error) {
	return MarshalKind(KindDialerNetwork, struct {
		Network string
		Address string
	}{
		Network: network,
		Address: address,
	})
}

func MarshalKindDialerPoller(poller string, dialers []json.RawMessage) (json.RawMessage, error) {
	return MarshalKind(KindDialerPoller, struct {
		Poller  string
		Dialers []json.RawMessage
	}{
		Poller:  poller,
		Dialers: dialers,
	})
}

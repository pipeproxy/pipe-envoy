// DO NOT EDIT! Code generated.

package bind

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PipeConfig struct {
	Pipe       Service
	Init       []Once
	Components []PipeComponent
}

type PipeComponent interface {
	isPipeComponent()
	json.Marshaler
}

type RawPipeComponent []byte

func (RawPipeComponent) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawPipeComponent) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawPipeComponent) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawPipeComponent: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NamePipeComponent struct {
	Name string
	PipeComponent
}

func (NamePipeComponent) isPipeComponent() {}

// MarshalJSON returns n as the JSON encoding of n.
func (n NamePipeComponent) MarshalJSON() ([]byte, error) {
	data, err := n.PipeComponent.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefPipeComponent string

func (RefPipeComponent) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of r.
func (r RefPipeComponent) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", r)), nil
}

type CodecDecoder interface {
	isCodecDecoder()
	PipeComponent
}

type RawCodecDecoder []byte

func (RawCodecDecoder) isCodecDecoder()  {}
func (RawCodecDecoder) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecDecoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecDecoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecDecoder: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameCodecDecoder struct {
	Name string
	CodecDecoder
}

func (NameCodecDecoder) isCodecDecoder()  {}
func (NameCodecDecoder) isPipeComponent() {}

func (n NameCodecDecoder) MarshalJSON() ([]byte, error) {
	data, err := n.CodecDecoder.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefCodecDecoder string

func (RefCodecDecoder) isCodecDecoder()  {}
func (RefCodecDecoder) isPipeComponent() {}

func (m RefCodecDecoder) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// CodecDecoderBase32Config github.com/wzshiming/pipe/codec.Decoder@base32
type CodecDecoderBase32Config struct {
	Encoding string
}

func (CodecDecoderBase32Config) isCodecDecoder()  {}
func (CodecDecoderBase32Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBase32Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Decoder@base32"
	type t CodecDecoderBase32Config
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecDecoderBase64Config github.com/wzshiming/pipe/codec.Decoder@base64
type CodecDecoderBase64Config struct {
	Encoding string
}

func (CodecDecoderBase64Config) isCodecDecoder()  {}
func (CodecDecoderBase64Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBase64Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Decoder@base64"
	type t CodecDecoderBase64Config
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecDecoderBzip2 github.com/wzshiming/pipe/codec.Decoder@bzip2
type CodecDecoderBzip2 struct {
}

func (CodecDecoderBzip2) isCodecDecoder()  {}
func (CodecDecoderBzip2) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBzip2) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Decoder@bzip2"
	type t CodecDecoderBzip2
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecDecoderGzip github.com/wzshiming/pipe/codec.Decoder@gzip
type CodecDecoderGzip struct {
}

func (CodecDecoderGzip) isCodecDecoder()  {}
func (CodecDecoderGzip) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderGzip) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Decoder@gzip"
	type t CodecDecoderGzip
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecDecoderHex github.com/wzshiming/pipe/codec.Decoder@hex
type CodecDecoderHex struct {
}

func (CodecDecoderHex) isCodecDecoder()  {}
func (CodecDecoderHex) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderHex) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Decoder@hex"
	type t CodecDecoderHex
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type CodecEncoder interface {
	isCodecEncoder()
	PipeComponent
}

type RawCodecEncoder []byte

func (RawCodecEncoder) isCodecEncoder()  {}
func (RawCodecEncoder) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecEncoder) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecEncoder) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecEncoder: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameCodecEncoder struct {
	Name string
	CodecEncoder
}

func (NameCodecEncoder) isCodecEncoder()  {}
func (NameCodecEncoder) isPipeComponent() {}

func (n NameCodecEncoder) MarshalJSON() ([]byte, error) {
	data, err := n.CodecEncoder.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefCodecEncoder string

func (RefCodecEncoder) isCodecEncoder()  {}
func (RefCodecEncoder) isPipeComponent() {}

func (m RefCodecEncoder) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// CodecEncoderBase32Config github.com/wzshiming/pipe/codec.Encoder@base32
type CodecEncoderBase32Config struct {
	Encoding string
}

func (CodecEncoderBase32Config) isCodecEncoder()  {}
func (CodecEncoderBase32Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderBase32Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Encoder@base32"
	type t CodecEncoderBase32Config
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecEncoderBase64Config github.com/wzshiming/pipe/codec.Encoder@base64
type CodecEncoderBase64Config struct {
	Encoding string
}

func (CodecEncoderBase64Config) isCodecEncoder()  {}
func (CodecEncoderBase64Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderBase64Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Encoder@base64"
	type t CodecEncoderBase64Config
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecEncoderGzip github.com/wzshiming/pipe/codec.Encoder@gzip
type CodecEncoderGzip struct {
}

func (CodecEncoderGzip) isCodecEncoder()  {}
func (CodecEncoderGzip) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderGzip) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Encoder@gzip"
	type t CodecEncoderGzip
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// CodecEncoderHex github.com/wzshiming/pipe/codec.Encoder@hex
type CodecEncoderHex struct {
}

func (CodecEncoderHex) isCodecEncoder()  {}
func (CodecEncoderHex) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderHex) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Encoder@hex"
	type t CodecEncoderHex
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type CodecMarshaler interface {
	isCodecMarshaler()
	PipeComponent
}

type RawCodecMarshaler []byte

func (RawCodecMarshaler) isCodecMarshaler() {}
func (RawCodecMarshaler) isPipeComponent()  {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecMarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecMarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecMarshaler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameCodecMarshaler struct {
	Name string
	CodecMarshaler
}

func (NameCodecMarshaler) isCodecMarshaler() {}
func (NameCodecMarshaler) isPipeComponent()  {}

func (n NameCodecMarshaler) MarshalJSON() ([]byte, error) {
	data, err := n.CodecMarshaler.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefCodecMarshaler string

func (RefCodecMarshaler) isCodecMarshaler() {}
func (RefCodecMarshaler) isPipeComponent()  {}

func (m RefCodecMarshaler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// CodecMarshalerJson github.com/wzshiming/pipe/codec.Marshaler@json
type CodecMarshalerJson struct {
}

func (CodecMarshalerJson) isCodecMarshaler() {}
func (CodecMarshalerJson) isPipeComponent()  {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecMarshalerJson) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Marshaler@json"
	type t CodecMarshalerJson
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type CodecUnmarshaler interface {
	isCodecUnmarshaler()
	PipeComponent
}

type RawCodecUnmarshaler []byte

func (RawCodecUnmarshaler) isCodecUnmarshaler() {}
func (RawCodecUnmarshaler) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawCodecUnmarshaler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawCodecUnmarshaler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameCodecUnmarshaler struct {
	Name string
	CodecUnmarshaler
}

func (NameCodecUnmarshaler) isCodecUnmarshaler() {}
func (NameCodecUnmarshaler) isPipeComponent()    {}

func (n NameCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	data, err := n.CodecUnmarshaler.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefCodecUnmarshaler string

func (RefCodecUnmarshaler) isCodecUnmarshaler() {}
func (RefCodecUnmarshaler) isPipeComponent()    {}

func (m RefCodecUnmarshaler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// CodecUnmarshalerJson github.com/wzshiming/pipe/codec.Unmarshaler@json
type CodecUnmarshalerJson struct {
}

func (CodecUnmarshalerJson) isCodecUnmarshaler() {}
func (CodecUnmarshalerJson) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecUnmarshalerJson) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/codec.Unmarshaler@json"
	type t CodecUnmarshalerJson
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type Dialer interface {
	isDialer()
	PipeComponent
}

type RawDialer []byte

func (RawDialer) isDialer()        {}
func (RawDialer) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawDialer) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawDialer) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawDialer: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameDialer struct {
	Name string
	Dialer
}

func (NameDialer) isDialer()        {}
func (NameDialer) isPipeComponent() {}

func (n NameDialer) MarshalJSON() ([]byte, error) {
	data, err := n.Dialer.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefDialer string

func (RefDialer) isDialer()        {}
func (RefDialer) isPipeComponent() {}

func (m RefDialer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// DialerNetworkConfig github.com/wzshiming/pipe/dialer.Dialer@network
type DialerNetworkConfig struct {
	Network string
	Address string
}

func (DialerNetworkConfig) isDialer()        {}
func (DialerNetworkConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DialerNetworkConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/dialer.Dialer@network"
	type t DialerNetworkConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// DialerPollerConfig github.com/wzshiming/pipe/dialer.Dialer@poller
type DialerPollerConfig struct {
	Poller  string
	Dialers []Dialer
}

func (DialerPollerConfig) isDialer()        {}
func (DialerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DialerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/dialer.Dialer@poller"
	type t DialerPollerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// DialerTlsConfig github.com/wzshiming/pipe/dialer.Dialer@tls
type DialerTlsConfig struct {
	Dialer Dialer
	TLS    TLS
}

func (DialerTlsConfig) isDialer()        {}
func (DialerTlsConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m DialerTlsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/dialer.Dialer@tls"
	type t DialerTlsConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type ListenerListenConfig interface {
	isListenerListenConfig()
	PipeComponent
}

type RawListenerListenConfig []byte

func (RawListenerListenConfig) isListenerListenConfig() {}
func (RawListenerListenConfig) isPipeComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawListenerListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawListenerListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawListenerListenConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameListenerListenConfig struct {
	Name string
	ListenerListenConfig
}

func (NameListenerListenConfig) isListenerListenConfig() {}
func (NameListenerListenConfig) isPipeComponent()        {}

func (n NameListenerListenConfig) MarshalJSON() ([]byte, error) {
	data, err := n.ListenerListenConfig.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefListenerListenConfig string

func (RefListenerListenConfig) isListenerListenConfig() {}
func (RefListenerListenConfig) isPipeComponent()        {}

func (m RefListenerListenConfig) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// ListenerListenConfigMultiConfig github.com/wzshiming/pipe/listener.ListenConfig@multi
type ListenerListenConfigMultiConfig struct {
	Multi []ListenerListenConfig
}

func (ListenerListenConfigMultiConfig) isListenerListenConfig() {}
func (ListenerListenConfigMultiConfig) isPipeComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ListenerListenConfigMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/listener.ListenConfig@multi"
	type t ListenerListenConfigMultiConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// ListenerListenConfigNetworkConfig github.com/wzshiming/pipe/listener.ListenConfig@network
type ListenerListenConfigNetworkConfig struct {
	Network string
	Address string
}

func (ListenerListenConfigNetworkConfig) isListenerListenConfig() {}
func (ListenerListenConfigNetworkConfig) isPipeComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ListenerListenConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/listener.ListenConfig@network"
	type t ListenerListenConfigNetworkConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// ListenerListenConfigTlsConfig github.com/wzshiming/pipe/listener.ListenConfig@tls
type ListenerListenConfigTlsConfig struct {
	ListenConfig ListenerListenConfig
	TLS          TLS
}

func (ListenerListenConfigTlsConfig) isListenerListenConfig() {}
func (ListenerListenConfigTlsConfig) isPipeComponent()        {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ListenerListenConfigTlsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/listener.ListenConfig@tls"
	type t ListenerListenConfigTlsConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type Once interface {
	isOnce()
	PipeComponent
}

type RawOnce []byte

func (RawOnce) isOnce()          {}
func (RawOnce) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawOnce) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawOnce) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawOnce: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameOnce struct {
	Name string
	Once
}

func (NameOnce) isOnce()          {}
func (NameOnce) isPipeComponent() {}

func (n NameOnce) MarshalJSON() ([]byte, error) {
	data, err := n.Once.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefOnce string

func (RefOnce) isOnce()          {}
func (RefOnce) isPipeComponent() {}

func (m RefOnce) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// OnceAccessLogConfig github.com/wzshiming/pipe/once.Once@access_log
type OnceAccessLogConfig struct {
	Name    string
	NodeID  string
	LogName string
	Dialer  Dialer
}

func (OnceAccessLogConfig) isOnce()          {}
func (OnceAccessLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceAccessLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/once.Once@access_log"
	type t OnceAccessLogConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// OnceAdsConfig github.com/wzshiming/pipe/once.Once@ads
type OnceAdsConfig struct {
	Name   string
	NodeID string
	Dialer Dialer
}

func (OnceAdsConfig) isOnce()          {}
func (OnceAdsConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceAdsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/once.Once@ads"
	type t OnceAdsConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// OnceMessageConfig github.com/wzshiming/pipe/once.Once@message
type OnceMessageConfig struct {
	Message string
}

func (OnceMessageConfig) isOnce()          {}
func (OnceMessageConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceMessageConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/once.Once@message"
	type t OnceMessageConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// OnceXdsConfig github.com/wzshiming/pipe/once.Once@xds
type OnceXdsConfig struct {
	XDS string
	ADS Once
}

func (OnceXdsConfig) isOnce()          {}
func (OnceXdsConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceXdsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/once.Once@xds"
	type t OnceXdsConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type Service interface {
	isService()
	PipeComponent
}

type RawService []byte

func (RawService) isService()       {}
func (RawService) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawService) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawService) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawService: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameService struct {
	Name string
	Service
}

func (NameService) isService()       {}
func (NameService) isPipeComponent() {}

func (n NameService) MarshalJSON() ([]byte, error) {
	data, err := n.Service.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefService string

func (RefService) isService()       {}
func (RefService) isPipeComponent() {}

func (m RefService) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// ServiceMultiConfig github.com/wzshiming/pipe/service.Service@multi
type ServiceMultiConfig struct {
	Multi []Service
}

func (ServiceMultiConfig) isService()       {}
func (ServiceMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/service.Service@multi"
	type t ServiceMultiConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// ServiceNone github.com/wzshiming/pipe/service.Service@none
type ServiceNone struct {
}

func (ServiceNone) isService()       {}
func (ServiceNone) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceNone) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/service.Service@none"
	type t ServiceNone
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// ServiceServerConfig github.com/wzshiming/pipe/service.Service@server
type ServiceServerConfig struct {
	Listener ListenerListenConfig
	Handler  StreamHandler
}

func (ServiceServerConfig) isService()       {}
func (ServiceServerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceServerConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/service.Service@server"
	type t ServiceServerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type StreamHandler interface {
	isStreamHandler()
	PipeComponent
}

type RawStreamHandler []byte

func (RawStreamHandler) isStreamHandler() {}
func (RawStreamHandler) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameStreamHandler struct {
	Name string
	StreamHandler
}

func (NameStreamHandler) isStreamHandler() {}
func (NameStreamHandler) isPipeComponent() {}

func (n NameStreamHandler) MarshalJSON() ([]byte, error) {
	data, err := n.StreamHandler.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefStreamHandler string

func (RefStreamHandler) isStreamHandler() {}
func (RefStreamHandler) isPipeComponent() {}

func (m RefStreamHandler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// StreamHandlerForwardConfig github.com/wzshiming/pipe/stream.Handler@forward
type StreamHandlerForwardConfig struct {
	Dialer Dialer
}

func (StreamHandlerForwardConfig) isStreamHandler() {}
func (StreamHandlerForwardConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerForwardConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@forward"
	type t StreamHandlerForwardConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerHttpConfig github.com/wzshiming/pipe/stream.Handler@http
type StreamHandlerHttpConfig struct {
	Handler HttpHandler
	TLS     TLS
}

func (StreamHandlerHttpConfig) isStreamHandler() {}
func (StreamHandlerHttpConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerHttpConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@http"
	type t StreamHandlerHttpConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerMultiConfig github.com/wzshiming/pipe/stream.Handler@multi
type StreamHandlerMultiConfig struct {
	Multi []StreamHandler
}

func (StreamHandlerMultiConfig) isStreamHandler() {}
func (StreamHandlerMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@multi"
	type t StreamHandlerMultiConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerMuxConfig github.com/wzshiming/pipe/stream.Handler@mux
type StreamHandlerMuxConfig struct {
	Routes   []StreamHandlerMuxRoute
	NotFound StreamHandler
}
type StreamHandlerMuxRoute struct {
	Pattern string
	Regexp  string
	Prefix  string
	Handler StreamHandler
}

func (StreamHandlerMuxConfig) isStreamHandler() {}
func (StreamHandlerMuxConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerMuxConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@mux"
	type t StreamHandlerMuxConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerPollerConfig github.com/wzshiming/pipe/stream.Handler@poller
type StreamHandlerPollerConfig struct {
	Poller   string
	Handlers []StreamHandler
}

func (StreamHandlerPollerConfig) isStreamHandler() {}
func (StreamHandlerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@poller"
	type t StreamHandlerPollerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerTlsDownConfig github.com/wzshiming/pipe/stream.Handler@tls_down
type StreamHandlerTlsDownConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func (StreamHandlerTlsDownConfig) isStreamHandler() {}
func (StreamHandlerTlsDownConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerTlsDownConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@tls_down"
	type t StreamHandlerTlsDownConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerTlsUpConfig github.com/wzshiming/pipe/stream.Handler@tls_up
type StreamHandlerTlsUpConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func (StreamHandlerTlsUpConfig) isStreamHandler() {}
func (StreamHandlerTlsUpConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerTlsUpConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@tls_up"
	type t StreamHandlerTlsUpConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// StreamHandlerWeightedConfig github.com/wzshiming/pipe/stream.Handler@weighted
type StreamHandlerWeightedConfig struct {
	Weighted []StreamHandlerWeightedWeighted
}
type StreamHandlerWeightedWeighted struct {
	Weight  uint
	Handler StreamHandler
}

func (StreamHandlerWeightedConfig) isStreamHandler() {}
func (StreamHandlerWeightedConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerWeightedConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/stream.Handler@weighted"
	type t StreamHandlerWeightedConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type TLS interface {
	isTLS()
	PipeComponent
}

type RawTLS []byte

func (RawTLS) isTLS()           {}
func (RawTLS) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawTLS) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawTLS) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawTLS: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameTLS struct {
	Name string
	TLS
}

func (NameTLS) isTLS()           {}
func (NameTLS) isPipeComponent() {}

func (n NameTLS) MarshalJSON() ([]byte, error) {
	data, err := n.TLS.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefTLS string

func (RefTLS) isTLS()           {}
func (RefTLS) isPipeComponent() {}

func (m RefTLS) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// TLSAcmeConfig github.com/wzshiming/pipe/tls.TLS@acme
type TLSAcmeConfig struct {
	Domains  []string
	CacheDir string
}

func (TLSAcmeConfig) isTLS()           {}
func (TLSAcmeConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSAcmeConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/tls.TLS@acme"
	type t TLSAcmeConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// TLSFromConfig github.com/wzshiming/pipe/tls.TLS@from
type TLSFromConfig struct {
	Domain string
	Cert   Input
	Key    Input
}

func (TLSFromConfig) isTLS()           {}
func (TLSFromConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSFromConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/tls.TLS@from"
	type t TLSFromConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// TLSMergeConfig github.com/wzshiming/pipe/tls.TLS@merge
type TLSMergeConfig struct {
	Merge []TLS
}

func (TLSMergeConfig) isTLS()           {}
func (TLSMergeConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSMergeConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/tls.TLS@merge"
	type t TLSMergeConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// TLSSelfSigned github.com/wzshiming/pipe/tls.TLS@self_signed
type TLSSelfSigned struct {
}

func (TLSSelfSigned) isTLS()           {}
func (TLSSelfSigned) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSSelfSigned) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/tls.TLS@self_signed"
	type t TLSSelfSigned
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// TLSValidationConfig github.com/wzshiming/pipe/tls.TLS@validation
type TLSValidationConfig struct {
	Ca Input
}

func (TLSValidationConfig) isTLS()           {}
func (TLSValidationConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSValidationConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/tls.TLS@validation"
	type t TLSValidationConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type Input interface {
	isInput()
	PipeComponent
}

type RawInput []byte

func (RawInput) isInput()         {}
func (RawInput) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawInput) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawInput) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawInput: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameInput struct {
	Name string
	Input
}

func (NameInput) isInput()         {}
func (NameInput) isPipeComponent() {}

func (n NameInput) MarshalJSON() ([]byte, error) {
	data, err := n.Input.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefInput string

func (RefInput) isInput()         {}
func (RefInput) isPipeComponent() {}

func (m RefInput) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// InputFileConfig io.ReadCloser@file
type InputFileConfig struct {
	Path string
}

func (InputFileConfig) isInput()         {}
func (InputFileConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m InputFileConfig) MarshalJSON() ([]byte, error) {
	const kind = "io.ReadCloser@file"
	type t InputFileConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// InputInlineConfig io.ReadCloser@inline
type InputInlineConfig struct {
	Data string
}

func (InputInlineConfig) isInput()         {}
func (InputInlineConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m InputInlineConfig) MarshalJSON() ([]byte, error) {
	const kind = "io.ReadCloser@inline"
	type t InputInlineConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type Output interface {
	isOutput()
	PipeComponent
}

type RawOutput []byte

func (RawOutput) isOutput()        {}
func (RawOutput) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawOutput) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawOutput) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawOutput: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameOutput struct {
	Name string
	Output
}

func (NameOutput) isOutput()        {}
func (NameOutput) isPipeComponent() {}

func (n NameOutput) MarshalJSON() ([]byte, error) {
	data, err := n.Output.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefOutput string

func (RefOutput) isOutput()        {}
func (RefOutput) isPipeComponent() {}

func (m RefOutput) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// OutputFileConfig io.WriteCloser@file
type OutputFileConfig struct {
	Path string
}

func (OutputFileConfig) isOutput()        {}
func (OutputFileConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OutputFileConfig) MarshalJSON() ([]byte, error) {
	const kind = "io.WriteCloser@file"
	type t OutputFileConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

type HttpHandler interface {
	isHttpHandler()
	PipeComponent
}

type RawHttpHandler []byte

func (RawHttpHandler) isHttpHandler()   {}
func (RawHttpHandler) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHttpHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHttpHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHttpHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameHttpHandler struct {
	Name string
	HttpHandler
}

func (NameHttpHandler) isHttpHandler()   {}
func (NameHttpHandler) isPipeComponent() {}

func (n NameHttpHandler) MarshalJSON() ([]byte, error) {
	data, err := n.HttpHandler.MarshalJSON()
	if err != nil {
		return nil, err
	}

	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Name\":%q}", n.Name))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Name\":%q,", n.Name)), data[1:]...)
		}
	}

	return data, nil
}

type RefHttpHandler string

func (RefHttpHandler) isHttpHandler()   {}
func (RefHttpHandler) isPipeComponent() {}

func (m RefHttpHandler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// HttpHandlerAccessLogConfig net/http.Handler@access_log
type HttpHandlerAccessLogConfig struct {
	AccessLog Once
	Handler   HttpHandler
}

func (HttpHandlerAccessLogConfig) isHttpHandler()   {}
func (HttpHandlerAccessLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerAccessLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@access_log"
	type t HttpHandlerAccessLogConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerAddRequestHeaderConfig net/http.Handler@add_request_header
type HttpHandlerAddRequestHeaderConfig struct {
	Key   string
	Value string
}

func (HttpHandlerAddRequestHeaderConfig) isHttpHandler()   {}
func (HttpHandlerAddRequestHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerAddRequestHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@add_request_header"
	type t HttpHandlerAddRequestHeaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerAddResponseHeaderConfig net/http.Handler@add_response_header
type HttpHandlerAddResponseHeaderConfig struct {
	Key   string
	Value string
}

func (HttpHandlerAddResponseHeaderConfig) isHttpHandler()   {}
func (HttpHandlerAddResponseHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerAddResponseHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@add_response_header"
	type t HttpHandlerAddResponseHeaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerCompressConfig net/http.Handler@compress
type HttpHandlerCompressConfig struct {
	Level   int
	Handler HttpHandler
}

func (HttpHandlerCompressConfig) isHttpHandler()   {}
func (HttpHandlerCompressConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerCompressConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@compress"
	type t HttpHandlerCompressConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerConfigDump net/http.Handler@config_dump
type HttpHandlerConfigDump struct {
}

func (HttpHandlerConfigDump) isHttpHandler()   {}
func (HttpHandlerConfigDump) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerConfigDump) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@config_dump"
	type t HttpHandlerConfigDump
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerDirectConfig net/http.Handler@direct
type HttpHandlerDirectConfig struct {
	Code int
	Body Input
}

func (HttpHandlerDirectConfig) isHttpHandler()   {}
func (HttpHandlerDirectConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerDirectConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@direct"
	type t HttpHandlerDirectConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerExpvar net/http.Handler@expvar
type HttpHandlerExpvar struct {
}

func (HttpHandlerExpvar) isHttpHandler()   {}
func (HttpHandlerExpvar) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerExpvar) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@expvar"
	type t HttpHandlerExpvar
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerFileConfig net/http.Handler@file
type HttpHandlerFileConfig struct {
	Root string
}

func (HttpHandlerFileConfig) isHttpHandler()   {}
func (HttpHandlerFileConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerFileConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@file"
	type t HttpHandlerFileConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerForwardConfig net/http.Handler@forward
type HttpHandlerForwardConfig struct {
	Dialer Dialer
	Pass   string
}

func (HttpHandlerForwardConfig) isHttpHandler()   {}
func (HttpHandlerForwardConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerForwardConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@forward"
	type t HttpHandlerForwardConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerH2CConfig net/http.Handler@h2c
type HttpHandlerH2CConfig struct {
	Handler HttpHandler
}

func (HttpHandlerH2CConfig) isHttpHandler()   {}
func (HttpHandlerH2CConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerH2CConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@h2c"
	type t HttpHandlerH2CConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerLogConfig net/http.Handler@log
type HttpHandlerLogConfig struct {
	Output  Output
	Handler HttpHandler
}

func (HttpHandlerLogConfig) isHttpHandler()   {}
func (HttpHandlerLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@log"
	type t HttpHandlerLogConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerMultiConfig net/http.Handler@multi
type HttpHandlerMultiConfig struct {
	Multi []HttpHandler
}

func (HttpHandlerMultiConfig) isHttpHandler()   {}
func (HttpHandlerMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@multi"
	type t HttpHandlerMultiConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerMuxConfig net/http.Handler@mux
type HttpHandlerMuxConfig struct {
	Routes   []HttpHandlerMuxRoute
	NotFound HttpHandler
}
type HttpHandlerMuxRoute struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler HttpHandler
}

func (HttpHandlerMuxConfig) isHttpHandler()   {}
func (HttpHandlerMuxConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerMuxConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@mux"
	type t HttpHandlerMuxConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerPollerConfig net/http.Handler@poller
type HttpHandlerPollerConfig struct {
	Poller   string
	Handlers []HttpHandler
}

func (HttpHandlerPollerConfig) isHttpHandler()   {}
func (HttpHandlerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@poller"
	type t HttpHandlerPollerConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerPprof net/http.Handler@pprof
type HttpHandlerPprof struct {
}

func (HttpHandlerPprof) isHttpHandler()   {}
func (HttpHandlerPprof) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerPprof) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@pprof"
	type t HttpHandlerPprof
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerRedirectConfig net/http.Handler@redirect
type HttpHandlerRedirectConfig struct {
	Code     int
	Location string
}

func (HttpHandlerRedirectConfig) isHttpHandler()   {}
func (HttpHandlerRedirectConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerRedirectConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@redirect"
	type t HttpHandlerRedirectConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerRemoveRequestHeaderConfig net/http.Handler@remove_request_header
type HttpHandlerRemoveRequestHeaderConfig struct {
	Key string
}

func (HttpHandlerRemoveRequestHeaderConfig) isHttpHandler()   {}
func (HttpHandlerRemoveRequestHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerRemoveRequestHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@remove_request_header"
	type t HttpHandlerRemoveRequestHeaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerRemoveResponseHeaderConfig net/http.Handler@remove_response_header
type HttpHandlerRemoveResponseHeaderConfig struct {
	Key string
}

func (HttpHandlerRemoveResponseHeaderConfig) isHttpHandler()   {}
func (HttpHandlerRemoveResponseHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerRemoveResponseHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@remove_response_header"
	type t HttpHandlerRemoveResponseHeaderConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

// HttpHandlerWeightedConfig net/http.Handler@weighted
type HttpHandlerWeightedConfig struct {
	Weighted []HttpHandlerWeightedWeighted
}
type HttpHandlerWeightedWeighted struct {
	Weight  uint
	Handler HttpHandler
}

func (HttpHandlerWeightedConfig) isHttpHandler()   {}
func (HttpHandlerWeightedConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HttpHandlerWeightedConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@weighted"
	type t HttpHandlerWeightedConfig
	data, err := json.Marshal(t(m))
	if err != nil {
		return nil, err
	}
	if data[0] == '{' {
		if len(data) == 2 {
			data = []byte(fmt.Sprintf("{\"@Kind\":%q}", kind))
		} else {
			data = append([]byte(fmt.Sprintf("{\"@Kind\":%q,", kind)), data[1:]...)
		}
	}
	return data, nil
}

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

// CodecDecoderBase32Config github.com/wzshiming/pipe/pipe/codec.Decoder@base32
type CodecDecoderBase32Config struct {
	Encoding string
}

func (CodecDecoderBase32Config) isCodecDecoder()  {}
func (CodecDecoderBase32Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBase32Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@base32"
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

// CodecDecoderBase64Config github.com/wzshiming/pipe/pipe/codec.Decoder@base64
type CodecDecoderBase64Config struct {
	Encoding string
}

func (CodecDecoderBase64Config) isCodecDecoder()  {}
func (CodecDecoderBase64Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBase64Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@base64"
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

// CodecDecoderBzip2 github.com/wzshiming/pipe/pipe/codec.Decoder@bzip2
type CodecDecoderBzip2 struct {
}

func (CodecDecoderBzip2) isCodecDecoder()  {}
func (CodecDecoderBzip2) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderBzip2) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@bzip2"
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

// CodecDecoderGzip github.com/wzshiming/pipe/pipe/codec.Decoder@gzip
type CodecDecoderGzip struct {
}

func (CodecDecoderGzip) isCodecDecoder()  {}
func (CodecDecoderGzip) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderGzip) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@gzip"
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

// CodecDecoderHex github.com/wzshiming/pipe/pipe/codec.Decoder@hex
type CodecDecoderHex struct {
}

func (CodecDecoderHex) isCodecDecoder()  {}
func (CodecDecoderHex) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderHex) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@hex"
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

// CodecDecoderLoadConfig github.com/wzshiming/pipe/pipe/codec.Decoder@load
type CodecDecoderLoadConfig struct {
	Load Input
}

func (CodecDecoderLoadConfig) isCodecDecoder()  {}
func (CodecDecoderLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecDecoderLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Decoder@load"
	type t CodecDecoderLoadConfig
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

// CodecEncoderBase32Config github.com/wzshiming/pipe/pipe/codec.Encoder@base32
type CodecEncoderBase32Config struct {
	Encoding string
}

func (CodecEncoderBase32Config) isCodecEncoder()  {}
func (CodecEncoderBase32Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderBase32Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Encoder@base32"
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

// CodecEncoderBase64Config github.com/wzshiming/pipe/pipe/codec.Encoder@base64
type CodecEncoderBase64Config struct {
	Encoding string
}

func (CodecEncoderBase64Config) isCodecEncoder()  {}
func (CodecEncoderBase64Config) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderBase64Config) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Encoder@base64"
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

// CodecEncoderGzip github.com/wzshiming/pipe/pipe/codec.Encoder@gzip
type CodecEncoderGzip struct {
}

func (CodecEncoderGzip) isCodecEncoder()  {}
func (CodecEncoderGzip) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderGzip) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Encoder@gzip"
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

// CodecEncoderHex github.com/wzshiming/pipe/pipe/codec.Encoder@hex
type CodecEncoderHex struct {
}

func (CodecEncoderHex) isCodecEncoder()  {}
func (CodecEncoderHex) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderHex) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Encoder@hex"
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

// CodecEncoderLoadConfig github.com/wzshiming/pipe/pipe/codec.Encoder@load
type CodecEncoderLoadConfig struct {
	Load Input
}

func (CodecEncoderLoadConfig) isCodecEncoder()  {}
func (CodecEncoderLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecEncoderLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Encoder@load"
	type t CodecEncoderLoadConfig
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

// CodecMarshalerJSON github.com/wzshiming/pipe/pipe/codec.Marshaler@json
type CodecMarshalerJSON struct {
}

func (CodecMarshalerJSON) isCodecMarshaler() {}
func (CodecMarshalerJSON) isPipeComponent()  {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecMarshalerJSON) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Marshaler@json"
	type t CodecMarshalerJSON
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

// CodecMarshalerLoadConfig github.com/wzshiming/pipe/pipe/codec.Marshaler@load
type CodecMarshalerLoadConfig struct {
	Load Input
}

func (CodecMarshalerLoadConfig) isCodecMarshaler() {}
func (CodecMarshalerLoadConfig) isPipeComponent()  {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecMarshalerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Marshaler@load"
	type t CodecMarshalerLoadConfig
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

// CodecUnmarshalerJSON github.com/wzshiming/pipe/pipe/codec.Unmarshaler@json
type CodecUnmarshalerJSON struct {
}

func (CodecUnmarshalerJSON) isCodecUnmarshaler() {}
func (CodecUnmarshalerJSON) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecUnmarshalerJSON) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Unmarshaler@json"
	type t CodecUnmarshalerJSON
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

// CodecUnmarshalerLoadConfig github.com/wzshiming/pipe/pipe/codec.Unmarshaler@load
type CodecUnmarshalerLoadConfig struct {
	Load Input
}

func (CodecUnmarshalerLoadConfig) isCodecUnmarshaler() {}
func (CodecUnmarshalerLoadConfig) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m CodecUnmarshalerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/codec.Unmarshaler@load"
	type t CodecUnmarshalerLoadConfig
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

// OnceAccessLogConfig github.com/wzshiming/pipe/pipe/once.Once@access_log
type OnceAccessLogConfig struct {
	Name    string
	NodeID  string
	LogName string
	Dialer  StreamDialer
}

func (OnceAccessLogConfig) isOnce()          {}
func (OnceAccessLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceAccessLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/once.Once@access_log"
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

// OnceAdsConfig github.com/wzshiming/pipe/pipe/once.Once@ads
type OnceAdsConfig struct {
	Name   string
	NodeID string
	Dialer StreamDialer
}

func (OnceAdsConfig) isOnce()          {}
func (OnceAdsConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceAdsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/once.Once@ads"
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

// OnceLoadConfig github.com/wzshiming/pipe/pipe/once.Once@load
type OnceLoadConfig struct {
	Load Input
}

func (OnceLoadConfig) isOnce()          {}
func (OnceLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/once.Once@load"
	type t OnceLoadConfig
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

// OnceMessageConfig github.com/wzshiming/pipe/pipe/once.Once@message
type OnceMessageConfig struct {
	Message string
}

func (OnceMessageConfig) isOnce()          {}
func (OnceMessageConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceMessageConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/once.Once@message"
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

// OnceXdsConfig github.com/wzshiming/pipe/pipe/once.Once@xds
type OnceXdsConfig struct {
	XDS string
	ADS Once
}

func (OnceXdsConfig) isOnce()          {}
func (OnceXdsConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OnceXdsConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/once.Once@xds"
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

type ProtocolHandler interface {
	isProtocolHandler()
	PipeComponent
}

type RawProtocolHandler []byte

func (RawProtocolHandler) isProtocolHandler() {}
func (RawProtocolHandler) isPipeComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawProtocolHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawProtocolHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawProtocolHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameProtocolHandler struct {
	Name string
	ProtocolHandler
}

func (NameProtocolHandler) isProtocolHandler() {}
func (NameProtocolHandler) isPipeComponent()   {}

func (n NameProtocolHandler) MarshalJSON() ([]byte, error) {
	data, err := n.ProtocolHandler.MarshalJSON()
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

type RefProtocolHandler string

func (RefProtocolHandler) isProtocolHandler() {}
func (RefProtocolHandler) isPipeComponent()   {}

func (m RefProtocolHandler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// ProtocolHandlerLoadConfig github.com/wzshiming/pipe/pipe/protocol.Handler@load
type ProtocolHandlerLoadConfig struct {
	Load Input
}

func (ProtocolHandlerLoadConfig) isProtocolHandler() {}
func (ProtocolHandlerLoadConfig) isPipeComponent()   {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ProtocolHandlerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/protocol.Handler@load"
	type t ProtocolHandlerLoadConfig
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

// ServiceLoadConfig github.com/wzshiming/pipe/pipe/service.Service@load
type ServiceLoadConfig struct {
	Load Input
}

func (ServiceLoadConfig) isService()       {}
func (ServiceLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/service.Service@load"
	type t ServiceLoadConfig
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

// ServiceMultiConfig github.com/wzshiming/pipe/pipe/service.Service@multi
type ServiceMultiConfig struct {
	Multi []Service
}

func (ServiceMultiConfig) isService()       {}
func (ServiceMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/service.Service@multi"
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

// ServiceNone github.com/wzshiming/pipe/pipe/service.Service@none
type ServiceNone struct {
}

func (ServiceNone) isService()       {}
func (ServiceNone) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceNone) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/service.Service@none"
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

// ServiceStreamConfig github.com/wzshiming/pipe/pipe/service.Service@stream
type ServiceStreamConfig struct {
	Listener StreamListenConfig
	Handler  StreamHandler
}

func (ServiceStreamConfig) isService()       {}
func (ServiceStreamConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m ServiceStreamConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/service.Service@stream"
	type t ServiceStreamConfig
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

// StreamHandlerForwardConfig github.com/wzshiming/pipe/pipe/stream.Handler@forward
type StreamHandlerForwardConfig struct {
	Dialer StreamDialer
}

func (StreamHandlerForwardConfig) isStreamHandler() {}
func (StreamHandlerForwardConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerForwardConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@forward"
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

// StreamHandlerHTTPConfig github.com/wzshiming/pipe/pipe/stream.Handler@http
type StreamHandlerHTTPConfig struct {
	Handler HTTPHandler
	TLS     TLS
}

func (StreamHandlerHTTPConfig) isStreamHandler() {}
func (StreamHandlerHTTPConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerHTTPConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@http"
	type t StreamHandlerHTTPConfig
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

// StreamHandlerLoadConfig github.com/wzshiming/pipe/pipe/stream.Handler@load
type StreamHandlerLoadConfig struct {
	Load Input
}

func (StreamHandlerLoadConfig) isStreamHandler() {}
func (StreamHandlerLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@load"
	type t StreamHandlerLoadConfig
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

// StreamHandlerMultiConfig github.com/wzshiming/pipe/pipe/stream.Handler@multi
type StreamHandlerMultiConfig struct {
	Multi []StreamHandler
}

func (StreamHandlerMultiConfig) isStreamHandler() {}
func (StreamHandlerMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@multi"
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

// StreamHandlerMuxConfig github.com/wzshiming/pipe/pipe/stream.Handler@mux
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
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@mux"
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

// StreamHandlerPollerConfig github.com/wzshiming/pipe/pipe/stream.Handler@poller
type StreamHandlerPollerConfig struct {
	Poller   StreamHandlerPollerPollerEnum
	Handlers []StreamHandler
}
type StreamHandlerPollerPollerEnum string

const (
	StreamHandlerPollerPollerEnumEnumRoundRobin StreamHandlerPollerPollerEnum = "round_robin"
	StreamHandlerPollerPollerEnumEnumRandom     StreamHandlerPollerPollerEnum = "random"
)

func (StreamHandlerPollerConfig) isStreamHandler() {}
func (StreamHandlerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@poller"
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

// StreamHandlerTLSDownConfig github.com/wzshiming/pipe/pipe/stream.Handler@tls_down
type StreamHandlerTLSDownConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func (StreamHandlerTLSDownConfig) isStreamHandler() {}
func (StreamHandlerTLSDownConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerTLSDownConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@tls_down"
	type t StreamHandlerTLSDownConfig
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

// StreamHandlerTLSUpConfig github.com/wzshiming/pipe/pipe/stream.Handler@tls_up
type StreamHandlerTLSUpConfig struct {
	Handler StreamHandler
	TLS     TLS
}

func (StreamHandlerTLSUpConfig) isStreamHandler() {}
func (StreamHandlerTLSUpConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamHandlerTLSUpConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@tls_up"
	type t StreamHandlerTLSUpConfig
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

// StreamHandlerWeightedConfig github.com/wzshiming/pipe/pipe/stream.Handler@weighted
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
	const kind = "github.com/wzshiming/pipe/pipe/stream.Handler@weighted"
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

type StreamDialer interface {
	isStreamDialer()
	PipeComponent
}

type RawStreamDialer []byte

func (RawStreamDialer) isStreamDialer()  {}
func (RawStreamDialer) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamDialer) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamDialer) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamDialer: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameStreamDialer struct {
	Name string
	StreamDialer
}

func (NameStreamDialer) isStreamDialer()  {}
func (NameStreamDialer) isPipeComponent() {}

func (n NameStreamDialer) MarshalJSON() ([]byte, error) {
	data, err := n.StreamDialer.MarshalJSON()
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

type RefStreamDialer string

func (RefStreamDialer) isStreamDialer()  {}
func (RefStreamDialer) isPipeComponent() {}

func (m RefStreamDialer) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// StreamDialerLoadConfig github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@load
type StreamDialerLoadConfig struct {
	Load Input
}

func (StreamDialerLoadConfig) isStreamDialer()  {}
func (StreamDialerLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamDialerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@load"
	type t StreamDialerLoadConfig
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

// StreamDialerNetworkConfig github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@network
type StreamDialerNetworkConfig struct {
	Network StreamDialerNetworkNetworkEnum
	Address string
}
type StreamDialerNetworkNetworkEnum string

const (
	StreamDialerNetworkNetworkEnumEnumUnix StreamDialerNetworkNetworkEnum = "unix"
	StreamDialerNetworkNetworkEnumEnumTCP6 StreamDialerNetworkNetworkEnum = "tcp6"
	StreamDialerNetworkNetworkEnumEnumTCP4 StreamDialerNetworkNetworkEnum = "tcp4"
	StreamDialerNetworkNetworkEnumEnumTCP  StreamDialerNetworkNetworkEnum = "tcp"
)

func (StreamDialerNetworkConfig) isStreamDialer()  {}
func (StreamDialerNetworkConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamDialerNetworkConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@network"
	type t StreamDialerNetworkConfig
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

// StreamDialerPollerConfig github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@poller
type StreamDialerPollerConfig struct {
	Poller  StreamDialerPollerPollerEnum
	Dialers []StreamDialer
}
type StreamDialerPollerPollerEnum string

const (
	StreamDialerPollerPollerEnumEnumRoundRobin StreamDialerPollerPollerEnum = "round_robin"
	StreamDialerPollerPollerEnumEnumRandom     StreamDialerPollerPollerEnum = "random"
)

func (StreamDialerPollerConfig) isStreamDialer()  {}
func (StreamDialerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamDialerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@poller"
	type t StreamDialerPollerConfig
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

// StreamDialerTLSConfig github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@tls
type StreamDialerTLSConfig struct {
	Dialer StreamDialer
	TLS    TLS
}

func (StreamDialerTLSConfig) isStreamDialer()  {}
func (StreamDialerTLSConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamDialerTLSConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/dialer.Dialer@tls"
	type t StreamDialerTLSConfig
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

type StreamListenConfig interface {
	isStreamListenConfig()
	PipeComponent
}

type RawStreamListenConfig []byte

func (RawStreamListenConfig) isStreamListenConfig() {}
func (RawStreamListenConfig) isPipeComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawStreamListenConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawStreamListenConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawStreamListenConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameStreamListenConfig struct {
	Name string
	StreamListenConfig
}

func (NameStreamListenConfig) isStreamListenConfig() {}
func (NameStreamListenConfig) isPipeComponent()      {}

func (n NameStreamListenConfig) MarshalJSON() ([]byte, error) {
	data, err := n.StreamListenConfig.MarshalJSON()
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

type RefStreamListenConfig string

func (RefStreamListenConfig) isStreamListenConfig() {}
func (RefStreamListenConfig) isPipeComponent()      {}

func (m RefStreamListenConfig) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// StreamListenConfigLoadConfig github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@load
type StreamListenConfigLoadConfig struct {
	Load Input
}

func (StreamListenConfigLoadConfig) isStreamListenConfig() {}
func (StreamListenConfigLoadConfig) isPipeComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamListenConfigLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@load"
	type t StreamListenConfigLoadConfig
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

// StreamListenConfigNetworkConfig github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@network
type StreamListenConfigNetworkConfig struct {
	Network StreamListenConfigNetworkNetworkEnum
	Address string
}
type StreamListenConfigNetworkNetworkEnum string

const (
	StreamListenConfigNetworkNetworkEnumEnumUnix StreamListenConfigNetworkNetworkEnum = "unix"
	StreamListenConfigNetworkNetworkEnumEnumTCP6 StreamListenConfigNetworkNetworkEnum = "tcp6"
	StreamListenConfigNetworkNetworkEnumEnumTCP4 StreamListenConfigNetworkNetworkEnum = "tcp4"
	StreamListenConfigNetworkNetworkEnumEnumTCP  StreamListenConfigNetworkNetworkEnum = "tcp"
)

func (StreamListenConfigNetworkConfig) isStreamListenConfig() {}
func (StreamListenConfigNetworkConfig) isPipeComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamListenConfigNetworkConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@network"
	type t StreamListenConfigNetworkConfig
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

// StreamListenConfigTLSConfig github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@tls
type StreamListenConfigTLSConfig struct {
	ListenConfig StreamListenConfig
	TLS          TLS
}

func (StreamListenConfigTLSConfig) isStreamListenConfig() {}
func (StreamListenConfigTLSConfig) isPipeComponent()      {}

// MarshalJSON returns m as the JSON encoding of m.
func (m StreamListenConfigTLSConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/stream/listener.ListenConfig@tls"
	type t StreamListenConfigTLSConfig
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

// TLSAcmeConfig github.com/wzshiming/pipe/pipe/tls.TLS@acme
type TLSAcmeConfig struct {
	Domains  []string
	CacheDir string
}

func (TLSAcmeConfig) isTLS()           {}
func (TLSAcmeConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSAcmeConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@acme"
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

// TLSFromConfig github.com/wzshiming/pipe/pipe/tls.TLS@from
type TLSFromConfig struct {
	Domain string
	Cert   Input
	Key    Input
}

func (TLSFromConfig) isTLS()           {}
func (TLSFromConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSFromConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@from"
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

// TLSLoadConfig github.com/wzshiming/pipe/pipe/tls.TLS@load
type TLSLoadConfig struct {
	Load Input
}

func (TLSLoadConfig) isTLS()           {}
func (TLSLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@load"
	type t TLSLoadConfig
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

// TLSMergeConfig github.com/wzshiming/pipe/pipe/tls.TLS@merge
type TLSMergeConfig struct {
	Merge []TLS
}

func (TLSMergeConfig) isTLS()           {}
func (TLSMergeConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSMergeConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@merge"
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

// TLSSelfSigned github.com/wzshiming/pipe/pipe/tls.TLS@self_signed
type TLSSelfSigned struct {
}

func (TLSSelfSigned) isTLS()           {}
func (TLSSelfSigned) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSSelfSigned) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@self_signed"
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

// TLSValidationConfig github.com/wzshiming/pipe/pipe/tls.TLS@validation
type TLSValidationConfig struct {
	Ca Input
}

func (TLSValidationConfig) isTLS()           {}
func (TLSValidationConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m TLSValidationConfig) MarshalJSON() ([]byte, error) {
	const kind = "github.com/wzshiming/pipe/pipe/tls.TLS@validation"
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

// InputLoadConfig io.ReadCloser@load
type InputLoadConfig struct {
	Load Input
}

func (InputLoadConfig) isInput()         {}
func (InputLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m InputLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "io.ReadCloser@load"
	type t InputLoadConfig
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

// OutputLoadConfig io.WriteCloser@load
type OutputLoadConfig struct {
	Load Input
}

func (OutputLoadConfig) isOutput()        {}
func (OutputLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m OutputLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "io.WriteCloser@load"
	type t OutputLoadConfig
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

type HTTPHandler interface {
	isHTTPHandler()
	PipeComponent
}

type RawHTTPHandler []byte

func (RawHTTPHandler) isHTTPHandler()   {}
func (RawHTTPHandler) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHTTPHandler) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHTTPHandler) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHTTPHandler: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameHTTPHandler struct {
	Name string
	HTTPHandler
}

func (NameHTTPHandler) isHTTPHandler()   {}
func (NameHTTPHandler) isPipeComponent() {}

func (n NameHTTPHandler) MarshalJSON() ([]byte, error) {
	data, err := n.HTTPHandler.MarshalJSON()
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

type RefHTTPHandler string

func (RefHTTPHandler) isHTTPHandler()   {}
func (RefHTTPHandler) isPipeComponent() {}

func (m RefHTTPHandler) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// HTTPHandlerAccessLogConfig net/http.Handler@access_log
type HTTPHandlerAccessLogConfig struct {
	AccessLog Once
	Handler   HTTPHandler
}

func (HTTPHandlerAccessLogConfig) isHTTPHandler()   {}
func (HTTPHandlerAccessLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerAccessLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@access_log"
	type t HTTPHandlerAccessLogConfig
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

// HTTPHandlerAddRequestHeaderConfig net/http.Handler@add_request_header
type HTTPHandlerAddRequestHeaderConfig struct {
	Key   string
	Value string
}

func (HTTPHandlerAddRequestHeaderConfig) isHTTPHandler()   {}
func (HTTPHandlerAddRequestHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerAddRequestHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@add_request_header"
	type t HTTPHandlerAddRequestHeaderConfig
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

// HTTPHandlerAddResponseHeaderConfig net/http.Handler@add_response_header
type HTTPHandlerAddResponseHeaderConfig struct {
	Key   string
	Value string
}

func (HTTPHandlerAddResponseHeaderConfig) isHTTPHandler()   {}
func (HTTPHandlerAddResponseHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerAddResponseHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@add_response_header"
	type t HTTPHandlerAddResponseHeaderConfig
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

// HTTPHandlerCompressConfig net/http.Handler@compress
type HTTPHandlerCompressConfig struct {
	Level   int
	Handler HTTPHandler
}

func (HTTPHandlerCompressConfig) isHTTPHandler()   {}
func (HTTPHandlerCompressConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerCompressConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@compress"
	type t HTTPHandlerCompressConfig
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

// HTTPHandlerConfigDump net/http.Handler@config_dump
type HTTPHandlerConfigDump struct {
}

func (HTTPHandlerConfigDump) isHTTPHandler()   {}
func (HTTPHandlerConfigDump) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerConfigDump) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@config_dump"
	type t HTTPHandlerConfigDump
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

// HTTPHandlerDirectConfig net/http.Handler@direct
type HTTPHandlerDirectConfig struct {
	Code int
	Body Input
}

func (HTTPHandlerDirectConfig) isHTTPHandler()   {}
func (HTTPHandlerDirectConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerDirectConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@direct"
	type t HTTPHandlerDirectConfig
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

// HTTPHandlerExpvar net/http.Handler@expvar
type HTTPHandlerExpvar struct {
}

func (HTTPHandlerExpvar) isHTTPHandler()   {}
func (HTTPHandlerExpvar) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerExpvar) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@expvar"
	type t HTTPHandlerExpvar
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

// HTTPHandlerFileConfig net/http.Handler@file
type HTTPHandlerFileConfig struct {
	Root string
}

func (HTTPHandlerFileConfig) isHTTPHandler()   {}
func (HTTPHandlerFileConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerFileConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@file"
	type t HTTPHandlerFileConfig
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

// HTTPHandlerForwardConfig net/http.Handler@forward
type HTTPHandlerForwardConfig struct {
	RoundTripper HTTPRoundTripper
	URL          string
}

func (HTTPHandlerForwardConfig) isHTTPHandler()   {}
func (HTTPHandlerForwardConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerForwardConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@forward"
	type t HTTPHandlerForwardConfig
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

// HTTPHandlerH2CConfig net/http.Handler@h2c
type HTTPHandlerH2CConfig struct {
	Handler HTTPHandler
}

func (HTTPHandlerH2CConfig) isHTTPHandler()   {}
func (HTTPHandlerH2CConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerH2CConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@h2c"
	type t HTTPHandlerH2CConfig
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

// HTTPHandlerLoadConfig net/http.Handler@load
type HTTPHandlerLoadConfig struct {
	Load Input
}

func (HTTPHandlerLoadConfig) isHTTPHandler()   {}
func (HTTPHandlerLoadConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@load"
	type t HTTPHandlerLoadConfig
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

// HTTPHandlerLogConfig net/http.Handler@log
type HTTPHandlerLogConfig struct {
	Output  Output
	Handler HTTPHandler
}

func (HTTPHandlerLogConfig) isHTTPHandler()   {}
func (HTTPHandlerLogConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerLogConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@log"
	type t HTTPHandlerLogConfig
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

// HTTPHandlerMultiConfig net/http.Handler@multi
type HTTPHandlerMultiConfig struct {
	Multi []HTTPHandler
}

func (HTTPHandlerMultiConfig) isHTTPHandler()   {}
func (HTTPHandlerMultiConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerMultiConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@multi"
	type t HTTPHandlerMultiConfig
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

// HTTPHandlerMuxConfig net/http.Handler@mux
type HTTPHandlerMuxConfig struct {
	Routes   []HTTPHandlerMuxRoute
	NotFound HTTPHandler
}
type HTTPHandlerMuxRoute struct {
	Prefix  string
	Path    string
	Regexp  string
	Handler HTTPHandler
}

func (HTTPHandlerMuxConfig) isHTTPHandler()   {}
func (HTTPHandlerMuxConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerMuxConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@mux"
	type t HTTPHandlerMuxConfig
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

// HTTPHandlerPollerConfig net/http.Handler@poller
type HTTPHandlerPollerConfig struct {
	Poller   HTTPHandlerPollerPollerEnum
	Handlers []HTTPHandler
}
type HTTPHandlerPollerPollerEnum string

const (
	HTTPHandlerPollerPollerEnumEnumRoundRobin HTTPHandlerPollerPollerEnum = "round_robin"
	HTTPHandlerPollerPollerEnumEnumRandom     HTTPHandlerPollerPollerEnum = "random"
)

func (HTTPHandlerPollerConfig) isHTTPHandler()   {}
func (HTTPHandlerPollerConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerPollerConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@poller"
	type t HTTPHandlerPollerConfig
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

// HTTPHandlerPprof net/http.Handler@pprof
type HTTPHandlerPprof struct {
}

func (HTTPHandlerPprof) isHTTPHandler()   {}
func (HTTPHandlerPprof) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerPprof) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@pprof"
	type t HTTPHandlerPprof
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

// HTTPHandlerRedirectConfig net/http.Handler@redirect
type HTTPHandlerRedirectConfig struct {
	Code     int
	Location string
}

func (HTTPHandlerRedirectConfig) isHTTPHandler()   {}
func (HTTPHandlerRedirectConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerRedirectConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@redirect"
	type t HTTPHandlerRedirectConfig
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

// HTTPHandlerRemoveRequestHeaderConfig net/http.Handler@remove_request_header
type HTTPHandlerRemoveRequestHeaderConfig struct {
	Key string
}

func (HTTPHandlerRemoveRequestHeaderConfig) isHTTPHandler()   {}
func (HTTPHandlerRemoveRequestHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerRemoveRequestHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@remove_request_header"
	type t HTTPHandlerRemoveRequestHeaderConfig
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

// HTTPHandlerRemoveResponseHeaderConfig net/http.Handler@remove_response_header
type HTTPHandlerRemoveResponseHeaderConfig struct {
	Key string
}

func (HTTPHandlerRemoveResponseHeaderConfig) isHTTPHandler()   {}
func (HTTPHandlerRemoveResponseHeaderConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerRemoveResponseHeaderConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@remove_response_header"
	type t HTTPHandlerRemoveResponseHeaderConfig
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

// HTTPHandlerWeightedConfig net/http.Handler@weighted
type HTTPHandlerWeightedConfig struct {
	Weighted []HTTPHandlerWeightedWeighted
}
type HTTPHandlerWeightedWeighted struct {
	Weight  uint
	Handler HTTPHandler
}

func (HTTPHandlerWeightedConfig) isHTTPHandler()   {}
func (HTTPHandlerWeightedConfig) isPipeComponent() {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPHandlerWeightedConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.Handler@weighted"
	type t HTTPHandlerWeightedConfig
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

type HTTPRoundTripper interface {
	isHTTPRoundTripper()
	PipeComponent
}

type RawHTTPRoundTripper []byte

func (RawHTTPRoundTripper) isHTTPRoundTripper() {}
func (RawHTTPRoundTripper) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m RawHTTPRoundTripper) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON sets *m to a copy of data.
func (m *RawHTTPRoundTripper) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("RawHTTPRoundTripper: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[:0], data...)
	return nil
}

type NameHTTPRoundTripper struct {
	Name string
	HTTPRoundTripper
}

func (NameHTTPRoundTripper) isHTTPRoundTripper() {}
func (NameHTTPRoundTripper) isPipeComponent()    {}

func (n NameHTTPRoundTripper) MarshalJSON() ([]byte, error) {
	data, err := n.HTTPRoundTripper.MarshalJSON()
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

type RefHTTPRoundTripper string

func (RefHTTPRoundTripper) isHTTPRoundTripper() {}
func (RefHTTPRoundTripper) isPipeComponent()    {}

func (m RefHTTPRoundTripper) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"@Ref\":%q}", m)), nil
}

// HTTPRoundTripperLoadConfig net/http.RoundTripper@load
type HTTPRoundTripperLoadConfig struct {
	Load Input
}

func (HTTPRoundTripperLoadConfig) isHTTPRoundTripper() {}
func (HTTPRoundTripperLoadConfig) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPRoundTripperLoadConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.RoundTripper@load"
	type t HTTPRoundTripperLoadConfig
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

// HTTPRoundTripperTransportConfig net/http.RoundTripper@transport
type HTTPRoundTripperTransportConfig struct {
	TLS    TLS
	Dialer StreamDialer
}

func (HTTPRoundTripperTransportConfig) isHTTPRoundTripper() {}
func (HTTPRoundTripperTransportConfig) isPipeComponent()    {}

// MarshalJSON returns m as the JSON encoding of m.
func (m HTTPRoundTripperTransportConfig) MarshalJSON() ([]byte, error) {
	const kind = "net/http.RoundTripper@transport"
	type t HTTPRoundTripperTransportConfig
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

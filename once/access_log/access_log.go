package access_log

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_data_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v2"
	envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/wzshiming/envoy/internal/client/access_log"
	"github.com/wzshiming/envoy/internal/logger"
)

type AccessLog struct {
	cli        *access_log.Client
	once       sync.Once
	ch         chan *envoy_data_accesslog_v2.HTTPAccessLogEntry
	flush      chan struct{}
	logName    string
	bufferSize int
}

func NewAccessLog(logName string, buffer int, config *access_log.Config) (*AccessLog, error) {
	if config == nil {
		config = &access_log.Config{}
	}
	a := &AccessLog{}
	a.bufferSize = buffer
	a.logName = logName
	a.ch = make(chan *envoy_data_accesslog_v2.HTTPAccessLogEntry, a.bufferSize)
	a.flush = make(chan struct{})
	cli, err := access_log.NewClient("", config)
	if err != nil {
		return nil, err
	}
	a.cli = cli

	return a, nil
}

func (a *AccessLog) Do(ctx context.Context) error {
	a.do()
	return nil
}

func (a *AccessLog) do() {
	a.once.Do(func() {
		a.start()
	})
}

func (a *AccessLog) start() {
	logger.Info("start access log")
	err := a.cli.Start()
	if err != nil {
		logger.Fatalln(err)
		return
	}

	go a.sendChannel()
}

func (a *AccessLog) sendChannel() {
	buf := make([]*envoy_data_accesslog_v2.HTTPAccessLogEntry, 0, a.bufferSize)
	i := time.Second / 10
	tick := time.NewTimer(i)
	for {
		select {
		case <-tick.C:
		case <-a.flush:
		case d := <-a.ch:
			buf = append(buf, d)
			if len(buf) < cap(buf)/2 {
				continue
			}
		}
		if len(buf) == 0 {
			continue
		}
		tick.Reset(i)

		err := a.cli.SendHttpLog(a.logName, &envoy_service_accesslog_v2.StreamAccessLogsMessage_HTTPAccessLogEntries{
			LogEntry: buf,
		})
		if err != nil {
			logger.Errorf("send log: %s", err)
		}
		buf = buf[:0]
	}
}

func (a *AccessLog) SendHttpLog(log *envoy_data_accesslog_v2.HTTPAccessLogEntry) {
	a.do()

	select {
	case a.ch <- log:
	default:
		a.flush <- struct{}{}
		a.SendHttpLog(log)
	}
}

func (a *AccessLog) Handler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		warp := &warpHTTPAccessLogResponseWriter{
			ResponseWriter: rw,
		}

		handler.ServeHTTP(warp, r)
		reqProperties := getHTTPAccessLogResponse(r)
		a.SendHttpLog(&envoy_data_accesslog_v2.HTTPAccessLogEntry{
			Response:        &warp.properties,
			Request:         reqProperties,
			ProtocolVersion: getVersion(r.ProtoMajor, r.ProtoMinor),
		})
	})
}

func getVersion(major, minor int) envoy_data_accesslog_v2.HTTPAccessLogEntry_HTTPVersion {
	switch major {
	case 1:
		switch minor {
		case 0:
			return envoy_data_accesslog_v2.HTTPAccessLogEntry_HTTP10
		case 1:
			return envoy_data_accesslog_v2.HTTPAccessLogEntry_HTTP11
		}
	case 2:
		switch minor {
		case 0:
			return envoy_data_accesslog_v2.HTTPAccessLogEntry_HTTP2
		}
	case 3:
		switch minor {
		case 0:
			return envoy_data_accesslog_v2.HTTPAccessLogEntry_HTTP3
		}
	}
	return envoy_data_accesslog_v2.HTTPAccessLogEntry_PROTOCOL_UNSPECIFIED
}

func getHTTPAccessLogResponse(r *http.Request) *envoy_data_accesslog_v2.HTTPRequestProperties {

	forwardedFor, port, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		if prior, ok := r.Header["X-Forwarded-For"]; ok {
			forwardedFor = strings.Join(prior, ", ") + ", " + forwardedFor
		}
	}

	var portValue uint64 = 80

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
		portValue = 443
	}

	if port != "" {
		portValue, _ = strconv.ParseUint(port, 10, 32)
	}

	return &envoy_data_accesslog_v2.HTTPRequestProperties{
		RequestMethod:    envoy_api_v2_core.RequestMethod(envoy_api_v2_core.RequestMethod_value[r.Method]),
		Scheme:           scheme,
		Port:             &wrappers.UInt32Value{Value: uint32(portValue)},
		Path:             r.URL.Path,
		UserAgent:        r.UserAgent(),
		Referer:          r.Referer(),
		RequestHeaders:   getHeader(r.Header),
		RequestBodyBytes: uint64(r.ContentLength),
		ForwardedFor:     forwardedFor,
	}
}

type warpHTTPAccessLogResponseWriter struct {
	http.ResponseWriter
	properties envoy_data_accesslog_v2.HTTPResponseProperties
}

func (w *warpHTTPAccessLogResponseWriter) WriteHeader(c int) {
	w.properties.ResponseCode = &wrappers.UInt32Value{
		Value: uint32(c),
	}
	w.ResponseWriter.WriteHeader(c)
}

func (w *warpHTTPAccessLogResponseWriter) Write(b []byte) (int, error) {
	w.properties.ResponseBodyBytes += uint64(len(b))
	w.properties.ResponseHeaders = getHeader(w.ResponseWriter.Header())
	return w.ResponseWriter.Write(b)
}

func getHeader(header http.Header) map[string]string {
	m := map[string]string{}
	for key, val := range header {
		m[key] = val[0]
	}
	return m
}

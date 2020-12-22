package convert

//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/core/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/listener/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/route/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_dst/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cors/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3
//go:generate go run ../../hack/convert github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3

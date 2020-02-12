module github.com/wzshiming/envoy

go 1.13

require (
	github.com/envoyproxy/go-control-plane v0.9.2
	github.com/golang/protobuf v1.3.3
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/gotype v0.6.3
	github.com/wzshiming/pipe v0.0.7
	google.golang.org/grpc v1.27.0
	sigs.k8s.io/yaml v1.1.0
)

replace (
	github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.9.2
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.3
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200210222208-86ce3cb69678
	golang.org/x/lint => golang.org/x/lint v0.0.0-20200130185559-910be7a94367
	golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sync => golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200212091648-12a6c2dcc1e4
	golang.org/x/tools => golang.org/x/tools v0.0.0-20200211205636-11eff242d136
	google.golang.org/genproto => google.golang.org/genproto v0.0.0-20200212121238-0849286c0f0e
	google.golang.org/grpc => google.golang.org/grpc v1.27.1
)

module github.com/pipeproxy/pipe-xds

go 1.15

replace (
	github.com/docker/distribution => github.com/docker/distribution v2.7.1+incompatible
	github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190815234213-e83c0a1c26c8
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5
	github.com/moby/term => github.com/moby/term v0.0.0-20201110203204-bea5bbe245bf
)

require (
	github.com/cncf/udpa/go v0.0.0-20201001150855-7e6fe0510fb5
	github.com/envoyproxy/go-control-plane v0.9.8-0.20201019204000-12785f608982
	github.com/golang/protobuf v1.4.3
	github.com/pipeproxy/pipe v0.7.2
	github.com/spf13/cobra v1.1.1 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/wzshiming/gotype v0.6.3
	github.com/wzshiming/lockfile v0.1.0
	github.com/wzshiming/logger v0.0.0-20201218143744-5ae3f93dcd65
	github.com/wzshiming/notify v0.0.5
	google.golang.org/grpc v1.33.1
	google.golang.org/protobuf v1.25.0
	istio.io/api v0.0.0-20201120175956-c2df7c41fd8e
	istio.io/istio v0.0.0-20201118224433-c87a4c874df2
	sigs.k8s.io/yaml v1.2.0
)

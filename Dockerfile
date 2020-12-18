FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe-xds/
COPY . .
ENV CGO_ENABLED=0
RUN go install ./cmd/...

FROM alpine
RUN apk add -U --no-cache curl iptables ip6tables
COPY --from=pipeproxy/pipe:v0.7.2 /usr/local/bin/pipe /usr/local/bin/pipe
COPY --from=istio/proxyv2:1.8.0 /usr/local/bin/pilot-agent /usr/local/bin/pilot-agent
COPY --from=istio/proxyv2:1.8.0 /var/lib/istio/envoy/ /var/lib/istio/envoy/
COPY --from=builder /go/bin/envoy /usr/local/bin/envoy
ENTRYPOINT [ "/usr/local/bin/pilot-agent" ]

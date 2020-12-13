FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe-xds/
COPY . .
ENV CGO_ENABLED=0
RUN go install ./cmd/pipe-xds

FROM alpine
RUN apk add -U --no-cache curl iptables ip6tables
COPY --from=pipeproxy/pipe:v0.7.1 /usr/local/bin/pipe /usr/local/bin/
COPY --from=builder /go/bin/pipe-xds /usr/local/bin/
WORKDIR /etc/istio/proxy
ENTRYPOINT [ "pipe-xds" ]

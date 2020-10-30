FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe-xds/
COPY . .
ENV CGO_ENABLED=0
RUN go install ./cmd/pipe-xds

FROM alpine
COPY --from=builder /go/bin/pipe-xds /usr/local/bin/
COPY --from=pipeproxy/pipe:v0.4.5 /usr/local/bin/pipe /usr/local/bin/
WORKDIR /etc/istio/proxy/
ENTRYPOINT [ "pipe-xds", "-u", "istiod.istio-system:15010" ]

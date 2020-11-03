FROM golang:alpine AS builder
WORKDIR /go/src/github.com/pipeproxy/pipe-xds/
COPY . .
ENV CGO_ENABLED=0
RUN go install ./cmd/pipe-xds

FROM alpine
COPY --from=pipeproxy/pipe:v0.4.10 /usr/local/bin/pipe /usr/local/bin/
COPY --from=builder /go/bin/pipe-xds /usr/local/bin/
WORKDIR /etc/istio/proxy/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk add -U --no-cache curl
ENTRYPOINT [ "pipe-xds", "-u", "istiod.istio-system:15010" ]

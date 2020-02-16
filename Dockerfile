FROM golang:alpine AS builder
WORKDIR /go/src/github.com/wzshiming/envoy/
COPY . .
RUN go install ./cmd/envoy

FROM alpine
COPY --from=builder /go/bin/envoy /usr/local/bin/
ENTRYPOINT [ "envoy" ]

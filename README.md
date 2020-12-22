# Pipe xDS

[![Build Status](https://travis-ci.org/pipeproxy/pipe-xds.svg?branch=master)](https://travis-ci.org/pipeproxy/pipe-xds)
[![Go Report Card](https://goreportcard.com/badge/github.com/pipeproxy/pipe-xds)](https://goreportcard.com/report/github.com/pipeproxy/pipe-xds)
[![GoDoc](https://godoc.org/github.com/pipeproxy/pipe-xds?status.svg)](https://godoc.org/github.com/pipeproxy/pipe-xds)
[![Docker Automated build](https://img.shields.io/docker/cloud/automated/pipeproxy/pipe-xds.svg)](https://hub.docker.com/r/pipeproxy/pipe-xds)
[![GitHub license](https://img.shields.io/github/license/pipeproxy/pipe-xds.svg)](https://github.com/pipeproxy/pipe-xds/blob/master/LICENSE)


As far as possible compatible with xDS data plane.  

## Install

``` shell
istioctl install --set "values.global.proxy.image=pipeproxy/pipe-xds:v0.2.10" --set "values.global.proxy_init.image=pipeproxy/pipe-xds:v0.2.10"
```

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/pipeproxy/pipe-xds/blob/master/LICENSE) for the full license text.

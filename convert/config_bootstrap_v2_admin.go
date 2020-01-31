package convert

import (
	"encoding/json"

	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_bootstrap_v2_Admin(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Admin) (string, error) {
	name, err := Convert_api_v2_core_AddressListener(conf, c.Address)
	if err != nil {
		return "", err
	}
	listenerRef, err := config.MarshalRef(name)
	if err != nil {
		return "", err
	}

	routes := []*config.Route{
		{
			Path:    "/expvar/",
			Handler: config.MustMarshalKind("expvar", nil),
		},
		{
			Prefix:  "/pprof/",
			Handler: config.MustMarshalKind("pprof", nil),
		},
		{
			Prefix:  "/config_dump/",
			Handler: config.MustMarshalKind("config_dump", nil),
		},
	}

	dHeader, err := config.MarshalKindHttpHandlerAddResponseHeader("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		return "", err
	}

	const body = `
<pre>
<a href="/expvar/">/expvar/</a>
<a href="/pprof/">/pprof/</a>
<a href="/config_dump/">/config_dump/</a>
</pre>
`
	dBody, err := config.MarshalKindHttpHandlerDirect(200, body)
	if err != nil {
		return "", err
	}

	d, err := config.MarshalKindHttpHandlerMulti([]json.RawMessage{dHeader, dBody})
	if err != nil {
		return "", err
	}

	d, err = config.MarshalKindHttpHandlerMux(routes, d)
	if err != nil {
		return "", err
	}

	output, err := config.MarshalKindOutputFile(c.AccessLogPath)
	if err != nil {
		return "", err
	}

	d, err = config.MarshalKindHttpHandlerLog(output, d)
	if err != nil {
		return "", err
	}

	d, err = config.MarshalKindStreamHandlerHTTP(d, nil)
	if err != nil {
		return "", err
	}

	d, err = config.MarshalKindServiceServer(listenerRef, d)
	if err != nil {
		return "", err
	}

	name, err = conf.RegisterComponents("", d)
	if err != nil {
		return "", err
	}

	err = conf.RegisterService(name)
	if err != nil {
		return "", err
	}
	return name, nil
}

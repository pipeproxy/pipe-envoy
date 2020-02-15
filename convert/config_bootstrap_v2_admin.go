package convert

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_bootstrap_v2_Admin(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Admin) (bind.Service, error) {
	listener, err := Convert_api_v2_core_AddressListener(conf, c.Address)
	if err != nil {
		return nil, err
	}

	d := bind.ServiceStreamConfig{
		Listener: listener,
		Handler: bind.StreamHandlerHTTPConfig{
			Handler: bind.HTTPHandlerLogConfig{
				Output:  bind.OutputFileConfig{Path: c.AccessLogPath},
				Handler: adminHandler,
			},
			TLS: nil,
		},
	}

	ref, err := conf.RegisterComponents("", d)
	if err != nil {
		return nil, err
	}

	err = conf.RegisterService(ref)
	if err != nil {
		return nil, err
	}
	return bind.RefService(ref), nil
}

var adminHandler = bind.HTTPHandlerMuxConfig{
	Routes: []bind.HTTPHandlerMuxRoute{
		{
			Path:    "/expvar/",
			Handler: bind.HTTPHandlerExpvar{},
		},
		{
			Prefix:  "/pprof/",
			Handler: bind.HTTPHandlerPprof{},
		},
		{
			Prefix:  "/config_dump/",
			Handler: bind.HTTPHandlerConfigDump{},
		},
	},
	NotFound: bind.HTTPHandlerMultiConfig{
		Multi: []bind.HTTPHandler{
			bind.HTTPHandlerAddResponseHeaderConfig{
				Key:   "Content-Type",
				Value: "text/html; charset=utf-8",
			},
			bind.HTTPHandlerDirectConfig{
				Code: 200,
				Body: bind.InputInlineConfig{
					Data: `
<pre>
<a href="/expvar/">/expvar/</a>
<a href="/pprof/">/pprof/</a>
<a href="/config_dump/">/config_dump/</a>
</pre>
`,
				},
			},
		}},
}

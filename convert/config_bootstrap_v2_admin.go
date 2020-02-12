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

	d := bind.ServiceServerConfig{
		Listener: listener,
		Handler: bind.StreamHandlerHttpConfig{
			Handler: bind.HttpHandlerLogConfig{
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

var adminHandler = bind.HttpHandlerMuxConfig{
	Routes: []bind.HttpHandlerMuxRoute{
		{
			Path:    "/expvar/",
			Handler: bind.HttpHandlerExpvar{},
		},
		{
			Prefix:  "/pprof/",
			Handler: bind.HttpHandlerPprof{},
		},
		{
			Prefix:  "/config_dump/",
			Handler: bind.HttpHandlerConfigDump{},
		},
	},
	NotFound: bind.HttpHandlerMultiConfig{
		Multi: []bind.HttpHandler{
			bind.HttpHandlerAddResponseHeaderConfig{
				Key:   "Content-Type",
				Value: "text/html; charset=utf-8",
			},
			bind.HttpHandlerDirectConfig{
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

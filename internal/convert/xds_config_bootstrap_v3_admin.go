package convert

import (
	"net/http"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/internal/config"
	"github.com/pipeproxy/pipe/bind"
	pipeconfig "github.com/pipeproxy/pipe/config"
)

func Convert_config_bootstrap_v3_Admin(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.Admin) (string, error) {
	if c.Address == nil {
		return "", nil
	}
	_, address, err := Convert_config_core_v3_Address(conf, c.Address)
	if err != nil {
		return "", err
	}
	handler := BuildAdminWithHTTPHandler()

	if c.AccessLogPath != "" {
		handler = pipeconfig.BuildHTTPLog(c.AccessLogPath, handler)
	}

	name := "_admin"
	svc := pipeconfig.BuildH1WithService(address, handler)
	conf.RegisterLDS(name, bind.TagsServiceConfig{
		Service: svc,
		Tag:     name,
	}, c)
	return "", nil
}

func BuildAdminWithHTTPHandler() bind.HTTPHandler {
	return bind.PathNetHTTPHandlerConfig{
		Paths: []bind.PathNetHTTPHandlerRoute{
			{
				Path: "/",
				Handler: bind.MultiNetHTTPHandlerConfig{
					Multi: []bind.HTTPHandler{
						pipeconfig.BuildContentTypeHTMLWithHTTPHandler(),
						bind.DirectNetHTTPHandlerConfig{
							Code: http.StatusOK,
							Body: bind.InlineIoReaderConfig{
								Data: `/pprof/
/expvar
/quitquitquit
/config_dump
/stats/prometheus
/stats
`,
							},
						},
					},
				},
			},
			{
				Path: "/stats",
				Handler: bind.DirectNetHTTPHandlerConfig{
					Code: 200,
					Body: bind.InlineIoReaderConfig{
						Data: `# TODO
cluster_manager.cds.update_rejected: 0
cluster_manager.cds.update_success: 1
listener_manager.lds.update_rejected: 0
listener_manager.lds.update_success: 1
listener_manager.workers_started: 1
server.state: 0
`,
					},
				},
			},
			{
				Prefix:  "/pprof/",
				Handler: bind.PprofNetHTTPHandler{},
			},
			{
				Path:    "/expvar",
				Handler: bind.ExpvarNetHTTPHandler{},
			},
			{
				Path: "/quitquitquit",
				Handler: bind.MethodNetHTTPHandlerConfig{
					Methods: []bind.MethodNetHTTPHandlerRoute{
						{
							Method:  bind.MethodNetHTTPHandlerMethodEnumMethodPost,
							Handler: bind.QuitNetHTTPHandler{},
						},
					},
				},
			},
			{
				Path:    "/config_dump",
				Handler: bind.ConfigDumpNetHTTPHandlerConfig{},
			},
			{
				Path:    "/stats/prometheus",
				Handler: bind.MetricsNetHTTPHandler{},
			},
		},
	}
}

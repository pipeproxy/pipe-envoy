package convert

import (
	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/config"
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
	handler := pipeconfig.BuildAdminWithHTTPHandler()

	if c.AccessLogPath != "" {
		handler = pipeconfig.BuildHTTPLog(c.AccessLogPath, handler)
	}

	svc := pipeconfig.BuildH1WithService(address, handler)
	conf.RegisterLDS("admin", svc, c)
	return "", nil
}

package convert

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bind"
	"github.com/wzshiming/envoy/config"
)

func Convert_config_bootstrap_v2_Bootstrap_DynamicResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap_DynamicResources) (string, error) {

	if c.AdsConfig != nil {
		_, err := Convert_api_v2_core_ApiConfigSource(conf, c.AdsConfig)
		if err != nil {
			return "", err
		}
	}

	if c.CdsConfig != nil {
		cdsName, err := Convert_api_v2_core_ConfigSource(conf, c.CdsConfig)
		if err != nil {
			return "", err
		}

		if cdsName != "" {
			xds := bind.OnceXdsConfig{
				XDS: "cds",
				ADS: bind.RefOnce(cdsName),
			}
			_, err = conf.RegisterInit(xds)
			if err != nil {
				return "", err
			}
		}
	}

	if c.LdsConfig != nil {
		ldsName, err := Convert_api_v2_core_ConfigSource(conf, c.LdsConfig)
		if err != nil {
			return "", err
		}
		if ldsName != "" {
			xds := bind.OnceXdsConfig{
				XDS: "lds",
				ADS: bind.RefOnce(ldsName),
			}
			if err != nil {
				return "", err
			}
			_, err = conf.RegisterInit(xds)
			if err != nil {
				return "", err
			}
		}
	}

	return "", nil
}

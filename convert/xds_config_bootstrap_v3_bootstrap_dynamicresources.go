package convert

import (
	"log"

	envoy_config_bootstrap_v3 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/encoding"
)

func Convert_config_bootstrap_v3_Bootstrap_DynamicResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v3.Bootstrap_DynamicResources) (string, error) {

	if c.AdsConfig != nil {
		ads, err := Convert_config_core_v3_ApiConfigSource(conf, c.AdsConfig)
		if err != nil {
			return "", err
		}
		conf.RegisterCDS("ads", ads, c.AdsConfig)
		conf.RegisterADS("ads", ads)
	}

	if c.CdsConfig != nil {
		cds, err := Convert_config_core_v3_ConfigSource(conf, c.CdsConfig)
		if err != nil {
			return "", err
		}
		conf.RegisterADS("cds", cds)
	}

	if c.LdsConfig != nil {
		lds, err := Convert_config_core_v3_ConfigSource(conf, c.LdsConfig)
		if err != nil {
			return "", err
		}
		conf.RegisterADS("lds", lds)
	}

	data, _ := encoding.Marshal(c)
	log.Printf("[TODO] envoy_config_bootstrap_v3.Bootstrap_DynamicResources %s\n", string(data))
	return "", nil
}

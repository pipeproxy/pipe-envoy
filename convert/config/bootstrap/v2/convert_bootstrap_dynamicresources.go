package convert_config_bootstrap_v2

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/config"
	convert_api_v2_core "github.com/wzshiming/envoy/convert/api/v2/core"
)

func Convert_Bootstrap_DynamicResources(conf *config.ConfigCtx, c *envoy_config_bootstrap_v2.Bootstrap_DynamicResources) (string, error) {
	adsName := ""

	if c.AdsConfig != nil {
		name, err := convert_api_v2_core.Convert_ApiConfigSource(conf, c.AdsConfig)
		if err != nil {
			return "", err
		}
		adsName = name
	}

	if c.CdsConfig != nil {
		cdsName, err := convert_api_v2_core.Convert_ConfigSource(conf, c.CdsConfig)
		if err != nil {
			return "", err
		}

		if cdsName != "" {
			if cdsName == config.KindOnceADS {
				cdsName = adsName
			}

			cdsRef, err := config.MarshalRef(cdsName)
			if err != nil {
				return "", err
			}
			xds, err := config.MarshalKindOnceXDS("cds", cdsRef)
			if err != nil {
				return "", err
			}
			_, err = conf.RegisterInit(xds)
			if err != nil {
				return "", err
			}
		}
	}

	if c.LdsConfig != nil {
		ldsName, err := convert_api_v2_core.Convert_ConfigSource(conf, c.LdsConfig)
		if err != nil {
			return "", err
		}
		if ldsName != "" {
			if ldsName == config.KindOnceADS {
				ldsName = adsName
			}

			ldsRef, err := config.MarshalRef(ldsName)
			if err != nil {
				return "", err
			}
			xds, err := config.MarshalKindOnceXDS("lds", ldsRef)
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

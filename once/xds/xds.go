package xds

import (
	"context"
	"fmt"

	"github.com/wzshiming/envoy/once/ads"
)

type XDS struct {
	ads *ads.ADS
	xds string
}

func NewXDS(ads *ads.ADS, xds string) *XDS {
	return &XDS{
		ads: ads,
		xds: xds,
	}
}

func (x *XDS) Do(ctx context.Context) error {
	err := x.ads.Do(ctx)
	if err != nil {
		return err
	}
	switch x.xds {
	case "cds":
		return x.ads.StartCDS()
	case "lds":
		return x.ads.StartLDS()
	}
	return fmt.Errorf("%q is not define in XDS", x.xds)
}

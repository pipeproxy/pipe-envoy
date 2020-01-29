package xds

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/wzshiming/pipe"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/golang/protobuf/jsonpb"
	"github.com/wzshiming/envoy/ads"
	"github.com/wzshiming/envoy/config"
	convert_api_v2 "github.com/wzshiming/envoy/convert/api/v2"
	"github.com/wzshiming/envoy/internal/logger"
)

var marshal = &jsonpb.Marshaler{OrigName: true, Indent: "  "}

type XDS struct {
	ads  *ads.Client
	conf *config.ConfigCtx
	once sync.Once
	ctx  context.Context
}

func NewXDS(conf *config.ConfigCtx, adsConfig *ads.Config) (*XDS, error) {
	x := &XDS{}
	if adsConfig == nil {
		adsConfig = &ads.Config{}
	}
	x.conf = conf

	adsConfig.HandleCDS = x.handleCDS
	adsConfig.HandleRDS = x.handleRDS
	adsConfig.HandleLDS = x.handleLDS
	adsConfig.HandleEDS = x.handleEDS

	cli, err := ads.Dial("", "", adsConfig)
	if err != nil {
		return nil, err
	}
	x.ads = cli

	return x, nil
}

func (x *XDS) Do(ctx context.Context) {
	x.once.Do(func() {
		x.ctx = ctx
		go x.do(ctx)
	})
}

func (x *XDS) do(ctx context.Context) {
	logger.Info("start xds")
	err := x.ads.Run()
	if err != nil {
		logger.Fatalln(err)
	}
}

func (x *XDS) handleCDS(cds []*envoy_api_v2.Cluster) {

	for _, cluster := range cds {
		name, err := convert_api_v2.Convert_Cluster(x.conf, cluster)
		if err != nil {
			logger.Error(err)
		}
		_ = name
	}

	if len(x.conf.EDS) != 0 {
		x.ads.SendRsc(ads.EndpointType, x.conf.EDS)
	}
}

func (x *XDS) handleEDS(eds []*envoy_api_v2.ClusterLoadAssignment) {
	for _, endpoint := range eds {
		name, err := convert_api_v2.Convert_ClusterLoadAssignment(x.conf, endpoint)
		if err != nil {
			logger.Error(err)
		}
		_ = name
	}

	_ = x.ads.SendRsc(ads.ListenerType, nil)
}

func (x *XDS) handleLDS(lds []*envoy_api_v2.Listener) {

	for _, listener := range lds {

		name, err := convert_api_v2.Convert_Listener(x.conf, listener)
		if err != nil {
			logger.Error(err)
			return
		}
		if name != "" {
			err = x.conf.RegisterService(name)
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}

	if len(x.conf.RDS) != 0 {
		x.ads.SendRsc(ads.RouteType, x.conf.RDS)
	}
}

func (x *XDS) handleRDS(rds []*envoy_api_v2.RouteConfiguration) {
	for _, route := range rds {

		name, err := convert_api_v2.Convert_RouteConfiguration(x.conf, route)
		if err != nil {
			logger.Error(err)
			return
		}

		_ = name
	}

	err := x.reload()
	if err != nil {
		logger.Error(err)
	}
}

func (x *XDS) reload() error {
	pipeConfig, _ := json.MarshalIndent(x.conf, "", " ")

	logger.Info(string(pipeConfig))

	conf, err := json.Marshal(x.conf)
	if err != nil {
		return err
	}
	p, ok := pipe.GetPipeWithContext(x.ctx)
	if !ok {
		return fmt.Errorf("not get pipe")
	}

	err = p.Reload(conf)
	if err != nil {
		return err
	}
	return nil
}

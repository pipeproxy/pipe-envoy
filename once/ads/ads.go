package ads

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	"github.com/wzshiming/envoy/config"
	"github.com/wzshiming/envoy/config/clean"
	"github.com/wzshiming/envoy/convert"
	"github.com/wzshiming/envoy/internal/client/ads"
	"github.com/wzshiming/envoy/internal/logger"
	"github.com/wzshiming/pipe"
)

type ADS struct {
	ads  *ads.Client
	conf *config.ConfigCtx
	once sync.Once
	ctx  context.Context

	ch   chan struct{}
	lock uint32
}

func NewADS(conf *config.ConfigCtx, adsConfig *ads.Config) (*ADS, error) {
	a := &ADS{}
	if adsConfig == nil {
		adsConfig = &ads.Config{}
	}
	a.ch = make(chan struct{}, 1)
	a.conf = conf

	adsConfig.HandleCDS = a.handleCDS
	adsConfig.HandleRDS = a.handleRDS
	adsConfig.HandleLDS = a.handleLDS
	adsConfig.HandleEDS = a.handleEDS
	adsConfig.HandleSDS = a.handleSDS

	cli, err := ads.NewClient("", "", adsConfig)
	if err != nil {
		return nil, err
	}
	a.ads = cli

	return a, nil
}

func (a *ADS) Do(ctx context.Context) error {
	if a.ctx == nil {
		a.ctx = ctx
	}
	a.do()
	return nil
}

func (a *ADS) do() {
	a.once.Do(func() {
		a.start()
	})
}

func (a *ADS) start() {
	logger.Info("start xds")
	err := a.ads.Start()
	if err != nil {
		logger.Fatalln(err)
	}
}

func (a *ADS) StartCDS() error {
	a.do()
	err := a.ads.SendRsc(ads.ClusterType, nil)
	if err != nil {
		return err
	}
	a.keepRsc()
	return nil
}

func (a *ADS) StartLDS() error {
	a.do()
	err := a.ads.SendRsc(ads.ListenerType, nil)
	if err != nil {
		return err
	}
	a.keepRsc()
	return nil
}

func (a *ADS) waitRsc() bool {
	if !atomic.CompareAndSwapUint32(&a.lock, 0, 1) {
		return false
	}
	defer atomic.StoreUint32(&a.lock, 0)
	b := false
	for {
		select {
		case <-time.After(time.Second / 50):
			return b
		case <-a.ch:
			b = true
		}
	}
}

func (a *ADS) keepRsc() {

	eds := a.conf.ResetEDS()
	if len(eds) != 0 {
		a.ads.SendRsc(ads.EndpointType, eds)
	}

	rds := a.conf.ResetRDS()
	if len(rds) != 0 {
		a.ads.SendRsc(ads.RouteType, rds)
	}

	sds := a.conf.ResetSDS()
	if len(sds) != 0 {
		a.ads.SendRsc(ads.SecretType, sds)
	}

	select {
	case a.ch <- struct{}{}:
	default:
	}

	if a.waitRsc() {
		go func() {
			err := a.reload()
			if err != nil {
				logger.Errorf("reload error %s", err.Error())
			}
		}()
	}
}

func (a *ADS) handleCDS(cds []*envoy_api_v2.Cluster) {
	for _, cluster := range cds {
		name, err := convert.Convert_api_v2_Cluster(a.conf, cluster)
		if err != nil {
			logger.Error(err)
		}
		_ = name
	}

	a.keepRsc()
}

func (a *ADS) handleEDS(eds []*envoy_api_v2.ClusterLoadAssignment) {
	for _, endpoint := range eds {
		name, err := convert.Convert_api_v2_ClusterLoadAssignment(a.conf, endpoint)
		if err != nil {
			logger.Error(err)
		}
		_ = name
	}

	a.keepRsc()
}

func (a *ADS) handleLDS(lds []*envoy_api_v2.Listener) {
	for _, listener := range lds {

		name, err := convert.Convert_api_v2_Listener(a.conf, listener)
		if err != nil {
			logger.Error(err)
			return
		}
		if name != "" {
			err = a.conf.RegisterService(name)
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}

	a.keepRsc()
}

func (a *ADS) handleRDS(rds []*envoy_api_v2.RouteConfiguration) {
	for _, route := range rds {

		name, err := convert.Convert_api_v2_RouteConfiguration(a.conf, route)
		if err != nil {
			logger.Error(err)
			return
		}

		_ = name
	}

	a.keepRsc()
}

func (a *ADS) handleSDS(sds []*envoy_api_v2_auth.Secret) {
	for _, secret := range sds {
		name, err := convert.Convert_api_v2_auth_Secret(a.conf, secret)
		if err != nil {
			logger.Error(err)
			return
		}
		_ = name
	}

	a.keepRsc()
}

func (a *ADS) reload() error {
	conf, err := json.Marshal(a.conf)
	if err != nil {
		return err
	}
	p, ok := pipe.GetPipeWithContext(a.ctx)
	if !ok {
		return fmt.Errorf("not get pipe")
	}

	conf0, err := clean.Clean(conf)
	if err != nil {
		logger.Error(err)
	} else {
		conf = conf0
	}

	err = p.Reload(conf)
	if err != nil {
		return err
	}
	return nil
}

package bootstrap

import (
	envoy_config_bootstrap_v2 "github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2"
	"github.com/wzshiming/envoy/bootstrap/admin"
	"github.com/wzshiming/envoy/bootstrap/resources/dynamic"
	"github.com/wzshiming/envoy/bootstrap/resources/static"
)

type Bootstrap struct {
	admin            *admin.Admin
	staticResources  *static.StaticResources
	dynamicResources *dynamic.DynamicResources
}

func NewBootstrap(config *envoy_config_bootstrap_v2.Bootstrap) (*Bootstrap, error) {
	b := &Bootstrap{}
	admin, err := admin.NewAdmin(config.Admin)
	if err != nil {
		return nil, err
	}
	b.admin = admin
	staticResources, err := static.NewStaticResources(config.StaticResources)
	if err != nil {
		return nil, err
	}
	b.staticResources = staticResources

	dynamicResources, err := dynamic.NewDynamicResources(config.DynamicResources)
	if err != nil {
		return nil, err
	}
	b.dynamicResources = dynamicResources

	return b, nil
}

func (b *Bootstrap) Start() error {
	err := b.admin.Start()
	if err != nil {
		return err
	}
	err = b.staticResources.Start()
	if err != nil {
		return err
	}
	err = b.dynamicResources.Start()
	if err != nil {
		return err
	}

	return nil
}

type BootstrapCtx struct {
	Bootstrap *Bootstrap
}

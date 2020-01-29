package convert

import (
	"context"
	"encoding/json"

	"github.com/wzshiming/envoy/config"
	convert_config_bootstrap_v2 "github.com/wzshiming/envoy/convert/config/bootstrap/v2"
)

type xdsCtxKeyType int

func GetXdsWithContext(ctx context.Context) (*config.ConfigCtx, bool) {
	i := ctx.Value(xdsCtxKeyType(0))
	if i == nil {
		return nil, false
	}
	p, ok := i.(*config.ConfigCtx)
	return p, ok
}

func ConvertXDS(ctx context.Context, data []byte) (context.Context, []byte, error) {
	conf, err := config.UnmarshalBootstrap(data)
	if err != nil {
		return nil, nil, err
	}

	c := &config.ConfigCtx{}
	c.Ctx = context.WithValue(ctx, xdsCtxKeyType(0), c)

	_, err = convert_config_bootstrap_v2.Convert_Bootstrap(c, conf)
	if err != nil {
		return nil, nil, err
	}

	pipeConfig, err := json.Marshal(c)
	if err != nil {
		return nil, nil, err
	}
	return c.Ctx, pipeConfig, nil
}

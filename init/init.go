package init

import (
	_ "github.com/wzshiming/envoy/pipe/http/access_log"
	_ "github.com/wzshiming/envoy/pipe/once/access_log"
	_ "github.com/wzshiming/envoy/pipe/once/ads"
	_ "github.com/wzshiming/envoy/pipe/once/xds"
	_ "github.com/wzshiming/envoy/pipe/service/none"
	_ "github.com/wzshiming/envoy/pipe/tls/merge"
	_ "github.com/wzshiming/envoy/pipe/tls/validation"
)

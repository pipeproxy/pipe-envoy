package init

import (
	_ "github.com/wzshiming/pipe/init"

	_ "github.com/wzshiming/envoy/http/access_log"
	_ "github.com/wzshiming/envoy/once/access_log"
	_ "github.com/wzshiming/envoy/once/ads"
	_ "github.com/wzshiming/envoy/once/xds"
	_ "github.com/wzshiming/envoy/service/none"
)

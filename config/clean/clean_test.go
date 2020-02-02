package clean

import (
	"reflect"
	"testing"
)

func TestClean(t *testing.T) {
	type args struct {
		config []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			args: args{[]byte(`{"Pipe":{"@Kind":"github.com/wzshiming/pipe/service.Service@multi","Multi":[{"@Ref":"auto@3c10ae54c139cc01b1791f8e5811f88c"},{"@Ref":"xds@listener-9000"},{"@Ref":"xds@listener-9001"},{"@Ref":"xds@listener-9002"},{"@Ref":"xds@listener-9003"}]},"Init":[{"@Kind":"github.com/wzshiming/pipe/once.Once@xds","XDS":"cds","ADS":{"@Ref":"ads@ads"}},{"@Kind":"github.com/wzshiming/pipe/once.Once@xds","XDS":"lds","ADS":{"@Ref":"ads@ads"}}],"Components":[{"@Name":"ads@ads","@Kind":"github.com/wzshiming/pipe/once.Once@ads","NodeID":"test-id","Forward":{"@Ref":"xds@xds_cluster"}},{"@Name":"auto@00b3b4375ebd3de04048f67a803d7f8b","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v1-3"}]},{"@Name":"auto@08731d6594ef41c80f182bc67a6d72ff","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v1-1.route"}]},{"@Name":"auto@15ea8246dd1c2388c3808de404bfc155","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:19000"},{"@Name":"auto@194156017b71b5e651f6ee463de8a1e6","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v1-0.route"}]},{"@Name":"auto@2d126de606e5dbc9594315973d0682c7","@Kind":"github.com/wzshiming/pipe/once.Once@access_log","NodeID":"test-id","LogName":"echo","Forward":{"@Ref":"xds@access_log_cluster"}},{"@Name":"auto@313412820b9d5dfeb3f4ece7c45f667f","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v2-0","Forward":{"@Ref":"xds@cluster-v2-0"}},{"@Name":"auto@36752de623441f5430cef8deeafa9141","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v2-0.route"}]},{"@Name":"auto@3c10ae54c139cc01b1791f8e5811f88c","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"auto@15ea8246dd1c2388c3808de404bfc155"},"Handler":{"@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Kind":"net/http.Handler@log","Output":{"@Kind":"io.WriteCloser@file","Path":"/dev/null"},"Handler":{"@Kind":"net/http.Handler@mux","Routes":[{"Path":"/expvar/","Handler":{"@Kind":"expvar"}},{"Prefix":"/pprof/","Handler":{"@Kind":"pprof"}},{"Prefix":"/config_dump/","Handler":{"@Kind":"config_dump"}}],"NotFound":{"@Kind":"net/http.Handler@multi","Multi":[{"@Kind":"net/http.Handler@add_response_header","Key":"Content-Type","Value":"text/html; charset=utf-8"},{"@Kind":"net/http.Handler@direct","Code":200,"Body":{"@Kind":"io.ReadCloser@inline","Data":"\n\u003cpre\u003e\n\u003ca href=\"/expvar/\"\u003e/expvar/\u003c/a\u003e\n\u003ca href=\"/pprof/\"\u003e/pprof/\u003c/a\u003e\n\u003ca href=\"/config_dump/\"\u003e/config_dump/\u003c/a\u003e\n\u003c/pre\u003e\n"}}]}}}}},{"@Name":"auto@4378e4b69d5ac8b9c39ea81faafe211c","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@a9e526278857a5e4c6cbb41de28abbf5"}]},{"@Name":"auto@4a1cdacea53e437ddbe9591e75db903a","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18000"},{"@Name":"auto@599bbd2f4f12f9905de3282340a3496e","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v1-0"}},{"@Name":"auto@5ec5302935e12dad723c80c5c56c74e6","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v2-2"}]},{"@Name":"auto@654d40207ac22167f5c7c7ece62f306a","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v2-0"}},{"@Name":"auto@6f9088714c22ebf0e4199fb53c4963ea","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v1-0","Forward":{"@Ref":"xds@cluster-v1-0"}},{"@Name":"auto@775360bade731cb7179ebcb4943117e8","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v0-0"}},{"@Name":"auto@787738150d84d5ab78f7149fd1f66bae","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v0-0.route"}]},{"@Name":"auto@7ab61da0d3ed8dcbd1e23413888d24cb","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v2-1"}},{"@Name":"auto@8047db928362be3c2c0d8aa53156173e","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v2-1.route"}]},{"@Name":"auto@8e50b72ce4058095da1b98dbc2d38f55","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v0-1.route"}]},{"@Name":"auto@a9e526278857a5e4c6cbb41de28abbf5","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18080"},{"@Name":"auto@aaea91e287eb1b4500d43de4ccb716cc","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9002"},{"@Name":"auto@b053866037465be2e985d1d0f55b145a","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18090"},{"@Name":"auto@b57215e080166a496a3fb950b9f1adf1","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v0-0","Forward":{"@Ref":"xds@cluster-v0-0"}},{"@Name":"auto@bde5961ddf9f2c9343cdf318efa55461","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v2-1","Forward":{"@Ref":"xds@cluster-v2-1"}},{"@Name":"auto@c1425a25ea1e7039dc7413308f1a626c","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v0-1","Forward":{"@Ref":"xds@cluster-v0-1"}},{"@Name":"auto@c78e74eb43c1eee28149704ab6ee0294","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v1-1","Forward":{"@Ref":"xds@cluster-v1-1"}},{"@Name":"auto@cb7d249ae76fd1e335430ca712671b63","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9000"},{"@Name":"auto@e362ca838e9c7d81bc4cfd9cbfb05c7a","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v2-3"}]},{"@Name":"auto@e62f3956a75f6753704c885a8cb92ade","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v1-2"}]},{"@Name":"auto@eb419710d81389aa033c27c6ae7859e0","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v1-1"}},{"@Name":"auto@ed6d25ffea7811381f2e440761e5751a","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9001"},{"@Name":"auto@f07cce60d9608c52da0d38deacd12225","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v0-2"}]},{"@Name":"auto@f360dde2ee5829dfc4e7436ee92d9c1c","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9003"},{"@Name":"auto@f8bfecc868b58d187c80c9973d347591","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v0-1"}},{"@Name":"auto@ff5c800e1d7278608190d739fc2a97e3","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v0-3"}]},{"@Name":"xds@access_log_cluster","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@b053866037465be2e985d1d0f55b145a"}]},{"@Name":"xds@cluster-v0-0","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v0-1","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v0-2","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v0-3","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v1-0","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v1-1","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v1-2","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v1-3","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-0","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-1","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-2","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-3","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@listener-9000","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9000.listener"},"Handler":{"@Ref":"xds@listener-9000.filter-chains"}},{"@Name":"xds@listener-9000.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@36752de623441f5430cef8deeafa9141"}]},{"@Name":"xds@listener-9000.listener","@Ref":"auto@cb7d249ae76fd1e335430ca712671b63"},{"@Name":"xds@listener-9001","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9001.listener"},"Handler":{"@Ref":"xds@listener-9001.filter-chains"}},{"@Name":"xds@listener-9001.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@8047db928362be3c2c0d8aa53156173e"}]},{"@Name":"xds@listener-9001.listener","@Ref":"auto@ed6d25ffea7811381f2e440761e5751a"},{"@Name":"xds@listener-9002","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9002.listener"},"Handler":{"@Ref":"xds@listener-9002.filter-chains"}},{"@Name":"xds@listener-9002.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@5ec5302935e12dad723c80c5c56c74e6"}]},{"@Name":"xds@listener-9002.listener","@Ref":"auto@aaea91e287eb1b4500d43de4ccb716cc"},{"@Name":"xds@listener-9003","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9003.listener"},"Handler":{"@Ref":"xds@listener-9003.filter-chains"}},{"@Name":"xds@listener-9003.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@e362ca838e9c7d81bc4cfd9cbfb05c7a"}]},{"@Name":"xds@listener-9003.listener","@Ref":"auto@f360dde2ee5829dfc4e7436ee92d9c1c"},{"@Name":"xds@route-v0-0","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v0-0.virtual-host"}]},{"@Name":"xds@route-v0-0.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@775360bade731cb7179ebcb4943117e8"}},{"@Name":"xds@route-v0-0.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@b57215e080166a496a3fb950b9f1adf1"}}]},{"@Name":"xds@route-v0-1","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v0-1.virtual-host"}]},{"@Name":"xds@route-v0-1.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@f8bfecc868b58d187c80c9973d347591"}},{"@Name":"xds@route-v0-1.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@c1425a25ea1e7039dc7413308f1a626c"}}]},{"@Name":"xds@route-v1-0","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v1-0.virtual-host"}]},{"@Name":"xds@route-v1-0.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@599bbd2f4f12f9905de3282340a3496e"}},{"@Name":"xds@route-v1-0.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@6f9088714c22ebf0e4199fb53c4963ea"}}]},{"@Name":"xds@route-v1-1","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v1-1.virtual-host"}]},{"@Name":"xds@route-v1-1.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@eb419710d81389aa033c27c6ae7859e0"}},{"@Name":"xds@route-v1-1.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@c78e74eb43c1eee28149704ab6ee0294"}}]},{"@Name":"xds@route-v2-0","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v2-0.virtual-host"}]},{"@Name":"xds@route-v2-0.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@654d40207ac22167f5c7c7ece62f306a"}},{"@Name":"xds@route-v2-0.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@313412820b9d5dfeb3f4ece7c45f667f"}}]},{"@Name":"xds@route-v2-1","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v2-1.virtual-host"}]},{"@Name":"xds@route-v2-1.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@7ab61da0d3ed8dcbd1e23413888d24cb"}},{"@Name":"xds@route-v2-1.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@bde5961ddf9f2c9343cdf318efa55461"}}]},{"@Name":"xds@xds_cluster","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4a1cdacea53e437ddbe9591e75db903a"}]}]}`)},
			want: []byte(`{"Pipe":{"@Kind":"github.com/wzshiming/pipe/service.Service@multi","Multi":[{"@Ref":"auto@3c10ae54c139cc01b1791f8e5811f88c"},{"@Ref":"xds@listener-9000"},{"@Ref":"xds@listener-9001"},{"@Ref":"xds@listener-9002"},{"@Ref":"xds@listener-9003"}]},"Init":[{"@Kind":"github.com/wzshiming/pipe/once.Once@xds","XDS":"cds","ADS":{"@Ref":"ads@ads"}},{"@Kind":"github.com/wzshiming/pipe/once.Once@xds","XDS":"lds","ADS":{"@Ref":"ads@ads"}}],"Components":[{"@Name":"ads@ads","@Kind":"github.com/wzshiming/pipe/once.Once@ads","NodeID":"test-id","Forward":{"@Ref":"xds@xds_cluster"}},{"@Name":"auto@15ea8246dd1c2388c3808de404bfc155","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:19000"},{"@Name":"auto@2d126de606e5dbc9594315973d0682c7","@Kind":"github.com/wzshiming/pipe/once.Once@access_log","NodeID":"test-id","LogName":"echo","Forward":{"@Ref":"xds@access_log_cluster"}},{"@Name":"auto@313412820b9d5dfeb3f4ece7c45f667f","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v2-0","Forward":{"@Ref":"xds@cluster-v2-0"}},{"@Name":"auto@36752de623441f5430cef8deeafa9141","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v2-0.route"}]},{"@Name":"auto@3c10ae54c139cc01b1791f8e5811f88c","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"auto@15ea8246dd1c2388c3808de404bfc155"},"Handler":{"@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Kind":"net/http.Handler@log","Output":{"@Kind":"io.WriteCloser@file","Path":"/dev/null"},"Handler":{"@Kind":"net/http.Handler@mux","Routes":[{"Path":"/expvar/","Handler":{"@Kind":"expvar"}},{"Prefix":"/pprof/","Handler":{"@Kind":"pprof"}},{"Prefix":"/config_dump/","Handler":{"@Kind":"config_dump"}}],"NotFound":{"@Kind":"net/http.Handler@multi","Multi":[{"@Kind":"net/http.Handler@add_response_header","Key":"Content-Type","Value":"text/html; charset=utf-8"},{"@Kind":"net/http.Handler@direct","Code":200,"Body":{"@Kind":"io.ReadCloser@inline","Data":"\n\u003cpre\u003e\n\u003ca href=\"/expvar/\"\u003e/expvar/\u003c/a\u003e\n\u003ca href=\"/pprof/\"\u003e/pprof/\u003c/a\u003e\n\u003ca href=\"/config_dump/\"\u003e/config_dump/\u003c/a\u003e\n\u003c/pre\u003e\n"}}]}}}}},{"@Name":"auto@4378e4b69d5ac8b9c39ea81faafe211c","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@a9e526278857a5e4c6cbb41de28abbf5"}]},{"@Name":"auto@4a1cdacea53e437ddbe9591e75db903a","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18000"},{"@Name":"auto@5ec5302935e12dad723c80c5c56c74e6","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v2-2"}]},{"@Name":"auto@654d40207ac22167f5c7c7ece62f306a","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v2-0"}},{"@Name":"auto@7ab61da0d3ed8dcbd1e23413888d24cb","@Kind":"net/http.Handler@access_log","AccessLog":{"@Ref":"auto@2d126de606e5dbc9594315973d0682c7"},"Handler":{"@Ref":"xds@route-v2-1"}},{"@Name":"auto@8047db928362be3c2c0d8aa53156173e","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@route-v2-1.route"}]},{"@Name":"auto@a9e526278857a5e4c6cbb41de28abbf5","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18080"},{"@Name":"auto@aaea91e287eb1b4500d43de4ccb716cc","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9002"},{"@Name":"auto@b053866037465be2e985d1d0f55b145a","@Kind":"github.com/wzshiming/pipe/stream.Handler@forward","Network":"tcp","Address":"127.0.0.1:18090"},{"@Name":"auto@bde5961ddf9f2c9343cdf318efa55461","@Kind":"net/http.Handler@forward","Pass":"http://cluster-v2-1","Forward":{"@Ref":"xds@cluster-v2-1"}},{"@Name":"auto@cb7d249ae76fd1e335430ca712671b63","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9000"},{"@Name":"auto@e362ca838e9c7d81bc4cfd9cbfb05c7a","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"xds@cluster-v2-3"}]},{"@Name":"auto@ed6d25ffea7811381f2e440761e5751a","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9001"},{"@Name":"auto@f360dde2ee5829dfc4e7436ee92d9c1c","@Kind":"github.com/wzshiming/pipe/listener.ListenConfig@network","Network":"tcp","Address":"127.0.0.1:9003"},{"@Name":"xds@access_log_cluster","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@b053866037465be2e985d1d0f55b145a"}]},{"@Name":"xds@cluster-v2-0","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-1","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-2","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@cluster-v2-3","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4378e4b69d5ac8b9c39ea81faafe211c"}]},{"@Name":"xds@listener-9000","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9000.listener"},"Handler":{"@Ref":"xds@listener-9000.filter-chains"}},{"@Name":"xds@listener-9000.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@36752de623441f5430cef8deeafa9141"}]},{"@Name":"xds@listener-9000.listener","@Ref":"auto@cb7d249ae76fd1e335430ca712671b63"},{"@Name":"xds@listener-9001","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9001.listener"},"Handler":{"@Ref":"xds@listener-9001.filter-chains"}},{"@Name":"xds@listener-9001.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@8047db928362be3c2c0d8aa53156173e"}]},{"@Name":"xds@listener-9001.listener","@Ref":"auto@ed6d25ffea7811381f2e440761e5751a"},{"@Name":"xds@listener-9002","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9002.listener"},"Handler":{"@Ref":"xds@listener-9002.filter-chains"}},{"@Name":"xds@listener-9002.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@5ec5302935e12dad723c80c5c56c74e6"}]},{"@Name":"xds@listener-9002.listener","@Ref":"auto@aaea91e287eb1b4500d43de4ccb716cc"},{"@Name":"xds@listener-9003","@Kind":"github.com/wzshiming/pipe/service.Service@server","Listener":{"@Ref":"xds@listener-9003.listener"},"Handler":{"@Ref":"xds@listener-9003.filter-chains"}},{"@Name":"xds@listener-9003.filter-chains","@Kind":"github.com/wzshiming/pipe/stream.Handler@multi","Multi":[{"@Ref":"auto@e362ca838e9c7d81bc4cfd9cbfb05c7a"}]},{"@Name":"xds@listener-9003.listener","@Ref":"auto@f360dde2ee5829dfc4e7436ee92d9c1c"},{"@Name":"xds@route-v2-0","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v2-0.virtual-host"}]},{"@Name":"xds@route-v2-0.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@654d40207ac22167f5c7c7ece62f306a"}},{"@Name":"xds@route-v2-0.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@313412820b9d5dfeb3f4ece7c45f667f"}}]},{"@Name":"xds@route-v2-1","@Kind":"net/http.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"xds@route-v2-1.virtual-host"}]},{"@Name":"xds@route-v2-1.route","@Kind":"github.com/wzshiming/pipe/stream.Handler@http","Handler":{"@Ref":"auto@7ab61da0d3ed8dcbd1e23413888d24cb"}},{"@Name":"xds@route-v2-1.virtual-host","@Kind":"net/http.Handler@mux","Routes":[{"Prefix":"/","Handler":{"@Ref":"auto@bde5961ddf9f2c9343cdf318efa55461"}}]},{"@Name":"xds@xds_cluster","@Kind":"github.com/wzshiming/pipe/stream.Handler@poller","Poller":"round_robin","Handlers":[{"@Ref":"auto@4a1cdacea53e437ddbe9591e75db903a"}]}]}`)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Clean(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("Clean() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clean() got = %v, want %v", got, tt.want)
			}
		})
	}
}

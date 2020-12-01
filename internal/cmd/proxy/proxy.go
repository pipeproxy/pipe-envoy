package proxy

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"syscall"
	"time"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/convert"
	"github.com/pipeproxy/pipe-xds/internal/adsc"
	"github.com/spf13/cobra"
	"github.com/wzshiming/notify"
	"github.com/wzshiming/xds/utils"
	"google.golang.org/grpc/grpclog"
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pkg/cmd"
	"istio.io/istio/pkg/config/constants"
	"istio.io/istio/security/pkg/stsservice/tokenmanager"
	"istio.io/pkg/env"
)

const (
	localHostIPv4 = "127.0.0.1"
	localHostIPv6 = "[::1]"
)

var (
	// Similar with ISTIO_META_, which is used to customize the node metadata - this customizes extra header.
	xdsHeaderPrefix = [...]string{"XDS_HEADER_", "ISTIO_META_"}
	certs           = ""
	metadata        = map[string]interface{}{
		"CLUSTER_ID": clusterIDVar.Get(),
	}
	tmp         = "./tmp/"
	ctx, cancel = context.WithCancel(context.Background())

	//role               = &model.Proxy{}
	DNSDomain string
	proxyType = model.SidecarProxy

	stsPort            int
	tokenManagerPlugin string

	meshConfigFile string

	// proxy config flags (named identically)
	serviceCluster         string
	proxyLogLevel          string
	proxyComponentLogLevel string
	concurrency            int
	templateFile           string
	outlierLogPath         string

	instanceIPVar   = env.RegisterStringVar("INSTANCE_IP", "", "")
	podNameVar      = env.RegisterStringVar("POD_NAME", "", "")
	podNamespaceVar = env.RegisterStringVar("POD_NAMESPACE", "", "")
	clusterIDVar    = env.RegisterStringVar("ISTIO_META_CLUSTER_ID", "", "")

	url = env.RegisterStringVar("ISTIOD_ADDRESS", "istiod.istio-system.svc:15010", "")

	proxyCmd = &cobra.Command{
		Use:   "proxy",
		Short: "Envoy proxy agent",
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			// Allow unknown flags for backward-compatibility.
			UnknownFlags: true,
		},
		RunE: func(c *cobra.Command, args []string) error {
			cmd.PrintFlags(c.Flags())
			grpclog.SetLoggerV2(grpclog.NewLoggerV2(ioutil.Discard, ioutil.Discard, ioutil.Discard))

			// Extract pod variables.
			podName := podNameVar.Get()
			podNamespace := podNamespaceVar.Get()
			podIP := net.ParseIP(instanceIPVar.Get()) // protobuf encoding of IP_ADDRESS type

			if len(args) > 0 {
				proxyType = model.NodeType(args[0])
				if !model.IsApplicationNodeType(proxyType) {
					return fmt.Errorf("Invalid role Type: " + string(proxyType))
				}
			}

			for _, env := range os.Environ() {
				kv := strings.SplitN(env, "=", 2)
				if len(kv) == 2 {
					for _, prefix := range xdsHeaderPrefix {
						if strings.HasPrefix(kv[0], prefix) {
							metadata[kv[0][len(xdsHeaderPrefix):]] = kv[1]
							break
						}
					}
				}
			}

			metadataBytes, _ := json.Marshal(metadata)
			metadataJSON := string(metadataBytes)
			log.Println("xds server", url)
			log.Println("metadata", metadataJSON)

			nodeId := fmt.Sprintf("%s~%s~%s.%s~%s", proxyType, podIP, podName, podNamespace, DNSDomain)

			notify.OnSlice([]os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}, cancel)

			var tlsConfig *tls.Config
			var err error
			if certs != "" {
				t, err := utils.TlsConfigFromDir(certs)
				if err != nil {
					log.Fatalln(err)
				}
				tlsConfig = t
			}

			err = os.MkdirAll(tmp, 0755)
			if err != nil {
				log.Fatal(err)
			}

			ads := adsc.NewADSC(url.Get(), tlsConfig, utils.NodeConfig{
				NodeID:   nodeId,
				Metadata: metadata,
			})

			cc := config.NewConfigCtx(ctx, ads, tmp, time.Second/2)

			ads.HandleCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
				for name, cluster := range clusters {
					d, err := convert.Convert_config_cluster_v3_Cluster(cc, cluster)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterCDS("cds."+name, d, cluster)
				}
			}
			ads.HandleEDSCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
				for name, cluster := range clusters {
					d, err := convert.Convert_config_cluster_v3_Cluster(cc, cluster)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterCDS("cds.eds."+name, d, cluster)
				}
			}
			ads.HandleEDS = func(endpoints map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment) {
				for name, endpoint := range endpoints {
					d, err := convert.Convert_config_endpoint_v3_ClusterLoadAssignment(cc, endpoint)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterEDS("eds."+name, d, endpoint)
				}
			}
			ads.HandleTcpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
				for name, listener := range listeners {
					d, err := convert.Convert_config_listener_v3_Listener(cc, listener)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterLDS("lds.tcp."+name, d, listener)
				}
			}
			ads.HandleHttpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
				for name, listener := range listeners {
					d, err := convert.Convert_config_listener_v3_Listener(cc, listener)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterLDS("lds.http."+name, d, listener)
				}
			}
			ads.HandleRDS = func(routes map[string]*envoy_config_route_v3.RouteConfiguration) {
				for name, route := range routes {
					d, err := convert.Convert_config_route_v3_RouteConfiguration(cc, route)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterRDS("rds.http."+name, d, route)
				}
			}
			ads.HandleSDS = func(secrets map[string]*envoy_extensions_transport_sockets_tls_v3.Secret) {
				for name, secret := range secrets {
					d, err := convert.Convert_extensions_transport_sockets_tls_v3_Secret(cc, secret)
					if err != nil {
						log.Println(err)
						continue
					}
					cc.RegisterSDS("sds."+name, d, secret)
				}
			}

			go func() {
				for {
					err = ads.Run(ctx)
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Second)
				}
			}()

			for ctx.Err() == nil {
				err = cc.Run(ctx)
				if err != nil {
					log.Println(err)
				}
			}
			return nil
		},
	}
)

func init() {
	proxyCmd.PersistentFlags().StringVar(&DNSDomain, "domain", "",
		"DNS domain suffix. If not provided uses ${POD_NAMESPACE}.svc.cluster.local")
	proxyCmd.PersistentFlags().StringVar(&meshConfigFile, "meshConfig", "./etc/istio/config/mesh",
		"File name for Istio mesh configuration. If not specified, a default mesh will be used. This may be overridden by "+
			"PROXY_CONFIG environment variable or proxy.istio.io/config annotation.")
	proxyCmd.PersistentFlags().IntVar(&stsPort, "stsPort", 0,
		"HTTP Port on which to serve Security Token Service (STS). If zero, STS service will not be provided.")
	proxyCmd.PersistentFlags().StringVar(&tokenManagerPlugin, "tokenManagerPlugin", tokenmanager.GoogleTokenExchange,
		"Token provider specific plugin name.")
	// Flags for proxy configuration
	proxyCmd.PersistentFlags().StringVar(&serviceCluster, "serviceCluster", constants.ServiceClusterName, "Service cluster")
	// Log levels are provided by the library https://github.com/gabime/spdlog, used by Envoy.
	proxyCmd.PersistentFlags().StringVar(&proxyLogLevel, "proxyLogLevel", "warning",
		fmt.Sprintf("The log level used to start the Envoy proxy (choose from {%s, %s, %s, %s, %s, %s, %s})",
			"trace", "debug", "info", "warning", "error", "critical", "off"))
	proxyCmd.PersistentFlags().IntVar(&concurrency, "concurrency", 0, "number of worker threads to run")
	proxyCmd.PersistentFlags().StringVar(&proxyComponentLogLevel, "proxyComponentLogLevel", "misc:error",
		"The component log level used to start the Envoy proxy")
	proxyCmd.PersistentFlags().StringVar(&templateFile, "templateFile", "",
		"Go template bootstrap config")
	proxyCmd.PersistentFlags().StringVar(&outlierLogPath, "outlierLogPath", "",
		"The log path for outlier detection")

}

func GetCommand() *cobra.Command {
	return proxyCmd
}

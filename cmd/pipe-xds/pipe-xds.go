package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"syscall"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/pipeproxy/pipe-xds/config"
	"github.com/pipeproxy/pipe-xds/convert"
	"github.com/pipeproxy/pipe-xds/internal/adsc"
	"github.com/wzshiming/notify"
	"github.com/wzshiming/xds/utils"
)

var (
	url      = "127.0.0.1:15010"
	certs    = ""
	nodeId   = getNodeType(os.Getenv("POD_NAME")) + "~$INSTANCE_IP~$POD_NAME.$POD_NAMESPACE~$POD_NAMESPACE.svc.$ISTIO_META_MESH_ID"
	metadata = map[string]interface{}{
		"CLUSTER_ID": "Kubernetes",
	}
	tmp         = "./tmp/"
	ctx, cancel = context.WithCancel(context.Background())
)

const (
	metaPrefix = "ISTIO_META_"
)

func init() {
	flag.StringVar(&url, "u", url, "xds server")
	flag.StringVar(&certs, "c", certs, "certs folder {cert-chain.pem,key.pem,root-cert.pem}")
	flag.StringVar(&nodeId, "n", nodeId, "node id")
	for _, env := range os.Environ() {
		kv := strings.SplitN(env, "=", 2)
		if len(kv) == 2 && strings.HasPrefix(kv[0], metaPrefix) {
			metadata[kv[0][len(metaPrefix):]] = kv[1]
		}
	}

	metadataBytes, _ := json.Marshal(metadata)
	metadataJSON := string(metadataBytes)
	flag.StringVar(&metadataJSON, "m", metadataJSON, "node metadata")
	flag.StringVar(&tmp, "t", tmp, "template config")
	flag.Parse()

	err := json.Unmarshal([]byte(metadataJSON), &metadata)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("xds server", url)
	nodeId = os.ExpandEnv(nodeId)
	log.Println("node id", nodeId)

	log.Println("metadata", metadataJSON)
}

func main() {
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

	ads := adsc.NewADSC(url, tlsConfig, utils.NodeConfig{
		NodeID:   nodeId,
		Metadata: metadata,
	})

	c := config.NewConfigCtx(ctx, ads, tmp)

	ads.HandleCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
		for name, cluster := range clusters {
			d, err := convert.Convert_config_cluster_v3_Cluster(c, cluster)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterCDS("cds."+name, d, cluster)
		}
	}
	ads.HandleEDSCDS = func(clusters map[string]*envoy_config_cluster_v3.Cluster) {
		for name, cluster := range clusters {
			d, err := convert.Convert_config_cluster_v3_Cluster(c, cluster)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterCDS("cds.eds."+name, d, cluster)
		}
	}
	ads.HandleEDS = func(endpoints map[string]*envoy_config_endpoint_v3.ClusterLoadAssignment) {
		for name, endpoint := range endpoints {
			d, err := convert.Convert_config_endpoint_v3_ClusterLoadAssignment(c, endpoint)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterEDS("eds."+name, d, endpoint)
		}
	}
	ads.HandleTcpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
		for name, listener := range listeners {
			d, err := convert.Convert_config_listener_v3_Listener(c, listener)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterLDS("lds.tcp."+name, d, listener)
		}
	}
	ads.HandleHttpLDS = func(listeners map[string]*envoy_config_listener_v3.Listener) {
		for name, listener := range listeners {
			d, err := convert.Convert_config_listener_v3_Listener(c, listener)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterLDS("lds.http."+name, d, listener)
		}
	}
	ads.HandleRDS = func(routes map[string]*envoy_config_route_v3.RouteConfiguration) {
		for name, route := range routes {
			d, err := convert.Convert_config_route_v3_RouteConfiguration(c, route)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterRDS("rds.http."+name, d, route)
		}
	}
	ads.HandleSDS = func(secrets map[string]*envoy_extensions_transport_sockets_tls_v3.Secret) {
		for name, secret := range secrets {
			d, err := convert.Convert_extensions_transport_sockets_tls_v3_Secret(c, secret)
			if err != nil {
				log.Println(err)
				continue
			}
			c.RegisterSDS("sds."+name, d, secret)
		}
	}

	err = c.Start(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = ads.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func getNodeType(podname string) string {
	const prefix = "istio-"
	if strings.HasPrefix(podname, prefix) {
		podname = podname[len(prefix):]
	}
	if strings.HasPrefix(podname, "ingressgateway") || strings.HasPrefix(podname, "egressgateway") {
		return "router"
	}
	if strings.HasPrefix(podname, "ingress") {
		return "ingress"
	}
	return "sidecar"
}

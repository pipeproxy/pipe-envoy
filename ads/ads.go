package ads

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"time"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_service_discovery_v2 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/golang/protobuf/proto"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Config for the Client connection.
type Config struct {
	NodeID string

	// Namespace defaults to 'default'
	Namespace string

	// Workload defaults to 'test'
	Workload string

	// NodeType defaults to sidecar. "ingress" and "router" are also supported.
	NodeType string

	// IP is currently the primary key used to locate inbound configs. It is sent by client,
	// must match a known endpoint IP. Tests can use a ServiceEntry to register fake IPs.
	IP string

	// Metadate includes additional metadata for the node
	Metadate *structpb.Struct

	ContextDialer func(context.Context, string) (net.Conn, error)

	HandleCDS func([]*envoy_api_v2.Cluster)
	HandleRDS func([]*envoy_api_v2.RouteConfiguration)
	HandleLDS func([]*envoy_api_v2.Listener)
	HandleEDS func([]*envoy_api_v2.ClusterLoadAssignment)
}

// Client implements a client for ADS.
type Client struct {
	stream envoy_service_discovery_v2.AggregatedDiscoveryService_StreamAggregatedResourcesClient

	conn *grpc.ClientConn

	nodeID string

	certDir string
	url     string

	metadata *structpb.Struct

	ContextDialer func(context.Context, string) (net.Conn, error)

	HandleCDS func([]*envoy_api_v2.Cluster)
	HandleRDS func([]*envoy_api_v2.RouteConfiguration)
	HandleLDS func([]*envoy_api_v2.Listener)
	HandleEDS func([]*envoy_api_v2.ClusterLoadAssignment)
}

const (
	typePrefix          = "type.googleapis.com/envoy.api.v2."
	discoveryTypePrefix = "type.googleapis.com/envoy.service.discovery.v2."

	// ClusterType is used for cluster discovery. first request received
	ClusterType = typePrefix + "Cluster"
	// EndpointType is used for Client endpoint discovery. second request.
	EndpointType = typePrefix + "ClusterLoadAssignment"
	// ListenerType is sent after clusters and endpoints.
	ListenerType = typePrefix + "Listener"
	// RouteType is sent after listeners.
	RouteType = typePrefix + "RouteConfiguration"

	SecretType = typePrefix + "auth.Secret"

	RuntimeType = discoveryTypePrefix + "Runtime"

	// AnyType is used only by ADS
	AnyType = ""
)

// Dial connects to a Client server, with optional TLS authentication if a cert dir is specified.
func Dial(url string, certDir string, opts *Config) (*Client, error) {
	ads := &Client{
		certDir: certDir,
		url:     url,
	}
	if opts == nil {
		opts = &Config{}
	}

	if opts.NodeID != "" {
		ads.nodeID = opts.NodeID
	} else {
		if opts.Namespace == "" {
			opts.Namespace = "default"
		}
		if opts.NodeType == "" {
			opts.NodeType = "sidecar"
		}
		if opts.IP == "" {
			opts.IP = getPrivateIPIfAvailable().String()
		}
		if opts.Workload == "" {
			opts.Workload = "test-1"
		}
		ads.nodeID = fmt.Sprintf("%s~%s~%s.%s~%s.svc.cluster.local", opts.NodeType, opts.IP,
			opts.Workload, opts.Namespace, opts.Namespace)
	}

	ads.metadata = opts.Metadate
	ads.ContextDialer = opts.ContextDialer
	ads.HandleCDS = opts.HandleCDS
	ads.HandleRDS = opts.HandleRDS
	ads.HandleLDS = opts.HandleLDS
	ads.HandleEDS = opts.HandleEDS

	return ads, nil
}

// Returns a private IP core, or unspecified IP (0.0.0.0) if no IP is available
func getPrivateIPIfAvailable() net.IP {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		default:
			continue
		}
		if !ip.IsLoopback() {
			return ip
		}
	}
	return net.IPv4zero
}

func tlsConfig(certDir string) (*tls.Config, error) {
	clientCert, err := tls.LoadX509KeyPair(filepath.Join(certDir+"cert-chain.pem"),
		filepath.Join(certDir, "key.pem"))
	if err != nil {
		return nil, err
	}

	serverCABytes, err := ioutil.ReadFile(filepath.Join(certDir, "root-cert.pem"))
	if err != nil {
		return nil, err
	}
	serverCAs := x509.NewCertPool()
	if ok := serverCAs.AppendCertsFromPEM(serverCABytes); !ok {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      serverCAs,
		ServerName:   "istio-pilot.istio-system",
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			return nil
		},
	}, nil
}

// Close the once.
func (a *Client) Close() error {
	if a.stream != nil {
		a.stream.CloseSend()
	}
	return a.conn.Close()
}

// Run the ADS client.
func (a *Client) Run() error {

	opts := []grpc.DialOption{}
	if len(a.certDir) != 0 {
		tlsCfg, err := tlsConfig(a.certDir)
		if err != nil {
			return err
		}
		creds := credentials.NewTLS(tlsCfg)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if a.ContextDialer != nil {
		opts = append(opts, grpc.WithContextDialer(a.ContextDialer))
	}
	conn, err := grpc.Dial(a.url, opts...)
	if err != nil {
		return err
	}
	a.conn = conn
	xds := envoy_service_discovery_v2.NewAggregatedDiscoveryServiceClient(a.conn)
	edsstr, err := xds.StreamAggregatedResources(context.Background())
	if err != nil {
		return err
	}
	a.stream = edsstr

	err = a.watch()
	if err != nil {
		return err
	}
	return a.handleRecv()
}

func (a *Client) handleRecv() error {
	for {
		msg, err := a.stream.Recv()
		if err != nil {
			return fmt.Errorf("connection closed %s: error: %w", a.nodeID, err)
		}
		// logger.Info(msg.TypeUrl)
		lds := []*envoy_api_v2.Listener{}
		cds := []*envoy_api_v2.Cluster{}
		rds := []*envoy_api_v2.RouteConfiguration{}
		eds := []*envoy_api_v2.ClusterLoadAssignment{}
		for _, rsc := range msg.Resources {
			valBytes := rsc.Value
			switch rsc.TypeUrl {
			case ListenerType:
				ll := &envoy_api_v2.Listener{}
				_ = proto.Unmarshal(valBytes, ll)
				lds = append(lds, ll)
			case ClusterType:
				ll := &envoy_api_v2.Cluster{}
				_ = proto.Unmarshal(valBytes, ll)
				cds = append(cds, ll)
			case EndpointType:
				ll := &envoy_api_v2.ClusterLoadAssignment{}
				_ = proto.Unmarshal(valBytes, ll)
				eds = append(eds, ll)
			case RouteType:
				ll := &envoy_api_v2.RouteConfiguration{}
				_ = proto.Unmarshal(valBytes, ll)
				rds = append(rds, ll)
			}
		}

		a.ack(msg)

		if len(cds) != 0 {
			a.handleCDS(cds)
		}
		if len(eds) != 0 {
			a.handleEDS(eds)
		}
		if len(lds) != 0 {
			a.handleLDS(lds)
		}
		if len(rds) != 0 {
			a.handleRDS(rds)
		}
	}
}

func (a *Client) handleCDS(cds []*envoy_api_v2.Cluster) {
	if a.HandleCDS != nil {
		a.HandleCDS(cds)
	}
}

func (a *Client) handleEDS(eds []*envoy_api_v2.ClusterLoadAssignment) {
	if a.HandleEDS != nil {
		a.HandleEDS(eds)
	}
}

func (a *Client) handleLDS(lds []*envoy_api_v2.Listener) {
	if a.HandleLDS != nil {
		a.HandleLDS(lds)
	}
}

func (a *Client) handleRDS(rds []*envoy_api_v2.RouteConfiguration) {
	if a.HandleRDS != nil {
		a.HandleRDS(rds)
	}
}

func (a *Client) node() *envoy_api_v2_core.Node {
	n := &envoy_api_v2_core.Node{
		Id:       a.nodeID,
		Metadata: a.metadata,
	}
	return n
}

func (a *Client) Send(req *envoy_api_v2.DiscoveryRequest) error {
	req.Node = a.node()
	return a.stream.Send(req)
}

// watch will start watching resources, starting with LDS. Based on the LDS response
// it will start watching RDS and CDS.
func (a *Client) watch() error {
	err := a.Send(&envoy_api_v2.DiscoveryRequest{
		ResponseNonce: time.Now().String(),
		TypeUrl:       ClusterType,
	})
	if err != nil {
		return fmt.Errorf("sending request: %w", err)
	}
	return nil
}

func (a *Client) SendRsc(typeurl string, rsc []string) error {
	return a.Send(&envoy_api_v2.DiscoveryRequest{
		ResponseNonce: "",
		TypeUrl:       typeurl,
		ResourceNames: rsc,
	})
}

func (a *Client) ack(msg *envoy_api_v2.DiscoveryResponse) error {
	return a.Send(&envoy_api_v2.DiscoveryRequest{
		ResponseNonce: msg.Nonce,
		TypeUrl:       msg.TypeUrl,
		VersionInfo:   msg.VersionInfo,
	})
}

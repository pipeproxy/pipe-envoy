package ads

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"

	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_service_discovery_v2 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/golang/protobuf/proto"
	"github.com/wzshiming/envoy/internal/node"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Config for the Client connection.
type Config struct {
	NodeConfig *node.Config

	ContextDialer func(context.Context, string) (net.Conn, error)

	HandleCDS func([]*envoy_api_v2.Cluster)
	HandleRDS func([]*envoy_api_v2.RouteConfiguration)
	HandleLDS func([]*envoy_api_v2.Listener)
	HandleEDS func([]*envoy_api_v2.ClusterLoadAssignment)
}

// Client implements a client for ADS.
type Client struct {
	stream        envoy_service_discovery_v2.AggregatedDiscoveryService_StreamAggregatedResourcesClient
	conn          *grpc.ClientConn
	nodeConfig    *node.Config
	contextDialer func(context.Context, string) (net.Conn, error)

	certDir string
	url     string

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

// NewClient connects to a Client server, with optional TLS authentication if a cert dir is specified.
func NewClient(url string, certDir string, opts *Config) (*Client, error) {
	ads := &Client{
		certDir: certDir,
		url:     url,
	}
	if opts == nil {
		opts = &Config{}
	}

	ads.nodeConfig = opts.NodeConfig
	ads.contextDialer = opts.ContextDialer
	ads.HandleCDS = opts.HandleCDS
	ads.HandleRDS = opts.HandleRDS
	ads.HandleLDS = opts.HandleLDS
	ads.HandleEDS = opts.HandleEDS

	return ads, nil
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
func (c *Client) Close() error {
	if c.stream != nil {
		c.stream.CloseSend()
	}
	return c.conn.Close()
}

// Run the ADS client.
func (c *Client) Run() error {
	err := c.run()
	if err != nil {
		return err
	}
	return c.handleRecv()
}

func (c *Client) Start() error {
	err := c.run()
	if err != nil {
		return err
	}
	go c.handleRecv()
	return nil
}

func (c *Client) run() error {
	opts := []grpc.DialOption{}
	if len(c.certDir) != 0 {
		tlsCfg, err := tlsConfig(c.certDir)
		if err != nil {
			return err
		}
		creds := credentials.NewTLS(tlsCfg)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if c.contextDialer != nil {
		opts = append(opts, grpc.WithContextDialer(c.contextDialer))
	}
	conn, err := grpc.Dial(c.url, opts...)
	if err != nil {
		return err
	}
	c.conn = conn
	xds := envoy_service_discovery_v2.NewAggregatedDiscoveryServiceClient(c.conn)
	edsstr, err := xds.StreamAggregatedResources(context.Background())
	if err != nil {
		return err
	}
	c.stream = edsstr
	return nil
}

func (c *Client) handleRecv() error {
	for {
		msg, err := c.stream.Recv()
		if err != nil {
			return fmt.Errorf("connection closed : error: %w", err)
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

		c.ack(msg)

		if len(cds) != 0 {
			c.handleCDS(cds)
		}
		if len(eds) != 0 {
			c.handleEDS(eds)
		}
		if len(lds) != 0 {
			c.handleLDS(lds)
		}
		if len(rds) != 0 {
			c.handleRDS(rds)
		}
	}
}

func (c *Client) handleCDS(cds []*envoy_api_v2.Cluster) {
	if c.HandleCDS != nil {
		c.HandleCDS(cds)
	}
}

func (c *Client) handleEDS(eds []*envoy_api_v2.ClusterLoadAssignment) {
	if c.HandleEDS != nil {
		c.HandleEDS(eds)
	}
}

func (c *Client) handleLDS(lds []*envoy_api_v2.Listener) {
	if c.HandleLDS != nil {
		c.HandleLDS(lds)
	}
}

func (c *Client) handleRDS(rds []*envoy_api_v2.RouteConfiguration) {
	if c.HandleRDS != nil {
		c.HandleRDS(rds)
	}
}

func (c *Client) Node() *envoy_api_v2_core.Node {
	return c.nodeConfig.Node()
}

func (c *Client) Send(req *envoy_api_v2.DiscoveryRequest) error {
	req.Node = c.Node()
	return c.stream.Send(req)
}

func (c *Client) SendRsc(typeurl string, rsc []string) error {
	return c.Send(&envoy_api_v2.DiscoveryRequest{
		ResponseNonce: "",
		TypeUrl:       typeurl,
		ResourceNames: rsc,
	})
}

func (c *Client) ack(msg *envoy_api_v2.DiscoveryResponse) error {
	return c.Send(&envoy_api_v2.DiscoveryRequest{
		ResponseNonce: msg.Nonce,
		TypeUrl:       msg.TypeUrl,
		VersionInfo:   msg.VersionInfo,
	})
}

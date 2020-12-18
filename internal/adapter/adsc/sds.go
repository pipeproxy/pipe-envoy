package adsc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_extensions_transport_sockets_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	envoy_service_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	envoy_service_secret_v3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// SDSConfig for the SDSClient connection.
type SDSConfig struct {
	Node           *envoy_config_core_v3.Node
	OnConnect      func(cli *SDSClient) error
	ContextDialer  func(ctx context.Context, address string) (net.Conn, error)
	HandleSDS      func(cli *SDSClient, secrets []*envoy_extensions_transport_sockets_tls_v3.Secret)
	HandleNotFound func(cli *SDSClient, others []*any.Any)
}

// SDSClient implements a client for xDS.
type SDSClient struct {
	mut sync.Mutex

	stream    envoy_service_secret_v3.SecretDiscoveryService_StreamSecretsClient
	conn      *grpc.ClientConn
	tlsConfig *tls.Config
	url       string
	isClose   bool

	// Last received message, by type
	received map[string]*cache

	SDSConfig
}

// NewSDSClient connects to a xDS server, with optional TLS authentication if a cert dir is specified.
func NewSDSClient(url string, tlsConfig *tls.Config, opts *SDSConfig) *SDSClient {
	ads := &SDSClient{
		tlsConfig: tlsConfig,
		url:       url,
		received:  map[string]*cache{},
	}
	if opts != nil {
		ads.SDSConfig = *opts
	}
	return ads
}

// Clone the once.
func (c *SDSClient) Clone() *SDSClient {
	return NewSDSClient(c.url, c.tlsConfig.Clone(), &c.SDSConfig)
}

// Close the once.
func (c *SDSClient) Close() error {
	c.isClose = true
	if c.stream != nil {
		c.stream.CloseSend()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Run the xDS client.
func (c *SDSClient) Run(ctx context.Context) error {
	err := c.run(ctx)
	if err != nil {
		return err
	}
	return c.handleRecv()
}

func (c *SDSClient) Start(ctx context.Context) error {
	err := c.run(ctx)
	if err != nil {
		return err
	}
	go c.handleRecv()
	return nil
}

func (c *SDSClient) run(ctx context.Context) error {
	opts := []grpc.DialOption{}
	if c.tlsConfig != nil {
		secret := credentials.NewTLS(c.tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(secret))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if c.ContextDialer != nil {
		opts = append(opts, grpc.WithContextDialer(c.ContextDialer))
	}
	conn, err := grpc.DialContext(ctx, c.url, opts...)
	if err != nil {
		return err
	}

	xds := envoy_service_secret_v3.NewSecretDiscoveryServiceClient(conn)

	stm, err := xds.StreamSecrets(ctx)
	if err != nil {
		return err
	}

	if c.conn != nil {
		c.conn.Close()
	}
	c.conn = conn
	if c.stream != nil {
		c.stream.CloseSend()
	}
	c.stream = stm
	if c.OnConnect != nil {
		return c.OnConnect(c)
	}
	return nil
}

func (c *SDSClient) handleRecv() error {
	c.isClose = false
	secrets := []*envoy_extensions_transport_sockets_tls_v3.Secret{}
	others := []*any.Any{}
	ctx := c.stream.Context()
	for {
		err := ctx.Err()
		if err != nil {
			return err
		}
		msg, err := c.stream.Recv()
		if err != nil {
			if code := status.Code(err); code == codes.Canceled || code == codes.DeadlineExceeded {
				return nil
			}
			return fmt.Errorf("connection closed : error: %w", err)
		}

		secrets = secrets[:0]
		others = others[:0]

		for _, rsc := range msg.Resources {
			switch rsc.TypeUrl {
			case SecretType:
				ll := &envoy_extensions_transport_sockets_tls_v3.Secret{}
				_ = proto.Unmarshal(rsc.Value, ll)
				secrets = append(secrets, ll)
			default:
				others = append(others, rsc)
			}
		}

		if len(secrets) != 0 && c.HandleSDS != nil {
			c.HandleSDS(c, secrets)
		}
		if len(others) != 0 && c.HandleNotFound != nil {
			c.HandleNotFound(c, others)
		}
		c.ack(msg)
	}
}

func (c *SDSClient) Node() *envoy_config_core_v3.Node {
	return c.SDSConfig.Node
}

func (c *SDSClient) Send(req *envoy_service_discovery_v3.DiscoveryRequest) error {
	req.Node = c.Node()
	return c.stream.Send(req)
}

func (c *SDSClient) SendRsc(typeURL string, rsc []string) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.received[typeURL] == nil {
		c.received[typeURL] = &cache{}
	}
	c.received[typeURL].Names = rsc
	version := c.received[typeURL].VersionInfo
	nonce := c.received[typeURL].Nonce
	return c.Send(&envoy_service_discovery_v3.DiscoveryRequest{
		ResponseNonce: nonce,
		TypeUrl:       typeURL,
		VersionInfo:   version,
		ResourceNames: rsc,
	})
}

func (c *SDSClient) ack(msg *envoy_service_discovery_v3.DiscoveryResponse) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.received[msg.TypeUrl] == nil {
		c.received[msg.TypeUrl] = &cache{}
	}
	c.received[msg.TypeUrl].VersionInfo = msg.VersionInfo
	c.received[msg.TypeUrl].Nonce = msg.Nonce
	rsc := c.received[msg.TypeUrl].Names
	return c.Send(&envoy_service_discovery_v3.DiscoveryRequest{
		ResponseNonce: msg.Nonce,
		TypeUrl:       msg.TypeUrl,
		VersionInfo:   msg.VersionInfo,
		ResourceNames: rsc,
	})
}

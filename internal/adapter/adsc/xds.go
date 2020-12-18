package adsc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"sync"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	envoy_service_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// XDSConfig for the XDSClient connection.
type XDSConfig struct {
	Node           *envoy_config_core_v3.Node
	OnConnect      func(cli *XDSClient) error
	ContextDialer  func(ctx context.Context, address string) (net.Conn, error)
	HandleCDS      func(cli *XDSClient, clusters []*envoy_config_cluster_v3.Cluster)
	HandleEDS      func(cli *XDSClient, endpoints []*envoy_config_endpoint_v3.ClusterLoadAssignment)
	HandleLDS      func(cli *XDSClient, listeners []*envoy_config_listener_v3.Listener)
	HandleRDS      func(cli *XDSClient, routes []*envoy_config_route_v3.RouteConfiguration)
	HandleNotFound func(cli *XDSClient, others []*any.Any)
}

// XDSClient implements a client for xDS.
type XDSClient struct {
	mut sync.Mutex

	stream    envoy_service_discovery_v3.AggregatedDiscoveryService_StreamAggregatedResourcesClient
	conn      *grpc.ClientConn
	tlsConfig *tls.Config
	url       string
	isClose   bool

	// Last received message, by type
	received map[string]*cache

	XDSConfig
}

// NewXDSClient connects to a xDS server, with optional TLS authentication if a cert dir is specified.
func NewXDSClient(url string, tlsConfig *tls.Config, opts *XDSConfig) *XDSClient {
	ads := &XDSClient{
		tlsConfig: tlsConfig,
		url:       url,
		received:  map[string]*cache{},
	}
	if opts != nil {
		ads.XDSConfig = *opts
	}
	return ads
}

// Clone the once.
func (c *XDSClient) Clone() *XDSClient {
	return NewXDSClient(c.url, c.tlsConfig.Clone(), &c.XDSConfig)
}

// Close the once.
func (c *XDSClient) Close() error {
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
func (c *XDSClient) Run(ctx context.Context) error {
	err := c.run(ctx)
	if err != nil {
		return err
	}
	return c.handleRecv()
}

func (c *XDSClient) Start(ctx context.Context) error {
	err := c.run(ctx)
	if err != nil {
		return err
	}
	go c.handleRecv()
	return nil
}

func (c *XDSClient) run(ctx context.Context) error {
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

	xds := envoy_service_discovery_v3.NewAggregatedDiscoveryServiceClient(conn)

	stm, err := xds.StreamAggregatedResources(ctx)
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

func (c *XDSClient) handleRecv() error {
	c.isClose = false
	clusters := []*envoy_config_cluster_v3.Cluster{}
	endpoints := []*envoy_config_endpoint_v3.ClusterLoadAssignment{}
	listeners := []*envoy_config_listener_v3.Listener{}
	routes := []*envoy_config_route_v3.RouteConfiguration{}
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

		clusters = clusters[:0]
		endpoints = endpoints[:0]
		listeners = listeners[:0]
		routes = routes[:0]
		others = others[:0]

		for _, rsc := range msg.Resources {
			switch rsc.TypeUrl {
			case ClusterType:
				ll := &envoy_config_cluster_v3.Cluster{}
				_ = proto.Unmarshal(rsc.Value, ll)
				clusters = append(clusters, ll)
			case EndpointType:
				ll := &envoy_config_endpoint_v3.ClusterLoadAssignment{}
				_ = proto.Unmarshal(rsc.Value, ll)
				endpoints = append(endpoints, ll)
			case ListenerType:
				ll := &envoy_config_listener_v3.Listener{}
				_ = proto.Unmarshal(rsc.Value, ll)
				listeners = append(listeners, ll)
			case RouteType:
				ll := &envoy_config_route_v3.RouteConfiguration{}
				_ = proto.Unmarshal(rsc.Value, ll)
				routes = append(routes, ll)
			default:
				others = append(others, rsc)
			}
		}

		if len(clusters) != 0 && c.HandleCDS != nil {
			c.HandleCDS(c, clusters)
		}
		if len(endpoints) != 0 && c.HandleEDS != nil {
			c.HandleEDS(c, endpoints)
		}
		if len(listeners) != 0 && c.HandleLDS != nil {
			c.HandleLDS(c, listeners)
		}
		if len(routes) != 0 && c.HandleRDS != nil {
			c.HandleRDS(c, routes)
		}
		if len(others) != 0 && c.HandleNotFound != nil {
			c.HandleNotFound(c, others)
		}
		c.ack(msg)
	}
}

func (c *XDSClient) Node() *envoy_config_core_v3.Node {
	return c.XDSConfig.Node
}

func (c *XDSClient) Send(req *envoy_service_discovery_v3.DiscoveryRequest) error {
	req.Node = c.Node()
	return c.stream.Send(req)
}

func (c *XDSClient) SendRsc(typeURL string, rsc []string) error {
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

func (c *XDSClient) ack(msg *envoy_service_discovery_v3.DiscoveryResponse) error {
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

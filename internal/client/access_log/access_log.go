package access_log

import (
	"context"
	"net"

	envoy_service_accesslog_v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/wzshiming/envoy/internal/node"
	"google.golang.org/grpc"
)

// Config for the Client connection.
type Config struct {
	NodeConfig *node.Config

	ContextDialer func(context.Context, string) (net.Conn, error)
}

type Client struct {
	stream        envoy_service_accesslog_v2.AccessLogService_StreamAccessLogsClient
	conn          *grpc.ClientConn
	nodeConfig    *node.Config
	contextDialer func(context.Context, string) (net.Conn, error)

	url string
}

// NewClient connects to a Client.
func NewClient(url string, opts *Config) (*Client, error) {
	cli := &Client{
		url: url,
	}

	cli.nodeConfig = opts.NodeConfig
	cli.contextDialer = opts.ContextDialer
	return cli, nil
}

func (c *Client) Start() error {
	opts := []grpc.DialOption{}

	opts = append(opts, grpc.WithInsecure())

	if c.contextDialer != nil {
		opts = append(opts, grpc.WithContextDialer(c.contextDialer))
	}
	conn, err := grpc.Dial(c.url, opts...)
	if err != nil {
		return err
	}

	c.conn = conn

	cli := envoy_service_accesslog_v2.NewAccessLogServiceClient(conn)
	stm, err := cli.StreamAccessLogs(context.Background())
	if err != nil {
		return err
	}
	c.stream = stm
	return nil
}

func (c *Client) Send(req *envoy_service_accesslog_v2.StreamAccessLogsMessage) error {
	return c.stream.Send(req)
}

func (c *Client) SendHttpLog(name string, req *envoy_service_accesslog_v2.StreamAccessLogsMessage_HTTPAccessLogEntries) error {
	identifier := envoy_service_accesslog_v2.StreamAccessLogsMessage_Identifier{
		Node:    c.nodeConfig.Node(),
		LogName: name,
	}
	return c.stream.Send(&envoy_service_accesslog_v2.StreamAccessLogsMessage{
		Identifier: &identifier,
		LogEntries: &envoy_service_accesslog_v2.StreamAccessLogsMessage_HttpLogs{
			HttpLogs: req,
		},
	})
}

func (c *Client) SendTcpLog(name string, req *envoy_service_accesslog_v2.StreamAccessLogsMessage_TCPAccessLogEntries) error {
	identifier := envoy_service_accesslog_v2.StreamAccessLogsMessage_Identifier{
		Node:    c.nodeConfig.Node(),
		LogName: name,
	}
	return c.stream.Send(&envoy_service_accesslog_v2.StreamAccessLogsMessage{
		Identifier: &identifier,
		LogEntries: &envoy_service_accesslog_v2.StreamAccessLogsMessage_TcpLogs{
			TcpLogs: req,
		},
	})
}

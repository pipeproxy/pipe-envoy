package node

import (
	"fmt"
	"net"

	envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	structpb "github.com/golang/protobuf/ptypes/struct"
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
}

func (c *Config) Node() *envoy_api_v2_core.Node {
	if c.NodeID == "" {
		if c.Namespace == "" {
			c.Namespace = "default"
		}
		if c.NodeType == "" {
			c.NodeType = "sidecar"
		}
		if c.IP == "" {
			c.IP = getPrivateIPIfAvailable().String()
		}
		if c.Workload == "" {
			c.Workload = "test-1"
		}
		c.NodeID = fmt.Sprintf("%s~%s~%s.%s~%s.svc.cluster.local", c.NodeType, c.IP, c.Workload, c.Namespace, c.Namespace)
	}

	return &envoy_api_v2_core.Node{
		Id:       c.NodeID,
		Metadata: c.Metadate,
	}
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

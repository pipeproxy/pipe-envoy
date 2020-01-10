package cluster

import (
	envoy_api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

type Cluster struct {
	//TransportSocketMatches []*Cluster_TransportSocketMatch `protobuf:"bytes,43,rep,name=transport_socket_matches,json=transportSocketMatches,proto3" json:"transport_socket_matches,omitempty"`
	//Name                   string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//AltStatName            string                          `protobuf:"bytes,28,opt,name=alt_stat_name,json=altStatName,proto3" json:"alt_stat_name,omitempty"`
	//// Types that are valid to be assigned to ClusterDiscoveryType:
	////	*Cluster_Type
	////	*Cluster_ClusterType
	//ClusterDiscoveryType          isCluster_ClusterDiscoveryType    `protobuf_oneof:"cluster_discovery_type"`
	//EdsClusterConfig              *Cluster_EdsClusterConfig         `protobuf:"bytes,3,opt,name=eds_cluster_config,json=edsClusterConfig,proto3" json:"eds_cluster_config,omitempty"`
	//ConnectTimeout                *duration.Duration                `protobuf:"bytes,4,opt,name=connect_timeout,json=connectTimeout,proto3" json:"connect_timeout,omitempty"`
	//PerConnectionBufferLimitBytes *wrappers.UInt32Value             `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	//LbPolicy                      Cluster_LbPolicy                  `protobuf:"varint,6,opt,name=lb_policy,json=lbPolicy,proto3,enum=envoy.api.v2.Cluster_LbPolicy" json:"lb_policy,omitempty"`
	//Hosts                         []*core.Address                   `protobuf:"bytes,7,rep,name=hosts,proto3" json:"hosts,omitempty"`
	//LoadAssignment                *ClusterLoadAssignment            `protobuf:"bytes,33,opt,name=load_assignment,json=loadAssignment,proto3" json:"load_assignment,omitempty"`
	//HealthChecks                  []*core.HealthCheck               `protobuf:"bytes,8,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	//MaxRequestsPerConnection      *wrappers.UInt32Value             `protobuf:"bytes,9,opt,name=max_requests_per_connection,json=maxRequestsPerConnection,proto3" json:"max_requests_per_connection,omitempty"`
	//CircuitBreakers               *cluster.CircuitBreakers          `protobuf:"bytes,10,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	//TlsContext                    *auth.UpstreamTlsContext          `protobuf:"bytes,11,opt,name=tls_context,json=tlsContext,proto3" json:"tls_context,omitempty"` // Deprecated: Do not use.
	//UpstreamHttpProtocolOptions   *core.UpstreamHttpProtocolOptions `protobuf:"bytes,46,opt,name=upstream_http_protocol_options,json=upstreamHttpProtocolOptions,proto3" json:"upstream_http_protocol_options,omitempty"`
	//CommonHttpProtocolOptions     *core.HttpProtocolOptions         `protobuf:"bytes,29,opt,name=common_http_protocol_options,json=commonHttpProtocolOptions,proto3" json:"common_http_protocol_options,omitempty"`
	//HttpProtocolOptions           *core.Http1ProtocolOptions        `protobuf:"bytes,13,opt,name=http_protocol_options,json=httpProtocolOptions,proto3" json:"http_protocol_options,omitempty"`
	//Http2ProtocolOptions          *core.Http2ProtocolOptions        `protobuf:"bytes,14,opt,name=http2_protocol_options,json=http2ProtocolOptions,proto3" json:"http2_protocol_options,omitempty"`
	//ExtensionProtocolOptions      map[string]*_struct.Struct        `protobuf:"bytes,35,rep,name=extension_protocol_options,json=extensionProtocolOptions,proto3" json:"extension_protocol_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	//TypedExtensionProtocolOptions map[string]*any.Any               `protobuf:"bytes,36,rep,name=typed_extension_protocol_options,json=typedExtensionProtocolOptions,proto3" json:"typed_extension_protocol_options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//DnsRefreshRate                *duration.Duration                `protobuf:"bytes,16,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	//DnsFailureRefreshRate         *Cluster_RefreshRate              `protobuf:"bytes,44,opt,name=dns_failure_refresh_rate,json=dnsFailureRefreshRate,proto3" json:"dns_failure_refresh_rate,omitempty"`
	//RespectDnsTtl                 bool                              `protobuf:"varint,39,opt,name=respect_dns_ttl,json=respectDnsTtl,proto3" json:"respect_dns_ttl,omitempty"`
	//DnsLookupFamily               Cluster_DnsLookupFamily           `protobuf:"varint,17,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v2.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	//DnsResolvers                  []*core.Address                   `protobuf:"bytes,18,rep,name=dns_resolvers,json=dnsResolvers,proto3" json:"dns_resolvers,omitempty"`
	//UseTcpForDnsLookups           bool                              `protobuf:"varint,45,opt,name=use_tcp_for_dns_lookups,json=useTcpForDnsLookups,proto3" json:"use_tcp_for_dns_lookups,omitempty"`
	//OutlierDetection              *cluster.OutlierDetection         `protobuf:"bytes,19,opt,name=outlier_detection,json=outlierDetection,proto3" json:"outlier_detection,omitempty"`
	//CleanupInterval               *duration.Duration                `protobuf:"bytes,20,opt,name=cleanup_interval,json=cleanupInterval,proto3" json:"cleanup_interval,omitempty"`
	//UpstreamBindConfig            *core.BindConfig                  `protobuf:"bytes,21,opt,name=upstream_bind_config,json=upstreamBindConfig,proto3" json:"upstream_bind_config,omitempty"`
	//LbSubsetConfig                *Cluster_LbSubsetConfig           `protobuf:"bytes,22,opt,name=lb_subset_config,json=lbSubsetConfig,proto3" json:"lb_subset_config,omitempty"`
	//// Types that are valid to be assigned to LbConfig:
	////	*Cluster_RingHashLbConfig_
	////	*Cluster_OriginalDstLbConfig_
	////	*Cluster_LeastRequestLbConfig_
	//LbConfig                            isCluster_LbConfig               `protobuf_oneof:"lb_config"`
	//CommonLbConfig                      *Cluster_CommonLbConfig          `protobuf:"bytes,27,opt,name=common_lb_config,json=commonLbConfig,proto3" json:"common_lb_config,omitempty"`
	//TransportSocket                     *core.TransportSocket            `protobuf:"bytes,24,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	//Metadata                            *core.Metadata                   `protobuf:"bytes,25,opt,name=metadata,proto3" json:"metadata,omitempty"`
	//ProtocolSelection                   Cluster_ClusterProtocolSelection `protobuf:"varint,26,opt,name=protocol_selection,json=protocolSelection,proto3,enum=envoy.api.v2.Cluster_ClusterProtocolSelection" json:"protocol_selection,omitempty"`
	//UpstreamConnectionOptions           *UpstreamConnectionOptions       `protobuf:"bytes,30,opt,name=upstream_connection_options,json=upstreamConnectionOptions,proto3" json:"upstream_connection_options,omitempty"`
	//CloseConnectionsOnHostHealthFailure bool                             `protobuf:"varint,31,opt,name=close_connections_on_host_health_failure,json=closeConnectionsOnHostHealthFailure,proto3" json:"close_connections_on_host_health_failure,omitempty"`
	//DrainConnectionsOnHostRemoval       bool                             `protobuf:"varint,32,opt,name=drain_connections_on_host_removal,json=drainConnectionsOnHostRemoval,proto3" json:"drain_connections_on_host_removal,omitempty"`
	//Filters                             []*cluster.Filter                `protobuf:"bytes,40,rep,name=filters,proto3" json:"filters,omitempty"`
	//LoadBalancingPolicy                 *LoadBalancingPolicy             `protobuf:"bytes,41,opt,name=load_balancing_policy,json=loadBalancingPolicy,proto3" json:"load_balancing_policy,omitempty"`
	//LrsServer                           *core.ConfigSource               `protobuf:"bytes,42,opt,name=lrs_server,json=lrsServer,proto3" json:"lrs_server,omitempty"`
}

func NewCluster(config *envoy_api_v2.Cluster) *Cluster {

}

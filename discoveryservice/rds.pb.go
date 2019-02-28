// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rds.proto

package discoveryservice

import (
	_ "envoy/api/v2"
	core "envoy/api/v2/core"
	route "envoy/api/v2/route"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#comment:next free field: 9]
type RouteConfiguration struct {
	// The name of the route configuration. For example, it might match
	// :ref:`route_config_name
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.Rds.route_config_name>` in
	// :ref:`envoy_api_msg_config.filter.network.http_connection_manager.v2.Rds`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// An array of virtual hosts that make up the route table.
	VirtualHosts []*route.VirtualHost `protobuf:"bytes,2,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// Optionally specifies a list of HTTP headers that the connection manager
	// will consider to be internal only. If they are found on external requests they will be cleaned
	// prior to filter invocation. See :ref:`config_http_conn_man_headers_x-envoy-internal` for more
	// information.
	InternalOnlyHeaders []string `protobuf:"bytes,3,rep,name=internal_only_headers,json=internalOnlyHeaders,proto3" json:"internal_only_headers,omitempty"`
	// Specifies a list of HTTP headers that should be added to each response that
	// the connection manager encodes. Headers specified at this level are applied
	// after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	ResponseHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each response
	// that the connection manager encodes.
	ResponseHeadersToRemove []string `protobuf:"bytes,5,rep,name=response_headers_to_remove,json=responseHeadersToRemove,proto3" json:"response_headers_to_remove,omitempty"`
	// Specifies a list of HTTP headers that should be added to each request
	// routed by the HTTP connection manager. Headers specified at this level are
	// applied after headers from any enclosed :ref:`envoy_api_msg_route.VirtualHost` or
	// :ref:`envoy_api_msg_route.RouteAction`. For more information, including details on
	// header value syntax, see the documentation on :ref:`custom request headers
	// <config_http_conn_man_headers_custom_request_headers>`.
	RequestHeadersToAdd []*core.HeaderValueOption `protobuf:"bytes,6,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Specifies a list of HTTP headers that should be removed from each request
	// routed by the HTTP connection manager.
	RequestHeadersToRemove []string `protobuf:"bytes,8,rep,name=request_headers_to_remove,json=requestHeadersToRemove,proto3" json:"request_headers_to_remove,omitempty"`
	// An optional boolean that specifies whether the clusters that the route
	// table refers to will be validated by the cluster manager. If set to true
	// and a route refers to a non-existent cluster, the route table will not
	// load. If set to false and a route refers to a non-existent cluster, the
	// route table will load and the router filter will return a 404 if the route
	// is selected at runtime. This setting defaults to true if the route table
	// is statically defined via the :ref:`route_config
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.route_config>`
	// option. This setting default to false if the route table is loaded dynamically via the
	// :ref:`rds
	// <envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.rds>`
	// option. Users may which to override the default behavior in certain cases (for example when
	// using CDS with a static route table).
	ValidateClusters     *wrappers.BoolValue `protobuf:"bytes,7,opt,name=validate_clusters,json=validateClusters,proto3" json:"validate_clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_6eac9d3f8ca57de4, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetVirtualHosts() []*route.VirtualHost {
	if m != nil {
		return m.VirtualHosts
	}
	return nil
}

func (m *RouteConfiguration) GetInternalOnlyHeaders() []string {
	if m != nil {
		return m.InternalOnlyHeaders
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.ResponseHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetResponseHeadersToRemove() []string {
	if m != nil {
		return m.ResponseHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToAdd() []*core.HeaderValueOption {
	if m != nil {
		return m.RequestHeadersToAdd
	}
	return nil
}

func (m *RouteConfiguration) GetRequestHeadersToRemove() []string {
	if m != nil {
		return m.RequestHeadersToRemove
	}
	return nil
}

func (m *RouteConfiguration) GetValidateClusters() *wrappers.BoolValue {
	if m != nil {
		return m.ValidateClusters
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.api.v2.RouteConfiguration")
}

func init() { proto.RegisterFile("rds.proto", fileDescriptor_6eac9d3f8ca57de4) }

var fileDescriptor_6eac9d3f8ca57de4 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xae, 0x9b, 0xfc, 0xd2, 0xe6, 0x92, 0x9f, 0xd4, 0x5e, 0xd3, 0x26, 0x8d, 0x50, 0x12, 0x45,
	0x0c, 0xa1, 0x83, 0x8d, 0xcc, 0x44, 0x99, 0x48, 0x11, 0x14, 0x96, 0x56, 0x2e, 0x74, 0xb5, 0x2e,
	0xf6, 0x4b, 0x62, 0xc9, 0xb9, 0x67, 0xee, 0xce, 0x86, 0xac, 0x4c, 0xcc, 0xf0, 0x4f, 0xf0, 0x37,
	0xc0, 0xc2, 0xc8, 0x8a, 0xd8, 0x19, 0x10, 0x03, 0xfc, 0x17, 0xc8, 0x67, 0x1b, 0xea, 0xb4, 0x48,
	0x08, 0xb1, 0x44, 0xcf, 0xfe, 0xbe, 0xef, 0x7d, 0xdf, 0xbb, 0x7b, 0x31, 0xa9, 0x0b, 0x5f, 0x9a,
	0x91, 0x40, 0x85, 0xb4, 0x09, 0x3c, 0xc1, 0xa5, 0xc9, 0xa2, 0xc0, 0x4c, 0xec, 0xee, 0x35, 0xfd,
	0x64, 0xb1, 0x28, 0xb0, 0x12, 0xdb, 0xf2, 0x50, 0x80, 0x35, 0x61, 0x12, 0x32, 0xee, 0x0a, 0xea,
	0x07, 0xd2, 0xc3, 0x04, 0xc4, 0x32, 0x47, 0x7b, 0x25, 0x54, 0x60, 0xac, 0x20, 0xfb, 0x2d, 0xd4,
	0x33, 0xc4, 0x59, 0x08, 0x9a, 0xc0, 0x38, 0x47, 0xc5, 0x54, 0x80, 0x5c, 0x16, 0xea, 0x1c, 0xd5,
	0x4f, 0x93, 0x78, 0x6a, 0x3d, 0x13, 0x2c, 0x8a, 0x40, 0x14, 0x78, 0x3b, 0x61, 0x61, 0xe0, 0x33,
	0x05, 0x56, 0x51, 0xe4, 0x40, 0x6b, 0x86, 0x33, 0xd4, 0xa5, 0x95, 0x56, 0xd9, 0xdb, 0xe1, 0xbb,
	0x2a, 0xa1, 0x4e, 0x6a, 0x7e, 0x84, 0x7c, 0x1a, 0xcc, 0x62, 0xa1, 0xcd, 0x28, 0x25, 0x55, 0xce,
	0x16, 0xd0, 0x31, 0x06, 0xc6, 0xa8, 0xee, 0xe8, 0x9a, 0x3e, 0x22, 0xff, 0x27, 0x81, 0x50, 0x31,
	0x0b, 0xdd, 0x39, 0x4a, 0x25, 0x3b, 0xeb, 0x83, 0xca, 0xa8, 0x61, 0xf7, 0xcd, 0x8b, 0x27, 0x63,
	0x66, 0x93, 0x9c, 0x67, 0xc4, 0x63, 0x94, 0x6a, 0x5c, 0xfd, 0xf0, 0xb9, 0xbf, 0xe6, 0x34, 0x93,
	0x5f, 0xaf, 0x24, 0xb5, 0xc9, 0x6e, 0xc0, 0x15, 0x08, 0xce, 0x42, 0x17, 0x79, 0xb8, 0x74, 0xe7,
	0xc0, 0x7c, 0x10, 0xb2, 0x53, 0x19, 0x54, 0x46, 0x75, 0x67, 0xa7, 0x00, 0x4f, 0x78, 0xb8, 0x3c,
	0xce, 0x20, 0x3a, 0x27, 0x6d, 0x01, 0x32, 0x42, 0x2e, 0xa1, 0xa0, 0xbb, 0x0a, 0x5d, 0xe6, 0xfb,
	0x9d, 0xaa, 0x4e, 0x72, 0xbd, 0x9c, 0x24, 0xbd, 0x15, 0x33, 0x13, 0x9f, 0xb3, 0x30, 0x86, 0x93,
	0x28, 0x1d, 0x6d, 0xdc, 0x78, 0xfb, 0xfd, 0x7d, 0xa5, 0xf6, 0xca, 0xa8, 0x6c, 0x7d, 0xdb, 0x70,
	0x5a, 0x45, 0xc7, 0xdc, 0xe4, 0x31, 0xde, 0xf5, 0x7d, 0x7a, 0x87, 0x74, 0xaf, 0x72, 0x12, 0xb0,
	0xc0, 0x04, 0x3a, 0xff, 0xe9, 0x88, 0xed, 0x4b, 0x4a, 0x47, 0xc3, 0x74, 0x4a, 0xf6, 0x04, 0x3c,
	0x8d, 0x41, 0xaa, 0xd5, 0x94, 0xb5, 0xbf, 0x4d, 0xb9, 0x93, 0x37, 0x2c, 0x85, 0xbc, 0x4d, 0xf6,
	0xaf, 0xf0, 0xc9, 0x33, 0x6e, 0xea, 0x8c, 0x7b, 0xab, 0xba, 0x3c, 0xe2, 0x03, 0xb2, 0x5d, 0x2c,
	0x87, 0xeb, 0x85, 0xb1, 0x54, 0xe9, 0xc9, 0x6f, 0x0c, 0x8c, 0x51, 0xc3, 0xee, 0x9a, 0xd9, 0x7e,
	0x99, 0xc5, 0x7e, 0x99, 0x63, 0xc4, 0x50, 0x27, 0x73, 0xb6, 0x0a, 0xd1, 0x51, 0xae, 0xb1, 0x3f,
	0xae, 0x93, 0x5d, 0xbd, 0x3d, 0xf7, 0x8a, 0x1d, 0x3f, 0x03, 0x91, 0x04, 0x1e, 0xd0, 0x27, 0xa4,
	0x79, 0xa6, 0x04, 0xb0, 0x85, 0x86, 0x25, 0xed, 0x95, 0xa7, 0xfe, 0xc9, 0x77, 0xb2, 0x84, 0xdd,
	0xfe, 0x6f, 0xf1, 0xec, 0x94, 0x87, 0x6b, 0x23, 0xe3, 0xa6, 0x41, 0x23, 0xb2, 0xfd, 0x90, 0x7b,
	0x02, 0x16, 0xc0, 0x15, 0x0b, 0xf3, 0xde, 0x37, 0xca, 0xda, 0x0b, 0x84, 0x4b, 0x36, 0x07, 0x7f,
	0x42, 0x2d, 0x39, 0x22, 0x69, 0xdc, 0x07, 0xe5, 0xcd, 0xff, 0xd5, 0x1c, 0xfd, 0x17, 0x9f, 0xbe,
	0xbe, 0x5e, 0xdf, 0x1f, 0xb6, 0x4a, 0x9f, 0x86, 0x43, 0xfd, 0xb7, 0x91, 0x87, 0xc6, 0xc1, 0xd8,
	0x7a, 0xf3, 0xa5, 0x67, 0x90, 0x6e, 0x80, 0x59, 0xa7, 0x48, 0xe0, 0xf3, 0x65, 0xa9, 0xe9, 0x78,
	0xd3, 0xf1, 0xe5, 0x69, 0x7a, 0x3d, 0xa7, 0xc6, 0x4b, 0xc3, 0x98, 0xd4, 0xf4, 0x55, 0xdd, 0xfa,
	0x11, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x50, 0xf6, 0x72, 0xad, 0x04, 0x00, 0x00,
}
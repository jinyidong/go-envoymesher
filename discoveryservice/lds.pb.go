// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lds.proto

package discoveryservice

import (
	_ "envoy/api/v2"
	core "envoy/api/v2/core"
	listener "envoy/api/v2/listener"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Listener_DrainType int32

const (
	// Drain in response to calling /healthcheck/fail admin endpoint (along with the health check
	// filter), listener removal/modification, and hot restart.
	Listener_DEFAULT Listener_DrainType = 0
	// Drain in response to listener removal/modification and hot restart. This setting does not
	// include /healthcheck/fail. This setting may be desirable if Envoy is hosting both ingress
	// and egress listeners.
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}

var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}

func (Listener_DrainType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14fceb022f959645, []int{0, 0}
}

// [#comment:next free field: 16]
type Listener struct {
	// The unique name by which this listener is known. If no name is provided,
	// Envoy will allocate an internal UUID for the listener. If the listener is to be dynamically
	// updated or removed via :ref:`LDS <config_listeners_lds>` a unique name must be provided.
	// By default, the maximum length of a listener's name is limited to 60 characters. This limit can
	// be increased by setting the :option:`--max-obj-name-len` command line argument to the desired
	// value.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The address that the listener should listen on. In general, the address must be unique, though
	// that is governed by the bind rules of the OS. E.g., multiple listeners can listen on port 0 on
	// Linux as the actual port will be allocated by the OS.
	Address *core.Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// A list of filter chains to consider for this listener. The
	// :ref:`FilterChain <envoy_api_msg_listener.FilterChain>` with the most specific
	// :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` criteria is used on a
	// connection.
	//
	// Example using SNI for filter chain selection can be found in the
	// :ref:`FAQ entry <faq_how_to_setup_sni>`.
	FilterChains []*listener.FilterChain `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains,proto3" json:"filter_chains,omitempty"`
	// If a connection is redirected using *iptables*, the port on which the proxy
	// receives it might be different from the original destination address. When this flag is set to
	// true, the listener hands off redirected connections to the listener associated with the
	// original destination address. If there is no listener associated with the original destination
	// address, the connection is handled by the listener that receives it. Defaults to false.
	//
	// .. attention::
	//
	//   This field is deprecated. Use :ref:`an original_dst <config_listener_filters_original_dst>`
	//   :ref:`listener filter <envoy_api_field_Listener.listener_filters>` instead.
	//
	//   Note that hand off to another listener is *NOT* performed without this flag. Once
	//   :ref:`FilterChainMatch <envoy_api_msg_listener.FilterChainMatch>` is implemented this flag
	//   will be removed, as filter chain matching can be used to select a filter chain based on the
	//   restored destination address.
	UseOriginalDst *wrappers.BoolValue `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst,proto3" json:"use_original_dst,omitempty"` // Deprecated: Do not use.
	// Soft limit on size of the listener’s new connection read and write buffers.
	// If unspecified, an implementation defined default is applied (1MiB).
	PerConnectionBufferLimitBytes *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes,proto3" json:"per_connection_buffer_limit_bytes,omitempty"`
	// Listener metadata.
	Metadata *core.Metadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// [#not-implemented-hide:]
	DeprecatedV1 *Listener_DeprecatedV1 `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"`
	// The type of draining to perform at a listener-wide level.
	DrainType Listener_DrainType `protobuf:"varint,8,opt,name=drain_type,json=drainType,proto3,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
	// Listener filters have the opportunity to manipulate and augment the connection metadata that
	// is used in connection filter chain matching, for example. These filters are run before any in
	// :ref:`filter_chains <envoy_api_field_Listener.filter_chains>`. Order matters as the
	// filters are processed sequentially right after a socket has been accepted by the listener, and
	// before a connection is created.
	ListenerFilters []*listener.ListenerFilter `protobuf:"bytes,9,rep,name=listener_filters,json=listenerFilters,proto3" json:"listener_filters,omitempty"`
	// The timeout to wait for all listener filters to complete operation. If the timeout is reached,
	// the accepted socket is closed without a connection being created. Specify 0 to disable the
	// timeout. If not specified, a default timeout of 15s is used.
	ListenerFiltersTimeout *duration.Duration `protobuf:"bytes,15,opt,name=listener_filters_timeout,json=listenerFiltersTimeout,proto3" json:"listener_filters_timeout,omitempty"`
	// Whether the listener should be set as a transparent socket.
	// When this flag is set to true, connections can be redirected to the listener using an
	// *iptables* *TPROXY* target, in which case the original source and destination addresses and
	// ports are preserved on accepted connections. This flag should be used in combination with
	// :ref:`an original_dst <config_listener_filters_original_dst>` :ref:`listener filter
	// <envoy_api_field_Listener.listener_filters>` to mark the connections' local addresses as
	// "restored." This can be used to hand off each redirected connection to another listener
	// associated with the connection's destination address. Direct connections to the socket without
	// using *TPROXY* cannot be distinguished from connections redirected using *TPROXY* and are
	// therefore treated as if they were redirected.
	// When this flag is set to false, the listener's socket is explicitly reset as non-transparent.
	// Setting this flag requires Envoy to run with the *CAP_NET_ADMIN* capability.
	// When this flag is not set (default), the socket is not modified, i.e. the transparent option
	// is neither set nor reset.
	Transparent *wrappers.BoolValue `protobuf:"bytes,10,opt,name=transparent,proto3" json:"transparent,omitempty"`
	// Whether the listener should set the *IP_FREEBIND* socket option. When this
	// flag is set to true, listeners can be bound to an IP address that is not
	// configured on the system running Envoy. When this flag is set to false, the
	// option *IP_FREEBIND* is disabled on the socket. When this flag is not set
	// (default), the socket is not modified, i.e. the option is neither enabled
	// nor disabled.
	Freebind *wrappers.BoolValue `protobuf:"bytes,11,opt,name=freebind,proto3" json:"freebind,omitempty"`
	// Additional socket options that may not be present in Envoy source code or
	// precompiled binaries.
	SocketOptions []*core.SocketOption `protobuf:"bytes,13,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	// Whether the listener should accept TCP Fast Open (TFO) connections.
	// When this flag is set to a value greater than 0, the option TCP_FASTOPEN is enabled on
	// the socket, with a queue length of the specified size
	// (see `details in RFC7413 <https://tools.ietf.org/html/rfc7413#section-5.1>`_).
	// When this flag is set to 0, the option TCP_FASTOPEN is disabled on the socket.
	// When this flag is not set (default), the socket is not modified,
	// i.e. the option is neither enabled nor disabled.
	//
	// On Linux, the net.ipv4.tcp_fastopen kernel parameter must include flag 0x2 to enable
	// TCP_FASTOPEN.
	// See `ip-sysctl.txt <https://www.kernel.org/doc/Documentation/networking/ip-sysctl.txt>`_.
	//
	// On macOS, only values of 0, 1, and unset are valid; other values may result in an error.
	// To set the queue length on macOS, set the net.inet.tcp.fastopen_backlog kernel parameter.
	TcpFastOpenQueueLength *wrappers.UInt32Value `protobuf:"bytes,12,opt,name=tcp_fast_open_queue_length,json=tcpFastOpenQueueLength,proto3" json:"tcp_fast_open_queue_length,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *Listener) Reset()         { *m = Listener{} }
func (m *Listener) String() string { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()    {}
func (*Listener) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fceb022f959645, []int{0}
}

func (m *Listener) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener.Unmarshal(m, b)
}
func (m *Listener) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener.Marshal(b, m, deterministic)
}
func (m *Listener) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener.Merge(m, src)
}
func (m *Listener) XXX_Size() int {
	return xxx_messageInfo_Listener.Size(m)
}
func (m *Listener) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener.DiscardUnknown(m)
}

var xxx_messageInfo_Listener proto.InternalMessageInfo

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*listener.FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

// Deprecated: Do not use.
func (m *Listener) GetUseOriginalDst() *wrappers.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

func (m *Listener) GetListenerFilters() []*listener.ListenerFilter {
	if m != nil {
		return m.ListenerFilters
	}
	return nil
}

func (m *Listener) GetListenerFiltersTimeout() *duration.Duration {
	if m != nil {
		return m.ListenerFiltersTimeout
	}
	return nil
}

func (m *Listener) GetTransparent() *wrappers.BoolValue {
	if m != nil {
		return m.Transparent
	}
	return nil
}

func (m *Listener) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *Listener) GetSocketOptions() []*core.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

func (m *Listener) GetTcpFastOpenQueueLength() *wrappers.UInt32Value {
	if m != nil {
		return m.TcpFastOpenQueueLength
	}
	return nil
}

// [#not-implemented-hide:]
type Listener_DeprecatedV1 struct {
	// Whether the listener should bind to the port. A listener that doesn't
	// bind can only receive connections redirected from other listeners that
	// set use_original_dst parameter to true. Default is true.
	//
	// [V2-API-DIFF] This is deprecated in v2, all Listeners will bind to their
	// port. An additional filter chain must be created for every original
	// destination port this listener may redirect to in v2, with the original
	// port specified in the FilterChainMatch destination_port field.
	//
	// [#comment:TODO(PiotrSikora): Remove this once verified that we no longer need it.]
	BindToPort           *wrappers.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort,proto3" json:"bind_to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Listener_DeprecatedV1) Reset()         { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()    {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_14fceb022f959645, []int{0, 0}
}

func (m *Listener_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listener_DeprecatedV1.Unmarshal(m, b)
}
func (m *Listener_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listener_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *Listener_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listener_DeprecatedV1.Merge(m, src)
}
func (m *Listener_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_Listener_DeprecatedV1.Size(m)
}
func (m *Listener_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_Listener_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_Listener_DeprecatedV1 proto.InternalMessageInfo

func (m *Listener_DeprecatedV1) GetBindToPort() *wrappers.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
}

func init() { proto.RegisterFile("lds.proto", fileDescriptor_14fceb022f959645) }

var fileDescriptor_14fceb022f959645 = []byte{
	// 842 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x41, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x33, 0x8e, 0xdb, 0xd8, 0x63, 0xc7, 0xb6, 0x46, 0xa8, 0x5d, 0x4c, 0x88, 0x8d, 0x01,
	0xc9, 0x70, 0x58, 0x53, 0x57, 0x02, 0xa9, 0xaa, 0x84, 0xba, 0x35, 0x56, 0x8b, 0x5c, 0x1c, 0x36,
	0x69, 0x68, 0x4e, 0xab, 0xf1, 0xee, 0xb3, 0x33, 0x62, 0x3d, 0x33, 0x9d, 0x99, 0x35, 0xf8, 0xca,
	0x09, 0x71, 0x84, 0x0b, 0x1f, 0x81, 0xcf, 0xc0, 0x89, 0x23, 0x77, 0xee, 0x20, 0x21, 0x2e, 0x88,
	0xcf, 0x80, 0x84, 0x76, 0xbc, 0x6b, 0xec, 0x26, 0x21, 0x1c, 0x7a, 0x7b, 0xe3, 0xf7, 0xff, 0xff,
	0x32, 0xfb, 0x7f, 0x6f, 0x37, 0xb8, 0x1c, 0x47, 0xda, 0x95, 0x4a, 0x18, 0x41, 0xaa, 0xc0, 0x17,
	0x62, 0xe9, 0x52, 0xc9, 0xdc, 0x45, 0xbf, 0xd9, 0xb2, 0xa7, 0x1e, 0x95, 0xac, 0xb7, 0xe8, 0xf7,
	0x42, 0xa1, 0xa0, 0x47, 0xa3, 0x48, 0x81, 0xce, 0xe4, 0xcd, 0x83, 0x8b, 0x82, 0x09, 0xd5, 0x70,
	0x69, 0x37, 0x62, 0x3a, 0x14, 0x0b, 0x50, 0xcb, 0xac, 0xfb, 0xd6, 0x56, 0x37, 0x66, 0xda, 0x00,
	0x07, 0xb5, 0x2e, 0x72, 0xc6, 0x4c, 0x88, 0x59, 0x0c, 0x56, 0x46, 0x39, 0x17, 0x86, 0x1a, 0x26,
	0x78, 0xfe, 0xf7, 0x0f, 0xb3, 0xae, 0x3d, 0x4d, 0x92, 0x69, 0x2f, 0x4a, 0x94, 0x15, 0x5c, 0xd5,
	0xff, 0x42, 0x51, 0x29, 0x41, 0xe5, 0xfe, 0xdb, 0x0b, 0x1a, 0xb3, 0x88, 0x1a, 0xe8, 0xe5, 0x45,
	0xd6, 0x78, 0x65, 0x26, 0x66, 0xc2, 0x96, 0xbd, 0xb4, 0x5a, 0xfd, 0xda, 0xf9, 0xbb, 0x84, 0x4b,
	0xa3, 0xec, 0x7e, 0x84, 0xe0, 0x22, 0xa7, 0x73, 0x70, 0x50, 0x1b, 0x75, 0xcb, 0xbe, 0xad, 0xc9,
	0x00, 0xef, 0x65, 0x01, 0x39, 0x85, 0x36, 0xea, 0x56, 0xfa, 0x4d, 0x77, 0x33, 0x50, 0x37, 0x4d,
	0xc8, 0x7d, 0xb0, 0x52, 0x78, 0xb5, 0x9f, 0x7f, 0x6d, 0xed, 0xfc, 0xf8, 0xe7, 0x4f, 0xbb, 0x37,
	0xbe, 0x41, 0x85, 0x06, 0xf2, 0x73, 0x2b, 0xf9, 0x0c, 0xef, 0x4f, 0x59, 0x6c, 0x40, 0x05, 0xe1,
	0x39, 0x65, 0x5c, 0x3b, 0xbb, 0xed, 0xdd, 0x6e, 0xa5, 0xdf, 0xd9, 0x66, 0xad, 0x83, 0x1a, 0x5a,
	0xed, 0xc3, 0x54, 0xba, 0xc1, 0xfc, 0x16, 0x15, 0x4a, 0xc8, 0xaf, 0x4e, 0xff, 0x6d, 0x6a, 0xf2,
	0x08, 0x37, 0x12, 0x0d, 0x81, 0x50, 0x6c, 0xc6, 0x38, 0x8d, 0x83, 0x48, 0x1b, 0xa7, 0x98, 0xdd,
	0x73, 0x95, 0x94, 0x9b, 0x27, 0xe5, 0x7a, 0x42, 0xc4, 0xa7, 0x34, 0x4e, 0xc0, 0x2b, 0x38, 0xc8,
	0xaf, 0x25, 0x1a, 0xc6, 0x99, 0x6d, 0xa0, 0x0d, 0x99, 0xe2, 0x37, 0x64, 0x7a, 0x3f, 0xc1, 0x39,
	0x84, 0x69, 0xe0, 0xc1, 0x24, 0x99, 0x4e, 0x41, 0x05, 0x31, 0x9b, 0x33, 0x13, 0x4c, 0x96, 0x06,
	0xb4, 0x73, 0xc3, 0xa2, 0x0f, 0x2e, 0xa0, 0x9f, 0x3e, 0xe6, 0xe6, 0x6e, 0xdf, 0xc2, 0xfd, 0xd7,
	0x25, 0xa8, 0x87, 0x6b, 0x8a, 0x67, 0x21, 0xa3, 0x94, 0xe1, 0xa5, 0x08, 0xf2, 0x01, 0x2e, 0xcd,
	0xc1, 0xd0, 0x88, 0x1a, 0xea, 0xdc, 0xb4, 0xb8, 0xd7, 0x2e, 0x49, 0xf4, 0x49, 0x26, 0xf1, 0xd7,
	0x62, 0xf2, 0x08, 0xef, 0x47, 0x20, 0x15, 0x84, 0xd4, 0x40, 0x14, 0x2c, 0xee, 0x38, 0x7b, 0xd6,
	0xfd, 0xe6, 0xb6, 0x3b, 0x1f, 0xa6, 0x3b, 0x58, 0x6b, 0x4f, 0xef, 0xf8, 0xd5, 0x68, 0xe3, 0x44,
	0x3e, 0xc4, 0x38, 0x52, 0x94, 0xf1, 0xc0, 0x2c, 0x25, 0x38, 0xa5, 0x36, 0xea, 0xd6, 0xfa, 0xed,
	0xab, 0x30, 0xa9, 0xf0, 0x64, 0x29, 0xc1, 0x2f, 0x47, 0x79, 0x49, 0x4e, 0x71, 0x23, 0x9f, 0x55,
	0xb0, 0x1a, 0x87, 0x76, 0xca, 0x76, 0xa2, 0x6f, 0x5f, 0x31, 0xd1, 0x9c, 0xb7, 0x9a, 0xac, 0x57,
	0x4c, 0x87, 0xea, 0xd7, 0xe3, 0xad, 0x5f, 0x35, 0x39, 0xc3, 0xce, 0x8b, 0xdc, 0xc0, 0xb0, 0x39,
	0x88, 0xc4, 0x38, 0x75, 0xfb, 0xb4, 0xaf, 0x5e, 0x88, 0x7e, 0x90, 0xbd, 0x1f, 0x5e, 0xf1, 0xfb,
	0xdf, 0x5a, 0xc8, 0xbf, 0xf5, 0x02, 0xf3, 0x64, 0x65, 0x27, 0xf7, 0x71, 0xc5, 0x28, 0xca, 0xb5,
	0xa4, 0x0a, 0xb8, 0x71, 0xf0, 0x75, 0x3b, 0xe2, 0x6f, 0xca, 0xc9, 0xfb, 0xb8, 0x34, 0x55, 0x00,
	0x13, 0xc6, 0x23, 0xa7, 0x72, 0xad, 0x75, 0xad, 0x25, 0x43, 0x5c, 0xd3, 0x22, 0xfc, 0x1c, 0x4c,
	0x20, 0xa4, 0x7d, 0xcb, 0x9d, 0x7d, 0x1b, 0x53, 0xeb, 0x92, 0x91, 0x1f, 0x5b, 0xe1, 0xd8, 0xea,
	0xfc, 0x7d, 0xbd, 0x71, 0xd2, 0xe4, 0x19, 0x6e, 0x9a, 0x50, 0x06, 0x53, 0xaa, 0x53, 0x12, 0xf0,
	0xe0, 0x79, 0x02, 0x09, 0x04, 0x31, 0xf0, 0x99, 0x39, 0x77, 0xaa, 0xff, 0x63, 0x2b, 0x6f, 0x99,
	0x50, 0x0e, 0xa9, 0x36, 0x63, 0x09, 0xfc, 0xd3, 0xd4, 0x3c, 0xb2, 0xde, 0xe6, 0x08, 0x57, 0x37,
	0x37, 0x85, 0xdc, 0xc7, 0xd5, 0xf4, 0xe6, 0x81, 0x11, 0x81, 0x14, 0xca, 0xd8, 0x6f, 0xc1, 0x7f,
	0x3f, 0x2d, 0x4e, 0xf5, 0x27, 0xe2, 0x48, 0x28, 0xd3, 0x79, 0x07, 0x97, 0xd7, 0x0b, 0x43, 0x2a,
	0x78, 0x6f, 0xf0, 0xd1, 0xf0, 0xc1, 0xd3, 0xd1, 0x49, 0x63, 0x87, 0xd4, 0x71, 0xe5, 0xc9, 0x78,
	0xf0, 0x78, 0x78, 0x16, 0x8c, 0x3f, 0x19, 0x9d, 0x35, 0xd0, 0xc7, 0xc5, 0x52, 0xad, 0x51, 0xef,
	0xff, 0x85, 0xb0, 0x93, 0xef, 0xc6, 0x20, 0xff, 0x9c, 0x1e, 0x83, 0x5a, 0xb0, 0x10, 0xc8, 0x33,
	0x5c, 0x3f, 0x36, 0x0a, 0xe8, 0x3c, 0x57, 0x68, 0x72, 0xb8, 0x1d, 0xdc, 0xda, 0xe2, 0xc3, 0xf3,
	0x04, 0xb4, 0x69, 0xb6, 0xae, 0xec, 0x6b, 0x29, 0xb8, 0x86, 0xce, 0x4e, 0x17, 0xbd, 0x87, 0x48,
	0x82, 0x6b, 0x43, 0x30, 0xe1, 0xf9, 0x4b, 0x04, 0x77, 0xbe, 0xfa, 0xe5, 0x8f, 0xef, 0x0a, 0x07,
	0x9d, 0xdb, 0x5b, 0xff, 0x19, 0xee, 0xe5, 0xeb, 0xa8, 0xef, 0xa1, 0x77, 0xbd, 0xde, 0x0f, 0xbf,
	0x1f, 0x22, 0xdc, 0x64, 0x62, 0x05, 0x93, 0x4a, 0x7c, 0xb9, 0xdc, 0xe2, 0x7a, 0xa5, 0x51, 0xa4,
	0x8f, 0xd2, 0x8c, 0x8f, 0xd0, 0xd7, 0x08, 0x4d, 0x6e, 0xda, 0xbc, 0xef, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x71, 0x62, 0x00, 0xd0, 0x06, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/listener/v3/quic_config.proto

package envoy_config_listener_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type QuicProtocolOptions struct {
	MaxConcurrentStreams   *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_concurrent_streams,json=maxConcurrentStreams,proto3" json:"max_concurrent_streams,omitempty"`
	IdleTimeout            *duration.Duration    `protobuf:"bytes,2,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	CryptoHandshakeTimeout *duration.Duration    `protobuf:"bytes,3,opt,name=crypto_handshake_timeout,json=cryptoHandshakeTimeout,proto3" json:"crypto_handshake_timeout,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *QuicProtocolOptions) Reset()         { *m = QuicProtocolOptions{} }
func (m *QuicProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*QuicProtocolOptions) ProtoMessage()    {}
func (*QuicProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_a504aa24c2bbc7cf, []int{0}
}

func (m *QuicProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuicProtocolOptions.Unmarshal(m, b)
}
func (m *QuicProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuicProtocolOptions.Marshal(b, m, deterministic)
}
func (m *QuicProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuicProtocolOptions.Merge(m, src)
}
func (m *QuicProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_QuicProtocolOptions.Size(m)
}
func (m *QuicProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QuicProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QuicProtocolOptions proto.InternalMessageInfo

func (m *QuicProtocolOptions) GetMaxConcurrentStreams() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConcurrentStreams
	}
	return nil
}

func (m *QuicProtocolOptions) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *QuicProtocolOptions) GetCryptoHandshakeTimeout() *duration.Duration {
	if m != nil {
		return m.CryptoHandshakeTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*QuicProtocolOptions)(nil), "envoy.config.listener.v3.QuicProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/listener/v3/quic_config.proto", fileDescriptor_a504aa24c2bbc7cf)
}

var fileDescriptor_a504aa24c2bbc7cf = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4d, 0x4e, 0xf3, 0x30,
	0x10, 0x86, 0xd5, 0x7e, 0xd2, 0xb7, 0x70, 0x91, 0x90, 0x02, 0xaa, 0x02, 0x42, 0x15, 0xb0, 0x40,
	0xfc, 0x48, 0x36, 0x6a, 0x76, 0xc0, 0xaa, 0x65, 0x01, 0x2b, 0x4a, 0x0b, 0x6c, 0x23, 0xd7, 0x71,
	0x5b, 0x8b, 0xc4, 0x63, 0xfc, 0x13, 0xda, 0x1b, 0x70, 0x06, 0x0e, 0xc1, 0x2d, 0xb8, 0x17, 0x8a,
	0x9d, 0xd2, 0x45, 0x41, 0x2c, 0x33, 0xf3, 0x3c, 0xef, 0x4c, 0x3c, 0xe8, 0x94, 0xcb, 0x12, 0x16,
	0x84, 0x81, 0x9c, 0x88, 0x29, 0xc9, 0x85, 0xb1, 0x5c, 0x72, 0x4d, 0xca, 0x84, 0xbc, 0x38, 0xc1,
	0xd2, 0x50, 0xc7, 0x4a, 0x83, 0x85, 0x28, 0xf6, 0x2c, 0xae, 0x6b, 0x4b, 0x16, 0x97, 0xc9, 0x6e,
	0x67, 0x0a, 0x30, 0xcd, 0x39, 0xf1, 0xdc, 0xd8, 0x4d, 0x48, 0xe6, 0x34, 0xb5, 0x02, 0x64, 0x30,
	0xd7, 0xfb, 0xaf, 0x9a, 0x2a, 0xc5, 0xb5, 0xa9, 0xfb, 0x07, 0x2e, 0x53, 0x94, 0x50, 0x29, 0xc1,
	0x7a, 0xcd, 0x90, 0x92, 0x6b, 0x23, 0x40, 0x0a, 0x59, 0x0f, 0x3f, 0xfc, 0x68, 0xa2, 0xad, 0x7b,
	0x27, 0xd8, 0xa0, 0xfa, 0x62, 0x90, 0xdf, 0x29, 0x0f, 0x46, 0x43, 0xd4, 0x2e, 0xe8, 0xbc, 0x5a,
	0x94, 0x39, 0xad, 0xb9, 0xb4, 0xa9, 0xb1, 0x9a, 0xd3, 0xc2, 0xc4, 0x8d, 0xfd, 0xc6, 0x71, 0xab,
	0xbb, 0x87, 0xc3, 0x6c, 0xbc, 0x9c, 0x8d, 0x1f, 0x6f, 0xa5, 0x4d, 0xba, 0x4f, 0x34, 0x77, 0x7c,
	0xb8, 0x5d, 0xd0, 0x79, 0xff, 0x5b, 0x1d, 0x05, 0x33, 0xba, 0x42, 0x1b, 0x22, 0xcb, 0x79, 0x6a,
	0x45, 0xc1, 0xc1, 0xd9, 0xb8, 0xe9, 0x93, 0x76, 0xd6, 0x92, 0xae, 0xeb, 0xbf, 0x1c, 0xb6, 0x2a,
	0xfc, 0x21, 0xd0, 0xd1, 0x08, 0xc5, 0x4c, 0x2f, 0x94, 0x85, 0x74, 0x46, 0x65, 0x66, 0x66, 0xf4,
	0x79, 0x95, 0xf4, 0xef, 0xaf, 0xa4, 0x76, 0x50, 0x6f, 0x96, 0x66, 0x1d, 0x7a, 0x71, 0xfe, 0xfe,
	0xf9, 0xd6, 0x39, 0x43, 0x27, 0xe1, 0x04, 0x54, 0x09, 0x5c, 0x76, 0x57, 0x27, 0xf8, 0xe1, 0x61,
	0x7a, 0x97, 0xe8, 0x48, 0x00, 0xf6, 0xbc, 0xd2, 0x30, 0x5f, 0xe0, 0xdf, 0xae, 0xd7, 0xdb, 0xac,
	0xf4, 0xbe, 0xaf, 0xfb, 0x90, 0x41, 0x63, 0xfc, 0xdf, 0x6f, 0x96, 0x7c, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x22, 0xcd, 0x1d, 0x57, 0x1f, 0x02, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/address.proto

package envoy_api_v2_core

import (
	fmt "fmt"

	proto "github.com/bwhour/go-grpc/lib/protobuf/proto"

	math "math"

	wrappers "github.com/bwhour/go-grpc/lib/protobuf/ptypes/wrappers"

	base "github.com/bwhour/go-grpc/lib/grpc/xds/internal/proto/envoy/api/v2/core/base"

	_ "github.com/bwhour/go-grpc/lib/grpc/xds/internal/proto/validate"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}
func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{1, 0}
}

type Pipe struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pipe) Reset()         { *m = Pipe{} }
func (m *Pipe) String() string { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()    {}
func (*Pipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{0}
}
func (m *Pipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pipe.Unmarshal(m, b)
}
func (m *Pipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pipe.Marshal(b, m, deterministic)
}
func (dst *Pipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pipe.Merge(dst, src)
}
func (m *Pipe) XXX_Size() int {
	return xxx_messageInfo_Pipe.Size(m)
}
func (m *Pipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Pipe.DiscardUnknown(m)
}

var xxx_messageInfo_Pipe proto.InternalMessageInfo

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.api.v2.core.SocketAddress_Protocol" json:"protocol,omitempty"`
	Address  string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier        isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	ResolverName         string                        `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	Ipv4Compat           bool                          `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SocketAddress) Reset()         { *m = SocketAddress{} }
func (m *SocketAddress) String() string { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()    {}
func (*SocketAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{1}
}
func (m *SocketAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketAddress.Unmarshal(m, b)
}
func (m *SocketAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketAddress.Marshal(b, m, deterministic)
}
func (dst *SocketAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketAddress.Merge(dst, src)
}
func (m *SocketAddress) XXX_Size() int {
	return xxx_messageInfo_SocketAddress.Size(m)
}
func (m *SocketAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketAddress.DiscardUnknown(m)
}

var xxx_messageInfo_SocketAddress proto.InternalMessageInfo

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}

func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SocketAddress_OneofMarshaler, _SocketAddress_OneofUnmarshaler, _SocketAddress_OneofSizer, []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

func _SocketAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NamedPort)
	case nil:
	default:
		return fmt.Errorf("SocketAddress.PortSpecifier has unexpected type %T", x)
	}
	return nil
}

func _SocketAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SocketAddress)
	switch tag {
	case 3: // port_specifier.port_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PortSpecifier = &SocketAddress_PortValue{uint32(x)}
		return true, err
	case 4: // port_specifier.named_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortSpecifier = &SocketAddress_NamedPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _SocketAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.NamedPort)))
		n += len(x.NamedPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type TcpKeepalive struct {
	KeepaliveProbes      *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=keepalive_probes,json=keepaliveProbes,proto3" json:"keepalive_probes,omitempty"`
	KeepaliveTime        *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=keepalive_time,json=keepaliveTime,proto3" json:"keepalive_time,omitempty"`
	KeepaliveInterval    *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=keepalive_interval,json=keepaliveInterval,proto3" json:"keepalive_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TcpKeepalive) Reset()         { *m = TcpKeepalive{} }
func (m *TcpKeepalive) String() string { return proto.CompactTextString(m) }
func (*TcpKeepalive) ProtoMessage()    {}
func (*TcpKeepalive) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{2}
}
func (m *TcpKeepalive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpKeepalive.Unmarshal(m, b)
}
func (m *TcpKeepalive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpKeepalive.Marshal(b, m, deterministic)
}
func (dst *TcpKeepalive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpKeepalive.Merge(dst, src)
}
func (m *TcpKeepalive) XXX_Size() int {
	return xxx_messageInfo_TcpKeepalive.Size(m)
}
func (m *TcpKeepalive) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpKeepalive.DiscardUnknown(m)
}

var xxx_messageInfo_TcpKeepalive proto.InternalMessageInfo

func (m *TcpKeepalive) GetKeepaliveProbes() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveProbes
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveTime() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveTime
	}
	return nil
}

func (m *TcpKeepalive) GetKeepaliveInterval() *wrappers.UInt32Value {
	if m != nil {
		return m.KeepaliveInterval
	}
	return nil
}

type BindConfig struct {
	SourceAddress        *SocketAddress       `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Freebind             *wrappers.BoolValue  `protobuf:"bytes,2,opt,name=freebind,proto3" json:"freebind,omitempty"`
	SocketOptions        []*base.SocketOption `protobuf:"bytes,3,rep,name=socket_options,json=socketOptions,proto3" json:"socket_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BindConfig) Reset()         { *m = BindConfig{} }
func (m *BindConfig) String() string { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()    {}
func (*BindConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{3}
}
func (m *BindConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindConfig.Unmarshal(m, b)
}
func (m *BindConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindConfig.Marshal(b, m, deterministic)
}
func (dst *BindConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindConfig.Merge(dst, src)
}
func (m *BindConfig) XXX_Size() int {
	return xxx_messageInfo_BindConfig.Size(m)
}
func (m *BindConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BindConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BindConfig proto.InternalMessageInfo

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *BindConfig) GetFreebind() *wrappers.BoolValue {
	if m != nil {
		return m.Freebind
	}
	return nil
}

func (m *BindConfig) GetSocketOptions() []*base.SocketOption {
	if m != nil {
		return m.SocketOptions
	}
	return nil
}

type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address              isAddress_Address `protobuf_oneof:"address"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{4}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}

func (*Address_Pipe) isAddress_Address() {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SocketAddress); err != nil {
			return err
		}
	case *Address_Pipe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipe); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Address.Address has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address.socket_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketAddress)
		err := b.DecodeMessage(msg)
		m.Address = &Address_SocketAddress{msg}
		return true, err
	case 2: // address.pipe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipe)
		err := b.DecodeMessage(msg)
		m.Address = &Address_Pipe{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		s := proto.Size(x.SocketAddress)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Address_Pipe:
		s := proto.Size(x.Pipe)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CidrRange struct {
	AddressPrefix        string                `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"`
	PrefixLen            *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CidrRange) Reset()         { *m = CidrRange{} }
func (m *CidrRange) String() string { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()    {}
func (*CidrRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_b91d58d2da3489da, []int{5}
}
func (m *CidrRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrRange.Unmarshal(m, b)
}
func (m *CidrRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrRange.Marshal(b, m, deterministic)
}
func (dst *CidrRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrRange.Merge(dst, src)
}
func (m *CidrRange) XXX_Size() int {
	return xxx_messageInfo_CidrRange.Size(m)
}
func (m *CidrRange) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrRange.DiscardUnknown(m)
}

var xxx_messageInfo_CidrRange proto.InternalMessageInfo

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.core.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.core.SocketAddress")
	proto.RegisterType((*TcpKeepalive)(nil), "envoy.api.v2.core.TcpKeepalive")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.core.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.core.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.core.CidrRange")
	proto.RegisterEnum("envoy.api.v2.core.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/address.proto", fileDescriptor_address_b91d58d2da3489da)
}

var fileDescriptor_address_b91d58d2da3489da = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x4f, 0xdb, 0x48,
	0x14, 0xcf, 0xc4, 0x01, 0x92, 0x17, 0x92, 0x0d, 0x73, 0xc1, 0x8a, 0xd8, 0x4d, 0x14, 0xb4, 0x52,
	0x16, 0xed, 0x3a, 0xbb, 0x61, 0xb5, 0x77, 0x9c, 0x55, 0x01, 0x51, 0xb5, 0xae, 0x81, 0x5e, 0xad,
	0x49, 0xf2, 0x92, 0x8e, 0x70, 0x3c, 0xa3, 0xb1, 0x71, 0xe1, 0x56, 0xf5, 0xd0, 0x43, 0xef, 0xfd,
	0x2e, 0x55, 0x4f, 0x7c, 0x87, 0x7e, 0x82, 0x1e, 0xf9, 0x14, 0x54, 0x33, 0xb6, 0x83, 0xda, 0xb4,
	0xa2, 0xbd, 0xcd, 0xbc, 0xf7, 0xfb, 0xfd, 0xe6, 0xf7, 0xfe, 0x0c, 0x74, 0x30, 0x4a, 0xc5, 0xf5,
	0x80, 0x49, 0x3e, 0x48, 0x87, 0x83, 0x89, 0x50, 0x38, 0x60, 0xd3, 0xa9, 0xc2, 0x38, 0x76, 0xa4,
	0x12, 0x89, 0xa0, 0x5b, 0x06, 0xe0, 0x30, 0xc9, 0x9d, 0x74, 0xe8, 0x68, 0x40, 0x7b, 0x67, 0x95,
	0x33, 0x66, 0x31, 0x66, 0x84, 0xf6, 0x6f, 0x73, 0x21, 0xe6, 0x21, 0x0e, 0xcc, 0x6d, 0x7c, 0x39,
	0x1b, 0xbc, 0x54, 0x4c, 0x4a, 0x54, 0xb9, 0x60, 0x7b, 0x3b, 0x65, 0x21, 0x9f, 0xb2, 0x04, 0x07,
	0xc5, 0x21, 0x4b, 0xf4, 0x7e, 0x87, 0x8a, 0xc7, 0x25, 0xd2, 0x5f, 0xa1, 0x22, 0x59, 0xf2, 0xc2,
	0x26, 0x5d, 0xd2, 0xaf, 0xb9, 0xb5, 0x0f, 0xb7, 0x37, 0x56, 0x45, 0x95, 0xbb, 0xc4, 0x37, 0xe1,
	0xde, 0xc7, 0x32, 0x34, 0x4e, 0xc5, 0xe4, 0x02, 0x93, 0x83, 0xcc, 0x28, 0x7d, 0x06, 0x55, 0xa3,
	0x30, 0x11, 0xa1, 0x21, 0x35, 0x87, 0x7f, 0x38, 0x2b, 0xae, 0x9d, 0x2f, 0x38, 0x8e, 0x97, 0x13,
	0x5c, 0xd0, 0xfa, 0x6b, 0xaf, 0x49, 0xb9, 0x45, 0xfc, 0xa5, 0x0c, 0xdd, 0x85, 0x8d, 0xbc, 0x0d,
	0x76, 0xf9, 0x6b, 0x1b, 0x45, 0x86, 0xfe, 0x09, 0x20, 0x85, 0x4a, 0x82, 0x94, 0x85, 0x97, 0x68,
	0x5b, 0x5d, 0xd2, 0x6f, 0xb8, 0x75, 0x8d, 0x5b, 0xdf, 0xab, 0xd8, 0x77, 0x77, 0xd6, 0x51, 0xc9,
	0xaf, 0x69, 0xc0, 0x73, 0x9d, 0xa7, 0x1d, 0x80, 0x88, 0x2d, 0x70, 0x1a, 0xe8, 0x90, 0x5d, 0xd1,
	0xaa, 0x1a, 0x60, 0x62, 0x9e, 0x50, 0x09, 0xdd, 0x85, 0x86, 0xc2, 0x58, 0x84, 0x29, 0xaa, 0x40,
	0x47, 0xed, 0x35, 0x8d, 0xf1, 0x37, 0x8b, 0xe0, 0x13, 0xb6, 0xd0, 0x2a, 0x75, 0x2e, 0xd3, 0x7f,
	0x83, 0x89, 0x58, 0x48, 0x96, 0xd8, 0xeb, 0x5d, 0xd2, 0xaf, 0xfa, 0xa0, 0x43, 0x23, 0x13, 0xe9,
	0xed, 0x40, 0xb5, 0xa8, 0x8d, 0x6e, 0x80, 0x75, 0x36, 0xf2, 0x5a, 0x25, 0x7d, 0x38, 0xff, 0xdf,
	0x6b, 0x11, 0x77, 0x1b, 0x9a, 0xc6, 0x72, 0x2c, 0x71, 0xc2, 0x67, 0x1c, 0x15, 0x5d, 0x7b, 0x7f,
	0x7b, 0x63, 0x91, 0xde, 0x2d, 0x81, 0xcd, 0xb3, 0x89, 0x3c, 0x41, 0x94, 0x2c, 0xe4, 0x29, 0xd2,
	0x43, 0x68, 0x5d, 0x14, 0x97, 0x40, 0x2a, 0x31, 0xc6, 0xd8, 0x34, 0xb7, 0x3e, 0xdc, 0x71, 0xb2,
	0x09, 0x3b, 0xc5, 0x84, 0x9d, 0xf3, 0xe3, 0x28, 0xd9, 0x1f, 0x9a, 0x32, 0xfd, 0x5f, 0x96, 0x2c,
	0xcf, 0x90, 0xe8, 0x08, 0x9a, 0xf7, 0x42, 0x09, 0x5f, 0xa0, 0xe9, 0xe8, 0x43, 0x32, 0x8d, 0x25,
	0xe7, 0x8c, 0x2f, 0x90, 0x9e, 0x00, 0xbd, 0x17, 0xe1, 0x51, 0x82, 0x2a, 0x65, 0xa1, 0x69, 0xf9,
	0x43, 0x42, 0x5b, 0x4b, 0xde, 0x71, 0x4e, 0xeb, 0x7d, 0x22, 0x00, 0x2e, 0x8f, 0xa6, 0x23, 0x11,
	0xcd, 0xf8, 0x9c, 0x9e, 0x42, 0x33, 0x16, 0x97, 0x6a, 0x82, 0x41, 0x31, 0xf2, 0xac, 0xce, 0xee,
	0x43, 0x4b, 0x94, 0xef, 0xce, 0x5b, 0xb3, 0x3b, 0x8d, 0x4c, 0xa3, 0xd8, 0xc9, 0xff, 0xa0, 0x3a,
	0x53, 0x88, 0x63, 0x1e, 0x4d, 0xf3, 0x7a, 0xdb, 0x2b, 0x36, 0x5d, 0x21, 0xc2, 0xcc, 0xe4, 0x12,
	0x4b, 0x1f, 0x69, 0x33, 0xfa, 0x8d, 0x40, 0xc8, 0x84, 0x8b, 0x28, 0xb6, 0xad, 0xae, 0xd5, 0xaf,
	0x0f, 0x3b, 0xdf, 0x35, 0xf3, 0xd4, 0xe0, 0xf4, 0xfb, 0xf7, 0xb7, 0xb8, 0xf7, 0x8e, 0xc0, 0x46,
	0xe1, 0xe5, 0x78, 0xa9, 0xf9, 0x93, 0x05, 0x1e, 0x95, 0x0a, 0xd9, 0x42, 0xea, 0x2f, 0xa8, 0x48,
	0x2e, 0x8b, 0x11, 0x6e, 0x7f, 0x43, 0x40, 0x7f, 0xe1, 0xa3, 0x92, 0x6f, 0x60, 0x6e, 0x6b, 0xf9,
	0x8d, 0x8a, 0x3d, 0x7b, 0x43, 0xa0, 0x36, 0xe2, 0x53, 0xe5, 0xb3, 0x68, 0x8e, 0xf4, 0x6f, 0x68,
	0xe6, 0xf9, 0x40, 0x2a, 0x9c, 0xf1, 0xab, 0xd5, 0x4f, 0xdf, 0xc8, 0x01, 0x9e, 0xc9, 0xd3, 0x43,
	0x80, 0x0c, 0x19, 0x84, 0x18, 0xfd, 0xc8, 0x26, 0xe5, 0x43, 0xda, 0xb3, 0xec, 0x57, 0xc4, 0xaf,
	0x65, 0xdc, 0xc7, 0x18, 0xb9, 0xff, 0x40, 0x87, 0x8b, 0xcc, 0xbf, 0x54, 0xe2, 0xea, 0x7a, 0xb5,
	0x14, 0x77, 0xf3, 0xa0, 0x78, 0x5a, 0x24, 0xc2, 0x23, 0xe3, 0x75, 0xa3, 0xbf, 0xff, 0x39, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x02, 0x9c, 0x89, 0x34, 0x05, 0x00, 0x00,
}

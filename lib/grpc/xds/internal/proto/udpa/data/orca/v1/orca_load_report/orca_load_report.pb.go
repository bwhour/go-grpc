// Code generated by protoc-gen-go. DO NOT EDIT.
// source: udpa/data/orca/v1/orca_load_report.proto

package v1

import (
	fmt "fmt"

	proto "github.com/bwhour/go-grpc/lib/protobuf/proto"

	math "math"

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

type OrcaLoadReport struct {
	CpuUtilization       float64            `protobuf:"fixed64,1,opt,name=cpu_utilization,json=cpuUtilization,proto3" json:"cpu_utilization,omitempty"`
	MemUtilization       float64            `protobuf:"fixed64,2,opt,name=mem_utilization,json=memUtilization,proto3" json:"mem_utilization,omitempty"`
	Rps                  uint64             `protobuf:"varint,3,opt,name=rps,proto3" json:"rps,omitempty"`
	RequestCost          map[string]float64 `protobuf:"bytes,4,rep,name=request_cost,json=requestCost,proto3" json:"request_cost,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	Utilization          map[string]float64 `protobuf:"bytes,5,rep,name=utilization,proto3" json:"utilization,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrcaLoadReport) Reset()         { *m = OrcaLoadReport{} }
func (m *OrcaLoadReport) String() string { return proto.CompactTextString(m) }
func (*OrcaLoadReport) ProtoMessage()    {}
func (*OrcaLoadReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_orca_load_report_f7a8d2b84dee17a7, []int{0}
}
func (m *OrcaLoadReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrcaLoadReport.Unmarshal(m, b)
}
func (m *OrcaLoadReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrcaLoadReport.Marshal(b, m, deterministic)
}
func (dst *OrcaLoadReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrcaLoadReport.Merge(dst, src)
}
func (m *OrcaLoadReport) XXX_Size() int {
	return xxx_messageInfo_OrcaLoadReport.Size(m)
}
func (m *OrcaLoadReport) XXX_DiscardUnknown() {
	xxx_messageInfo_OrcaLoadReport.DiscardUnknown(m)
}

var xxx_messageInfo_OrcaLoadReport proto.InternalMessageInfo

func (m *OrcaLoadReport) GetCpuUtilization() float64 {
	if m != nil {
		return m.CpuUtilization
	}
	return 0
}

func (m *OrcaLoadReport) GetMemUtilization() float64 {
	if m != nil {
		return m.MemUtilization
	}
	return 0
}

func (m *OrcaLoadReport) GetRps() uint64 {
	if m != nil {
		return m.Rps
	}
	return 0
}

func (m *OrcaLoadReport) GetRequestCost() map[string]float64 {
	if m != nil {
		return m.RequestCost
	}
	return nil
}

func (m *OrcaLoadReport) GetUtilization() map[string]float64 {
	if m != nil {
		return m.Utilization
	}
	return nil
}

func init() {
	proto.RegisterType((*OrcaLoadReport)(nil), "udpa.data.orca.v1.OrcaLoadReport")
	proto.RegisterMapType((map[string]float64)(nil), "udpa.data.orca.v1.OrcaLoadReport.RequestCostEntry")
	proto.RegisterMapType((map[string]float64)(nil), "udpa.data.orca.v1.OrcaLoadReport.UtilizationEntry")
}

func init() {
	proto.RegisterFile("udpa/data/orca/v1/orca_load_report.proto", fileDescriptor_orca_load_report_f7a8d2b84dee17a7)
}

var fileDescriptor_orca_load_report_f7a8d2b84dee17a7 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x4d, 0x29, 0x48,
	0xd4, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0xcf, 0x2f, 0x4a, 0x4e, 0xd4, 0x2f, 0x33, 0x04, 0xd3, 0xf1,
	0x39, 0xf9, 0x89, 0x29, 0xf1, 0x45, 0xa9, 0x05, 0xf9, 0x45, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x82, 0x20, 0x95, 0x7a, 0x20, 0x95, 0x7a, 0x20, 0x15, 0x7a, 0x65, 0x86, 0x52, 0xe2,
	0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0xad, 0xd2, 0x24, 0x16,
	0x2e, 0x3e, 0xff, 0xa2, 0xe4, 0x44, 0x9f, 0xfc, 0xc4, 0x94, 0x20, 0xb0, 0x21, 0x42, 0x9e, 0x5c,
	0xfc, 0xc9, 0x05, 0xa5, 0xf1, 0xa5, 0x25, 0x99, 0x39, 0x99, 0x55, 0x89, 0x25, 0x99, 0xf9, 0x79,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x8c, 0x4e, 0x0a, 0xbb, 0x5e, 0x1e, 0x60, 0xe6, 0x16, 0xe2, 0xd4,
	0x64, 0x80, 0x02, 0x28, 0x5f, 0x12, 0xc2, 0xfb, 0x60, 0x1f, 0xc4, 0x97, 0x5c, 0x50, 0x1a, 0x8a,
	0xd0, 0x07, 0x32, 0x2a, 0x37, 0x35, 0x17, 0xc5, 0x28, 0x26, 0x62, 0x8d, 0xca, 0x4d, 0xcd, 0x45,
	0x36, 0x4a, 0x80, 0x8b, 0xb9, 0xa8, 0xa0, 0x58, 0x82, 0x59, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc4,
	0x14, 0x0a, 0xe5, 0xe2, 0x29, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0x4f, 0xce, 0x2f, 0x2e,
	0x91, 0x60, 0x51, 0x60, 0xd6, 0xe0, 0x36, 0x32, 0xd2, 0xc3, 0xf0, 0xbd, 0x1e, 0xaa, 0x07, 0xf5,
	0x82, 0x20, 0xba, 0x9c, 0xf3, 0x8b, 0x4b, 0x5c, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xb8, 0x8b, 0x10,
	0x22, 0x42, 0x0d, 0x8c, 0x5c, 0xdc, 0xc8, 0x0e, 0x66, 0x25, 0xd6, 0x58, 0x24, 0xd7, 0x82, 0x8d,
	0x75, 0xd2, 0x02, 0x79, 0x4a, 0x60, 0x16, 0x23, 0xaf, 0x16, 0xba, 0x4f, 0x61, 0x82, 0x08, 0xef,
	0x22, 0x5b, 0x29, 0x65, 0xc7, 0x25, 0x80, 0xee, 0x46, 0x90, 0xff, 0xb3, 0x53, 0x2b, 0xc1, 0x31,
	0xc1, 0x19, 0x04, 0x62, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x42, 0x82, 0x34,
	0x08, 0xc2, 0xb1, 0x62, 0xb2, 0x60, 0x04, 0xe9, 0x47, 0x77, 0x0c, 0x29, 0xfa, 0x9d, 0x6c, 0xb8,
	0xe4, 0x33, 0xf3, 0xf5, 0x52, 0xf3, 0xca, 0xf2, 0x2b, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0x31, 0xfd,
	0xee, 0x24, 0x8c, 0xea, 0xf9, 0x00, 0x50, 0x62, 0x0a, 0x60, 0x8c, 0x62, 0x2a, 0x33, 0x4c, 0x62,
	0x03, 0xa7, 0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x12, 0xa0, 0x7f, 0xb1, 0x02,
	0x00, 0x00,
}

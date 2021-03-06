// Code generated by protoc-gen-go. DO NOT EDIT.
// source: udpa/service/orca/v1/orca.proto

package v1

import (
	fmt "fmt"

	proto "github.com/bwhour/go-grpc/lib/protobuf/proto"

	math "math"

	duration "github.com/bwhour/go-grpc/lib/protobuf/ptypes/duration"

	orca_load_report "github.com/bwhour/go-grpc/lib/grpc/xds/internal/proto/udpa/data/orca/v1/orca_load_report"

	_ "github.com/bwhour/go-grpc/lib/grpc/xds/internal/proto/validate"

	context "golang.org/x/net/context"

	grpc "github.com/bwhour/go-grpc/lib/grpc"
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

type OrcaLoadReportRequest struct {
	ReportInterval       *duration.Duration `protobuf:"bytes,1,opt,name=report_interval,json=reportInterval,proto3" json:"report_interval,omitempty"`
	RequestCostNames     []string           `protobuf:"bytes,2,rep,name=request_cost_names,json=requestCostNames,proto3" json:"request_cost_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrcaLoadReportRequest) Reset()         { *m = OrcaLoadReportRequest{} }
func (m *OrcaLoadReportRequest) String() string { return proto.CompactTextString(m) }
func (*OrcaLoadReportRequest) ProtoMessage()    {}
func (*OrcaLoadReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_orca_ca77e509304795c3, []int{0}
}
func (m *OrcaLoadReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrcaLoadReportRequest.Unmarshal(m, b)
}
func (m *OrcaLoadReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrcaLoadReportRequest.Marshal(b, m, deterministic)
}
func (dst *OrcaLoadReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrcaLoadReportRequest.Merge(dst, src)
}
func (m *OrcaLoadReportRequest) XXX_Size() int {
	return xxx_messageInfo_OrcaLoadReportRequest.Size(m)
}
func (m *OrcaLoadReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrcaLoadReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrcaLoadReportRequest proto.InternalMessageInfo

func (m *OrcaLoadReportRequest) GetReportInterval() *duration.Duration {
	if m != nil {
		return m.ReportInterval
	}
	return nil
}

func (m *OrcaLoadReportRequest) GetRequestCostNames() []string {
	if m != nil {
		return m.RequestCostNames
	}
	return nil
}

func init() {
	proto.RegisterType((*OrcaLoadReportRequest)(nil), "udpa.service.orca.v1.OrcaLoadReportRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OpenRcaServiceClient is the client API for OpenRcaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/github.com/bwhour/go-grpc/lib/grpc#ClientConn.NewStream.
type OpenRcaServiceClient interface {
	StreamCoreMetrics(ctx context.Context, in *OrcaLoadReportRequest, opts ...grpc.CallOption) (OpenRcaService_StreamCoreMetricsClient, error)
}

type openRcaServiceClient struct {
	cc *grpc.ClientConn
}

func NewOpenRcaServiceClient(cc *grpc.ClientConn) OpenRcaServiceClient {
	return &openRcaServiceClient{cc}
}

func (c *openRcaServiceClient) StreamCoreMetrics(ctx context.Context, in *OrcaLoadReportRequest, opts ...grpc.CallOption) (OpenRcaService_StreamCoreMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OpenRcaService_serviceDesc.Streams[0], "/udpa.service.orca.v1.OpenRcaService/StreamCoreMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &openRcaServiceStreamCoreMetricsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OpenRcaService_StreamCoreMetricsClient interface {
	Recv() (*orca_load_report.OrcaLoadReport, error)
	grpc.ClientStream
}

type openRcaServiceStreamCoreMetricsClient struct {
	grpc.ClientStream
}

func (x *openRcaServiceStreamCoreMetricsClient) Recv() (*orca_load_report.OrcaLoadReport, error) {
	m := new(orca_load_report.OrcaLoadReport)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OpenRcaServiceServer is the server API for OpenRcaService service.
type OpenRcaServiceServer interface {
	StreamCoreMetrics(*OrcaLoadReportRequest, OpenRcaService_StreamCoreMetricsServer) error
}

func RegisterOpenRcaServiceServer(s *grpc.Server, srv OpenRcaServiceServer) {
	s.RegisterService(&_OpenRcaService_serviceDesc, srv)
}

func _OpenRcaService_StreamCoreMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(OrcaLoadReportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OpenRcaServiceServer).StreamCoreMetrics(m, &openRcaServiceStreamCoreMetricsServer{stream})
}

type OpenRcaService_StreamCoreMetricsServer interface {
	Send(*orca_load_report.OrcaLoadReport) error
	grpc.ServerStream
}

type openRcaServiceStreamCoreMetricsServer struct {
	grpc.ServerStream
}

func (x *openRcaServiceStreamCoreMetricsServer) Send(m *orca_load_report.OrcaLoadReport) error {
	return x.ServerStream.SendMsg(m)
}

var _OpenRcaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "udpa.service.orca.v1.OpenRcaService",
	HandlerType: (*OpenRcaServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCoreMetrics",
			Handler:       _OpenRcaService_StreamCoreMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "udpa/service/orca/v1/orca.proto",
}

func init() {
	proto.RegisterFile("udpa/service/orca/v1/orca.proto", fileDescriptor_orca_ca77e509304795c3)
}

var fileDescriptor_orca_ca77e509304795c3 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x05, 0xa1, 0x2b, 0x54, 0x0d, 0x8a, 0xb5, 0x07, 0xad, 0x3d, 0x15, 0x94, 0x8d,
	0xad, 0xf8, 0x05, 0x5a, 0x2f, 0x82, 0xda, 0x92, 0xde, 0xbc, 0x84, 0x69, 0x32, 0x96, 0x85, 0x34,
	0x13, 0x67, 0x37, 0xab, 0xfd, 0x08, 0x7e, 0x6b, 0xc9, 0x6e, 0x7a, 0x10, 0xe2, 0x69, 0xff, 0xbc,
	0xf7, 0x9b, 0xc7, 0x3c, 0x71, 0x5d, 0x65, 0x25, 0x44, 0x1a, 0xd9, 0xaa, 0x14, 0x23, 0xe2, 0x14,
	0x22, 0x3b, 0x71, 0xa7, 0x2c, 0x99, 0x0c, 0x85, 0x67, 0xb5, 0x41, 0x36, 0x06, 0xe9, 0x04, 0x3b,
	0x19, 0x8c, 0x1d, 0x96, 0x81, 0x81, 0x3f, 0x4c, 0x92, 0x13, 0x64, 0x09, 0x63, 0x49, 0x6c, 0x3c,
	0x3f, 0xb8, 0xda, 0x10, 0x6d, 0x72, 0x8c, 0xdc, 0x6b, 0x5d, 0x7d, 0x44, 0x59, 0xc5, 0x60, 0x14,
	0x15, 0x8d, 0x7e, 0x61, 0x21, 0x57, 0x19, 0x18, 0x8c, 0xf6, 0x17, 0x2f, 0x8c, 0x7e, 0x02, 0x71,
	0xbe, 0xe0, 0x14, 0x5e, 0x08, 0xb2, 0xd8, 0x4d, 0x8c, 0xf1, 0xb3, 0x42, 0x6d, 0xc2, 0x99, 0x38,
	0xf6, 0x11, 0x89, 0x2a, 0x0c, 0xb2, 0x85, 0xbc, 0x1f, 0x0c, 0x83, 0xf1, 0xd1, 0xf4, 0x52, 0xfa,
	0x30, 0xb9, 0x0f, 0x93, 0x4f, 0x4d, 0x58, 0xdc, 0xf3, 0xc4, 0x73, 0x03, 0x84, 0x77, 0x22, 0x64,
	0x3f, 0x2e, 0x49, 0x49, 0x9b, 0xa4, 0x80, 0x2d, 0xea, 0x7e, 0x67, 0x78, 0x30, 0xee, 0xc6, 0x27,
	0x8d, 0x32, 0x27, 0x6d, 0xde, 0xea, 0xff, 0xe9, 0x97, 0xe8, 0x2d, 0x4a, 0x2c, 0xe2, 0x14, 0x56,
	0xbe, 0x88, 0x10, 0xc5, 0xe9, 0xca, 0x30, 0xc2, 0x76, 0x4e, 0x8c, 0xaf, 0x68, 0x58, 0xa5, 0x3a,
	0xbc, 0x95, 0x6d, 0x65, 0xc9, 0xd6, 0x2d, 0x06, 0x37, 0xde, 0x5c, 0x77, 0xf8, 0x8f, 0xf3, 0x3e,
	0x98, 0x3d, 0x8a, 0x91, 0x22, 0x89, 0x85, 0xa5, 0x5d, 0xc9, 0xf4, 0xbd, 0x6b, 0x0d, 0x98, 0x75,
	0x6b, 0x6e, 0x59, 0xef, 0xbc, 0x0c, 0xde, 0x3b, 0x76, 0xb2, 0x3e, 0x74, 0x05, 0x3c, 0xfc, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xac, 0x32, 0x62, 0x96, 0xde, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hlltc.proto

package hlltc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EstimateRequest struct {
	Sketch               []byte   `protobuf:"bytes,1,opt,name=sketch,proto3" json:"sketch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateRequest) Reset()         { *m = EstimateRequest{} }
func (m *EstimateRequest) String() string { return proto.CompactTextString(m) }
func (*EstimateRequest) ProtoMessage()    {}
func (*EstimateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870f57255fd25983, []int{0}
}

func (m *EstimateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateRequest.Unmarshal(m, b)
}
func (m *EstimateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateRequest.Marshal(b, m, deterministic)
}
func (m *EstimateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateRequest.Merge(m, src)
}
func (m *EstimateRequest) XXX_Size() int {
	return xxx_messageInfo_EstimateRequest.Size(m)
}
func (m *EstimateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateRequest proto.InternalMessageInfo

func (m *EstimateRequest) GetSketch() []byte {
	if m != nil {
		return m.Sketch
	}
	return nil
}

type EstimateReply struct {
	Sum                  uint64   `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstimateReply) Reset()         { *m = EstimateReply{} }
func (m *EstimateReply) String() string { return proto.CompactTextString(m) }
func (*EstimateReply) ProtoMessage()    {}
func (*EstimateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_870f57255fd25983, []int{1}
}

func (m *EstimateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateReply.Unmarshal(m, b)
}
func (m *EstimateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateReply.Marshal(b, m, deterministic)
}
func (m *EstimateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateReply.Merge(m, src)
}
func (m *EstimateReply) XXX_Size() int {
	return xxx_messageInfo_EstimateReply.Size(m)
}
func (m *EstimateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateReply.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateReply proto.InternalMessageInfo

func (m *EstimateReply) GetSum() uint64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type MergeRequest struct {
	Sketch               []byte   `protobuf:"bytes,1,opt,name=sketch,proto3" json:"sketch,omitempty"`
	Sketches             [][]byte `protobuf:"bytes,3,rep,name=sketches,proto3" json:"sketches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeRequest) Reset()         { *m = MergeRequest{} }
func (m *MergeRequest) String() string { return proto.CompactTextString(m) }
func (*MergeRequest) ProtoMessage()    {}
func (*MergeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_870f57255fd25983, []int{2}
}

func (m *MergeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeRequest.Unmarshal(m, b)
}
func (m *MergeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeRequest.Marshal(b, m, deterministic)
}
func (m *MergeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeRequest.Merge(m, src)
}
func (m *MergeRequest) XXX_Size() int {
	return xxx_messageInfo_MergeRequest.Size(m)
}
func (m *MergeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MergeRequest proto.InternalMessageInfo

func (m *MergeRequest) GetSketch() []byte {
	if m != nil {
		return m.Sketch
	}
	return nil
}

func (m *MergeRequest) GetSketches() [][]byte {
	if m != nil {
		return m.Sketches
	}
	return nil
}

type MergeReply struct {
	MergedSketch         []byte   `protobuf:"bytes,1,opt,name=merged_sketch,json=mergedSketch,proto3" json:"merged_sketch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeReply) Reset()         { *m = MergeReply{} }
func (m *MergeReply) String() string { return proto.CompactTextString(m) }
func (*MergeReply) ProtoMessage()    {}
func (*MergeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_870f57255fd25983, []int{3}
}

func (m *MergeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeReply.Unmarshal(m, b)
}
func (m *MergeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeReply.Marshal(b, m, deterministic)
}
func (m *MergeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeReply.Merge(m, src)
}
func (m *MergeReply) XXX_Size() int {
	return xxx_messageInfo_MergeReply.Size(m)
}
func (m *MergeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeReply.DiscardUnknown(m)
}

var xxx_messageInfo_MergeReply proto.InternalMessageInfo

func (m *MergeReply) GetMergedSketch() []byte {
	if m != nil {
		return m.MergedSketch
	}
	return nil
}

func init() {
	proto.RegisterType((*EstimateRequest)(nil), "EstimateRequest")
	proto.RegisterType((*EstimateReply)(nil), "EstimateReply")
	proto.RegisterType((*MergeRequest)(nil), "MergeRequest")
	proto.RegisterType((*MergeReply)(nil), "MergeReply")
}

func init() { proto.RegisterFile("hlltc.proto", fileDescriptor_870f57255fd25983) }

var fileDescriptor_870f57255fd25983 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc8, 0xc9, 0x29,
	0x49, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe4, 0xe2, 0x77, 0x2d, 0x2e, 0xc9, 0xcc,
	0x4d, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0xce,
	0x4e, 0x2d, 0x49, 0xce, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0xf2, 0x94, 0x14, 0xb9,
	0x78, 0x11, 0x4a, 0x0b, 0x72, 0x2a, 0x85, 0x04, 0xb8, 0x98, 0x8b, 0x4b, 0x73, 0xc1, 0xaa, 0x58,
	0x82, 0x40, 0x4c, 0x25, 0x27, 0x2e, 0x1e, 0xdf, 0xd4, 0xa2, 0x74, 0x42, 0x46, 0x09, 0x49, 0x71,
	0x71, 0x40, 0x58, 0xa9, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0x3c, 0x41, 0x70, 0xbe, 0x92, 0x21,
	0x17, 0x17, 0xd4, 0x0c, 0x90, 0x1d, 0xca, 0x5c, 0xbc, 0xb9, 0x20, 0x5e, 0x4a, 0x3c, 0x8a, 0x41,
	0x3c, 0x10, 0xc1, 0x60, 0xb0, 0x98, 0x91, 0x35, 0x17, 0x27, 0xd4, 0x65, 0xf9, 0x45, 0x42, 0x7a,
	0x5c, 0x1c, 0x30, 0x67, 0x0a, 0x09, 0xe8, 0xa1, 0x79, 0x4e, 0x8a, 0x4f, 0x0f, 0xc5, 0x0f, 0x4a,
	0x0c, 0x46, 0xfa, 0x5c, 0x6c, 0x60, 0xfb, 0x8a, 0x84, 0x54, 0xb9, 0x58, 0xc1, 0x2c, 0x21, 0x5e,
	0x3d, 0x64, 0x5f, 0x48, 0x71, 0xeb, 0x21, 0x1c, 0xa4, 0xc4, 0xe0, 0xc4, 0x1e, 0xc5, 0x0a, 0x0e,
	0xc1, 0x24, 0x36, 0x70, 0x10, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x18, 0x86, 0x29,
	0x51, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EstimatorClient is the client API for Estimator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EstimatorClient interface {
	Estimate(ctx context.Context, in *EstimateRequest, opts ...grpc.CallOption) (*EstimateReply, error)
}

type estimatorClient struct {
	cc *grpc.ClientConn
}

func NewEstimatorClient(cc *grpc.ClientConn) EstimatorClient {
	return &estimatorClient{cc}
}

func (c *estimatorClient) Estimate(ctx context.Context, in *EstimateRequest, opts ...grpc.CallOption) (*EstimateReply, error) {
	out := new(EstimateReply)
	err := c.cc.Invoke(ctx, "/Estimator/Estimate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EstimatorServer is the server API for Estimator service.
type EstimatorServer interface {
	Estimate(context.Context, *EstimateRequest) (*EstimateReply, error)
}

func RegisterEstimatorServer(s *grpc.Server, srv EstimatorServer) {
	s.RegisterService(&_Estimator_serviceDesc, srv)
}

func _Estimator_Estimate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EstimatorServer).Estimate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Estimator/Estimate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EstimatorServer).Estimate(ctx, req.(*EstimateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Estimator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Estimator",
	HandlerType: (*EstimatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Estimate",
			Handler:    _Estimator_Estimate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hlltc.proto",
}

// MergerClient is the client API for Merger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MergerClient interface {
	Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*MergeReply, error)
}

type mergerClient struct {
	cc *grpc.ClientConn
}

func NewMergerClient(cc *grpc.ClientConn) MergerClient {
	return &mergerClient{cc}
}

func (c *mergerClient) Merge(ctx context.Context, in *MergeRequest, opts ...grpc.CallOption) (*MergeReply, error) {
	out := new(MergeReply)
	err := c.cc.Invoke(ctx, "/Merger/Merge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MergerServer is the server API for Merger service.
type MergerServer interface {
	Merge(context.Context, *MergeRequest) (*MergeReply, error)
}

func RegisterMergerServer(s *grpc.Server, srv MergerServer) {
	s.RegisterService(&_Merger_serviceDesc, srv)
}

func _Merger_Merge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MergerServer).Merge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Merger/Merge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MergerServer).Merge(ctx, req.(*MergeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Merger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Merger",
	HandlerType: (*MergerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Merge",
			Handler:    _Merger_Merge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hlltc.proto",
}

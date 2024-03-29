// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/depend.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DependRequest struct {
	DependName           string   `protobuf:"bytes,1,opt,name=dependName,proto3" json:"dependName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DependRequest) Reset()         { *m = DependRequest{} }
func (m *DependRequest) String() string { return proto.CompactTextString(m) }
func (*DependRequest) ProtoMessage()    {}
func (*DependRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42b5bb14098327a, []int{0}
}

func (m *DependRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DependRequest.Unmarshal(m, b)
}
func (m *DependRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DependRequest.Marshal(b, m, deterministic)
}
func (m *DependRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DependRequest.Merge(m, src)
}
func (m *DependRequest) XXX_Size() int {
	return xxx_messageInfo_DependRequest.Size(m)
}
func (m *DependRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DependRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DependRequest proto.InternalMessageInfo

func (m *DependRequest) GetDependName() string {
	if m != nil {
		return m.DependName
	}
	return ""
}

type DependResponse struct {
	DependMessage        string   `protobuf:"bytes,1,opt,name=dependMessage,proto3" json:"dependMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DependResponse) Reset()         { *m = DependResponse{} }
func (m *DependResponse) String() string { return proto.CompactTextString(m) }
func (*DependResponse) ProtoMessage()    {}
func (*DependResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42b5bb14098327a, []int{1}
}

func (m *DependResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DependResponse.Unmarshal(m, b)
}
func (m *DependResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DependResponse.Marshal(b, m, deterministic)
}
func (m *DependResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DependResponse.Merge(m, src)
}
func (m *DependResponse) XXX_Size() int {
	return xxx_messageInfo_DependResponse.Size(m)
}
func (m *DependResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DependResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DependResponse proto.InternalMessageInfo

func (m *DependResponse) GetDependMessage() string {
	if m != nil {
		return m.DependMessage
	}
	return ""
}

type Bot struct {
	Run                  string   `protobuf:"bytes,1,opt,name=run,proto3" json:"run,omitempty"`
	Say                  []byte   `protobuf:"bytes,2,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bot) Reset()         { *m = Bot{} }
func (m *Bot) String() string { return proto.CompactTextString(m) }
func (*Bot) ProtoMessage()    {}
func (*Bot) Descriptor() ([]byte, []int) {
	return fileDescriptor_d42b5bb14098327a, []int{2}
}

func (m *Bot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bot.Unmarshal(m, b)
}
func (m *Bot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bot.Marshal(b, m, deterministic)
}
func (m *Bot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bot.Merge(m, src)
}
func (m *Bot) XXX_Size() int {
	return xxx_messageInfo_Bot.Size(m)
}
func (m *Bot) XXX_DiscardUnknown() {
	xxx_messageInfo_Bot.DiscardUnknown(m)
}

var xxx_messageInfo_Bot proto.InternalMessageInfo

func (m *Bot) GetRun() string {
	if m != nil {
		return m.Run
	}
	return ""
}

func (m *Bot) GetSay() []byte {
	if m != nil {
		return m.Say
	}
	return nil
}

func init() {
	proto.RegisterType((*DependRequest)(nil), "proto.DependRequest")
	proto.RegisterType((*DependResponse)(nil), "proto.DependResponse")
	proto.RegisterType((*Bot)(nil), "proto.Bot")
}

func init() { proto.RegisterFile("proto/depend.proto", fileDescriptor_d42b5bb14098327a) }

var fileDescriptor_d42b5bb14098327a = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0x2d, 0x48, 0xcd, 0x4b, 0xd1, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x92,
	0x3e, 0x17, 0xaf, 0x0b, 0x58, 0x38, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8e, 0x8b,
	0x0b, 0xa2, 0xce, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x49, 0x44,
	0xc9, 0x8c, 0x8b, 0x0f, 0xa6, 0xa1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x85, 0x8b, 0x17,
	0x22, 0xef, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x0e, 0xd3, 0x84, 0x2a, 0xa8, 0xa4, 0xc9, 0xc5, 0xec,
	0x94, 0x5f, 0x22, 0x24, 0xc0, 0xc5, 0x5c, 0x54, 0x9a, 0x07, 0x55, 0x02, 0x62, 0x82, 0x44, 0x8a,
	0x13, 0x2b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x40, 0x4c, 0x23, 0x1f, 0x2e, 0x1e, 0x88,
	0x15, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0x42, 0x36, 0x5c, 0xdc, 0x10, 0xbe, 0x47, 0x6a, 0x4e,
	0x4e, 0xbe, 0x90, 0x08, 0xc4, 0x07, 0x7a, 0x28, 0xee, 0x96, 0x12, 0x45, 0x13, 0x85, 0x38, 0x4e,
	0x89, 0x21, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x63, 0x22, 0xae, 0xf7,
	0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DependServerClient is the client API for DependServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DependServerClient interface {
	DependHello(ctx context.Context, in *DependRequest, opts ...grpc.CallOption) (*DependResponse, error)
}

type dependServerClient struct {
	cc *grpc.ClientConn
}

func NewDependServerClient(cc *grpc.ClientConn) DependServerClient {
	return &dependServerClient{cc}
}

func (c *dependServerClient) DependHello(ctx context.Context, in *DependRequest, opts ...grpc.CallOption) (*DependResponse, error) {
	out := new(DependResponse)
	err := c.cc.Invoke(ctx, "/proto.DependServer/DependHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependServerServer is the server API for DependServer service.
type DependServerServer interface {
	DependHello(context.Context, *DependRequest) (*DependResponse, error)
}

// UnimplementedDependServerServer can be embedded to have forward compatible implementations.
type UnimplementedDependServerServer struct {
}

func (*UnimplementedDependServerServer) DependHello(ctx context.Context, req *DependRequest) (*DependResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DependHello not implemented")
}

func RegisterDependServerServer(s *grpc.Server, srv DependServerServer) {
	s.RegisterService(&_DependServer_serviceDesc, srv)
}

func _DependServer_DependHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependServerServer).DependHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DependServer/DependHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependServerServer).DependHello(ctx, req.(*DependRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DependServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DependServer",
	HandlerType: (*DependServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DependHello",
			Handler:    _DependServer_DependHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/depend.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user/user.proto

package user

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

type GetUserRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type UserResponse struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed89022014131a74, []int{1}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "user.GetUserRequest")
	proto.RegisterType((*UserResponse)(nil), "user.UserResponse")
}

func init() { proto.RegisterFile("user/user.proto", fileDescriptor_ed89022014131a74) }

var fileDescriptor_ed89022014131a74 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2d, 0x4e, 0x2d,
	0xd2, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x06, 0x17,
	0x9f, 0x7b, 0x6a, 0x49, 0x68, 0x71, 0x6a, 0x51, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x18, 0x17, 0x1b, 0x48, 0xc6, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xca, 0x53,
	0xca, 0xe0, 0xe2, 0x81, 0x28, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5, 0xa5, 0x4e, 0x48, 0x88,
	0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x0a, 0x66, 0x0b, 0x89, 0x70, 0xb1, 0xa6,
	0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30, 0x83, 0x05, 0x21, 0x1c, 0x90, 0x09, 0x89, 0x65, 0x89, 0x25,
	0x89, 0x45, 0x12, 0x2c, 0x10, 0x13, 0x20, 0x3c, 0x23, 0x17, 0x2e, 0x6e, 0x90, 0x4d, 0xc1, 0xa9,
	0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xa6, 0x5c, 0xec, 0xe9, 0x10, 0x27, 0x0a, 0x89, 0xe8, 0x81,
	0x3d, 0x80, 0xea, 0x62, 0x29, 0x21, 0x88, 0x28, 0xb2, 0xeb, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xde,
	0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x33, 0x13, 0x3e, 0xf9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/getUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getUser",
			Handler:    _UserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
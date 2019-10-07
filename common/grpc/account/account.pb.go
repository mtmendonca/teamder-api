// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

package account

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

type LoginRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Provider             string   `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Avatar               string   `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRequest) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
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
	return fileDescriptor_d66906c5773c9d08, []int{2}
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

type Skill struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Skill) Reset()         { *m = Skill{} }
func (m *Skill) String() string { return proto.CompactTextString(m) }
func (*Skill) ProtoMessage()    {}
func (*Skill) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{3}
}

func (m *Skill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Skill.Unmarshal(m, b)
}
func (m *Skill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Skill.Marshal(b, m, deterministic)
}
func (m *Skill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Skill.Merge(m, src)
}
func (m *Skill) XXX_Size() int {
	return xxx_messageInfo_Skill.Size(m)
}
func (m *Skill) XXX_DiscardUnknown() {
	xxx_messageInfo_Skill.DiscardUnknown(m)
}

var xxx_messageInfo_Skill proto.InternalMessageInfo

func (m *Skill) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Skill) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type SetProfileRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Education            string   `protobuf:"bytes,2,opt,name=education,proto3" json:"education,omitempty"`
	Experience           string   `protobuf:"bytes,3,opt,name=experience,proto3" json:"experience,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	Skills               []*Skill `protobuf:"bytes,5,rep,name=skills,proto3" json:"skills,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetProfileRequest) Reset()         { *m = SetProfileRequest{} }
func (m *SetProfileRequest) String() string { return proto.CompactTextString(m) }
func (*SetProfileRequest) ProtoMessage()    {}
func (*SetProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{4}
}

func (m *SetProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetProfileRequest.Unmarshal(m, b)
}
func (m *SetProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetProfileRequest.Marshal(b, m, deterministic)
}
func (m *SetProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProfileRequest.Merge(m, src)
}
func (m *SetProfileRequest) XXX_Size() int {
	return xxx_messageInfo_SetProfileRequest.Size(m)
}
func (m *SetProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetProfileRequest proto.InternalMessageInfo

func (m *SetProfileRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *SetProfileRequest) GetEducation() string {
	if m != nil {
		return m.Education
	}
	return ""
}

func (m *SetProfileRequest) GetExperience() string {
	if m != nil {
		return m.Experience
	}
	return ""
}

func (m *SetProfileRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *SetProfileRequest) GetSkills() []*Skill {
	if m != nil {
		return m.Skills
	}
	return nil
}

type SetProfileResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetProfileResponse) Reset()         { *m = SetProfileResponse{} }
func (m *SetProfileResponse) String() string { return proto.CompactTextString(m) }
func (*SetProfileResponse) ProtoMessage()    {}
func (*SetProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{5}
}

func (m *SetProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetProfileResponse.Unmarshal(m, b)
}
func (m *SetProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetProfileResponse.Marshal(b, m, deterministic)
}
func (m *SetProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetProfileResponse.Merge(m, src)
}
func (m *SetProfileResponse) XXX_Size() int {
	return xxx_messageInfo_SetProfileResponse.Size(m)
}
func (m *SetProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetProfileResponse proto.InternalMessageInfo

func (m *SetProfileResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "account.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "account.LoginResponse")
	proto.RegisterType((*UserResponse)(nil), "account.UserResponse")
	proto.RegisterType((*Skill)(nil), "account.Skill")
	proto.RegisterType((*SetProfileRequest)(nil), "account.SetProfileRequest")
	proto.RegisterType((*SetProfileResponse)(nil), "account.SetProfileResponse")
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_d66906c5773c9d08) }

var fileDescriptor_d66906c5773c9d08 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xdf, 0x6e, 0x6b, 0xb7, 0x9d, 0x77, 0xef, 0xe0, 0x0d, 0x6e, 0x94, 0x2a, 0x32, 0x0a,
	0xca, 0xae, 0x26, 0xce, 0x1b, 0x6f, 0x05, 0x41, 0x04, 0x2f, 0xa4, 0xc3, 0x0f, 0x50, 0xb3, 0xa3,
	0x86, 0x65, 0x49, 0x4d, 0xd2, 0xe2, 0xb5, 0x5f, 0xc1, 0xcf, 0xe0, 0xf7, 0x94, 0xa5, 0x69, 0xf7,
	0x87, 0x7a, 0xd5, 0x3c, 0xcf, 0x09, 0xcf, 0xf9, 0xf5, 0xe4, 0xc0, 0x28, 0xa5, 0x54, 0xe6, 0xc2,
	0x5c, 0xb8, 0xef, 0x2c, 0x53, 0xd2, 0x48, 0xd2, 0x75, 0x32, 0xfe, 0xf4, 0x60, 0xf0, 0x20, 0x5f,
	0x99, 0x48, 0xf0, 0x3d, 0x47, 0x6d, 0xc8, 0x11, 0xf8, 0x46, 0xae, 0x50, 0x84, 0xde, 0xc4, 0x9b,
	0xf6, 0x93, 0x52, 0x90, 0x08, 0x7a, 0x99, 0x92, 0x05, 0x5b, 0xa2, 0x0a, 0x5b, 0xb6, 0x50, 0x6b,
	0x42, 0xa0, 0x23, 0xd2, 0x35, 0x86, 0x6d, 0xeb, 0xdb, 0xf3, 0x26, 0x05, 0xd7, 0x29, 0xe3, 0x61,
	0xa7, 0x4c, 0xb1, 0x82, 0x8c, 0x21, 0x48, 0x8b, 0xd4, 0xa4, 0x2a, 0xf4, 0xad, 0xed, 0x54, 0x7c,
	0x06, 0xff, 0x1c, 0x83, 0xce, 0xa4, 0xd0, 0xd8, 0x0c, 0x11, 0xbf, 0xc1, 0xe0, 0x49, 0xa3, 0xaa,
	0x6f, 0x8d, 0x21, 0xc8, 0x35, 0xaa, 0xfb, 0x5b, 0x77, 0xcd, 0xa9, 0x1a, 0xa8, 0xd5, 0x04, 0xd4,
	0x6e, 0x06, 0xea, 0xec, 0x01, 0x5d, 0x82, 0xbf, 0x58, 0x31, 0xce, 0xeb, 0x28, 0x6f, 0x3f, 0x8a,
	0x63, 0x81, 0xdc, 0xe5, 0x97, 0x22, 0xfe, 0xf6, 0xe0, 0xff, 0x02, 0xcd, 0xa3, 0x92, 0x2f, 0x8c,
	0x63, 0x35, 0xcd, 0xdf, 0x10, 0x4f, 0xa0, 0x8f, 0xcb, 0x9c, 0xa6, 0x86, 0x49, 0xe1, 0x72, 0xb6,
	0x06, 0x39, 0x05, 0xc0, 0x8f, 0x0c, 0x15, 0x43, 0x41, 0xab, 0xb9, 0xee, 0x38, 0x1b, 0x2a, 0x25,
	0x39, 0x3a, 0x68, 0x7b, 0x26, 0xe7, 0x10, 0xe8, 0x0d, 0xb2, 0x0e, 0xfd, 0x49, 0x7b, 0xfa, 0x77,
	0x3e, 0x9c, 0x55, 0x2f, 0x6e, 0xff, 0x24, 0x71, 0xd5, 0x78, 0x06, 0x64, 0x17, 0xd3, 0x8d, 0x32,
	0x84, 0xae, 0xce, 0x29, 0x45, 0xad, 0x2d, 0x68, 0x2f, 0xa9, 0xe4, 0xfc, 0xcb, 0x83, 0xe1, 0x4d,
	0x99, 0xb4, 0x40, 0x55, 0x30, 0x8a, 0xe4, 0x1a, 0x7c, 0xfb, 0x5c, 0x64, 0x54, 0xf7, 0xd8, 0x5d,
	0xa1, 0x68, 0x7c, 0x68, 0x97, 0x4d, 0xe2, 0x3f, 0xe4, 0x0e, 0x60, 0xdb, 0x9c, 0x44, 0x5b, 0xc4,
	0xc3, 0xc1, 0x45, 0xc7, 0x8d, 0xb5, 0x2a, 0xe8, 0x39, 0xb0, 0x6b, 0x7c, 0xf5, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x0a, 0x77, 0xd9, 0x69, 0xdf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	SetProfile(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*SetProfileResponse, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SetProfile(ctx context.Context, in *SetProfileRequest, opts ...grpc.CallOption) (*SetProfileResponse, error) {
	out := new(SetProfileResponse)
	err := c.cc.Invoke(ctx, "/account.AccountService/SetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	SetProfile(context.Context, *SetProfileRequest) (*SetProfileResponse, error)
}

// UnimplementedAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (*UnimplementedAccountServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAccountServiceServer) SetProfile(ctx context.Context, req *SetProfileRequest) (*SetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfile not implemented")
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/SetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SetProfile(ctx, req.(*SetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AccountService_Login_Handler,
		},
		{
			MethodName: "SetProfile",
			Handler:    _AccountService_SetProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}

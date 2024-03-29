// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package usecase

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

type User struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	EmailId              string   `protobuf:"bytes,3,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Token                string   `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *User) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UserCreateResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCreateResponse) Reset()         { *m = UserCreateResponse{} }
func (m *UserCreateResponse) String() string { return proto.CompactTextString(m) }
func (*UserCreateResponse) ProtoMessage()    {}
func (*UserCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreateResponse.Unmarshal(m, b)
}
func (m *UserCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreateResponse.Marshal(b, m, deterministic)
}
func (m *UserCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreateResponse.Merge(m, src)
}
func (m *UserCreateResponse) XXX_Size() int {
	return xxx_messageInfo_UserCreateResponse.Size(m)
}
func (m *UserCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreateResponse proto.InternalMessageInfo

func (m *UserCreateResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AllUserResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllUserResponse) Reset()         { *m = AllUserResponse{} }
func (m *AllUserResponse) String() string { return proto.CompactTextString(m) }
func (*AllUserResponse) ProtoMessage()    {}
func (*AllUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *AllUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllUserResponse.Unmarshal(m, b)
}
func (m *AllUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllUserResponse.Marshal(b, m, deterministic)
}
func (m *AllUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllUserResponse.Merge(m, src)
}
func (m *AllUserResponse) XXX_Size() int {
	return xxx_messageInfo_AllUserResponse.Size(m)
}
func (m *AllUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllUserResponse proto.InternalMessageInfo

func (m *AllUserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *AllUserResponse) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "usecase.User")
	proto.RegisterType((*UserCreateResponse)(nil), "usecase.UserCreateResponse")
	proto.RegisterType((*Request)(nil), "usecase.Request")
	proto.RegisterType((*Error)(nil), "usecase.Error")
	proto.RegisterType((*AllUserResponse)(nil), "usecase.AllUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x3d, 0x4f, 0xe3, 0x40,
	0x10, 0x86, 0x63, 0xc7, 0x1f, 0xc9, 0x44, 0x97, 0x3b, 0x8d, 0xae, 0xd8, 0xcb, 0xe9, 0x24, 0x9f,
	0x91, 0x50, 0xaa, 0x14, 0x41, 0x82, 0x8a, 0xc2, 0x42, 0x08, 0xd1, 0x1a, 0xd1, 0x12, 0x19, 0xef,
	0x14, 0x16, 0x8e, 0xd7, 0xec, 0xae, 0xa1, 0xe2, 0x67, 0xf1, 0xff, 0xd0, 0x8e, 0x1d, 0x8b, 0x40,
	0xb7, 0xf3, 0x3e, 0xf3, 0xf5, 0x8e, 0x16, 0xa0, 0x33, 0xa4, 0x37, 0xad, 0x56, 0x56, 0x61, 0xdc,
	0x19, 0x2a, 0x0b, 0x43, 0xe9, 0xbb, 0x07, 0xc1, 0xbd, 0x21, 0x8d, 0x4b, 0xf0, 0x2b, 0x29, 0xbc,
	0xc4, 0x5b, 0x4f, 0x73, 0xbf, 0x92, 0x88, 0x10, 0x34, 0xc5, 0x9e, 0x84, 0x9f, 0x78, 0xeb, 0x79,
	0xce, 0x6f, 0xfc, 0x03, 0x33, 0xda, 0x17, 0x55, 0xbd, 0xab, 0xa4, 0x98, 0xb2, 0x1e, 0x73, 0x7c,
	0x2b, 0x71, 0x05, 0xb3, 0xb6, 0x30, 0xe6, 0x55, 0x69, 0x29, 0x02, 0x46, 0x63, 0x8c, 0xff, 0x00,
	0x4a, 0x4d, 0x85, 0x25, 0xb9, 0x2b, 0xac, 0x08, 0x99, 0xce, 0x07, 0x25, 0xb3, 0x0e, 0x77, 0xad,
	0x3c, 0xe0, 0xa8, 0xc7, 0x83, 0x92, 0x59, 0xfc, 0x0d, 0xa1, 0x55, 0x4f, 0xd4, 0x88, 0x98, 0x49,
	0x1f, 0xa4, 0x17, 0x80, 0x6e, 0xed, 0x2b, 0xee, 0x92, 0x93, 0x69, 0x55, 0x63, 0x08, 0xff, 0x43,
	0xe0, 0x4c, 0xb2, 0x8d, 0xc5, 0xf6, 0xc7, 0x66, 0x70, 0xb9, 0x71, 0xa9, 0x39, 0xa3, 0x74, 0x0e,
	0x71, 0x4e, 0xcf, 0x1d, 0x19, 0x9b, 0x5e, 0x42, 0x78, 0xad, 0xb5, 0xd2, 0xce, 0x6b, 0xa9, 0x24,
	0x71, 0x59, 0x98, 0xf3, 0x1b, 0x13, 0x58, 0x48, 0x32, 0xa5, 0xae, 0x5a, 0x5b, 0xa9, 0x66, 0x38,
	0xc3, 0x67, 0x29, 0x7d, 0x80, 0x9f, 0x59, 0x5d, 0x73, 0xeb, 0xc3, 0xfc, 0x13, 0x08, 0xdd, 0x10,
	0x23, 0xbc, 0x64, 0xfa, 0x7d, 0x81, 0x9e, 0xe1, 0x29, 0x44, 0xe4, 0xc6, 0x1a, 0xe1, 0x73, 0xd6,
	0x72, 0xcc, 0xe2, 0x6d, 0xf2, 0x81, 0x6e, 0xdf, 0x60, 0xe1, 0xca, 0xee, 0x48, 0xbf, 0x54, 0x25,
	0xe1, 0x39, 0x44, 0xbd, 0x5b, 0x3c, 0x6e, 0xbb, 0xfa, 0x7b, 0x14, 0x1e, 0x5f, 0x24, 0x9d, 0xb8,
	0xba, 0x1b, 0xb2, 0x59, 0x5d, 0xe3, 0xaf, 0x31, 0x71, 0xb8, 0xc0, 0x4a, 0x8c, 0xca, 0x17, 0x27,
	0xe9, 0xe4, 0x31, 0xe2, 0x9f, 0x72, 0xf6, 0x11, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x33, 0xbc, 0x13,
	0x37, 0x02, 0x00, 0x00,
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
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserCreateResponse, error)
	GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserCreateResponse, error) {
	out := new(UserCreateResponse)
	err := c.cc.Invoke(ctx, "/usecase.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...grpc.CallOption) (*AllUserResponse, error) {
	out := new(AllUserResponse)
	err := c.cc.Invoke(ctx, "/usecase.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *User) (*UserCreateResponse, error)
	GetAll(context.Context, *Request) (*AllUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Create(ctx context.Context, req *User) (*UserCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedUserServiceServer) GetAll(ctx context.Context, req *Request) (*AllUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usecase.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usecase.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usecase.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

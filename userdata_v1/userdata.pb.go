// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userdata.proto

/*
Package userdata_v1 is a generated protocol buffer package.

It is generated from these files:
	userdata.proto

It has these top-level messages:
	CreateUserRequest
	GetUserRequest
	User
*/
package userdata_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateUserRequest struct {
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateUserRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *CreateUserRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type GetUserRequest struct {
	// Either id or Email is expected
	//
	// Types that are valid to be assigned to UniqueUserOneof:
	//	*GetUserRequest_Id
	//	*GetUserRequest_Email
	UniqueUserOneof isGetUserRequest_UniqueUserOneof `protobuf_oneof:"unique_user_oneof"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isGetUserRequest_UniqueUserOneof interface {
	isGetUserRequest_UniqueUserOneof()
}

type GetUserRequest_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,oneof"`
}
type GetUserRequest_Email struct {
	Email string `protobuf:"bytes,2,opt,name=email,oneof"`
}

func (*GetUserRequest_Id) isGetUserRequest_UniqueUserOneof()    {}
func (*GetUserRequest_Email) isGetUserRequest_UniqueUserOneof() {}

func (m *GetUserRequest) GetUniqueUserOneof() isGetUserRequest_UniqueUserOneof {
	if m != nil {
		return m.UniqueUserOneof
	}
	return nil
}

func (m *GetUserRequest) GetId() int32 {
	if x, ok := m.GetUniqueUserOneof().(*GetUserRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (m *GetUserRequest) GetEmail() string {
	if x, ok := m.GetUniqueUserOneof().(*GetUserRequest_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetUserRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetUserRequest_OneofMarshaler, _GetUserRequest_OneofUnmarshaler, _GetUserRequest_OneofSizer, []interface{}{
		(*GetUserRequest_Id)(nil),
		(*GetUserRequest_Email)(nil),
	}
}

func _GetUserRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetUserRequest)
	// unique_user_oneof
	switch x := m.UniqueUserOneof.(type) {
	case *GetUserRequest_Id:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Id))
	case *GetUserRequest_Email:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Email)
	case nil:
	default:
		return fmt.Errorf("GetUserRequest.UniqueUserOneof has unexpected type %T", x)
	}
	return nil
}

func _GetUserRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetUserRequest)
	switch tag {
	case 1: // unique_user_oneof.id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.UniqueUserOneof = &GetUserRequest_Id{int32(x)}
		return true, err
	case 2: // unique_user_oneof.email
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.UniqueUserOneof = &GetUserRequest_Email{x}
		return true, err
	default:
		return false, nil
	}
}

func _GetUserRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetUserRequest)
	// unique_user_oneof
	switch x := m.UniqueUserOneof.(type) {
	case *GetUserRequest_Id:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Id))
	case *GetUserRequest_Email:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Email)))
		n += len(x.Email)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type User struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "userdata.v1.CreateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "userdata.v1.GetUserRequest")
	proto.RegisterType((*User)(nil), "userdata.v1.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserData service

type UserDataClient interface {
	// Create a new user
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Get the user by id or email
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userDataClient struct {
	cc *grpc.ClientConn
}

func NewUserDataClient(cc *grpc.ClientConn) UserDataClient {
	return &userDataClient{cc}
}

func (c *userDataClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/userdata.v1.UserData/createUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/userdata.v1.UserData/getUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserData service

type UserDataServer interface {
	// Create a new user
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// Get the user by id or email
	GetUser(context.Context, *GetUserRequest) (*User, error)
}

func RegisterUserDataServer(s *grpc.Server, srv UserDataServer) {
	s.RegisterService(&_UserData_serviceDesc, srv)
}

func _UserData_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userdata.v1.UserData/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserData_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userdata.v1.UserData/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserData_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userdata.v1.UserData",
	HandlerType: (*UserDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUser",
			Handler:    _UserData_CreateUser_Handler,
		},
		{
			MethodName: "getUser",
			Handler:    _UserData_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userdata.proto",
}

func init() { proto.RegisterFile("userdata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x93, 0xb4, 0xd5, 0x66, 0x84, 0x60, 0x46, 0x91, 0x62, 0x51, 0x24, 0x27, 0x4f, 0x01,
	0xf5, 0xe8, 0xc9, 0x3f, 0x60, 0x4f, 0x0a, 0x01, 0xcf, 0x61, 0x34, 0x53, 0x5d, 0x68, 0x12, 0xbb,
	0xd9, 0xf5, 0x43, 0xf8, 0xa9, 0x65, 0xa7, 0xb6, 0x66, 0xf1, 0xcf, 0x71, 0xe6, 0xc7, 0x7b, 0xf3,
	0x78, 0x03, 0x89, 0xed, 0x58, 0x57, 0x64, 0x28, 0x7f, 0xd3, 0xad, 0x69, 0x71, 0x67, 0x33, 0xbf,
	0x9f, 0x65, 0x0c, 0xe9, 0x8d, 0x66, 0x32, 0xfc, 0xd8, 0xb1, 0x2e, 0x78, 0x69, 0xb9, 0x33, 0x78,
	0x04, 0x30, 0x57, 0xba, 0x33, 0x65, 0x43, 0x35, 0x4f, 0xc2, 0x93, 0xf0, 0x34, 0x2e, 0x62, 0xd9,
	0xdc, 0x53, 0xcd, 0x38, 0x85, 0x78, 0x41, 0x6b, 0x1a, 0x09, 0x1d, 0xbb, 0x85, 0xc0, 0x7d, 0x18,
	0x71, 0x4d, 0x6a, 0x31, 0x19, 0x08, 0x58, 0x0d, 0xd9, 0x03, 0x24, 0x77, 0x6c, 0xfa, 0x37, 0x76,
	0x21, 0x52, 0x95, 0x78, 0x8f, 0x66, 0x41, 0x11, 0xa9, 0x0a, 0x0f, 0xd6, 0x4a, 0xb1, 0x9c, 0x05,
	0x5f, 0xda, 0xeb, 0x3d, 0x48, 0x6d, 0xa3, 0x96, 0x96, 0x4b, 0x17, 0xbc, 0x6c, 0x1b, 0x6e, 0xe7,
	0xd9, 0x2b, 0x0c, 0x9d, 0x1b, 0x26, 0xdf, 0x36, 0x62, 0xe2, 0x47, 0x8f, 0xfe, 0x8d, 0x3e, 0xf8,
	0x2b, 0xfa, 0xb0, 0x17, 0xfd, 0xfc, 0x23, 0x84, 0xb1, 0x3b, 0x75, 0x4b, 0x86, 0xf0, 0x0a, 0xe0,
	0x79, 0x53, 0x17, 0x1e, 0xe7, 0xbd, 0x2a, 0xf3, 0x1f, 0x3d, 0x1e, 0xa6, 0x1e, 0x77, 0x24, 0x0b,
	0xf0, 0x12, 0xb6, 0x5f, 0x56, 0x55, 0xe0, 0xd4, 0xe3, 0x7e, 0x41, 0xbf, 0x8a, 0x9f, 0xb6, 0xe4,
	0x85, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xeb, 0xab, 0x04, 0xd4, 0x01, 0x00, 0x00,
}

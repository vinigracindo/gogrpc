// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: user.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	AddUserStream(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserStreamClient, error)
	AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/pb.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUserStream(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], "/pb.UserService/AddUserStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceAddUserStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_AddUserStreamClient interface {
	Recv() (*UserResultStream, error)
	grpc.ClientStream
}

type userServiceAddUserStreamClient struct {
	grpc.ClientStream
}

func (x *userServiceAddUserStreamClient) Recv() (*UserResultStream, error) {
	m := new(UserResultStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userServiceClient) AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[1], "/pb.UserService/AddUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceAddUsersClient{stream}
	return x, nil
}

type UserService_AddUsersClient interface {
	Send(*User) error
	CloseAndRecv() (*Users, error)
	grpc.ClientStream
}

type userServiceAddUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceAddUsersClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userServiceAddUsersClient) CloseAndRecv() (*Users, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Users)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	AddUserStream(*User, UserService_AddUserStreamServer) error
	AddUsers(UserService_AddUsersServer) error
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) AddUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserServiceServer) AddUserStream(*User, UserService_AddUserStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AddUserStream not implemented")
}
func (UnimplementedUserServiceServer) AddUsers(UserService_AddUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method AddUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUserStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).AddUserStream(m, &userServiceAddUserStreamServer{stream})
}

type UserService_AddUserStreamServer interface {
	Send(*UserResultStream) error
	grpc.ServerStream
}

type userServiceAddUserStreamServer struct {
	grpc.ServerStream
}

func (x *userServiceAddUserStreamServer) Send(m *UserResultStream) error {
	return x.ServerStream.SendMsg(m)
}

func _UserService_AddUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServiceServer).AddUsers(&userServiceAddUsersServer{stream})
}

type UserService_AddUsersServer interface {
	SendAndClose(*Users) error
	Recv() (*User, error)
	grpc.ServerStream
}

type userServiceAddUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceAddUsersServer) SendAndClose(m *Users) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userServiceAddUsersServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddUserStream",
			Handler:       _UserService_AddUserStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddUsers",
			Handler:       _UserService_AddUsers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}

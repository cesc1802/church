// Code generated by protoc-gen-go-grpc_user. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc_user v1.2.0
// - protoc             v3.6.1
// source: module/user_v1/grpc_user/user.proto

package grpc_user

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResgisterClient is the client API for Resgister service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResgisterClient interface {
	Resgister(ctx context.Context, in *RegisterModel, opts ...grpc.CallOption) (*empty.Empty, error)
}

type resgisterClient struct {
	cc grpc.ClientConnInterface
}

func NewResgisterClient(cc grpc.ClientConnInterface) ResgisterClient {
	return &resgisterClient{cc}
}

func (c *resgisterClient) Resgister(ctx context.Context, in *RegisterModel, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Resgister/Resgister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResgisterServer is the server API for Resgister service.
// All implementations must embed UnimplementedResgisterServer
// for forward compatibility
type ResgisterServer interface {
	Resgister(context.Context, *RegisterModel) (*empty.Empty, error)
	mustEmbedUnimplementedResgisterServer()
}

// UnimplementedResgisterServer must be embedded to have forward compatible implementations.
type UnimplementedResgisterServer struct {
}

func (UnimplementedResgisterServer) Resgister(context.Context, *RegisterModel) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resgister not implemented")
}
func (UnimplementedResgisterServer) mustEmbedUnimplementedResgisterServer() {}

// UnsafeResgisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResgisterServer will
// result in compilation errors.
type UnsafeResgisterServer interface {
	mustEmbedUnimplementedResgisterServer()
}

func RegisterResgisterServer(s grpc.ServiceRegistrar, srv ResgisterServer) {
	s.RegisterService(&Resgister_ServiceDesc, srv)
}

func _Resgister_Resgister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResgisterServer).Resgister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Resgister/Resgister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResgisterServer).Resgister(ctx, req.(*RegisterModel))
	}
	return interceptor(ctx, in, info, handler)
}

// Resgister_ServiceDesc is the grpc_user.ServiceDesc for Resgister service.
// It's only intended for direct use with grpc_user.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resgister_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Resgister",
	HandlerType: (*ResgisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resgister",
			Handler:    _Resgister_Resgister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/user_v1/grpc_user/user.proto",
}

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	Login(ctx context.Context, in *LoginModel, opts ...grpc.CallOption) (*JWT, error)
	CheckAuthen(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*empty.Empty, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginModel, opts ...grpc.CallOption) (*JWT, error) {
	out := new(JWT)
	err := c.cc.Invoke(ctx, "/authentication.Login/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) CheckAuthen(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/authentication.Login/CheckAuthen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	Login(context.Context, *LoginModel) (*JWT, error)
	CheckAuthen(context.Context, *JWT) (*empty.Empty, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) Login(context.Context, *LoginModel) (*JWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServer) CheckAuthen(context.Context, *JWT) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuthen not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_CheckAuthen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).CheckAuthen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.Login/CheckAuthen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).CheckAuthen(ctx, req.(*JWT))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc_user.ServiceDesc for Login service.
// It's only intended for direct use with grpc_user.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
		{
			MethodName: "CheckAuthen",
			Handler:    _Login_CheckAuthen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/user_v1/grpc_user/user.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Auth_Login_FullMethodName    = "/auth.Auth/Login"
	Auth_Upsert_FullMethodName   = "/auth.Auth/Upsert"
	Auth_Validate_FullMethodName = "/auth.Auth/Validate"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Login(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error)
	Upsert(ctx context.Context, in *AuthUpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Validate(ctx context.Context, in *AuthValidateRequest, opts ...grpc.CallOption) (*AuthValidateResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Login(ctx context.Context, in *AuthLoginRequest, opts ...grpc.CallOption) (*AuthLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthLoginResponse)
	err := c.cc.Invoke(ctx, Auth_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Upsert(ctx context.Context, in *AuthUpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Auth_Upsert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Validate(ctx context.Context, in *AuthValidateRequest, opts ...grpc.CallOption) (*AuthValidateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthValidateResponse)
	err := c.cc.Invoke(ctx, Auth_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	Login(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error)
	Upsert(context.Context, *AuthUpsertRequest) (*emptypb.Empty, error)
	Validate(context.Context, *AuthValidateRequest) (*AuthValidateResponse, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) Login(context.Context, *AuthLoginRequest) (*AuthLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServer) Upsert(context.Context, *AuthUpsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedAuthServer) Validate(context.Context, *AuthValidateRequest) (*AuthValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}
func (UnimplementedAuthServer) testEmbeddedByValue()              {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Login(ctx, req.(*AuthLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Upsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Upsert(ctx, req.(*AuthUpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Validate(ctx, req.(*AuthValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Auth_Login_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _Auth_Upsert_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _Auth_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/auth.proto",
}

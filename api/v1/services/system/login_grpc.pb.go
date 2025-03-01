// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: system/login.proto

package system

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LoginService_Captcha_FullMethodName      = "/api.v1.services.system.LoginService/Captcha"
	LoginService_CaptchaId_FullMethodName    = "/api.v1.services.system.LoginService/CaptchaId"
	LoginService_CaptchaImage_FullMethodName = "/api.v1.services.system.LoginService/CaptchaImage"
	LoginService_CaptchaAudio_FullMethodName = "/api.v1.services.system.LoginService/CaptchaAudio"
	LoginService_Login_FullMethodName        = "/api.v1.services.system.LoginService/Login"
	LoginService_Logout_FullMethodName       = "/api.v1.services.system.LoginService/Logout"
	LoginService_Register_FullMethodName     = "/api.v1.services.system.LoginService/Register"
	LoginService_TokenRefresh_FullMethodName = "/api.v1.services.system.LoginService/TokenRefresh"
)

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The login service definition.
type LoginServiceClient interface {
	Captcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error)
	CaptchaId(ctx context.Context, in *CaptchaIdRequest, opts ...grpc.CallOption) (*CaptchaIdResponse, error)
	CaptchaImage(ctx context.Context, in *CaptchaImageRequest, opts ...grpc.CallOption) (*CaptchaImageResponse, error)
	CaptchaAudio(ctx context.Context, in *CaptchaAudioRequest, opts ...grpc.CallOption) (*CaptchaAudioResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	TokenRefresh(ctx context.Context, in *TokenRefreshRequest, opts ...grpc.CallOption) (*TokenRefreshResponse, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Captcha(ctx context.Context, in *CaptchaRequest, opts ...grpc.CallOption) (*CaptchaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaptchaResponse)
	err := c.cc.Invoke(ctx, LoginService_Captcha_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) CaptchaId(ctx context.Context, in *CaptchaIdRequest, opts ...grpc.CallOption) (*CaptchaIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaptchaIdResponse)
	err := c.cc.Invoke(ctx, LoginService_CaptchaId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) CaptchaImage(ctx context.Context, in *CaptchaImageRequest, opts ...grpc.CallOption) (*CaptchaImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaptchaImageResponse)
	err := c.cc.Invoke(ctx, LoginService_CaptchaImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) CaptchaAudio(ctx context.Context, in *CaptchaAudioRequest, opts ...grpc.CallOption) (*CaptchaAudioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaptchaAudioResponse)
	err := c.cc.Invoke(ctx, LoginService_CaptchaAudio_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, LoginService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, LoginService_Logout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, LoginService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginServiceClient) TokenRefresh(ctx context.Context, in *TokenRefreshRequest, opts ...grpc.CallOption) (*TokenRefreshResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenRefreshResponse)
	err := c.cc.Invoke(ctx, LoginService_TokenRefresh_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility.
//
// The login service definition.
type LoginServiceServer interface {
	Captcha(context.Context, *CaptchaRequest) (*CaptchaResponse, error)
	CaptchaId(context.Context, *CaptchaIdRequest) (*CaptchaIdResponse, error)
	CaptchaImage(context.Context, *CaptchaImageRequest) (*CaptchaImageResponse, error)
	CaptchaAudio(context.Context, *CaptchaAudioRequest) (*CaptchaAudioResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	TokenRefresh(context.Context, *TokenRefreshRequest) (*TokenRefreshResponse, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLoginServiceServer struct{}

func (UnimplementedLoginServiceServer) Captcha(context.Context, *CaptchaRequest) (*CaptchaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Captcha not implemented")
}
func (UnimplementedLoginServiceServer) CaptchaId(context.Context, *CaptchaIdRequest) (*CaptchaIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaId not implemented")
}
func (UnimplementedLoginServiceServer) CaptchaImage(context.Context, *CaptchaImageRequest) (*CaptchaImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaImage not implemented")
}
func (UnimplementedLoginServiceServer) CaptchaAudio(context.Context, *CaptchaAudioRequest) (*CaptchaAudioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CaptchaAudio not implemented")
}
func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedLoginServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedLoginServiceServer) TokenRefresh(context.Context, *TokenRefreshRequest) (*TokenRefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenRefresh not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}
func (UnimplementedLoginServiceServer) testEmbeddedByValue()                      {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	// If the following call pancis, it indicates UnimplementedLoginServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Captcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Captcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_Captcha_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Captcha(ctx, req.(*CaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_CaptchaId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CaptchaId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_CaptchaId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CaptchaId(ctx, req.(*CaptchaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_CaptchaImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CaptchaImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_CaptchaImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CaptchaImage(ctx, req.(*CaptchaImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_CaptchaAudio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CaptchaAudioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).CaptchaAudio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_CaptchaAudio_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).CaptchaAudio(ctx, req.(*CaptchaAudioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoginService_TokenRefresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).TokenRefresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LoginService_TokenRefresh_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).TokenRefresh(ctx, req.(*TokenRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.services.system.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Captcha",
			Handler:    _LoginService_Captcha_Handler,
		},
		{
			MethodName: "CaptchaId",
			Handler:    _LoginService_CaptchaId_Handler,
		},
		{
			MethodName: "CaptchaImage",
			Handler:    _LoginService_CaptchaImage_Handler,
		},
		{
			MethodName: "CaptchaAudio",
			Handler:    _LoginService_CaptchaAudio_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _LoginService_Logout_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _LoginService_Register_Handler,
		},
		{
			MethodName: "TokenRefresh",
			Handler:    _LoginService_TokenRefresh_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/login.proto",
}

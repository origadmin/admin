// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: system/current.proto

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
	CurrentService_CurrentLogout_FullMethodName             = "/api.v1.services.system.CurrentService/CurrentLogout"
	CurrentService_UpdateCurrentUserPassword_FullMethodName = "/api.v1.services.system.CurrentService/UpdateCurrentUserPassword"
	CurrentService_UpdateCurrentUser_FullMethodName         = "/api.v1.services.system.CurrentService/UpdateCurrentUser"
	CurrentService_GetCurrentUser_FullMethodName            = "/api.v1.services.system.CurrentService/GetCurrentUser"
	CurrentService_ListCurrentResources_FullMethodName      = "/api.v1.services.system.CurrentService/ListCurrentResources"
	CurrentService_ListCurrentRoles_FullMethodName          = "/api.v1.services.system.CurrentService/ListCurrentRoles"
	CurrentService_UpdateCurrentSetting_FullMethodName      = "/api.v1.services.system.CurrentService/UpdateCurrentSetting"
	CurrentService_RefreshCurrentToken_FullMethodName       = "/api.v1.services.system.CurrentService/RefreshCurrentToken"
)

// CurrentServiceClient is the client API for CurrentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The current service definition.
type CurrentServiceClient interface {
	CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...grpc.CallOption) (*CurrentLogoutResponse, error)
	// UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordResponse, error)
	// UpdateCurrentUser Update the current user information
	UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserResponse, error)
	// GetCurrentUser Update the current user information
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error)
	// ListCurrentResources List the current user's menu
	ListCurrentResources(ctx context.Context, in *ListCurrentResourcesRequest, opts ...grpc.CallOption) (*ListCurrentResourcesResponse, error)
	// ListCurrentResources List the current user's menu
	ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...grpc.CallOption) (*ListCurrentRolesResponse, error)
	// UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentSettingResponse, error)
	// RefreshCurrentToken Refresh the current user's token
	RefreshCurrentToken(ctx context.Context, in *RefreshCurrentTokenRequest, opts ...grpc.CallOption) (*RefreshCurrentTokenResponse, error)
}

type currentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrentServiceClient(cc grpc.ClientConnInterface) CurrentServiceClient {
	return &currentServiceClient{cc}
}

func (c *currentServiceClient) CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...grpc.CallOption) (*CurrentLogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrentLogoutResponse)
	err := c.cc.Invoke(ctx, CurrentService_CurrentLogout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserPasswordResponse)
	err := c.cc.Invoke(ctx, CurrentService_UpdateCurrentUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserResponse)
	err := c.cc.Invoke(ctx, CurrentService_UpdateCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentUserResponse)
	err := c.cc.Invoke(ctx, CurrentService_GetCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) ListCurrentResources(ctx context.Context, in *ListCurrentResourcesRequest, opts ...grpc.CallOption) (*ListCurrentResourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentResourcesResponse)
	err := c.cc.Invoke(ctx, CurrentService_ListCurrentResources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...grpc.CallOption) (*ListCurrentRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentRolesResponse)
	err := c.cc.Invoke(ctx, CurrentService_ListCurrentRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentSettingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentSettingResponse)
	err := c.cc.Invoke(ctx, CurrentService_UpdateCurrentSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentServiceClient) RefreshCurrentToken(ctx context.Context, in *RefreshCurrentTokenRequest, opts ...grpc.CallOption) (*RefreshCurrentTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshCurrentTokenResponse)
	err := c.cc.Invoke(ctx, CurrentService_RefreshCurrentToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentServiceServer is the server API for CurrentService service.
// All implementations must embed UnimplementedCurrentServiceServer
// for forward compatibility.
//
// The current service definition.
type CurrentServiceServer interface {
	CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error)
	// UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error)
	// UpdateCurrentUser Update the current user information
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error)
	// GetCurrentUser Update the current user information
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error)
	// ListCurrentResources List the current user's menu
	ListCurrentResources(context.Context, *ListCurrentResourcesRequest) (*ListCurrentResourcesResponse, error)
	// ListCurrentResources List the current user's menu
	ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error)
	// UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error)
	// RefreshCurrentToken Refresh the current user's token
	RefreshCurrentToken(context.Context, *RefreshCurrentTokenRequest) (*RefreshCurrentTokenResponse, error)
	mustEmbedUnimplementedCurrentServiceServer()
}

// UnimplementedCurrentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCurrentServiceServer struct{}

func (UnimplementedCurrentServiceServer) CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentLogout not implemented")
}
func (UnimplementedCurrentServiceServer) UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUserPassword not implemented")
}
func (UnimplementedCurrentServiceServer) UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUser not implemented")
}
func (UnimplementedCurrentServiceServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedCurrentServiceServer) ListCurrentResources(context.Context, *ListCurrentResourcesRequest) (*ListCurrentResourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentResources not implemented")
}
func (UnimplementedCurrentServiceServer) ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentRoles not implemented")
}
func (UnimplementedCurrentServiceServer) UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentSetting not implemented")
}
func (UnimplementedCurrentServiceServer) RefreshCurrentToken(context.Context, *RefreshCurrentTokenRequest) (*RefreshCurrentTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCurrentToken not implemented")
}
func (UnimplementedCurrentServiceServer) mustEmbedUnimplementedCurrentServiceServer() {}
func (UnimplementedCurrentServiceServer) testEmbeddedByValue()                        {}

// UnsafeCurrentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrentServiceServer will
// result in compilation errors.
type UnsafeCurrentServiceServer interface {
	mustEmbedUnimplementedCurrentServiceServer()
}

func RegisterCurrentServiceServer(s grpc.ServiceRegistrar, srv CurrentServiceServer) {
	// If the following call pancis, it indicates UnimplementedCurrentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CurrentService_ServiceDesc, srv)
}

func _CurrentService_CurrentLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).CurrentLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_CurrentLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).CurrentLogout(ctx, req.(*CurrentLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_UpdateCurrentUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).UpdateCurrentUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_UpdateCurrentUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).UpdateCurrentUserPassword(ctx, req.(*UpdateCurrentUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_UpdateCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).UpdateCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_UpdateCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_ListCurrentResources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrentResourcesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).ListCurrentResources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_ListCurrentResources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).ListCurrentResources(ctx, req.(*ListCurrentResourcesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_ListCurrentRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrentRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).ListCurrentRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_ListCurrentRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).ListCurrentRoles(ctx, req.(*ListCurrentRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_UpdateCurrentSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).UpdateCurrentSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_UpdateCurrentSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).UpdateCurrentSetting(ctx, req.(*UpdateCurrentSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentService_RefreshCurrentToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshCurrentTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServiceServer).RefreshCurrentToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentService_RefreshCurrentToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServiceServer).RefreshCurrentToken(ctx, req.(*RefreshCurrentTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrentService_ServiceDesc is the grpc.ServiceDesc for CurrentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.services.system.CurrentService",
	HandlerType: (*CurrentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentLogout",
			Handler:    _CurrentService_CurrentLogout_Handler,
		},
		{
			MethodName: "UpdateCurrentUserPassword",
			Handler:    _CurrentService_UpdateCurrentUserPassword_Handler,
		},
		{
			MethodName: "UpdateCurrentUser",
			Handler:    _CurrentService_UpdateCurrentUser_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _CurrentService_GetCurrentUser_Handler,
		},
		{
			MethodName: "ListCurrentResources",
			Handler:    _CurrentService_ListCurrentResources_Handler,
		},
		{
			MethodName: "ListCurrentRoles",
			Handler:    _CurrentService_ListCurrentRoles_Handler,
		},
		{
			MethodName: "UpdateCurrentSetting",
			Handler:    _CurrentService_UpdateCurrentSetting_Handler,
		},
		{
			MethodName: "RefreshCurrentToken",
			Handler:    _CurrentService_RefreshCurrentToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/current.proto",
}

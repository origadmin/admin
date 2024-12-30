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
	CurrentAPI_CurrentLogout_FullMethodName             = "/api.v1.services.system.CurrentAPI/CurrentLogout"
	CurrentAPI_UpdateCurrentUserPassword_FullMethodName = "/api.v1.services.system.CurrentAPI/UpdateCurrentUserPassword"
	CurrentAPI_UpdateCurrentUser_FullMethodName         = "/api.v1.services.system.CurrentAPI/UpdateCurrentUser"
	CurrentAPI_GetCurrentUser_FullMethodName            = "/api.v1.services.system.CurrentAPI/GetCurrentUser"
	CurrentAPI_ListCurrentMenus_FullMethodName          = "/api.v1.services.system.CurrentAPI/ListCurrentMenus"
	CurrentAPI_ListCurrentRoles_FullMethodName          = "/api.v1.services.system.CurrentAPI/ListCurrentRoles"
	CurrentAPI_UpdateCurrentSetting_FullMethodName      = "/api.v1.services.system.CurrentAPI/UpdateCurrentSetting"
)

// CurrentAPIClient is the client API for CurrentAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The current service definition.
type CurrentAPIClient interface {
	CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...grpc.CallOption) (*CurrentLogoutResponse, error)
	// UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordResponse, error)
	// UpdateCurrentUser Update the current user information
	UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserResponse, error)
	// GetCurrentUser Update the current user information
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error)
	// ListCurrentMenus List the current user's menu
	ListCurrentMenus(ctx context.Context, in *ListCurrentMenusRequest, opts ...grpc.CallOption) (*ListCurrentMenusResponse, error)
	// ListCurrentMenus List the current user's menu
	ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...grpc.CallOption) (*ListCurrentRolesResponse, error)
	// UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentSettingResponse, error)
}

type currentAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrentAPIClient(cc grpc.ClientConnInterface) CurrentAPIClient {
	return &currentAPIClient{cc}
}

func (c *currentAPIClient) CurrentLogout(ctx context.Context, in *CurrentLogoutRequest, opts ...grpc.CallOption) (*CurrentLogoutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CurrentLogoutResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_CurrentLogout_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) UpdateCurrentUserPassword(ctx context.Context, in *UpdateCurrentUserPasswordRequest, opts ...grpc.CallOption) (*UpdateCurrentUserPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserPasswordResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_UpdateCurrentUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) UpdateCurrentUser(ctx context.Context, in *UpdateCurrentUserRequest, opts ...grpc.CallOption) (*UpdateCurrentUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentUserResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_UpdateCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCurrentUserResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_GetCurrentUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) ListCurrentMenus(ctx context.Context, in *ListCurrentMenusRequest, opts ...grpc.CallOption) (*ListCurrentMenusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentMenusResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_ListCurrentMenus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) ListCurrentRoles(ctx context.Context, in *ListCurrentRolesRequest, opts ...grpc.CallOption) (*ListCurrentRolesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCurrentRolesResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_ListCurrentRoles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentAPIClient) UpdateCurrentSetting(ctx context.Context, in *UpdateCurrentSettingRequest, opts ...grpc.CallOption) (*UpdateCurrentSettingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCurrentSettingResponse)
	err := c.cc.Invoke(ctx, CurrentAPI_UpdateCurrentSetting_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentAPIServer is the server API for CurrentAPI service.
// All implementations must embed UnimplementedCurrentAPIServer
// for forward compatibility.
//
// The current service definition.
type CurrentAPIServer interface {
	CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error)
	// UpdateCurrentUserPassword The user changes the password
	UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error)
	// UpdateCurrentUser Update the current user information
	UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error)
	// GetCurrentUser Update the current user information
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error)
	// ListCurrentMenus List the current user's menu
	ListCurrentMenus(context.Context, *ListCurrentMenusRequest) (*ListCurrentMenusResponse, error)
	// ListCurrentMenus List the current user's menu
	ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error)
	// UpdateCurrentSetting User settings are saved
	UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error)
	mustEmbedUnimplementedCurrentAPIServer()
}

// UnimplementedCurrentAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCurrentAPIServer struct{}

func (UnimplementedCurrentAPIServer) CurrentLogout(context.Context, *CurrentLogoutRequest) (*CurrentLogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentLogout not implemented")
}
func (UnimplementedCurrentAPIServer) UpdateCurrentUserPassword(context.Context, *UpdateCurrentUserPasswordRequest) (*UpdateCurrentUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUserPassword not implemented")
}
func (UnimplementedCurrentAPIServer) UpdateCurrentUser(context.Context, *UpdateCurrentUserRequest) (*UpdateCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentUser not implemented")
}
func (UnimplementedCurrentAPIServer) GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedCurrentAPIServer) ListCurrentMenus(context.Context, *ListCurrentMenusRequest) (*ListCurrentMenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentMenus not implemented")
}
func (UnimplementedCurrentAPIServer) ListCurrentRoles(context.Context, *ListCurrentRolesRequest) (*ListCurrentRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCurrentRoles not implemented")
}
func (UnimplementedCurrentAPIServer) UpdateCurrentSetting(context.Context, *UpdateCurrentSettingRequest) (*UpdateCurrentSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCurrentSetting not implemented")
}
func (UnimplementedCurrentAPIServer) mustEmbedUnimplementedCurrentAPIServer() {}
func (UnimplementedCurrentAPIServer) testEmbeddedByValue()                    {}

// UnsafeCurrentAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrentAPIServer will
// result in compilation errors.
type UnsafeCurrentAPIServer interface {
	mustEmbedUnimplementedCurrentAPIServer()
}

func RegisterCurrentAPIServer(s grpc.ServiceRegistrar, srv CurrentAPIServer) {
	// If the following call pancis, it indicates UnimplementedCurrentAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CurrentAPI_ServiceDesc, srv)
}

func _CurrentAPI_CurrentLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).CurrentLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_CurrentLogout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).CurrentLogout(ctx, req.(*CurrentLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_UpdateCurrentUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).UpdateCurrentUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_UpdateCurrentUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).UpdateCurrentUserPassword(ctx, req.(*UpdateCurrentUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_UpdateCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).UpdateCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_UpdateCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).UpdateCurrentUser(ctx, req.(*UpdateCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_GetCurrentUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_ListCurrentMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrentMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).ListCurrentMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_ListCurrentMenus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).ListCurrentMenus(ctx, req.(*ListCurrentMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_ListCurrentRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrentRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).ListCurrentRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_ListCurrentRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).ListCurrentRoles(ctx, req.(*ListCurrentRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrentAPI_UpdateCurrentSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrentSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentAPIServer).UpdateCurrentSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrentAPI_UpdateCurrentSetting_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentAPIServer).UpdateCurrentSetting(ctx, req.(*UpdateCurrentSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrentAPI_ServiceDesc is the grpc.ServiceDesc for CurrentAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrentAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.services.system.CurrentAPI",
	HandlerType: (*CurrentAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentLogout",
			Handler:    _CurrentAPI_CurrentLogout_Handler,
		},
		{
			MethodName: "UpdateCurrentUserPassword",
			Handler:    _CurrentAPI_UpdateCurrentUserPassword_Handler,
		},
		{
			MethodName: "UpdateCurrentUser",
			Handler:    _CurrentAPI_UpdateCurrentUser_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _CurrentAPI_GetCurrentUser_Handler,
		},
		{
			MethodName: "ListCurrentMenus",
			Handler:    _CurrentAPI_ListCurrentMenus_Handler,
		},
		{
			MethodName: "ListCurrentRoles",
			Handler:    _CurrentAPI_ListCurrentRoles_Handler,
		},
		{
			MethodName: "UpdateCurrentSetting",
			Handler:    _CurrentAPI_UpdateCurrentSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/current.proto",
}
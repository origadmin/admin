// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: basis/current.proto

package basis

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/origadmin/runtime/gen/go/security/jwt/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateCurrentSettingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateCurrentSettingRequest) Reset() {
	*x = UpdateCurrentSettingRequest{}
	mi := &file_basis_current_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentSettingRequest) ProtoMessage() {}

func (x *UpdateCurrentSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentSettingRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentSettingRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCurrentSettingRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCurrentSettingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCurrentSettingResponse) Reset() {
	*x = UpdateCurrentSettingResponse{}
	mi := &file_basis_current_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentSettingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentSettingResponse) ProtoMessage() {}

func (x *UpdateCurrentSettingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentSettingResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrentSettingResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{1}
}

type UpdateCurrentRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateCurrentRolesRequest) Reset() {
	*x = UpdateCurrentRolesRequest{}
	mi := &file_basis_current_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentRolesRequest) ProtoMessage() {}

func (x *UpdateCurrentRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentRolesRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentRolesRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCurrentRolesRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCurrentRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCurrentRolesResponse) Reset() {
	*x = UpdateCurrentRolesResponse{}
	mi := &file_basis_current_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentRolesResponse) ProtoMessage() {}

func (x *UpdateCurrentRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentRolesResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrentRolesResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{3}
}

type ListCurrentMenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ListCurrentMenusRequest) Reset() {
	*x = ListCurrentMenusRequest{}
	mi := &file_basis_current_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrentMenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrentMenusRequest) ProtoMessage() {}

func (x *ListCurrentMenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrentMenusRequest.ProtoReflect.Descriptor instead.
func (*ListCurrentMenusRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{4}
}

func (x *ListCurrentMenusRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListCurrentMenusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCurrentMenusResponse) Reset() {
	*x = ListCurrentMenusResponse{}
	mi := &file_basis_current_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCurrentMenusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCurrentMenusResponse) ProtoMessage() {}

func (x *ListCurrentMenusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCurrentMenusResponse.ProtoReflect.Descriptor instead.
func (*ListCurrentMenusResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{5}
}

type UpdateCurrentUserPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateCurrentUserPasswordRequest) Reset() {
	*x = UpdateCurrentUserPasswordRequest{}
	mi := &file_basis_current_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentUserPasswordRequest) ProtoMessage() {}

func (x *UpdateCurrentUserPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCurrentUserPasswordRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCurrentUserPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCurrentUserPasswordResponse) Reset() {
	*x = UpdateCurrentUserPasswordResponse{}
	mi := &file_basis_current_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentUserPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentUserPasswordResponse) ProtoMessage() {}

func (x *UpdateCurrentUserPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentUserPasswordResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrentUserPasswordResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{7}
}

type CurrentPasswordRestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CurrentPasswordRestRequest) Reset() {
	*x = CurrentPasswordRestRequest{}
	mi := &file_basis_current_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentPasswordRestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentPasswordRestRequest) ProtoMessage() {}

func (x *CurrentPasswordRestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentPasswordRestRequest.ProtoReflect.Descriptor instead.
func (*CurrentPasswordRestRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{8}
}

func (x *CurrentPasswordRestRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CurrentPasswordRestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CurrentPasswordRestResponse) Reset() {
	*x = CurrentPasswordRestResponse{}
	mi := &file_basis_current_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentPasswordRestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentPasswordRestResponse) ProtoMessage() {}

func (x *CurrentPasswordRestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentPasswordRestResponse.ProtoReflect.Descriptor instead.
func (*CurrentPasswordRestResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{9}
}

type UpdateCurrentUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateCurrentUserRequest) Reset() {
	*x = UpdateCurrentUserRequest{}
	mi := &file_basis_current_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentUserRequest) ProtoMessage() {}

func (x *UpdateCurrentUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateCurrentUserRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateCurrentUserRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpdateCurrentUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCurrentUserResponse) Reset() {
	*x = UpdateCurrentUserResponse{}
	mi := &file_basis_current_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCurrentUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCurrentUserResponse) ProtoMessage() {}

func (x *UpdateCurrentUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCurrentUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateCurrentUserResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{11}
}

type CurrentLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CurrentLogoutRequest) Reset() {
	*x = CurrentLogoutRequest{}
	mi := &file_basis_current_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentLogoutRequest) ProtoMessage() {}

func (x *CurrentLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentLogoutRequest.ProtoReflect.Descriptor instead.
func (*CurrentLogoutRequest) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{12}
}

func (x *CurrentLogoutRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type CurrentLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CurrentLogoutResponse) Reset() {
	*x = CurrentLogoutResponse{}
	mi := &file_basis_current_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CurrentLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentLogoutResponse) ProtoMessage() {}

func (x *CurrentLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_basis_current_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentLogoutResponse.ProtoReflect.Descriptor instead.
func (*CurrentLogoutResponse) Descriptor() ([]byte, []int) {
	return file_basis_current_proto_rawDescGZIP(), []int{13}
}

func (x *CurrentLogoutResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_basis_current_proto protoreflect.FileDescriptor

var file_basis_current_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x6a, 0x77, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x1b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1c, 0x0a, 0x1a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1a,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x0a, 0x20, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x23, 0x0a, 0x21, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x0a,
	0x1a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xaf, 0x07, 0x0a, 0x0a,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x12, 0x89, 0x01, 0x0a, 0x0d, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x2b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x73, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f,
	0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0xb4, 0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x93, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x0d, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x8b, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x73, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x42, 0xba, 0x01,
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x42, 0x0c, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x17, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x73, 0x3b, 0x62,
	0x61, 0x73, 0x69, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x53, 0x42, 0xaa, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x61,
	0x73, 0x69, 0x73, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x69, 0x73, 0xe2, 0x02, 0x21, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x61,
	0x73, 0x69, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x18, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_basis_current_proto_rawDescOnce sync.Once
	file_basis_current_proto_rawDescData = file_basis_current_proto_rawDesc
)

func file_basis_current_proto_rawDescGZIP() []byte {
	file_basis_current_proto_rawDescOnce.Do(func() {
		file_basis_current_proto_rawDescData = protoimpl.X.CompressGZIP(file_basis_current_proto_rawDescData)
	})
	return file_basis_current_proto_rawDescData
}

var file_basis_current_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_basis_current_proto_goTypes = []any{
	(*UpdateCurrentSettingRequest)(nil),       // 0: api.v1.services.basis.UpdateCurrentSettingRequest
	(*UpdateCurrentSettingResponse)(nil),      // 1: api.v1.services.basis.UpdateCurrentSettingResponse
	(*UpdateCurrentRolesRequest)(nil),         // 2: api.v1.services.basis.UpdateCurrentRolesRequest
	(*UpdateCurrentRolesResponse)(nil),        // 3: api.v1.services.basis.UpdateCurrentRolesResponse
	(*ListCurrentMenusRequest)(nil),           // 4: api.v1.services.basis.ListCurrentMenusRequest
	(*ListCurrentMenusResponse)(nil),          // 5: api.v1.services.basis.ListCurrentMenusResponse
	(*UpdateCurrentUserPasswordRequest)(nil),  // 6: api.v1.services.basis.UpdateCurrentUserPasswordRequest
	(*UpdateCurrentUserPasswordResponse)(nil), // 7: api.v1.services.basis.UpdateCurrentUserPasswordResponse
	(*CurrentPasswordRestRequest)(nil),        // 8: api.v1.services.basis.CurrentPasswordRestRequest
	(*CurrentPasswordRestResponse)(nil),       // 9: api.v1.services.basis.CurrentPasswordRestResponse
	(*UpdateCurrentUserRequest)(nil),          // 10: api.v1.services.basis.UpdateCurrentUserRequest
	(*UpdateCurrentUserResponse)(nil),         // 11: api.v1.services.basis.UpdateCurrentUserResponse
	(*CurrentLogoutRequest)(nil),              // 12: api.v1.services.basis.CurrentLogoutRequest
	(*CurrentLogoutResponse)(nil),             // 13: api.v1.services.basis.CurrentLogoutResponse
	(*anypb.Any)(nil),                         // 14: google.protobuf.Any
}
var file_basis_current_proto_depIdxs = []int32{
	14, // 0: api.v1.services.basis.UpdateCurrentSettingRequest.data:type_name -> google.protobuf.Any
	14, // 1: api.v1.services.basis.UpdateCurrentRolesRequest.data:type_name -> google.protobuf.Any
	14, // 2: api.v1.services.basis.ListCurrentMenusRequest.data:type_name -> google.protobuf.Any
	14, // 3: api.v1.services.basis.UpdateCurrentUserPasswordRequest.data:type_name -> google.protobuf.Any
	14, // 4: api.v1.services.basis.UpdateCurrentUserRequest.data:type_name -> google.protobuf.Any
	14, // 5: api.v1.services.basis.CurrentLogoutRequest.data:type_name -> google.protobuf.Any
	12, // 6: api.v1.services.basis.CurrentAPI.CurrentLogout:input_type -> api.v1.services.basis.CurrentLogoutRequest
	6,  // 7: api.v1.services.basis.CurrentAPI.UpdateCurrentUserPassword:input_type -> api.v1.services.basis.UpdateCurrentUserPasswordRequest
	10, // 8: api.v1.services.basis.CurrentAPI.UpdateCurrentUser:input_type -> api.v1.services.basis.UpdateCurrentUserRequest
	4,  // 9: api.v1.services.basis.CurrentAPI.ListCurrentMenus:input_type -> api.v1.services.basis.ListCurrentMenusRequest
	2,  // 10: api.v1.services.basis.CurrentAPI.UpdateCurrentRoles:input_type -> api.v1.services.basis.UpdateCurrentRolesRequest
	0,  // 11: api.v1.services.basis.CurrentAPI.UpdateCurrentSetting:input_type -> api.v1.services.basis.UpdateCurrentSettingRequest
	13, // 12: api.v1.services.basis.CurrentAPI.CurrentLogout:output_type -> api.v1.services.basis.CurrentLogoutResponse
	7,  // 13: api.v1.services.basis.CurrentAPI.UpdateCurrentUserPassword:output_type -> api.v1.services.basis.UpdateCurrentUserPasswordResponse
	11, // 14: api.v1.services.basis.CurrentAPI.UpdateCurrentUser:output_type -> api.v1.services.basis.UpdateCurrentUserResponse
	5,  // 15: api.v1.services.basis.CurrentAPI.ListCurrentMenus:output_type -> api.v1.services.basis.ListCurrentMenusResponse
	3,  // 16: api.v1.services.basis.CurrentAPI.UpdateCurrentRoles:output_type -> api.v1.services.basis.UpdateCurrentRolesResponse
	1,  // 17: api.v1.services.basis.CurrentAPI.UpdateCurrentSetting:output_type -> api.v1.services.basis.UpdateCurrentSettingResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_basis_current_proto_init() }
func file_basis_current_proto_init() {
	if File_basis_current_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basis_current_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_basis_current_proto_goTypes,
		DependencyIndexes: file_basis_current_proto_depIdxs,
		MessageInfos:      file_basis_current_proto_msgTypes,
	}.Build()
	File_basis_current_proto = out.File
	file_basis_current_proto_rawDesc = nil
	file_basis_current_proto_goTypes = nil
	file_basis_current_proto_depIdxs = nil
}
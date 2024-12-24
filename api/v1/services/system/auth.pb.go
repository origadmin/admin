// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: system/auth.proto

package system

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAuthResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of Auths to return.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The next_page_token value returned from a previous List request, if any.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// The current page number.
	Current int32 `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`
	// The no_paging is used to disable pagination.
	NoPaging bool `protobuf:"varint,4,opt,name=no_paging,json=noPaging,proto3" json:"no_paging,omitempty"`
}

func (x *ListAuthResourcesRequest) Reset() {
	*x = ListAuthResourcesRequest{}
	mi := &file_system_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthResourcesRequest) ProtoMessage() {}

func (x *ListAuthResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthResourcesRequest.ProtoReflect.Descriptor instead.
func (*ListAuthResourcesRequest) Descriptor() ([]byte, []int) {
	return file_system_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ListAuthResourcesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAuthResourcesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListAuthResourcesRequest) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *ListAuthResourcesRequest) GetNoPaging() bool {
	if x != nil {
		return x.NoPaging
	}
	return false
}

type ListAuthResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of Auths.
	Resources []*Resource `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	// The total number of Auths in the result set.
	TotalSize int32 `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
}

func (x *ListAuthResourcesResponse) Reset() {
	*x = ListAuthResourcesResponse{}
	mi := &file_system_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAuthResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuthResourcesResponse) ProtoMessage() {}

func (x *ListAuthResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuthResourcesResponse.ProtoReflect.Descriptor instead.
func (*ListAuthResourcesResponse) Descriptor() ([]byte, []int) {
	return file_system_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ListAuthResourcesResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *ListAuthResourcesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

var File_system_auth_proto protoreflect.FileDescriptor

var file_system_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x6f, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x7a, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a,
	0x65, 0x32, 0xa1, 0x01, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x68, 0x41, 0x50, 0x49, 0x12, 0x95, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x12, 0x13, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0xbe, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x19, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xa2, 0x02, 0x04, 0x41,
	0x56, 0x53, 0x53, 0xaa, 0x02, 0x16, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xca, 0x02, 0x16, 0x41,
	0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0xe2, 0x02, 0x22, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69,
	0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_auth_proto_rawDescOnce sync.Once
	file_system_auth_proto_rawDescData = file_system_auth_proto_rawDesc
)

func file_system_auth_proto_rawDescGZIP() []byte {
	file_system_auth_proto_rawDescOnce.Do(func() {
		file_system_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_auth_proto_rawDescData)
	})
	return file_system_auth_proto_rawDescData
}

var file_system_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_system_auth_proto_goTypes = []any{
	(*ListAuthResourcesRequest)(nil),  // 0: api.v1.services.system.ListAuthResourcesRequest
	(*ListAuthResourcesResponse)(nil), // 1: api.v1.services.system.ListAuthResourcesResponse
	(*Resource)(nil),                  // 2: api.v1.services.system.Resource
}
var file_system_auth_proto_depIdxs = []int32{
	2, // 0: api.v1.services.system.ListAuthResourcesResponse.resources:type_name -> api.v1.services.system.Resource
	0, // 1: api.v1.services.system.AuthAPI.ListAuthResources:input_type -> api.v1.services.system.ListAuthResourcesRequest
	1, // 2: api.v1.services.system.AuthAPI.ListAuthResources:output_type -> api.v1.services.system.ListAuthResourcesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_auth_proto_init() }
func file_system_auth_proto_init() {
	if File_system_auth_proto != nil {
		return
	}
	file_system_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_auth_proto_goTypes,
		DependencyIndexes: file_system_auth_proto_depIdxs,
		MessageInfos:      file_system_auth_proto_msgTypes,
	}.Build()
	File_system_auth_proto = out.File
	file_system_auth_proto_rawDesc = nil
	file_system_auth_proto_goTypes = nil
	file_system_auth_proto_depIdxs = nil
}

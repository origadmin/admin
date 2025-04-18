// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: configs/root_user.proto

package configs

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RootUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        bool   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Id             string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Username       string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password       string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Salt           string `protobuf:"bytes,5,opt,name=salt,proto3" json:"salt,omitempty"`
	Name           string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Email          string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Nickname       string `protobuf:"bytes,8,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar         string `protobuf:"bytes,9,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Mobile         string `protobuf:"bytes,10,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Description    string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	AutoCreate     bool   `protobuf:"varint,100,opt,name=auto_create,proto3" json:"auto_create,omitempty"`
	RandomPassword bool   `protobuf:"varint,101,opt,name=random_password,proto3" json:"random_password,omitempty"`
}

func (x *RootUser) Reset() {
	*x = RootUser{}
	mi := &file_configs_root_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RootUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RootUser) ProtoMessage() {}

func (x *RootUser) ProtoReflect() protoreflect.Message {
	mi := &file_configs_root_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RootUser.ProtoReflect.Descriptor instead.
func (*RootUser) Descriptor() ([]byte, []int) {
	return file_configs_root_user_proto_rawDescGZIP(), []int{0}
}

func (x *RootUser) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *RootUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RootUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RootUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RootUser) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *RootUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RootUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RootUser) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RootUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RootUser) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RootUser) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RootUser) GetAutoCreate() bool {
	if x != nil {
		return x.AutoCreate
	}
	return false
}

func (x *RootUser) GetRandomPassword() bool {
	if x != nil {
		return x.RandomPassword
	}
	return false
}

var File_configs_root_user_proto protoreflect.FileDescriptor

var file_configs_root_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x03, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x06, 0x18,
	0x20, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x73,
	0x61, 0x6c, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x06, 0x18, 0x0c, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x72,
	0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x2e,
	0x5a, 0x2c, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configs_root_user_proto_rawDescOnce sync.Once
	file_configs_root_user_proto_rawDescData = file_configs_root_user_proto_rawDesc
)

func file_configs_root_user_proto_rawDescGZIP() []byte {
	file_configs_root_user_proto_rawDescOnce.Do(func() {
		file_configs_root_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_configs_root_user_proto_rawDescData)
	})
	return file_configs_root_user_proto_rawDescData
}

var file_configs_root_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_configs_root_user_proto_goTypes = []any{
	(*RootUser)(nil), // 0: configs.api.RootUser
}
var file_configs_root_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_configs_root_user_proto_init() }
func file_configs_root_user_proto_init() {
	if File_configs_root_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_configs_root_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configs_root_user_proto_goTypes,
		DependencyIndexes: file_configs_root_user_proto_depIdxs,
		MessageInfos:      file_configs_root_user_proto_msgTypes,
	}.Build()
	File_configs_root_user_proto = out.File
	file_configs_root_user_proto_rawDesc = nil
	file_configs_root_user_proto_goTypes = nil
	file_configs_root_user_proto_depIdxs = nil
}

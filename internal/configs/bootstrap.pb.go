// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: configs/bootstrap.proto

package configs

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/origadmin/runtime/gen/go/config/v1"
	v11 "github.com/origadmin/runtime/gen/go/middleware/v1"
	v12 "github.com/origadmin/runtime/gen/go/security/v1"
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

type EntrySelectorConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global  bool   `protobuf:"varint,2,opt,name=global,proto3" json:"global,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EntrySelectorConfig) Reset() {
	*x = EntrySelectorConfig{}
	mi := &file_configs_bootstrap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntrySelectorConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntrySelectorConfig) ProtoMessage() {}

func (x *EntrySelectorConfig) ProtoReflect() protoreflect.Message {
	mi := &file_configs_bootstrap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntrySelectorConfig.ProtoReflect.Descriptor instead.
func (*EntrySelectorConfig) Descriptor() ([]byte, []int) {
	return file_configs_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *EntrySelectorConfig) GetGlobal() bool {
	if x != nil {
		return x.Global
	}
	return false
}

func (x *EntrySelectorConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EntrySelectorConfig) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the application name or service name for used
	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Mode       string            `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
	Version    string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	CryptoType string            `protobuf:"bytes,3,opt,name=crypto_type,proto3" json:"crypto_type,omitempty"`
	Servers    map[string]string `protobuf:"bytes,4,rep,name=servers,proto3" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id         string            `protobuf:"bytes,101,opt,name=id,proto3" json:"id,omitempty"`
	Entry      *Bootstrap_Entry  `protobuf:"bytes,100,opt,name=entry,proto3" json:"entry,omitempty"`
	Service    *v1.Service       `protobuf:"bytes,200,opt,name=service,proto3" json:"service,omitempty"`
	Data       *v1.Data          `protobuf:"bytes,300,opt,name=data,proto3" json:"data,omitempty"`
	Registry   *v1.Registry      `protobuf:"bytes,400,opt,name=registry,proto3" json:"registry,omitempty"`
	Middleware *v11.Middleware   `protobuf:"bytes,9,opt,name=middleware,proto3" json:"middleware,omitempty"`
	Authn      *v12.AuthN        `protobuf:"bytes,1000,opt,name=authn,proto3" json:"authn,omitempty"`
	Authz      *v12.AuthZ        `protobuf:"bytes,1001,opt,name=authz,proto3" json:"authz,omitempty"`
	Security   *v1.Security      `protobuf:"bytes,1002,opt,name=security,proto3" json:"security,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	mi := &file_configs_bootstrap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_configs_bootstrap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_configs_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *Bootstrap) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bootstrap) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Bootstrap) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Bootstrap) GetCryptoType() string {
	if x != nil {
		return x.CryptoType
	}
	return ""
}

func (x *Bootstrap) GetServers() map[string]string {
	if x != nil {
		return x.Servers
	}
	return nil
}

func (x *Bootstrap) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bootstrap) GetEntry() *Bootstrap_Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Bootstrap) GetService() *v1.Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Bootstrap) GetData() *v1.Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *v1.Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

func (x *Bootstrap) GetMiddleware() *v11.Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Bootstrap) GetAuthn() *v12.AuthN {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *Bootstrap) GetAuthz() *v12.AuthZ {
	if x != nil {
		return x.Authz
	}
	return nil
}

func (x *Bootstrap) GetSecurity() *v1.Security {
	if x != nil {
		return x.Security
	}
	return nil
}

type Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CryptoType string `protobuf:"bytes,1,opt,name=crypto_type,proto3" json:"crypto_type,omitempty"`
}

func (x *Settings) Reset() {
	*x = Settings{}
	mi := &file_configs_bootstrap_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Settings) ProtoMessage() {}

func (x *Settings) ProtoReflect() protoreflect.Message {
	mi := &file_configs_bootstrap_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Settings.ProtoReflect.Descriptor instead.
func (*Settings) Descriptor() ([]byte, []int) {
	return file_configs_bootstrap_proto_rawDescGZIP(), []int{2}
}

func (x *Settings) GetCryptoType() string {
	if x != nil {
		return x.CryptoType
	}
	return ""
}

// Entry
type Bootstrap_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme string           `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Grpc   *v1.Service_GRPC `protobuf:"bytes,10,opt,name=grpc,proto3" json:"grpc,omitempty"`
	Http   *v1.Service_HTTP `protobuf:"bytes,20,opt,name=http,proto3" json:"http,omitempty"` //    config.v1.Service.GINS gins = 30 [json_name = "gins"];
}

func (x *Bootstrap_Entry) Reset() {
	*x = Bootstrap_Entry{}
	mi := &file_configs_bootstrap_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bootstrap_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap_Entry) ProtoMessage() {}

func (x *Bootstrap_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_configs_bootstrap_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap_Entry.ProtoReflect.Descriptor instead.
func (*Bootstrap_Entry) Descriptor() ([]byte, []int) {
	return file_configs_bootstrap_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Bootstrap_Entry) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Bootstrap_Entry) GetGrpc() *v1.Service_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Bootstrap_Entry) GetHttp() *v1.Service_HTTP {
	if x != nil {
		return x.Http
	}
	return nil
}

var File_configs_bootstrap_proto protoreflect.FileDescriptor

var file_configs_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x8e, 0x06, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x52, 0x09,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0xac,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x08,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x90, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x4e, 0x52, 0x05, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x18, 0xe9, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x5a, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x12,
	0x30, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0xea, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x79, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x12, 0x2b, 0x0a, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x54,
	0x54, 0x50, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x22, 0x2c, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_configs_bootstrap_proto_rawDescOnce sync.Once
	file_configs_bootstrap_proto_rawDescData = file_configs_bootstrap_proto_rawDesc
)

func file_configs_bootstrap_proto_rawDescGZIP() []byte {
	file_configs_bootstrap_proto_rawDescOnce.Do(func() {
		file_configs_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_configs_bootstrap_proto_rawDescData)
	})
	return file_configs_bootstrap_proto_rawDescData
}

var file_configs_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_configs_bootstrap_proto_goTypes = []any{
	(*EntrySelectorConfig)(nil), // 0: configs.api.EntrySelectorConfig
	(*Bootstrap)(nil),           // 1: configs.api.Bootstrap
	(*Settings)(nil),            // 2: configs.api.Settings
	nil,                         // 3: configs.api.Bootstrap.ServersEntry
	(*Bootstrap_Entry)(nil),     // 4: configs.api.Bootstrap.Entry
	(*v1.Service)(nil),          // 5: config.v1.Service
	(*v1.Data)(nil),             // 6: config.v1.Data
	(*v1.Registry)(nil),         // 7: config.v1.Registry
	(*v11.Middleware)(nil),      // 8: middleware.v1.Build
	(*v12.AuthN)(nil),           // 9: security.v1.AuthN
	(*v12.AuthZ)(nil),           // 10: security.v1.AuthZ
	(*v1.Security)(nil),         // 11: config.v1.Security
	(*v1.Service_GRPC)(nil),     // 12: config.v1.Service.GRPC
	(*v1.Service_HTTP)(nil),     // 13: config.v1.Service.HTTP
}
var file_configs_bootstrap_proto_depIdxs = []int32{
	3,  // 0: configs.api.Bootstrap.servers:type_name -> configs.api.Bootstrap.ServersEntry
	4,  // 1: configs.api.Bootstrap.entry:type_name -> configs.api.Bootstrap.Entry
	5,  // 2: configs.api.Bootstrap.service:type_name -> config.v1.Service
	6,  // 3: configs.api.Bootstrap.data:type_name -> config.v1.Data
	7,  // 4: configs.api.Bootstrap.registry:type_name -> config.v1.Registry
	8,  // 5: configs.api.Bootstrap.middleware:type_name -> middleware.v1.Build
	9,  // 6: configs.api.Bootstrap.authn:type_name -> security.v1.AuthN
	10, // 7: configs.api.Bootstrap.authz:type_name -> security.v1.AuthZ
	11, // 8: configs.api.Bootstrap.security:type_name -> config.v1.Security
	12, // 9: configs.api.Bootstrap.Entry.grpc:type_name -> config.v1.Service.GRPC
	13, // 10: configs.api.Bootstrap.Entry.http:type_name -> config.v1.Service.HTTP
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_configs_bootstrap_proto_init() }
func file_configs_bootstrap_proto_init() {
	if File_configs_bootstrap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_configs_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_configs_bootstrap_proto_goTypes,
		DependencyIndexes: file_configs_bootstrap_proto_depIdxs,
		MessageInfos:      file_configs_bootstrap_proto_msgTypes,
	}.Build()
	File_configs_bootstrap_proto = out.File
	file_configs_bootstrap_proto_rawDesc = nil
	file_configs_bootstrap_proto_goTypes = nil
	file_configs_bootstrap_proto_depIdxs = nil
}

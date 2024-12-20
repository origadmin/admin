// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: system/error.proto

package system

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type SystemErrorReason int32

const (
	SystemErrorReason_SYSTEM_ERROR_REASON_UNSPECIFIED    SystemErrorReason = 0
	SystemErrorReason_SYSTEM_ERROR_REASON_USER_NOT_FOUND SystemErrorReason = 2001
)

// Enum value maps for SystemErrorReason.
var (
	SystemErrorReason_name = map[int32]string{
		0:    "SYSTEM_ERROR_REASON_UNSPECIFIED",
		2001: "SYSTEM_ERROR_REASON_USER_NOT_FOUND",
	}
	SystemErrorReason_value = map[string]int32{
		"SYSTEM_ERROR_REASON_UNSPECIFIED":    0,
		"SYSTEM_ERROR_REASON_USER_NOT_FOUND": 2001,
	}
)

func (x SystemErrorReason) Enum() *SystemErrorReason {
	p := new(SystemErrorReason)
	*p = x
	return p
}

func (x SystemErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_system_error_proto_enumTypes[0].Descriptor()
}

func (SystemErrorReason) Type() protoreflect.EnumType {
	return &file_system_error_proto_enumTypes[0]
}

func (x SystemErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemErrorReason.Descriptor instead.
func (SystemErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_system_error_proto_rawDescGZIP(), []int{0}
}

var File_system_error_proto protoreflect.FileDescriptor

var file_system_error_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x13, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x6d, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2d, 0x0a, 0x22, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0xd1, 0x0f, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03,
	0x42, 0xbf, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x42,
	0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x3b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x53, 0x53, 0xaa,
	0x02, 0x16, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xca, 0x02, 0x16, 0x41, 0x70, 0x69, 0x5c, 0x56,
	0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0xe2, 0x02, 0x22, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_error_proto_rawDescOnce sync.Once
	file_system_error_proto_rawDescData = file_system_error_proto_rawDesc
)

func file_system_error_proto_rawDescGZIP() []byte {
	file_system_error_proto_rawDescOnce.Do(func() {
		file_system_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_error_proto_rawDescData)
	})
	return file_system_error_proto_rawDescData
}

var file_system_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_system_error_proto_goTypes = []any{
	(SystemErrorReason)(0), // 0: api.v1.services.system.SystemErrorReason
}
var file_system_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_error_proto_init() }
func file_system_error_proto_init() {
	if File_system_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_system_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_error_proto_goTypes,
		DependencyIndexes: file_system_error_proto_depIdxs,
		EnumInfos:         file_system_error_proto_enumTypes,
	}.Build()
	File_system_error_proto = out.File
	file_system_error_proto_rawDesc = nil
	file_system_error_proto_goTypes = nil
	file_system_error_proto_depIdxs = nil
}

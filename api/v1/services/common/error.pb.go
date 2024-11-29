// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: common/error.proto

package common

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

type CommonErrorReason int32

const (
	CommonErrorReason_COMMON_ERROR_REASON_UNSPECIFIED    CommonErrorReason = 0
	CommonErrorReason_COMMON_ERROR_REASON_USER_NOT_FOUND CommonErrorReason = 1
)

// Enum value maps for CommonErrorReason.
var (
	CommonErrorReason_name = map[int32]string{
		0: "COMMON_ERROR_REASON_UNSPECIFIED",
		1: "COMMON_ERROR_REASON_USER_NOT_FOUND",
	}
	CommonErrorReason_value = map[string]int32{
		"COMMON_ERROR_REASON_UNSPECIFIED":    0,
		"COMMON_ERROR_REASON_USER_NOT_FOUND": 1,
	}
)

func (x CommonErrorReason) Enum() *CommonErrorReason {
	p := new(CommonErrorReason)
	*p = x
	return p
}

func (x CommonErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_common_error_proto_enumTypes[0].Descriptor()
}

func (CommonErrorReason) Type() protoreflect.EnumType {
	return &file_common_error_proto_enumTypes[0]
}

func (x CommonErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonErrorReason.Descriptor instead.
func (CommonErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_common_error_proto_rawDescGZIP(), []int{0}
}

var File_common_error_proto protoreflect.FileDescriptor

var file_common_error_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x13, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x66, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4d, 0x4d, 0x4f, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x43,
	0x4f, 0x4d, 0x4d, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x01, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xbf, 0x01, 0x0a, 0x1a, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x19, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x53, 0x43, 0xaa, 0x02, 0x16, 0x41, 0x70, 0x69, 0x2e, 0x56,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0xca, 0x02, 0x16, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x22, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x19, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_error_proto_rawDescOnce sync.Once
	file_common_error_proto_rawDescData = file_common_error_proto_rawDesc
)

func file_common_error_proto_rawDescGZIP() []byte {
	file_common_error_proto_rawDescOnce.Do(func() {
		file_common_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_error_proto_rawDescData)
	})
	return file_common_error_proto_rawDescData
}

var file_common_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_error_proto_goTypes = []any{
	(CommonErrorReason)(0), // 0: api.v1.services.common.CommonErrorReason
}
var file_common_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_error_proto_init() }
func file_common_error_proto_init() {
	if File_common_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_error_proto_goTypes,
		DependencyIndexes: file_common_error_proto_depIdxs,
		EnumInfos:         file_common_error_proto_enumTypes,
	}.Build()
	File_common_error_proto = out.File
	file_common_error_proto_rawDesc = nil
	file_common_error_proto_goTypes = nil
	file_common_error_proto_depIdxs = nil
}

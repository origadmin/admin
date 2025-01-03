// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: basis/error.proto

package basis

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

type BasisErrorReason int32

const (
	BasisErrorReason_BASIS_ERROR_REASON_UNSPECIFIED          BasisErrorReason = 0
	BasisErrorReason_BASIS_ERROR_REASON_CAPTCHA_ID_NOT_FOUND BasisErrorReason = 1001
	BasisErrorReason_BASIS_ERROR_REASON_INVALID_CAPTCHA_ID   BasisErrorReason = 1002
	BasisErrorReason_BASIS_ERROR_REASON_INVALID_CAPTCHA_CODE BasisErrorReason = 1003
	BasisErrorReason_BASIS_ERROR_REASON_INVALID_TOKEN        BasisErrorReason = 1004
	BasisErrorReason_BASIS_ERROR_REASON_INVALID_USERNAME     BasisErrorReason = 1005
	BasisErrorReason_BASIS_ERROR_REASON_INVALID_PASSWORD     BasisErrorReason = 1006
)

// Enum value maps for BasisErrorReason.
var (
	BasisErrorReason_name = map[int32]string{
		0:    "BASIS_ERROR_REASON_UNSPECIFIED",
		1001: "BASIS_ERROR_REASON_CAPTCHA_ID_NOT_FOUND",
		1002: "BASIS_ERROR_REASON_INVALID_CAPTCHA_ID",
		1003: "BASIS_ERROR_REASON_INVALID_CAPTCHA_CODE",
		1004: "BASIS_ERROR_REASON_INVALID_TOKEN",
		1005: "BASIS_ERROR_REASON_INVALID_USERNAME",
		1006: "BASIS_ERROR_REASON_INVALID_PASSWORD",
	}
	BasisErrorReason_value = map[string]int32{
		"BASIS_ERROR_REASON_UNSPECIFIED":          0,
		"BASIS_ERROR_REASON_CAPTCHA_ID_NOT_FOUND": 1001,
		"BASIS_ERROR_REASON_INVALID_CAPTCHA_ID":   1002,
		"BASIS_ERROR_REASON_INVALID_CAPTCHA_CODE": 1003,
		"BASIS_ERROR_REASON_INVALID_TOKEN":        1004,
		"BASIS_ERROR_REASON_INVALID_USERNAME":     1005,
		"BASIS_ERROR_REASON_INVALID_PASSWORD":     1006,
	}
)

func (x BasisErrorReason) Enum() *BasisErrorReason {
	p := new(BasisErrorReason)
	*p = x
	return p
}

func (x BasisErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BasisErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_basis_error_proto_enumTypes[0].Descriptor()
}

func (BasisErrorReason) Type() protoreflect.EnumType {
	return &file_basis_error_proto_enumTypes[0]
}

func (x BasisErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BasisErrorReason.Descriptor instead.
func (BasisErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_basis_error_proto_rawDescGZIP(), []int{0}
}

var File_basis_error_proto protoreflect.FileDescriptor

var file_basis_error_proto_rawDesc = []byte{
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x73, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0xe3, 0x02, 0x0a, 0x10, 0x42, 0x61, 0x73, 0x69, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x32, 0x0a, 0x27, 0x42, 0x41, 0x53, 0x49,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x43,
	0x41, 0x50, 0x54, 0x43, 0x48, 0x41, 0x5f, 0x49, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0xe9, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x30, 0x0a, 0x25,
	0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45, 0x41, 0x53,
	0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x50, 0x54, 0x43,
	0x48, 0x41, 0x5f, 0x49, 0x44, 0x10, 0xea, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x32,
	0x0a, 0x27, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x41, 0x50,
	0x54, 0x43, 0x48, 0x41, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0xeb, 0x07, 0x1a, 0x04, 0xa8, 0x45,
	0x90, 0x03, 0x12, 0x2b, 0x0a, 0x20, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xec, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12,
	0x2e, 0x0a, 0x23, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0xed, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12,
	0x2e, 0x0a, 0x23, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x52,
	0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41,
	0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0xee, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0xb8, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x73, 0x42, 0x0a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x17, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x69, 0x73, 0x3b, 0x62, 0x61, 0x73, 0x69, 0x73, 0xa2, 0x02, 0x04, 0x41, 0x56, 0x53,
	0x42, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x73, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x69,
	0x73, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x69, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56, 0x31, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x73, 0x69, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basis_error_proto_rawDescOnce sync.Once
	file_basis_error_proto_rawDescData = file_basis_error_proto_rawDesc
)

func file_basis_error_proto_rawDescGZIP() []byte {
	file_basis_error_proto_rawDescOnce.Do(func() {
		file_basis_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_basis_error_proto_rawDescData)
	})
	return file_basis_error_proto_rawDescData
}

var file_basis_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_basis_error_proto_goTypes = []any{
	(BasisErrorReason)(0), // 0: api.v1.services.basis.BasisErrorReason
}
var file_basis_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basis_error_proto_init() }
func file_basis_error_proto_init() {
	if File_basis_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basis_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basis_error_proto_goTypes,
		DependencyIndexes: file_basis_error_proto_depIdxs,
		EnumInfos:         file_basis_error_proto_enumTypes,
	}.Build()
	File_basis_error_proto = out.File
	file_basis_error_proto_rawDesc = nil
	file_basis_error_proto_goTypes = nil
	file_basis_error_proto_depIdxs = nil
}

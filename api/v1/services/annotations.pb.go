// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: annotations.proto

package services

import (
	_ "github.com/google/gnostic/openapiv3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_annotations_proto protoreflect.FileDescriptor

var file_annotations_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x24, 0x67, 0x6e, 0x6f, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xae, 0x04, 0xba, 0x47, 0x8f,
	0x03, 0x12, 0x8c, 0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x20,
	0x41, 0x50, 0x49, 0x12, 0x5f, 0x41, 0x20, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x2c, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x69, 0x62, 0x6c, 0x65, 0x2c, 0x20, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x6e, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x66, 0x75, 0x6c, 0x6c, 0x2d, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x20, 0x52, 0x42, 0x41, 0x43, 0x20, 0x73, 0x63, 0x61,
	0x66, 0x66, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x22, 0x40, 0x0a, 0x07, 0x47, 0x6f, 0x64, 0x43, 0x6f, 0x6e, 0x67, 0x12,
	0x1c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17, 0x77,
	0x61, 0x69, 0x74, 0x66, 0x6f, 0x72, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x40, 0x67, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x2a, 0x3f, 0x0a, 0x03, 0x4d, 0x49, 0x54, 0x12, 0x38, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2f,
	0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x32, 0x17, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x18, 0x0a, 0x16, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x68, 0x6f, 0x73, 0x74, 0x3a, 0x31, 0x30, 0x30, 0x38, 0x30, 0x1a, 0x19, 0x0a, 0x17, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a,
	0x31, 0x30, 0x30, 0x38, 0x30, 0x2a, 0x49, 0x3a, 0x47, 0x0a, 0x18, 0x0a, 0x05, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x12, 0x0f, 0x0a, 0x0d, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x2a, 0x05, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x0a, 0x2b, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x1f, 0x0a, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2,
	0x02, 0x03, 0x41, 0x56, 0x53, 0xaa, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xca, 0x02, 0x0f, 0x41, 0x70, 0x69, 0x5c, 0x56, 0x31,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xe2, 0x02, 0x1b, 0x41, 0x70, 0x69, 0x5c,
	0x56, 0x31, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x56,
	0x31, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_annotations_proto_goTypes = []any{}
var file_annotations_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_annotations_proto_init() }
func file_annotations_proto_init() {
	if File_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_annotations_proto_goTypes,
		DependencyIndexes: file_annotations_proto_depIdxs,
	}.Build()
	File_annotations_proto = out.File
	file_annotations_proto_rawDesc = nil
	file_annotations_proto_goTypes = nil
	file_annotations_proto_depIdxs = nil
}

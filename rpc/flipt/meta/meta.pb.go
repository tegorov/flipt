// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        (unknown)
// source: meta/meta.proto

package meta

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_meta_meta_proto protoreflect.FileDescriptor

var file_meta_meta_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f,
	0x64, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8c, 0x02, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x43, 0x92, 0x41,
	0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x21, 0x47, 0x65, 0x74,
	0x20, 0x46, 0x6c, 0x69, 0x70, 0x74, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x20,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x2a, 0x11,
	0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x71, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x38, 0x92, 0x41, 0x35, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x1f, 0x47, 0x65, 0x74, 0x20, 0x46,
	0x6c, 0x69, 0x70, 0x74, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x2a, 0x08, 0x67, 0x65, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0xce, 0x03, 0x92, 0x41, 0xa8, 0x03, 0x12, 0xaa, 0x01, 0x0a, 0x13,
	0x46, 0x6c, 0x69, 0x70, 0x74, 0x20, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x20, 0x41,
	0x50, 0x49, 0x73, 0x22, 0x3d, 0x0a, 0x0a, 0x46, 0x6c, 0x69, 0x70, 0x74, 0x20, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x66,
	0x6c, 0x69, 0x70, 0x74, 0x1a, 0x0c, 0x64, 0x65, 0x76, 0x40, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e,
	0x69, 0x6f, 0x2a, 0x4c, 0x0a, 0x0b, 0x4d, 0x49, 0x54, 0x20, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x66,
	0x6c, 0x69, 0x70, 0x74, 0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45,
	0x32, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x52, 0x63, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x5c, 0x0a, 0x3d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x63, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x20, 0x28, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x29, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5a, 0x2a, 0x0a, 0x28, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x13, 0x08, 0x02,
	0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x02, 0x62, 0x17, 0x0a, 0x15, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x72, 0x27, 0x0a, 0x0a, 0x46, 0x6c,
	0x69, 0x70, 0x74, 0x20, 0x44, 0x6f, 0x63, 0x73, 0x12, 0x19, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x64,
	0x6f, 0x63, 0x73, 0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2e, 0x69, 0x6f,
	0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x66, 0x6c, 0x69, 0x70, 0x74,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_meta_meta_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),     // 0: google.protobuf.Empty
	(*httpbody.HttpBody)(nil), // 1: google.api.HttpBody
}
var file_meta_meta_proto_depIdxs = []int32{
	0, // 0: flipt.meta.MetadataService.GetConfiguration:input_type -> google.protobuf.Empty
	0, // 1: flipt.meta.MetadataService.GetInfo:input_type -> google.protobuf.Empty
	1, // 2: flipt.meta.MetadataService.GetConfiguration:output_type -> google.api.HttpBody
	1, // 3: flipt.meta.MetadataService.GetInfo:output_type -> google.api.HttpBody
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_meta_proto_init() }
func file_meta_meta_proto_init() {
	if File_meta_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_meta_meta_proto_goTypes,
		DependencyIndexes: file_meta_meta_proto_depIdxs,
	}.Build()
	File_meta_meta_proto = out.File
	file_meta_meta_proto_rawDesc = nil
	file_meta_meta_proto_goTypes = nil
	file_meta_meta_proto_depIdxs = nil
}
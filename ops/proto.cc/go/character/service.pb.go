// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: character/service.proto

package character

import (
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

var File_character_service_proto protoreflect.FileDescriptor

var file_character_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x69, 0x67, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x1a, 0x17, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdb, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x12, 0x27, 0x2e, 0x62, 0x69,
	0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e,
	0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x26, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x69, 0x67,
	0x73, 0x74, 0x61, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x64, 0x74, 0x68, 0x65, 0x72, 0x68, 0x74, 0x75, 0x6e, 0x2f, 0x4c, 0x65, 0x61, 0x72,
	0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_character_service_proto_goTypes = []interface{}{
	(*AllCharactersRequest)(nil),  // 0: bigstar.character.AllCharactersRequest
	(*GetCharacterRequest)(nil),   // 1: bigstar.character.GetCharacterRequest
	(*AllCharactersResponse)(nil), // 2: bigstar.character.AllCharactersResponse
	(*GetCharacterResponse)(nil),  // 3: bigstar.character.GetCharacterResponse
}
var file_character_service_proto_depIdxs = []int32{
	0, // 0: bigstar.character.CharacterService.GetCharacters:input_type -> bigstar.character.AllCharactersRequest
	1, // 1: bigstar.character.CharacterService.GetCharacterById:input_type -> bigstar.character.GetCharacterRequest
	2, // 2: bigstar.character.CharacterService.GetCharacters:output_type -> bigstar.character.AllCharactersResponse
	3, // 3: bigstar.character.CharacterService.GetCharacterById:output_type -> bigstar.character.GetCharacterResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_character_service_proto_init() }
func file_character_service_proto_init() {
	if File_character_service_proto != nil {
		return
	}
	file_character_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_character_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_character_service_proto_goTypes,
		DependencyIndexes: file_character_service_proto_depIdxs,
	}.Build()
	File_character_service_proto = out.File
	file_character_service_proto_rawDesc = nil
	file_character_service_proto_goTypes = nil
	file_character_service_proto_depIdxs = nil
}

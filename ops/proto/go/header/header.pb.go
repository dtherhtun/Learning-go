// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: header/header.proto

package header

import (
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

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_header_header_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_header_header_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_header_header_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Span     *UUID `protobuf:"bytes,1,opt,name=span,proto3" json:"span,omitempty"`
	RootSpan *UUID `protobuf:"bytes,2,opt,name=root_span,json=rootSpan,proto3" json:"root_span,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_header_header_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_header_header_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_header_header_proto_rawDescGZIP(), []int{1}
}

func (x *Header) GetSpan() *UUID {
	if x != nil {
		return x.Span
	}
	return nil
}

func (x *Header) GetRootSpan() *UUID {
	if x != nil {
		return x.RootSpan
	}
	return nil
}

var File_header_header_proto protoreflect.FileDescriptor

var file_header_header_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x22, 0x1c,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x57, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x52, 0x04, 0x73, 0x70, 0x61, 0x6e, 0x12, 0x2a, 0x0a, 0x09, 0x72, 0x6f, 0x6f,
	0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62,
	0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x08, 0x72, 0x6f, 0x6f,
	0x74, 0x53, 0x70, 0x61, 0x6e, 0x42, 0x61, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x42, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x74, 0x68, 0x65, 0x72, 0x68, 0x74, 0x75, 0x6e, 0x2f, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e,
	0x67, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_header_header_proto_rawDescOnce sync.Once
	file_header_header_proto_rawDescData = file_header_header_proto_rawDesc
)

func file_header_header_proto_rawDescGZIP() []byte {
	file_header_header_proto_rawDescOnce.Do(func() {
		file_header_header_proto_rawDescData = protoimpl.X.CompressGZIP(file_header_header_proto_rawDescData)
	})
	return file_header_header_proto_rawDescData
}

var file_header_header_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_header_header_proto_goTypes = []interface{}{
	(*UUID)(nil),   // 0: bigstar.UUID
	(*Header)(nil), // 1: bigstar.Header
}
var file_header_header_proto_depIdxs = []int32{
	0, // 0: bigstar.Header.span:type_name -> bigstar.UUID
	0, // 1: bigstar.Header.root_span:type_name -> bigstar.UUID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_header_header_proto_init() }
func file_header_header_proto_init() {
	if File_header_header_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_header_header_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_header_header_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_header_header_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_header_header_proto_goTypes,
		DependencyIndexes: file_header_header_proto_depIdxs,
		MessageInfos:      file_header_header_proto_msgTypes,
	}.Build()
	File_header_header_proto = out.File
	file_header_header_proto_rawDesc = nil
	file_header_header_proto_goTypes = nil
	file_header_header_proto_depIdxs = nil
}

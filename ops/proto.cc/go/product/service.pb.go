// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: product/service.proto

package product

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

var File_product_service_proto protoreflect.FileDescriptor

var file_product_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0x97, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x58, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x23, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x22,
	0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x69, 0x67, 0x73, 0x74,
	0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x74, 0x68, 0x65, 0x72, 0x68, 0x74, 0x75,
	0x6e, 0x2f, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x67, 0x6f, 0x2f, 0x6f, 0x70,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_product_service_proto_goTypes = []interface{}{
	(*AllProductsRequest)(nil),  // 0: bigstar.product.AllProductsRequest
	(*GetProductRequest)(nil),   // 1: bigstar.product.GetProductRequest
	(*SearchRequest)(nil),       // 2: bigstar.product.SearchRequest
	(*AllProductsResponse)(nil), // 3: bigstar.product.AllProductsResponse
	(*GetProductResponse)(nil),  // 4: bigstar.product.GetProductResponse
	(*SearchResponse)(nil),      // 5: bigstar.product.SearchResponse
}
var file_product_service_proto_depIdxs = []int32{
	0, // 0: bigstar.product.ProductService.GetProducts:input_type -> bigstar.product.AllProductsRequest
	1, // 1: bigstar.product.ProductService.GetProductById:input_type -> bigstar.product.GetProductRequest
	2, // 2: bigstar.product.ProductService.ProductSearch:input_type -> bigstar.product.SearchRequest
	3, // 3: bigstar.product.ProductService.GetProducts:output_type -> bigstar.product.AllProductsResponse
	4, // 4: bigstar.product.ProductService.GetProductById:output_type -> bigstar.product.GetProductResponse
	5, // 5: bigstar.product.ProductService.ProductSearch:output_type -> bigstar.product.SearchResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_service_proto_init() }
func file_product_service_proto_init() {
	if File_product_service_proto != nil {
		return
	}
	file_product_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_service_proto_goTypes,
		DependencyIndexes: file_product_service_proto_depIdxs,
	}.Build()
	File_product_service_proto = out.File
	file_product_service_proto_rawDesc = nil
	file_product_service_proto_goTypes = nil
	file_product_service_proto_depIdxs = nil
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: character/service.proto

package character

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CharacterService_GetCharacters_FullMethodName    = "/bigstar.character.CharacterService/GetCharacters"
	CharacterService_GetCharacterById_FullMethodName = "/bigstar.character.CharacterService/GetCharacterById"
)

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterServiceClient interface {
	GetCharacters(ctx context.Context, in *AllCharactersRequest, opts ...grpc.CallOption) (*AllCharactersResponse, error)
	GetCharacterById(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*GetCharacterResponse, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) GetCharacters(ctx context.Context, in *AllCharactersRequest, opts ...grpc.CallOption) (*AllCharactersResponse, error) {
	out := new(AllCharactersResponse)
	err := c.cc.Invoke(ctx, CharacterService_GetCharacters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) GetCharacterById(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*GetCharacterResponse, error) {
	out := new(GetCharacterResponse)
	err := c.cc.Invoke(ctx, CharacterService_GetCharacterById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
// All implementations must embed UnimplementedCharacterServiceServer
// for forward compatibility
type CharacterServiceServer interface {
	GetCharacters(context.Context, *AllCharactersRequest) (*AllCharactersResponse, error)
	GetCharacterById(context.Context, *GetCharacterRequest) (*GetCharacterResponse, error)
	mustEmbedUnimplementedCharacterServiceServer()
}

// UnimplementedCharacterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (UnimplementedCharacterServiceServer) GetCharacters(context.Context, *AllCharactersRequest) (*AllCharactersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacters not implemented")
}
func (UnimplementedCharacterServiceServer) GetCharacterById(context.Context, *GetCharacterRequest) (*GetCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacterById not implemented")
}
func (UnimplementedCharacterServiceServer) mustEmbedUnimplementedCharacterServiceServer() {}

// UnsafeCharacterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServiceServer will
// result in compilation errors.
type UnsafeCharacterServiceServer interface {
	mustEmbedUnimplementedCharacterServiceServer()
}

func RegisterCharacterServiceServer(s grpc.ServiceRegistrar, srv CharacterServiceServer) {
	s.RegisterService(&CharacterService_ServiceDesc, srv)
}

func _CharacterService_GetCharacters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllCharactersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_GetCharacters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacters(ctx, req.(*AllCharactersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_GetCharacterById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacterById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CharacterService_GetCharacterById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacterById(ctx, req.(*GetCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterService_ServiceDesc is the grpc.ServiceDesc for CharacterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bigstar.character.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharacters",
			Handler:    _CharacterService_GetCharacters_Handler,
		},
		{
			MethodName: "GetCharacterById",
			Handler:    _CharacterService_GetCharacterById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "character/service.proto",
}
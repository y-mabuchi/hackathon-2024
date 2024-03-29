// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: family.proto

package generated

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

// FamilyServiceClient is the client API for FamilyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FamilyServiceClient interface {
	FindOneFamily(ctx context.Context, in *FindOneFamilyRequest, opts ...grpc.CallOption) (*FindOneFamilyResponse, error)
	CreateFamily(ctx context.Context, in *CreateFamilyRequest, opts ...grpc.CallOption) (*CreateFamilyResponse, error)
	UpdateFamily(ctx context.Context, in *UpdateFamilyRequest, opts ...grpc.CallOption) (*UpdateFamilyResponse, error)
	DeleteFamily(ctx context.Context, in *DeleteFamilyRequest, opts ...grpc.CallOption) (*DeleteFamilyResponse, error)
}

type familyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFamilyServiceClient(cc grpc.ClientConnInterface) FamilyServiceClient {
	return &familyServiceClient{cc}
}

func (c *familyServiceClient) FindOneFamily(ctx context.Context, in *FindOneFamilyRequest, opts ...grpc.CallOption) (*FindOneFamilyResponse, error) {
	out := new(FindOneFamilyResponse)
	err := c.cc.Invoke(ctx, "/family.FamilyService/FindOneFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyServiceClient) CreateFamily(ctx context.Context, in *CreateFamilyRequest, opts ...grpc.CallOption) (*CreateFamilyResponse, error) {
	out := new(CreateFamilyResponse)
	err := c.cc.Invoke(ctx, "/family.FamilyService/CreateFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyServiceClient) UpdateFamily(ctx context.Context, in *UpdateFamilyRequest, opts ...grpc.CallOption) (*UpdateFamilyResponse, error) {
	out := new(UpdateFamilyResponse)
	err := c.cc.Invoke(ctx, "/family.FamilyService/UpdateFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *familyServiceClient) DeleteFamily(ctx context.Context, in *DeleteFamilyRequest, opts ...grpc.CallOption) (*DeleteFamilyResponse, error) {
	out := new(DeleteFamilyResponse)
	err := c.cc.Invoke(ctx, "/family.FamilyService/DeleteFamily", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FamilyServiceServer is the server API for FamilyService service.
// All implementations must embed UnimplementedFamilyServiceServer
// for forward compatibility
type FamilyServiceServer interface {
	FindOneFamily(context.Context, *FindOneFamilyRequest) (*FindOneFamilyResponse, error)
	CreateFamily(context.Context, *CreateFamilyRequest) (*CreateFamilyResponse, error)
	UpdateFamily(context.Context, *UpdateFamilyRequest) (*UpdateFamilyResponse, error)
	DeleteFamily(context.Context, *DeleteFamilyRequest) (*DeleteFamilyResponse, error)
	mustEmbedUnimplementedFamilyServiceServer()
}

// UnimplementedFamilyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFamilyServiceServer struct {
}

func (UnimplementedFamilyServiceServer) FindOneFamily(context.Context, *FindOneFamilyRequest) (*FindOneFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneFamily not implemented")
}
func (UnimplementedFamilyServiceServer) CreateFamily(context.Context, *CreateFamilyRequest) (*CreateFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFamily not implemented")
}
func (UnimplementedFamilyServiceServer) UpdateFamily(context.Context, *UpdateFamilyRequest) (*UpdateFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFamily not implemented")
}
func (UnimplementedFamilyServiceServer) DeleteFamily(context.Context, *DeleteFamilyRequest) (*DeleteFamilyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFamily not implemented")
}
func (UnimplementedFamilyServiceServer) mustEmbedUnimplementedFamilyServiceServer() {}

// UnsafeFamilyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FamilyServiceServer will
// result in compilation errors.
type UnsafeFamilyServiceServer interface {
	mustEmbedUnimplementedFamilyServiceServer()
}

func RegisterFamilyServiceServer(s grpc.ServiceRegistrar, srv FamilyServiceServer) {
	s.RegisterService(&FamilyService_ServiceDesc, srv)
}

func _FamilyService_FindOneFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilyServiceServer).FindOneFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/family.FamilyService/FindOneFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilyServiceServer).FindOneFamily(ctx, req.(*FindOneFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FamilyService_CreateFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilyServiceServer).CreateFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/family.FamilyService/CreateFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilyServiceServer).CreateFamily(ctx, req.(*CreateFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FamilyService_UpdateFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilyServiceServer).UpdateFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/family.FamilyService/UpdateFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilyServiceServer).UpdateFamily(ctx, req.(*UpdateFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FamilyService_DeleteFamily_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFamilyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FamilyServiceServer).DeleteFamily(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/family.FamilyService/DeleteFamily",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FamilyServiceServer).DeleteFamily(ctx, req.(*DeleteFamilyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FamilyService_ServiceDesc is the grpc.ServiceDesc for FamilyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FamilyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "family.FamilyService",
	HandlerType: (*FamilyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOneFamily",
			Handler:    _FamilyService_FindOneFamily_Handler,
		},
		{
			MethodName: "CreateFamily",
			Handler:    _FamilyService_CreateFamily_Handler,
		},
		{
			MethodName: "UpdateFamily",
			Handler:    _FamilyService_UpdateFamily_Handler,
		},
		{
			MethodName: "DeleteFamily",
			Handler:    _FamilyService_DeleteFamily_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "family.proto",
}

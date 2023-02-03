// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.0
// source: grpc/proto/images_handler.proto

package api

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

// ImagesHandlerClient is the client API for ImagesHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImagesHandlerClient interface {
	SetImage(ctx context.Context, in *SetImageRequest, opts ...grpc.CallOption) (*SetImageResponse, error)
	GetImagesList(ctx context.Context, in *GetImagesListRequest, opts ...grpc.CallOption) (*GetImagesListResponse, error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error)
}

type imagesHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewImagesHandlerClient(cc grpc.ClientConnInterface) ImagesHandlerClient {
	return &imagesHandlerClient{cc}
}

func (c *imagesHandlerClient) SetImage(ctx context.Context, in *SetImageRequest, opts ...grpc.CallOption) (*SetImageResponse, error) {
	out := new(SetImageResponse)
	err := c.cc.Invoke(ctx, "/proto.ImagesHandler/SetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesHandlerClient) GetImagesList(ctx context.Context, in *GetImagesListRequest, opts ...grpc.CallOption) (*GetImagesListResponse, error) {
	out := new(GetImagesListResponse)
	err := c.cc.Invoke(ctx, "/proto.ImagesHandler/GetImagesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesHandlerClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*GetImageResponse, error) {
	out := new(GetImageResponse)
	err := c.cc.Invoke(ctx, "/proto.ImagesHandler/GetImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImagesHandlerServer is the server API for ImagesHandler service.
// All implementations must embed UnimplementedImagesHandlerServer
// for forward compatibility
type ImagesHandlerServer interface {
	SetImage(context.Context, *SetImageRequest) (*SetImageResponse, error)
	GetImagesList(context.Context, *GetImagesListRequest) (*GetImagesListResponse, error)
	GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error)
	mustEmbedUnimplementedImagesHandlerServer()
}

// UnimplementedImagesHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedImagesHandlerServer struct {
}

func (UnimplementedImagesHandlerServer) SetImage(context.Context, *SetImageRequest) (*SetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetImage not implemented")
}
func (UnimplementedImagesHandlerServer) GetImagesList(context.Context, *GetImagesListRequest) (*GetImagesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImagesList not implemented")
}
func (UnimplementedImagesHandlerServer) GetImage(context.Context, *GetImageRequest) (*GetImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedImagesHandlerServer) mustEmbedUnimplementedImagesHandlerServer() {}

// UnsafeImagesHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImagesHandlerServer will
// result in compilation errors.
type UnsafeImagesHandlerServer interface {
	mustEmbedUnimplementedImagesHandlerServer()
}

func RegisterImagesHandlerServer(s grpc.ServiceRegistrar, srv ImagesHandlerServer) {
	s.RegisterService(&ImagesHandler_ServiceDesc, srv)
}

func _ImagesHandler_SetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesHandlerServer).SetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImagesHandler/SetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesHandlerServer).SetImage(ctx, req.(*SetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImagesHandler_GetImagesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImagesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesHandlerServer).GetImagesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImagesHandler/GetImagesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesHandlerServer).GetImagesList(ctx, req.(*GetImagesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImagesHandler_GetImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesHandlerServer).GetImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImagesHandler/GetImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesHandlerServer).GetImage(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImagesHandler_ServiceDesc is the grpc.ServiceDesc for ImagesHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImagesHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ImagesHandler",
	HandlerType: (*ImagesHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetImage",
			Handler:    _ImagesHandler_SetImage_Handler,
		},
		{
			MethodName: "GetImagesList",
			Handler:    _ImagesHandler_GetImagesList_Handler,
		},
		{
			MethodName: "GetImage",
			Handler:    _ImagesHandler_GetImage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/images_handler.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: container.proto

package container

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

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerServiceClient interface {
	// 准备值
	PrepareValue(ctx context.Context, in *PrepareReq, opts ...grpc.CallOption) (*PrepareRes, error)
	// 获取值
	FetchValue(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchRes, error)
	// Append 添加值
	AppendValue(ctx context.Context, in *AppendReq, opts ...grpc.CallOption) (*AppendRes, error)
	// watch完成标记
	WatchValue(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRes, error)
}

type containerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerServiceClient(cc grpc.ClientConnInterface) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) PrepareValue(ctx context.Context, in *PrepareReq, opts ...grpc.CallOption) (*PrepareRes, error) {
	out := new(PrepareRes)
	err := c.cc.Invoke(ctx, "/container.ContainerService/PrepareValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) FetchValue(ctx context.Context, in *FetchReq, opts ...grpc.CallOption) (*FetchRes, error) {
	out := new(FetchRes)
	err := c.cc.Invoke(ctx, "/container.ContainerService/FetchValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) AppendValue(ctx context.Context, in *AppendReq, opts ...grpc.CallOption) (*AppendRes, error) {
	out := new(AppendRes)
	err := c.cc.Invoke(ctx, "/container.ContainerService/AppendValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *containerServiceClient) WatchValue(ctx context.Context, in *WatchReq, opts ...grpc.CallOption) (*WatchRes, error) {
	out := new(WatchRes)
	err := c.cc.Invoke(ctx, "/container.ContainerService/WatchValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContainerServiceServer is the server API for ContainerService service.
// All implementations must embed UnimplementedContainerServiceServer
// for forward compatibility
type ContainerServiceServer interface {
	// 准备值
	PrepareValue(context.Context, *PrepareReq) (*PrepareRes, error)
	// 获取值
	FetchValue(context.Context, *FetchReq) (*FetchRes, error)
	// Append 添加值
	AppendValue(context.Context, *AppendReq) (*AppendRes, error)
	// watch完成标记
	WatchValue(context.Context, *WatchReq) (*WatchRes, error)
	mustEmbedUnimplementedContainerServiceServer()
}

// UnimplementedContainerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (UnimplementedContainerServiceServer) PrepareValue(context.Context, *PrepareReq) (*PrepareRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareValue not implemented")
}
func (UnimplementedContainerServiceServer) FetchValue(context.Context, *FetchReq) (*FetchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchValue not implemented")
}
func (UnimplementedContainerServiceServer) AppendValue(context.Context, *AppendReq) (*AppendRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendValue not implemented")
}
func (UnimplementedContainerServiceServer) WatchValue(context.Context, *WatchReq) (*WatchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchValue not implemented")
}
func (UnimplementedContainerServiceServer) mustEmbedUnimplementedContainerServiceServer() {}

// UnsafeContainerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerServiceServer will
// result in compilation errors.
type UnsafeContainerServiceServer interface {
	mustEmbedUnimplementedContainerServiceServer()
}

func RegisterContainerServiceServer(s grpc.ServiceRegistrar, srv ContainerServiceServer) {
	s.RegisterService(&ContainerService_ServiceDesc, srv)
}

func _ContainerService_PrepareValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).PrepareValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.ContainerService/PrepareValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).PrepareValue(ctx, req.(*PrepareReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_FetchValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).FetchValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.ContainerService/FetchValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).FetchValue(ctx, req.(*FetchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_AppendValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).AppendValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.ContainerService/AppendValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).AppendValue(ctx, req.(*AppendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContainerService_WatchValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContainerServiceServer).WatchValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/container.ContainerService/WatchValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContainerServiceServer).WatchValue(ctx, req.(*WatchReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ContainerService_ServiceDesc is the grpc.ServiceDesc for ContainerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "container.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareValue",
			Handler:    _ContainerService_PrepareValue_Handler,
		},
		{
			MethodName: "FetchValue",
			Handler:    _ContainerService_FetchValue_Handler,
		},
		{
			MethodName: "AppendValue",
			Handler:    _ContainerService_AppendValue_Handler,
		},
		{
			MethodName: "WatchValue",
			Handler:    _ContainerService_WatchValue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "container.proto",
}

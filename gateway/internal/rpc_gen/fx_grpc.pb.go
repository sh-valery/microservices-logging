// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: fx.proto

package rpc_gen

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

// FxServiceClient is the client API for FxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FxServiceClient interface {
	GetFxRate(ctx context.Context, in *FxServiceRequest, opts ...grpc.CallOption) (*FxServiceResponse, error)
}

type fxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFxServiceClient(cc grpc.ClientConnInterface) FxServiceClient {
	return &fxServiceClient{cc}
}

func (c *fxServiceClient) GetFxRate(ctx context.Context, in *FxServiceRequest, opts ...grpc.CallOption) (*FxServiceResponse, error) {
	out := new(FxServiceResponse)
	err := c.cc.Invoke(ctx, "/FxService/GetFxRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FxServiceServer is the server API for FxService service.
// All implementations must embed UnimplementedFxServiceServer
// for forward compatibility
type FxServiceServer interface {
	GetFxRate(context.Context, *FxServiceRequest) (*FxServiceResponse, error)
	mustEmbedUnimplementedFxServiceServer()
}

// UnimplementedFxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFxServiceServer struct {
}

func (UnimplementedFxServiceServer) GetFxRate(context.Context, *FxServiceRequest) (*FxServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFxRate not implemented")
}
func (UnimplementedFxServiceServer) mustEmbedUnimplementedFxServiceServer() {}

// UnsafeFxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FxServiceServer will
// result in compilation errors.
type UnsafeFxServiceServer interface {
	mustEmbedUnimplementedFxServiceServer()
}

func RegisterFxServiceServer(s grpc.ServiceRegistrar, srv FxServiceServer) {
	s.RegisterService(&FxService_ServiceDesc, srv)
}

func _FxService_GetFxRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FxServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FxServiceServer).GetFxRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FxService/GetFxRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FxServiceServer).GetFxRate(ctx, req.(*FxServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FxService_ServiceDesc is the grpc.ServiceDesc for FxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FxService",
	HandlerType: (*FxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFxRate",
			Handler:    _FxService_GetFxRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fx.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: deliveryOptions.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DeliveryOptionsService_CalculateDeliveryOptions_FullMethodName = "/proto.DeliveryOptionsService/CalculateDeliveryOptions"
)

// DeliveryOptionsServiceClient is the client API for DeliveryOptionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeliveryOptionsServiceClient interface {
	CalculateDeliveryOptions(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
}

type deliveryOptionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeliveryOptionsServiceClient(cc grpc.ClientConnInterface) DeliveryOptionsServiceClient {
	return &deliveryOptionsServiceClient{cc}
}

func (c *deliveryOptionsServiceClient) CalculateDeliveryOptions(ctx context.Context, in *DeliveryRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, DeliveryOptionsService_CalculateDeliveryOptions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeliveryOptionsServiceServer is the server API for DeliveryOptionsService service.
// All implementations must embed UnimplementedDeliveryOptionsServiceServer
// for forward compatibility.
type DeliveryOptionsServiceServer interface {
	CalculateDeliveryOptions(context.Context, *DeliveryRequest) (*DeliveryResponse, error)
	mustEmbedUnimplementedDeliveryOptionsServiceServer()
}

// UnimplementedDeliveryOptionsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeliveryOptionsServiceServer struct{}

func (UnimplementedDeliveryOptionsServiceServer) CalculateDeliveryOptions(context.Context, *DeliveryRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateDeliveryOptions not implemented")
}
func (UnimplementedDeliveryOptionsServiceServer) mustEmbedUnimplementedDeliveryOptionsServiceServer() {
}
func (UnimplementedDeliveryOptionsServiceServer) testEmbeddedByValue() {}

// UnsafeDeliveryOptionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeliveryOptionsServiceServer will
// result in compilation errors.
type UnsafeDeliveryOptionsServiceServer interface {
	mustEmbedUnimplementedDeliveryOptionsServiceServer()
}

func RegisterDeliveryOptionsServiceServer(s grpc.ServiceRegistrar, srv DeliveryOptionsServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeliveryOptionsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeliveryOptionsService_ServiceDesc, srv)
}

func _DeliveryOptionsService_CalculateDeliveryOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeliveryOptionsServiceServer).CalculateDeliveryOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeliveryOptionsService_CalculateDeliveryOptions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeliveryOptionsServiceServer).CalculateDeliveryOptions(ctx, req.(*DeliveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeliveryOptionsService_ServiceDesc is the grpc.ServiceDesc for DeliveryOptionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeliveryOptionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DeliveryOptionsService",
	HandlerType: (*DeliveryOptionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalculateDeliveryOptions",
			Handler:    _DeliveryOptionsService_CalculateDeliveryOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deliveryOptions.proto",
}

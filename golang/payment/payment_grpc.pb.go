// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0
// source: payment/payment.proto

package payment

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
	Payment_Create_FullMethodName = "/Payment/Create"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentResponse)
	err := c.cc.Invoke(ctx, Payment_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility.
type PaymentServer interface {
	Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServer struct{}

func (UnimplementedPaymentServer) Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}
func (UnimplementedPaymentServer) testEmbeddedByValue()                 {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).Create(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Payment_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/payment.proto",
}

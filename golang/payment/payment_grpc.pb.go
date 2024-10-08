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
	Payment_CreatePaymentMethod_FullMethodName = "/Payment/CreatePaymentMethod"
	Payment_GetPaymentMethod_FullMethodName    = "/Payment/GetPaymentMethod"
	Payment_UpdatePaymentMethod_FullMethodName = "/Payment/UpdatePaymentMethod"
	Payment_DeletePaymentMethod_FullMethodName = "/Payment/DeletePaymentMethod"
	Payment_Create_FullMethodName              = "/Payment/Create"
	Payment_Get_FullMethodName                 = "/Payment/Get"
)

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	CreatePaymentMethod(ctx context.Context, in *CreatePaymentMethodRequest, opts ...grpc.CallOption) (*CreatePaymentMethodResponse, error)
	GetPaymentMethod(ctx context.Context, in *GetPaymentMethodRequest, opts ...grpc.CallOption) (*GetPaymentMethodResponse, error)
	UpdatePaymentMethod(ctx context.Context, in *UpdatePaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePaymentMethodResponse, error)
	DeletePaymentMethod(ctx context.Context, in *DeletePaymentMethodRequest, opts ...grpc.CallOption) (*DeletePaymentMethodResponse, error)
	Create(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*CreatePaymentResponse, error)
	Get(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) CreatePaymentMethod(ctx context.Context, in *CreatePaymentMethodRequest, opts ...grpc.CallOption) (*CreatePaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePaymentMethodResponse)
	err := c.cc.Invoke(ctx, Payment_CreatePaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetPaymentMethod(ctx context.Context, in *GetPaymentMethodRequest, opts ...grpc.CallOption) (*GetPaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentMethodResponse)
	err := c.cc.Invoke(ctx, Payment_GetPaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) UpdatePaymentMethod(ctx context.Context, in *UpdatePaymentMethodRequest, opts ...grpc.CallOption) (*UpdatePaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePaymentMethodResponse)
	err := c.cc.Invoke(ctx, Payment_UpdatePaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) DeletePaymentMethod(ctx context.Context, in *DeletePaymentMethodRequest, opts ...grpc.CallOption) (*DeletePaymentMethodResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePaymentMethodResponse)
	err := c.cc.Invoke(ctx, Payment_DeletePaymentMethod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *paymentClient) Get(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*GetPaymentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPaymentResponse)
	err := c.cc.Invoke(ctx, Payment_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility.
type PaymentServer interface {
	CreatePaymentMethod(context.Context, *CreatePaymentMethodRequest) (*CreatePaymentMethodResponse, error)
	GetPaymentMethod(context.Context, *GetPaymentMethodRequest) (*GetPaymentMethodResponse, error)
	UpdatePaymentMethod(context.Context, *UpdatePaymentMethodRequest) (*UpdatePaymentMethodResponse, error)
	DeletePaymentMethod(context.Context, *DeletePaymentMethodRequest) (*DeletePaymentMethodResponse, error)
	Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error)
	Get(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServer struct{}

func (UnimplementedPaymentServer) CreatePaymentMethod(context.Context, *CreatePaymentMethodRequest) (*CreatePaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePaymentMethod not implemented")
}
func (UnimplementedPaymentServer) GetPaymentMethod(context.Context, *GetPaymentMethodRequest) (*GetPaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentMethod not implemented")
}
func (UnimplementedPaymentServer) UpdatePaymentMethod(context.Context, *UpdatePaymentMethodRequest) (*UpdatePaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePaymentMethod not implemented")
}
func (UnimplementedPaymentServer) DeletePaymentMethod(context.Context, *DeletePaymentMethodRequest) (*DeletePaymentMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePaymentMethod not implemented")
}
func (UnimplementedPaymentServer) Create(context.Context, *CreatePaymentRequest) (*CreatePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentServer) Get(context.Context, *GetPaymentRequest) (*GetPaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
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

func _Payment_CreatePaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).CreatePaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_CreatePaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).CreatePaymentMethod(ctx, req.(*CreatePaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetPaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetPaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_GetPaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetPaymentMethod(ctx, req.(*GetPaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_UpdatePaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).UpdatePaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_UpdatePaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).UpdatePaymentMethod(ctx, req.(*UpdatePaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_DeletePaymentMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePaymentMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).DeletePaymentMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_DeletePaymentMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).DeletePaymentMethod(ctx, req.(*DeletePaymentMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
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

func _Payment_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Payment_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).Get(ctx, req.(*GetPaymentRequest))
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
			MethodName: "CreatePaymentMethod",
			Handler:    _Payment_CreatePaymentMethod_Handler,
		},
		{
			MethodName: "GetPaymentMethod",
			Handler:    _Payment_GetPaymentMethod_Handler,
		},
		{
			MethodName: "UpdatePaymentMethod",
			Handler:    _Payment_UpdatePaymentMethod_Handler,
		},
		{
			MethodName: "DeletePaymentMethod",
			Handler:    _Payment_DeletePaymentMethod_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Payment_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Payment_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment/payment.proto",
}

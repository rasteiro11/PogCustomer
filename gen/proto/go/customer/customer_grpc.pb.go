// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: customer/customer.proto

package customer

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
	CustomerService_GetUser_FullMethodName           = "/pog.customer.CustomerService/GetUser"
	CustomerService_GetUserByDocument_FullMethodName = "/pog.customer.CustomerService/GetUserByDocument"
	CustomerService_VerifySession_FullMethodName     = "/pog.customer.CustomerService/VerifySession"
)

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUserByDocument(ctx context.Context, in *GetUserByDocumentRequest, opts ...grpc.CallOption) (*GetUserByDocumentResponse, error)
	VerifySession(ctx context.Context, in *VerifySessionRequest, opts ...grpc.CallOption) (*VerifySessionResponse, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, CustomerService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) GetUserByDocument(ctx context.Context, in *GetUserByDocumentRequest, opts ...grpc.CallOption) (*GetUserByDocumentResponse, error) {
	out := new(GetUserByDocumentResponse)
	err := c.cc.Invoke(ctx, CustomerService_GetUserByDocument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) VerifySession(ctx context.Context, in *VerifySessionRequest, opts ...grpc.CallOption) (*VerifySessionResponse, error) {
	out := new(VerifySessionResponse)
	err := c.cc.Invoke(ctx, CustomerService_VerifySession_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations should embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUserByDocument(context.Context, *GetUserByDocumentRequest) (*GetUserByDocumentResponse, error)
	VerifySession(context.Context, *VerifySessionRequest) (*VerifySessionResponse, error)
}

// UnimplementedCustomerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCustomerServiceServer) GetUserByDocument(context.Context, *GetUserByDocumentRequest) (*GetUserByDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByDocument not implemented")
}
func (UnimplementedCustomerServiceServer) VerifySession(context.Context, *VerifySessionRequest) (*VerifySessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySession not implemented")
}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_GetUserByDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).GetUserByDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_GetUserByDocument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).GetUserByDocument(ctx, req.(*GetUserByDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_VerifySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).VerifySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CustomerService_VerifySession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).VerifySession(ctx, req.(*VerifySessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pog.customer.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _CustomerService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByDocument",
			Handler:    _CustomerService_GetUserByDocument_Handler,
		},
		{
			MethodName: "VerifySession",
			Handler:    _CustomerService_VerifySession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer/customer.proto",
}

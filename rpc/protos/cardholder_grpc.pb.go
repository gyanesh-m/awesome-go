// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: protos/cardholder.proto

package v1

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

// CardHolderServiceClient is the client API for CardHolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardHolderServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
}

type cardHolderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardHolderServiceClient(cc grpc.ClientConnInterface) CardHolderServiceClient {
	return &cardHolderServiceClient{cc}
}

func (c *cardHolderServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/cardholders.v1.CardHolderService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardHolderServiceClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/cardholders.v1.CardHolderService/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardHolderServiceServer is the server API for CardHolderService service.
// All implementations must embed UnimplementedCardHolderServiceServer
// for forward compatibility
type CardHolderServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	mustEmbedUnimplementedCardHolderServiceServer()
}

// UnimplementedCardHolderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardHolderServiceServer struct {
}

func (UnimplementedCardHolderServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedCardHolderServiceServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedCardHolderServiceServer) mustEmbedUnimplementedCardHolderServiceServer() {}

// UnsafeCardHolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardHolderServiceServer will
// result in compilation errors.
type UnsafeCardHolderServiceServer interface {
	mustEmbedUnimplementedCardHolderServiceServer()
}

func RegisterCardHolderServiceServer(s grpc.ServiceRegistrar, srv CardHolderServiceServer) {
	s.RegisterService(&CardHolderService_ServiceDesc, srv)
}

func _CardHolderService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardHolderServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cardholders.v1.CardHolderService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardHolderServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardHolderService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardHolderServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cardholders.v1.CardHolderService/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardHolderServiceServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardHolderService_ServiceDesc is the grpc.ServiceDesc for CardHolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardHolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cardholders.v1.CardHolderService",
	HandlerType: (*CardHolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _CardHolderService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _CardHolderService_GetAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/cardholder.proto",
}

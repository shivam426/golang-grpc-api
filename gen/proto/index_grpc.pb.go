// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// CorpapiClient is the client API for Corpapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CorpapiClient interface {
	CreateUser(ctx context.Context, in *CreateRequestResponse, opts ...grpc.CallOption) (*CreateRequestResponse, error)
	FetchUser(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type corpapiClient struct {
	cc grpc.ClientConnInterface
}

func NewCorpapiClient(cc grpc.ClientConnInterface) CorpapiClient {
	return &corpapiClient{cc}
}

func (c *corpapiClient) CreateUser(ctx context.Context, in *CreateRequestResponse, opts ...grpc.CallOption) (*CreateRequestResponse, error) {
	out := new(CreateRequestResponse)
	err := c.cc.Invoke(ctx, "/main.corpapi/createUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *corpapiClient) FetchUser(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/main.corpapi/fetchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CorpapiServer is the server API for Corpapi service.
// All implementations must embed UnimplementedCorpapiServer
// for forward compatibility
type CorpapiServer interface {
	CreateUser(context.Context, *CreateRequestResponse) (*CreateRequestResponse, error)
	FetchUser(context.Context, *ReadRequest) (*ReadResponse, error)
	mustEmbedUnimplementedCorpapiServer()
}

// UnimplementedCorpapiServer must be embedded to have forward compatible implementations.
type UnimplementedCorpapiServer struct {
}

func (UnimplementedCorpapiServer) CreateUser(context.Context, *CreateRequestResponse) (*CreateRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedCorpapiServer) FetchUser(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchUser not implemented")
}
func (UnimplementedCorpapiServer) mustEmbedUnimplementedCorpapiServer() {}

// UnsafeCorpapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CorpapiServer will
// result in compilation errors.
type UnsafeCorpapiServer interface {
	mustEmbedUnimplementedCorpapiServer()
}

func RegisterCorpapiServer(s grpc.ServiceRegistrar, srv CorpapiServer) {
	s.RegisterService(&Corpapi_ServiceDesc, srv)
}

func _Corpapi_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorpapiServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.corpapi/createUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorpapiServer).CreateUser(ctx, req.(*CreateRequestResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _Corpapi_FetchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CorpapiServer).FetchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.corpapi/fetchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CorpapiServer).FetchUser(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Corpapi_ServiceDesc is the grpc.ServiceDesc for Corpapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Corpapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.corpapi",
	HandlerType: (*CorpapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createUser",
			Handler:    _Corpapi_CreateUser_Handler,
		},
		{
			MethodName: "fetchUser",
			Handler:    _Corpapi_FetchUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index.proto",
}

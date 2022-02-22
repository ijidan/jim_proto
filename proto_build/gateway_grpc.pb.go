// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: gateway.proto

package proto_build

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

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	UnRegister(ctx context.Context, in *UnRegisterRequest, opts ...grpc.CallOption) (*UnRegisterResponse, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (GatewayService_SendMessageClient, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/message.GatewayService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) UnRegister(ctx context.Context, in *UnRegisterRequest, opts ...grpc.CallOption) (*UnRegisterResponse, error) {
	out := new(UnRegisterResponse)
	err := c.cc.Invoke(ctx, "/message.GatewayService/UnRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (GatewayService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &GatewayService_ServiceDesc.Streams[0], "/message.GatewayService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayServiceSendMessageClient{stream}
	return x, nil
}

type GatewayService_SendMessageClient interface {
	Send(*SendMessageRequest) error
	Recv() (*SendMessageResponse, error)
	grpc.ClientStream
}

type gatewayServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *gatewayServiceSendMessageClient) Send(m *SendMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayServiceSendMessageClient) Recv() (*SendMessageResponse, error) {
	m := new(SendMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UnRegister(context.Context, *UnRegisterRequest) (*UnRegisterResponse, error)
	SendMessage(GatewayService_SendMessageServer) error
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGatewayServiceServer) UnRegister(context.Context, *UnRegisterRequest) (*UnRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnRegister not implemented")
}
func (UnimplementedGatewayServiceServer) SendMessage(GatewayService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.GatewayService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_UnRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).UnRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.GatewayService/UnRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).UnRegister(ctx, req.(*UnRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayServiceServer).SendMessage(&gatewayServiceSendMessageServer{stream})
}

type GatewayService_SendMessageServer interface {
	Send(*SendMessageResponse) error
	Recv() (*SendMessageRequest, error)
	grpc.ServerStream
}

type gatewayServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *gatewayServiceSendMessageServer) Send(m *SendMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayServiceSendMessageServer) Recv() (*SendMessageRequest, error) {
	m := new(SendMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GatewayService_Register_Handler,
		},
		{
			MethodName: "UnRegister",
			Handler:    _GatewayService_UnRegister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMessage",
			Handler:       _GatewayService_SendMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gateway.proto",
}

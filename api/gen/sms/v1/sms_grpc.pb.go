// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sms/v1/sms.proto

package smsv1

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
	SmsService_Send_FullMethodName = "/sms.v1.SmsService/Send"
)

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, SmsService_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations must embed UnimplementedSmsServiceServer
// for forward compatibility
type SmsServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	mustEmbedUnimplementedSmsServiceServer()
}

// UnimplementedSmsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmsServiceServer struct {
}

func (UnimplementedSmsServiceServer) Send(context.Context, *SendRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedSmsServiceServer) mustEmbedUnimplementedSmsServiceServer() {}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	s.RegisterService(&SmsService_ServiceDesc, srv)
}

func _SmsService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsService_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsService_ServiceDesc is the grpc.ServiceDesc for SmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sms.v1.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _SmsService_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sms/v1/sms.proto",
}

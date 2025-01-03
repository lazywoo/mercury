// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: crontask/v1/crontask.proto

package crontaskv1

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
	CronTaskService_Preempt_FullMethodName       = "/crontask.v1.CronTaskService/Preempt"
	CronTaskService_ResetNextTime_FullMethodName = "/crontask.v1.CronTaskService/ResetNextTime"
	CronTaskService_AddTask_FullMethodName       = "/crontask.v1.CronTaskService/AddTask"
)

// CronTaskServiceClient is the client API for CronTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CronTaskServiceClient interface {
	Preempt(ctx context.Context, in *PreemptRequest, opts ...grpc.CallOption) (*PreemptResponse, error)
	ResetNextTime(ctx context.Context, in *ResetNextTimeRequest, opts ...grpc.CallOption) (*ResetNextTimeResponse, error)
	AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error)
}

type cronTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCronTaskServiceClient(cc grpc.ClientConnInterface) CronTaskServiceClient {
	return &cronTaskServiceClient{cc}
}

func (c *cronTaskServiceClient) Preempt(ctx context.Context, in *PreemptRequest, opts ...grpc.CallOption) (*PreemptResponse, error) {
	out := new(PreemptResponse)
	err := c.cc.Invoke(ctx, CronTaskService_Preempt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronTaskServiceClient) ResetNextTime(ctx context.Context, in *ResetNextTimeRequest, opts ...grpc.CallOption) (*ResetNextTimeResponse, error) {
	out := new(ResetNextTimeResponse)
	err := c.cc.Invoke(ctx, CronTaskService_ResetNextTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronTaskServiceClient) AddTask(ctx context.Context, in *AddTaskRequest, opts ...grpc.CallOption) (*AddTaskResponse, error) {
	out := new(AddTaskResponse)
	err := c.cc.Invoke(ctx, CronTaskService_AddTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronTaskServiceServer is the server API for CronTaskService service.
// All implementations must embed UnimplementedCronTaskServiceServer
// for forward compatibility
type CronTaskServiceServer interface {
	Preempt(context.Context, *PreemptRequest) (*PreemptResponse, error)
	ResetNextTime(context.Context, *ResetNextTimeRequest) (*ResetNextTimeResponse, error)
	AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error)
	mustEmbedUnimplementedCronTaskServiceServer()
}

// UnimplementedCronTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCronTaskServiceServer struct {
}

func (UnimplementedCronTaskServiceServer) Preempt(context.Context, *PreemptRequest) (*PreemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preempt not implemented")
}
func (UnimplementedCronTaskServiceServer) ResetNextTime(context.Context, *ResetNextTimeRequest) (*ResetNextTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetNextTime not implemented")
}
func (UnimplementedCronTaskServiceServer) AddTask(context.Context, *AddTaskRequest) (*AddTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTask not implemented")
}
func (UnimplementedCronTaskServiceServer) mustEmbedUnimplementedCronTaskServiceServer() {}

// UnsafeCronTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CronTaskServiceServer will
// result in compilation errors.
type UnsafeCronTaskServiceServer interface {
	mustEmbedUnimplementedCronTaskServiceServer()
}

func RegisterCronTaskServiceServer(s grpc.ServiceRegistrar, srv CronTaskServiceServer) {
	s.RegisterService(&CronTaskService_ServiceDesc, srv)
}

func _CronTaskService_Preempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronTaskServiceServer).Preempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronTaskService_Preempt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronTaskServiceServer).Preempt(ctx, req.(*PreemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronTaskService_ResetNextTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetNextTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronTaskServiceServer).ResetNextTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronTaskService_ResetNextTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronTaskServiceServer).ResetNextTime(ctx, req.(*ResetNextTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronTaskService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronTaskServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CronTaskService_AddTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronTaskServiceServer).AddTask(ctx, req.(*AddTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CronTaskService_ServiceDesc is the grpc.ServiceDesc for CronTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CronTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crontask.v1.CronTaskService",
	HandlerType: (*CronTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Preempt",
			Handler:    _CronTaskService_Preempt_Handler,
		},
		{
			MethodName: "ResetNextTime",
			Handler:    _CronTaskService_ResetNextTime_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _CronTaskService_AddTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crontask/v1/crontask.proto",
}

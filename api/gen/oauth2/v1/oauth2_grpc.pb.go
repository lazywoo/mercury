// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: oauth2/v1/oauth2.proto

package oauth2v1

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
	Oauth2Service_AuthURL_FullMethodName    = "/oauth2.v1.Oauth2Service/AuthURL"
	Oauth2Service_VerifyCode_FullMethodName = "/oauth2.v1.Oauth2Service/VerifyCode"
)

// Oauth2ServiceClient is the client API for Oauth2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Oauth2ServiceClient interface {
	AuthURL(ctx context.Context, in *AuthURLRequest, opts ...grpc.CallOption) (*AuthURLResponse, error)
	VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error)
}

type oauth2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOauth2ServiceClient(cc grpc.ClientConnInterface) Oauth2ServiceClient {
	return &oauth2ServiceClient{cc}
}

func (c *oauth2ServiceClient) AuthURL(ctx context.Context, in *AuthURLRequest, opts ...grpc.CallOption) (*AuthURLResponse, error) {
	out := new(AuthURLResponse)
	err := c.cc.Invoke(ctx, Oauth2Service_AuthURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2ServiceClient) VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error) {
	out := new(VerifyCodeResponse)
	err := c.cc.Invoke(ctx, Oauth2Service_VerifyCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Oauth2ServiceServer is the server API for Oauth2Service service.
// All implementations must embed UnimplementedOauth2ServiceServer
// for forward compatibility
type Oauth2ServiceServer interface {
	AuthURL(context.Context, *AuthURLRequest) (*AuthURLResponse, error)
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error)
	mustEmbedUnimplementedOauth2ServiceServer()
}

// UnimplementedOauth2ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOauth2ServiceServer struct {
}

func (UnimplementedOauth2ServiceServer) AuthURL(context.Context, *AuthURLRequest) (*AuthURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthURL not implemented")
}
func (UnimplementedOauth2ServiceServer) VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}
func (UnimplementedOauth2ServiceServer) mustEmbedUnimplementedOauth2ServiceServer() {}

// UnsafeOauth2ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Oauth2ServiceServer will
// result in compilation errors.
type UnsafeOauth2ServiceServer interface {
	mustEmbedUnimplementedOauth2ServiceServer()
}

func RegisterOauth2ServiceServer(s grpc.ServiceRegistrar, srv Oauth2ServiceServer) {
	s.RegisterService(&Oauth2Service_ServiceDesc, srv)
}

func _Oauth2Service_AuthURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2ServiceServer).AuthURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Service_AuthURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2ServiceServer).AuthURL(ctx, req.(*AuthURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2Service_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2ServiceServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Oauth2Service_VerifyCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2ServiceServer).VerifyCode(ctx, req.(*VerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth2Service_ServiceDesc is the grpc.ServiceDesc for Oauth2Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth2Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth2.v1.Oauth2Service",
	HandlerType: (*Oauth2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthURL",
			Handler:    _Oauth2Service_AuthURL_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _Oauth2Service_VerifyCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth2/v1/oauth2.proto",
}

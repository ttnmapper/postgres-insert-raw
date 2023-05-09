// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: lorawan-stack/api/identityserver.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EntityAccessClient is the client API for EntityAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EntityAccessClient interface {
	// AuthInfo returns information about the authentication that is used on the request.
	AuthInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthInfoResponse, error)
}

type entityAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityAccessClient(cc grpc.ClientConnInterface) EntityAccessClient {
	return &entityAccessClient{cc}
}

func (c *entityAccessClient) AuthInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthInfoResponse, error) {
	out := new(AuthInfoResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EntityAccess/AuthInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityAccessServer is the server API for EntityAccess service.
// All implementations must embed UnimplementedEntityAccessServer
// for forward compatibility
type EntityAccessServer interface {
	// AuthInfo returns information about the authentication that is used on the request.
	AuthInfo(context.Context, *emptypb.Empty) (*AuthInfoResponse, error)
	mustEmbedUnimplementedEntityAccessServer()
}

// UnimplementedEntityAccessServer must be embedded to have forward compatible implementations.
type UnimplementedEntityAccessServer struct {
}

func (UnimplementedEntityAccessServer) AuthInfo(context.Context, *emptypb.Empty) (*AuthInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthInfo not implemented")
}
func (UnimplementedEntityAccessServer) mustEmbedUnimplementedEntityAccessServer() {}

// UnsafeEntityAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EntityAccessServer will
// result in compilation errors.
type UnsafeEntityAccessServer interface {
	mustEmbedUnimplementedEntityAccessServer()
}

func RegisterEntityAccessServer(s grpc.ServiceRegistrar, srv EntityAccessServer) {
	s.RegisterService(&EntityAccess_ServiceDesc, srv)
}

func _EntityAccess_AuthInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityAccessServer).AuthInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EntityAccess/AuthInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityAccessServer).AuthInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// EntityAccess_ServiceDesc is the grpc.ServiceDesc for EntityAccess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EntityAccess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EntityAccess",
	HandlerType: (*EntityAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthInfo",
			Handler:    _EntityAccess_AuthInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/identityserver.proto",
}

// IsClient is the client API for Is service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IsClient interface {
	// Get the configuration of the Identity Server. The response is typically used
	// to enable or disable features in a user interface.
	GetConfiguration(ctx context.Context, in *GetIsConfigurationRequest, opts ...grpc.CallOption) (*GetIsConfigurationResponse, error)
}

type isClient struct {
	cc grpc.ClientConnInterface
}

func NewIsClient(cc grpc.ClientConnInterface) IsClient {
	return &isClient{cc}
}

func (c *isClient) GetConfiguration(ctx context.Context, in *GetIsConfigurationRequest, opts ...grpc.CallOption) (*GetIsConfigurationResponse, error) {
	out := new(GetIsConfigurationResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Is/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsServer is the server API for Is service.
// All implementations must embed UnimplementedIsServer
// for forward compatibility
type IsServer interface {
	// Get the configuration of the Identity Server. The response is typically used
	// to enable or disable features in a user interface.
	GetConfiguration(context.Context, *GetIsConfigurationRequest) (*GetIsConfigurationResponse, error)
	mustEmbedUnimplementedIsServer()
}

// UnimplementedIsServer must be embedded to have forward compatible implementations.
type UnimplementedIsServer struct {
}

func (UnimplementedIsServer) GetConfiguration(context.Context, *GetIsConfigurationRequest) (*GetIsConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedIsServer) mustEmbedUnimplementedIsServer() {}

// UnsafeIsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IsServer will
// result in compilation errors.
type UnsafeIsServer interface {
	mustEmbedUnimplementedIsServer()
}

func RegisterIsServer(s grpc.ServiceRegistrar, srv IsServer) {
	s.RegisterService(&Is_ServiceDesc, srv)
}

func _Is_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIsConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IsServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Is/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IsServer).GetConfiguration(ctx, req.(*GetIsConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Is_ServiceDesc is the grpc.ServiceDesc for Is service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Is_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Is",
	HandlerType: (*IsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfiguration",
			Handler:    _Is_GetConfiguration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/identityserver.proto",
}

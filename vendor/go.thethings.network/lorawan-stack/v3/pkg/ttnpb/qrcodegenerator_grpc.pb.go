// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: lorawan-stack/api/qrcodegenerator.proto

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

// EndDeviceQRCodeGeneratorClient is the client API for EndDeviceQRCodeGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndDeviceQRCodeGeneratorClient interface {
	// Return the QR code format.
	GetFormat(ctx context.Context, in *GetQRCodeFormatRequest, opts ...grpc.CallOption) (*QRCodeFormat, error)
	// Returns the supported formats.
	ListFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*QRCodeFormats, error)
	// Generates a QR code.
	Generate(ctx context.Context, in *GenerateEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*GenerateQRCodeResponse, error)
	// Parse QR Codes of known formats and return the information contained within.
	Parse(ctx context.Context, in *ParseEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*ParseEndDeviceQRCodeResponse, error)
}

type endDeviceQRCodeGeneratorClient struct {
	cc grpc.ClientConnInterface
}

func NewEndDeviceQRCodeGeneratorClient(cc grpc.ClientConnInterface) EndDeviceQRCodeGeneratorClient {
	return &endDeviceQRCodeGeneratorClient{cc}
}

func (c *endDeviceQRCodeGeneratorClient) GetFormat(ctx context.Context, in *GetQRCodeFormatRequest, opts ...grpc.CallOption) (*QRCodeFormat, error) {
	out := new(QRCodeFormat)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/GetFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceQRCodeGeneratorClient) ListFormats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*QRCodeFormats, error) {
	out := new(QRCodeFormats)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/ListFormats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceQRCodeGeneratorClient) Generate(ctx context.Context, in *GenerateEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*GenerateQRCodeResponse, error) {
	out := new(GenerateQRCodeResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endDeviceQRCodeGeneratorClient) Parse(ctx context.Context, in *ParseEndDeviceQRCodeRequest, opts ...grpc.CallOption) (*ParseEndDeviceQRCodeResponse, error) {
	out := new(ParseEndDeviceQRCodeResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Parse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndDeviceQRCodeGeneratorServer is the server API for EndDeviceQRCodeGenerator service.
// All implementations must embed UnimplementedEndDeviceQRCodeGeneratorServer
// for forward compatibility
type EndDeviceQRCodeGeneratorServer interface {
	// Return the QR code format.
	GetFormat(context.Context, *GetQRCodeFormatRequest) (*QRCodeFormat, error)
	// Returns the supported formats.
	ListFormats(context.Context, *emptypb.Empty) (*QRCodeFormats, error)
	// Generates a QR code.
	Generate(context.Context, *GenerateEndDeviceQRCodeRequest) (*GenerateQRCodeResponse, error)
	// Parse QR Codes of known formats and return the information contained within.
	Parse(context.Context, *ParseEndDeviceQRCodeRequest) (*ParseEndDeviceQRCodeResponse, error)
	mustEmbedUnimplementedEndDeviceQRCodeGeneratorServer()
}

// UnimplementedEndDeviceQRCodeGeneratorServer must be embedded to have forward compatible implementations.
type UnimplementedEndDeviceQRCodeGeneratorServer struct {
}

func (UnimplementedEndDeviceQRCodeGeneratorServer) GetFormat(context.Context, *GetQRCodeFormatRequest) (*QRCodeFormat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFormat not implemented")
}
func (UnimplementedEndDeviceQRCodeGeneratorServer) ListFormats(context.Context, *emptypb.Empty) (*QRCodeFormats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFormats not implemented")
}
func (UnimplementedEndDeviceQRCodeGeneratorServer) Generate(context.Context, *GenerateEndDeviceQRCodeRequest) (*GenerateQRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedEndDeviceQRCodeGeneratorServer) Parse(context.Context, *ParseEndDeviceQRCodeRequest) (*ParseEndDeviceQRCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Parse not implemented")
}
func (UnimplementedEndDeviceQRCodeGeneratorServer) mustEmbedUnimplementedEndDeviceQRCodeGeneratorServer() {
}

// UnsafeEndDeviceQRCodeGeneratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndDeviceQRCodeGeneratorServer will
// result in compilation errors.
type UnsafeEndDeviceQRCodeGeneratorServer interface {
	mustEmbedUnimplementedEndDeviceQRCodeGeneratorServer()
}

func RegisterEndDeviceQRCodeGeneratorServer(s grpc.ServiceRegistrar, srv EndDeviceQRCodeGeneratorServer) {
	s.RegisterService(&EndDeviceQRCodeGenerator_ServiceDesc, srv)
}

func _EndDeviceQRCodeGenerator_GetFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQRCodeFormatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).GetFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/GetFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).GetFormat(ctx, req.(*GetQRCodeFormatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceQRCodeGenerator_ListFormats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).ListFormats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/ListFormats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).ListFormats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceQRCodeGenerator_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateEndDeviceQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).Generate(ctx, req.(*GenerateEndDeviceQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EndDeviceQRCodeGenerator_Parse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseEndDeviceQRCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndDeviceQRCodeGeneratorServer).Parse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.EndDeviceQRCodeGenerator/Parse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndDeviceQRCodeGeneratorServer).Parse(ctx, req.(*ParseEndDeviceQRCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EndDeviceQRCodeGenerator_ServiceDesc is the grpc.ServiceDesc for EndDeviceQRCodeGenerator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EndDeviceQRCodeGenerator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.EndDeviceQRCodeGenerator",
	HandlerType: (*EndDeviceQRCodeGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFormat",
			Handler:    _EndDeviceQRCodeGenerator_GetFormat_Handler,
		},
		{
			MethodName: "ListFormats",
			Handler:    _EndDeviceQRCodeGenerator_ListFormats_Handler,
		},
		{
			MethodName: "Generate",
			Handler:    _EndDeviceQRCodeGenerator_Generate_Handler,
		},
		{
			MethodName: "Parse",
			Handler:    _EndDeviceQRCodeGenerator_Parse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/qrcodegenerator.proto",
}

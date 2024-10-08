// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: metalstack/io/accounting/api/v1/volume.proto

package v1

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
	VolumeService_Added_FullMethodName    = "/metalstack.io.accounting.api.v1.VolumeService/Added"
	VolumeService_Modified_FullMethodName = "/metalstack.io.accounting.api.v1.VolumeService/Modified"
	VolumeService_Deleted_FullMethodName  = "/metalstack.io.accounting.api.v1.VolumeService/Deleted"
	VolumeService_Usage_FullMethodName    = "/metalstack.io.accounting.api.v1.VolumeService/Usage"
)

// VolumeServiceClient is the client API for VolumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolumeServiceClient interface {
	Added(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error)
	Modified(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error)
	Deleted(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error)
	Usage(ctx context.Context, in *VolumeUsageRequest, opts ...grpc.CallOption) (*VolumeUsageResponse, error)
}

type volumeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumeServiceClient(cc grpc.ClientConnInterface) VolumeServiceClient {
	return &volumeServiceClient{cc}
}

func (c *volumeServiceClient) Added(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, VolumeService_Added_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) Modified(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, VolumeService_Modified_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) Deleted(ctx context.Context, in *VolumeReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, VolumeService_Deleted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) Usage(ctx context.Context, in *VolumeUsageRequest, opts ...grpc.CallOption) (*VolumeUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VolumeUsageResponse)
	err := c.cc.Invoke(ctx, VolumeService_Usage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeServiceServer is the server API for VolumeService service.
// All implementations should embed UnimplementedVolumeServiceServer
// for forward compatibility.
type VolumeServiceServer interface {
	Added(context.Context, *VolumeReport) (*Empty, error)
	Modified(context.Context, *VolumeReport) (*Empty, error)
	Deleted(context.Context, *VolumeReport) (*Empty, error)
	Usage(context.Context, *VolumeUsageRequest) (*VolumeUsageResponse, error)
}

// UnimplementedVolumeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVolumeServiceServer struct{}

func (UnimplementedVolumeServiceServer) Added(context.Context, *VolumeReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Added not implemented")
}
func (UnimplementedVolumeServiceServer) Modified(context.Context, *VolumeReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modified not implemented")
}
func (UnimplementedVolumeServiceServer) Deleted(context.Context, *VolumeReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deleted not implemented")
}
func (UnimplementedVolumeServiceServer) Usage(context.Context, *VolumeUsageRequest) (*VolumeUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}
func (UnimplementedVolumeServiceServer) testEmbeddedByValue() {}

// UnsafeVolumeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolumeServiceServer will
// result in compilation errors.
type UnsafeVolumeServiceServer interface {
	mustEmbedUnimplementedVolumeServiceServer()
}

func RegisterVolumeServiceServer(s grpc.ServiceRegistrar, srv VolumeServiceServer) {
	// If the following call pancis, it indicates UnimplementedVolumeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VolumeService_ServiceDesc, srv)
}

func _VolumeService_Added_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).Added(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_Added_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).Added(ctx, req.(*VolumeReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_Modified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).Modified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_Modified_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).Modified(ctx, req.(*VolumeReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_Deleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).Deleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_Deleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).Deleted(ctx, req.(*VolumeReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_Usage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).Usage(ctx, req.(*VolumeUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VolumeService_ServiceDesc is the grpc.ServiceDesc for VolumeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolumeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.accounting.api.v1.VolumeService",
	HandlerType: (*VolumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Added",
			Handler:    _VolumeService_Added_Handler,
		},
		{
			MethodName: "Modified",
			Handler:    _VolumeService_Modified_Handler,
		},
		{
			MethodName: "Deleted",
			Handler:    _VolumeService_Deleted_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _VolumeService_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/accounting/api/v1/volume.proto",
}

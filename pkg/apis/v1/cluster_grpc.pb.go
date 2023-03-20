// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: metalstack/io/accounting/api/v1/cluster.proto

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

const (
	ClusterService_Added_FullMethodName    = "/metalstack.io.accounting.api.v1.ClusterService/Added"
	ClusterService_Modified_FullMethodName = "/metalstack.io.accounting.api.v1.ClusterService/Modified"
	ClusterService_Deleted_FullMethodName  = "/metalstack.io.accounting.api.v1.ClusterService/Deleted"
	ClusterService_Usage_FullMethodName    = "/metalstack.io.accounting.api.v1.ClusterService/Usage"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	Added(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error)
	Modified(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error)
	Deleted(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error)
	Usage(ctx context.Context, in *ClusterUsageRequest, opts ...grpc.CallOption) (*ClusterUsageResponse, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Added(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ClusterService_Added_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Modified(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ClusterService_Modified_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Deleted(ctx context.Context, in *ClusterReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ClusterService_Deleted_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Usage(ctx context.Context, in *ClusterUsageRequest, opts ...grpc.CallOption) (*ClusterUsageResponse, error) {
	out := new(ClusterUsageResponse)
	err := c.cc.Invoke(ctx, ClusterService_Usage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	Added(context.Context, *ClusterReport) (*Empty, error)
	Modified(context.Context, *ClusterReport) (*Empty, error)
	Deleted(context.Context, *ClusterReport) (*Empty, error)
	Usage(context.Context, *ClusterUsageRequest) (*ClusterUsageResponse, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

func (UnimplementedClusterServiceServer) Added(context.Context, *ClusterReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Added not implemented")
}
func (UnimplementedClusterServiceServer) Modified(context.Context, *ClusterReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modified not implemented")
}
func (UnimplementedClusterServiceServer) Deleted(context.Context, *ClusterReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deleted not implemented")
}
func (UnimplementedClusterServiceServer) Usage(context.Context, *ClusterUsageRequest) (*ClusterUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_Added_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Added(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Added_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Added(ctx, req.(*ClusterReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Modified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Modified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Modified_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Modified(ctx, req.(*ClusterReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Deleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Deleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Deleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Deleted(ctx, req.(*ClusterReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Usage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Usage(ctx, req.(*ClusterUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.accounting.api.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Added",
			Handler:    _ClusterService_Added_Handler,
		},
		{
			MethodName: "Modified",
			Handler:    _ClusterService_Modified_Handler,
		},
		{
			MethodName: "Deleted",
			Handler:    _ClusterService_Deleted_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _ClusterService_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/accounting/api/v1/cluster.proto",
}

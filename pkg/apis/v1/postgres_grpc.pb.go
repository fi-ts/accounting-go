// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PostgresServiceClient is the client API for PostgresService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostgresServiceClient interface {
	Added(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error)
	Modified(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error)
	Deleted(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error)
	Usage(ctx context.Context, in *PostgresUsageRequest, opts ...grpc.CallOption) (*PostgresUsageResponse, error)
}

type postgresServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostgresServiceClient(cc grpc.ClientConnInterface) PostgresServiceClient {
	return &postgresServiceClient{cc}
}

func (c *postgresServiceClient) Added(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/metalstack.io.accounting.api.v1.PostgresService/Added", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresServiceClient) Modified(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/metalstack.io.accounting.api.v1.PostgresService/Modified", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresServiceClient) Deleted(ctx context.Context, in *PostgresReport, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/metalstack.io.accounting.api.v1.PostgresService/Deleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgresServiceClient) Usage(ctx context.Context, in *PostgresUsageRequest, opts ...grpc.CallOption) (*PostgresUsageResponse, error) {
	out := new(PostgresUsageResponse)
	err := c.cc.Invoke(ctx, "/metalstack.io.accounting.api.v1.PostgresService/Usage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgresServiceServer is the server API for PostgresService service.
// All implementations should embed UnimplementedPostgresServiceServer
// for forward compatibility
type PostgresServiceServer interface {
	Added(context.Context, *PostgresReport) (*Empty, error)
	Modified(context.Context, *PostgresReport) (*Empty, error)
	Deleted(context.Context, *PostgresReport) (*Empty, error)
	Usage(context.Context, *PostgresUsageRequest) (*PostgresUsageResponse, error)
}

// UnimplementedPostgresServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPostgresServiceServer struct {
}

func (UnimplementedPostgresServiceServer) Added(context.Context, *PostgresReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Added not implemented")
}
func (UnimplementedPostgresServiceServer) Modified(context.Context, *PostgresReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modified not implemented")
}
func (UnimplementedPostgresServiceServer) Deleted(context.Context, *PostgresReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deleted not implemented")
}
func (UnimplementedPostgresServiceServer) Usage(context.Context, *PostgresUsageRequest) (*PostgresUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}

// UnsafePostgresServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostgresServiceServer will
// result in compilation errors.
type UnsafePostgresServiceServer interface {
	mustEmbedUnimplementedPostgresServiceServer()
}

func RegisterPostgresServiceServer(s grpc.ServiceRegistrar, srv PostgresServiceServer) {
	s.RegisterService(&PostgresService_ServiceDesc, srv)
}

func _PostgresService_Added_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).Added(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.accounting.api.v1.PostgresService/Added",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).Added(ctx, req.(*PostgresReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresService_Modified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).Modified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.accounting.api.v1.PostgresService/Modified",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).Modified(ctx, req.(*PostgresReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresService_Deleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).Deleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.accounting.api.v1.PostgresService/Deleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).Deleted(ctx, req.(*PostgresReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgresService_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgresUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgresServiceServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metalstack.io.accounting.api.v1.PostgresService/Usage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgresServiceServer).Usage(ctx, req.(*PostgresUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostgresService_ServiceDesc is the grpc.ServiceDesc for PostgresService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostgresService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.accounting.api.v1.PostgresService",
	HandlerType: (*PostgresServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Added",
			Handler:    _PostgresService_Added_Handler,
		},
		{
			MethodName: "Modified",
			Handler:    _PostgresService_Modified_Handler,
		},
		{
			MethodName: "Deleted",
			Handler:    _PostgresService_Deleted_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _PostgresService_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/accounting/api/v1/postgres.proto",
}
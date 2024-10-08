// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: metalstack/io/accounting/api/v1/machine_reservation.proto

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
	MachineReservationService_Added_FullMethodName    = "/metalstack.io.accounting.api.v1.MachineReservationService/Added"
	MachineReservationService_Modified_FullMethodName = "/metalstack.io.accounting.api.v1.MachineReservationService/Modified"
	MachineReservationService_Deleted_FullMethodName  = "/metalstack.io.accounting.api.v1.MachineReservationService/Deleted"
	MachineReservationService_Usage_FullMethodName    = "/metalstack.io.accounting.api.v1.MachineReservationService/Usage"
)

// MachineReservationServiceClient is the client API for MachineReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineReservationServiceClient interface {
	Added(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error)
	Modified(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error)
	Deleted(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error)
	Usage(ctx context.Context, in *MachineReservationUsageRequest, opts ...grpc.CallOption) (*MachineReservationUsageResponse, error)
}

type machineReservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineReservationServiceClient(cc grpc.ClientConnInterface) MachineReservationServiceClient {
	return &machineReservationServiceClient{cc}
}

func (c *machineReservationServiceClient) Added(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MachineReservationService_Added_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineReservationServiceClient) Modified(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MachineReservationService_Modified_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineReservationServiceClient) Deleted(ctx context.Context, in *MachineReservationReport, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, MachineReservationService_Deleted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineReservationServiceClient) Usage(ctx context.Context, in *MachineReservationUsageRequest, opts ...grpc.CallOption) (*MachineReservationUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MachineReservationUsageResponse)
	err := c.cc.Invoke(ctx, MachineReservationService_Usage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineReservationServiceServer is the server API for MachineReservationService service.
// All implementations should embed UnimplementedMachineReservationServiceServer
// for forward compatibility.
type MachineReservationServiceServer interface {
	Added(context.Context, *MachineReservationReport) (*Empty, error)
	Modified(context.Context, *MachineReservationReport) (*Empty, error)
	Deleted(context.Context, *MachineReservationReport) (*Empty, error)
	Usage(context.Context, *MachineReservationUsageRequest) (*MachineReservationUsageResponse, error)
}

// UnimplementedMachineReservationServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMachineReservationServiceServer struct{}

func (UnimplementedMachineReservationServiceServer) Added(context.Context, *MachineReservationReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Added not implemented")
}
func (UnimplementedMachineReservationServiceServer) Modified(context.Context, *MachineReservationReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modified not implemented")
}
func (UnimplementedMachineReservationServiceServer) Deleted(context.Context, *MachineReservationReport) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deleted not implemented")
}
func (UnimplementedMachineReservationServiceServer) Usage(context.Context, *MachineReservationUsageRequest) (*MachineReservationUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Usage not implemented")
}
func (UnimplementedMachineReservationServiceServer) testEmbeddedByValue() {}

// UnsafeMachineReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineReservationServiceServer will
// result in compilation errors.
type UnsafeMachineReservationServiceServer interface {
	mustEmbedUnimplementedMachineReservationServiceServer()
}

func RegisterMachineReservationServiceServer(s grpc.ServiceRegistrar, srv MachineReservationServiceServer) {
	// If the following call pancis, it indicates UnimplementedMachineReservationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MachineReservationService_ServiceDesc, srv)
}

func _MachineReservationService_Added_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineReservationReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineReservationServiceServer).Added(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineReservationService_Added_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineReservationServiceServer).Added(ctx, req.(*MachineReservationReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineReservationService_Modified_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineReservationReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineReservationServiceServer).Modified(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineReservationService_Modified_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineReservationServiceServer).Modified(ctx, req.(*MachineReservationReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineReservationService_Deleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineReservationReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineReservationServiceServer).Deleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineReservationService_Deleted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineReservationServiceServer).Deleted(ctx, req.(*MachineReservationReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineReservationService_Usage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineReservationUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineReservationServiceServer).Usage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MachineReservationService_Usage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineReservationServiceServer).Usage(ctx, req.(*MachineReservationUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineReservationService_ServiceDesc is the grpc.ServiceDesc for MachineReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metalstack.io.accounting.api.v1.MachineReservationService",
	HandlerType: (*MachineReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Added",
			Handler:    _MachineReservationService_Added_Handler,
		},
		{
			MethodName: "Modified",
			Handler:    _MachineReservationService_Modified_Handler,
		},
		{
			MethodName: "Deleted",
			Handler:    _MachineReservationService_Deleted_Handler,
		},
		{
			MethodName: "Usage",
			Handler:    _MachineReservationService_Usage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metalstack/io/accounting/api/v1/machine_reservation.proto",
}

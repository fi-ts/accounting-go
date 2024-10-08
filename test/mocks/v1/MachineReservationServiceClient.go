// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
)

// MachineReservationServiceClient is an autogenerated mock type for the MachineReservationServiceClient type
type MachineReservationServiceClient struct {
	mock.Mock
}

// Added provides a mock function with given fields: ctx, in, opts
func (_m *MachineReservationServiceClient) Added(ctx context.Context, in *v1.MachineReservationReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Added")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) (*v1.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deleted provides a mock function with given fields: ctx, in, opts
func (_m *MachineReservationServiceClient) Deleted(ctx context.Context, in *v1.MachineReservationReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Deleted")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) (*v1.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modified provides a mock function with given fields: ctx, in, opts
func (_m *MachineReservationServiceClient) Modified(ctx context.Context, in *v1.MachineReservationReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Modified")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) (*v1.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReservationReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: ctx, in, opts
func (_m *MachineReservationServiceClient) Usage(ctx context.Context, in *v1.MachineReservationUsageRequest, opts ...grpc.CallOption) (*v1.MachineReservationUsageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Usage")
	}

	var r0 *v1.MachineReservationUsageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationUsageRequest, ...grpc.CallOption) (*v1.MachineReservationUsageResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReservationUsageRequest, ...grpc.CallOption) *v1.MachineReservationUsageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MachineReservationUsageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReservationUsageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMachineReservationServiceClient creates a new instance of MachineReservationServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMachineReservationServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MachineReservationServiceClient {
	mock := &MachineReservationServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

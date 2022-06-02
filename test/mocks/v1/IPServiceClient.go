// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
)

// IPServiceClient is an autogenerated mock type for the IPServiceClient type
type IPServiceClient struct {
	mock.Mock
}

// Added provides a mock function with given fields: ctx, in, opts
func (_m *IPServiceClient) Added(ctx context.Context, in *v1.IPReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IPReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IPReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deleted provides a mock function with given fields: ctx, in, opts
func (_m *IPServiceClient) Deleted(ctx context.Context, in *v1.IPReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IPReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IPReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modified provides a mock function with given fields: ctx, in, opts
func (_m *IPServiceClient) Modified(ctx context.Context, in *v1.IPReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IPReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IPReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: ctx, in, opts
func (_m *IPServiceClient) Usage(ctx context.Context, in *v1.IPUsageRequest, opts ...grpc.CallOption) (*v1.IPUsageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.IPUsageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.IPUsageRequest, ...grpc.CallOption) *v1.IPUsageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.IPUsageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.IPUsageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewIPServiceClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewIPServiceClient creates a new instance of IPServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIPServiceClient(t NewIPServiceClientT) *IPServiceClient {
	mock := &IPServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
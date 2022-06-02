// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
)

// PostgresServiceClient is an autogenerated mock type for the PostgresServiceClient type
type PostgresServiceClient struct {
	mock.Mock
}

// Added provides a mock function with given fields: ctx, in, opts
func (_m *PostgresServiceClient) Added(ctx context.Context, in *v1.PostgresReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deleted provides a mock function with given fields: ctx, in, opts
func (_m *PostgresServiceClient) Deleted(ctx context.Context, in *v1.PostgresReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modified provides a mock function with given fields: ctx, in, opts
func (_m *PostgresServiceClient) Modified(ctx context.Context, in *v1.PostgresReport, opts ...grpc.CallOption) (*v1.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Empty
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) *v1.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PostgresReport, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: ctx, in, opts
func (_m *PostgresServiceClient) Usage(ctx context.Context, in *v1.PostgresUsageRequest, opts ...grpc.CallOption) (*v1.PostgresUsageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.PostgresUsageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.PostgresUsageRequest, ...grpc.CallOption) *v1.PostgresUsageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PostgresUsageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.PostgresUsageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewPostgresServiceClientT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostgresServiceClient creates a new instance of PostgresServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostgresServiceClient(t NewPostgresServiceClientT) *PostgresServiceClient {
	mock := &PostgresServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
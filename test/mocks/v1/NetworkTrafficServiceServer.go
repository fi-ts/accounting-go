// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	mock "github.com/stretchr/testify/mock"
)

// NetworkTrafficServiceServer is an autogenerated mock type for the NetworkTrafficServiceServer type
type NetworkTrafficServiceServer struct {
	mock.Mock
}

// Modified provides a mock function with given fields: _a0, _a1
func (_m *NetworkTrafficServiceServer) Modified(_a0 context.Context, _a1 *v1.NetworkTrafficReport) (*v1.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Modified")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NetworkTrafficReport) (*v1.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NetworkTrafficReport) *v1.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NetworkTrafficReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: _a0, _a1
func (_m *NetworkTrafficServiceServer) Usage(_a0 context.Context, _a1 *v1.NetworkUsageRequest) (*v1.NetworkUsageResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Usage")
	}

	var r0 *v1.NetworkUsageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NetworkUsageRequest) (*v1.NetworkUsageResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.NetworkUsageRequest) *v1.NetworkUsageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.NetworkUsageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.NetworkUsageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNetworkTrafficServiceServer creates a new instance of NetworkTrafficServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNetworkTrafficServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *NetworkTrafficServiceServer {
	mock := &NetworkTrafficServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

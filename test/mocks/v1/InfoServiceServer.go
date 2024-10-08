// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	mock "github.com/stretchr/testify/mock"
)

// InfoServiceServer is an autogenerated mock type for the InfoServiceServer type
type InfoServiceServer struct {
	mock.Mock
}

// Projects provides a mock function with given fields: _a0, _a1
func (_m *InfoServiceServer) Projects(_a0 context.Context, _a1 *v1.ProjectInfoRequest) (*v1.ProjectInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Projects")
	}

	var r0 *v1.ProjectInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ProjectInfoRequest) (*v1.ProjectInfoResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.ProjectInfoRequest) *v1.ProjectInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.ProjectInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.ProjectInfoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Tenants provides a mock function with given fields: _a0, _a1
func (_m *InfoServiceServer) Tenants(_a0 context.Context, _a1 *v1.TenantInfoRequest) (*v1.TenantInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Tenants")
	}

	var r0 *v1.TenantInfoResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.TenantInfoRequest) (*v1.TenantInfoResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.TenantInfoRequest) *v1.TenantInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.TenantInfoResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.TenantInfoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInfoServiceServer creates a new instance of InfoServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfoServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *InfoServiceServer {
	mock := &InfoServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

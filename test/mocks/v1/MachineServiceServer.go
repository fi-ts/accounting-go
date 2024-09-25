// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import (
	context "context"

	v1 "github.com/fi-ts/accounting-go/pkg/apis/v1"
	mock "github.com/stretchr/testify/mock"
)

// MachineServiceServer is an autogenerated mock type for the MachineServiceServer type
type MachineServiceServer struct {
	mock.Mock
}

// Added provides a mock function with given fields: _a0, _a1
func (_m *MachineServiceServer) Added(_a0 context.Context, _a1 *v1.MachineReport) (*v1.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Added")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) (*v1.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) *v1.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deleted provides a mock function with given fields: _a0, _a1
func (_m *MachineServiceServer) Deleted(_a0 context.Context, _a1 *v1.MachineReport) (*v1.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Deleted")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) (*v1.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) *v1.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Modified provides a mock function with given fields: _a0, _a1
func (_m *MachineServiceServer) Modified(_a0 context.Context, _a1 *v1.MachineReport) (*v1.Empty, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Modified")
	}

	var r0 *v1.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) (*v1.Empty, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineReport) *v1.Empty); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineReport) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usage provides a mock function with given fields: _a0, _a1
func (_m *MachineServiceServer) Usage(_a0 context.Context, _a1 *v1.MachineUsageRequest) (*v1.MachineUsageResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Usage")
	}

	var r0 *v1.MachineUsageResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineUsageRequest) (*v1.MachineUsageResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1.MachineUsageRequest) *v1.MachineUsageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.MachineUsageResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1.MachineUsageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMachineServiceServer creates a new instance of MachineServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMachineServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MachineServiceServer {
	mock := &MachineServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafePodServiceServer is an autogenerated mock type for the UnsafePodServiceServer type
type UnsafePodServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedPodServiceServer provides a mock function with given fields:
func (_m *UnsafePodServiceServer) mustEmbedUnimplementedPodServiceServer() {
	_m.Called()
}

type mockConstructorTestingTNewUnsafePodServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewUnsafePodServiceServer creates a new instance of UnsafePodServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUnsafePodServiceServer(t mockConstructorTestingTNewUnsafePodServiceServer) *UnsafePodServiceServer {
	mock := &UnsafePodServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

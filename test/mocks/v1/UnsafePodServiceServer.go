// Code generated by mockery v2.38.0. DO NOT EDIT.

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

// NewUnsafePodServiceServer creates a new instance of UnsafePodServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafePodServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafePodServiceServer {
	mock := &UnsafePodServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.46.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafePostgresServiceServer is an autogenerated mock type for the UnsafePostgresServiceServer type
type UnsafePostgresServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedPostgresServiceServer provides a mock function with given fields:
func (_m *UnsafePostgresServiceServer) mustEmbedUnimplementedPostgresServiceServer() {
	_m.Called()
}

// NewUnsafePostgresServiceServer creates a new instance of UnsafePostgresServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafePostgresServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafePostgresServiceServer {
	mock := &UnsafePostgresServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

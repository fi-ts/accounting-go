// Code generated by mockery v2.46.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UnsafeProductOptionServiceServer is an autogenerated mock type for the UnsafeProductOptionServiceServer type
type UnsafeProductOptionServiceServer struct {
	mock.Mock
}

// mustEmbedUnimplementedProductOptionServiceServer provides a mock function with given fields:
func (_m *UnsafeProductOptionServiceServer) mustEmbedUnimplementedProductOptionServiceServer() {
	_m.Called()
}

// NewUnsafeProductOptionServiceServer creates a new instance of UnsafeProductOptionServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUnsafeProductOptionServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UnsafeProductOptionServiceServer {
	mock := &UnsafeProductOptionServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

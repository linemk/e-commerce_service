// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Methods is an autogenerated mock type for the Methods type
type Methods struct {
	mock.Mock
}

// getAllProducts provides a mock function with given fields: w, r
func (_m *Methods) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// getProductsById provides a mock function with given fields: w, r
func (_m *Methods) GetProductsById(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// postAndPutProducts provides a mock function with given fields: w, r
func (_m *Methods) PostAndPutProducts(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewMethods creates a new instance of Methods. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMethods(t interface {
	mock.TestingT
	Cleanup(func())
}) *Methods {
	mock := &Methods{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
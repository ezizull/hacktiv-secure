// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	book "secure/challenge-3/domain/book"

	mock "github.com/stretchr/testify/mock"
)

// BookMocks is an autogenerated mock type for the BookMocks type
type BookMocks struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *BookMocks) Create(_a0 *book.Book) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*book.Book) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *BookMocks) Delete(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *BookMocks) Get(_a0 int) (*book.Book, error) {
	ret := _m.Called(_a0)

	var r0 *book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*book.Book, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *book.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*book.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *BookMocks) GetAll() ([]*book.Book, error) {
	ret := _m.Called()

	var r0 []*book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*book.Book, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*book.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*book.Book)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *BookMocks) GetByID(_a0 int) (*book.Book, error) {
	ret := _m.Called(_a0)

	var r0 *book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*book.Book, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *book.Book); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*book.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMap provides a mock function with given fields: _a0
func (_m *BookMocks) GetByMap(_a0 map[string]interface{}) map[string]interface{} {
	ret := _m.Called(_a0)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *BookMocks) Update(_a0 int, _a1 map[string]interface{}) (*book.Book, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *book.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) (*book.Book, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) *book.Book); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*book.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBookMocks interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookMocks creates a new instance of BookMocks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookMocks(t mockConstructorTestingTNewBookMocks) *BookMocks {
	mock := &BookMocks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

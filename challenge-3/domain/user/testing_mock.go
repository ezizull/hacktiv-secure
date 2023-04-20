// Code generated by mockery v2.20.0. DO NOT EDIT.

package user

import mock "github.com/stretchr/testify/mock"

// MockTesting is an autogenerated mock type for the Testing type
type MockTesting struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *MockTesting) Create(_a0 *User) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *MockTesting) Delete(_a0 int) error {
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
func (_m *MockTesting) Get(_a0 int) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
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
func (_m *MockTesting) GetAll() ([]*User, error) {
	ret := _m.Called()

	var r0 []*User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*User)
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
func (_m *MockTesting) GetByID(_a0 int) (*User, error) {
	ret := _m.Called(_a0)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(int) *User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
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
func (_m *MockTesting) GetByMap(_a0 map[string]interface{}) map[string]interface{} {
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
func (_m *MockTesting) Update(_a0 int, _a1 map[string]interface{}) (*User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *User
	var r1 error
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) (*User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(int, map[string]interface{}) *User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*User)
		}
	}

	if rf, ok := ret.Get(1).(func(int, map[string]interface{}) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockTesting interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTesting creates a new instance of MockTesting. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTesting(t mockConstructorTestingTNewMockTesting) *MockTesting {
	mock := &MockTesting{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

package book

import (
	domainBook "secure/challenge-3/domain/book"
	"time"

	"github.com/stretchr/testify/mock"
)

type BookMock struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: date
func (m *BookMock) GetAll(date time.Time) ([]*domainBook.TestBook, error) {
	ret := m.Called(date)

	var r0 []*domainBook.TestBook
	if rf, ok := ret.Get(0).(func(time.Time) []*domainBook.TestBook); ok {
		r0 = rf(date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domainBook.TestBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GeByID provides a mock function with given fields:
func (m *BookMock) GeByID() (*domainBook.TestBook, error) {
	ret := m.Called()

	var r0 *domainBook.TestBook
	if rf, ok := ret.Get(0).(func() *domainBook.TestBook); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainBook.TestBook)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

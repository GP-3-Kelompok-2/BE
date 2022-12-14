// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	homestay "github.com/GP-3-Kelompok-2/airbnb-app-project/features/homestay"
	mock "github.com/stretchr/testify/mock"
)

// HomestayRepository is an autogenerated mock type for the RepositoryInterface type
type HomestayRepository struct {
	mock.Mock
}

// DeleteHomestay provides a mock function with given fields: id
func (_m *HomestayRepository) DeleteHomestay(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllHomestays provides a mock function with given fields: keyword
func (_m *HomestayRepository) GetAllHomestays(keyword string) ([]homestay.HomestayCore, error) {
	ret := _m.Called(keyword)

	var r0 []homestay.HomestayCore
	if rf, ok := ret.Get(0).(func(string) []homestay.HomestayCore); ok {
		r0 = rf(keyword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]homestay.HomestayCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(keyword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHomestayById provides a mock function with given fields: id
func (_m *HomestayRepository) GetHomestayById(id uint) (homestay.HomestayCore, error) {
	ret := _m.Called(id)

	var r0 homestay.HomestayCore
	if rf, ok := ret.Get(0).(func(uint) homestay.HomestayCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(homestay.HomestayCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertHomestay provides a mock function with given fields: data
func (_m *HomestayRepository) InsertHomestay(data homestay.HomestayCore) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(homestay.HomestayCore) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(homestay.HomestayCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateHomestay provides a mock function with given fields: input, id
func (_m *HomestayRepository) UpdateHomestay(input homestay.HomestayCore, id uint) (homestay.HomestayCore, error) {
	ret := _m.Called(input, id)

	var r0 homestay.HomestayCore
	if rf, ok := ret.Get(0).(func(homestay.HomestayCore, uint) homestay.HomestayCore); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Get(0).(homestay.HomestayCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(homestay.HomestayCore, uint) error); ok {
		r1 = rf(input, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHomestayRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewHomestayRepository creates a new instance of HomestayRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHomestayRepository(t mockConstructorTestingTNewHomestayRepository) *HomestayRepository {
	mock := &HomestayRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

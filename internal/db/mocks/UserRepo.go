// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	models "github.com/dropdevrahul/simple-api/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the UserRepo type
type UserRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: d, u
func (_m *UserRepo) Create(d *models.DB, u *models.User) error {
	ret := _m.Called(d, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.DB, *models.User) error); ok {
		r0 = rf(d, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: d, email, u
func (_m *UserRepo) GetByEmail(d *models.DB, email string, u *models.User) error {
	ret := _m.Called(d, email, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.DB, string, *models.User) error); ok {
		r0 = rf(d, email, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepo(t mockConstructorTestingTNewUserRepo) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

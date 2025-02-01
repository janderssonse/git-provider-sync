// SPDX-FileCopyrightText: 2025 Josef Andersson
//
// SPDX-License-Identifier: EUPL-1.2

// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	retryablehttp "github.com/hashicorp/go-retryablehttp"
	mock "github.com/stretchr/testify/mock"
)

// RequestOptionFunc is an autogenerated mock type for the RequestOptionFunc type
type RequestOptionFunc struct {
	mock.Mock
}

type RequestOptionFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestOptionFunc) EXPECT() *RequestOptionFunc_Expecter {
	return &RequestOptionFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *RequestOptionFunc) Execute(_a0 *retryablehttp.Request) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*retryablehttp.Request) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestOptionFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type RequestOptionFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *retryablehttp.Request
func (_e *RequestOptionFunc_Expecter) Execute(_a0 interface{}) *RequestOptionFunc_Execute_Call {
	return &RequestOptionFunc_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *RequestOptionFunc_Execute_Call) Run(run func(_a0 *retryablehttp.Request)) *RequestOptionFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*retryablehttp.Request))
	})
	return _c
}

func (_c *RequestOptionFunc_Execute_Call) Return(_a0 error) *RequestOptionFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestOptionFunc_Execute_Call) RunAndReturn(run func(*retryablehttp.Request) error) *RequestOptionFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestOptionFunc creates a new instance of RequestOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestOptionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestOptionFunc {
	mock := &RequestOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

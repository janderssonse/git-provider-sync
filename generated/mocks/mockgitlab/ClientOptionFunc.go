// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ClientOptionFunc is an autogenerated mock type for the ClientOptionFunc type
type ClientOptionFunc struct {
	mock.Mock
}

type ClientOptionFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientOptionFunc) EXPECT() *ClientOptionFunc_Expecter {
	return &ClientOptionFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *ClientOptionFunc) Execute(_a0 *gitlab.Client) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gitlab.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClientOptionFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ClientOptionFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *gitlab.Client
func (_e *ClientOptionFunc_Expecter) Execute(_a0 interface{}) *ClientOptionFunc_Execute_Call {
	return &ClientOptionFunc_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *ClientOptionFunc_Execute_Call) Run(run func(_a0 *gitlab.Client)) *ClientOptionFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gitlab.Client))
	})
	return _c
}

func (_c *ClientOptionFunc_Execute_Call) Return(_a0 error) *ClientOptionFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ClientOptionFunc_Execute_Call) RunAndReturn(run func(*gitlab.Client) error) *ClientOptionFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientOptionFunc creates a new instance of ClientOptionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientOptionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientOptionFunc {
	mock := &ClientOptionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

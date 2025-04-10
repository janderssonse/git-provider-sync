// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// VersionServiceInterface is an autogenerated mock type for the VersionServiceInterface type
type VersionServiceInterface struct {
	mock.Mock
}

type VersionServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *VersionServiceInterface) EXPECT() *VersionServiceInterface_Expecter {
	return &VersionServiceInterface_Expecter{mock: &_m.Mock}
}

// GetVersion provides a mock function with given fields: options
func (_m *VersionServiceInterface) GetVersion(options ...gitlab.RequestOptionFunc) (*gitlab.Version, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetVersion")
	}

	var r0 *gitlab.Version
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) (*gitlab.Version, *gitlab.Response, error)); ok {
		return rf(options...)
	}
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) *gitlab.Version); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Version)
		}
	}

	if rf, ok := ret.Get(1).(func(...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// VersionServiceInterface_GetVersion_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVersion'
type VersionServiceInterface_GetVersion_Call struct {
	*mock.Call
}

// GetVersion is a helper method to define mock.On call
//   - options ...gitlab.RequestOptionFunc
func (_e *VersionServiceInterface_Expecter) GetVersion(options ...interface{}) *VersionServiceInterface_GetVersion_Call {
	return &VersionServiceInterface_GetVersion_Call{Call: _e.mock.On("GetVersion",
		append([]interface{}{}, options...)...)}
}

func (_c *VersionServiceInterface_GetVersion_Call) Run(run func(options ...gitlab.RequestOptionFunc)) *VersionServiceInterface_GetVersion_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *VersionServiceInterface_GetVersion_Call) Return(_a0 *gitlab.Version, _a1 *gitlab.Response, _a2 error) *VersionServiceInterface_GetVersion_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *VersionServiceInterface_GetVersion_Call) RunAndReturn(run func(...gitlab.RequestOptionFunc) (*gitlab.Version, *gitlab.Response, error)) *VersionServiceInterface_GetVersion_Call {
	_c.Call.Return(run)
	return _c
}

// NewVersionServiceInterface creates a new instance of VersionServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVersionServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *VersionServiceInterface {
	mock := &VersionServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

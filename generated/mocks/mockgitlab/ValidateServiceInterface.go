// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ValidateServiceInterface is an autogenerated mock type for the ValidateServiceInterface type
type ValidateServiceInterface struct {
	mock.Mock
}

type ValidateServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ValidateServiceInterface) EXPECT() *ValidateServiceInterface_Expecter {
	return &ValidateServiceInterface_Expecter{mock: &_m.Mock}
}

// ProjectLint provides a mock function with given fields: pid, opt, options
func (_m *ValidateServiceInterface) ProjectLint(pid interface{}, opt *gitlab.ProjectLintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ProjectLint")
	}

	var r0 *gitlab.ProjectLintResult
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectLintOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectLintOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectLintResult); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectLintResult)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ProjectLintOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ProjectLintOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ValidateServiceInterface_ProjectLint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProjectLint'
type ValidateServiceInterface_ProjectLint_Call struct {
	*mock.Call
}

// ProjectLint is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ProjectLintOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ValidateServiceInterface_Expecter) ProjectLint(pid interface{}, opt interface{}, options ...interface{}) *ValidateServiceInterface_ProjectLint_Call {
	return &ValidateServiceInterface_ProjectLint_Call{Call: _e.mock.On("ProjectLint",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ValidateServiceInterface_ProjectLint_Call) Run(run func(pid interface{}, opt *gitlab.ProjectLintOptions, options ...gitlab.RequestOptionFunc)) *ValidateServiceInterface_ProjectLint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ProjectLintOptions), variadicArgs...)
	})
	return _c
}

func (_c *ValidateServiceInterface_ProjectLint_Call) Return(_a0 *gitlab.ProjectLintResult, _a1 *gitlab.Response, _a2 error) *ValidateServiceInterface_ProjectLint_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ValidateServiceInterface_ProjectLint_Call) RunAndReturn(run func(interface{}, *gitlab.ProjectLintOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error)) *ValidateServiceInterface_ProjectLint_Call {
	_c.Call.Return(run)
	return _c
}

// ProjectNamespaceLint provides a mock function with given fields: pid, opt, options
func (_m *ValidateServiceInterface) ProjectNamespaceLint(pid interface{}, opt *gitlab.ProjectNamespaceLintOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ProjectNamespaceLint")
	}

	var r0 *gitlab.ProjectLintResult
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectNamespaceLintOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectNamespaceLintOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectLintResult); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectLintResult)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ProjectNamespaceLintOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ProjectNamespaceLintOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ValidateServiceInterface_ProjectNamespaceLint_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProjectNamespaceLint'
type ValidateServiceInterface_ProjectNamespaceLint_Call struct {
	*mock.Call
}

// ProjectNamespaceLint is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ProjectNamespaceLintOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ValidateServiceInterface_Expecter) ProjectNamespaceLint(pid interface{}, opt interface{}, options ...interface{}) *ValidateServiceInterface_ProjectNamespaceLint_Call {
	return &ValidateServiceInterface_ProjectNamespaceLint_Call{Call: _e.mock.On("ProjectNamespaceLint",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ValidateServiceInterface_ProjectNamespaceLint_Call) Run(run func(pid interface{}, opt *gitlab.ProjectNamespaceLintOptions, options ...gitlab.RequestOptionFunc)) *ValidateServiceInterface_ProjectNamespaceLint_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ProjectNamespaceLintOptions), variadicArgs...)
	})
	return _c
}

func (_c *ValidateServiceInterface_ProjectNamespaceLint_Call) Return(_a0 *gitlab.ProjectLintResult, _a1 *gitlab.Response, _a2 error) *ValidateServiceInterface_ProjectNamespaceLint_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ValidateServiceInterface_ProjectNamespaceLint_Call) RunAndReturn(run func(interface{}, *gitlab.ProjectNamespaceLintOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectLintResult, *gitlab.Response, error)) *ValidateServiceInterface_ProjectNamespaceLint_Call {
	_c.Call.Return(run)
	return _c
}

// NewValidateServiceInterface creates a new instance of ValidateServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewValidateServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ValidateServiceInterface {
	mock := &ValidateServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

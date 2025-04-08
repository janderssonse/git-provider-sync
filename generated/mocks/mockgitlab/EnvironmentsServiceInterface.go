// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// EnvironmentsServiceInterface is an autogenerated mock type for the EnvironmentsServiceInterface type
type EnvironmentsServiceInterface struct {
	mock.Mock
}

type EnvironmentsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EnvironmentsServiceInterface) EXPECT() *EnvironmentsServiceInterface_Expecter {
	return &EnvironmentsServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateEnvironment provides a mock function with given fields: pid, opt, options
func (_m *EnvironmentsServiceInterface) CreateEnvironment(pid interface{}, opt *gitlab.CreateEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateEnvironment")
	}

	var r0 *gitlab.Environment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Environment); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateEnvironmentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvironmentsServiceInterface_CreateEnvironment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEnvironment'
type EnvironmentsServiceInterface_CreateEnvironment_Call struct {
	*mock.Call
}

// CreateEnvironment is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateEnvironmentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) CreateEnvironment(pid interface{}, opt interface{}, options ...interface{}) *EnvironmentsServiceInterface_CreateEnvironment_Call {
	return &EnvironmentsServiceInterface_CreateEnvironment_Call{Call: _e.mock.On("CreateEnvironment",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_CreateEnvironment_Call) Run(run func(pid interface{}, opt *gitlab.CreateEnvironmentOptions, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_CreateEnvironment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateEnvironmentOptions), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_CreateEnvironment_Call) Return(_a0 *gitlab.Environment, _a1 *gitlab.Response, _a2 error) *EnvironmentsServiceInterface_CreateEnvironment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EnvironmentsServiceInterface_CreateEnvironment_Call) RunAndReturn(run func(interface{}, *gitlab.CreateEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)) *EnvironmentsServiceInterface_CreateEnvironment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEnvironment provides a mock function with given fields: pid, environment, options
func (_m *EnvironmentsServiceInterface) DeleteEnvironment(pid interface{}, environment int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, environment)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEnvironment")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, environment, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, environment, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, environment, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnvironmentsServiceInterface_DeleteEnvironment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEnvironment'
type EnvironmentsServiceInterface_DeleteEnvironment_Call struct {
	*mock.Call
}

// DeleteEnvironment is a helper method to define mock.On call
//   - pid interface{}
//   - environment int
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) DeleteEnvironment(pid interface{}, environment interface{}, options ...interface{}) *EnvironmentsServiceInterface_DeleteEnvironment_Call {
	return &EnvironmentsServiceInterface_DeleteEnvironment_Call{Call: _e.mock.On("DeleteEnvironment",
		append([]interface{}{pid, environment}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_DeleteEnvironment_Call) Run(run func(pid interface{}, environment int, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_DeleteEnvironment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_DeleteEnvironment_Call) Return(_a0 *gitlab.Response, _a1 error) *EnvironmentsServiceInterface_DeleteEnvironment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EnvironmentsServiceInterface_DeleteEnvironment_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *EnvironmentsServiceInterface_DeleteEnvironment_Call {
	_c.Call.Return(run)
	return _c
}

// EditEnvironment provides a mock function with given fields: pid, environment, opt, options
func (_m *EnvironmentsServiceInterface) EditEnvironment(pid interface{}, environment int, opt *gitlab.EditEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, environment, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EditEnvironment")
	}

	var r0 *gitlab.Environment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)); ok {
		return rf(pid, environment, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Environment); ok {
		r0 = rf(pid, environment, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.EditEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, environment, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.EditEnvironmentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, environment, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvironmentsServiceInterface_EditEnvironment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditEnvironment'
type EnvironmentsServiceInterface_EditEnvironment_Call struct {
	*mock.Call
}

// EditEnvironment is a helper method to define mock.On call
//   - pid interface{}
//   - environment int
//   - opt *gitlab.EditEnvironmentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) EditEnvironment(pid interface{}, environment interface{}, opt interface{}, options ...interface{}) *EnvironmentsServiceInterface_EditEnvironment_Call {
	return &EnvironmentsServiceInterface_EditEnvironment_Call{Call: _e.mock.On("EditEnvironment",
		append([]interface{}{pid, environment, opt}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_EditEnvironment_Call) Run(run func(pid interface{}, environment int, opt *gitlab.EditEnvironmentOptions, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_EditEnvironment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.EditEnvironmentOptions), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_EditEnvironment_Call) Return(_a0 *gitlab.Environment, _a1 *gitlab.Response, _a2 error) *EnvironmentsServiceInterface_EditEnvironment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EnvironmentsServiceInterface_EditEnvironment_Call) RunAndReturn(run func(interface{}, int, *gitlab.EditEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)) *EnvironmentsServiceInterface_EditEnvironment_Call {
	_c.Call.Return(run)
	return _c
}

// GetEnvironment provides a mock function with given fields: pid, environment, options
func (_m *EnvironmentsServiceInterface) GetEnvironment(pid interface{}, environment int, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, environment)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetEnvironment")
	}

	var r0 *gitlab.Environment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)); ok {
		return rf(pid, environment, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Environment); ok {
		r0 = rf(pid, environment, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, environment, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, environment, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvironmentsServiceInterface_GetEnvironment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEnvironment'
type EnvironmentsServiceInterface_GetEnvironment_Call struct {
	*mock.Call
}

// GetEnvironment is a helper method to define mock.On call
//   - pid interface{}
//   - environment int
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) GetEnvironment(pid interface{}, environment interface{}, options ...interface{}) *EnvironmentsServiceInterface_GetEnvironment_Call {
	return &EnvironmentsServiceInterface_GetEnvironment_Call{Call: _e.mock.On("GetEnvironment",
		append([]interface{}{pid, environment}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_GetEnvironment_Call) Run(run func(pid interface{}, environment int, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_GetEnvironment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_GetEnvironment_Call) Return(_a0 *gitlab.Environment, _a1 *gitlab.Response, _a2 error) *EnvironmentsServiceInterface_GetEnvironment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EnvironmentsServiceInterface_GetEnvironment_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)) *EnvironmentsServiceInterface_GetEnvironment_Call {
	_c.Call.Return(run)
	return _c
}

// ListEnvironments provides a mock function with given fields: pid, opts, options
func (_m *EnvironmentsServiceInterface) ListEnvironments(pid interface{}, opts *gitlab.ListEnvironmentsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Environment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListEnvironments")
	}

	var r0 []*gitlab.Environment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListEnvironmentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Environment, *gitlab.Response, error)); ok {
		return rf(pid, opts, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListEnvironmentsOptions, ...gitlab.RequestOptionFunc) []*gitlab.Environment); ok {
		r0 = rf(pid, opts, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListEnvironmentsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opts, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListEnvironmentsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opts, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvironmentsServiceInterface_ListEnvironments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListEnvironments'
type EnvironmentsServiceInterface_ListEnvironments_Call struct {
	*mock.Call
}

// ListEnvironments is a helper method to define mock.On call
//   - pid interface{}
//   - opts *gitlab.ListEnvironmentsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) ListEnvironments(pid interface{}, opts interface{}, options ...interface{}) *EnvironmentsServiceInterface_ListEnvironments_Call {
	return &EnvironmentsServiceInterface_ListEnvironments_Call{Call: _e.mock.On("ListEnvironments",
		append([]interface{}{pid, opts}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_ListEnvironments_Call) Run(run func(pid interface{}, opts *gitlab.ListEnvironmentsOptions, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_ListEnvironments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListEnvironmentsOptions), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_ListEnvironments_Call) Return(_a0 []*gitlab.Environment, _a1 *gitlab.Response, _a2 error) *EnvironmentsServiceInterface_ListEnvironments_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EnvironmentsServiceInterface_ListEnvironments_Call) RunAndReturn(run func(interface{}, *gitlab.ListEnvironmentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Environment, *gitlab.Response, error)) *EnvironmentsServiceInterface_ListEnvironments_Call {
	_c.Call.Return(run)
	return _c
}

// StopEnvironment provides a mock function with given fields: pid, environmentID, opt, options
func (_m *EnvironmentsServiceInterface) StopEnvironment(pid interface{}, environmentID int, opt *gitlab.StopEnvironmentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, environmentID, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StopEnvironment")
	}

	var r0 *gitlab.Environment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.StopEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)); ok {
		return rf(pid, environmentID, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.StopEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Environment); ok {
		r0 = rf(pid, environmentID, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Environment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.StopEnvironmentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, environmentID, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.StopEnvironmentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, environmentID, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnvironmentsServiceInterface_StopEnvironment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopEnvironment'
type EnvironmentsServiceInterface_StopEnvironment_Call struct {
	*mock.Call
}

// StopEnvironment is a helper method to define mock.On call
//   - pid interface{}
//   - environmentID int
//   - opt *gitlab.StopEnvironmentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *EnvironmentsServiceInterface_Expecter) StopEnvironment(pid interface{}, environmentID interface{}, opt interface{}, options ...interface{}) *EnvironmentsServiceInterface_StopEnvironment_Call {
	return &EnvironmentsServiceInterface_StopEnvironment_Call{Call: _e.mock.On("StopEnvironment",
		append([]interface{}{pid, environmentID, opt}, options...)...)}
}

func (_c *EnvironmentsServiceInterface_StopEnvironment_Call) Run(run func(pid interface{}, environmentID int, opt *gitlab.StopEnvironmentOptions, options ...gitlab.RequestOptionFunc)) *EnvironmentsServiceInterface_StopEnvironment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.StopEnvironmentOptions), variadicArgs...)
	})
	return _c
}

func (_c *EnvironmentsServiceInterface_StopEnvironment_Call) Return(_a0 *gitlab.Environment, _a1 *gitlab.Response, _a2 error) *EnvironmentsServiceInterface_StopEnvironment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *EnvironmentsServiceInterface_StopEnvironment_Call) RunAndReturn(run func(interface{}, int, *gitlab.StopEnvironmentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Environment, *gitlab.Response, error)) *EnvironmentsServiceInterface_StopEnvironment_Call {
	_c.Call.Return(run)
	return _c
}

// NewEnvironmentsServiceInterface creates a new instance of EnvironmentsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEnvironmentsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EnvironmentsServiceInterface {
	mock := &EnvironmentsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

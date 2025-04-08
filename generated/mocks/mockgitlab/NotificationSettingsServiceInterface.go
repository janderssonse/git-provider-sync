// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// NotificationSettingsServiceInterface is an autogenerated mock type for the NotificationSettingsServiceInterface type
type NotificationSettingsServiceInterface struct {
	mock.Mock
}

type NotificationSettingsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *NotificationSettingsServiceInterface) EXPECT() *NotificationSettingsServiceInterface_Expecter {
	return &NotificationSettingsServiceInterface_Expecter{mock: &_m.Mock}
}

// GetGlobalSettings provides a mock function with given fields: options
func (_m *NotificationSettingsServiceInterface) GetGlobalSettings(options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGlobalSettings")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(options...)
	}
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
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

// NotificationSettingsServiceInterface_GetGlobalSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGlobalSettings'
type NotificationSettingsServiceInterface_GetGlobalSettings_Call struct {
	*mock.Call
}

// GetGlobalSettings is a helper method to define mock.On call
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) GetGlobalSettings(options ...interface{}) *NotificationSettingsServiceInterface_GetGlobalSettings_Call {
	return &NotificationSettingsServiceInterface_GetGlobalSettings_Call{Call: _e.mock.On("GetGlobalSettings",
		append([]interface{}{}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_GetGlobalSettings_Call) Run(run func(options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_GetGlobalSettings_Call {
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

func (_c *NotificationSettingsServiceInterface_GetGlobalSettings_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_GetGlobalSettings_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_GetGlobalSettings_Call) RunAndReturn(run func(...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_GetGlobalSettings_Call {
	_c.Call.Return(run)
	return _c
}

// GetSettingsForGroup provides a mock function with given fields: gid, options
func (_m *NotificationSettingsServiceInterface) GetSettingsForGroup(gid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSettingsForGroup")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(gid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(gid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationSettingsServiceInterface_GetSettingsForGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSettingsForGroup'
type NotificationSettingsServiceInterface_GetSettingsForGroup_Call struct {
	*mock.Call
}

// GetSettingsForGroup is a helper method to define mock.On call
//   - gid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) GetSettingsForGroup(gid interface{}, options ...interface{}) *NotificationSettingsServiceInterface_GetSettingsForGroup_Call {
	return &NotificationSettingsServiceInterface_GetSettingsForGroup_Call{Call: _e.mock.On("GetSettingsForGroup",
		append([]interface{}{gid}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForGroup_Call) Run(run func(gid interface{}, options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_GetSettingsForGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForGroup_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_GetSettingsForGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForGroup_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_GetSettingsForGroup_Call {
	_c.Call.Return(run)
	return _c
}

// GetSettingsForProject provides a mock function with given fields: pid, options
func (_m *NotificationSettingsServiceInterface) GetSettingsForProject(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSettingsForProject")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(pid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(pid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationSettingsServiceInterface_GetSettingsForProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSettingsForProject'
type NotificationSettingsServiceInterface_GetSettingsForProject_Call struct {
	*mock.Call
}

// GetSettingsForProject is a helper method to define mock.On call
//   - pid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) GetSettingsForProject(pid interface{}, options ...interface{}) *NotificationSettingsServiceInterface_GetSettingsForProject_Call {
	return &NotificationSettingsServiceInterface_GetSettingsForProject_Call{Call: _e.mock.On("GetSettingsForProject",
		append([]interface{}{pid}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForProject_Call) Run(run func(pid interface{}, options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_GetSettingsForProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForProject_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_GetSettingsForProject_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_GetSettingsForProject_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_GetSettingsForProject_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateGlobalSettings provides a mock function with given fields: opt, options
func (_m *NotificationSettingsServiceInterface) UpdateGlobalSettings(opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateGlobalSettings")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(*gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(opt, options...)
	}
	if rf, ok := ret.Get(0).(func(*gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(*gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(*gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationSettingsServiceInterface_UpdateGlobalSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateGlobalSettings'
type NotificationSettingsServiceInterface_UpdateGlobalSettings_Call struct {
	*mock.Call
}

// UpdateGlobalSettings is a helper method to define mock.On call
//   - opt *gitlab.NotificationSettingsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) UpdateGlobalSettings(opt interface{}, options ...interface{}) *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call {
	return &NotificationSettingsServiceInterface_UpdateGlobalSettings_Call{Call: _e.mock.On("UpdateGlobalSettings",
		append([]interface{}{opt}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call) Run(run func(opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(*gitlab.NotificationSettingsOptions), variadicArgs...)
	})
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call) RunAndReturn(run func(*gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_UpdateGlobalSettings_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSettingsForGroup provides a mock function with given fields: gid, opt, options
func (_m *NotificationSettingsServiceInterface) UpdateSettingsForGroup(gid interface{}, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSettingsForGroup")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSettingsForGroup'
type NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call struct {
	*mock.Call
}

// UpdateSettingsForGroup is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.NotificationSettingsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) UpdateSettingsForGroup(gid interface{}, opt interface{}, options ...interface{}) *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call {
	return &NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call{Call: _e.mock.On("UpdateSettingsForGroup",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call) Run(run func(gid interface{}, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.NotificationSettingsOptions), variadicArgs...)
	})
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call) RunAndReturn(run func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_UpdateSettingsForGroup_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSettingsForProject provides a mock function with given fields: pid, opt, options
func (_m *NotificationSettingsServiceInterface) UpdateSettingsForProject(pid interface{}, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSettingsForProject")
	}

	var r0 *gitlab.NotificationSettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.NotificationSettings); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.NotificationSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NotificationSettingsServiceInterface_UpdateSettingsForProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSettingsForProject'
type NotificationSettingsServiceInterface_UpdateSettingsForProject_Call struct {
	*mock.Call
}

// UpdateSettingsForProject is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.NotificationSettingsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *NotificationSettingsServiceInterface_Expecter) UpdateSettingsForProject(pid interface{}, opt interface{}, options ...interface{}) *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call {
	return &NotificationSettingsServiceInterface_UpdateSettingsForProject_Call{Call: _e.mock.On("UpdateSettingsForProject",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call) Run(run func(pid interface{}, opt *gitlab.NotificationSettingsOptions, options ...gitlab.RequestOptionFunc)) *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.NotificationSettingsOptions), variadicArgs...)
	})
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call) Return(_a0 *gitlab.NotificationSettings, _a1 *gitlab.Response, _a2 error) *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call) RunAndReturn(run func(interface{}, *gitlab.NotificationSettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.NotificationSettings, *gitlab.Response, error)) *NotificationSettingsServiceInterface_UpdateSettingsForProject_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotificationSettingsServiceInterface creates a new instance of NotificationSettingsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationSettingsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationSettingsServiceInterface {
	mock := &NotificationSettingsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

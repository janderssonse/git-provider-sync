// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ProjectSecuritySettingsServiceInterface is an autogenerated mock type for the ProjectSecuritySettingsServiceInterface type
type ProjectSecuritySettingsServiceInterface struct {
	mock.Mock
}

type ProjectSecuritySettingsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectSecuritySettingsServiceInterface) EXPECT() *ProjectSecuritySettingsServiceInterface_Expecter {
	return &ProjectSecuritySettingsServiceInterface_Expecter{mock: &_m.Mock}
}

// ListProjectSecuritySettings provides a mock function with given fields: pid, options
func (_m *ProjectSecuritySettingsServiceInterface) ListProjectSecuritySettings(pid interface{}, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectSecuritySettings")
	}

	var r0 *gitlab.ProjectSecuritySettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error)); ok {
		return rf(pid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.ProjectSecuritySettings); ok {
		r0 = rf(pid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectSecuritySettings)
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

// ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectSecuritySettings'
type ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call struct {
	*mock.Call
}

// ListProjectSecuritySettings is a helper method to define mock.On call
//   - pid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSecuritySettingsServiceInterface_Expecter) ListProjectSecuritySettings(pid interface{}, options ...interface{}) *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call {
	return &ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call{Call: _e.mock.On("ListProjectSecuritySettings",
		append([]interface{}{pid}, options...)...)}
}

func (_c *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call) Run(run func(pid interface{}, options ...gitlab.RequestOptionFunc)) *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call {
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

func (_c *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call) Return(_a0 *gitlab.ProjectSecuritySettings, _a1 *gitlab.Response, _a2 error) *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error)) *ProjectSecuritySettingsServiceInterface_ListProjectSecuritySettings_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSecretPushProtectionEnabledSetting provides a mock function with given fields: pid, opt, options
func (_m *ProjectSecuritySettingsServiceInterface) UpdateSecretPushProtectionEnabledSetting(pid interface{}, opt gitlab.UpdateProjectSecuritySettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecretPushProtectionEnabledSetting")
	}

	var r0 *gitlab.ProjectSecuritySettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.UpdateProjectSecuritySettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.UpdateProjectSecuritySettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectSecuritySettings); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectSecuritySettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, gitlab.UpdateProjectSecuritySettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, gitlab.UpdateProjectSecuritySettingsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSecretPushProtectionEnabledSetting'
type ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call struct {
	*mock.Call
}

// UpdateSecretPushProtectionEnabledSetting is a helper method to define mock.On call
//   - pid interface{}
//   - opt gitlab.UpdateProjectSecuritySettingsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSecuritySettingsServiceInterface_Expecter) UpdateSecretPushProtectionEnabledSetting(pid interface{}, opt interface{}, options ...interface{}) *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	return &ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call{Call: _e.mock.On("UpdateSecretPushProtectionEnabledSetting",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) Run(run func(pid interface{}, opt gitlab.UpdateProjectSecuritySettingsOptions, options ...gitlab.RequestOptionFunc)) *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(gitlab.UpdateProjectSecuritySettingsOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) Return(_a0 *gitlab.ProjectSecuritySettings, _a1 *gitlab.Response, _a2 error) *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) RunAndReturn(run func(interface{}, gitlab.UpdateProjectSecuritySettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectSecuritySettings, *gitlab.Response, error)) *ProjectSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectSecuritySettingsServiceInterface creates a new instance of ProjectSecuritySettingsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectSecuritySettingsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectSecuritySettingsServiceInterface {
	mock := &ProjectSecuritySettingsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GroupSecuritySettingsServiceInterface is an autogenerated mock type for the GroupSecuritySettingsServiceInterface type
type GroupSecuritySettingsServiceInterface struct {
	mock.Mock
}

type GroupSecuritySettingsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupSecuritySettingsServiceInterface) EXPECT() *GroupSecuritySettingsServiceInterface_Expecter {
	return &GroupSecuritySettingsServiceInterface_Expecter{mock: &_m.Mock}
}

// UpdateSecretPushProtectionEnabledSetting provides a mock function with given fields: gid, opt, options
func (_m *GroupSecuritySettingsServiceInterface) UpdateSecretPushProtectionEnabledSetting(gid interface{}, opt gitlab.UpdateGroupSecuritySettingsOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupSecuritySettings, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecretPushProtectionEnabledSetting")
	}

	var r0 *gitlab.GroupSecuritySettings
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.UpdateGroupSecuritySettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupSecuritySettings, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.UpdateGroupSecuritySettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupSecuritySettings); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupSecuritySettings)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, gitlab.UpdateGroupSecuritySettingsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, gitlab.UpdateGroupSecuritySettingsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSecretPushProtectionEnabledSetting'
type GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call struct {
	*mock.Call
}

// UpdateSecretPushProtectionEnabledSetting is a helper method to define mock.On call
//   - gid interface{}
//   - opt gitlab.UpdateGroupSecuritySettingsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupSecuritySettingsServiceInterface_Expecter) UpdateSecretPushProtectionEnabledSetting(gid interface{}, opt interface{}, options ...interface{}) *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	return &GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call{Call: _e.mock.On("UpdateSecretPushProtectionEnabledSetting",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) Run(run func(gid interface{}, opt gitlab.UpdateGroupSecuritySettingsOptions, options ...gitlab.RequestOptionFunc)) *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(gitlab.UpdateGroupSecuritySettingsOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) Return(_a0 *gitlab.GroupSecuritySettings, _a1 *gitlab.Response, _a2 error) *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call) RunAndReturn(run func(interface{}, gitlab.UpdateGroupSecuritySettingsOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupSecuritySettings, *gitlab.Response, error)) *GroupSecuritySettingsServiceInterface_UpdateSecretPushProtectionEnabledSetting_Call {
	_c.Call.Return(run)
	return _c
}

// NewGroupSecuritySettingsServiceInterface creates a new instance of GroupSecuritySettingsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupSecuritySettingsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupSecuritySettingsServiceInterface {
	mock := &GroupSecuritySettingsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

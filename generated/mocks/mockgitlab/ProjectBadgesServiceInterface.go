// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ProjectBadgesServiceInterface is an autogenerated mock type for the ProjectBadgesServiceInterface type
type ProjectBadgesServiceInterface struct {
	mock.Mock
}

type ProjectBadgesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectBadgesServiceInterface) EXPECT() *ProjectBadgesServiceInterface_Expecter {
	return &ProjectBadgesServiceInterface_Expecter{mock: &_m.Mock}
}

// AddProjectBadge provides a mock function with given fields: pid, opt, options
func (_m *ProjectBadgesServiceInterface) AddProjectBadge(pid interface{}, opt *gitlab.AddProjectBadgeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddProjectBadge")
	}

	var r0 *gitlab.ProjectBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddProjectBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddProjectBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectBadge); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.AddProjectBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.AddProjectBadgeOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectBadgesServiceInterface_AddProjectBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProjectBadge'
type ProjectBadgesServiceInterface_AddProjectBadge_Call struct {
	*mock.Call
}

// AddProjectBadge is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.AddProjectBadgeOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) AddProjectBadge(pid interface{}, opt interface{}, options ...interface{}) *ProjectBadgesServiceInterface_AddProjectBadge_Call {
	return &ProjectBadgesServiceInterface_AddProjectBadge_Call{Call: _e.mock.On("AddProjectBadge",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_AddProjectBadge_Call) Run(run func(pid interface{}, opt *gitlab.AddProjectBadgeOptions, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_AddProjectBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.AddProjectBadgeOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectBadgesServiceInterface_AddProjectBadge_Call) Return(_a0 *gitlab.ProjectBadge, _a1 *gitlab.Response, _a2 error) *ProjectBadgesServiceInterface_AddProjectBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectBadgesServiceInterface_AddProjectBadge_Call) RunAndReturn(run func(interface{}, *gitlab.AddProjectBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)) *ProjectBadgesServiceInterface_AddProjectBadge_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectBadge provides a mock function with given fields: pid, badge, options
func (_m *ProjectBadgesServiceInterface) DeleteProjectBadge(pid interface{}, badge int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, badge)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectBadge")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, badge, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, badge, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, badge, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectBadgesServiceInterface_DeleteProjectBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectBadge'
type ProjectBadgesServiceInterface_DeleteProjectBadge_Call struct {
	*mock.Call
}

// DeleteProjectBadge is a helper method to define mock.On call
//   - pid interface{}
//   - badge int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) DeleteProjectBadge(pid interface{}, badge interface{}, options ...interface{}) *ProjectBadgesServiceInterface_DeleteProjectBadge_Call {
	return &ProjectBadgesServiceInterface_DeleteProjectBadge_Call{Call: _e.mock.On("DeleteProjectBadge",
		append([]interface{}{pid, badge}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_DeleteProjectBadge_Call) Run(run func(pid interface{}, badge int, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_DeleteProjectBadge_Call {
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

func (_c *ProjectBadgesServiceInterface_DeleteProjectBadge_Call) Return(_a0 *gitlab.Response, _a1 error) *ProjectBadgesServiceInterface_DeleteProjectBadge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectBadgesServiceInterface_DeleteProjectBadge_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ProjectBadgesServiceInterface_DeleteProjectBadge_Call {
	_c.Call.Return(run)
	return _c
}

// EditProjectBadge provides a mock function with given fields: pid, badge, opt, options
func (_m *ProjectBadgesServiceInterface) EditProjectBadge(pid interface{}, badge int, opt *gitlab.EditProjectBadgeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, badge, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EditProjectBadge")
	}

	var r0 *gitlab.ProjectBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditProjectBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)); ok {
		return rf(pid, badge, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditProjectBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectBadge); ok {
		r0 = rf(pid, badge, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.EditProjectBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, badge, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.EditProjectBadgeOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, badge, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectBadgesServiceInterface_EditProjectBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditProjectBadge'
type ProjectBadgesServiceInterface_EditProjectBadge_Call struct {
	*mock.Call
}

// EditProjectBadge is a helper method to define mock.On call
//   - pid interface{}
//   - badge int
//   - opt *gitlab.EditProjectBadgeOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) EditProjectBadge(pid interface{}, badge interface{}, opt interface{}, options ...interface{}) *ProjectBadgesServiceInterface_EditProjectBadge_Call {
	return &ProjectBadgesServiceInterface_EditProjectBadge_Call{Call: _e.mock.On("EditProjectBadge",
		append([]interface{}{pid, badge, opt}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_EditProjectBadge_Call) Run(run func(pid interface{}, badge int, opt *gitlab.EditProjectBadgeOptions, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_EditProjectBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.EditProjectBadgeOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectBadgesServiceInterface_EditProjectBadge_Call) Return(_a0 *gitlab.ProjectBadge, _a1 *gitlab.Response, _a2 error) *ProjectBadgesServiceInterface_EditProjectBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectBadgesServiceInterface_EditProjectBadge_Call) RunAndReturn(run func(interface{}, int, *gitlab.EditProjectBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)) *ProjectBadgesServiceInterface_EditProjectBadge_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectBadge provides a mock function with given fields: pid, badge, options
func (_m *ProjectBadgesServiceInterface) GetProjectBadge(pid interface{}, badge int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, badge)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectBadge")
	}

	var r0 *gitlab.ProjectBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)); ok {
		return rf(pid, badge, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.ProjectBadge); ok {
		r0 = rf(pid, badge, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, badge, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, badge, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectBadgesServiceInterface_GetProjectBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectBadge'
type ProjectBadgesServiceInterface_GetProjectBadge_Call struct {
	*mock.Call
}

// GetProjectBadge is a helper method to define mock.On call
//   - pid interface{}
//   - badge int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) GetProjectBadge(pid interface{}, badge interface{}, options ...interface{}) *ProjectBadgesServiceInterface_GetProjectBadge_Call {
	return &ProjectBadgesServiceInterface_GetProjectBadge_Call{Call: _e.mock.On("GetProjectBadge",
		append([]interface{}{pid, badge}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_GetProjectBadge_Call) Run(run func(pid interface{}, badge int, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_GetProjectBadge_Call {
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

func (_c *ProjectBadgesServiceInterface_GetProjectBadge_Call) Return(_a0 *gitlab.ProjectBadge, _a1 *gitlab.Response, _a2 error) *ProjectBadgesServiceInterface_GetProjectBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectBadgesServiceInterface_GetProjectBadge_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)) *ProjectBadgesServiceInterface_GetProjectBadge_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectBadges provides a mock function with given fields: pid, opt, options
func (_m *ProjectBadgesServiceInterface) ListProjectBadges(pid interface{}, opt *gitlab.ListProjectBadgesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectBadges")
	}

	var r0 []*gitlab.ProjectBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectBadgesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectBadge, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectBadgesOptions, ...gitlab.RequestOptionFunc) []*gitlab.ProjectBadge); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ProjectBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectBadgesOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectBadgesOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectBadgesServiceInterface_ListProjectBadges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectBadges'
type ProjectBadgesServiceInterface_ListProjectBadges_Call struct {
	*mock.Call
}

// ListProjectBadges is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListProjectBadgesOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) ListProjectBadges(pid interface{}, opt interface{}, options ...interface{}) *ProjectBadgesServiceInterface_ListProjectBadges_Call {
	return &ProjectBadgesServiceInterface_ListProjectBadges_Call{Call: _e.mock.On("ListProjectBadges",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_ListProjectBadges_Call) Run(run func(pid interface{}, opt *gitlab.ListProjectBadgesOptions, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_ListProjectBadges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListProjectBadgesOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectBadgesServiceInterface_ListProjectBadges_Call) Return(_a0 []*gitlab.ProjectBadge, _a1 *gitlab.Response, _a2 error) *ProjectBadgesServiceInterface_ListProjectBadges_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectBadgesServiceInterface_ListProjectBadges_Call) RunAndReturn(run func(interface{}, *gitlab.ListProjectBadgesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectBadge, *gitlab.Response, error)) *ProjectBadgesServiceInterface_ListProjectBadges_Call {
	_c.Call.Return(run)
	return _c
}

// PreviewProjectBadge provides a mock function with given fields: pid, opt, options
func (_m *ProjectBadgesServiceInterface) PreviewProjectBadge(pid interface{}, opt *gitlab.ProjectBadgePreviewOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PreviewProjectBadge")
	}

	var r0 *gitlab.ProjectBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectBadgePreviewOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ProjectBadgePreviewOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectBadge); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ProjectBadgePreviewOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ProjectBadgePreviewOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectBadgesServiceInterface_PreviewProjectBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PreviewProjectBadge'
type ProjectBadgesServiceInterface_PreviewProjectBadge_Call struct {
	*mock.Call
}

// PreviewProjectBadge is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ProjectBadgePreviewOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectBadgesServiceInterface_Expecter) PreviewProjectBadge(pid interface{}, opt interface{}, options ...interface{}) *ProjectBadgesServiceInterface_PreviewProjectBadge_Call {
	return &ProjectBadgesServiceInterface_PreviewProjectBadge_Call{Call: _e.mock.On("PreviewProjectBadge",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectBadgesServiceInterface_PreviewProjectBadge_Call) Run(run func(pid interface{}, opt *gitlab.ProjectBadgePreviewOptions, options ...gitlab.RequestOptionFunc)) *ProjectBadgesServiceInterface_PreviewProjectBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ProjectBadgePreviewOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectBadgesServiceInterface_PreviewProjectBadge_Call) Return(_a0 *gitlab.ProjectBadge, _a1 *gitlab.Response, _a2 error) *ProjectBadgesServiceInterface_PreviewProjectBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectBadgesServiceInterface_PreviewProjectBadge_Call) RunAndReturn(run func(interface{}, *gitlab.ProjectBadgePreviewOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectBadge, *gitlab.Response, error)) *ProjectBadgesServiceInterface_PreviewProjectBadge_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectBadgesServiceInterface creates a new instance of ProjectBadgesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectBadgesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectBadgesServiceInterface {
	mock := &ProjectBadgesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

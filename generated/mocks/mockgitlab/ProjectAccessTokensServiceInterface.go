// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ProjectAccessTokensServiceInterface is an autogenerated mock type for the ProjectAccessTokensServiceInterface type
type ProjectAccessTokensServiceInterface struct {
	mock.Mock
}

type ProjectAccessTokensServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectAccessTokensServiceInterface) EXPECT() *ProjectAccessTokensServiceInterface_Expecter {
	return &ProjectAccessTokensServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateProjectAccessToken provides a mock function with given fields: pid, opt, options
func (_m *ProjectAccessTokensServiceInterface) CreateProjectAccessToken(pid interface{}, opt *gitlab.CreateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectAccessToken")
	}

	var r0 *gitlab.ProjectAccessToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectAccessToken); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectAccessToken'
type ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call struct {
	*mock.Call
}

// CreateProjectAccessToken is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateProjectAccessTokenOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) CreateProjectAccessToken(pid interface{}, opt interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call {
	return &ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call{Call: _e.mock.On("CreateProjectAccessToken",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call) Run(run func(pid interface{}, opt *gitlab.CreateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateProjectAccessTokenOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call) Return(_a0 *gitlab.ProjectAccessToken, _a1 *gitlab.Response, _a2 error) *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call) RunAndReturn(run func(interface{}, *gitlab.CreateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *ProjectAccessTokensServiceInterface_CreateProjectAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectAccessToken provides a mock function with given fields: pid, id, options
func (_m *ProjectAccessTokensServiceInterface) GetProjectAccessToken(pid interface{}, id int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectAccessToken")
	}

	var r0 *gitlab.ProjectAccessToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)); ok {
		return rf(pid, id, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.ProjectAccessToken); ok {
		r0 = rf(pid, id, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, id, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, id, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectAccessToken'
type ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call struct {
	*mock.Call
}

// GetProjectAccessToken is a helper method to define mock.On call
//   - pid interface{}
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) GetProjectAccessToken(pid interface{}, id interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call {
	return &ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call{Call: _e.mock.On("GetProjectAccessToken",
		append([]interface{}{pid, id}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call) Run(run func(pid interface{}, id int, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call {
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

func (_c *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call) Return(_a0 *gitlab.ProjectAccessToken, _a1 *gitlab.Response, _a2 error) *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *ProjectAccessTokensServiceInterface_GetProjectAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectAccessTokens provides a mock function with given fields: pid, opt, options
func (_m *ProjectAccessTokensServiceInterface) ListProjectAccessTokens(pid interface{}, opt *gitlab.ListProjectAccessTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectAccessTokens")
	}

	var r0 []*gitlab.ProjectAccessToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) []*gitlab.ProjectAccessToken); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ProjectAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectAccessTokens'
type ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call struct {
	*mock.Call
}

// ListProjectAccessTokens is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListProjectAccessTokensOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) ListProjectAccessTokens(pid interface{}, opt interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call {
	return &ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call{Call: _e.mock.On("ListProjectAccessTokens",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call) Run(run func(pid interface{}, opt *gitlab.ListProjectAccessTokensOptions, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListProjectAccessTokensOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call) Return(_a0 []*gitlab.ProjectAccessToken, _a1 *gitlab.Response, _a2 error) *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call) RunAndReturn(run func(interface{}, *gitlab.ListProjectAccessTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectAccessToken, *gitlab.Response, error)) *ProjectAccessTokensServiceInterface_ListProjectAccessTokens_Call {
	_c.Call.Return(run)
	return _c
}

// RevokeProjectAccessToken provides a mock function with given fields: pid, id, options
func (_m *ProjectAccessTokensServiceInterface) RevokeProjectAccessToken(pid interface{}, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RevokeProjectAccessToken")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, id, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, id, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, id, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeProjectAccessToken'
type ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call struct {
	*mock.Call
}

// RevokeProjectAccessToken is a helper method to define mock.On call
//   - pid interface{}
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) RevokeProjectAccessToken(pid interface{}, id interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call {
	return &ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call{Call: _e.mock.On("RevokeProjectAccessToken",
		append([]interface{}{pid, id}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call) Run(run func(pid interface{}, id int, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call {
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

func (_c *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call) Return(_a0 *gitlab.Response, _a1 error) *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ProjectAccessTokensServiceInterface_RevokeProjectAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// RotateProjectAccessToken provides a mock function with given fields: pid, id, opt, options
func (_m *ProjectAccessTokensServiceInterface) RotateProjectAccessToken(pid interface{}, id int, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, id, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RotateProjectAccessToken")
	}

	var r0 *gitlab.ProjectAccessToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)); ok {
		return rf(pid, id, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectAccessToken); ok {
		r0 = rf(pid, id, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, id, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, id, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RotateProjectAccessToken'
type ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call struct {
	*mock.Call
}

// RotateProjectAccessToken is a helper method to define mock.On call
//   - pid interface{}
//   - id int
//   - opt *gitlab.RotateProjectAccessTokenOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) RotateProjectAccessToken(pid interface{}, id interface{}, opt interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call {
	return &ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call{Call: _e.mock.On("RotateProjectAccessToken",
		append([]interface{}{pid, id, opt}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call) Run(run func(pid interface{}, id int, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.RotateProjectAccessTokenOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call) Return(_a0 *gitlab.ProjectAccessToken, _a1 *gitlab.Response, _a2 error) *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call) RunAndReturn(run func(interface{}, int, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *ProjectAccessTokensServiceInterface_RotateProjectAccessToken_Call {
	_c.Call.Return(run)
	return _c
}

// RotateProjectAccessTokenSelf provides a mock function with given fields: pid, opt, options
func (_m *ProjectAccessTokensServiceInterface) RotateProjectAccessTokenSelf(pid interface{}, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RotateProjectAccessTokenSelf")
	}

	var r0 *gitlab.ProjectAccessToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectAccessToken); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectAccessToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RotateProjectAccessTokenSelf'
type ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call struct {
	*mock.Call
}

// RotateProjectAccessTokenSelf is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.RotateProjectAccessTokenOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectAccessTokensServiceInterface_Expecter) RotateProjectAccessTokenSelf(pid interface{}, opt interface{}, options ...interface{}) *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call {
	return &ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call{Call: _e.mock.On("RotateProjectAccessTokenSelf",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call) Run(run func(pid interface{}, opt *gitlab.RotateProjectAccessTokenOptions, options ...gitlab.RequestOptionFunc)) *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.RotateProjectAccessTokenOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call) Return(_a0 *gitlab.ProjectAccessToken, _a1 *gitlab.Response, _a2 error) *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call) RunAndReturn(run func(interface{}, *gitlab.RotateProjectAccessTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectAccessToken, *gitlab.Response, error)) *ProjectAccessTokensServiceInterface_RotateProjectAccessTokenSelf_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectAccessTokensServiceInterface creates a new instance of ProjectAccessTokensServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectAccessTokensServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectAccessTokensServiceInterface {
	mock := &ProjectAccessTokensServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

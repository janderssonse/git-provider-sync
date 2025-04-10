// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ProjectSnippetsServiceInterface is an autogenerated mock type for the ProjectSnippetsServiceInterface type
type ProjectSnippetsServiceInterface struct {
	mock.Mock
}

type ProjectSnippetsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectSnippetsServiceInterface) EXPECT() *ProjectSnippetsServiceInterface_Expecter {
	return &ProjectSnippetsServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateSnippet provides a mock function with given fields: pid, opt, options
func (_m *ProjectSnippetsServiceInterface) CreateSnippet(pid interface{}, opt *gitlab.CreateProjectSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateSnippet")
	}

	var r0 *gitlab.Snippet
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) *gitlab.Snippet); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Snippet)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSnippetsServiceInterface_CreateSnippet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateSnippet'
type ProjectSnippetsServiceInterface_CreateSnippet_Call struct {
	*mock.Call
}

// CreateSnippet is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateProjectSnippetOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) CreateSnippet(pid interface{}, opt interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_CreateSnippet_Call {
	return &ProjectSnippetsServiceInterface_CreateSnippet_Call{Call: _e.mock.On("CreateSnippet",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_CreateSnippet_Call) Run(run func(pid interface{}, opt *gitlab.CreateProjectSnippetOptions, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_CreateSnippet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateProjectSnippetOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectSnippetsServiceInterface_CreateSnippet_Call) Return(_a0 *gitlab.Snippet, _a1 *gitlab.Response, _a2 error) *ProjectSnippetsServiceInterface_CreateSnippet_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_CreateSnippet_Call) RunAndReturn(run func(interface{}, *gitlab.CreateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *ProjectSnippetsServiceInterface_CreateSnippet_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSnippet provides a mock function with given fields: pid, snippet, options
func (_m *ProjectSnippetsServiceInterface) DeleteSnippet(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, snippet)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSnippet")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, snippet, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, snippet, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, snippet, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProjectSnippetsServiceInterface_DeleteSnippet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSnippet'
type ProjectSnippetsServiceInterface_DeleteSnippet_Call struct {
	*mock.Call
}

// DeleteSnippet is a helper method to define mock.On call
//   - pid interface{}
//   - snippet int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) DeleteSnippet(pid interface{}, snippet interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_DeleteSnippet_Call {
	return &ProjectSnippetsServiceInterface_DeleteSnippet_Call{Call: _e.mock.On("DeleteSnippet",
		append([]interface{}{pid, snippet}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_DeleteSnippet_Call) Run(run func(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_DeleteSnippet_Call {
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

func (_c *ProjectSnippetsServiceInterface_DeleteSnippet_Call) Return(_a0 *gitlab.Response, _a1 error) *ProjectSnippetsServiceInterface_DeleteSnippet_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_DeleteSnippet_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ProjectSnippetsServiceInterface_DeleteSnippet_Call {
	_c.Call.Return(run)
	return _c
}

// GetSnippet provides a mock function with given fields: pid, snippet, options
func (_m *ProjectSnippetsServiceInterface) GetSnippet(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, snippet)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetSnippet")
	}

	var r0 *gitlab.Snippet
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)); ok {
		return rf(pid, snippet, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Snippet); ok {
		r0 = rf(pid, snippet, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Snippet)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, snippet, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, snippet, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSnippetsServiceInterface_GetSnippet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSnippet'
type ProjectSnippetsServiceInterface_GetSnippet_Call struct {
	*mock.Call
}

// GetSnippet is a helper method to define mock.On call
//   - pid interface{}
//   - snippet int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) GetSnippet(pid interface{}, snippet interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_GetSnippet_Call {
	return &ProjectSnippetsServiceInterface_GetSnippet_Call{Call: _e.mock.On("GetSnippet",
		append([]interface{}{pid, snippet}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_GetSnippet_Call) Run(run func(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_GetSnippet_Call {
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

func (_c *ProjectSnippetsServiceInterface_GetSnippet_Call) Return(_a0 *gitlab.Snippet, _a1 *gitlab.Response, _a2 error) *ProjectSnippetsServiceInterface_GetSnippet_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_GetSnippet_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *ProjectSnippetsServiceInterface_GetSnippet_Call {
	_c.Call.Return(run)
	return _c
}

// ListSnippets provides a mock function with given fields: pid, opt, options
func (_m *ProjectSnippetsServiceInterface) ListSnippets(pid interface{}, opt *gitlab.ListProjectSnippetsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListSnippets")
	}

	var r0 []*gitlab.Snippet
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) []*gitlab.Snippet); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Snippet)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSnippetsServiceInterface_ListSnippets_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSnippets'
type ProjectSnippetsServiceInterface_ListSnippets_Call struct {
	*mock.Call
}

// ListSnippets is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListProjectSnippetsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) ListSnippets(pid interface{}, opt interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_ListSnippets_Call {
	return &ProjectSnippetsServiceInterface_ListSnippets_Call{Call: _e.mock.On("ListSnippets",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_ListSnippets_Call) Run(run func(pid interface{}, opt *gitlab.ListProjectSnippetsOptions, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_ListSnippets_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListProjectSnippetsOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectSnippetsServiceInterface_ListSnippets_Call) Return(_a0 []*gitlab.Snippet, _a1 *gitlab.Response, _a2 error) *ProjectSnippetsServiceInterface_ListSnippets_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_ListSnippets_Call) RunAndReturn(run func(interface{}, *gitlab.ListProjectSnippetsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Snippet, *gitlab.Response, error)) *ProjectSnippetsServiceInterface_ListSnippets_Call {
	_c.Call.Return(run)
	return _c
}

// SnippetContent provides a mock function with given fields: pid, snippet, options
func (_m *ProjectSnippetsServiceInterface) SnippetContent(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, snippet)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SnippetContent")
	}

	var r0 []byte
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)); ok {
		return rf(pid, snippet, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) []byte); ok {
		r0 = rf(pid, snippet, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, snippet, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, snippet, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSnippetsServiceInterface_SnippetContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SnippetContent'
type ProjectSnippetsServiceInterface_SnippetContent_Call struct {
	*mock.Call
}

// SnippetContent is a helper method to define mock.On call
//   - pid interface{}
//   - snippet int
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) SnippetContent(pid interface{}, snippet interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_SnippetContent_Call {
	return &ProjectSnippetsServiceInterface_SnippetContent_Call{Call: _e.mock.On("SnippetContent",
		append([]interface{}{pid, snippet}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_SnippetContent_Call) Run(run func(pid interface{}, snippet int, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_SnippetContent_Call {
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

func (_c *ProjectSnippetsServiceInterface_SnippetContent_Call) Return(_a0 []byte, _a1 *gitlab.Response, _a2 error) *ProjectSnippetsServiceInterface_SnippetContent_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_SnippetContent_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *ProjectSnippetsServiceInterface_SnippetContent_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSnippet provides a mock function with given fields: pid, snippet, opt, options
func (_m *ProjectSnippetsServiceInterface) UpdateSnippet(pid interface{}, snippet int, opt *gitlab.UpdateProjectSnippetOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, snippet, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSnippet")
	}

	var r0 *gitlab.Snippet
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)); ok {
		return rf(pid, snippet, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) *gitlab.Snippet); ok {
		r0 = rf(pid, snippet, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Snippet)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, snippet, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, snippet, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectSnippetsServiceInterface_UpdateSnippet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSnippet'
type ProjectSnippetsServiceInterface_UpdateSnippet_Call struct {
	*mock.Call
}

// UpdateSnippet is a helper method to define mock.On call
//   - pid interface{}
//   - snippet int
//   - opt *gitlab.UpdateProjectSnippetOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectSnippetsServiceInterface_Expecter) UpdateSnippet(pid interface{}, snippet interface{}, opt interface{}, options ...interface{}) *ProjectSnippetsServiceInterface_UpdateSnippet_Call {
	return &ProjectSnippetsServiceInterface_UpdateSnippet_Call{Call: _e.mock.On("UpdateSnippet",
		append([]interface{}{pid, snippet, opt}, options...)...)}
}

func (_c *ProjectSnippetsServiceInterface_UpdateSnippet_Call) Run(run func(pid interface{}, snippet int, opt *gitlab.UpdateProjectSnippetOptions, options ...gitlab.RequestOptionFunc)) *ProjectSnippetsServiceInterface_UpdateSnippet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.UpdateProjectSnippetOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectSnippetsServiceInterface_UpdateSnippet_Call) Return(_a0 *gitlab.Snippet, _a1 *gitlab.Response, _a2 error) *ProjectSnippetsServiceInterface_UpdateSnippet_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectSnippetsServiceInterface_UpdateSnippet_Call) RunAndReturn(run func(interface{}, int, *gitlab.UpdateProjectSnippetOptions, ...gitlab.RequestOptionFunc) (*gitlab.Snippet, *gitlab.Response, error)) *ProjectSnippetsServiceInterface_UpdateSnippet_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectSnippetsServiceInterface creates a new instance of ProjectSnippetsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectSnippetsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectSnippetsServiceInterface {
	mock := &ProjectSnippetsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

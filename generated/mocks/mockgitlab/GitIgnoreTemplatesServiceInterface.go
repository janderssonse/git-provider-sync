// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GitIgnoreTemplatesServiceInterface is an autogenerated mock type for the GitIgnoreTemplatesServiceInterface type
type GitIgnoreTemplatesServiceInterface struct {
	mock.Mock
}

type GitIgnoreTemplatesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GitIgnoreTemplatesServiceInterface) EXPECT() *GitIgnoreTemplatesServiceInterface_Expecter {
	return &GitIgnoreTemplatesServiceInterface_Expecter{mock: &_m.Mock}
}

// GetTemplate provides a mock function with given fields: _a0, _a1
func (_m *GitIgnoreTemplatesServiceInterface) GetTemplate(_a0 string, _a1 ...gitlab.RequestOptionFunc) (*gitlab.GitIgnoreTemplate, *gitlab.Response, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplate")
	}

	var r0 *gitlab.GitIgnoreTemplate
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(string, ...gitlab.RequestOptionFunc) (*gitlab.GitIgnoreTemplate, *gitlab.Response, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(string, ...gitlab.RequestOptionFunc) *gitlab.GitIgnoreTemplate); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GitIgnoreTemplate)
		}
	}

	if rf, ok := ret.Get(1).(func(string, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(_a0, _a1...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(string, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(_a0, _a1...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GitIgnoreTemplatesServiceInterface_GetTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplate'
type GitIgnoreTemplatesServiceInterface_GetTemplate_Call struct {
	*mock.Call
}

// GetTemplate is a helper method to define mock.On call
//   - _a0 string
//   - _a1 ...gitlab.RequestOptionFunc
func (_e *GitIgnoreTemplatesServiceInterface_Expecter) GetTemplate(_a0 interface{}, _a1 ...interface{}) *GitIgnoreTemplatesServiceInterface_GetTemplate_Call {
	return &GitIgnoreTemplatesServiceInterface_GetTemplate_Call{Call: _e.mock.On("GetTemplate",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *GitIgnoreTemplatesServiceInterface_GetTemplate_Call) Run(run func(_a0 string, _a1 ...gitlab.RequestOptionFunc)) *GitIgnoreTemplatesServiceInterface_GetTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *GitIgnoreTemplatesServiceInterface_GetTemplate_Call) Return(_a0 *gitlab.GitIgnoreTemplate, _a1 *gitlab.Response, _a2 error) *GitIgnoreTemplatesServiceInterface_GetTemplate_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GitIgnoreTemplatesServiceInterface_GetTemplate_Call) RunAndReturn(run func(string, ...gitlab.RequestOptionFunc) (*gitlab.GitIgnoreTemplate, *gitlab.Response, error)) *GitIgnoreTemplatesServiceInterface_GetTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// ListTemplates provides a mock function with given fields: _a0, _a1
func (_m *GitIgnoreTemplatesServiceInterface) ListTemplates(_a0 *gitlab.ListTemplatesOptions, _a1 ...gitlab.RequestOptionFunc) ([]*gitlab.GitIgnoreTemplateListItem, *gitlab.Response, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTemplates")
	}

	var r0 []*gitlab.GitIgnoreTemplateListItem
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(*gitlab.ListTemplatesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GitIgnoreTemplateListItem, *gitlab.Response, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(*gitlab.ListTemplatesOptions, ...gitlab.RequestOptionFunc) []*gitlab.GitIgnoreTemplateListItem); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.GitIgnoreTemplateListItem)
		}
	}

	if rf, ok := ret.Get(1).(func(*gitlab.ListTemplatesOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(_a0, _a1...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(*gitlab.ListTemplatesOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(_a0, _a1...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GitIgnoreTemplatesServiceInterface_ListTemplates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTemplates'
type GitIgnoreTemplatesServiceInterface_ListTemplates_Call struct {
	*mock.Call
}

// ListTemplates is a helper method to define mock.On call
//   - _a0 *gitlab.ListTemplatesOptions
//   - _a1 ...gitlab.RequestOptionFunc
func (_e *GitIgnoreTemplatesServiceInterface_Expecter) ListTemplates(_a0 interface{}, _a1 ...interface{}) *GitIgnoreTemplatesServiceInterface_ListTemplates_Call {
	return &GitIgnoreTemplatesServiceInterface_ListTemplates_Call{Call: _e.mock.On("ListTemplates",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *GitIgnoreTemplatesServiceInterface_ListTemplates_Call) Run(run func(_a0 *gitlab.ListTemplatesOptions, _a1 ...gitlab.RequestOptionFunc)) *GitIgnoreTemplatesServiceInterface_ListTemplates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(*gitlab.ListTemplatesOptions), variadicArgs...)
	})
	return _c
}

func (_c *GitIgnoreTemplatesServiceInterface_ListTemplates_Call) Return(_a0 []*gitlab.GitIgnoreTemplateListItem, _a1 *gitlab.Response, _a2 error) *GitIgnoreTemplatesServiceInterface_ListTemplates_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GitIgnoreTemplatesServiceInterface_ListTemplates_Call) RunAndReturn(run func(*gitlab.ListTemplatesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GitIgnoreTemplateListItem, *gitlab.Response, error)) *GitIgnoreTemplatesServiceInterface_ListTemplates_Call {
	_c.Call.Return(run)
	return _c
}

// NewGitIgnoreTemplatesServiceInterface creates a new instance of GitIgnoreTemplatesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGitIgnoreTemplatesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GitIgnoreTemplatesServiceInterface {
	mock := &GitIgnoreTemplatesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

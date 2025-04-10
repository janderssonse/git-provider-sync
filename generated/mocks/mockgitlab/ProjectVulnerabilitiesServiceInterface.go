// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ProjectVulnerabilitiesServiceInterface is an autogenerated mock type for the ProjectVulnerabilitiesServiceInterface type
type ProjectVulnerabilitiesServiceInterface struct {
	mock.Mock
}

type ProjectVulnerabilitiesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ProjectVulnerabilitiesServiceInterface) EXPECT() *ProjectVulnerabilitiesServiceInterface_Expecter {
	return &ProjectVulnerabilitiesServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateVulnerability provides a mock function with given fields: pid, opt, options
func (_m *ProjectVulnerabilitiesServiceInterface) CreateVulnerability(pid interface{}, opt *gitlab.CreateVulnerabilityOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateVulnerability")
	}

	var r0 *gitlab.ProjectVulnerability
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) *gitlab.ProjectVulnerability); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ProjectVulnerability)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVulnerability'
type ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call struct {
	*mock.Call
}

// CreateVulnerability is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateVulnerabilityOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectVulnerabilitiesServiceInterface_Expecter) CreateVulnerability(pid interface{}, opt interface{}, options ...interface{}) *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call {
	return &ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call{Call: _e.mock.On("CreateVulnerability",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call) Run(run func(pid interface{}, opt *gitlab.CreateVulnerabilityOptions, options ...gitlab.RequestOptionFunc)) *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateVulnerabilityOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call) Return(_a0 *gitlab.ProjectVulnerability, _a1 *gitlab.Response, _a2 error) *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call) RunAndReturn(run func(interface{}, *gitlab.CreateVulnerabilityOptions, ...gitlab.RequestOptionFunc) (*gitlab.ProjectVulnerability, *gitlab.Response, error)) *ProjectVulnerabilitiesServiceInterface_CreateVulnerability_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectVulnerabilities provides a mock function with given fields: pid, opt, options
func (_m *ProjectVulnerabilitiesServiceInterface) ListProjectVulnerabilities(pid interface{}, opt *gitlab.ListProjectVulnerabilitiesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectVulnerabilities")
	}

	var r0 []*gitlab.ProjectVulnerability
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) []*gitlab.ProjectVulnerability); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ProjectVulnerability)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectVulnerabilities'
type ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call struct {
	*mock.Call
}

// ListProjectVulnerabilities is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListProjectVulnerabilitiesOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ProjectVulnerabilitiesServiceInterface_Expecter) ListProjectVulnerabilities(pid interface{}, opt interface{}, options ...interface{}) *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call {
	return &ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call{Call: _e.mock.On("ListProjectVulnerabilities",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call) Run(run func(pid interface{}, opt *gitlab.ListProjectVulnerabilitiesOptions, options ...gitlab.RequestOptionFunc)) *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListProjectVulnerabilitiesOptions), variadicArgs...)
	})
	return _c
}

func (_c *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call) Return(_a0 []*gitlab.ProjectVulnerability, _a1 *gitlab.Response, _a2 error) *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call) RunAndReturn(run func(interface{}, *gitlab.ListProjectVulnerabilitiesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.ProjectVulnerability, *gitlab.Response, error)) *ProjectVulnerabilitiesServiceInterface_ListProjectVulnerabilities_Call {
	_c.Call.Return(run)
	return _c
}

// NewProjectVulnerabilitiesServiceInterface creates a new instance of ProjectVulnerabilitiesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProjectVulnerabilitiesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProjectVulnerabilitiesServiceInterface {
	mock := &ProjectVulnerabilitiesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

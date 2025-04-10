// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ResourceGroupServiceInterface is an autogenerated mock type for the ResourceGroupServiceInterface type
type ResourceGroupServiceInterface struct {
	mock.Mock
}

type ResourceGroupServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceGroupServiceInterface) EXPECT() *ResourceGroupServiceInterface_Expecter {
	return &ResourceGroupServiceInterface_Expecter{mock: &_m.Mock}
}

// EditAnExistingResourceGroup provides a mock function with given fields: pid, key, opts, options
func (_m *ResourceGroupServiceInterface) EditAnExistingResourceGroup(pid interface{}, key string, opts *gitlab.EditAnExistingResourceGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, key, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EditAnExistingResourceGroup")
	}

	var r0 *gitlab.ResourceGroup
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)); ok {
		return rf(pid, key, opts, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) *gitlab.ResourceGroup); ok {
		r0 = rf(pid, key, opts, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ResourceGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, key, opts, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, key, opts, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditAnExistingResourceGroup'
type ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call struct {
	*mock.Call
}

// EditAnExistingResourceGroup is a helper method to define mock.On call
//   - pid interface{}
//   - key string
//   - opts *gitlab.EditAnExistingResourceGroupOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ResourceGroupServiceInterface_Expecter) EditAnExistingResourceGroup(pid interface{}, key interface{}, opts interface{}, options ...interface{}) *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call {
	return &ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call{Call: _e.mock.On("EditAnExistingResourceGroup",
		append([]interface{}{pid, key, opts}, options...)...)}
}

func (_c *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call) Run(run func(pid interface{}, key string, opts *gitlab.EditAnExistingResourceGroupOptions, options ...gitlab.RequestOptionFunc)) *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), args[2].(*gitlab.EditAnExistingResourceGroupOptions), variadicArgs...)
	})
	return _c
}

func (_c *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call) Return(_a0 *gitlab.ResourceGroup, _a1 *gitlab.Response, _a2 error) *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call) RunAndReturn(run func(interface{}, string, *gitlab.EditAnExistingResourceGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *ResourceGroupServiceInterface_EditAnExistingResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// GetASpecificResourceGroup provides a mock function with given fields: pid, key, options
func (_m *ResourceGroupServiceInterface) GetASpecificResourceGroup(pid interface{}, key string, options ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetASpecificResourceGroup")
	}

	var r0 *gitlab.ResourceGroup
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)); ok {
		return rf(pid, key, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) *gitlab.ResourceGroup); ok {
		r0 = rf(pid, key, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ResourceGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, key, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, string, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, key, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResourceGroupServiceInterface_GetASpecificResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetASpecificResourceGroup'
type ResourceGroupServiceInterface_GetASpecificResourceGroup_Call struct {
	*mock.Call
}

// GetASpecificResourceGroup is a helper method to define mock.On call
//   - pid interface{}
//   - key string
//   - options ...gitlab.RequestOptionFunc
func (_e *ResourceGroupServiceInterface_Expecter) GetASpecificResourceGroup(pid interface{}, key interface{}, options ...interface{}) *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call {
	return &ResourceGroupServiceInterface_GetASpecificResourceGroup_Call{Call: _e.mock.On("GetASpecificResourceGroup",
		append([]interface{}{pid, key}, options...)...)}
}

func (_c *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call) Run(run func(pid interface{}, key string, options ...gitlab.RequestOptionFunc)) *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call) Return(_a0 *gitlab.ResourceGroup, _a1 *gitlab.Response, _a2 error) *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call) RunAndReturn(run func(interface{}, string, ...gitlab.RequestOptionFunc) (*gitlab.ResourceGroup, *gitlab.Response, error)) *ResourceGroupServiceInterface_GetASpecificResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllResourceGroupsForAProject provides a mock function with given fields: pid, options
func (_m *ResourceGroupServiceInterface) GetAllResourceGroupsForAProject(pid interface{}, options ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAllResourceGroupsForAProject")
	}

	var r0 []*gitlab.ResourceGroup
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error)); ok {
		return rf(pid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) []*gitlab.ResourceGroup); ok {
		r0 = rf(pid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ResourceGroup)
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

// ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllResourceGroupsForAProject'
type ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call struct {
	*mock.Call
}

// GetAllResourceGroupsForAProject is a helper method to define mock.On call
//   - pid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *ResourceGroupServiceInterface_Expecter) GetAllResourceGroupsForAProject(pid interface{}, options ...interface{}) *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call {
	return &ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call{Call: _e.mock.On("GetAllResourceGroupsForAProject",
		append([]interface{}{pid}, options...)...)}
}

func (_c *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call) Run(run func(pid interface{}, options ...gitlab.RequestOptionFunc)) *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call {
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

func (_c *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call) Return(_a0 []*gitlab.ResourceGroup, _a1 *gitlab.Response, _a2 error) *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.ResourceGroup, *gitlab.Response, error)) *ResourceGroupServiceInterface_GetAllResourceGroupsForAProject_Call {
	_c.Call.Return(run)
	return _c
}

// ListUpcomingJobsForASpecificResourceGroup provides a mock function with given fields: pid, key, options
func (_m *ResourceGroupServiceInterface) ListUpcomingJobsForASpecificResourceGroup(pid interface{}, key string, options ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListUpcomingJobsForASpecificResourceGroup")
	}

	var r0 []*gitlab.Job
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)); ok {
		return rf(pid, key, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) []*gitlab.Job); ok {
		r0 = rf(pid, key, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, key, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, string, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, key, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListUpcomingJobsForASpecificResourceGroup'
type ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call struct {
	*mock.Call
}

// ListUpcomingJobsForASpecificResourceGroup is a helper method to define mock.On call
//   - pid interface{}
//   - key string
//   - options ...gitlab.RequestOptionFunc
func (_e *ResourceGroupServiceInterface_Expecter) ListUpcomingJobsForASpecificResourceGroup(pid interface{}, key interface{}, options ...interface{}) *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call {
	return &ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call{Call: _e.mock.On("ListUpcomingJobsForASpecificResourceGroup",
		append([]interface{}{pid, key}, options...)...)}
}

func (_c *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call) Run(run func(pid interface{}, key string, options ...gitlab.RequestOptionFunc)) *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call) Return(_a0 []*gitlab.Job, _a1 *gitlab.Response, _a2 error) *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call) RunAndReturn(run func(interface{}, string, ...gitlab.RequestOptionFunc) ([]*gitlab.Job, *gitlab.Response, error)) *ResourceGroupServiceInterface_ListUpcomingJobsForASpecificResourceGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewResourceGroupServiceInterface creates a new instance of ResourceGroupServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewResourceGroupServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ResourceGroupServiceInterface {
	mock := &ResourceGroupServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

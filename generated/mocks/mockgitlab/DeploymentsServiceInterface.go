// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// DeploymentsServiceInterface is an autogenerated mock type for the DeploymentsServiceInterface type
type DeploymentsServiceInterface struct {
	mock.Mock
}

type DeploymentsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *DeploymentsServiceInterface) EXPECT() *DeploymentsServiceInterface_Expecter {
	return &DeploymentsServiceInterface_Expecter{mock: &_m.Mock}
}

// ApproveOrRejectProjectDeployment provides a mock function with given fields: pid, deployment, opt, options
func (_m *DeploymentsServiceInterface) ApproveOrRejectProjectDeployment(pid interface{}, deployment int, opt *gitlab.ApproveOrRejectProjectDeploymentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, deployment, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ApproveOrRejectProjectDeployment")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.ApproveOrRejectProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, deployment, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.ApproveOrRejectProjectDeploymentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, deployment, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.ApproveOrRejectProjectDeploymentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, deployment, opt, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApproveOrRejectProjectDeployment'
type DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call struct {
	*mock.Call
}

// ApproveOrRejectProjectDeployment is a helper method to define mock.On call
//   - pid interface{}
//   - deployment int
//   - opt *gitlab.ApproveOrRejectProjectDeploymentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) ApproveOrRejectProjectDeployment(pid interface{}, deployment interface{}, opt interface{}, options ...interface{}) *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call {
	return &DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call{Call: _e.mock.On("ApproveOrRejectProjectDeployment",
		append([]interface{}{pid, deployment, opt}, options...)...)}
}

func (_c *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call) Run(run func(pid interface{}, deployment int, opt *gitlab.ApproveOrRejectProjectDeploymentOptions, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.ApproveOrRejectProjectDeploymentOptions), variadicArgs...)
	})
	return _c
}

func (_c *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call) Return(_a0 *gitlab.Response, _a1 error) *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call) RunAndReturn(run func(interface{}, int, *gitlab.ApproveOrRejectProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *DeploymentsServiceInterface_ApproveOrRejectProjectDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProjectDeployment provides a mock function with given fields: pid, opt, options
func (_m *DeploymentsServiceInterface) CreateProjectDeployment(pid interface{}, opt *gitlab.CreateProjectDeploymentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateProjectDeployment")
	}

	var r0 *gitlab.Deployment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) *gitlab.Deployment); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeploymentsServiceInterface_CreateProjectDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProjectDeployment'
type DeploymentsServiceInterface_CreateProjectDeployment_Call struct {
	*mock.Call
}

// CreateProjectDeployment is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateProjectDeploymentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) CreateProjectDeployment(pid interface{}, opt interface{}, options ...interface{}) *DeploymentsServiceInterface_CreateProjectDeployment_Call {
	return &DeploymentsServiceInterface_CreateProjectDeployment_Call{Call: _e.mock.On("CreateProjectDeployment",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *DeploymentsServiceInterface_CreateProjectDeployment_Call) Run(run func(pid interface{}, opt *gitlab.CreateProjectDeploymentOptions, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_CreateProjectDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateProjectDeploymentOptions), variadicArgs...)
	})
	return _c
}

func (_c *DeploymentsServiceInterface_CreateProjectDeployment_Call) Return(_a0 *gitlab.Deployment, _a1 *gitlab.Response, _a2 error) *DeploymentsServiceInterface_CreateProjectDeployment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DeploymentsServiceInterface_CreateProjectDeployment_Call) RunAndReturn(run func(interface{}, *gitlab.CreateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)) *DeploymentsServiceInterface_CreateProjectDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProjectDeployment provides a mock function with given fields: pid, deployment, options
func (_m *DeploymentsServiceInterface) DeleteProjectDeployment(pid interface{}, deployment int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, deployment)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProjectDeployment")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, deployment, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, deployment, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, deployment, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeploymentsServiceInterface_DeleteProjectDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProjectDeployment'
type DeploymentsServiceInterface_DeleteProjectDeployment_Call struct {
	*mock.Call
}

// DeleteProjectDeployment is a helper method to define mock.On call
//   - pid interface{}
//   - deployment int
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) DeleteProjectDeployment(pid interface{}, deployment interface{}, options ...interface{}) *DeploymentsServiceInterface_DeleteProjectDeployment_Call {
	return &DeploymentsServiceInterface_DeleteProjectDeployment_Call{Call: _e.mock.On("DeleteProjectDeployment",
		append([]interface{}{pid, deployment}, options...)...)}
}

func (_c *DeploymentsServiceInterface_DeleteProjectDeployment_Call) Run(run func(pid interface{}, deployment int, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_DeleteProjectDeployment_Call {
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

func (_c *DeploymentsServiceInterface_DeleteProjectDeployment_Call) Return(_a0 *gitlab.Response, _a1 error) *DeploymentsServiceInterface_DeleteProjectDeployment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeploymentsServiceInterface_DeleteProjectDeployment_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *DeploymentsServiceInterface_DeleteProjectDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// GetProjectDeployment provides a mock function with given fields: pid, deployment, options
func (_m *DeploymentsServiceInterface) GetProjectDeployment(pid interface{}, deployment int, options ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, deployment)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectDeployment")
	}

	var r0 *gitlab.Deployment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)); ok {
		return rf(pid, deployment, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Deployment); ok {
		r0 = rf(pid, deployment, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, deployment, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, deployment, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeploymentsServiceInterface_GetProjectDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProjectDeployment'
type DeploymentsServiceInterface_GetProjectDeployment_Call struct {
	*mock.Call
}

// GetProjectDeployment is a helper method to define mock.On call
//   - pid interface{}
//   - deployment int
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) GetProjectDeployment(pid interface{}, deployment interface{}, options ...interface{}) *DeploymentsServiceInterface_GetProjectDeployment_Call {
	return &DeploymentsServiceInterface_GetProjectDeployment_Call{Call: _e.mock.On("GetProjectDeployment",
		append([]interface{}{pid, deployment}, options...)...)}
}

func (_c *DeploymentsServiceInterface_GetProjectDeployment_Call) Run(run func(pid interface{}, deployment int, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_GetProjectDeployment_Call {
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

func (_c *DeploymentsServiceInterface_GetProjectDeployment_Call) Return(_a0 *gitlab.Deployment, _a1 *gitlab.Response, _a2 error) *DeploymentsServiceInterface_GetProjectDeployment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DeploymentsServiceInterface_GetProjectDeployment_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)) *DeploymentsServiceInterface_GetProjectDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// ListProjectDeployments provides a mock function with given fields: pid, opts, options
func (_m *DeploymentsServiceInterface) ListProjectDeployments(pid interface{}, opts *gitlab.ListProjectDeploymentsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Deployment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectDeployments")
	}

	var r0 []*gitlab.Deployment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectDeploymentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Deployment, *gitlab.Response, error)); ok {
		return rf(pid, opts, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectDeploymentsOptions, ...gitlab.RequestOptionFunc) []*gitlab.Deployment); ok {
		r0 = rf(pid, opts, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectDeploymentsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opts, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectDeploymentsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opts, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeploymentsServiceInterface_ListProjectDeployments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListProjectDeployments'
type DeploymentsServiceInterface_ListProjectDeployments_Call struct {
	*mock.Call
}

// ListProjectDeployments is a helper method to define mock.On call
//   - pid interface{}
//   - opts *gitlab.ListProjectDeploymentsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) ListProjectDeployments(pid interface{}, opts interface{}, options ...interface{}) *DeploymentsServiceInterface_ListProjectDeployments_Call {
	return &DeploymentsServiceInterface_ListProjectDeployments_Call{Call: _e.mock.On("ListProjectDeployments",
		append([]interface{}{pid, opts}, options...)...)}
}

func (_c *DeploymentsServiceInterface_ListProjectDeployments_Call) Run(run func(pid interface{}, opts *gitlab.ListProjectDeploymentsOptions, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_ListProjectDeployments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListProjectDeploymentsOptions), variadicArgs...)
	})
	return _c
}

func (_c *DeploymentsServiceInterface_ListProjectDeployments_Call) Return(_a0 []*gitlab.Deployment, _a1 *gitlab.Response, _a2 error) *DeploymentsServiceInterface_ListProjectDeployments_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DeploymentsServiceInterface_ListProjectDeployments_Call) RunAndReturn(run func(interface{}, *gitlab.ListProjectDeploymentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Deployment, *gitlab.Response, error)) *DeploymentsServiceInterface_ListProjectDeployments_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProjectDeployment provides a mock function with given fields: pid, deployment, opt, options
func (_m *DeploymentsServiceInterface) UpdateProjectDeployment(pid interface{}, deployment int, opt *gitlab.UpdateProjectDeploymentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, deployment, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProjectDeployment")
	}

	var r0 *gitlab.Deployment
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)); ok {
		return rf(pid, deployment, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) *gitlab.Deployment); ok {
		r0 = rf(pid, deployment, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Deployment)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.UpdateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, deployment, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.UpdateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, deployment, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DeploymentsServiceInterface_UpdateProjectDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProjectDeployment'
type DeploymentsServiceInterface_UpdateProjectDeployment_Call struct {
	*mock.Call
}

// UpdateProjectDeployment is a helper method to define mock.On call
//   - pid interface{}
//   - deployment int
//   - opt *gitlab.UpdateProjectDeploymentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *DeploymentsServiceInterface_Expecter) UpdateProjectDeployment(pid interface{}, deployment interface{}, opt interface{}, options ...interface{}) *DeploymentsServiceInterface_UpdateProjectDeployment_Call {
	return &DeploymentsServiceInterface_UpdateProjectDeployment_Call{Call: _e.mock.On("UpdateProjectDeployment",
		append([]interface{}{pid, deployment, opt}, options...)...)}
}

func (_c *DeploymentsServiceInterface_UpdateProjectDeployment_Call) Run(run func(pid interface{}, deployment int, opt *gitlab.UpdateProjectDeploymentOptions, options ...gitlab.RequestOptionFunc)) *DeploymentsServiceInterface_UpdateProjectDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.UpdateProjectDeploymentOptions), variadicArgs...)
	})
	return _c
}

func (_c *DeploymentsServiceInterface_UpdateProjectDeployment_Call) Return(_a0 *gitlab.Deployment, _a1 *gitlab.Response, _a2 error) *DeploymentsServiceInterface_UpdateProjectDeployment_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *DeploymentsServiceInterface_UpdateProjectDeployment_Call) RunAndReturn(run func(interface{}, int, *gitlab.UpdateProjectDeploymentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Deployment, *gitlab.Response, error)) *DeploymentsServiceInterface_UpdateProjectDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeploymentsServiceInterface creates a new instance of DeploymentsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeploymentsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeploymentsServiceInterface {
	mock := &DeploymentsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

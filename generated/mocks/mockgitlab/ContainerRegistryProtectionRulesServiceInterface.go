// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ContainerRegistryProtectionRulesServiceInterface is an autogenerated mock type for the ContainerRegistryProtectionRulesServiceInterface type
type ContainerRegistryProtectionRulesServiceInterface struct {
	mock.Mock
}

type ContainerRegistryProtectionRulesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ContainerRegistryProtectionRulesServiceInterface) EXPECT() *ContainerRegistryProtectionRulesServiceInterface_Expecter {
	return &ContainerRegistryProtectionRulesServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateContainerRegistryProtectionRule provides a mock function with given fields: pid, opt, options
func (_m *ContainerRegistryProtectionRulesServiceInterface) CreateContainerRegistryProtectionRule(pid interface{}, opt *gitlab.CreateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateContainerRegistryProtectionRule")
	}

	var r0 *gitlab.ContainerRegistryProtectionRule
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) *gitlab.ContainerRegistryProtectionRule); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ContainerRegistryProtectionRule)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateContainerRegistryProtectionRule'
type ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call struct {
	*mock.Call
}

// CreateContainerRegistryProtectionRule is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CreateContainerRegistryProtectionRuleOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ContainerRegistryProtectionRulesServiceInterface_Expecter) CreateContainerRegistryProtectionRule(pid interface{}, opt interface{}, options ...interface{}) *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call {
	return &ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call{Call: _e.mock.On("CreateContainerRegistryProtectionRule",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call) Run(run func(pid interface{}, opt *gitlab.CreateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc)) *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateContainerRegistryProtectionRuleOptions), variadicArgs...)
	})
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call) Return(_a0 *gitlab.ContainerRegistryProtectionRule, _a1 *gitlab.Response, _a2 error) *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call) RunAndReturn(run func(interface{}, *gitlab.CreateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *ContainerRegistryProtectionRulesServiceInterface_CreateContainerRegistryProtectionRule_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteContainerRegistryProtectionRule provides a mock function with given fields: pid, ruleID, options
func (_m *ContainerRegistryProtectionRulesServiceInterface) DeleteContainerRegistryProtectionRule(pid interface{}, ruleID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, ruleID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteContainerRegistryProtectionRule")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, ruleID, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, ruleID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, ruleID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteContainerRegistryProtectionRule'
type ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call struct {
	*mock.Call
}

// DeleteContainerRegistryProtectionRule is a helper method to define mock.On call
//   - pid interface{}
//   - ruleID int
//   - options ...gitlab.RequestOptionFunc
func (_e *ContainerRegistryProtectionRulesServiceInterface_Expecter) DeleteContainerRegistryProtectionRule(pid interface{}, ruleID interface{}, options ...interface{}) *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call {
	return &ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call{Call: _e.mock.On("DeleteContainerRegistryProtectionRule",
		append([]interface{}{pid, ruleID}, options...)...)}
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call) Run(run func(pid interface{}, ruleID int, options ...gitlab.RequestOptionFunc)) *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call {
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

func (_c *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call) Return(_a0 *gitlab.Response, _a1 error) *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ContainerRegistryProtectionRulesServiceInterface_DeleteContainerRegistryProtectionRule_Call {
	_c.Call.Return(run)
	return _c
}

// ListContainerRegistryProtectionRules provides a mock function with given fields: pid, options
func (_m *ContainerRegistryProtectionRulesServiceInterface) ListContainerRegistryProtectionRules(pid interface{}, options ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListContainerRegistryProtectionRules")
	}

	var r0 []*gitlab.ContainerRegistryProtectionRule
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)); ok {
		return rf(pid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) []*gitlab.ContainerRegistryProtectionRule); ok {
		r0 = rf(pid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ContainerRegistryProtectionRule)
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

// ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListContainerRegistryProtectionRules'
type ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call struct {
	*mock.Call
}

// ListContainerRegistryProtectionRules is a helper method to define mock.On call
//   - pid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *ContainerRegistryProtectionRulesServiceInterface_Expecter) ListContainerRegistryProtectionRules(pid interface{}, options ...interface{}) *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call {
	return &ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call{Call: _e.mock.On("ListContainerRegistryProtectionRules",
		append([]interface{}{pid}, options...)...)}
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call) Run(run func(pid interface{}, options ...gitlab.RequestOptionFunc)) *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call {
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

func (_c *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call) Return(_a0 []*gitlab.ContainerRegistryProtectionRule, _a1 *gitlab.Response, _a2 error) *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *ContainerRegistryProtectionRulesServiceInterface_ListContainerRegistryProtectionRules_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateContainerRegistryProtectionRule provides a mock function with given fields: pid, ruleID, opt, options
func (_m *ContainerRegistryProtectionRulesServiceInterface) UpdateContainerRegistryProtectionRule(pid interface{}, ruleID int, opt *gitlab.UpdateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, ruleID, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateContainerRegistryProtectionRule")
	}

	var r0 *gitlab.ContainerRegistryProtectionRule
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)); ok {
		return rf(pid, ruleID, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) *gitlab.ContainerRegistryProtectionRule); ok {
		r0 = rf(pid, ruleID, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ContainerRegistryProtectionRule)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, ruleID, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, ruleID, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateContainerRegistryProtectionRule'
type ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call struct {
	*mock.Call
}

// UpdateContainerRegistryProtectionRule is a helper method to define mock.On call
//   - pid interface{}
//   - ruleID int
//   - opt *gitlab.UpdateContainerRegistryProtectionRuleOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ContainerRegistryProtectionRulesServiceInterface_Expecter) UpdateContainerRegistryProtectionRule(pid interface{}, ruleID interface{}, opt interface{}, options ...interface{}) *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call {
	return &ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call{Call: _e.mock.On("UpdateContainerRegistryProtectionRule",
		append([]interface{}{pid, ruleID, opt}, options...)...)}
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call) Run(run func(pid interface{}, ruleID int, opt *gitlab.UpdateContainerRegistryProtectionRuleOptions, options ...gitlab.RequestOptionFunc)) *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.UpdateContainerRegistryProtectionRuleOptions), variadicArgs...)
	})
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call) Return(_a0 *gitlab.ContainerRegistryProtectionRule, _a1 *gitlab.Response, _a2 error) *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call) RunAndReturn(run func(interface{}, int, *gitlab.UpdateContainerRegistryProtectionRuleOptions, ...gitlab.RequestOptionFunc) (*gitlab.ContainerRegistryProtectionRule, *gitlab.Response, error)) *ContainerRegistryProtectionRulesServiceInterface_UpdateContainerRegistryProtectionRule_Call {
	_c.Call.Return(run)
	return _c
}

// NewContainerRegistryProtectionRulesServiceInterface creates a new instance of ContainerRegistryProtectionRulesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContainerRegistryProtectionRulesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContainerRegistryProtectionRulesServiceInterface {
	mock := &ContainerRegistryProtectionRulesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

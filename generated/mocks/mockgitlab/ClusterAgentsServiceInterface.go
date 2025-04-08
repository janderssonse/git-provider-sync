// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// ClusterAgentsServiceInterface is an autogenerated mock type for the ClusterAgentsServiceInterface type
type ClusterAgentsServiceInterface struct {
	mock.Mock
}

type ClusterAgentsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ClusterAgentsServiceInterface) EXPECT() *ClusterAgentsServiceInterface_Expecter {
	return &ClusterAgentsServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateAgentToken provides a mock function with given fields: pid, aid, opt, options
func (_m *ClusterAgentsServiceInterface) CreateAgentToken(pid interface{}, aid int, opt *gitlab.CreateAgentTokenOptions, options ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, aid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateAgentToken")
	}

	var r0 *gitlab.AgentToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.CreateAgentTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error)); ok {
		return rf(pid, aid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.CreateAgentTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.AgentToken); ok {
		r0 = rf(pid, aid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.AgentToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.CreateAgentTokenOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, aid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.CreateAgentTokenOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, aid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterAgentsServiceInterface_CreateAgentToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAgentToken'
type ClusterAgentsServiceInterface_CreateAgentToken_Call struct {
	*mock.Call
}

// CreateAgentToken is a helper method to define mock.On call
//   - pid interface{}
//   - aid int
//   - opt *gitlab.CreateAgentTokenOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) CreateAgentToken(pid interface{}, aid interface{}, opt interface{}, options ...interface{}) *ClusterAgentsServiceInterface_CreateAgentToken_Call {
	return &ClusterAgentsServiceInterface_CreateAgentToken_Call{Call: _e.mock.On("CreateAgentToken",
		append([]interface{}{pid, aid, opt}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_CreateAgentToken_Call) Run(run func(pid interface{}, aid int, opt *gitlab.CreateAgentTokenOptions, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_CreateAgentToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.CreateAgentTokenOptions), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_CreateAgentToken_Call) Return(_a0 *gitlab.AgentToken, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_CreateAgentToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_CreateAgentToken_Call) RunAndReturn(run func(interface{}, int, *gitlab.CreateAgentTokenOptions, ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error)) *ClusterAgentsServiceInterface_CreateAgentToken_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAgent provides a mock function with given fields: pid, id, options
func (_m *ClusterAgentsServiceInterface) DeleteAgent(pid interface{}, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAgent")
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

// ClusterAgentsServiceInterface_DeleteAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAgent'
type ClusterAgentsServiceInterface_DeleteAgent_Call struct {
	*mock.Call
}

// DeleteAgent is a helper method to define mock.On call
//   - pid interface{}
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) DeleteAgent(pid interface{}, id interface{}, options ...interface{}) *ClusterAgentsServiceInterface_DeleteAgent_Call {
	return &ClusterAgentsServiceInterface_DeleteAgent_Call{Call: _e.mock.On("DeleteAgent",
		append([]interface{}{pid, id}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_DeleteAgent_Call) Run(run func(pid interface{}, id int, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_DeleteAgent_Call {
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

func (_c *ClusterAgentsServiceInterface_DeleteAgent_Call) Return(_a0 *gitlab.Response, _a1 error) *ClusterAgentsServiceInterface_DeleteAgent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClusterAgentsServiceInterface_DeleteAgent_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ClusterAgentsServiceInterface_DeleteAgent_Call {
	_c.Call.Return(run)
	return _c
}

// GetAgent provides a mock function with given fields: pid, id, options
func (_m *ClusterAgentsServiceInterface) GetAgent(pid interface{}, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAgent")
	}

	var r0 *gitlab.Agent
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error)); ok {
		return rf(pid, id, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Agent); ok {
		r0 = rf(pid, id, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Agent)
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

// ClusterAgentsServiceInterface_GetAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAgent'
type ClusterAgentsServiceInterface_GetAgent_Call struct {
	*mock.Call
}

// GetAgent is a helper method to define mock.On call
//   - pid interface{}
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) GetAgent(pid interface{}, id interface{}, options ...interface{}) *ClusterAgentsServiceInterface_GetAgent_Call {
	return &ClusterAgentsServiceInterface_GetAgent_Call{Call: _e.mock.On("GetAgent",
		append([]interface{}{pid, id}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_GetAgent_Call) Run(run func(pid interface{}, id int, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_GetAgent_Call {
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

func (_c *ClusterAgentsServiceInterface_GetAgent_Call) Return(_a0 *gitlab.Agent, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_GetAgent_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_GetAgent_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error)) *ClusterAgentsServiceInterface_GetAgent_Call {
	_c.Call.Return(run)
	return _c
}

// GetAgentToken provides a mock function with given fields: pid, aid, id, options
func (_m *ClusterAgentsServiceInterface) GetAgentToken(pid interface{}, aid int, id int, options ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, aid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAgentToken")
	}

	var r0 *gitlab.AgentToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error)); ok {
		return rf(pid, aid, id, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) *gitlab.AgentToken); ok {
		r0 = rf(pid, aid, id, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.AgentToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, aid, id, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, aid, id, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterAgentsServiceInterface_GetAgentToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAgentToken'
type ClusterAgentsServiceInterface_GetAgentToken_Call struct {
	*mock.Call
}

// GetAgentToken is a helper method to define mock.On call
//   - pid interface{}
//   - aid int
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) GetAgentToken(pid interface{}, aid interface{}, id interface{}, options ...interface{}) *ClusterAgentsServiceInterface_GetAgentToken_Call {
	return &ClusterAgentsServiceInterface_GetAgentToken_Call{Call: _e.mock.On("GetAgentToken",
		append([]interface{}{pid, aid, id}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_GetAgentToken_Call) Run(run func(pid interface{}, aid int, id int, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_GetAgentToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(int), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_GetAgentToken_Call) Return(_a0 *gitlab.AgentToken, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_GetAgentToken_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_GetAgentToken_Call) RunAndReturn(run func(interface{}, int, int, ...gitlab.RequestOptionFunc) (*gitlab.AgentToken, *gitlab.Response, error)) *ClusterAgentsServiceInterface_GetAgentToken_Call {
	_c.Call.Return(run)
	return _c
}

// ListAgentTokens provides a mock function with given fields: pid, aid, opt, options
func (_m *ClusterAgentsServiceInterface) ListAgentTokens(pid interface{}, aid int, opt *gitlab.ListAgentTokensOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.AgentToken, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, aid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAgentTokens")
	}

	var r0 []*gitlab.AgentToken
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.ListAgentTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AgentToken, *gitlab.Response, error)); ok {
		return rf(pid, aid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.ListAgentTokensOptions, ...gitlab.RequestOptionFunc) []*gitlab.AgentToken); ok {
		r0 = rf(pid, aid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.AgentToken)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.ListAgentTokensOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, aid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.ListAgentTokensOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, aid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterAgentsServiceInterface_ListAgentTokens_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAgentTokens'
type ClusterAgentsServiceInterface_ListAgentTokens_Call struct {
	*mock.Call
}

// ListAgentTokens is a helper method to define mock.On call
//   - pid interface{}
//   - aid int
//   - opt *gitlab.ListAgentTokensOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) ListAgentTokens(pid interface{}, aid interface{}, opt interface{}, options ...interface{}) *ClusterAgentsServiceInterface_ListAgentTokens_Call {
	return &ClusterAgentsServiceInterface_ListAgentTokens_Call{Call: _e.mock.On("ListAgentTokens",
		append([]interface{}{pid, aid, opt}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_ListAgentTokens_Call) Run(run func(pid interface{}, aid int, opt *gitlab.ListAgentTokensOptions, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_ListAgentTokens_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.ListAgentTokensOptions), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_ListAgentTokens_Call) Return(_a0 []*gitlab.AgentToken, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_ListAgentTokens_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_ListAgentTokens_Call) RunAndReturn(run func(interface{}, int, *gitlab.ListAgentTokensOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.AgentToken, *gitlab.Response, error)) *ClusterAgentsServiceInterface_ListAgentTokens_Call {
	_c.Call.Return(run)
	return _c
}

// ListAgents provides a mock function with given fields: pid, opt, options
func (_m *ClusterAgentsServiceInterface) ListAgents(pid interface{}, opt *gitlab.ListAgentsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Agent, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListAgents")
	}

	var r0 []*gitlab.Agent
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListAgentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Agent, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListAgentsOptions, ...gitlab.RequestOptionFunc) []*gitlab.Agent); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Agent)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListAgentsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListAgentsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterAgentsServiceInterface_ListAgents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAgents'
type ClusterAgentsServiceInterface_ListAgents_Call struct {
	*mock.Call
}

// ListAgents is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListAgentsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) ListAgents(pid interface{}, opt interface{}, options ...interface{}) *ClusterAgentsServiceInterface_ListAgents_Call {
	return &ClusterAgentsServiceInterface_ListAgents_Call{Call: _e.mock.On("ListAgents",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_ListAgents_Call) Run(run func(pid interface{}, opt *gitlab.ListAgentsOptions, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_ListAgents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListAgentsOptions), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_ListAgents_Call) Return(_a0 []*gitlab.Agent, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_ListAgents_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_ListAgents_Call) RunAndReturn(run func(interface{}, *gitlab.ListAgentsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Agent, *gitlab.Response, error)) *ClusterAgentsServiceInterface_ListAgents_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterAgent provides a mock function with given fields: pid, opt, options
func (_m *ClusterAgentsServiceInterface) RegisterAgent(pid interface{}, opt *gitlab.RegisterAgentOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RegisterAgent")
	}

	var r0 *gitlab.Agent
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.RegisterAgentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.RegisterAgentOptions, ...gitlab.RequestOptionFunc) *gitlab.Agent); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Agent)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.RegisterAgentOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.RegisterAgentOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClusterAgentsServiceInterface_RegisterAgent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterAgent'
type ClusterAgentsServiceInterface_RegisterAgent_Call struct {
	*mock.Call
}

// RegisterAgent is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.RegisterAgentOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) RegisterAgent(pid interface{}, opt interface{}, options ...interface{}) *ClusterAgentsServiceInterface_RegisterAgent_Call {
	return &ClusterAgentsServiceInterface_RegisterAgent_Call{Call: _e.mock.On("RegisterAgent",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_RegisterAgent_Call) Run(run func(pid interface{}, opt *gitlab.RegisterAgentOptions, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_RegisterAgent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.RegisterAgentOptions), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_RegisterAgent_Call) Return(_a0 *gitlab.Agent, _a1 *gitlab.Response, _a2 error) *ClusterAgentsServiceInterface_RegisterAgent_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ClusterAgentsServiceInterface_RegisterAgent_Call) RunAndReturn(run func(interface{}, *gitlab.RegisterAgentOptions, ...gitlab.RequestOptionFunc) (*gitlab.Agent, *gitlab.Response, error)) *ClusterAgentsServiceInterface_RegisterAgent_Call {
	_c.Call.Return(run)
	return _c
}

// RevokeAgentToken provides a mock function with given fields: pid, aid, id, options
func (_m *ClusterAgentsServiceInterface) RevokeAgentToken(pid interface{}, aid int, id int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, aid, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RevokeAgentToken")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, aid, id, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, aid, id, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, aid, id, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClusterAgentsServiceInterface_RevokeAgentToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokeAgentToken'
type ClusterAgentsServiceInterface_RevokeAgentToken_Call struct {
	*mock.Call
}

// RevokeAgentToken is a helper method to define mock.On call
//   - pid interface{}
//   - aid int
//   - id int
//   - options ...gitlab.RequestOptionFunc
func (_e *ClusterAgentsServiceInterface_Expecter) RevokeAgentToken(pid interface{}, aid interface{}, id interface{}, options ...interface{}) *ClusterAgentsServiceInterface_RevokeAgentToken_Call {
	return &ClusterAgentsServiceInterface_RevokeAgentToken_Call{Call: _e.mock.On("RevokeAgentToken",
		append([]interface{}{pid, aid, id}, options...)...)}
}

func (_c *ClusterAgentsServiceInterface_RevokeAgentToken_Call) Run(run func(pid interface{}, aid int, id int, options ...gitlab.RequestOptionFunc)) *ClusterAgentsServiceInterface_RevokeAgentToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(int), variadicArgs...)
	})
	return _c
}

func (_c *ClusterAgentsServiceInterface_RevokeAgentToken_Call) Return(_a0 *gitlab.Response, _a1 error) *ClusterAgentsServiceInterface_RevokeAgentToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClusterAgentsServiceInterface_RevokeAgentToken_Call) RunAndReturn(run func(interface{}, int, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *ClusterAgentsServiceInterface_RevokeAgentToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewClusterAgentsServiceInterface creates a new instance of ClusterAgentsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClusterAgentsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClusterAgentsServiceInterface {
	mock := &ClusterAgentsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

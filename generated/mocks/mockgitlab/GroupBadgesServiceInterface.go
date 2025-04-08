// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GroupBadgesServiceInterface is an autogenerated mock type for the GroupBadgesServiceInterface type
type GroupBadgesServiceInterface struct {
	mock.Mock
}

type GroupBadgesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupBadgesServiceInterface) EXPECT() *GroupBadgesServiceInterface_Expecter {
	return &GroupBadgesServiceInterface_Expecter{mock: &_m.Mock}
}

// AddGroupBadge provides a mock function with given fields: gid, opt, options
func (_m *GroupBadgesServiceInterface) AddGroupBadge(gid interface{}, opt *gitlab.AddGroupBadgeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddGroupBadge")
	}

	var r0 *gitlab.GroupBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddGroupBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddGroupBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupBadge); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.AddGroupBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.AddGroupBadgeOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupBadgesServiceInterface_AddGroupBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGroupBadge'
type GroupBadgesServiceInterface_AddGroupBadge_Call struct {
	*mock.Call
}

// AddGroupBadge is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.AddGroupBadgeOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) AddGroupBadge(gid interface{}, opt interface{}, options ...interface{}) *GroupBadgesServiceInterface_AddGroupBadge_Call {
	return &GroupBadgesServiceInterface_AddGroupBadge_Call{Call: _e.mock.On("AddGroupBadge",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_AddGroupBadge_Call) Run(run func(gid interface{}, opt *gitlab.AddGroupBadgeOptions, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_AddGroupBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.AddGroupBadgeOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupBadgesServiceInterface_AddGroupBadge_Call) Return(_a0 *gitlab.GroupBadge, _a1 *gitlab.Response, _a2 error) *GroupBadgesServiceInterface_AddGroupBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupBadgesServiceInterface_AddGroupBadge_Call) RunAndReturn(run func(interface{}, *gitlab.AddGroupBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)) *GroupBadgesServiceInterface_AddGroupBadge_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteGroupBadge provides a mock function with given fields: gid, badge, options
func (_m *GroupBadgesServiceInterface) DeleteGroupBadge(gid interface{}, badge int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, badge)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteGroupBadge")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(gid, badge, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(gid, badge, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(gid, badge, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupBadgesServiceInterface_DeleteGroupBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteGroupBadge'
type GroupBadgesServiceInterface_DeleteGroupBadge_Call struct {
	*mock.Call
}

// DeleteGroupBadge is a helper method to define mock.On call
//   - gid interface{}
//   - badge int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) DeleteGroupBadge(gid interface{}, badge interface{}, options ...interface{}) *GroupBadgesServiceInterface_DeleteGroupBadge_Call {
	return &GroupBadgesServiceInterface_DeleteGroupBadge_Call{Call: _e.mock.On("DeleteGroupBadge",
		append([]interface{}{gid, badge}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_DeleteGroupBadge_Call) Run(run func(gid interface{}, badge int, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_DeleteGroupBadge_Call {
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

func (_c *GroupBadgesServiceInterface_DeleteGroupBadge_Call) Return(_a0 *gitlab.Response, _a1 error) *GroupBadgesServiceInterface_DeleteGroupBadge_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GroupBadgesServiceInterface_DeleteGroupBadge_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *GroupBadgesServiceInterface_DeleteGroupBadge_Call {
	_c.Call.Return(run)
	return _c
}

// EditGroupBadge provides a mock function with given fields: gid, badge, opt, options
func (_m *GroupBadgesServiceInterface) EditGroupBadge(gid interface{}, badge int, opt *gitlab.EditGroupBadgeOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, badge, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EditGroupBadge")
	}

	var r0 *gitlab.GroupBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditGroupBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)); ok {
		return rf(gid, badge, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditGroupBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupBadge); ok {
		r0 = rf(gid, badge, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.EditGroupBadgeOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, badge, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.EditGroupBadgeOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, badge, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupBadgesServiceInterface_EditGroupBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditGroupBadge'
type GroupBadgesServiceInterface_EditGroupBadge_Call struct {
	*mock.Call
}

// EditGroupBadge is a helper method to define mock.On call
//   - gid interface{}
//   - badge int
//   - opt *gitlab.EditGroupBadgeOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) EditGroupBadge(gid interface{}, badge interface{}, opt interface{}, options ...interface{}) *GroupBadgesServiceInterface_EditGroupBadge_Call {
	return &GroupBadgesServiceInterface_EditGroupBadge_Call{Call: _e.mock.On("EditGroupBadge",
		append([]interface{}{gid, badge, opt}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_EditGroupBadge_Call) Run(run func(gid interface{}, badge int, opt *gitlab.EditGroupBadgeOptions, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_EditGroupBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.EditGroupBadgeOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupBadgesServiceInterface_EditGroupBadge_Call) Return(_a0 *gitlab.GroupBadge, _a1 *gitlab.Response, _a2 error) *GroupBadgesServiceInterface_EditGroupBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupBadgesServiceInterface_EditGroupBadge_Call) RunAndReturn(run func(interface{}, int, *gitlab.EditGroupBadgeOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)) *GroupBadgesServiceInterface_EditGroupBadge_Call {
	_c.Call.Return(run)
	return _c
}

// GetGroupBadge provides a mock function with given fields: gid, badge, options
func (_m *GroupBadgesServiceInterface) GetGroupBadge(gid interface{}, badge int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, badge)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupBadge")
	}

	var r0 *gitlab.GroupBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)); ok {
		return rf(gid, badge, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.GroupBadge); ok {
		r0 = rf(gid, badge, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, badge, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, badge, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupBadgesServiceInterface_GetGroupBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupBadge'
type GroupBadgesServiceInterface_GetGroupBadge_Call struct {
	*mock.Call
}

// GetGroupBadge is a helper method to define mock.On call
//   - gid interface{}
//   - badge int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) GetGroupBadge(gid interface{}, badge interface{}, options ...interface{}) *GroupBadgesServiceInterface_GetGroupBadge_Call {
	return &GroupBadgesServiceInterface_GetGroupBadge_Call{Call: _e.mock.On("GetGroupBadge",
		append([]interface{}{gid, badge}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_GetGroupBadge_Call) Run(run func(gid interface{}, badge int, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_GetGroupBadge_Call {
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

func (_c *GroupBadgesServiceInterface_GetGroupBadge_Call) Return(_a0 *gitlab.GroupBadge, _a1 *gitlab.Response, _a2 error) *GroupBadgesServiceInterface_GetGroupBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupBadgesServiceInterface_GetGroupBadge_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)) *GroupBadgesServiceInterface_GetGroupBadge_Call {
	_c.Call.Return(run)
	return _c
}

// ListGroupBadges provides a mock function with given fields: gid, opt, options
func (_m *GroupBadgesServiceInterface) ListGroupBadges(gid interface{}, opt *gitlab.ListGroupBadgesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupBadges")
	}

	var r0 []*gitlab.GroupBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListGroupBadgesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupBadge, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListGroupBadgesOptions, ...gitlab.RequestOptionFunc) []*gitlab.GroupBadge); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.GroupBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListGroupBadgesOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListGroupBadgesOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupBadgesServiceInterface_ListGroupBadges_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGroupBadges'
type GroupBadgesServiceInterface_ListGroupBadges_Call struct {
	*mock.Call
}

// ListGroupBadges is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.ListGroupBadgesOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) ListGroupBadges(gid interface{}, opt interface{}, options ...interface{}) *GroupBadgesServiceInterface_ListGroupBadges_Call {
	return &GroupBadgesServiceInterface_ListGroupBadges_Call{Call: _e.mock.On("ListGroupBadges",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_ListGroupBadges_Call) Run(run func(gid interface{}, opt *gitlab.ListGroupBadgesOptions, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_ListGroupBadges_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListGroupBadgesOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupBadgesServiceInterface_ListGroupBadges_Call) Return(_a0 []*gitlab.GroupBadge, _a1 *gitlab.Response, _a2 error) *GroupBadgesServiceInterface_ListGroupBadges_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupBadgesServiceInterface_ListGroupBadges_Call) RunAndReturn(run func(interface{}, *gitlab.ListGroupBadgesOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupBadge, *gitlab.Response, error)) *GroupBadgesServiceInterface_ListGroupBadges_Call {
	_c.Call.Return(run)
	return _c
}

// PreviewGroupBadge provides a mock function with given fields: gid, opt, options
func (_m *GroupBadgesServiceInterface) PreviewGroupBadge(gid interface{}, opt *gitlab.GroupBadgePreviewOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PreviewGroupBadge")
	}

	var r0 *gitlab.GroupBadge
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.GroupBadgePreviewOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.GroupBadgePreviewOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupBadge); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupBadge)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.GroupBadgePreviewOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.GroupBadgePreviewOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupBadgesServiceInterface_PreviewGroupBadge_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PreviewGroupBadge'
type GroupBadgesServiceInterface_PreviewGroupBadge_Call struct {
	*mock.Call
}

// PreviewGroupBadge is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.GroupBadgePreviewOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupBadgesServiceInterface_Expecter) PreviewGroupBadge(gid interface{}, opt interface{}, options ...interface{}) *GroupBadgesServiceInterface_PreviewGroupBadge_Call {
	return &GroupBadgesServiceInterface_PreviewGroupBadge_Call{Call: _e.mock.On("PreviewGroupBadge",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupBadgesServiceInterface_PreviewGroupBadge_Call) Run(run func(gid interface{}, opt *gitlab.GroupBadgePreviewOptions, options ...gitlab.RequestOptionFunc)) *GroupBadgesServiceInterface_PreviewGroupBadge_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.GroupBadgePreviewOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupBadgesServiceInterface_PreviewGroupBadge_Call) Return(_a0 *gitlab.GroupBadge, _a1 *gitlab.Response, _a2 error) *GroupBadgesServiceInterface_PreviewGroupBadge_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupBadgesServiceInterface_PreviewGroupBadge_Call) RunAndReturn(run func(interface{}, *gitlab.GroupBadgePreviewOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupBadge, *gitlab.Response, error)) *GroupBadgesServiceInterface_PreviewGroupBadge_Call {
	_c.Call.Return(run)
	return _c
}

// NewGroupBadgesServiceInterface creates a new instance of GroupBadgesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupBadgesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupBadgesServiceInterface {
	mock := &GroupBadgesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

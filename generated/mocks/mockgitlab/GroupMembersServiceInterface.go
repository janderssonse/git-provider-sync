// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GroupMembersServiceInterface is an autogenerated mock type for the GroupMembersServiceInterface type
type GroupMembersServiceInterface struct {
	mock.Mock
}

type GroupMembersServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupMembersServiceInterface) EXPECT() *GroupMembersServiceInterface_Expecter {
	return &GroupMembersServiceInterface_Expecter{mock: &_m.Mock}
}

// AddGroupMember provides a mock function with given fields: gid, opt, options
func (_m *GroupMembersServiceInterface) AddGroupMember(gid interface{}, opt *gitlab.AddGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddGroupMember")
	}

	var r0 *gitlab.GroupMember
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddGroupMemberOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupMember); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupMember)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.AddGroupMemberOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.AddGroupMemberOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupMembersServiceInterface_AddGroupMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddGroupMember'
type GroupMembersServiceInterface_AddGroupMember_Call struct {
	*mock.Call
}

// AddGroupMember is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.AddGroupMemberOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) AddGroupMember(gid interface{}, opt interface{}, options ...interface{}) *GroupMembersServiceInterface_AddGroupMember_Call {
	return &GroupMembersServiceInterface_AddGroupMember_Call{Call: _e.mock.On("AddGroupMember",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupMembersServiceInterface_AddGroupMember_Call) Run(run func(gid interface{}, opt *gitlab.AddGroupMemberOptions, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_AddGroupMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.AddGroupMemberOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupMembersServiceInterface_AddGroupMember_Call) Return(_a0 *gitlab.GroupMember, _a1 *gitlab.Response, _a2 error) *GroupMembersServiceInterface_AddGroupMember_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupMembersServiceInterface_AddGroupMember_Call) RunAndReturn(run func(interface{}, *gitlab.AddGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)) *GroupMembersServiceInterface_AddGroupMember_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteShareWithGroup provides a mock function with given fields: gid, groupID, options
func (_m *GroupMembersServiceInterface) DeleteShareWithGroup(gid interface{}, groupID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteShareWithGroup")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(gid, groupID, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(gid, groupID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(gid, groupID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupMembersServiceInterface_DeleteShareWithGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteShareWithGroup'
type GroupMembersServiceInterface_DeleteShareWithGroup_Call struct {
	*mock.Call
}

// DeleteShareWithGroup is a helper method to define mock.On call
//   - gid interface{}
//   - groupID int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) DeleteShareWithGroup(gid interface{}, groupID interface{}, options ...interface{}) *GroupMembersServiceInterface_DeleteShareWithGroup_Call {
	return &GroupMembersServiceInterface_DeleteShareWithGroup_Call{Call: _e.mock.On("DeleteShareWithGroup",
		append([]interface{}{gid, groupID}, options...)...)}
}

func (_c *GroupMembersServiceInterface_DeleteShareWithGroup_Call) Run(run func(gid interface{}, groupID int, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_DeleteShareWithGroup_Call {
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

func (_c *GroupMembersServiceInterface_DeleteShareWithGroup_Call) Return(_a0 *gitlab.Response, _a1 error) *GroupMembersServiceInterface_DeleteShareWithGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GroupMembersServiceInterface_DeleteShareWithGroup_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *GroupMembersServiceInterface_DeleteShareWithGroup_Call {
	_c.Call.Return(run)
	return _c
}

// EditGroupMember provides a mock function with given fields: gid, user, opt, options
func (_m *GroupMembersServiceInterface) EditGroupMember(gid interface{}, user int, opt *gitlab.EditGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, user, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EditGroupMember")
	}

	var r0 *gitlab.GroupMember
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)); ok {
		return rf(gid, user, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.EditGroupMemberOptions, ...gitlab.RequestOptionFunc) *gitlab.GroupMember); ok {
		r0 = rf(gid, user, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupMember)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.EditGroupMemberOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, user, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, *gitlab.EditGroupMemberOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, user, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupMembersServiceInterface_EditGroupMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditGroupMember'
type GroupMembersServiceInterface_EditGroupMember_Call struct {
	*mock.Call
}

// EditGroupMember is a helper method to define mock.On call
//   - gid interface{}
//   - user int
//   - opt *gitlab.EditGroupMemberOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) EditGroupMember(gid interface{}, user interface{}, opt interface{}, options ...interface{}) *GroupMembersServiceInterface_EditGroupMember_Call {
	return &GroupMembersServiceInterface_EditGroupMember_Call{Call: _e.mock.On("EditGroupMember",
		append([]interface{}{gid, user, opt}, options...)...)}
}

func (_c *GroupMembersServiceInterface_EditGroupMember_Call) Run(run func(gid interface{}, user int, opt *gitlab.EditGroupMemberOptions, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_EditGroupMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.EditGroupMemberOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupMembersServiceInterface_EditGroupMember_Call) Return(_a0 *gitlab.GroupMember, _a1 *gitlab.Response, _a2 error) *GroupMembersServiceInterface_EditGroupMember_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupMembersServiceInterface_EditGroupMember_Call) RunAndReturn(run func(interface{}, int, *gitlab.EditGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)) *GroupMembersServiceInterface_EditGroupMember_Call {
	_c.Call.Return(run)
	return _c
}

// GetGroupMember provides a mock function with given fields: gid, user, options
func (_m *GroupMembersServiceInterface) GetGroupMember(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, user)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupMember")
	}

	var r0 *gitlab.GroupMember
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)); ok {
		return rf(gid, user, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.GroupMember); ok {
		r0 = rf(gid, user, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupMember)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, user, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, user, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupMembersServiceInterface_GetGroupMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupMember'
type GroupMembersServiceInterface_GetGroupMember_Call struct {
	*mock.Call
}

// GetGroupMember is a helper method to define mock.On call
//   - gid interface{}
//   - user int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) GetGroupMember(gid interface{}, user interface{}, options ...interface{}) *GroupMembersServiceInterface_GetGroupMember_Call {
	return &GroupMembersServiceInterface_GetGroupMember_Call{Call: _e.mock.On("GetGroupMember",
		append([]interface{}{gid, user}, options...)...)}
}

func (_c *GroupMembersServiceInterface_GetGroupMember_Call) Run(run func(gid interface{}, user int, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_GetGroupMember_Call {
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

func (_c *GroupMembersServiceInterface_GetGroupMember_Call) Return(_a0 *gitlab.GroupMember, _a1 *gitlab.Response, _a2 error) *GroupMembersServiceInterface_GetGroupMember_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupMembersServiceInterface_GetGroupMember_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)) *GroupMembersServiceInterface_GetGroupMember_Call {
	_c.Call.Return(run)
	return _c
}

// GetInheritedGroupMember provides a mock function with given fields: gid, user, options
func (_m *GroupMembersServiceInterface) GetInheritedGroupMember(gid interface{}, user int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, user)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetInheritedGroupMember")
	}

	var r0 *gitlab.GroupMember
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)); ok {
		return rf(gid, user, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.GroupMember); ok {
		r0 = rf(gid, user, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupMember)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, user, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, user, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupMembersServiceInterface_GetInheritedGroupMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInheritedGroupMember'
type GroupMembersServiceInterface_GetInheritedGroupMember_Call struct {
	*mock.Call
}

// GetInheritedGroupMember is a helper method to define mock.On call
//   - gid interface{}
//   - user int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) GetInheritedGroupMember(gid interface{}, user interface{}, options ...interface{}) *GroupMembersServiceInterface_GetInheritedGroupMember_Call {
	return &GroupMembersServiceInterface_GetInheritedGroupMember_Call{Call: _e.mock.On("GetInheritedGroupMember",
		append([]interface{}{gid, user}, options...)...)}
}

func (_c *GroupMembersServiceInterface_GetInheritedGroupMember_Call) Run(run func(gid interface{}, user int, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_GetInheritedGroupMember_Call {
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

func (_c *GroupMembersServiceInterface_GetInheritedGroupMember_Call) Return(_a0 *gitlab.GroupMember, _a1 *gitlab.Response, _a2 error) *GroupMembersServiceInterface_GetInheritedGroupMember_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupMembersServiceInterface_GetInheritedGroupMember_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupMember, *gitlab.Response, error)) *GroupMembersServiceInterface_GetInheritedGroupMember_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveGroupMember provides a mock function with given fields: gid, user, opt, options
func (_m *GroupMembersServiceInterface) RemoveGroupMember(gid interface{}, user int, opt *gitlab.RemoveGroupMemberOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, user, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveGroupMember")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.RemoveGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(gid, user, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, *gitlab.RemoveGroupMemberOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(gid, user, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, *gitlab.RemoveGroupMemberOptions, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(gid, user, opt, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupMembersServiceInterface_RemoveGroupMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveGroupMember'
type GroupMembersServiceInterface_RemoveGroupMember_Call struct {
	*mock.Call
}

// RemoveGroupMember is a helper method to define mock.On call
//   - gid interface{}
//   - user int
//   - opt *gitlab.RemoveGroupMemberOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) RemoveGroupMember(gid interface{}, user interface{}, opt interface{}, options ...interface{}) *GroupMembersServiceInterface_RemoveGroupMember_Call {
	return &GroupMembersServiceInterface_RemoveGroupMember_Call{Call: _e.mock.On("RemoveGroupMember",
		append([]interface{}{gid, user, opt}, options...)...)}
}

func (_c *GroupMembersServiceInterface_RemoveGroupMember_Call) Run(run func(gid interface{}, user int, opt *gitlab.RemoveGroupMemberOptions, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_RemoveGroupMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(int), args[2].(*gitlab.RemoveGroupMemberOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupMembersServiceInterface_RemoveGroupMember_Call) Return(_a0 *gitlab.Response, _a1 error) *GroupMembersServiceInterface_RemoveGroupMember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GroupMembersServiceInterface_RemoveGroupMember_Call) RunAndReturn(run func(interface{}, int, *gitlab.RemoveGroupMemberOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *GroupMembersServiceInterface_RemoveGroupMember_Call {
	_c.Call.Return(run)
	return _c
}

// ShareWithGroup provides a mock function with given fields: gid, opt, options
func (_m *GroupMembersServiceInterface) ShareWithGroup(gid interface{}, opt *gitlab.ShareWithGroupOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ShareWithGroup")
	}

	var r0 *gitlab.Group
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ShareWithGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ShareWithGroupOptions, ...gitlab.RequestOptionFunc) *gitlab.Group); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ShareWithGroupOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ShareWithGroupOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupMembersServiceInterface_ShareWithGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShareWithGroup'
type GroupMembersServiceInterface_ShareWithGroup_Call struct {
	*mock.Call
}

// ShareWithGroup is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.ShareWithGroupOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupMembersServiceInterface_Expecter) ShareWithGroup(gid interface{}, opt interface{}, options ...interface{}) *GroupMembersServiceInterface_ShareWithGroup_Call {
	return &GroupMembersServiceInterface_ShareWithGroup_Call{Call: _e.mock.On("ShareWithGroup",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupMembersServiceInterface_ShareWithGroup_Call) Run(run func(gid interface{}, opt *gitlab.ShareWithGroupOptions, options ...gitlab.RequestOptionFunc)) *GroupMembersServiceInterface_ShareWithGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ShareWithGroupOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupMembersServiceInterface_ShareWithGroup_Call) Return(_a0 *gitlab.Group, _a1 *gitlab.Response, _a2 error) *GroupMembersServiceInterface_ShareWithGroup_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupMembersServiceInterface_ShareWithGroup_Call) RunAndReturn(run func(interface{}, *gitlab.ShareWithGroupOptions, ...gitlab.RequestOptionFunc) (*gitlab.Group, *gitlab.Response, error)) *GroupMembersServiceInterface_ShareWithGroup_Call {
	_c.Call.Return(run)
	return _c
}

// NewGroupMembersServiceInterface creates a new instance of GroupMembersServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupMembersServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupMembersServiceInterface {
	mock := &GroupMembersServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

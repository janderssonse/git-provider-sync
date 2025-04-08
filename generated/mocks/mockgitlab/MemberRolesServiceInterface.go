// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// MemberRolesServiceInterface is an autogenerated mock type for the MemberRolesServiceInterface type
type MemberRolesServiceInterface struct {
	mock.Mock
}

type MemberRolesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MemberRolesServiceInterface) EXPECT() *MemberRolesServiceInterface_Expecter {
	return &MemberRolesServiceInterface_Expecter{mock: &_m.Mock}
}

// CreateInstanceMemberRole provides a mock function with given fields: opt, options
func (_m *MemberRolesServiceInterface) CreateInstanceMemberRole(opt *gitlab.CreateMemberRoleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateInstanceMemberRole")
	}

	var r0 *gitlab.MemberRole
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(*gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error)); ok {
		return rf(opt, options...)
	}
	if rf, ok := ret.Get(0).(func(*gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) *gitlab.MemberRole); ok {
		r0 = rf(opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.MemberRole)
		}
	}

	if rf, ok := ret.Get(1).(func(*gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(*gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MemberRolesServiceInterface_CreateInstanceMemberRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateInstanceMemberRole'
type MemberRolesServiceInterface_CreateInstanceMemberRole_Call struct {
	*mock.Call
}

// CreateInstanceMemberRole is a helper method to define mock.On call
//   - opt *gitlab.CreateMemberRoleOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) CreateInstanceMemberRole(opt interface{}, options ...interface{}) *MemberRolesServiceInterface_CreateInstanceMemberRole_Call {
	return &MemberRolesServiceInterface_CreateInstanceMemberRole_Call{Call: _e.mock.On("CreateInstanceMemberRole",
		append([]interface{}{opt}, options...)...)}
}

func (_c *MemberRolesServiceInterface_CreateInstanceMemberRole_Call) Run(run func(opt *gitlab.CreateMemberRoleOptions, options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_CreateInstanceMemberRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(*gitlab.CreateMemberRoleOptions), variadicArgs...)
	})
	return _c
}

func (_c *MemberRolesServiceInterface_CreateInstanceMemberRole_Call) Return(_a0 *gitlab.MemberRole, _a1 *gitlab.Response, _a2 error) *MemberRolesServiceInterface_CreateInstanceMemberRole_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MemberRolesServiceInterface_CreateInstanceMemberRole_Call) RunAndReturn(run func(*gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error)) *MemberRolesServiceInterface_CreateInstanceMemberRole_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMemberRole provides a mock function with given fields: gid, opt, options
func (_m *MemberRolesServiceInterface) CreateMemberRole(gid interface{}, opt *gitlab.CreateMemberRoleOptions, options ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMemberRole")
	}

	var r0 *gitlab.MemberRole
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) *gitlab.MemberRole); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.MemberRole)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MemberRolesServiceInterface_CreateMemberRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMemberRole'
type MemberRolesServiceInterface_CreateMemberRole_Call struct {
	*mock.Call
}

// CreateMemberRole is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.CreateMemberRoleOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) CreateMemberRole(gid interface{}, opt interface{}, options ...interface{}) *MemberRolesServiceInterface_CreateMemberRole_Call {
	return &MemberRolesServiceInterface_CreateMemberRole_Call{Call: _e.mock.On("CreateMemberRole",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *MemberRolesServiceInterface_CreateMemberRole_Call) Run(run func(gid interface{}, opt *gitlab.CreateMemberRoleOptions, options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_CreateMemberRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CreateMemberRoleOptions), variadicArgs...)
	})
	return _c
}

func (_c *MemberRolesServiceInterface_CreateMemberRole_Call) Return(_a0 *gitlab.MemberRole, _a1 *gitlab.Response, _a2 error) *MemberRolesServiceInterface_CreateMemberRole_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MemberRolesServiceInterface_CreateMemberRole_Call) RunAndReturn(run func(interface{}, *gitlab.CreateMemberRoleOptions, ...gitlab.RequestOptionFunc) (*gitlab.MemberRole, *gitlab.Response, error)) *MemberRolesServiceInterface_CreateMemberRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteInstanceMemberRole provides a mock function with given fields: memberRoleID, options
func (_m *MemberRolesServiceInterface) DeleteInstanceMemberRole(memberRoleID int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, memberRoleID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteInstanceMemberRole")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(memberRoleID, options...)
	}
	if rf, ok := ret.Get(0).(func(int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(memberRoleID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(memberRoleID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemberRolesServiceInterface_DeleteInstanceMemberRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteInstanceMemberRole'
type MemberRolesServiceInterface_DeleteInstanceMemberRole_Call struct {
	*mock.Call
}

// DeleteInstanceMemberRole is a helper method to define mock.On call
//   - memberRoleID int
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) DeleteInstanceMemberRole(memberRoleID interface{}, options ...interface{}) *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call {
	return &MemberRolesServiceInterface_DeleteInstanceMemberRole_Call{Call: _e.mock.On("DeleteInstanceMemberRole",
		append([]interface{}{memberRoleID}, options...)...)}
}

func (_c *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call) Run(run func(memberRoleID int, options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(int), variadicArgs...)
	})
	return _c
}

func (_c *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call) Return(_a0 *gitlab.Response, _a1 error) *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call) RunAndReturn(run func(int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MemberRolesServiceInterface_DeleteInstanceMemberRole_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteMemberRole provides a mock function with given fields: gid, memberRole, options
func (_m *MemberRolesServiceInterface) DeleteMemberRole(gid interface{}, memberRole int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, memberRole)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMemberRole")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(gid, memberRole, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(gid, memberRole, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(gid, memberRole, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MemberRolesServiceInterface_DeleteMemberRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteMemberRole'
type MemberRolesServiceInterface_DeleteMemberRole_Call struct {
	*mock.Call
}

// DeleteMemberRole is a helper method to define mock.On call
//   - gid interface{}
//   - memberRole int
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) DeleteMemberRole(gid interface{}, memberRole interface{}, options ...interface{}) *MemberRolesServiceInterface_DeleteMemberRole_Call {
	return &MemberRolesServiceInterface_DeleteMemberRole_Call{Call: _e.mock.On("DeleteMemberRole",
		append([]interface{}{gid, memberRole}, options...)...)}
}

func (_c *MemberRolesServiceInterface_DeleteMemberRole_Call) Run(run func(gid interface{}, memberRole int, options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_DeleteMemberRole_Call {
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

func (_c *MemberRolesServiceInterface_DeleteMemberRole_Call) Return(_a0 *gitlab.Response, _a1 error) *MemberRolesServiceInterface_DeleteMemberRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MemberRolesServiceInterface_DeleteMemberRole_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *MemberRolesServiceInterface_DeleteMemberRole_Call {
	_c.Call.Return(run)
	return _c
}

// ListInstanceMemberRoles provides a mock function with given fields: options
func (_m *MemberRolesServiceInterface) ListInstanceMemberRoles(options ...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListInstanceMemberRoles")
	}

	var r0 []*gitlab.MemberRole
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error)); ok {
		return rf(options...)
	}
	if rf, ok := ret.Get(0).(func(...gitlab.RequestOptionFunc) []*gitlab.MemberRole); ok {
		r0 = rf(options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.MemberRole)
		}
	}

	if rf, ok := ret.Get(1).(func(...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MemberRolesServiceInterface_ListInstanceMemberRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListInstanceMemberRoles'
type MemberRolesServiceInterface_ListInstanceMemberRoles_Call struct {
	*mock.Call
}

// ListInstanceMemberRoles is a helper method to define mock.On call
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) ListInstanceMemberRoles(options ...interface{}) *MemberRolesServiceInterface_ListInstanceMemberRoles_Call {
	return &MemberRolesServiceInterface_ListInstanceMemberRoles_Call{Call: _e.mock.On("ListInstanceMemberRoles",
		append([]interface{}{}, options...)...)}
}

func (_c *MemberRolesServiceInterface_ListInstanceMemberRoles_Call) Run(run func(options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_ListInstanceMemberRoles_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MemberRolesServiceInterface_ListInstanceMemberRoles_Call) Return(_a0 []*gitlab.MemberRole, _a1 *gitlab.Response, _a2 error) *MemberRolesServiceInterface_ListInstanceMemberRoles_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MemberRolesServiceInterface_ListInstanceMemberRoles_Call) RunAndReturn(run func(...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error)) *MemberRolesServiceInterface_ListInstanceMemberRoles_Call {
	_c.Call.Return(run)
	return _c
}

// ListMemberRoles provides a mock function with given fields: gid, options
func (_m *MemberRolesServiceInterface) ListMemberRoles(gid interface{}, options ...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListMemberRoles")
	}

	var r0 []*gitlab.MemberRole
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error)); ok {
		return rf(gid, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, ...gitlab.RequestOptionFunc) []*gitlab.MemberRole); ok {
		r0 = rf(gid, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.MemberRole)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MemberRolesServiceInterface_ListMemberRoles_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListMemberRoles'
type MemberRolesServiceInterface_ListMemberRoles_Call struct {
	*mock.Call
}

// ListMemberRoles is a helper method to define mock.On call
//   - gid interface{}
//   - options ...gitlab.RequestOptionFunc
func (_e *MemberRolesServiceInterface_Expecter) ListMemberRoles(gid interface{}, options ...interface{}) *MemberRolesServiceInterface_ListMemberRoles_Call {
	return &MemberRolesServiceInterface_ListMemberRoles_Call{Call: _e.mock.On("ListMemberRoles",
		append([]interface{}{gid}, options...)...)}
}

func (_c *MemberRolesServiceInterface_ListMemberRoles_Call) Run(run func(gid interface{}, options ...gitlab.RequestOptionFunc)) *MemberRolesServiceInterface_ListMemberRoles_Call {
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

func (_c *MemberRolesServiceInterface_ListMemberRoles_Call) Return(_a0 []*gitlab.MemberRole, _a1 *gitlab.Response, _a2 error) *MemberRolesServiceInterface_ListMemberRoles_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MemberRolesServiceInterface_ListMemberRoles_Call) RunAndReturn(run func(interface{}, ...gitlab.RequestOptionFunc) ([]*gitlab.MemberRole, *gitlab.Response, error)) *MemberRolesServiceInterface_ListMemberRoles_Call {
	_c.Call.Return(run)
	return _c
}

// NewMemberRolesServiceInterface creates a new instance of MemberRolesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemberRolesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MemberRolesServiceInterface {
	mock := &MemberRolesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

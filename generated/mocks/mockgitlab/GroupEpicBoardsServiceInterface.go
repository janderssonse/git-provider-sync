// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

// GroupEpicBoardsServiceInterface is an autogenerated mock type for the GroupEpicBoardsServiceInterface type
type GroupEpicBoardsServiceInterface struct {
	mock.Mock
}

type GroupEpicBoardsServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupEpicBoardsServiceInterface) EXPECT() *GroupEpicBoardsServiceInterface_Expecter {
	return &GroupEpicBoardsServiceInterface_Expecter{mock: &_m.Mock}
}

// GetGroupEpicBoard provides a mock function with given fields: gid, board, options
func (_m *GroupEpicBoardsServiceInterface) GetGroupEpicBoard(gid interface{}, board int, options ...gitlab.RequestOptionFunc) (*gitlab.GroupEpicBoard, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, board)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetGroupEpicBoard")
	}

	var r0 *gitlab.GroupEpicBoard
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupEpicBoard, *gitlab.Response, error)); ok {
		return rf(gid, board, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.GroupEpicBoard); ok {
		r0 = rf(gid, board, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupEpicBoard)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, board, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, board, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupEpicBoard'
type GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call struct {
	*mock.Call
}

// GetGroupEpicBoard is a helper method to define mock.On call
//   - gid interface{}
//   - board int
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupEpicBoardsServiceInterface_Expecter) GetGroupEpicBoard(gid interface{}, board interface{}, options ...interface{}) *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call {
	return &GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call{Call: _e.mock.On("GetGroupEpicBoard",
		append([]interface{}{gid, board}, options...)...)}
}

func (_c *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call) Run(run func(gid interface{}, board int, options ...gitlab.RequestOptionFunc)) *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call {
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

func (_c *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call) Return(_a0 *gitlab.GroupEpicBoard, _a1 *gitlab.Response, _a2 error) *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call) RunAndReturn(run func(interface{}, int, ...gitlab.RequestOptionFunc) (*gitlab.GroupEpicBoard, *gitlab.Response, error)) *GroupEpicBoardsServiceInterface_GetGroupEpicBoard_Call {
	_c.Call.Return(run)
	return _c
}

// ListGroupEpicBoards provides a mock function with given fields: gid, opt, options
func (_m *GroupEpicBoardsServiceInterface) ListGroupEpicBoards(gid interface{}, opt *gitlab.ListGroupEpicBoardsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.GroupEpicBoard, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, gid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListGroupEpicBoards")
	}

	var r0 []*gitlab.GroupEpicBoard
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListGroupEpicBoardsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupEpicBoard, *gitlab.Response, error)); ok {
		return rf(gid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListGroupEpicBoardsOptions, ...gitlab.RequestOptionFunc) []*gitlab.GroupEpicBoard); ok {
		r0 = rf(gid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.GroupEpicBoard)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListGroupEpicBoardsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(gid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListGroupEpicBoardsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(gid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListGroupEpicBoards'
type GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call struct {
	*mock.Call
}

// ListGroupEpicBoards is a helper method to define mock.On call
//   - gid interface{}
//   - opt *gitlab.ListGroupEpicBoardsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *GroupEpicBoardsServiceInterface_Expecter) ListGroupEpicBoards(gid interface{}, opt interface{}, options ...interface{}) *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call {
	return &GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call{Call: _e.mock.On("ListGroupEpicBoards",
		append([]interface{}{gid, opt}, options...)...)}
}

func (_c *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call) Run(run func(gid interface{}, opt *gitlab.ListGroupEpicBoardsOptions, options ...gitlab.RequestOptionFunc)) *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListGroupEpicBoardsOptions), variadicArgs...)
	})
	return _c
}

func (_c *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call) Return(_a0 []*gitlab.GroupEpicBoard, _a1 *gitlab.Response, _a2 error) *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call) RunAndReturn(run func(interface{}, *gitlab.ListGroupEpicBoardsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.GroupEpicBoard, *gitlab.Response, error)) *GroupEpicBoardsServiceInterface_ListGroupEpicBoards_Call {
	_c.Call.Return(run)
	return _c
}

// NewGroupEpicBoardsServiceInterface creates a new instance of GroupEpicBoardsServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGroupEpicBoardsServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *GroupEpicBoardsServiceInterface {
	mock := &GroupEpicBoardsServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

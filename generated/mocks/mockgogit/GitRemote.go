// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	model "itiquette/git-provider-sync/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// GitRemote is an autogenerated mock type for the GitRemote type
type GitRemote struct {
	mock.Mock
}

type GitRemote_Expecter struct {
	mock *mock.Mock
}

func (_m *GitRemote) EXPECT() *GitRemote_Expecter {
	return &GitRemote_Expecter{mock: &_m.Mock}
}

// CreateRemote provides a mock function with given fields: name, url, isMirror
func (_m *GitRemote) CreateRemote(name string, url string, isMirror bool) error {
	ret := _m.Called(name, url, isMirror)

	if len(ret) == 0 {
		panic("no return value specified for CreateRemote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool) error); ok {
		r0 = rf(name, url, isMirror)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GitRemote_CreateRemote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateRemote'
type GitRemote_CreateRemote_Call struct {
	*mock.Call
}

// CreateRemote is a helper method to define mock.On call
//   - name string
//   - url string
//   - isMirror bool
func (_e *GitRemote_Expecter) CreateRemote(name interface{}, url interface{}, isMirror interface{}) *GitRemote_CreateRemote_Call {
	return &GitRemote_CreateRemote_Call{Call: _e.mock.On("CreateRemote", name, url, isMirror)}
}

func (_c *GitRemote_CreateRemote_Call) Run(run func(name string, url string, isMirror bool)) *GitRemote_CreateRemote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *GitRemote_CreateRemote_Call) Return(_a0 error) *GitRemote_CreateRemote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GitRemote_CreateRemote_Call) RunAndReturn(run func(string, string, bool) error) *GitRemote_CreateRemote_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteRemote provides a mock function with given fields: name
func (_m *GitRemote) DeleteRemote(name string) error {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRemote")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GitRemote_DeleteRemote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteRemote'
type GitRemote_DeleteRemote_Call struct {
	*mock.Call
}

// DeleteRemote is a helper method to define mock.On call
//   - name string
func (_e *GitRemote_Expecter) DeleteRemote(name interface{}) *GitRemote_DeleteRemote_Call {
	return &GitRemote_DeleteRemote_Call{Call: _e.mock.On("DeleteRemote", name)}
}

func (_c *GitRemote_DeleteRemote_Call) Run(run func(name string)) *GitRemote_DeleteRemote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GitRemote_DeleteRemote_Call) Return(_a0 error) *GitRemote_DeleteRemote_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *GitRemote_DeleteRemote_Call) RunAndReturn(run func(string) error) *GitRemote_DeleteRemote_Call {
	_c.Call.Return(run)
	return _c
}

// Remote provides a mock function with given fields: name
func (_m *GitRemote) Remote(name string) (model.Remote, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for Remote")
	}

	var r0 model.Remote
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (model.Remote, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) model.Remote); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(model.Remote)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GitRemote_Remote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remote'
type GitRemote_Remote_Call struct {
	*mock.Call
}

// Remote is a helper method to define mock.On call
//   - name string
func (_e *GitRemote_Expecter) Remote(name interface{}) *GitRemote_Remote_Call {
	return &GitRemote_Remote_Call{Call: _e.mock.On("Remote", name)}
}

func (_c *GitRemote_Remote_Call) Run(run func(name string)) *GitRemote_Remote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GitRemote_Remote_Call) Return(_a0 model.Remote, _a1 error) *GitRemote_Remote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GitRemote_Remote_Call) RunAndReturn(run func(string) (model.Remote, error)) *GitRemote_Remote_Call {
	_c.Call.Return(run)
	return _c
}

// NewGitRemote creates a new instance of GitRemote. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGitRemote(t interface {
	mock.TestingT
	Cleanup(func())
}) *GitRemote {
	mock := &GitRemote{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// SPDX-FileCopyrightText: 2025 Josef Andersson
//
// SPDX-License-Identifier: EUPL-1.2

// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	interfaces "itiquette/git-provider-sync/internal/interfaces"

	mock "github.com/stretchr/testify/mock"

	model "itiquette/git-provider-sync/internal/model"
)

// TargetWriter is an autogenerated mock type for the TargetWriter type
type TargetWriter struct {
	mock.Mock
}

type TargetWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *TargetWriter) EXPECT() *TargetWriter_Expecter {
	return &TargetWriter_Expecter{mock: &_m.Mock}
}

// Push provides a mock function with given fields: ctx, repository, opt
func (_m *TargetWriter) Push(ctx context.Context, repository interfaces.GitRepository, opt model.PushOption) error {
	ret := _m.Called(ctx, repository, opt)

	if len(ret) == 0 {
		panic("no return value specified for Push")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.GitRepository, model.PushOption) error); ok {
		r0 = rf(ctx, repository, opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TargetWriter_Push_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Push'
type TargetWriter_Push_Call struct {
	*mock.Call
}

// Push is a helper method to define mock.On call
//   - ctx context.Context
//   - repository interfaces.GitRepository
//   - opt model.PushOption
func (_e *TargetWriter_Expecter) Push(ctx interface{}, repository interface{}, opt interface{}) *TargetWriter_Push_Call {
	return &TargetWriter_Push_Call{Call: _e.mock.On("Push", ctx, repository, opt)}
}

func (_c *TargetWriter_Push_Call) Run(run func(ctx context.Context, repository interfaces.GitRepository, opt model.PushOption)) *TargetWriter_Push_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interfaces.GitRepository), args[2].(model.PushOption))
	})
	return _c
}

func (_c *TargetWriter_Push_Call) Return(_a0 error) *TargetWriter_Push_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TargetWriter_Push_Call) RunAndReturn(run func(context.Context, interfaces.GitRepository, model.PushOption) error) *TargetWriter_Push_Call {
	_c.Call.Return(run)
	return _c
}

// NewTargetWriter creates a new instance of TargetWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTargetWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *TargetWriter {
	mock := &TargetWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

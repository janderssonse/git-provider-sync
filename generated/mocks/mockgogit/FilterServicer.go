// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	functiondefinition "itiquette/git-provider-sync/internal/functiondefinition"
	interfaces "itiquette/git-provider-sync/internal/interfaces"

	mock "github.com/stretchr/testify/mock"

	model "itiquette/git-provider-sync/internal/model"
)

// FilterServicer is an autogenerated mock type for the FilterServicer type
type FilterServicer struct {
	mock.Mock
}

type FilterServicer_Expecter struct {
	mock *mock.Mock
}

func (_m *FilterServicer) EXPECT() *FilterServicer_Expecter {
	return &FilterServicer_Expecter{mock: &_m.Mock}
}

// FilterProjectinfos provides a mock function with given fields: ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval
func (_m *FilterServicer) FilterProjectinfos(ctx context.Context, opt model.ProviderOption, projectinfos []model.ProjectInfo, filterExcludedIncludedFunc functiondefinition.FilterIncludedExcludedFunc, isInInterval interfaces.IsInIntervalFunc) ([]model.ProjectInfo, error) {
	ret := _m.Called(ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval)

	if len(ret) == 0 {
		panic("no return value specified for FilterProjectinfos")
	}

	var r0 []model.ProjectInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.ProviderOption, []model.ProjectInfo, functiondefinition.FilterIncludedExcludedFunc, interfaces.IsInIntervalFunc) ([]model.ProjectInfo, error)); ok {
		return rf(ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.ProviderOption, []model.ProjectInfo, functiondefinition.FilterIncludedExcludedFunc, interfaces.IsInIntervalFunc) []model.ProjectInfo); ok {
		r0 = rf(ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ProjectInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.ProviderOption, []model.ProjectInfo, functiondefinition.FilterIncludedExcludedFunc, interfaces.IsInIntervalFunc) error); ok {
		r1 = rf(ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterServicer_FilterProjectinfos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterProjectinfos'
type FilterServicer_FilterProjectinfos_Call struct {
	*mock.Call
}

// FilterProjectinfos is a helper method to define mock.On call
//   - ctx context.Context
//   - opt model.ProviderOption
//   - projectinfos []model.ProjectInfo
//   - filterExcludedIncludedFunc functiondefinition.FilterIncludedExcludedFunc
//   - isInInterval interfaces.IsInIntervalFunc
func (_e *FilterServicer_Expecter) FilterProjectinfos(ctx interface{}, opt interface{}, projectinfos interface{}, filterExcludedIncludedFunc interface{}, isInInterval interface{}) *FilterServicer_FilterProjectinfos_Call {
	return &FilterServicer_FilterProjectinfos_Call{Call: _e.mock.On("FilterProjectinfos", ctx, opt, projectinfos, filterExcludedIncludedFunc, isInInterval)}
}

func (_c *FilterServicer_FilterProjectinfos_Call) Run(run func(ctx context.Context, opt model.ProviderOption, projectinfos []model.ProjectInfo, filterExcludedIncludedFunc functiondefinition.FilterIncludedExcludedFunc, isInInterval interfaces.IsInIntervalFunc)) *FilterServicer_FilterProjectinfos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.ProviderOption), args[2].([]model.ProjectInfo), args[3].(functiondefinition.FilterIncludedExcludedFunc), args[4].(interfaces.IsInIntervalFunc))
	})
	return _c
}

func (_c *FilterServicer_FilterProjectinfos_Call) Return(_a0 []model.ProjectInfo, _a1 error) *FilterServicer_FilterProjectinfos_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FilterServicer_FilterProjectinfos_Call) RunAndReturn(run func(context.Context, model.ProviderOption, []model.ProjectInfo, functiondefinition.FilterIncludedExcludedFunc, interfaces.IsInIntervalFunc) ([]model.ProjectInfo, error)) *FilterServicer_FilterProjectinfos_Call {
	_c.Call.Return(run)
	return _c
}

// NewFilterServicer creates a new instance of FilterServicer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilterServicer(t interface {
	mock.TestingT
	Cleanup(func())
}) *FilterServicer {
	mock := &FilterServicer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

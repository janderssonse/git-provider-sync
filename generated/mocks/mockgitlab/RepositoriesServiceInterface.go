// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	io "io"

	gitlab "gitlab.com/gitlab-org/api/client-go"

	mock "github.com/stretchr/testify/mock"
)

// RepositoriesServiceInterface is an autogenerated mock type for the RepositoriesServiceInterface type
type RepositoriesServiceInterface struct {
	mock.Mock
}

type RepositoriesServiceInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *RepositoriesServiceInterface) EXPECT() *RepositoriesServiceInterface_Expecter {
	return &RepositoriesServiceInterface_Expecter{mock: &_m.Mock}
}

// AddChangelog provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) AddChangelog(pid interface{}, opt *gitlab.AddChangelogOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddChangelog")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddChangelogOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.AddChangelogOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.AddChangelogOptions, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, opt, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoriesServiceInterface_AddChangelog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddChangelog'
type RepositoriesServiceInterface_AddChangelog_Call struct {
	*mock.Call
}

// AddChangelog is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.AddChangelogOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) AddChangelog(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_AddChangelog_Call {
	return &RepositoriesServiceInterface_AddChangelog_Call{Call: _e.mock.On("AddChangelog",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_AddChangelog_Call) Run(run func(pid interface{}, opt *gitlab.AddChangelogOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_AddChangelog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.AddChangelogOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_AddChangelog_Call) Return(_a0 *gitlab.Response, _a1 error) *RepositoriesServiceInterface_AddChangelog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoriesServiceInterface_AddChangelog_Call) RunAndReturn(run func(interface{}, *gitlab.AddChangelogOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *RepositoriesServiceInterface_AddChangelog_Call {
	_c.Call.Return(run)
	return _c
}

// Archive provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) Archive(pid interface{}, opt *gitlab.ArchiveOptions, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Archive")
	}

	var r0 []byte
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) []byte); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_Archive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Archive'
type RepositoriesServiceInterface_Archive_Call struct {
	*mock.Call
}

// Archive is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ArchiveOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) Archive(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_Archive_Call {
	return &RepositoriesServiceInterface_Archive_Call{Call: _e.mock.On("Archive",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_Archive_Call) Run(run func(pid interface{}, opt *gitlab.ArchiveOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_Archive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ArchiveOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_Archive_Call) Return(_a0 []byte, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_Archive_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_Archive_Call) RunAndReturn(run func(interface{}, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *RepositoriesServiceInterface_Archive_Call {
	_c.Call.Return(run)
	return _c
}

// Blob provides a mock function with given fields: pid, sha, options
func (_m *RepositoriesServiceInterface) Blob(pid interface{}, sha string, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, sha)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Blob")
	}

	var r0 []byte
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)); ok {
		return rf(pid, sha, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) []byte); ok {
		r0 = rf(pid, sha, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, sha, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, string, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, sha, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_Blob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Blob'
type RepositoriesServiceInterface_Blob_Call struct {
	*mock.Call
}

// Blob is a helper method to define mock.On call
//   - pid interface{}
//   - sha string
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) Blob(pid interface{}, sha interface{}, options ...interface{}) *RepositoriesServiceInterface_Blob_Call {
	return &RepositoriesServiceInterface_Blob_Call{Call: _e.mock.On("Blob",
		append([]interface{}{pid, sha}, options...)...)}
}

func (_c *RepositoriesServiceInterface_Blob_Call) Run(run func(pid interface{}, sha string, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_Blob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_Blob_Call) Return(_a0 []byte, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_Blob_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_Blob_Call) RunAndReturn(run func(interface{}, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *RepositoriesServiceInterface_Blob_Call {
	_c.Call.Return(run)
	return _c
}

// Compare provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) Compare(pid interface{}, opt *gitlab.CompareOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Compare, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Compare")
	}

	var r0 *gitlab.Compare
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CompareOptions, ...gitlab.RequestOptionFunc) (*gitlab.Compare, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.CompareOptions, ...gitlab.RequestOptionFunc) *gitlab.Compare); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Compare)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.CompareOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.CompareOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_Compare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Compare'
type RepositoriesServiceInterface_Compare_Call struct {
	*mock.Call
}

// Compare is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.CompareOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) Compare(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_Compare_Call {
	return &RepositoriesServiceInterface_Compare_Call{Call: _e.mock.On("Compare",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_Compare_Call) Run(run func(pid interface{}, opt *gitlab.CompareOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_Compare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.CompareOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_Compare_Call) Return(_a0 *gitlab.Compare, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_Compare_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_Compare_Call) RunAndReturn(run func(interface{}, *gitlab.CompareOptions, ...gitlab.RequestOptionFunc) (*gitlab.Compare, *gitlab.Response, error)) *RepositoriesServiceInterface_Compare_Call {
	_c.Call.Return(run)
	return _c
}

// Contributors provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) Contributors(pid interface{}, opt *gitlab.ListContributorsOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.Contributor, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Contributors")
	}

	var r0 []*gitlab.Contributor
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListContributorsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Contributor, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListContributorsOptions, ...gitlab.RequestOptionFunc) []*gitlab.Contributor); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Contributor)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListContributorsOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListContributorsOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_Contributors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contributors'
type RepositoriesServiceInterface_Contributors_Call struct {
	*mock.Call
}

// Contributors is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListContributorsOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) Contributors(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_Contributors_Call {
	return &RepositoriesServiceInterface_Contributors_Call{Call: _e.mock.On("Contributors",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_Contributors_Call) Run(run func(pid interface{}, opt *gitlab.ListContributorsOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_Contributors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListContributorsOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_Contributors_Call) Return(_a0 []*gitlab.Contributor, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_Contributors_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_Contributors_Call) RunAndReturn(run func(interface{}, *gitlab.ListContributorsOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.Contributor, *gitlab.Response, error)) *RepositoriesServiceInterface_Contributors_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateChangelogData provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) GenerateChangelogData(pid interface{}, opt gitlab.GenerateChangelogDataOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ChangelogData, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GenerateChangelogData")
	}

	var r0 *gitlab.ChangelogData
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.GenerateChangelogDataOptions, ...gitlab.RequestOptionFunc) (*gitlab.ChangelogData, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, gitlab.GenerateChangelogDataOptions, ...gitlab.RequestOptionFunc) *gitlab.ChangelogData); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.ChangelogData)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, gitlab.GenerateChangelogDataOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, gitlab.GenerateChangelogDataOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_GenerateChangelogData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateChangelogData'
type RepositoriesServiceInterface_GenerateChangelogData_Call struct {
	*mock.Call
}

// GenerateChangelogData is a helper method to define mock.On call
//   - pid interface{}
//   - opt gitlab.GenerateChangelogDataOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) GenerateChangelogData(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_GenerateChangelogData_Call {
	return &RepositoriesServiceInterface_GenerateChangelogData_Call{Call: _e.mock.On("GenerateChangelogData",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_GenerateChangelogData_Call) Run(run func(pid interface{}, opt gitlab.GenerateChangelogDataOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_GenerateChangelogData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(gitlab.GenerateChangelogDataOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_GenerateChangelogData_Call) Return(_a0 *gitlab.ChangelogData, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_GenerateChangelogData_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_GenerateChangelogData_Call) RunAndReturn(run func(interface{}, gitlab.GenerateChangelogDataOptions, ...gitlab.RequestOptionFunc) (*gitlab.ChangelogData, *gitlab.Response, error)) *RepositoriesServiceInterface_GenerateChangelogData_Call {
	_c.Call.Return(run)
	return _c
}

// ListTree provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) ListTree(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListTree")
	}

	var r0 []*gitlab.TreeNode
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListTreeOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListTreeOptions, ...gitlab.RequestOptionFunc) []*gitlab.TreeNode); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.TreeNode)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListTreeOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListTreeOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_ListTree_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTree'
type RepositoriesServiceInterface_ListTree_Call struct {
	*mock.Call
}

// ListTree is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.ListTreeOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) ListTree(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_ListTree_Call {
	return &RepositoriesServiceInterface_ListTree_Call{Call: _e.mock.On("ListTree",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_ListTree_Call) Run(run func(pid interface{}, opt *gitlab.ListTreeOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_ListTree_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.ListTreeOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_ListTree_Call) Return(_a0 []*gitlab.TreeNode, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_ListTree_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_ListTree_Call) RunAndReturn(run func(interface{}, *gitlab.ListTreeOptions, ...gitlab.RequestOptionFunc) ([]*gitlab.TreeNode, *gitlab.Response, error)) *RepositoriesServiceInterface_ListTree_Call {
	_c.Call.Return(run)
	return _c
}

// MergeBase provides a mock function with given fields: pid, opt, options
func (_m *RepositoriesServiceInterface) MergeBase(pid interface{}, opt *gitlab.MergeBaseOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Commit, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MergeBase")
	}

	var r0 *gitlab.Commit
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.MergeBaseOptions, ...gitlab.RequestOptionFunc) (*gitlab.Commit, *gitlab.Response, error)); ok {
		return rf(pid, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.MergeBaseOptions, ...gitlab.RequestOptionFunc) *gitlab.Commit); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.MergeBaseOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.MergeBaseOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_MergeBase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MergeBase'
type RepositoriesServiceInterface_MergeBase_Call struct {
	*mock.Call
}

// MergeBase is a helper method to define mock.On call
//   - pid interface{}
//   - opt *gitlab.MergeBaseOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) MergeBase(pid interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_MergeBase_Call {
	return &RepositoriesServiceInterface_MergeBase_Call{Call: _e.mock.On("MergeBase",
		append([]interface{}{pid, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_MergeBase_Call) Run(run func(pid interface{}, opt *gitlab.MergeBaseOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_MergeBase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(*gitlab.MergeBaseOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_MergeBase_Call) Return(_a0 *gitlab.Commit, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_MergeBase_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_MergeBase_Call) RunAndReturn(run func(interface{}, *gitlab.MergeBaseOptions, ...gitlab.RequestOptionFunc) (*gitlab.Commit, *gitlab.Response, error)) *RepositoriesServiceInterface_MergeBase_Call {
	_c.Call.Return(run)
	return _c
}

// RawBlobContent provides a mock function with given fields: pid, sha, options
func (_m *RepositoriesServiceInterface) RawBlobContent(pid interface{}, sha string, options ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, sha)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RawBlobContent")
	}

	var r0 []byte
	var r1 *gitlab.Response
	var r2 error
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)); ok {
		return rf(pid, sha, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, string, ...gitlab.RequestOptionFunc) []byte); ok {
		r0 = rf(pid, sha, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, string, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, sha, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(2).(func(interface{}, string, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, sha, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RepositoriesServiceInterface_RawBlobContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RawBlobContent'
type RepositoriesServiceInterface_RawBlobContent_Call struct {
	*mock.Call
}

// RawBlobContent is a helper method to define mock.On call
//   - pid interface{}
//   - sha string
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) RawBlobContent(pid interface{}, sha interface{}, options ...interface{}) *RepositoriesServiceInterface_RawBlobContent_Call {
	return &RepositoriesServiceInterface_RawBlobContent_Call{Call: _e.mock.On("RawBlobContent",
		append([]interface{}{pid, sha}, options...)...)}
}

func (_c *RepositoriesServiceInterface_RawBlobContent_Call) Run(run func(pid interface{}, sha string, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_RawBlobContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_RawBlobContent_Call) Return(_a0 []byte, _a1 *gitlab.Response, _a2 error) *RepositoriesServiceInterface_RawBlobContent_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *RepositoriesServiceInterface_RawBlobContent_Call) RunAndReturn(run func(interface{}, string, ...gitlab.RequestOptionFunc) ([]byte, *gitlab.Response, error)) *RepositoriesServiceInterface_RawBlobContent_Call {
	_c.Call.Return(run)
	return _c
}

// StreamArchive provides a mock function with given fields: pid, w, opt, options
func (_m *RepositoriesServiceInterface) StreamArchive(pid interface{}, w io.Writer, opt *gitlab.ArchiveOptions, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, w, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StreamArchive")
	}

	var r0 *gitlab.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}, io.Writer, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)); ok {
		return rf(pid, w, opt, options...)
	}
	if rf, ok := ret.Get(0).(func(interface{}, io.Writer, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r0 = rf(pid, w, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}, io.Writer, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) error); ok {
		r1 = rf(pid, w, opt, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoriesServiceInterface_StreamArchive_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamArchive'
type RepositoriesServiceInterface_StreamArchive_Call struct {
	*mock.Call
}

// StreamArchive is a helper method to define mock.On call
//   - pid interface{}
//   - w io.Writer
//   - opt *gitlab.ArchiveOptions
//   - options ...gitlab.RequestOptionFunc
func (_e *RepositoriesServiceInterface_Expecter) StreamArchive(pid interface{}, w interface{}, opt interface{}, options ...interface{}) *RepositoriesServiceInterface_StreamArchive_Call {
	return &RepositoriesServiceInterface_StreamArchive_Call{Call: _e.mock.On("StreamArchive",
		append([]interface{}{pid, w, opt}, options...)...)}
}

func (_c *RepositoriesServiceInterface_StreamArchive_Call) Run(run func(pid interface{}, w io.Writer, opt *gitlab.ArchiveOptions, options ...gitlab.RequestOptionFunc)) *RepositoriesServiceInterface_StreamArchive_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]gitlab.RequestOptionFunc, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(gitlab.RequestOptionFunc)
			}
		}
		run(args[0].(interface{}), args[1].(io.Writer), args[2].(*gitlab.ArchiveOptions), variadicArgs...)
	})
	return _c
}

func (_c *RepositoriesServiceInterface_StreamArchive_Call) Return(_a0 *gitlab.Response, _a1 error) *RepositoriesServiceInterface_StreamArchive_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoriesServiceInterface_StreamArchive_Call) RunAndReturn(run func(interface{}, io.Writer, *gitlab.ArchiveOptions, ...gitlab.RequestOptionFunc) (*gitlab.Response, error)) *RepositoriesServiceInterface_StreamArchive_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepositoriesServiceInterface creates a new instance of RepositoriesServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoriesServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoriesServiceInterface {
	mock := &RepositoriesServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

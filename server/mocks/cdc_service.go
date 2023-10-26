// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	request "github.com/zilliztech/milvus-cdc/server/model/request"
)

// CDCService is an autogenerated mock type for the CDCService type
type CDCService struct {
	mock.Mock
}

type CDCService_Expecter struct {
	mock *mock.Mock
}

func (_m *CDCService) EXPECT() *CDCService_Expecter {
	return &CDCService_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *CDCService) Create(_a0 *request.CreateRequest) (*request.CreateResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.CreateRequest) (*request.CreateResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.CreateRequest) *request.CreateResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.CreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type CDCService_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//  - _a0 *request.CreateRequest
func (_e *CDCService_Expecter) Create(_a0 interface{}) *CDCService_Create_Call {
	return &CDCService_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *CDCService_Create_Call) Run(run func(_a0 *request.CreateRequest)) *CDCService_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.CreateRequest))
	})
	return _c
}

func (_c *CDCService_Create_Call) Return(_a0 *request.CreateResponse, _a1 error) *CDCService_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_Create_Call) RunAndReturn(run func(*request.CreateRequest) (*request.CreateResponse, error)) *CDCService_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0
func (_m *CDCService) Delete(_a0 *request.DeleteRequest) (*request.DeleteResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.DeleteResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.DeleteRequest) (*request.DeleteResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.DeleteRequest) *request.DeleteResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.DeleteResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.DeleteRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CDCService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//  - _a0 *request.DeleteRequest
func (_e *CDCService_Expecter) Delete(_a0 interface{}) *CDCService_Delete_Call {
	return &CDCService_Delete_Call{Call: _e.mock.On("Delete", _a0)}
}

func (_c *CDCService_Delete_Call) Run(run func(_a0 *request.DeleteRequest)) *CDCService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.DeleteRequest))
	})
	return _c
}

func (_c *CDCService_Delete_Call) Return(_a0 *request.DeleteResponse, _a1 error) *CDCService_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_Delete_Call) RunAndReturn(run func(*request.DeleteRequest) (*request.DeleteResponse, error)) *CDCService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: _a0
func (_m *CDCService) Get(_a0 *request.GetRequest) (*request.GetResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.GetResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.GetRequest) (*request.GetResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.GetRequest) *request.GetResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.GetResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.GetRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CDCService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 *request.GetRequest
func (_e *CDCService_Expecter) Get(_a0 interface{}) *CDCService_Get_Call {
	return &CDCService_Get_Call{Call: _e.mock.On("Get", _a0)}
}

func (_c *CDCService_Get_Call) Run(run func(_a0 *request.GetRequest)) *CDCService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.GetRequest))
	})
	return _c
}

func (_c *CDCService_Get_Call) Return(_a0 *request.GetResponse, _a1 error) *CDCService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_Get_Call) RunAndReturn(run func(*request.GetRequest) (*request.GetResponse, error)) *CDCService_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetPosition provides a mock function with given fields: req
func (_m *CDCService) GetPosition(req *request.GetPositionRequest) (*request.GetPositionResponse, error) {
	ret := _m.Called(req)

	var r0 *request.GetPositionResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.GetPositionRequest) (*request.GetPositionResponse, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*request.GetPositionRequest) *request.GetPositionResponse); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.GetPositionResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.GetPositionRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_GetPosition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPosition'
type CDCService_GetPosition_Call struct {
	*mock.Call
}

// GetPosition is a helper method to define mock.On call
//  - req *request.GetPositionRequest
func (_e *CDCService_Expecter) GetPosition(req interface{}) *CDCService_GetPosition_Call {
	return &CDCService_GetPosition_Call{Call: _e.mock.On("GetPosition", req)}
}

func (_c *CDCService_GetPosition_Call) Run(run func(req *request.GetPositionRequest)) *CDCService_GetPosition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.GetPositionRequest))
	})
	return _c
}

func (_c *CDCService_GetPosition_Call) Return(_a0 *request.GetPositionResponse, _a1 error) *CDCService_GetPosition_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_GetPosition_Call) RunAndReturn(run func(*request.GetPositionRequest) (*request.GetPositionResponse, error)) *CDCService_GetPosition_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: _a0
func (_m *CDCService) List(_a0 *request.ListRequest) (*request.ListResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.ListResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.ListRequest) (*request.ListResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.ListRequest) *request.ListResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.ListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.ListRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type CDCService_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//  - _a0 *request.ListRequest
func (_e *CDCService_Expecter) List(_a0 interface{}) *CDCService_List_Call {
	return &CDCService_List_Call{Call: _e.mock.On("List", _a0)}
}

func (_c *CDCService_List_Call) Run(run func(_a0 *request.ListRequest)) *CDCService_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.ListRequest))
	})
	return _c
}

func (_c *CDCService_List_Call) Return(_a0 *request.ListResponse, _a1 error) *CDCService_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_List_Call) RunAndReturn(run func(*request.ListRequest) (*request.ListResponse, error)) *CDCService_List_Call {
	_c.Call.Return(run)
	return _c
}

// Pause provides a mock function with given fields: _a0
func (_m *CDCService) Pause(_a0 *request.PauseRequest) (*request.PauseResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.PauseResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.PauseRequest) (*request.PauseResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.PauseRequest) *request.PauseResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.PauseResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.PauseRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_Pause_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pause'
type CDCService_Pause_Call struct {
	*mock.Call
}

// Pause is a helper method to define mock.On call
//  - _a0 *request.PauseRequest
func (_e *CDCService_Expecter) Pause(_a0 interface{}) *CDCService_Pause_Call {
	return &CDCService_Pause_Call{Call: _e.mock.On("Pause", _a0)}
}

func (_c *CDCService_Pause_Call) Run(run func(_a0 *request.PauseRequest)) *CDCService_Pause_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.PauseRequest))
	})
	return _c
}

func (_c *CDCService_Pause_Call) Return(_a0 *request.PauseResponse, _a1 error) *CDCService_Pause_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_Pause_Call) RunAndReturn(run func(*request.PauseRequest) (*request.PauseResponse, error)) *CDCService_Pause_Call {
	_c.Call.Return(run)
	return _c
}

// ReloadTask provides a mock function with given fields:
func (_m *CDCService) ReloadTask() {
	_m.Called()
}

// CDCService_ReloadTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReloadTask'
type CDCService_ReloadTask_Call struct {
	*mock.Call
}

// ReloadTask is a helper method to define mock.On call
func (_e *CDCService_Expecter) ReloadTask() *CDCService_ReloadTask_Call {
	return &CDCService_ReloadTask_Call{Call: _e.mock.On("ReloadTask")}
}

func (_c *CDCService_ReloadTask_Call) Run(run func()) *CDCService_ReloadTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CDCService_ReloadTask_Call) Return() *CDCService_ReloadTask_Call {
	_c.Call.Return()
	return _c
}

func (_c *CDCService_ReloadTask_Call) RunAndReturn(run func()) *CDCService_ReloadTask_Call {
	_c.Call.Return(run)
	return _c
}

// Resume provides a mock function with given fields: _a0
func (_m *CDCService) Resume(_a0 *request.ResumeRequest) (*request.ResumeResponse, error) {
	ret := _m.Called(_a0)

	var r0 *request.ResumeResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*request.ResumeRequest) (*request.ResumeResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*request.ResumeRequest) *request.ResumeResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.ResumeResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*request.ResumeRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CDCService_Resume_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resume'
type CDCService_Resume_Call struct {
	*mock.Call
}

// Resume is a helper method to define mock.On call
//  - _a0 *request.ResumeRequest
func (_e *CDCService_Expecter) Resume(_a0 interface{}) *CDCService_Resume_Call {
	return &CDCService_Resume_Call{Call: _e.mock.On("Resume", _a0)}
}

func (_c *CDCService_Resume_Call) Run(run func(_a0 *request.ResumeRequest)) *CDCService_Resume_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*request.ResumeRequest))
	})
	return _c
}

func (_c *CDCService_Resume_Call) Return(_a0 *request.ResumeResponse, _a1 error) *CDCService_Resume_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CDCService_Resume_Call) RunAndReturn(run func(*request.ResumeRequest) (*request.ResumeResponse, error)) *CDCService_Resume_Call {
	_c.Call.Return(run)
	return _c
}

// NewCDCService creates a new instance of CDCService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCDCService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CDCService {
	mock := &CDCService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	io "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io"

	mock "github.com/stretchr/testify/mock"
)

// OutputReader is an autogenerated mock type for the OutputReader type
type OutputReader struct {
	mock.Mock
}

type OutputReader_DeckExists struct {
	*mock.Call
}

func (_m OutputReader_DeckExists) Return(_a0 bool, _a1 error) *OutputReader_DeckExists {
	return &OutputReader_DeckExists{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OutputReader) OnDeckExists(ctx context.Context) *OutputReader_DeckExists {
	c_call := _m.On("DeckExists", ctx)
	return &OutputReader_DeckExists{Call: c_call}
}

func (_m *OutputReader) OnDeckExistsMatch(matchers ...interface{}) *OutputReader_DeckExists {
	c_call := _m.On("DeckExists", matchers...)
	return &OutputReader_DeckExists{Call: c_call}
}

// DeckExists provides a mock function with given fields: ctx
func (_m *OutputReader) DeckExists(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OutputReader_Exists struct {
	*mock.Call
}

func (_m OutputReader_Exists) Return(_a0 bool, _a1 error) *OutputReader_Exists {
	return &OutputReader_Exists{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OutputReader) OnExists(ctx context.Context) *OutputReader_Exists {
	c_call := _m.On("Exists", ctx)
	return &OutputReader_Exists{Call: c_call}
}

func (_m *OutputReader) OnExistsMatch(matchers ...interface{}) *OutputReader_Exists {
	c_call := _m.On("Exists", matchers...)
	return &OutputReader_Exists{Call: c_call}
}

// Exists provides a mock function with given fields: ctx
func (_m *OutputReader) Exists(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OutputReader_IsError struct {
	*mock.Call
}

func (_m OutputReader_IsError) Return(_a0 bool, _a1 error) *OutputReader_IsError {
	return &OutputReader_IsError{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OutputReader) OnIsError(ctx context.Context) *OutputReader_IsError {
	c_call := _m.On("IsError", ctx)
	return &OutputReader_IsError{Call: c_call}
}

func (_m *OutputReader) OnIsErrorMatch(matchers ...interface{}) *OutputReader_IsError {
	c_call := _m.On("IsError", matchers...)
	return &OutputReader_IsError{Call: c_call}
}

// IsError provides a mock function with given fields: ctx
func (_m *OutputReader) IsError(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OutputReader_IsFile struct {
	*mock.Call
}

func (_m OutputReader_IsFile) Return(_a0 bool) *OutputReader_IsFile {
	return &OutputReader_IsFile{Call: _m.Call.Return(_a0)}
}

func (_m *OutputReader) OnIsFile(ctx context.Context) *OutputReader_IsFile {
	c_call := _m.On("IsFile", ctx)
	return &OutputReader_IsFile{Call: c_call}
}

func (_m *OutputReader) OnIsFileMatch(matchers ...interface{}) *OutputReader_IsFile {
	c_call := _m.On("IsFile", matchers...)
	return &OutputReader_IsFile{Call: c_call}
}

// IsFile provides a mock function with given fields: ctx
func (_m *OutputReader) IsFile(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type OutputReader_Read struct {
	*mock.Call
}

func (_m OutputReader_Read) Return(_a0 *core.LiteralMap, _a1 *io.ExecutionError, _a2 error) *OutputReader_Read {
	return &OutputReader_Read{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *OutputReader) OnRead(ctx context.Context) *OutputReader_Read {
	c_call := _m.On("Read", ctx)
	return &OutputReader_Read{Call: c_call}
}

func (_m *OutputReader) OnReadMatch(matchers ...interface{}) *OutputReader_Read {
	c_call := _m.On("Read", matchers...)
	return &OutputReader_Read{Call: c_call}
}

// Read provides a mock function with given fields: ctx
func (_m *OutputReader) Read(ctx context.Context) (*core.LiteralMap, *io.ExecutionError, error) {
	ret := _m.Called(ctx)

	var r0 *core.LiteralMap
	if rf, ok := ret.Get(0).(func(context.Context) *core.LiteralMap); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.LiteralMap)
		}
	}

	var r1 *io.ExecutionError
	if rf, ok := ret.Get(1).(func(context.Context) *io.ExecutionError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*io.ExecutionError)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type OutputReader_ReadError struct {
	*mock.Call
}

func (_m OutputReader_ReadError) Return(_a0 io.ExecutionError, _a1 error) *OutputReader_ReadError {
	return &OutputReader_ReadError{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OutputReader) OnReadError(ctx context.Context) *OutputReader_ReadError {
	c_call := _m.On("ReadError", ctx)
	return &OutputReader_ReadError{Call: c_call}
}

func (_m *OutputReader) OnReadErrorMatch(matchers ...interface{}) *OutputReader_ReadError {
	c_call := _m.On("ReadError", matchers...)
	return &OutputReader_ReadError{Call: c_call}
}

// ReadError provides a mock function with given fields: ctx
func (_m *OutputReader) ReadError(ctx context.Context) (io.ExecutionError, error) {
	ret := _m.Called(ctx)

	var r0 io.ExecutionError
	if rf, ok := ret.Get(0).(func(context.Context) io.ExecutionError); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(io.ExecutionError)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

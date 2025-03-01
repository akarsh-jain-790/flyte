// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// ExecutableWorkflowStatus is an autogenerated mock type for the ExecutableWorkflowStatus type
type ExecutableWorkflowStatus struct {
	mock.Mock
}

type ExecutableWorkflowStatus_ConstructNodeDataDir struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_ConstructNodeDataDir) Return(_a0 storage.DataReference, _a1 error) *ExecutableWorkflowStatus_ConstructNodeDataDir {
	return &ExecutableWorkflowStatus_ConstructNodeDataDir{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ExecutableWorkflowStatus) OnConstructNodeDataDir(ctx context.Context, name string) *ExecutableWorkflowStatus_ConstructNodeDataDir {
	c_call := _m.On("ConstructNodeDataDir", ctx, name)
	return &ExecutableWorkflowStatus_ConstructNodeDataDir{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnConstructNodeDataDirMatch(matchers ...interface{}) *ExecutableWorkflowStatus_ConstructNodeDataDir {
	c_call := _m.On("ConstructNodeDataDir", matchers...)
	return &ExecutableWorkflowStatus_ConstructNodeDataDir{Call: c_call}
}

// ConstructNodeDataDir provides a mock function with given fields: ctx, name
func (_m *ExecutableWorkflowStatus) ConstructNodeDataDir(ctx context.Context, name string) (storage.DataReference, error) {
	ret := _m.Called(ctx, name)

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func(context.Context, string) storage.DataReference); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ExecutableWorkflowStatus_GetDataDir struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetDataDir) Return(_a0 storage.DataReference) *ExecutableWorkflowStatus_GetDataDir {
	return &ExecutableWorkflowStatus_GetDataDir{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetDataDir() *ExecutableWorkflowStatus_GetDataDir {
	c_call := _m.On("GetDataDir")
	return &ExecutableWorkflowStatus_GetDataDir{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetDataDirMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetDataDir {
	c_call := _m.On("GetDataDir", matchers...)
	return &ExecutableWorkflowStatus_GetDataDir{Call: c_call}
}

// GetDataDir provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetDataDir() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type ExecutableWorkflowStatus_GetExecutionError struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetExecutionError) Return(_a0 *core.ExecutionError) *ExecutableWorkflowStatus_GetExecutionError {
	return &ExecutableWorkflowStatus_GetExecutionError{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetExecutionError() *ExecutableWorkflowStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError")
	return &ExecutableWorkflowStatus_GetExecutionError{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetExecutionErrorMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetExecutionError {
	c_call := _m.On("GetExecutionError", matchers...)
	return &ExecutableWorkflowStatus_GetExecutionError{Call: c_call}
}

// GetExecutionError provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

type ExecutableWorkflowStatus_GetLastUpdatedAt struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetLastUpdatedAt) Return(_a0 *v1.Time) *ExecutableWorkflowStatus_GetLastUpdatedAt {
	return &ExecutableWorkflowStatus_GetLastUpdatedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetLastUpdatedAt() *ExecutableWorkflowStatus_GetLastUpdatedAt {
	c_call := _m.On("GetLastUpdatedAt")
	return &ExecutableWorkflowStatus_GetLastUpdatedAt{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetLastUpdatedAtMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetLastUpdatedAt {
	c_call := _m.On("GetLastUpdatedAt", matchers...)
	return &ExecutableWorkflowStatus_GetLastUpdatedAt{Call: c_call}
}

// GetLastUpdatedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetLastUpdatedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableWorkflowStatus_GetMessage struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetMessage) Return(_a0 string) *ExecutableWorkflowStatus_GetMessage {
	return &ExecutableWorkflowStatus_GetMessage{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetMessage() *ExecutableWorkflowStatus_GetMessage {
	c_call := _m.On("GetMessage")
	return &ExecutableWorkflowStatus_GetMessage{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetMessageMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetMessage {
	c_call := _m.On("GetMessage", matchers...)
	return &ExecutableWorkflowStatus_GetMessage{Call: c_call}
}

// GetMessage provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetMessage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type ExecutableWorkflowStatus_GetNodeExecutionStatus struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetNodeExecutionStatus) Return(_a0 v1alpha1.ExecutableNodeStatus) *ExecutableWorkflowStatus_GetNodeExecutionStatus {
	return &ExecutableWorkflowStatus_GetNodeExecutionStatus{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetNodeExecutionStatus(ctx context.Context, id string) *ExecutableWorkflowStatus_GetNodeExecutionStatus {
	c_call := _m.On("GetNodeExecutionStatus", ctx, id)
	return &ExecutableWorkflowStatus_GetNodeExecutionStatus{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetNodeExecutionStatusMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetNodeExecutionStatus {
	c_call := _m.On("GetNodeExecutionStatus", matchers...)
	return &ExecutableWorkflowStatus_GetNodeExecutionStatus{Call: c_call}
}

// GetNodeExecutionStatus provides a mock function with given fields: ctx, id
func (_m *ExecutableWorkflowStatus) GetNodeExecutionStatus(ctx context.Context, id string) v1alpha1.ExecutableNodeStatus {
	ret := _m.Called(ctx, id)

	var r0 v1alpha1.ExecutableNodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, string) v1alpha1.ExecutableNodeStatus); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNodeStatus)
		}
	}

	return r0
}

type ExecutableWorkflowStatus_GetOutputReference struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetOutputReference) Return(_a0 storage.DataReference) *ExecutableWorkflowStatus_GetOutputReference {
	return &ExecutableWorkflowStatus_GetOutputReference{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetOutputReference() *ExecutableWorkflowStatus_GetOutputReference {
	c_call := _m.On("GetOutputReference")
	return &ExecutableWorkflowStatus_GetOutputReference{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetOutputReferenceMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetOutputReference {
	c_call := _m.On("GetOutputReference", matchers...)
	return &ExecutableWorkflowStatus_GetOutputReference{Call: c_call}
}

// GetOutputReference provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetOutputReference() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type ExecutableWorkflowStatus_GetPhase struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetPhase) Return(_a0 v1alpha1.WorkflowPhase) *ExecutableWorkflowStatus_GetPhase {
	return &ExecutableWorkflowStatus_GetPhase{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetPhase() *ExecutableWorkflowStatus_GetPhase {
	c_call := _m.On("GetPhase")
	return &ExecutableWorkflowStatus_GetPhase{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetPhaseMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetPhase {
	c_call := _m.On("GetPhase", matchers...)
	return &ExecutableWorkflowStatus_GetPhase{Call: c_call}
}

// GetPhase provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetPhase() v1alpha1.WorkflowPhase {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowPhase
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowPhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowPhase)
	}

	return r0
}

type ExecutableWorkflowStatus_GetStartedAt struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetStartedAt) Return(_a0 *v1.Time) *ExecutableWorkflowStatus_GetStartedAt {
	return &ExecutableWorkflowStatus_GetStartedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetStartedAt() *ExecutableWorkflowStatus_GetStartedAt {
	c_call := _m.On("GetStartedAt")
	return &ExecutableWorkflowStatus_GetStartedAt{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetStartedAtMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetStartedAt {
	c_call := _m.On("GetStartedAt", matchers...)
	return &ExecutableWorkflowStatus_GetStartedAt{Call: c_call}
}

// GetStartedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetStartedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

type ExecutableWorkflowStatus_GetStoppedAt struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_GetStoppedAt) Return(_a0 *v1.Time) *ExecutableWorkflowStatus_GetStoppedAt {
	return &ExecutableWorkflowStatus_GetStoppedAt{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnGetStoppedAt() *ExecutableWorkflowStatus_GetStoppedAt {
	c_call := _m.On("GetStoppedAt")
	return &ExecutableWorkflowStatus_GetStoppedAt{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnGetStoppedAtMatch(matchers ...interface{}) *ExecutableWorkflowStatus_GetStoppedAt {
	c_call := _m.On("GetStoppedAt", matchers...)
	return &ExecutableWorkflowStatus_GetStoppedAt{Call: c_call}
}

// GetStoppedAt provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) GetStoppedAt() *v1.Time {
	ret := _m.Called()

	var r0 *v1.Time
	if rf, ok := ret.Get(0).(func() *v1.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Time)
		}
	}

	return r0
}

// IncFailedAttempts provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) IncFailedAttempts() {
	_m.Called()
}

type ExecutableWorkflowStatus_IsTerminated struct {
	*mock.Call
}

func (_m ExecutableWorkflowStatus_IsTerminated) Return(_a0 bool) *ExecutableWorkflowStatus_IsTerminated {
	return &ExecutableWorkflowStatus_IsTerminated{Call: _m.Call.Return(_a0)}
}

func (_m *ExecutableWorkflowStatus) OnIsTerminated() *ExecutableWorkflowStatus_IsTerminated {
	c_call := _m.On("IsTerminated")
	return &ExecutableWorkflowStatus_IsTerminated{Call: c_call}
}

func (_m *ExecutableWorkflowStatus) OnIsTerminatedMatch(matchers ...interface{}) *ExecutableWorkflowStatus_IsTerminated {
	c_call := _m.On("IsTerminated", matchers...)
	return &ExecutableWorkflowStatus_IsTerminated{Call: c_call}
}

// IsTerminated provides a mock function with given fields:
func (_m *ExecutableWorkflowStatus) IsTerminated() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *ExecutableWorkflowStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetMessage provides a mock function with given fields: msg
func (_m *ExecutableWorkflowStatus) SetMessage(msg string) {
	_m.Called(msg)
}

// SetOutputReference provides a mock function with given fields: reference
func (_m *ExecutableWorkflowStatus) SetOutputReference(reference storage.DataReference) {
	_m.Called(reference)
}

// UpdatePhase provides a mock function with given fields: p, msg, err
func (_m *ExecutableWorkflowStatus) UpdatePhase(p v1alpha1.WorkflowPhase, msg string, err *core.ExecutionError) {
	_m.Called(p, msg, err)
}

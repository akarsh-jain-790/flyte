// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/flyteorg/flyte/flytestdlib/storage"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/flyteorg/flyte/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MutableNodeStatus is an autogenerated mock type for the MutableNodeStatus type
type MutableNodeStatus struct {
	mock.Mock
}

// ClearArrayNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearArrayNodeStatus() {
	_m.Called()
}

// ClearDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearDynamicNodeStatus() {
	_m.Called()
}

// ClearGateNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearGateNodeStatus() {
	_m.Called()
}

// ClearLastAttemptStartedAt provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearLastAttemptStartedAt() {
	_m.Called()
}

// ClearSubNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearSubNodeStatus() {
	_m.Called()
}

// ClearTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearTaskStatus() {
	_m.Called()
}

// ClearWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) ClearWorkflowStatus() {
	_m.Called()
}

type MutableNodeStatus_GetArrayNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetArrayNodeStatus) Return(_a0 v1alpha1.MutableArrayNodeStatus) *MutableNodeStatus_GetArrayNodeStatus {
	return &MutableNodeStatus_GetArrayNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetArrayNodeStatus() *MutableNodeStatus_GetArrayNodeStatus {
	c_call := _m.On("GetArrayNodeStatus")
	return &MutableNodeStatus_GetArrayNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetArrayNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetArrayNodeStatus {
	c_call := _m.On("GetArrayNodeStatus", matchers...)
	return &MutableNodeStatus_GetArrayNodeStatus{Call: c_call}
}

// GetArrayNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetArrayNodeStatus() v1alpha1.MutableArrayNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableArrayNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableArrayNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableArrayNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetBranchStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *MutableNodeStatus_GetBranchStatus {
	return &MutableNodeStatus_GetBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetBranchStatus() *MutableNodeStatus_GetBranchStatus {
	c_call := _m.On("GetBranchStatus")
	return &MutableNodeStatus_GetBranchStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetBranchStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetBranchStatus {
	c_call := _m.On("GetBranchStatus", matchers...)
	return &MutableNodeStatus_GetBranchStatus{Call: c_call}
}

// GetBranchStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetDynamicNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *MutableNodeStatus_GetDynamicNodeStatus {
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetDynamicNodeStatus() *MutableNodeStatus_GetDynamicNodeStatus {
	c_call := _m.On("GetDynamicNodeStatus")
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetDynamicNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetDynamicNodeStatus {
	c_call := _m.On("GetDynamicNodeStatus", matchers...)
	return &MutableNodeStatus_GetDynamicNodeStatus{Call: c_call}
}

// GetDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetGateNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetGateNodeStatus) Return(_a0 v1alpha1.MutableGateNodeStatus) *MutableNodeStatus_GetGateNodeStatus {
	return &MutableNodeStatus_GetGateNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetGateNodeStatus() *MutableNodeStatus_GetGateNodeStatus {
	c_call := _m.On("GetGateNodeStatus")
	return &MutableNodeStatus_GetGateNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetGateNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetGateNodeStatus {
	c_call := _m.On("GetGateNodeStatus", matchers...)
	return &MutableNodeStatus_GetGateNodeStatus{Call: c_call}
}

// GetGateNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetGateNodeStatus() v1alpha1.MutableGateNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableGateNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableGateNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableGateNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateArrayNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateArrayNodeStatus) Return(_a0 v1alpha1.MutableArrayNodeStatus) *MutableNodeStatus_GetOrCreateArrayNodeStatus {
	return &MutableNodeStatus_GetOrCreateArrayNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateArrayNodeStatus() *MutableNodeStatus_GetOrCreateArrayNodeStatus {
	c_call := _m.On("GetOrCreateArrayNodeStatus")
	return &MutableNodeStatus_GetOrCreateArrayNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateArrayNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateArrayNodeStatus {
	c_call := _m.On("GetOrCreateArrayNodeStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateArrayNodeStatus{Call: c_call}
}

// GetOrCreateArrayNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateArrayNodeStatus() v1alpha1.MutableArrayNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableArrayNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableArrayNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableArrayNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateBranchStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateBranchStatus) Return(_a0 v1alpha1.MutableBranchNodeStatus) *MutableNodeStatus_GetOrCreateBranchStatus {
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateBranchStatus() *MutableNodeStatus_GetOrCreateBranchStatus {
	c_call := _m.On("GetOrCreateBranchStatus")
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateBranchStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateBranchStatus {
	c_call := _m.On("GetOrCreateBranchStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateBranchStatus{Call: c_call}
}

// GetOrCreateBranchStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateBranchStatus() v1alpha1.MutableBranchNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableBranchNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableBranchNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableBranchNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateDynamicNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateDynamicNodeStatus) Return(_a0 v1alpha1.MutableDynamicNodeStatus) *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateDynamicNodeStatus() *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c_call := _m.On("GetOrCreateDynamicNodeStatus")
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateDynamicNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateDynamicNodeStatus {
	c_call := _m.On("GetOrCreateDynamicNodeStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateDynamicNodeStatus{Call: c_call}
}

// GetOrCreateDynamicNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateDynamicNodeStatus() v1alpha1.MutableDynamicNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableDynamicNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableDynamicNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableDynamicNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateGateNodeStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateGateNodeStatus) Return(_a0 v1alpha1.MutableGateNodeStatus) *MutableNodeStatus_GetOrCreateGateNodeStatus {
	return &MutableNodeStatus_GetOrCreateGateNodeStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateGateNodeStatus() *MutableNodeStatus_GetOrCreateGateNodeStatus {
	c_call := _m.On("GetOrCreateGateNodeStatus")
	return &MutableNodeStatus_GetOrCreateGateNodeStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateGateNodeStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateGateNodeStatus {
	c_call := _m.On("GetOrCreateGateNodeStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateGateNodeStatus{Call: c_call}
}

// GetOrCreateGateNodeStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateGateNodeStatus() v1alpha1.MutableGateNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableGateNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableGateNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableGateNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateTaskStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *MutableNodeStatus_GetOrCreateTaskStatus {
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateTaskStatus() *MutableNodeStatus_GetOrCreateTaskStatus {
	c_call := _m.On("GetOrCreateTaskStatus")
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateTaskStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateTaskStatus {
	c_call := _m.On("GetOrCreateTaskStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateTaskStatus{Call: c_call}
}

// GetOrCreateTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetOrCreateWorkflowStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetOrCreateWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *MutableNodeStatus_GetOrCreateWorkflowStatus {
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetOrCreateWorkflowStatus() *MutableNodeStatus_GetOrCreateWorkflowStatus {
	c_call := _m.On("GetOrCreateWorkflowStatus")
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetOrCreateWorkflowStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetOrCreateWorkflowStatus {
	c_call := _m.On("GetOrCreateWorkflowStatus", matchers...)
	return &MutableNodeStatus_GetOrCreateWorkflowStatus{Call: c_call}
}

// GetOrCreateWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetOrCreateWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetTaskStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetTaskStatus) Return(_a0 v1alpha1.MutableTaskNodeStatus) *MutableNodeStatus_GetTaskStatus {
	return &MutableNodeStatus_GetTaskStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetTaskStatus() *MutableNodeStatus_GetTaskStatus {
	c_call := _m.On("GetTaskStatus")
	return &MutableNodeStatus_GetTaskStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetTaskStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetTaskStatus {
	c_call := _m.On("GetTaskStatus", matchers...)
	return &MutableNodeStatus_GetTaskStatus{Call: c_call}
}

// GetTaskStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetTaskStatus() v1alpha1.MutableTaskNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableTaskNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableTaskNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableTaskNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_GetWorkflowStatus struct {
	*mock.Call
}

func (_m MutableNodeStatus_GetWorkflowStatus) Return(_a0 v1alpha1.MutableWorkflowNodeStatus) *MutableNodeStatus_GetWorkflowStatus {
	return &MutableNodeStatus_GetWorkflowStatus{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnGetWorkflowStatus() *MutableNodeStatus_GetWorkflowStatus {
	c_call := _m.On("GetWorkflowStatus")
	return &MutableNodeStatus_GetWorkflowStatus{Call: c_call}
}

func (_m *MutableNodeStatus) OnGetWorkflowStatusMatch(matchers ...interface{}) *MutableNodeStatus_GetWorkflowStatus {
	c_call := _m.On("GetWorkflowStatus", matchers...)
	return &MutableNodeStatus_GetWorkflowStatus{Call: c_call}
}

// GetWorkflowStatus provides a mock function with given fields:
func (_m *MutableNodeStatus) GetWorkflowStatus() v1alpha1.MutableWorkflowNodeStatus {
	ret := _m.Called()

	var r0 v1alpha1.MutableWorkflowNodeStatus
	if rf, ok := ret.Get(0).(func() v1alpha1.MutableWorkflowNodeStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.MutableWorkflowNodeStatus)
		}
	}

	return r0
}

type MutableNodeStatus_IncrementAttempts struct {
	*mock.Call
}

func (_m MutableNodeStatus_IncrementAttempts) Return(_a0 uint32) *MutableNodeStatus_IncrementAttempts {
	return &MutableNodeStatus_IncrementAttempts{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnIncrementAttempts() *MutableNodeStatus_IncrementAttempts {
	c_call := _m.On("IncrementAttempts")
	return &MutableNodeStatus_IncrementAttempts{Call: c_call}
}

func (_m *MutableNodeStatus) OnIncrementAttemptsMatch(matchers ...interface{}) *MutableNodeStatus_IncrementAttempts {
	c_call := _m.On("IncrementAttempts", matchers...)
	return &MutableNodeStatus_IncrementAttempts{Call: c_call}
}

// IncrementAttempts provides a mock function with given fields:
func (_m *MutableNodeStatus) IncrementAttempts() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type MutableNodeStatus_IncrementSystemFailures struct {
	*mock.Call
}

func (_m MutableNodeStatus_IncrementSystemFailures) Return(_a0 uint32) *MutableNodeStatus_IncrementSystemFailures {
	return &MutableNodeStatus_IncrementSystemFailures{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnIncrementSystemFailures() *MutableNodeStatus_IncrementSystemFailures {
	c_call := _m.On("IncrementSystemFailures")
	return &MutableNodeStatus_IncrementSystemFailures{Call: c_call}
}

func (_m *MutableNodeStatus) OnIncrementSystemFailuresMatch(matchers ...interface{}) *MutableNodeStatus_IncrementSystemFailures {
	c_call := _m.On("IncrementSystemFailures", matchers...)
	return &MutableNodeStatus_IncrementSystemFailures{Call: c_call}
}

// IncrementSystemFailures provides a mock function with given fields:
func (_m *MutableNodeStatus) IncrementSystemFailures() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type MutableNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableNodeStatus_IsDirty) Return(_a0 bool) *MutableNodeStatus_IsDirty {
	return &MutableNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableNodeStatus) OnIsDirty() *MutableNodeStatus_IsDirty {
	c_call := _m.On("IsDirty")
	return &MutableNodeStatus_IsDirty{Call: c_call}
}

func (_m *MutableNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableNodeStatus_IsDirty {
	c_call := _m.On("IsDirty", matchers...)
	return &MutableNodeStatus_IsDirty{Call: c_call}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ResetDirty provides a mock function with given fields:
func (_m *MutableNodeStatus) ResetDirty() {
	_m.Called()
}

// SetCached provides a mock function with given fields:
func (_m *MutableNodeStatus) SetCached() {
	_m.Called()
}

// SetDataDir provides a mock function with given fields: _a0
func (_m *MutableNodeStatus) SetDataDir(_a0 storage.DataReference) {
	_m.Called(_a0)
}

// SetOutputDir provides a mock function with given fields: d
func (_m *MutableNodeStatus) SetOutputDir(d storage.DataReference) {
	_m.Called(d)
}

// SetParentNodeID provides a mock function with given fields: n
func (_m *MutableNodeStatus) SetParentNodeID(n *string) {
	_m.Called(n)
}

// SetParentTaskID provides a mock function with given fields: t
func (_m *MutableNodeStatus) SetParentTaskID(t *core.TaskExecutionIdentifier) {
	_m.Called(t)
}

// UpdatePhase provides a mock function with given fields: phase, occurredAt, reason, err
func (_m *MutableNodeStatus) UpdatePhase(phase v1alpha1.NodePhase, occurredAt v1.Time, reason string, err *core.ExecutionError) {
	_m.Called(phase, occurredAt, reason, err)
}

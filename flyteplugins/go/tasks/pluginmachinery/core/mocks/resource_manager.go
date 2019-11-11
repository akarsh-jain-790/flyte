// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
	mock "github.com/stretchr/testify/mock"
)

// ResourceManager is an autogenerated mock type for the ResourceManager type
type ResourceManager struct {
	mock.Mock
}

type ResourceManager_AllocateResource struct {
	*mock.Call
}

func (_m ResourceManager_AllocateResource) Return(_a0 core.AllocationStatus, _a1 error) *ResourceManager_AllocateResource {
	return &ResourceManager_AllocateResource{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *ResourceManager) OnAllocateResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) *ResourceManager_AllocateResource {
	c := _m.On("AllocateResource", ctx, namespace, allocationToken)
	return &ResourceManager_AllocateResource{Call: c}
}

func (_m *ResourceManager) OnAllocateResourceMatch(matchers ...interface{}) *ResourceManager_AllocateResource {
	c := _m.On("AllocateResource", matchers...)
	return &ResourceManager_AllocateResource{Call: c}
}

// AllocateResource provides a mock function with given fields: ctx, namespace, allocationToken
func (_m *ResourceManager) AllocateResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) (core.AllocationStatus, error) {
	ret := _m.Called(ctx, namespace, allocationToken)

	var r0 core.AllocationStatus
	if rf, ok := ret.Get(0).(func(context.Context, core.ResourceNamespace, string) core.AllocationStatus); ok {
		r0 = rf(ctx, namespace, allocationToken)
	} else {
		r0 = ret.Get(0).(core.AllocationStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.ResourceNamespace, string) error); ok {
		r1 = rf(ctx, namespace, allocationToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type ResourceManager_ReleaseResource struct {
	*mock.Call
}

func (_m ResourceManager_ReleaseResource) Return(_a0 error) *ResourceManager_ReleaseResource {
	return &ResourceManager_ReleaseResource{Call: _m.Call.Return(_a0)}
}

func (_m *ResourceManager) OnReleaseResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) *ResourceManager_ReleaseResource {
	c := _m.On("ReleaseResource", ctx, namespace, allocationToken)
	return &ResourceManager_ReleaseResource{Call: c}
}

func (_m *ResourceManager) OnReleaseResourceMatch(matchers ...interface{}) *ResourceManager_ReleaseResource {
	c := _m.On("ReleaseResource", matchers...)
	return &ResourceManager_ReleaseResource{Call: c}
}

// ReleaseResource provides a mock function with given fields: ctx, namespace, allocationToken
func (_m *ResourceManager) ReleaseResource(ctx context.Context, namespace core.ResourceNamespace, allocationToken string) error {
	ret := _m.Called(ctx, namespace, allocationToken)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.ResourceNamespace, string) error); ok {
		r0 = rf(ctx, namespace, allocationToken)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

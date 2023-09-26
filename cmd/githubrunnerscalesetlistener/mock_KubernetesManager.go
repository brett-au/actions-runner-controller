// Code generated by mockery v2.33.2. DO NOT EDIT.

package main

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockKubernetesManager is an autogenerated mock type for the KubernetesManager type
type MockKubernetesManager struct {
	mock.Mock
}

// ScaleEphemeralRunnerSet provides a mock function with given fields: ctx, namespace, resourceName, runnerCount
func (_m *MockKubernetesManager) ScaleEphemeralRunnerSet(ctx context.Context, namespace string, resourceName string, runnerCount int) error {
	ret := _m.Called(ctx, namespace, resourceName, runnerCount)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int) error); ok {
		r0 = rf(ctx, namespace, resourceName, runnerCount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateEphemeralRunnerWithJobInfo provides a mock function with given fields: ctx, namespace, resourceName, ownerName, repositoryName, jobWorkflowRef, jobDisplayName, jobRequestId, workflowRunId
func (_m *MockKubernetesManager) UpdateEphemeralRunnerWithJobInfo(ctx context.Context, namespace string, resourceName string, ownerName string, repositoryName string, jobWorkflowRef string, jobDisplayName string, jobRequestId int64, workflowRunId int64) error {
	ret := _m.Called(ctx, namespace, resourceName, ownerName, repositoryName, jobWorkflowRef, jobDisplayName, jobRequestId, workflowRunId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string, string, int64, int64) error); ok {
		r0 = rf(ctx, namespace, resourceName, ownerName, repositoryName, jobWorkflowRef, jobDisplayName, jobRequestId, workflowRunId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockKubernetesManager creates a new instance of MockKubernetesManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKubernetesManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKubernetesManager {
	mock := &MockKubernetesManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
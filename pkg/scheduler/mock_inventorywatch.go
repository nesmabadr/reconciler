// Code generated by mockery 2.7.4. DO NOT EDIT.

package scheduler

import (
	context "context"

	cluster "github.com/kyma-incubator/reconciler/pkg/cluster"

	mock "github.com/stretchr/testify/mock"
)

// MockInventoryWatcher is an autogenerated mock type for the InventoryWatcher type
type MockInventoryWatcher struct {
	mock.Mock
}

// Inventory provides a mock function with given fields:
func (_m *MockInventoryWatcher) Inventory() cluster.Inventory {
	ret := _m.Called()

	var r0 cluster.Inventory
	if rf, ok := ret.Get(0).(func() cluster.Inventory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.Inventory)
		}
	}

	return r0
}

// Run provides a mock function with given fields: ctx, informer
func (_m *MockInventoryWatcher) Run(ctx context.Context, informer InventoryQueue) error {
	ret := _m.Called(ctx, informer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, InventoryQueue) error); ok {
		r0 = rf(ctx, informer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
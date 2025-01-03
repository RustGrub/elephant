// Code generated by mockery v2.46.2. DO NOT EDIT.

package shardedpg

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockShardPicker is an autogenerated mock type for the ShardPicker type
type MockShardPicker struct {
	mock.Mock
}

type MockShardPicker_Expecter struct {
	mock *mock.Mock
}

func (_m *MockShardPicker) EXPECT() *MockShardPicker_Expecter {
	return &MockShardPicker_Expecter{mock: &_m.Mock}
}

// Pick provides a mock function with given fields: ctx, key
func (_m *MockShardPicker) Pick(ctx context.Context, key string) byte {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Pick")
	}

	var r0 byte
	if rf, ok := ret.Get(0).(func(context.Context, string) byte); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(byte)
	}

	return r0
}

// MockShardPicker_Pick_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pick'
type MockShardPicker_Pick_Call struct {
	*mock.Call
}

// Pick is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *MockShardPicker_Expecter) Pick(ctx interface{}, key interface{}) *MockShardPicker_Pick_Call {
	return &MockShardPicker_Pick_Call{Call: _e.mock.On("Pick", ctx, key)}
}

func (_c *MockShardPicker_Pick_Call) Run(run func(ctx context.Context, key string)) *MockShardPicker_Pick_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockShardPicker_Pick_Call) Return(_a0 byte) *MockShardPicker_Pick_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockShardPicker_Pick_Call) RunAndReturn(run func(context.Context, string) byte) *MockShardPicker_Pick_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockShardPicker creates a new instance of MockShardPicker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockShardPicker(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockShardPicker {
	mock := &MockShardPicker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

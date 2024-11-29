// Code generated by mockery v2.40.3. DO NOT EDIT.

package regular

import (
	context "context"

	pgconn "github.com/jackc/pgx/v5/pgconn"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

type MockDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDB) EXPECT() *MockDB_Expecter {
	return &MockDB_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: ctx
func (_m *MockDB) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type MockDB_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDB_Expecter) Begin(ctx interface{}) *MockDB_Begin_Call {
	return &MockDB_Begin_Call{Call: _e.mock.On("Begin", ctx)}
}

func (_c *MockDB_Begin_Call) Run(run func(ctx context.Context)) *MockDB_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDB_Begin_Call) Return(_a0 pgx.Tx, _a1 error) *MockDB_Begin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_Begin_Call) RunAndReturn(run func(context.Context) (pgx.Tx, error)) *MockDB_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Exec provides a mock function with given fields: ctx, query, args
func (_m *MockDB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 pgconn.CommandTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgconn.CommandTag); ok {
		r0 = rf(ctx, query, args...)
	} else {
		r0 = ret.Get(0).(pgconn.CommandTag)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockDB_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockDB_Expecter) Exec(ctx interface{}, query interface{}, args ...interface{}) *MockDB_Exec_Call {
	return &MockDB_Exec_Call{Call: _e.mock.On("Exec",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockDB_Exec_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockDB_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_Exec_Call) Return(_a0 pgconn.CommandTag, _a1 error) *MockDB_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_Exec_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgconn.CommandTag, error)) *MockDB_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, query, args
func (_m *MockDB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 pgx.Rows
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) (pgx.Rows, error)); ok {
		return rf(ctx, query, args...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Rows); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Rows)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type MockDB_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockDB_Expecter) Query(ctx interface{}, query interface{}, args ...interface{}) *MockDB_Query_Call {
	return &MockDB_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockDB_Query_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockDB_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_Query_Call) Return(_a0 pgx.Rows, _a1 error) *MockDB_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_Query_Call) RunAndReturn(run func(context.Context, string, ...interface{}) (pgx.Rows, error)) *MockDB_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryRow provides a mock function with given fields: ctx, query, args
func (_m *MockDB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryRow")
	}

	var r0 pgx.Row
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) pgx.Row); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Row)
		}
	}

	return r0
}

// MockDB_QueryRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryRow'
type MockDB_QueryRow_Call struct {
	*mock.Call
}

// QueryRow is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *MockDB_Expecter) QueryRow(ctx interface{}, query interface{}, args ...interface{}) *MockDB_QueryRow_Call {
	return &MockDB_QueryRow_Call{Call: _e.mock.On("QueryRow",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *MockDB_QueryRow_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *MockDB_QueryRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *MockDB_QueryRow_Call) Return(_a0 pgx.Row) *MockDB_QueryRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDB_QueryRow_Call) RunAndReturn(run func(context.Context, string, ...interface{}) pgx.Row) *MockDB_QueryRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDB creates a new instance of MockDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDB {
	mock := &MockDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
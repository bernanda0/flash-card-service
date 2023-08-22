package sqlc

import (
	"context"
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockDBTX struct {
	mock.Mock
}

func (m *MockDBTX) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	called := m.Called(ctx, query, args)
	return called.Get(0).(sql.Result), called.Error(1)
}

func (m *MockDBTX) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	called := m.Called(ctx, query)
	return called.Get(0).(*sql.Stmt), called.Error(1)
}

func (m *MockDBTX) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	called := m.Called(ctx, query, args)
	return called.Get(0).(*sql.Rows), called.Error(1)
}

func (m *MockDBTX) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	called := m.Called(ctx, query, args)
	return called.Get(0).(*sql.Row)
}

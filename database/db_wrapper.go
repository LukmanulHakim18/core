package database

import (
	"context"
	"database/sql"
	"time"

	"go.elastic.co/apm/module/apmsql/v2"
)

// DBWrapper wraps *sql.DB to add hook
type DBWrapper struct {
	*sql.DB
	hook *metricHook
}

// NewDBClientV2 initializes the database connection with apmsql and prometheus metrics
func NewDBClientV2(host, userName, password, sslMode, port, databaseName string) (*DBWrapper, error) {
	connString := getConnectionString(host, userName, password, sslMode, port, databaseName)
	db, err := apmsql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return &DBWrapper{db, newMetricHook()}, nil
}

// ExecContext overrides sql.DB's ExecContext
func (w *DBWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	startTime := time.Now()

	res, err := w.DB.ExecContext(ctx, query, args...)

	w.hook.AfterProcess(ctx, startTime, err, query, args...)
	return res, err
}

// QueryContext overrides sql.DB's QueryContext
func (w *DBWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	startTime := time.Now()

	res, err := w.DB.QueryContext(ctx, query, args...)

	w.hook.AfterProcess(ctx, startTime, err, query, args...)
	return res, err
}

func (w *DBWrapper) Exec(query string, args ...any) (sql.Result, error) {
	startTime := time.Now()

	res, err := w.DB.Exec(query, args...)

	w.hook.AfterProcess(context.Background(), startTime, err, query, args...)
	return res, err
}

func (w *DBWrapper) Query(query string, args ...any) (*sql.Rows, error) {
	startTime := time.Now()

	res, err := w.DB.Query(query, args...)

	w.hook.AfterProcess(context.Background(), startTime, err, query, args...)
	return res, err

}

func (w *DBWrapper) QueryRow(query string, args ...any) *sql.Row {
	startTime := time.Now()

	w.hook.AfterProcess(context.Background(), startTime, nil, query, args...)
	res := w.DB.QueryRow(query, args...)

	return res
}

func (w *DBWrapper) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	startTime := time.Now()

	w.hook.AfterProcess(context.Background(), startTime, nil, query, args...)
	res := w.DB.QueryRowContext(ctx, query, args...)

	return res
}

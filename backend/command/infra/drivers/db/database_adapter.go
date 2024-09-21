package db

import "context"

type DB interface {
    Query(ctx context.Context, query string, args ...interface{}) (Rows, error)
    Exec(ctx context.Context, query string, args ...interface{}) (Result, error)
    Close() error
}

type Rows interface {
    Next() bool
    Scan(dest ...interface{}) error
    Close() error
}

type Result interface {
    RowsAffected() (int64, error)
}
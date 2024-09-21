package db

import (
    "context"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/jackc/pgconn"
)

type PgxAdapter struct {
    pool *pgxpool.Pool
}

func NewPgxAdapter(connString string) (*PgxAdapter, error) {
    pool, err := pgxpool.Connect(context.Background(), connString)
    if err != nil {
        return nil, err
    }

    return &PgxAdapter{
        pool: pool,
    }, nil
}

func (p *PgxAdapter) Query(ctx context.Context, query string, args ...interface{}) (Rows, error) {
    rows, err := p.pool.Query(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    return &PgxRows{rows: rows}, nil
}

func (p *PgxAdapter) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
    result, err := p.pool.Exec(ctx, query, args...)
    if err != nil {
        return nil, err
    }
    return &PgxResult{result}, nil
}

func (p *PgxAdapter) Close() error {
    p.pool.Close()
    return nil
}

type PgxRows struct {
    rows pgx.Rows
}

func (r *PgxRows) Next() bool {
    return r.rows.Next()
}

func (r *PgxRows) Scan(dest ...interface{}) error {
    return r.rows.Scan(dest...)
}

func (r *PgxRows) Close() error {
    r.rows.Close()
    return nil
}

type PgxResult struct {
    commandTag pgconn.CommandTag
}

func (r *PgxResult) RowsAffected() (int64, error) {
    return r.commandTag.RowsAffected(), nil
}

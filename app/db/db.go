package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct {
	DB *sql.DB
}

func New(db *sql.DB) *Queries {
	return &Queries{DB: db}
}

func (q *Queries) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := q.DB.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	nq := New(q.DB)
	err = fn(nq)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

package store

import (
	"database/sql"

	appdb "github.com/like-org/common/app/db"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*appdb.Queries
	db *sql.DB
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: appdb.New(db),
	}
}

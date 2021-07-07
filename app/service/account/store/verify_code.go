package store

import (
	"context"
	"time"

	"github.com/like-org/common/app/service/account/model"
)

type ArgCreateVerifyCode struct {
	Key       string
	Code      int
	ExpiredAt time.Time
}

const createVerifyCode = `
INSERT INTO verify_codes (
    key,
    code,
    expired_at
) VALUES (
    $1, $2, $3
) RETURNING id, key, code, expired_at, created_at
`

func (store *SQLStore) CreateVerifyCode(ctx context.Context, arg *ArgCreateVerifyCode) (*model.VerifyCode, error) {
	row := store.DB.QueryRowContext(ctx, createVerifyCode, arg.Key, arg.Code, arg.ExpiredAt)

	var i model.VerifyCode
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Code,
		&i.ExpiredAt,
		&i.CreatedAt,
	)

	return &i, err
}

const getVerifyCode = `
SELECT * FROM verify_codes 
WHERE key = $1 ORDER BY created_at DESC LIMIT 1
`

func (store *SQLStore) GetVerifyCode(ctx context.Context, key string) (*model.VerifyCode, error) {
	row := store.DB.QueryRowContext(ctx, getVerifyCode, key)

	var i model.VerifyCode
	err := row.Scan(
		&i.ID,
		&i.Key,
		&i.Code,
		&i.ExpiredAt,
		&i.CreatedAt,
	)

	return &i, err
}

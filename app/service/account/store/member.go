package store

import (
	"context"

	"github.com/like-org/common/app/service/account/model"
)

const createMember = `
INSERT INTO members (
    mobile,
    password,
    avatar
) VALUES (
    $1, $2, $3
) RETURNING id, mobile, avatar, status, created_at
`

type ArgCreateMember struct {
	Mobile   string
	Password string
	Avatar   string
}

func (store *SQLStore) CreateMember(ctx context.Context, arg *ArgCreateMember) (*model.Member, error) {
	row := store.DB.QueryRowContext(ctx, createMember, arg.Mobile, arg.Password, arg.Avatar)

	var i model.Member
	err := row.Scan(
		&i.ID,
		&i.Mobile,
		&i.Avatar,
		&i.Status,
		&i.CreatedAt,
	)

	return &i, err
}

const getMemberByID = `
SELECT id, mobile, username, avatar, status, password, created_at FROM members 
WHERE id = $1 LIMIT 1
`

func (store *SQLStore) GetMemberByID(ctx context.Context, ID int) (*model.Member, error) {
	row := store.DB.QueryRowContext(ctx, getMemberByID, ID)

	var i model.Member
	err := row.Scan(
		&i.ID,
		&i.Mobile,
		&i.Username,
		&i.Avatar,
		&i.Status,
		&i.Password,
		&i.CreatedAt,
	)

	return &i, err
}

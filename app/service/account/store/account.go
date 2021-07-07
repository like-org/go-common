package store

import (
	"context"

	"github.com/like-org/common/app/db"
	"github.com/like-org/common/app/service/account/model"
)

const createAccount = `
INSERT INTO accounts (
    amount,
	type
) VALUES (
    $1, $2
) RETURNING id, amount, type, status, created_at
`

type ArgCreateAccount struct {
	Amount int
	Type   int
}

func (store *SQLStore) CreateAccount(ctx context.Context, arg *ArgCreateAccount) (*model.Account, error) {

	row := store.DB.QueryRowContext(ctx, createAccount, arg.Amount, arg.Type)

	var i model.Account
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.Type,
		&i.Status,
		&i.CreatedAt,
	)

	return &i, err
}

const createAccountRef = `
INSERT INTO acc_mb_ref (
    acc_id,
	mb_id
) VALUES (
    $1, $2
) RETURNING id
`

type ArgCreateAccountRef struct {
	AccountID int
	MemberID  int
}

func (q *SQLStore) CreateAccountRef(ctx context.Context, arg *ArgCreateAccountRef) error {

	row := q.DB.QueryRowContext(ctx, createAccountRef, arg.AccountID, arg.MemberID)

	var i int
	err := row.Scan(
		&i,
	)

	return err
}

type ArgCreateAccountTx struct {
	Amount   int
	Type     int
	MemberID int
}

func (store *SQLStore) CreateAccountTx(ctx context.Context, arg *ArgCreateAccountTx) (*model.Account, error) {
	var acc model.Account
	err := store.ExecTx(ctx, func(q *db.Queries) error {
		acc, err := store.CreateAccount(ctx, &ArgCreateAccount{Amount: arg.Amount, Type: arg.Type})
		if err != nil {
			return err
		}
		err = store.CreateAccountRef(ctx, &ArgCreateAccountRef{AccountID: acc.ID, MemberID: arg.MemberID})
		if err != nil {
			return err
		}
		return nil
	})
	return &acc, err
}

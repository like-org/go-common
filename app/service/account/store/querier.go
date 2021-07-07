package store

import (
	"context"

	"github.com/like-org/common/app/service/account/model"
)

type Querier interface {
	CreateMember(ctx context.Context, arg *ArgCreateMember) (*model.Member, error)
	CreateVerifyCode(ctx context.Context, arg *ArgCreateVerifyCode) (*model.VerifyCode, error)
	GetVerifyCode(ctx context.Context, key string) (*model.VerifyCode, error)
	GetMemberByID(ctx context.Context, ID int) (*model.Member, error)
	CreateAccountTx(ctx context.Context, arg *ArgCreateAccountTx) (*model.Account, error)
}

var _ Querier = (*SQLStore)(nil)

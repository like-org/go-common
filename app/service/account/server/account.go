package service

import (
	"context"

	"github.com/like-org/common/app/service/account/pb"
	db "github.com/like-org/common/app/service/account/store"
)

func (srv *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountReq) (*pb.CreateAccountResp, error) {
	arg := &db.ArgCreateAccountTx{
		MemberID: int(req.GetMemberId()),
		Amount:   int(req.GetAmount()),
		Type:     int(req.GetType()),
	}

	acc, err := srv.store.CreateAccountTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResp{Account: &pb.Account{
		Id:        uint64(acc.ID),
		Amount:    int64(acc.Amount),
		Type:      int32(acc.Type),
		Status:    int32(acc.Status),
		UpdatedAt: acc.UpdatedAt.Format("2006-01-02 15:04:05"),
		CreatedAt: acc.CreatedAt.Format("2006-01-02 15:04:05"),
	}}, nil
}

package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/like-org/common/app/service/account/pb"
	db "github.com/like-org/common/app/service/account/store"
)

func (srv *Server) CreateMemberByMobile(ctx context.Context, req *pb.CreateMemberByMobileReq) (*pb.CreateMemberByMobileResp, error) {

	arg := &db.ArgCreateMember{
		Mobile:   req.GetMobile(),
		Avatar:   req.GetAvatar(),
		Password: req.GetPassword(),
	}
	mb, err := srv.store.CreateMember(ctx, arg)
	if err != nil {
		if strings.Contains(err.Error(), "pq: duplicate key") {
			return nil, fmt.Errorf("手机号码已存在")
		}
		return nil, err
	}

	resp := &pb.CreateMemberByMobileResp{
		Member: &pb.Member{
			Id:        uint64(mb.ID),
			Mobile:    mb.Mobile,
			CreatedAt: mb.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	_, err = srv.CreateAccount(ctx, &pb.CreateAccountReq{
		Amount:   0,
		Type:     0,
		MemberId: int64(mb.ID),
	})
	if err != nil {
		log.Printf("CreateMemberByMobile CreateAccount background failure: %v", err)
	}

	return resp, nil
}

func (srv *Server) GetMemeberById(ctx context.Context, req *pb.GetMemeberByIdReq) (*pb.GetMemeberByIdResp, error) {
	log.Print("Account service --> member GetMemberById call")
	mb, err := srv.store.GetMemberByID(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	mm := &pb.Member{
		Id:        uint64(mb.ID),
		Mobile:    mb.Mobile,
		Username:  mb.Username,
		CreatedAt: mb.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	return &pb.GetMemeberByIdResp{
		Member: mm,
	}, nil
}

package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/like-org/common/app/service/account/pb"
	db "github.com/like-org/common/app/service/account/store"

	"github.com/like-org/common/lib/util"
)

const (
	codeDuration  = time.Minute * 5
	defaultAvatar = "/default.jpg"
)

func (srv *Server) SendVerifyCode(ctx context.Context, req *pb.SendVerifyCodeReq) (*pb.SendVerifyCodeResp, error) {
	key := req.GetKey()
	code := util.RandomInt(1000, 9999)

	arg := &db.ArgCreateVerifyCode{
		Key:       key,
		Code:      code,
		ExpiredAt: time.Now().Add(codeDuration),
	}

	_, err := srv.store.CreateVerifyCode(ctx, arg)
	if err != nil {
		return nil, err
	}

	if strings.Contains(key, "@") { // email
		// send email code
	} else {
		// send mobile code
	}

	return &pb.SendVerifyCodeResp{Ok: true}, nil
}

func (srv *Server) RegisterByMobile(ctx context.Context, req *pb.RegisterByMobileReq) (*pb.RegisterByMobileResp, error) {
	mobile := req.GetMobile()
	code := req.GetCode()

	verifyCode, err := srv.store.GetVerifyCode(ctx, mobile)
	if err != nil {
		log.Printf("RegisterMyMobile GetVerifyCode Error: %v", err)
		return nil, err
	}

	if verifyCode.Code != int(code) ||
		verifyCode.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("无效的验证码")
	}

	hashPwd, err := util.HashPassword(util.RandomPassword())
	if err != nil {
		return nil, err
	}

	resp, err := srv.CreateMemberByMobile(ctx, &pb.CreateMemberByMobileReq{
		Mobile:   mobile,
		Password: hashPwd,
		Avatar:   defaultAvatar,
	})

	if err != nil {
		return nil, err
	}

	token := db.GenerateToken(int(resp.GetMember().GetId()), hashPwd)

	return &pb.RegisterByMobileResp{
		Member: resp.GetMember(),
		Token:  token,
	}, nil
}

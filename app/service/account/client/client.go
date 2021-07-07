package client

import (
	"context"
	"fmt"
	"log"

	"github.com/like-org/common/app/service/account/model"
	"github.com/like-org/common/app/service/account/pb"
	"google.golang.org/grpc"
)

type AccountClient struct {
	server pb.ServiceClient
}

func NewAccountClient(cc *grpc.ClientConn) *AccountClient {
	server := pb.NewServiceClient(cc)
	return &AccountClient{server}
}

func (a *AccountClient) GetMemberById(ctx context.Context, uid int) (*model.Member, error) {
	log.Print("Account client --> member GetMemberById call")
	log.Print(uid)
	resp, err := a.server.GetMemeberById(ctx, &pb.GetMemeberByIdReq{
		Id: int64(uid),
	})
	if err != nil {
		log.Printf("Account client --> member GetMemberById call %v", err)
		return nil, err
	}
	m := &model.Member{
		ID:       int(resp.GetMember().GetId()),
		Username: resp.GetMember().Username,
	}
	return m, nil
}

func (a *AccountClient) RegisterByMobile(ctx context.Context, mobile string, code int) (*model.Member, error) {
	req := &pb.RegisterByMobileReq{
		Mobile: mobile,
		Code:   uint32(code),
	}
	resp, err := a.server.RegisterByMobile(ctx, req)
	if err != nil {
		return nil, err
	}
	member := resp.GetMember()

	return &model.Member{
		ID:       int(member.GetId()),
		Username: resp.GetMember().Username,
	}, nil
}

func (a *AccountClient) SendVerifyCode(ctx context.Context, key string) error {
	req := &pb.SendVerifyCodeReq{
		Key: key,
	}
	resp, err := a.server.SendVerifyCode(ctx, req)
	if err != nil {
		return err
	}

	if !resp.GetOk() {
		return fmt.Errorf("send verify code failure")
	}
	return nil
}

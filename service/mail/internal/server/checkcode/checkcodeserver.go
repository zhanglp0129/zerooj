// Code generated by goctl. DO NOT EDIT.
// Source: mail.proto

package server

import (
	"context"

	"zerooj/service/mail/internal/logic/checkcode"
	"zerooj/service/mail/internal/svc"
	"zerooj/service/mail/pb/user"
)

type CheckCodeServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedCheckCodeServer
}

func NewCheckCodeServer(svcCtx *svc.ServiceContext) *CheckCodeServer {
	return &CheckCodeServer{
		svcCtx: svcCtx,
	}
}

// 发送邮箱验证码
func (s *CheckCodeServer) SendMailCheckCode(ctx context.Context, in *user.SendMailCheckCodeReq) (*user.SendMailCheckCodeResp, error) {
	l := checkcodelogic.NewSendMailCheckCodeLogic(ctx, s.svcCtx)
	return l.SendMailCheckCode(in)
}

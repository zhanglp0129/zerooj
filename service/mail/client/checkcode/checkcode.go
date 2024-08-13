// Code generated by goctl. DO NOT EDIT.
// Source: mail.proto

package checkcode

import (
	"context"

	"zerooj/service/mail/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendMailCheckCodeReq  = user.SendMailCheckCodeReq
	SendMailCheckCodeResp = user.SendMailCheckCodeResp

	CheckCode interface {
		// 发送邮箱验证码
		SendMailCheckCode(ctx context.Context, in *SendMailCheckCodeReq, opts ...grpc.CallOption) (*SendMailCheckCodeResp, error)
	}

	defaultCheckCode struct {
		cli zrpc.Client
	}
)

func NewCheckCode(cli zrpc.Client) CheckCode {
	return &defaultCheckCode{
		cli: cli,
	}
}

// 发送邮箱验证码
func (m *defaultCheckCode) SendMailCheckCode(ctx context.Context, in *SendMailCheckCodeReq, opts ...grpc.CallOption) (*SendMailCheckCodeResp, error) {
	client := user.NewCheckCodeClient(m.cli.Conn())
	return client.SendMailCheckCode(ctx, in, opts...)
}

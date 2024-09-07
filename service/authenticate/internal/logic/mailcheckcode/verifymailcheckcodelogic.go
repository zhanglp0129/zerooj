package mailcheckcodelogic

import (
	"context"
	"zerooj/common/constant"

	"zerooj/service/authenticate/internal/svc"
	"zerooj/service/authenticate/pb/authenticate"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyMailCheckCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyMailCheckCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyMailCheckCodeLogic {
	return &VerifyMailCheckCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验
func (l *VerifyMailCheckCodeLogic) VerifyMailCheckCode(in *authenticate.VerifyMailCheckCodeReq) (*authenticate.VerifyMailCheckCodeResp, error) {
	rdb := l.svcCtx.RDB
	key := in.RedisKey
	trueCheckCode, err := rdb.Get(context.Background(), key).Result()
	if err != nil || in.CheckCode != trueCheckCode {
		return nil, constant.MailCheckCodeError
	}

	return &authenticate.VerifyMailCheckCodeResp{}, nil
}

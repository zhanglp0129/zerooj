package mailcheckcodelogic

import (
	"context"
	"time"

	"zerooj/service/authenticate/internal/svc"
	"zerooj/service/authenticate/pb/authenticate"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishMailCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFinishMailCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishMailCheckLogic {
	return &FinishMailCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱验证码校验完毕
func (l *FinishMailCheckLogic) FinishMailCheck(in *authenticate.FinishMailCheckReq) (*authenticate.FinishMailCheckResp, error) {
	rdb := l.svcCtx.RDB
	key := in.RedisKey
	// 删除邮箱验证码，不需要很高的一致性
	go func() {
		// 失败后重试
		for i := 0; i < 3; i++ {
			err := rdb.Del(context.Background(), key).Err()
			if err == nil {
				break
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return &authenticate.FinishMailCheckResp{}, nil
}

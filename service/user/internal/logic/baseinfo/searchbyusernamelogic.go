package baseinfologic

import (
	"context"

	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchByUsernameLogic {
	return &SearchByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据用户名搜索用户，并缓存
func (l *SearchByUsernameLogic) SearchByUsername(in *user.SearchByUsernameReq) (*user.SearchByUsernameResp, error) {
	// todo: add your logic here and delete this line

	return &user.SearchByUsernameResp{}, nil
}

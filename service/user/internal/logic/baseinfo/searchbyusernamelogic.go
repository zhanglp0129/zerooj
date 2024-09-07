package baseinfologic

import (
	"context"
	"zerooj/service/user/models"

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

// 根据用户名搜索用户
func (l *SearchByUsernameLogic) SearchByUsername(in *user.SearchByUsernameReq) (*user.SearchByUsernameResp, error) {
	db := l.svcCtx.DB
	u := models.User{}
	err := db.Select("id").Where("username = ?", in.Username).Take(&u).Error
	if err != nil {
		return nil, err
	}

	return &user.SearchByUsernameResp{
		Id: u.ID,
	}, nil
}

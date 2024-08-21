package profilelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhanglp0129/redis_cache"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/models"
	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户简介
func (l *GetUserProfileLogic) GetUserProfile(in *user.GetUserProfileReq) (*user.GetUserProfileResp, error) {
	// 带着缓存查询
	rdb := l.svcCtx.RDB
	key := fmt.Sprintf("/cache/user/get_user_profile/%d", in.Id)
	var model user.GetUserProfileResp
	_, err := redis_cache.QueryWithCache(rdb, key, &model, func() (*user.GetUserProfileResp, error) {
		// 先查找用户简介表
		var p models.UserProfile
		db := l.svcCtx.DB
		err := db.Where("user_id = ?", in.Id).Take(&p).Error

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没有用户简介，需要插入默认数据
			p = models.UserProfile{
				UserID:   in.Id,
				Birthday: time.Now(),
			}
			err = db.Create(&p).Error
			if err != nil {
				return nil, err
			}
			model.Birthday = timestamppb.New(p.Birthday)
		} else {
			// 有用户简介，查询其他信息
			u := models.User{}
			err = db.Preload("Websites").Preload("Skills").Take(&u, in.Id).Error
			if err != nil {
				return nil, err
			}
			// 将查询到的数据赋给model
			model = user.GetUserProfileResp{
				Nickname:             p.Nickname,
				AvatarURL:            p.AvatarURL,
				Gender:               uint32(p.Gender),
				Birthday:             timestamppb.New(p.Birthday),
				PersonalIntroduction: p.PersonalIntroduction,
			}
			for _, website := range u.Websites {
				model.Websites = append(model.Websites, &user.PersonalWebsite{
					Id:   website.ID,
					Name: website.Name,
					Url:  website.URL,
				})
			}
			for _, skill := range u.Skills {
				model.Skills = append(model.Skills, &user.Skill{
					Id:   skill.ID,
					Name: skill.Name,
				})
			}
		}

		return &model, nil
	})
	if err != nil {
		return nil, err
	}

	return &model, nil
}

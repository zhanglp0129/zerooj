// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package follow

import (
	"context"

	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCityReq                    = user.AddCityReq
	AddCityResp                   = user.AddCityResp
	AddSkillReq                   = user.AddSkillReq
	AddSkillResp                  = user.AddSkillResp
	AddUserPersonalWebsiteReq     = user.AddUserPersonalWebsiteReq
	AddUserPersonalWebsiteResp    = user.AddUserPersonalWebsiteResp
	AddUserSkillReq               = user.AddUserSkillReq
	AddUserSkillResp              = user.AddUserSkillResp
	City                          = user.City
	DeleteCityReq                 = user.DeleteCityReq
	DeleteCityResp                = user.DeleteCityResp
	DeleteSkillReq                = user.DeleteSkillReq
	DeleteSkillResp               = user.DeleteSkillResp
	DeleteUserPersonalWebsiteReq  = user.DeleteUserPersonalWebsiteReq
	DeleteUserPersonalWebsiteResp = user.DeleteUserPersonalWebsiteResp
	DeleteUserSkillReq            = user.DeleteUserSkillReq
	DeleteUserSkillResp           = user.DeleteUserSkillResp
	FollowUserInfo                = user.FollowUserInfo
	FollowUserReq                 = user.FollowUserReq
	FollowUserResp                = user.FollowUserResp
	ForgetPasswordReq             = user.ForgetPasswordReq
	ForgetPasswordResp            = user.ForgetPasswordResp
	GetBaseInfoReq                = user.GetBaseInfoReq
	GetBaseInfoResp               = user.GetBaseInfoResp
	GetFansReq                    = user.GetFansReq
	GetFansResp                   = user.GetFansResp
	GetFollowingsReq              = user.GetFollowingsReq
	GetFollowingsResp             = user.GetFollowingsResp
	GetPermissionReq              = user.GetPermissionReq
	GetPermissionResp             = user.GetPermissionResp
	GetUserInfoReq                = user.GetUserInfoReq
	GetUserInfoResp               = user.GetUserInfoResp
	MustDeleteCityReq             = user.MustDeleteCityReq
	MustDeleteCityResp            = user.MustDeleteCityResp
	MustDeleteSkillReq            = user.MustDeleteSkillReq
	MustDeleteSkillResp           = user.MustDeleteSkillResp
	PersonalWebsite               = user.PersonalWebsite
	Skill                         = user.Skill
	UpdateEmailReq                = user.UpdateEmailReq
	UpdateEmailResp               = user.UpdateEmailResp
	UpdatePasswordReq             = user.UpdatePasswordReq
	UpdatePasswordResp            = user.UpdatePasswordResp
	UpdatePermissionReq           = user.UpdatePermissionReq
	UpdatePermissionResp          = user.UpdatePermissionResp
	UpdateUserProfileReq          = user.UpdateUserProfileReq
	UpdateUserProfileResp         = user.UpdateUserProfileResp
	UpdateUsernameReq             = user.UpdateUsernameReq
	UpdateUsernameResp            = user.UpdateUsernameResp
	UserLoginReq                  = user.UserLoginReq
	UserLoginResp                 = user.UserLoginResp
	UserRegisterReq               = user.UserRegisterReq
	UserRegisterResp              = user.UserRegisterResp

	Follow interface {
		// 获取所有关注，分页查询
		GetFollowings(ctx context.Context, in *GetFollowingsReq, opts ...grpc.CallOption) (*GetFollowingsResp, error)
		// 获取所有粉丝，分页查询
		GetFans(ctx context.Context, in *GetFansReq, opts ...grpc.CallOption) (*GetFansResp, error)
		// 关注其他用户
		FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*FollowUserResp, error)
	}

	defaultFollow struct {
		cli zrpc.Client
	}
)

func NewFollow(cli zrpc.Client) Follow {
	return &defaultFollow{
		cli: cli,
	}
}

// 获取所有关注，分页查询
func (m *defaultFollow) GetFollowings(ctx context.Context, in *GetFollowingsReq, opts ...grpc.CallOption) (*GetFollowingsResp, error) {
	client := user.NewFollowClient(m.cli.Conn())
	return client.GetFollowings(ctx, in, opts...)
}

// 获取所有粉丝，分页查询
func (m *defaultFollow) GetFans(ctx context.Context, in *GetFansReq, opts ...grpc.CallOption) (*GetFansResp, error) {
	client := user.NewFollowClient(m.cli.Conn())
	return client.GetFans(ctx, in, opts...)
}

// 关注其他用户
func (m *defaultFollow) FollowUser(ctx context.Context, in *FollowUserReq, opts ...grpc.CallOption) (*FollowUserResp, error) {
	client := user.NewFollowClient(m.cli.Conn())
	return client.FollowUser(ctx, in, opts...)
}

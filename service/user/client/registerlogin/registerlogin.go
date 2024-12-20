// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package registerlogin

import (
	"context"

	"zerooj/service/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddSkillReq                   = user.AddSkillReq
	AddSkillResp                  = user.AddSkillResp
	AddUserPersonalWebsiteReq     = user.AddUserPersonalWebsiteReq
	AddUserPersonalWebsiteResp    = user.AddUserPersonalWebsiteResp
	AddUserSkillReq               = user.AddUserSkillReq
	AddUserSkillResp              = user.AddUserSkillResp
	DeleteSkillReq                = user.DeleteSkillReq
	DeleteSkillResp               = user.DeleteSkillResp
	DeleteUserPersonalWebsiteReq  = user.DeleteUserPersonalWebsiteReq
	DeleteUserPersonalWebsiteResp = user.DeleteUserPersonalWebsiteResp
	DeleteUserSkillReq            = user.DeleteUserSkillReq
	DeleteUserSkillResp           = user.DeleteUserSkillResp
	FollowUserInfo                = user.FollowUserInfo
	FollowUserReq                 = user.FollowUserReq
	FollowUserResp                = user.FollowUserResp
	GetBaseInfoReq                = user.GetBaseInfoReq
	GetBaseInfoResp               = user.GetBaseInfoResp
	GetFansReq                    = user.GetFansReq
	GetFansResp                   = user.GetFansResp
	GetFollowingsReq              = user.GetFollowingsReq
	GetFollowingsResp             = user.GetFollowingsResp
	GetPermissionReq              = user.GetPermissionReq
	GetPermissionResp             = user.GetPermissionResp
	GetUserProfileReq             = user.GetUserProfileReq
	GetUserProfileResp            = user.GetUserProfileResp
	MustDeleteSkillReq            = user.MustDeleteSkillReq
	MustDeleteSkillResp           = user.MustDeleteSkillResp
	PersonalWebsite               = user.PersonalWebsite
	SearchByUsernameReq           = user.SearchByUsernameReq
	SearchByUsernameResp          = user.SearchByUsernameResp
	SearchSkillsReq               = user.SearchSkillsReq
	SearchSkillsResp              = user.SearchSkillsResp
	SkillInfo                     = user.SkillInfo
	UnfollowUserReq               = user.UnfollowUserReq
	UnfollowUserResp              = user.UnfollowUserResp
	UpdateBaseInfoReq             = user.UpdateBaseInfoReq
	UpdateBaseInfoResp            = user.UpdateBaseInfoResp
	UpdatePermissionReq           = user.UpdatePermissionReq
	UpdatePermissionResp          = user.UpdatePermissionResp
	UpdateSkillReq                = user.UpdateSkillReq
	UpdateSkillResp               = user.UpdateSkillResp
	UpdateUserProfileReq          = user.UpdateUserProfileReq
	UpdateUserProfileResp         = user.UpdateUserProfileResp
	UserLoginReq                  = user.UserLoginReq
	UserLoginResp                 = user.UserLoginResp
	UserRegisterReq               = user.UserRegisterReq
	UserRegisterResp              = user.UserRegisterResp

	RegisterLogin interface {
		// 用户注册
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
		// 用户登录
		UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	}

	defaultRegisterLogin struct {
		cli zrpc.Client
	}
)

func NewRegisterLogin(cli zrpc.Client) RegisterLogin {
	return &defaultRegisterLogin{
		cli: cli,
	}
}

// 用户注册
func (m *defaultRegisterLogin) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	client := user.NewRegisterLoginClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}

// 用户登录
func (m *defaultRegisterLogin) UserLogin(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := user.NewRegisterLoginClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

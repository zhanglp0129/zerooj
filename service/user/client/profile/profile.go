// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package profile

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
	ForgetPasswordReq             = user.ForgetPasswordReq
	ForgetPasswordResp            = user.ForgetPasswordResp
	GetAllSkillsReq               = user.GetAllSkillsReq
	GetAllSkillsResp              = user.GetAllSkillsResp
	GetBaseInfoReq                = user.GetBaseInfoReq
	GetBaseInfoResp               = user.GetBaseInfoResp
	GetFansReq                    = user.GetFansReq
	GetFansResp                   = user.GetFansResp
	GetFollowingsReq              = user.GetFollowingsReq
	GetFollowingsResp             = user.GetFollowingsResp
	GetUserProfileReq             = user.GetUserProfileReq
	GetUserProfileResp            = user.GetUserProfileResp
	MustDeleteSkillReq            = user.MustDeleteSkillReq
	MustDeleteSkillResp           = user.MustDeleteSkillResp
	PersonalWebsite               = user.PersonalWebsite
	SearchByUsernameReq           = user.SearchByUsernameReq
	SearchByUsernameResp          = user.SearchByUsernameResp
	Skill                         = user.Skill
	UnfollowUserReq               = user.UnfollowUserReq
	UnfollowUserResp              = user.UnfollowUserResp
	UpdateEmailReq                = user.UpdateEmailReq
	UpdateEmailResp               = user.UpdateEmailResp
	UpdatePasswordReq             = user.UpdatePasswordReq
	UpdatePasswordResp            = user.UpdatePasswordResp
	UpdatePermissionReq           = user.UpdatePermissionReq
	UpdatePermissionResp          = user.UpdatePermissionResp
	UpdateSkillReq                = user.UpdateSkillReq
	UpdateSkillResp               = user.UpdateSkillResp
	UpdateUserProfileReq          = user.UpdateUserProfileReq
	UpdateUserProfileResp         = user.UpdateUserProfileResp
	UpdateUsernameReq             = user.UpdateUsernameReq
	UpdateUsernameResp            = user.UpdateUsernameResp
	UserLoginReq                  = user.UserLoginReq
	UserLoginResp                 = user.UserLoginResp
	UserRegisterReq               = user.UserRegisterReq
	UserRegisterResp              = user.UserRegisterResp

	Profile interface {
		// 获取用户简介
		GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileResp, error)
		// 修改用户简介，不包括个人网站和用户技能
		UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileResp, error)
		// 添加个人网站，最多5个
		AddUserPersonalWebsite(ctx context.Context, in *AddUserPersonalWebsiteReq, opts ...grpc.CallOption) (*AddUserPersonalWebsiteResp, error)
		// 删除个人网站
		DeleteUserPersonalWebsite(ctx context.Context, in *DeleteUserPersonalWebsiteReq, opts ...grpc.CallOption) (*DeleteUserPersonalWebsiteResp, error)
		// 添加用户技能，最多10个
		AddUserSkill(ctx context.Context, in *AddUserSkillReq, opts ...grpc.CallOption) (*AddUserSkillResp, error)
		// 删除用户技能
		DeleteUserSkill(ctx context.Context, in *DeleteUserSkillReq, opts ...grpc.CallOption) (*DeleteUserSkillResp, error)
	}

	defaultProfile struct {
		cli zrpc.Client
	}
)

func NewProfile(cli zrpc.Client) Profile {
	return &defaultProfile{
		cli: cli,
	}
}

// 获取用户简介
func (m *defaultProfile) GetUserProfile(ctx context.Context, in *GetUserProfileReq, opts ...grpc.CallOption) (*GetUserProfileResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.GetUserProfile(ctx, in, opts...)
}

// 修改用户简介，不包括个人网站和用户技能
func (m *defaultProfile) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...grpc.CallOption) (*UpdateUserProfileResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.UpdateUserProfile(ctx, in, opts...)
}

// 添加个人网站，最多5个
func (m *defaultProfile) AddUserPersonalWebsite(ctx context.Context, in *AddUserPersonalWebsiteReq, opts ...grpc.CallOption) (*AddUserPersonalWebsiteResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.AddUserPersonalWebsite(ctx, in, opts...)
}

// 删除个人网站
func (m *defaultProfile) DeleteUserPersonalWebsite(ctx context.Context, in *DeleteUserPersonalWebsiteReq, opts ...grpc.CallOption) (*DeleteUserPersonalWebsiteResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.DeleteUserPersonalWebsite(ctx, in, opts...)
}

// 添加用户技能，最多10个
func (m *defaultProfile) AddUserSkill(ctx context.Context, in *AddUserSkillReq, opts ...grpc.CallOption) (*AddUserSkillResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.AddUserSkill(ctx, in, opts...)
}

// 删除用户技能
func (m *defaultProfile) DeleteUserSkill(ctx context.Context, in *DeleteUserSkillReq, opts ...grpc.CallOption) (*DeleteUserSkillResp, error) {
	client := user.NewProfileClient(m.cli.Conn())
	return client.DeleteUserSkill(ctx, in, opts...)
}

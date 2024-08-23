// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package other

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
	SearchSkillsReq               = user.SearchSkillsReq
	SearchSkillsResp              = user.SearchSkillsResp
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

	Other interface {
		// 添加技能，需要客服权限
		AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*AddSkillResp, error)
		// 修改技能，需要客服权限
		UpdateSkill(ctx context.Context, in *UpdateSkillReq, opts ...grpc.CallOption) (*UpdateSkillResp, error)
		// 搜索技能
		SearchSkills(ctx context.Context, in *SearchSkillsReq, opts ...grpc.CallOption) (*SearchSkillsResp, error)
		// 删除技能，需要客服权限
		DeleteSkill(ctx context.Context, in *DeleteSkillReq, opts ...grpc.CallOption) (*DeleteSkillResp, error)
		// 强行删除技能，必须要管理员权限
		MustDeleteSkill(ctx context.Context, in *MustDeleteSkillReq, opts ...grpc.CallOption) (*MustDeleteSkillResp, error)
	}

	defaultOther struct {
		cli zrpc.Client
	}
)

func NewOther(cli zrpc.Client) Other {
	return &defaultOther{
		cli: cli,
	}
}

// 添加技能，需要客服权限
func (m *defaultOther) AddSkill(ctx context.Context, in *AddSkillReq, opts ...grpc.CallOption) (*AddSkillResp, error) {
	client := user.NewOtherClient(m.cli.Conn())
	return client.AddSkill(ctx, in, opts...)
}

// 修改技能，需要客服权限
func (m *defaultOther) UpdateSkill(ctx context.Context, in *UpdateSkillReq, opts ...grpc.CallOption) (*UpdateSkillResp, error) {
	client := user.NewOtherClient(m.cli.Conn())
	return client.UpdateSkill(ctx, in, opts...)
}

// 搜索技能
func (m *defaultOther) SearchSkills(ctx context.Context, in *SearchSkillsReq, opts ...grpc.CallOption) (*SearchSkillsResp, error) {
	client := user.NewOtherClient(m.cli.Conn())
	return client.SearchSkills(ctx, in, opts...)
}

// 删除技能，需要客服权限
func (m *defaultOther) DeleteSkill(ctx context.Context, in *DeleteSkillReq, opts ...grpc.CallOption) (*DeleteSkillResp, error) {
	client := user.NewOtherClient(m.cli.Conn())
	return client.DeleteSkill(ctx, in, opts...)
}

// 强行删除技能，必须要管理员权限
func (m *defaultOther) MustDeleteSkill(ctx context.Context, in *MustDeleteSkillReq, opts ...grpc.CallOption) (*MustDeleteSkillResp, error) {
	client := user.NewOtherClient(m.cli.Conn())
	return client.MustDeleteSkill(ctx, in, opts...)
}

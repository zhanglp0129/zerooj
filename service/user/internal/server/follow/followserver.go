// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"zerooj/service/user/internal/logic/follow"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"
)

type FollowServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedFollowServer
}

func NewFollowServer(svcCtx *svc.ServiceContext) *FollowServer {
	return &FollowServer{
		svcCtx: svcCtx,
	}
}

// 获取所有关注，分页查询
func (s *FollowServer) GetFollowings(ctx context.Context, in *user.GetFollowingsReq) (*user.GetFollowingsResp, error) {
	l := followlogic.NewGetFollowingsLogic(ctx, s.svcCtx)
	return l.GetFollowings(in)
}

// 获取所有粉丝，分页查询
func (s *FollowServer) GetFans(ctx context.Context, in *user.GetFansReq) (*user.GetFansResp, error) {
	l := followlogic.NewGetFansLogic(ctx, s.svcCtx)
	return l.GetFans(in)
}

// 关注其他用户
func (s *FollowServer) FollowUser(ctx context.Context, in *user.FollowUserReq) (*user.FollowUserResp, error) {
	l := followlogic.NewFollowUserLogic(ctx, s.svcCtx)
	return l.FollowUser(in)
}
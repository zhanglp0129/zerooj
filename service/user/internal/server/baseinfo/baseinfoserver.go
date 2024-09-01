// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"zerooj/service/user/internal/logic/baseinfo"
	"zerooj/service/user/internal/svc"
	"zerooj/service/user/pb/user"
)

type BaseInfoServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedBaseInfoServer
}

func NewBaseInfoServer(svcCtx *svc.ServiceContext) *BaseInfoServer {
	return &BaseInfoServer{
		svcCtx: svcCtx,
	}
}

// 获取用户基本信息，并缓存
func (s *BaseInfoServer) GetBaseInfo(ctx context.Context, in *user.GetBaseInfoReq) (*user.GetBaseInfoResp, error) {
	l := baseinfologic.NewGetBaseInfoLogic(ctx, s.svcCtx)
	return l.GetBaseInfo(in)
}

// 修改用户基本信息，字段为空表示不修改
func (s *BaseInfoServer) UpdateBaseInfo(ctx context.Context, in *user.UpdateBaseInfoReq) (*user.UpdateBaseInfoResp, error) {
	l := baseinfologic.NewUpdateBaseInfoLogic(ctx, s.svcCtx)
	return l.UpdateBaseInfo(in)
}

// 根据用户名搜索用户，并缓存
func (s *BaseInfoServer) SearchByUsername(ctx context.Context, in *user.SearchByUsernameReq) (*user.SearchByUsernameResp, error) {
	l := baseinfologic.NewSearchByUsernameLogic(ctx, s.svcCtx)
	return l.SearchByUsername(in)
}

// 修改用户权限
func (s *BaseInfoServer) UpdatePermission(ctx context.Context, in *user.UpdatePermissionReq) (*user.UpdatePermissionResp, error) {
	l := baseinfologic.NewUpdatePermissionLogic(ctx, s.svcCtx)
	return l.UpdatePermission(in)
}

// 获取用户权限
func (s *BaseInfoServer) GetPermission(ctx context.Context, in *user.GetPermissionReq) (*user.GetPermissionResp, error) {
	l := baseinfologic.NewGetPermissionLogic(ctx, s.svcCtx)
	return l.GetPermission(in)
}

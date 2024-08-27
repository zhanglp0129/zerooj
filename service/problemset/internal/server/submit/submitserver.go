// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package server

import (
	"context"

	"zerooj/service/problemset/internal/logic/submit"
	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"
)

type SubmitServer struct {
	svcCtx *svc.ServiceContext
	problemset.UnimplementedSubmitServer
}

func NewSubmitServer(svcCtx *svc.ServiceContext) *SubmitServer {
	return &SubmitServer{
		svcCtx: svcCtx,
	}
}

// 用户提交代码
func (s *SubmitServer) SubmitCode(ctx context.Context, in *problemset.SubmitCodeReq) (*problemset.SubmitCodeResp, error) {
	l := submitlogic.NewSubmitCodeLogic(ctx, s.svcCtx)
	return l.SubmitCode(in)
}

// 获取题目的提交数
func (s *SubmitServer) GetProblemSubmitCount(ctx context.Context, in *problemset.GetProblemSubmitCountReq) (*problemset.GetProblemSubmitCountResp, error) {
	l := submitlogic.NewGetProblemSubmitCountLogic(ctx, s.svcCtx)
	return l.GetProblemSubmitCount(in)
}

// 分页获取用户提交
func (s *SubmitServer) GetUserSubmit(ctx context.Context, in *problemset.GetUserSubmitReq) (*problemset.GetUserSubmitResp, error) {
	l := submitlogic.NewGetUserSubmitLogic(ctx, s.svcCtx)
	return l.GetUserSubmit(in)
}

// 获取用户某一道题的全部提交
func (s *SubmitServer) GetUserProblemSubmit(ctx context.Context, in *problemset.GetUserProblemSubmitReq) (*problemset.GetUserProblemSubmitResp, error) {
	l := submitlogic.NewGetUserProblemSubmitLogic(ctx, s.svcCtx)
	return l.GetUserProblemSubmit(in)
}

// 获取通过id提交记录
func (s *SubmitServer) GetSubmitById(ctx context.Context, in *problemset.GetSubmitByIdReq) (*problemset.GetSubmitByIdResp, error) {
	l := submitlogic.NewGetSubmitByIdLogic(ctx, s.svcCtx)
	return l.GetSubmitById(in)
}
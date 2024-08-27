// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package server

import (
	"context"

	"zerooj/service/problemset/internal/logic/example"
	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"
)

type ExampleServer struct {
	svcCtx *svc.ServiceContext
	problemset.UnimplementedExampleServer
}

func NewExampleServer(svcCtx *svc.ServiceContext) *ExampleServer {
	return &ExampleServer{
		svcCtx: svcCtx,
	}
}

// 添加样例
func (s *ExampleServer) AddExample(ctx context.Context, in *problemset.AddExampleReq) (*problemset.AddExampleResp, error) {
	l := examplelogic.NewAddExampleLogic(ctx, s.svcCtx)
	return l.AddExample(in)
}

// 删除样例
func (s *ExampleServer) DeleteExample(ctx context.Context, in *problemset.DeleteExampleReq) (*problemset.DeleteExampleResp, error) {
	l := examplelogic.NewDeleteExampleLogic(ctx, s.svcCtx)
	return l.DeleteExample(in)
}

// 修改样例
func (s *ExampleServer) UpdateExample(ctx context.Context, in *problemset.UpdateExampleReq) (*problemset.UpdateExampleResp, error) {
	l := examplelogic.NewUpdateExampleLogic(ctx, s.svcCtx)
	return l.UpdateExample(in)
}

// 获取样例
func (s *ExampleServer) GetExample(ctx context.Context, in *problemset.GetExampleReq) (*problemset.GetExampleResp, error) {
	l := examplelogic.NewGetExampleLogic(ctx, s.svcCtx)
	return l.GetExample(in)
}

// 获取题目的所有样例
func (s *ExampleServer) GetProblemExamples(ctx context.Context, in *problemset.GetProblemExamplesReq) (*problemset.GetProblemExamplesResp, error) {
	l := examplelogic.NewGetProblemExamplesLogic(ctx, s.svcCtx)
	return l.GetProblemExamples(in)
}
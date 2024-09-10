// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package server

import (
	"context"

	"zerooj/service/problemset/internal/logic/hint"
	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"
)

type HintServer struct {
	svcCtx *svc.ServiceContext
	problemset.UnimplementedHintServer
}

func NewHintServer(svcCtx *svc.ServiceContext) *HintServer {
	return &HintServer{
		svcCtx: svcCtx,
	}
}

// 添加提示
func (s *HintServer) AddHint(stream problemset.Hint_AddHintServer) error {
	l := hintlogic.NewAddHintLogic(stream.Context(), s.svcCtx)
	return l.AddHint(stream)
}

// 删除提示
func (s *HintServer) DeleteHint(ctx context.Context, in *problemset.DeleteHintReq) (*problemset.DeleteHintResp, error) {
	l := hintlogic.NewDeleteHintLogic(ctx, s.svcCtx)
	return l.DeleteHint(in)
}

// 修改提示
func (s *HintServer) UpdateHint(stream problemset.Hint_UpdateHintServer) error {
	l := hintlogic.NewUpdateHintLogic(stream.Context(), s.svcCtx)
	return l.UpdateHint(stream)
}

// 获取提示
func (s *HintServer) GetHint(in *problemset.GetHintReq, stream problemset.Hint_GetHintServer) error {
	l := hintlogic.NewGetHintLogic(stream.Context(), s.svcCtx)
	return l.GetHint(in, stream)
}

// 获取题目所有提示
func (s *HintServer) GetProblemHints(in *problemset.GetProblemHintsReq, stream problemset.Hint_GetProblemHintsServer) error {
	l := hintlogic.NewGetProblemHintsLogic(stream.Context(), s.svcCtx)
	return l.GetProblemHints(in, stream)
}
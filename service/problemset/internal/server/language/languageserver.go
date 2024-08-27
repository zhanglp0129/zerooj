// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package server

import (
	"context"

	"zerooj/service/problemset/internal/logic/language"
	"zerooj/service/problemset/internal/svc"
	"zerooj/service/problemset/pb/problemset"
)

type LanguageServer struct {
	svcCtx *svc.ServiceContext
	problemset.UnimplementedLanguageServer
}

func NewLanguageServer(svcCtx *svc.ServiceContext) *LanguageServer {
	return &LanguageServer{
		svcCtx: svcCtx,
	}
}

// 添加语言
func (s *LanguageServer) AddLanguage(ctx context.Context, in *problemset.AddLanguageReq) (*problemset.AddLanguageResp, error) {
	l := languagelogic.NewAddLanguageLogic(ctx, s.svcCtx)
	return l.AddLanguage(in)
}

// 删除语言
func (s *LanguageServer) DeleteLanguage(ctx context.Context, in *problemset.DeleteLanguageReq) (*problemset.DeleteLanguageResp, error) {
	l := languagelogic.NewDeleteLanguageLogic(ctx, s.svcCtx)
	return l.DeleteLanguage(in)
}

// 更新语言
func (s *LanguageServer) UpdateLanguage(ctx context.Context, in *problemset.UpdateLanguageReq) (*problemset.UpdateLanguageResp, error) {
	l := languagelogic.NewUpdateLanguageLogic(ctx, s.svcCtx)
	return l.UpdateLanguage(in)
}

// 获取所有语言
func (s *LanguageServer) GetLanguages(ctx context.Context, in *problemset.GetLanguagesReq) (*problemset.GetLanguagesResp, error) {
	l := languagelogic.NewGetLanguagesLogic(ctx, s.svcCtx)
	return l.GetLanguages(in)
}

// 根据id获取语言
func (s *LanguageServer) GetLanguageById(ctx context.Context, in *problemset.GetLanguageByIdReq) (*problemset.GetLanguageByIdResp, error) {
	l := languagelogic.NewGetLanguageByIdLogic(ctx, s.svcCtx)
	return l.GetLanguageById(in)
}

// 给题目增加语言
func (s *LanguageServer) ProblemAddLanguages(ctx context.Context, in *problemset.ProblemAddLanguagesReq) (*problemset.ProblemAddLanguagesResp, error) {
	l := languagelogic.NewProblemAddLanguagesLogic(ctx, s.svcCtx)
	return l.ProblemAddLanguages(in)
}

// 题目删除语言
func (s *LanguageServer) ProblemDeleteLanguages(ctx context.Context, in *problemset.ProblemDeleteLanguagesReq) (*problemset.ProblemDeleteLanguagesResp, error) {
	l := languagelogic.NewProblemDeleteLanguagesLogic(ctx, s.svcCtx)
	return l.ProblemDeleteLanguages(in)
}

// 获取题目的所有语言
func (s *LanguageServer) GetProblemLanguages(ctx context.Context, in *problemset.GetProblemLanguagesReq) (*problemset.GetProblemLanguagesResp, error) {
	l := languagelogic.NewGetProblemLanguagesLogic(ctx, s.svcCtx)
	return l.GetProblemLanguages(in)
}
// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package language

import (
	"context"

	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddExampleReq              = problemset.AddExampleReq
	AddExampleResp             = problemset.AddExampleResp
	AddHintReq                 = problemset.AddHintReq
	AddHintResp                = problemset.AddHintResp
	AddJudgeDataReq            = problemset.AddJudgeDataReq
	AddJudgeDataResp           = problemset.AddJudgeDataResp
	AddLanguageReq             = problemset.AddLanguageReq
	AddLanguageResp            = problemset.AddLanguageResp
	AddProblemReq              = problemset.AddProblemReq
	AddProblemResp             = problemset.AddProblemResp
	AddTagsReq                 = problemset.AddTagsReq
	AddTagsResp                = problemset.AddTagsResp
	DeleteExampleReq           = problemset.DeleteExampleReq
	DeleteExampleResp          = problemset.DeleteExampleResp
	DeleteHintReq              = problemset.DeleteHintReq
	DeleteHintResp             = problemset.DeleteHintResp
	DeleteJudgeDataReq         = problemset.DeleteJudgeDataReq
	DeleteJudgeDataResp        = problemset.DeleteJudgeDataResp
	DeleteLanguageReq          = problemset.DeleteLanguageReq
	DeleteLanguageResp         = problemset.DeleteLanguageResp
	DeleteProblemReq           = problemset.DeleteProblemReq
	DeleteProblemResp          = problemset.DeleteProblemResp
	DeleteTagsReq              = problemset.DeleteTagsReq
	DeleteTagsResp             = problemset.DeleteTagsResp
	ExampleInfo                = problemset.ExampleInfo
	GetAllTagsReq              = problemset.GetAllTagsReq
	GetAllTagsResp             = problemset.GetAllTagsResp
	GetExampleReq              = problemset.GetExampleReq
	GetExampleResp             = problemset.GetExampleResp
	GetHintReq                 = problemset.GetHintReq
	GetHintResp                = problemset.GetHintResp
	GetJudgeDataReq            = problemset.GetJudgeDataReq
	GetJudgeDataResp           = problemset.GetJudgeDataResp
	GetLanguageByIdReq         = problemset.GetLanguageByIdReq
	GetLanguageByIdResp        = problemset.GetLanguageByIdResp
	GetLanguagesReq            = problemset.GetLanguagesReq
	GetLanguagesResp           = problemset.GetLanguagesResp
	GetProblemContentReq       = problemset.GetProblemContentReq
	GetProblemContentResp      = problemset.GetProblemContentResp
	GetProblemExamplesReq      = problemset.GetProblemExamplesReq
	GetProblemExamplesResp     = problemset.GetProblemExamplesResp
	GetProblemHintsReq         = problemset.GetProblemHintsReq
	GetProblemHintsResp        = problemset.GetProblemHintsResp
	GetProblemLanguagesReq     = problemset.GetProblemLanguagesReq
	GetProblemLanguagesResp    = problemset.GetProblemLanguagesResp
	GetProblemSubmitCountReq   = problemset.GetProblemSubmitCountReq
	GetProblemSubmitCountResp  = problemset.GetProblemSubmitCountResp
	GetProblemTagsReq          = problemset.GetProblemTagsReq
	GetProblemTagsResp         = problemset.GetProblemTagsResp
	GetSubmitByIdReq           = problemset.GetSubmitByIdReq
	GetSubmitByIdResp          = problemset.GetSubmitByIdResp
	GetUserProblemSubmitReq    = problemset.GetUserProblemSubmitReq
	GetUserProblemSubmitResp   = problemset.GetUserProblemSubmitResp
	GetUserSubmitReq           = problemset.GetUserSubmitReq
	GetUserSubmitResp          = problemset.GetUserSubmitResp
	HintInfo                   = problemset.HintInfo
	JudgeDataInfo              = problemset.JudgeDataInfo
	LanguageInfo               = problemset.LanguageInfo
	MustDeleteTagsReq          = problemset.MustDeleteTagsReq
	MustDeleteTagsResp         = problemset.MustDeleteTagsResp
	ProblemAddLanguagesReq     = problemset.ProblemAddLanguagesReq
	ProblemAddLanguagesResp    = problemset.ProblemAddLanguagesResp
	ProblemAddTagsReq          = problemset.ProblemAddTagsReq
	ProblemAddTagsResp         = problemset.ProblemAddTagsResp
	ProblemDeleteLanguagesReq  = problemset.ProblemDeleteLanguagesReq
	ProblemDeleteLanguagesResp = problemset.ProblemDeleteLanguagesResp
	ProblemDeleteTagsReq       = problemset.ProblemDeleteTagsReq
	ProblemDeleteTagsResp      = problemset.ProblemDeleteTagsResp
	ProblemInfo                = problemset.ProblemInfo
	ProblemLanguagesModel      = problemset.ProblemLanguagesModel
	SearchProblemReq           = problemset.SearchProblemReq
	SearchProblemResp          = problemset.SearchProblemResp
	SubmitCodeReq              = problemset.SubmitCodeReq
	SubmitCodeResp             = problemset.SubmitCodeResp
	SubmitInfo                 = problemset.SubmitInfo
	TagInfo                    = problemset.TagInfo
	UpdateExampleReq           = problemset.UpdateExampleReq
	UpdateExampleResp          = problemset.UpdateExampleResp
	UpdateHintReq              = problemset.UpdateHintReq
	UpdateHintResp             = problemset.UpdateHintResp
	UpdateJudgeDataReq         = problemset.UpdateJudgeDataReq
	UpdateJudgeDataResp        = problemset.UpdateJudgeDataResp
	UpdateLanguageReq          = problemset.UpdateLanguageReq
	UpdateLanguageResp         = problemset.UpdateLanguageResp
	UpdateProblemReq           = problemset.UpdateProblemReq
	UpdateProblemResp          = problemset.UpdateProblemResp
	UpdateTagsReq              = problemset.UpdateTagsReq
	UpdateTagsResp             = problemset.UpdateTagsResp

	Language interface {
		// 添加语言
		AddLanguage(ctx context.Context, in *AddLanguageReq, opts ...grpc.CallOption) (*AddLanguageResp, error)
		// 删除语言
		DeleteLanguage(ctx context.Context, in *DeleteLanguageReq, opts ...grpc.CallOption) (*DeleteLanguageResp, error)
		// 更新语言
		UpdateLanguage(ctx context.Context, in *UpdateLanguageReq, opts ...grpc.CallOption) (*UpdateLanguageResp, error)
		// 获取所有语言
		GetLanguages(ctx context.Context, in *GetLanguagesReq, opts ...grpc.CallOption) (*GetLanguagesResp, error)
		// 根据id获取语言
		GetLanguageById(ctx context.Context, in *GetLanguageByIdReq, opts ...grpc.CallOption) (*GetLanguageByIdResp, error)
		// 给题目增加语言
		ProblemAddLanguages(ctx context.Context, in *ProblemAddLanguagesReq, opts ...grpc.CallOption) (*ProblemAddLanguagesResp, error)
		// 题目删除语言
		ProblemDeleteLanguages(ctx context.Context, in *ProblemDeleteLanguagesReq, opts ...grpc.CallOption) (*ProblemDeleteLanguagesResp, error)
		// 获取题目的所有语言
		GetProblemLanguages(ctx context.Context, in *GetProblemLanguagesReq, opts ...grpc.CallOption) (*GetProblemLanguagesResp, error)
	}

	defaultLanguage struct {
		cli zrpc.Client
	}
)

func NewLanguage(cli zrpc.Client) Language {
	return &defaultLanguage{
		cli: cli,
	}
}

// 添加语言
func (m *defaultLanguage) AddLanguage(ctx context.Context, in *AddLanguageReq, opts ...grpc.CallOption) (*AddLanguageResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.AddLanguage(ctx, in, opts...)
}

// 删除语言
func (m *defaultLanguage) DeleteLanguage(ctx context.Context, in *DeleteLanguageReq, opts ...grpc.CallOption) (*DeleteLanguageResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.DeleteLanguage(ctx, in, opts...)
}

// 更新语言
func (m *defaultLanguage) UpdateLanguage(ctx context.Context, in *UpdateLanguageReq, opts ...grpc.CallOption) (*UpdateLanguageResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.UpdateLanguage(ctx, in, opts...)
}

// 获取所有语言
func (m *defaultLanguage) GetLanguages(ctx context.Context, in *GetLanguagesReq, opts ...grpc.CallOption) (*GetLanguagesResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.GetLanguages(ctx, in, opts...)
}

// 根据id获取语言
func (m *defaultLanguage) GetLanguageById(ctx context.Context, in *GetLanguageByIdReq, opts ...grpc.CallOption) (*GetLanguageByIdResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.GetLanguageById(ctx, in, opts...)
}

// 给题目增加语言
func (m *defaultLanguage) ProblemAddLanguages(ctx context.Context, in *ProblemAddLanguagesReq, opts ...grpc.CallOption) (*ProblemAddLanguagesResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.ProblemAddLanguages(ctx, in, opts...)
}

// 题目删除语言
func (m *defaultLanguage) ProblemDeleteLanguages(ctx context.Context, in *ProblemDeleteLanguagesReq, opts ...grpc.CallOption) (*ProblemDeleteLanguagesResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.ProblemDeleteLanguages(ctx, in, opts...)
}

// 获取题目的所有语言
func (m *defaultLanguage) GetProblemLanguages(ctx context.Context, in *GetProblemLanguagesReq, opts ...grpc.CallOption) (*GetProblemLanguagesResp, error) {
	client := problemset.NewLanguageClient(m.cli.Conn())
	return client.GetProblemLanguages(ctx, in, opts...)
}

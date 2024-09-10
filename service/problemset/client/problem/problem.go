// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package problem

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
	AddTagReq                  = problemset.AddTagReq
	AddTagResp                 = problemset.AddTagResp
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
	DeleteTagReq               = problemset.DeleteTagReq
	DeleteTagResp              = problemset.DeleteTagResp
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
	MustDeleteTagReq           = problemset.MustDeleteTagReq
	MustDeleteTagResp          = problemset.MustDeleteTagResp
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
	UpdateTagReq               = problemset.UpdateTagReq
	UpdateTagResp              = problemset.UpdateTagResp

	Problem interface {
		// 添加问题
		AddProblem(ctx context.Context, opts ...grpc.CallOption) (problemset.Problem_AddProblemClient, error)
		// 删除问题
		DeleteProblem(ctx context.Context, in *DeleteProblemReq, opts ...grpc.CallOption) (*DeleteProblemResp, error)
		// 获取问题信息，可缓存
		GetProblemContent(ctx context.Context, in *GetProblemContentReq, opts ...grpc.CallOption) (problemset.Problem_GetProblemContentClient, error)
		// 更新问题
		UpdateProblem(ctx context.Context, opts ...grpc.CallOption) (problemset.Problem_UpdateProblemClient, error)
		// 分页搜索题目
		SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error)
	}

	defaultProblem struct {
		cli zrpc.Client
	}
)

func NewProblem(cli zrpc.Client) Problem {
	return &defaultProblem{
		cli: cli,
	}
}

// 添加问题
func (m *defaultProblem) AddProblem(ctx context.Context, opts ...grpc.CallOption) (problemset.Problem_AddProblemClient, error) {
	client := problemset.NewProblemClient(m.cli.Conn())
	return client.AddProblem(ctx, opts...)
}

// 删除问题
func (m *defaultProblem) DeleteProblem(ctx context.Context, in *DeleteProblemReq, opts ...grpc.CallOption) (*DeleteProblemResp, error) {
	client := problemset.NewProblemClient(m.cli.Conn())
	return client.DeleteProblem(ctx, in, opts...)
}

// 获取问题信息，可缓存
func (m *defaultProblem) GetProblemContent(ctx context.Context, in *GetProblemContentReq, opts ...grpc.CallOption) (problemset.Problem_GetProblemContentClient, error) {
	client := problemset.NewProblemClient(m.cli.Conn())
	return client.GetProblemContent(ctx, in, opts...)
}

// 更新问题
func (m *defaultProblem) UpdateProblem(ctx context.Context, opts ...grpc.CallOption) (problemset.Problem_UpdateProblemClient, error) {
	client := problemset.NewProblemClient(m.cli.Conn())
	return client.UpdateProblem(ctx, opts...)
}

// 分页搜索题目
func (m *defaultProblem) SearchProblem(ctx context.Context, in *SearchProblemReq, opts ...grpc.CallOption) (*SearchProblemResp, error) {
	client := problemset.NewProblemClient(m.cli.Conn())
	return client.SearchProblem(ctx, in, opts...)
}

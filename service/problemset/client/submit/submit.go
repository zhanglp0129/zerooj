// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package submit

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

	Submit interface {
		// 用户提交代码
		SubmitCode(ctx context.Context, in *SubmitCodeReq, opts ...grpc.CallOption) (*SubmitCodeResp, error)
		// 获取题目的提交数
		GetProblemSubmitCount(ctx context.Context, in *GetProblemSubmitCountReq, opts ...grpc.CallOption) (*GetProblemSubmitCountResp, error)
		// 分页获取用户提交
		GetUserSubmit(ctx context.Context, in *GetUserSubmitReq, opts ...grpc.CallOption) (*GetUserSubmitResp, error)
		// 获取用户某一道题的全部提交
		GetUserProblemSubmit(ctx context.Context, in *GetUserProblemSubmitReq, opts ...grpc.CallOption) (*GetUserProblemSubmitResp, error)
		// 获取通过id提交记录
		GetSubmitById(ctx context.Context, in *GetSubmitByIdReq, opts ...grpc.CallOption) (*GetSubmitByIdResp, error)
	}

	defaultSubmit struct {
		cli zrpc.Client
	}
)

func NewSubmit(cli zrpc.Client) Submit {
	return &defaultSubmit{
		cli: cli,
	}
}

// 用户提交代码
func (m *defaultSubmit) SubmitCode(ctx context.Context, in *SubmitCodeReq, opts ...grpc.CallOption) (*SubmitCodeResp, error) {
	client := problemset.NewSubmitClient(m.cli.Conn())
	return client.SubmitCode(ctx, in, opts...)
}

// 获取题目的提交数
func (m *defaultSubmit) GetProblemSubmitCount(ctx context.Context, in *GetProblemSubmitCountReq, opts ...grpc.CallOption) (*GetProblemSubmitCountResp, error) {
	client := problemset.NewSubmitClient(m.cli.Conn())
	return client.GetProblemSubmitCount(ctx, in, opts...)
}

// 分页获取用户提交
func (m *defaultSubmit) GetUserSubmit(ctx context.Context, in *GetUserSubmitReq, opts ...grpc.CallOption) (*GetUserSubmitResp, error) {
	client := problemset.NewSubmitClient(m.cli.Conn())
	return client.GetUserSubmit(ctx, in, opts...)
}

// 获取用户某一道题的全部提交
func (m *defaultSubmit) GetUserProblemSubmit(ctx context.Context, in *GetUserProblemSubmitReq, opts ...grpc.CallOption) (*GetUserProblemSubmitResp, error) {
	client := problemset.NewSubmitClient(m.cli.Conn())
	return client.GetUserProblemSubmit(ctx, in, opts...)
}

// 获取通过id提交记录
func (m *defaultSubmit) GetSubmitById(ctx context.Context, in *GetSubmitByIdReq, opts ...grpc.CallOption) (*GetSubmitByIdResp, error) {
	client := problemset.NewSubmitClient(m.cli.Conn())
	return client.GetSubmitById(ctx, in, opts...)
}

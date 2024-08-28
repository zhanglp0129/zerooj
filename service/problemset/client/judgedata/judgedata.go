// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package judgedata

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

	JudgeData interface {
		// 添加测评数据
		AddJudgeData(ctx context.Context, opts ...grpc.CallOption) (problemset.JudgeData_AddJudgeDataClient, error)
		// 删除测评数据
		DeleteJudgeData(ctx context.Context, in *DeleteJudgeDataReq, opts ...grpc.CallOption) (*DeleteJudgeDataResp, error)
		// 修改测评数据
		UpdateJudgeData(ctx context.Context, opts ...grpc.CallOption) (problemset.JudgeData_UpdateJudgeDataClient, error)
		// 获取题目的测评数据，返回minio对象名称，可缓存
		GetJudgeData(ctx context.Context, in *GetJudgeDataReq, opts ...grpc.CallOption) (problemset.JudgeData_GetJudgeDataClient, error)
	}

	defaultJudgeData struct {
		cli zrpc.Client
	}
)

func NewJudgeData(cli zrpc.Client) JudgeData {
	return &defaultJudgeData{
		cli: cli,
	}
}

// 添加测评数据
func (m *defaultJudgeData) AddJudgeData(ctx context.Context, opts ...grpc.CallOption) (problemset.JudgeData_AddJudgeDataClient, error) {
	client := problemset.NewJudgeDataClient(m.cli.Conn())
	return client.AddJudgeData(ctx, opts...)
}

// 删除测评数据
func (m *defaultJudgeData) DeleteJudgeData(ctx context.Context, in *DeleteJudgeDataReq, opts ...grpc.CallOption) (*DeleteJudgeDataResp, error) {
	client := problemset.NewJudgeDataClient(m.cli.Conn())
	return client.DeleteJudgeData(ctx, in, opts...)
}

// 修改测评数据
func (m *defaultJudgeData) UpdateJudgeData(ctx context.Context, opts ...grpc.CallOption) (problemset.JudgeData_UpdateJudgeDataClient, error) {
	client := problemset.NewJudgeDataClient(m.cli.Conn())
	return client.UpdateJudgeData(ctx, opts...)
}

// 获取题目的测评数据，返回minio对象名称，可缓存
func (m *defaultJudgeData) GetJudgeData(ctx context.Context, in *GetJudgeDataReq, opts ...grpc.CallOption) (problemset.JudgeData_GetJudgeDataClient, error) {
	client := problemset.NewJudgeDataClient(m.cli.Conn())
	return client.GetJudgeData(ctx, in, opts...)
}

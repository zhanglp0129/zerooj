// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package example

import (
	"context"

	"zerooj/service/problemset/pb/problemset"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddExampleReq          = problemset.AddExampleReq
	AddExampleResp         = problemset.AddExampleResp
	AddHintReq             = problemset.AddHintReq
	AddHintResp            = problemset.AddHintResp
	AddJudgeDataReq        = problemset.AddJudgeDataReq
	AddJudgeDataResp       = problemset.AddJudgeDataResp
	AddProblemReq          = problemset.AddProblemReq
	AddProblemResp         = problemset.AddProblemResp
	AddTagsReq             = problemset.AddTagsReq
	AddTagsResp            = problemset.AddTagsResp
	DeleteExampleReq       = problemset.DeleteExampleReq
	DeleteExampleResp      = problemset.DeleteExampleResp
	DeleteHintReq          = problemset.DeleteHintReq
	DeleteHintResp         = problemset.DeleteHintResp
	DeleteJudgeDataReq     = problemset.DeleteJudgeDataReq
	DeleteJudgeDataResp    = problemset.DeleteJudgeDataResp
	DeleteProblemReq       = problemset.DeleteProblemReq
	DeleteProblemResp      = problemset.DeleteProblemResp
	DeleteTagsReq          = problemset.DeleteTagsReq
	DeleteTagsResp         = problemset.DeleteTagsResp
	ExampleInfo            = problemset.ExampleInfo
	GetAllTagsReq          = problemset.GetAllTagsReq
	GetAllTagsResp         = problemset.GetAllTagsResp
	GetExampleReq          = problemset.GetExampleReq
	GetExampleResp         = problemset.GetExampleResp
	GetHintReq             = problemset.GetHintReq
	GetHintResp            = problemset.GetHintResp
	GetJudgeDataReq        = problemset.GetJudgeDataReq
	GetJudgeDataResp       = problemset.GetJudgeDataResp
	GetProblemContentReq   = problemset.GetProblemContentReq
	GetProblemContentResp  = problemset.GetProblemContentResp
	GetProblemExamplesReq  = problemset.GetProblemExamplesReq
	GetProblemExamplesResp = problemset.GetProblemExamplesResp
	GetProblemHintsReq     = problemset.GetProblemHintsReq
	GetProblemHintsResp    = problemset.GetProblemHintsResp
	GetProblemTagsReq      = problemset.GetProblemTagsReq
	GetProblemTagsResp     = problemset.GetProblemTagsResp
	HintInfo               = problemset.HintInfo
	JudgeDataInfo          = problemset.JudgeDataInfo
	MustDeleteTagsReq      = problemset.MustDeleteTagsReq
	MustDeleteTagsResp     = problemset.MustDeleteTagsResp
	ProblemAddTagsReq      = problemset.ProblemAddTagsReq
	ProblemAddTagsResp     = problemset.ProblemAddTagsResp
	ProblemDeleteTagsReq   = problemset.ProblemDeleteTagsReq
	ProblemDeleteTagsResp  = problemset.ProblemDeleteTagsResp
	SearchProblemReq       = problemset.SearchProblemReq
	SearchProblemResp      = problemset.SearchProblemResp
	TagInfo                = problemset.TagInfo
	UpdateExampleReq       = problemset.UpdateExampleReq
	UpdateExampleResp      = problemset.UpdateExampleResp
	UpdateHintReq          = problemset.UpdateHintReq
	UpdateHintResp         = problemset.UpdateHintResp
	UpdateJudgeDataReq     = problemset.UpdateJudgeDataReq
	UpdateJudgeDataResp    = problemset.UpdateJudgeDataResp
	UpdateProblemReq       = problemset.UpdateProblemReq
	UpdateProblemResp      = problemset.UpdateProblemResp
	UpdateTagsReq          = problemset.UpdateTagsReq
	UpdateTagsResp         = problemset.UpdateTagsResp

	Example interface {
		// 添加样例，客服权限
		AddExample(ctx context.Context, in *AddExampleReq, opts ...grpc.CallOption) (*AddExampleResp, error)
		// 删除样例，客服权限
		DeleteExample(ctx context.Context, in *DeleteExampleReq, opts ...grpc.CallOption) (*DeleteExampleResp, error)
		// 修改样例，客服权限
		UpdateExample(ctx context.Context, in *UpdateExampleReq, opts ...grpc.CallOption) (*UpdateExampleResp, error)
		// 获取样例
		GetExample(ctx context.Context, in *GetExampleReq, opts ...grpc.CallOption) (*GetExampleResp, error)
		// 获取题目的所有样例
		GetProblemExamples(ctx context.Context, in *GetProblemExamplesReq, opts ...grpc.CallOption) (*GetProblemExamplesResp, error)
	}

	defaultExample struct {
		cli zrpc.Client
	}
)

func NewExample(cli zrpc.Client) Example {
	return &defaultExample{
		cli: cli,
	}
}

// 添加样例，客服权限
func (m *defaultExample) AddExample(ctx context.Context, in *AddExampleReq, opts ...grpc.CallOption) (*AddExampleResp, error) {
	client := problemset.NewExampleClient(m.cli.Conn())
	return client.AddExample(ctx, in, opts...)
}

// 删除样例，客服权限
func (m *defaultExample) DeleteExample(ctx context.Context, in *DeleteExampleReq, opts ...grpc.CallOption) (*DeleteExampleResp, error) {
	client := problemset.NewExampleClient(m.cli.Conn())
	return client.DeleteExample(ctx, in, opts...)
}

// 修改样例，客服权限
func (m *defaultExample) UpdateExample(ctx context.Context, in *UpdateExampleReq, opts ...grpc.CallOption) (*UpdateExampleResp, error) {
	client := problemset.NewExampleClient(m.cli.Conn())
	return client.UpdateExample(ctx, in, opts...)
}

// 获取样例
func (m *defaultExample) GetExample(ctx context.Context, in *GetExampleReq, opts ...grpc.CallOption) (*GetExampleResp, error) {
	client := problemset.NewExampleClient(m.cli.Conn())
	return client.GetExample(ctx, in, opts...)
}

// 获取题目的所有样例
func (m *defaultExample) GetProblemExamples(ctx context.Context, in *GetProblemExamplesReq, opts ...grpc.CallOption) (*GetProblemExamplesResp, error) {
	client := problemset.NewExampleClient(m.cli.Conn())
	return client.GetProblemExamples(ctx, in, opts...)
}
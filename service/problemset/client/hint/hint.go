// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package hint

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

	Hint interface {
		// 添加提示，客服权限
		AddHint(ctx context.Context, in *AddHintReq, opts ...grpc.CallOption) (*AddHintResp, error)
		// 删除提示，客服权限
		DeleteHint(ctx context.Context, in *DeleteHintReq, opts ...grpc.CallOption) (*DeleteHintResp, error)
		// 修改提示，客服权限
		UpdateHint(ctx context.Context, in *UpdateHintReq, opts ...grpc.CallOption) (*UpdateHintResp, error)
		// 获取提示
		GetHint(ctx context.Context, in *GetHintReq, opts ...grpc.CallOption) (*GetHintResp, error)
		// 获取题目所有提示
		GetProblemHints(ctx context.Context, in *GetProblemHintsReq, opts ...grpc.CallOption) (*GetProblemHintsResp, error)
	}

	defaultHint struct {
		cli zrpc.Client
	}
)

func NewHint(cli zrpc.Client) Hint {
	return &defaultHint{
		cli: cli,
	}
}

// 添加提示，客服权限
func (m *defaultHint) AddHint(ctx context.Context, in *AddHintReq, opts ...grpc.CallOption) (*AddHintResp, error) {
	client := problemset.NewHintClient(m.cli.Conn())
	return client.AddHint(ctx, in, opts...)
}

// 删除提示，客服权限
func (m *defaultHint) DeleteHint(ctx context.Context, in *DeleteHintReq, opts ...grpc.CallOption) (*DeleteHintResp, error) {
	client := problemset.NewHintClient(m.cli.Conn())
	return client.DeleteHint(ctx, in, opts...)
}

// 修改提示，客服权限
func (m *defaultHint) UpdateHint(ctx context.Context, in *UpdateHintReq, opts ...grpc.CallOption) (*UpdateHintResp, error) {
	client := problemset.NewHintClient(m.cli.Conn())
	return client.UpdateHint(ctx, in, opts...)
}

// 获取提示
func (m *defaultHint) GetHint(ctx context.Context, in *GetHintReq, opts ...grpc.CallOption) (*GetHintResp, error) {
	client := problemset.NewHintClient(m.cli.Conn())
	return client.GetHint(ctx, in, opts...)
}

// 获取题目所有提示
func (m *defaultHint) GetProblemHints(ctx context.Context, in *GetProblemHintsReq, opts ...grpc.CallOption) (*GetProblemHintsResp, error) {
	client := problemset.NewHintClient(m.cli.Conn())
	return client.GetProblemHints(ctx, in, opts...)
}
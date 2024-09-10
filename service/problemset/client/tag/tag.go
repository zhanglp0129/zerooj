// Code generated by goctl. DO NOT EDIT.
// Source: problemset.proto

package tag

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

	Tag interface {
		// 添加标签
		AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error)
		// 删除标签
		DeleteTag(ctx context.Context, in *DeleteTagReq, opts ...grpc.CallOption) (*DeleteTagResp, error)
		// 强行删除标签
		MustDeleteTag(ctx context.Context, in *MustDeleteTagReq, opts ...grpc.CallOption) (*MustDeleteTagResp, error)
		// 更新标签
		UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error)
		// 获取所有标签，可缓存
		GetAllTags(ctx context.Context, in *GetAllTagsReq, opts ...grpc.CallOption) (*GetAllTagsResp, error)
		// 给题目添加标签，最多5个
		ProblemAddTags(ctx context.Context, in *ProblemAddTagsReq, opts ...grpc.CallOption) (*ProblemAddTagsResp, error)
		// 给题目删除标签
		ProblemDeleteTags(ctx context.Context, in *ProblemDeleteTagsReq, opts ...grpc.CallOption) (*ProblemDeleteTagsResp, error)
		// 获取一个题目的所有标签，可缓存
		GetProblemTags(ctx context.Context, in *GetProblemTagsReq, opts ...grpc.CallOption) (*GetProblemTagsResp, error)
	}

	defaultTag struct {
		cli zrpc.Client
	}
)

func NewTag(cli zrpc.Client) Tag {
	return &defaultTag{
		cli: cli,
	}
}

// 添加标签
func (m *defaultTag) AddTag(ctx context.Context, in *AddTagReq, opts ...grpc.CallOption) (*AddTagResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.AddTag(ctx, in, opts...)
}

// 删除标签
func (m *defaultTag) DeleteTag(ctx context.Context, in *DeleteTagReq, opts ...grpc.CallOption) (*DeleteTagResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.DeleteTag(ctx, in, opts...)
}

// 强行删除标签
func (m *defaultTag) MustDeleteTag(ctx context.Context, in *MustDeleteTagReq, opts ...grpc.CallOption) (*MustDeleteTagResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.MustDeleteTag(ctx, in, opts...)
}

// 更新标签
func (m *defaultTag) UpdateTag(ctx context.Context, in *UpdateTagReq, opts ...grpc.CallOption) (*UpdateTagResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.UpdateTag(ctx, in, opts...)
}

// 获取所有标签，可缓存
func (m *defaultTag) GetAllTags(ctx context.Context, in *GetAllTagsReq, opts ...grpc.CallOption) (*GetAllTagsResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.GetAllTags(ctx, in, opts...)
}

// 给题目添加标签，最多5个
func (m *defaultTag) ProblemAddTags(ctx context.Context, in *ProblemAddTagsReq, opts ...grpc.CallOption) (*ProblemAddTagsResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.ProblemAddTags(ctx, in, opts...)
}

// 给题目删除标签
func (m *defaultTag) ProblemDeleteTags(ctx context.Context, in *ProblemDeleteTagsReq, opts ...grpc.CallOption) (*ProblemDeleteTagsResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.ProblemDeleteTags(ctx, in, opts...)
}

// 获取一个题目的所有标签，可缓存
func (m *defaultTag) GetProblemTags(ctx context.Context, in *GetProblemTagsReq, opts ...grpc.CallOption) (*GetProblemTagsResp, error) {
	client := problemset.NewTagClient(m.cli.Conn())
	return client.GetProblemTags(ctx, in, opts...)
}
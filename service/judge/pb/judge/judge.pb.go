// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: judge.proto

package judge

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JudgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId   int64  `protobuf:"varint,1,opt,name=problemId,proto3" json:"problemId,omitempty"`
	TimeLimit   int64  `protobuf:"varint,2,opt,name=timeLimit,proto3" json:"timeLimit,omitempty"`
	MemoryLimit int64  `protobuf:"varint,3,opt,name=memoryLimit,proto3" json:"memoryLimit,omitempty"`
	Code        string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *JudgeReq) Reset() {
	*x = JudgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeReq) ProtoMessage() {}

func (x *JudgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeReq.ProtoReflect.Descriptor instead.
func (*JudgeReq) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{0}
}

func (x *JudgeReq) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *JudgeReq) GetTimeLimit() int64 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *JudgeReq) GetMemoryLimit() int64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *JudgeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type JudgeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 测评结果
	Result uint32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	// 运行时间
	RunTime int64 `protobuf:"varint,2,opt,name=runTime,proto3" json:"runTime,omitempty"`
	// 消耗内存
	MemoryConsumption int64 `protobuf:"varint,3,opt,name=memoryConsumption,proto3" json:"memoryConsumption,omitempty"`
}

func (x *JudgeResult) Reset() {
	*x = JudgeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeResult) ProtoMessage() {}

func (x *JudgeResult) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeResult.ProtoReflect.Descriptor instead.
func (*JudgeResult) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{1}
}

func (x *JudgeResult) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *JudgeResult) GetRunTime() int64 {
	if x != nil {
		return x.RunTime
	}
	return 0
}

func (x *JudgeResult) GetMemoryConsumption() int64 {
	if x != nil {
		return x.MemoryConsumption
	}
	return 0
}

type JudgeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *JudgeData) Reset() {
	*x = JudgeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeData) ProtoMessage() {}

func (x *JudgeData) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeData.ProtoReflect.Descriptor instead.
func (*JudgeData) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{2}
}

func (x *JudgeData) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *JudgeData) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type JudgeWithDataReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JudgeData   []*JudgeData `protobuf:"bytes,4,rep,name=judgeData,proto3" json:"judgeData,omitempty"`
	TimeLimit   int64        `protobuf:"varint,1,opt,name=timeLimit,proto3" json:"timeLimit,omitempty"`
	MemoryLimit int64        `protobuf:"varint,2,opt,name=memoryLimit,proto3" json:"memoryLimit,omitempty"`
	Code        string       `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *JudgeWithDataReq) Reset() {
	*x = JudgeWithDataReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeWithDataReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeWithDataReq) ProtoMessage() {}

func (x *JudgeWithDataReq) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeWithDataReq.ProtoReflect.Descriptor instead.
func (*JudgeWithDataReq) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{3}
}

func (x *JudgeWithDataReq) GetJudgeData() []*JudgeData {
	if x != nil {
		return x.JudgeData
	}
	return nil
}

func (x *JudgeWithDataReq) GetTimeLimit() int64 {
	if x != nil {
		return x.TimeLimit
	}
	return 0
}

func (x *JudgeWithDataReq) GetMemoryLimit() int64 {
	if x != nil {
		return x.MemoryLimit
	}
	return 0
}

func (x *JudgeWithDataReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_judge_proto protoreflect.FileDescriptor

var file_judge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6a,
	0x75, 0x64, 0x67, 0x65, 0x22, 0x7c, 0x0a, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x6d, 0x0a, 0x0b, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x39, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x96, 0x01, 0x0a,
	0x10, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x12, 0x2e, 0x0a, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4a, 0x75, 0x64,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x77, 0x0a, 0x05, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x12, 0x2e,
	0x0a, 0x05, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e,
	0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x01, 0x12, 0x3e,
	0x0a, 0x0d, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x17, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x6a, 0x75, 0x64, 0x67, 0x65,
	0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x28, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_judge_proto_rawDescOnce sync.Once
	file_judge_proto_rawDescData = file_judge_proto_rawDesc
)

func file_judge_proto_rawDescGZIP() []byte {
	file_judge_proto_rawDescOnce.Do(func() {
		file_judge_proto_rawDescData = protoimpl.X.CompressGZIP(file_judge_proto_rawDescData)
	})
	return file_judge_proto_rawDescData
}

var file_judge_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_judge_proto_goTypes = []any{
	(*JudgeReq)(nil),         // 0: judge.JudgeReq
	(*JudgeResult)(nil),      // 1: judge.JudgeResult
	(*JudgeData)(nil),        // 2: judge.JudgeData
	(*JudgeWithDataReq)(nil), // 3: judge.JudgeWithDataReq
}
var file_judge_proto_depIdxs = []int32{
	2, // 0: judge.JudgeWithDataReq.judgeData:type_name -> judge.JudgeData
	0, // 1: judge.Judge.Judge:input_type -> judge.JudgeReq
	3, // 2: judge.Judge.JudgeWithData:input_type -> judge.JudgeWithDataReq
	1, // 3: judge.Judge.Judge:output_type -> judge.JudgeResult
	1, // 4: judge.Judge.JudgeWithData:output_type -> judge.JudgeResult
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_judge_proto_init() }
func file_judge_proto_init() {
	if File_judge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judge_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*JudgeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_judge_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*JudgeResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_judge_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*JudgeData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_judge_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*JudgeWithDataReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_judge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_judge_proto_goTypes,
		DependencyIndexes: file_judge_proto_depIdxs,
		MessageInfos:      file_judge_proto_msgTypes,
	}.Build()
	File_judge_proto = out.File
	file_judge_proto_rawDesc = nil
	file_judge_proto_goTypes = nil
	file_judge_proto_depIdxs = nil
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: mail.proto

package mail

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

type SendMailCheckCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// 用户操作，例如：注册账号、忘记密码等，该字段将会被嵌入发送的邮件中
	UserOperation string `protobuf:"bytes,2,opt,name=userOperation,proto3" json:"userOperation,omitempty"`
}

func (x *SendMailCheckCodeReq) Reset() {
	*x = SendMailCheckCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailCheckCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailCheckCodeReq) ProtoMessage() {}

func (x *SendMailCheckCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailCheckCodeReq.ProtoReflect.Descriptor instead.
func (*SendMailCheckCodeReq) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{0}
}

func (x *SendMailCheckCodeReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendMailCheckCodeReq) GetUserOperation() string {
	if x != nil {
		return x.UserOperation
	}
	return ""
}

type SendMailCheckCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckCode string `protobuf:"bytes,1,opt,name=checkCode,proto3" json:"checkCode,omitempty"`
}

func (x *SendMailCheckCodeResp) Reset() {
	*x = SendMailCheckCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailCheckCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailCheckCodeResp) ProtoMessage() {}

func (x *SendMailCheckCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_mail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailCheckCodeResp.ProtoReflect.Descriptor instead.
func (*SendMailCheckCodeResp) Descriptor() ([]byte, []int) {
	return file_mail_proto_rawDescGZIP(), []int{1}
}

func (x *SendMailCheckCodeResp) GetCheckCode() string {
	if x != nil {
		return x.CheckCode
	}
	return ""
}

var File_mail_proto protoreflect.FileDescriptor

var file_mail_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x52, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x24, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x59, 0x0a,
	0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6d, 0x61,
	0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mail_proto_rawDescOnce sync.Once
	file_mail_proto_rawDescData = file_mail_proto_rawDesc
)

func file_mail_proto_rawDescGZIP() []byte {
	file_mail_proto_rawDescOnce.Do(func() {
		file_mail_proto_rawDescData = protoimpl.X.CompressGZIP(file_mail_proto_rawDescData)
	})
	return file_mail_proto_rawDescData
}

var file_mail_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mail_proto_goTypes = []any{
	(*SendMailCheckCodeReq)(nil),  // 0: mail.SendMailCheckCodeReq
	(*SendMailCheckCodeResp)(nil), // 1: mail.SendMailCheckCodeResp
}
var file_mail_proto_depIdxs = []int32{
	0, // 0: mail.CheckCode.SendMailCheckCode:input_type -> mail.SendMailCheckCodeReq
	1, // 1: mail.CheckCode.SendMailCheckCode:output_type -> mail.SendMailCheckCodeResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mail_proto_init() }
func file_mail_proto_init() {
	if File_mail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailCheckCodeReq); i {
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
		file_mail_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SendMailCheckCodeResp); i {
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
			RawDescriptor: file_mail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mail_proto_goTypes,
		DependencyIndexes: file_mail_proto_depIdxs,
		MessageInfos:      file_mail_proto_msgTypes,
	}.Build()
	File_mail_proto = out.File
	file_mail_proto_rawDesc = nil
	file_mail_proto_goTypes = nil
	file_mail_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: sms/v1/simple_message.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 错误码的取值范围应该在 0 < code <= 600 之间, 超出范围将抛出异常
type ErrorReason int32

const (
	ErrorReason_NOT_FOUND ErrorReason = 0
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"NOT_FOUND": 0,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_sms_v1_simple_message_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_sms_v1_simple_message_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{0}
}

// 认证头
type Authentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AppId
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// AppId对应的访问令牌
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Authentication) Reset() {
	*x = Authentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Authentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Authentication) ProtoMessage() {}

func (x *Authentication) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Authentication.ProtoReflect.Descriptor instead.
func (*Authentication) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{0}
}

func (x *Authentication) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Authentication) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 文本消息
type TextMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 认证头
	Auth *Authentication `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	// 消息内容
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// 消息优先级
	Priority int32 `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	// 定时发送时间 （时间戳, 大于当前时间, 24小时以内）
	AtTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=at_time,json=atTime,proto3,oneof" json:"at_time,omitempty"` //[(validate.rules).timestamp = {gt_now: true, within: {seconds: 86400}}];
	// 手机号列表
	Phones []string `protobuf:"bytes,10,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *TextMessageRequest) Reset() {
	*x = TextMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessageRequest) ProtoMessage() {}

func (x *TextMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessageRequest.ProtoReflect.Descriptor instead.
func (*TextMessageRequest) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{1}
}

func (x *TextMessageRequest) GetAuth() *Authentication {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *TextMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextMessageRequest) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *TextMessageRequest) GetAtTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AtTime
	}
	return nil
}

func (x *TextMessageRequest) GetPhones() []string {
	if x != nil {
		return x.Phones
	}
	return nil
}

// 模板消息
type TemplateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 认证头
	Auth *Authentication `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	// 模板编号,如 "T00001", "TUOS01" 等
	TemplateId string `protobuf:"bytes,2,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	// 定时发送时间 （时间戳, 大于当前时间, 24小时以内）
	AtTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=at_time,json=atTime,proto3,oneof" json:"at_time,omitempty"` //[(validate.rules).timestamp = {gt_now: true, within: {seconds: 86400}}];
	// 手机号列表
	Phones []string `protobuf:"bytes,10,rep,name=phones,proto3" json:"phones,omitempty"`
	// 模板占位符对应参数 key=占位符，value=替换值，如key="acc_no" value="6222020105063281"
	Args map[string]string `protobuf:"bytes,11,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TemplateMessageRequest) Reset() {
	*x = TemplateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateMessageRequest) ProtoMessage() {}

func (x *TemplateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateMessageRequest.ProtoReflect.Descriptor instead.
func (*TemplateMessageRequest) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{2}
}

func (x *TemplateMessageRequest) GetAuth() *Authentication {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *TemplateMessageRequest) GetTemplateId() string {
	if x != nil {
		return x.TemplateId
	}
	return ""
}

func (x *TemplateMessageRequest) GetAtTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AtTime
	}
	return nil
}

func (x *TemplateMessageRequest) GetPhones() []string {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *TemplateMessageRequest) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

// 异步查询发送结果
type AsyncResultQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 认证头
	Auth *Authentication `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
	// 查询编号
	QueryId uint64 `protobuf:"varint,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *AsyncResultQueryRequest) Reset() {
	*x = AsyncResultQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncResultQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncResultQueryRequest) ProtoMessage() {}

func (x *AsyncResultQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncResultQueryRequest.ProtoReflect.Descriptor instead.
func (*AsyncResultQueryRequest) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{3}
}

func (x *AsyncResultQueryRequest) GetAuth() *Authentication {
	if x != nil {
		return x.Auth
	}
	return nil
}

func (x *AsyncResultQueryRequest) GetQueryId() uint64 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

// 响应消息公共状态头
type ReplyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码, 200 成功，其他失败
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 状态描述信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ReplyStatus) Reset() {
	*x = ReplyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyStatus) ProtoMessage() {}

func (x *ReplyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyStatus.ProtoReflect.Descriptor instead.
func (*ReplyStatus) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ReplyStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 发送消息的同步应答消息
type SendMessageReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应消息头
	Status *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 用于异步查询结果的查询编号
	QueryId string `protobuf:"bytes,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
}

func (x *SendMessageReply) Reset() {
	*x = SendMessageReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReply) ProtoMessage() {}

func (x *SendMessageReply) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReply.ProtoReflect.Descriptor instead.
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{5}
}

func (x *SendMessageReply) GetStatus() *ReplyStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SendMessageReply) GetQueryId() string {
	if x != nil {
		return x.QueryId
	}
	return ""
}

// 查询结果
type AsyncResultQueryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 响应消息头
	Status *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// 查询编号
	QueryId uint64 `protobuf:"varint,2,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	// 查询结果列表
	Results []*AsyncResultQueryReply_ResultList `protobuf:"bytes,10,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *AsyncResultQueryReply) Reset() {
	*x = AsyncResultQueryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncResultQueryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncResultQueryReply) ProtoMessage() {}

func (x *AsyncResultQueryReply) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncResultQueryReply.ProtoReflect.Descriptor instead.
func (*AsyncResultQueryReply) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{6}
}

func (x *AsyncResultQueryReply) GetStatus() *ReplyStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AsyncResultQueryReply) GetQueryId() uint64 {
	if x != nil {
		return x.QueryId
	}
	return 0
}

func (x *AsyncResultQueryReply) GetResults() []*AsyncResultQueryReply_ResultList {
	if x != nil {
		return x.Results
	}
	return nil
}

type AsyncResultQueryReply_ResultList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	// 短信发送流水号
	SequenceId uint64 `protobuf:"varint,2,opt,name=sequence_id,json=sequenceId,proto3" json:"sequence_id,omitempty"`
	// 运营商网关响应
	Result uint32 `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	// 运营商网关短信编号
	MsgId string `protobuf:"bytes,4,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	// 消息发送时间
	SendTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	// 运营商网关响应时间
	ResponseTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// 状态报告接收到的时间
	ReportTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=report_time,json=reportTime,proto3,oneof" json:"report_time,omitempty"`
	// 状态报告内容
	Report *string `protobuf:"bytes,8,opt,name=report,proto3,oneof" json:"report,omitempty"`
}

func (x *AsyncResultQueryReply_ResultList) Reset() {
	*x = AsyncResultQueryReply_ResultList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sms_v1_simple_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncResultQueryReply_ResultList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncResultQueryReply_ResultList) ProtoMessage() {}

func (x *AsyncResultQueryReply_ResultList) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_simple_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncResultQueryReply_ResultList.ProtoReflect.Descriptor instead.
func (*AsyncResultQueryReply_ResultList) Descriptor() ([]byte, []int) {
	return file_sms_v1_simple_message_proto_rawDescGZIP(), []int{6, 0}
}

func (x *AsyncResultQueryReply_ResultList) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AsyncResultQueryReply_ResultList) GetSequenceId() uint64 {
	if x != nil {
		return x.SequenceId
	}
	return 0
}

func (x *AsyncResultQueryReply_ResultList) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *AsyncResultQueryReply_ResultList) GetMsgId() string {
	if x != nil {
		return x.MsgId
	}
	return ""
}

func (x *AsyncResultQueryReply_ResultList) GetSendTime() *timestamppb.Timestamp {
	if x != nil {
		return x.SendTime
	}
	return nil
}

func (x *AsyncResultQueryReply_ResultList) GetResponseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ResponseTime
	}
	return nil
}

func (x *AsyncResultQueryReply_ResultList) GetReportTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ReportTime
	}
	return nil
}

func (x *AsyncResultQueryReply_ResultList) GetReport() string {
	if x != nil && x.Report != nil {
		return *x.Report
	}
	return ""
}

var File_sms_v1_simple_message_proto protoreflect.FileDescriptor

var file_sms_v1_simple_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x10, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x40, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf9, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x09, 0x28, 0x00, 0x52, 0x08, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x06, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x08, 0x01, 0x10, 0x63, 0x18, 0x01, 0x52, 0x06, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0xe4, 0x02, 0x0a, 0x16, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x06, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x48, 0x00, 0x52, 0x06, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0c, 0xfa,
	0x42, 0x09, 0x92, 0x01, 0x06, 0x08, 0x01, 0x10, 0x63, 0x18, 0x01, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x10, 0xfa, 0x42, 0x0d,
	0x9a, 0x01, 0x0a, 0x08, 0x01, 0x10, 0x10, 0x2a, 0x04, 0x72, 0x02, 0x18, 0x40, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x17, 0x41, 0x73, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12,
	0x22, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x5a, 0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x22, 0x8c, 0x04, 0x0a,
	0x15, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x64, 0x12, 0x42,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x1a, 0xe6, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x15, 0x0a, 0x06, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2a, 0x28, 0x0a, 0x0b, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x09, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a,
	0x04, 0xa0, 0x45, 0xf4, 0x03, 0x32, 0xd3, 0x02, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x62, 0x0a,
	0x0f, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x1a, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e,
	0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x6e, 0x0a, 0x13, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x73, 0x6d, 0x73,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0x78, 0x0a, 0x10, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x73, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f,
	0x73, 0x6d, 0x73, 0x2f, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x2f, 0x7b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x28, 0x0a, 0x0a, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x18, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2d, 0x73, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sms_v1_simple_message_proto_rawDescOnce sync.Once
	file_sms_v1_simple_message_proto_rawDescData = file_sms_v1_simple_message_proto_rawDesc
)

func file_sms_v1_simple_message_proto_rawDescGZIP() []byte {
	file_sms_v1_simple_message_proto_rawDescOnce.Do(func() {
		file_sms_v1_simple_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_v1_simple_message_proto_rawDescData)
	})
	return file_sms_v1_simple_message_proto_rawDescData
}

var file_sms_v1_simple_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sms_v1_simple_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sms_v1_simple_message_proto_goTypes = []interface{}{
	(ErrorReason)(0),                         // 0: sms.v1.ErrorReason
	(*Authentication)(nil),                   // 1: sms.v1.Authentication
	(*TextMessageRequest)(nil),               // 2: sms.v1.TextMessageRequest
	(*TemplateMessageRequest)(nil),           // 3: sms.v1.TemplateMessageRequest
	(*AsyncResultQueryRequest)(nil),          // 4: sms.v1.AsyncResultQueryRequest
	(*ReplyStatus)(nil),                      // 5: sms.v1.ReplyStatus
	(*SendMessageReply)(nil),                 // 6: sms.v1.SendMessageReply
	(*AsyncResultQueryReply)(nil),            // 7: sms.v1.AsyncResultQueryReply
	nil,                                      // 8: sms.v1.TemplateMessageRequest.ArgsEntry
	(*AsyncResultQueryReply_ResultList)(nil), // 9: sms.v1.AsyncResultQueryReply.ResultList
	(*timestamppb.Timestamp)(nil),            // 10: google.protobuf.Timestamp
}
var file_sms_v1_simple_message_proto_depIdxs = []int32{
	1,  // 0: sms.v1.TextMessageRequest.auth:type_name -> sms.v1.Authentication
	10, // 1: sms.v1.TextMessageRequest.at_time:type_name -> google.protobuf.Timestamp
	1,  // 2: sms.v1.TemplateMessageRequest.auth:type_name -> sms.v1.Authentication
	10, // 3: sms.v1.TemplateMessageRequest.at_time:type_name -> google.protobuf.Timestamp
	8,  // 4: sms.v1.TemplateMessageRequest.args:type_name -> sms.v1.TemplateMessageRequest.ArgsEntry
	1,  // 5: sms.v1.AsyncResultQueryRequest.auth:type_name -> sms.v1.Authentication
	5,  // 6: sms.v1.SendMessageReply.status:type_name -> sms.v1.ReplyStatus
	5,  // 7: sms.v1.AsyncResultQueryReply.status:type_name -> sms.v1.ReplyStatus
	9,  // 8: sms.v1.AsyncResultQueryReply.results:type_name -> sms.v1.AsyncResultQueryReply.ResultList
	10, // 9: sms.v1.AsyncResultQueryReply.ResultList.send_time:type_name -> google.protobuf.Timestamp
	10, // 10: sms.v1.AsyncResultQueryReply.ResultList.response_time:type_name -> google.protobuf.Timestamp
	10, // 11: sms.v1.AsyncResultQueryReply.ResultList.report_time:type_name -> google.protobuf.Timestamp
	2,  // 12: sms.v1.Sms.TextMessageSend:input_type -> sms.v1.TextMessageRequest
	3,  // 13: sms.v1.Sms.TemplateMessageSend:input_type -> sms.v1.TemplateMessageRequest
	4,  // 14: sms.v1.Sms.AsyncResultQuery:input_type -> sms.v1.AsyncResultQueryRequest
	6,  // 15: sms.v1.Sms.TextMessageSend:output_type -> sms.v1.SendMessageReply
	6,  // 16: sms.v1.Sms.TemplateMessageSend:output_type -> sms.v1.SendMessageReply
	7,  // 17: sms.v1.Sms.AsyncResultQuery:output_type -> sms.v1.AsyncResultQueryReply
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_sms_v1_simple_message_proto_init() }
func file_sms_v1_simple_message_proto_init() {
	if File_sms_v1_simple_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sms_v1_simple_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Authentication); i {
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
		file_sms_v1_simple_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMessageRequest); i {
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
		file_sms_v1_simple_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateMessageRequest); i {
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
		file_sms_v1_simple_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncResultQueryRequest); i {
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
		file_sms_v1_simple_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyStatus); i {
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
		file_sms_v1_simple_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReply); i {
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
		file_sms_v1_simple_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncResultQueryReply); i {
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
		file_sms_v1_simple_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncResultQueryReply_ResultList); i {
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
	file_sms_v1_simple_message_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_sms_v1_simple_message_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_sms_v1_simple_message_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sms_v1_simple_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_v1_simple_message_proto_goTypes,
		DependencyIndexes: file_sms_v1_simple_message_proto_depIdxs,
		EnumInfos:         file_sms_v1_simple_message_proto_enumTypes,
		MessageInfos:      file_sms_v1_simple_message_proto_msgTypes,
	}.Build()
	File_sms_v1_simple_message_proto = out.File
	file_sms_v1_simple_message_proto_rawDesc = nil
	file_sms_v1_simple_message_proto_goTypes = nil
	file_sms_v1_simple_message_proto_depIdxs = nil
}

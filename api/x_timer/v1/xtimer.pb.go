// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.1
// source: x_timer/v1/xtimer.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type NotifyHTTPParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url     string            `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Method  string            `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body    string            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *NotifyHTTPParam) Reset() {
	*x = NotifyHTTPParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyHTTPParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyHTTPParam) ProtoMessage() {}

func (x *NotifyHTTPParam) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyHTTPParam.ProtoReflect.Descriptor instead.
func (*NotifyHTTPParam) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{2}
}

func (x *NotifyHTTPParam) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *NotifyHTTPParam) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *NotifyHTTPParam) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *NotifyHTTPParam) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreateTimerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App             string           `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Name            string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status          int32            `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Cron            string           `protobuf:"bytes,4,opt,name=cron,proto3" json:"cron,omitempty"`
	NotifyHTTPParam *NotifyHTTPParam `protobuf:"bytes,5,opt,name=notifyHTTPParam,proto3" json:"notifyHTTPParam,omitempty"`
}

func (x *CreateTimerRequest) Reset() {
	*x = CreateTimerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTimerRequest) ProtoMessage() {}

func (x *CreateTimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTimerRequest.ProtoReflect.Descriptor instead.
func (*CreateTimerRequest) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTimerRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *CreateTimerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTimerRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateTimerRequest) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *CreateTimerRequest) GetNotifyHTTPParam() *NotifyHTTPParam {
	if x != nil {
		return x.NotifyHTTPParam
	}
	return nil
}

type CreateTimerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTimerReply) Reset() {
	*x = CreateTimerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTimerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTimerReply) ProtoMessage() {}

func (x *CreateTimerReply) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTimerReply.ProtoReflect.Descriptor instead.
func (*CreateTimerReply) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTimerReply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ActiveTimerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	App    string `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Id     int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ActiveTimerRequest) Reset() {
	*x = ActiveTimerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveTimerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveTimerRequest) ProtoMessage() {}

func (x *ActiveTimerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveTimerRequest.ProtoReflect.Descriptor instead.
func (*ActiveTimerRequest) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{5}
}

func (x *ActiveTimerRequest) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *ActiveTimerRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActiveTimerRequest) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ActiveTimerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ActiveTimerReply) Reset() {
	*x = ActiveTimerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_x_timer_v1_xtimer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveTimerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveTimerReply) ProtoMessage() {}

func (x *ActiveTimerReply) ProtoReflect() protoreflect.Message {
	mi := &file_x_timer_v1_xtimer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveTimerReply.ProtoReflect.Descriptor instead.
func (*ActiveTimerReply) Descriptor() ([]byte, []int) {
	return file_x_timer_v1_xtimer_proto_rawDescGZIP(), []int{6}
}

func (x *ActiveTimerReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ActiveTimerReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_x_timer_v1_xtimer_proto protoreflect.FileDescriptor

var file_x_timer_v1_xtimer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x78, 0x74, 0x69,
	0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x78, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xcf, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x42, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xad, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x48, 0x54, 0x54, 0x50, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x22, 0x22, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x70, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3c, 0x0a, 0x10, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xa8, 0x02, 0x0a, 0x06, 0x58, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x52,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x78, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x64, 0x0a, 0x0b, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a,
	0x22, 0x0c, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x42, 0x1a,
	0x5a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_x_timer_v1_xtimer_proto_rawDescOnce sync.Once
	file_x_timer_v1_xtimer_proto_rawDescData = file_x_timer_v1_xtimer_proto_rawDesc
)

func file_x_timer_v1_xtimer_proto_rawDescGZIP() []byte {
	file_x_timer_v1_xtimer_proto_rawDescOnce.Do(func() {
		file_x_timer_v1_xtimer_proto_rawDescData = protoimpl.X.CompressGZIP(file_x_timer_v1_xtimer_proto_rawDescData)
	})
	return file_x_timer_v1_xtimer_proto_rawDescData
}

var file_x_timer_v1_xtimer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_x_timer_v1_xtimer_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),       // 0: x_timer.v1.HelloRequest
	(*HelloReply)(nil),         // 1: x_timer.v1.HelloReply
	(*NotifyHTTPParam)(nil),    // 2: x_timer.v1.NotifyHTTPParam
	(*CreateTimerRequest)(nil), // 3: x_timer.v1.CreateTimerRequest
	(*CreateTimerReply)(nil),   // 4: x_timer.v1.CreateTimerReply
	(*ActiveTimerRequest)(nil), // 5: x_timer.v1.ActiveTimerRequest
	(*ActiveTimerReply)(nil),   // 6: x_timer.v1.ActiveTimerReply
	nil,                        // 7: x_timer.v1.NotifyHTTPParam.HeadersEntry
}
var file_x_timer_v1_xtimer_proto_depIdxs = []int32{
	7, // 0: x_timer.v1.NotifyHTTPParam.headers:type_name -> x_timer.v1.NotifyHTTPParam.HeadersEntry
	2, // 1: x_timer.v1.CreateTimerRequest.notifyHTTPParam:type_name -> x_timer.v1.NotifyHTTPParam
	0, // 2: x_timer.v1.XTimer.SayHello:input_type -> x_timer.v1.HelloRequest
	3, // 3: x_timer.v1.XTimer.CreateTimer:input_type -> x_timer.v1.CreateTimerRequest
	5, // 4: x_timer.v1.XTimer.ActiveTimer:input_type -> x_timer.v1.ActiveTimerRequest
	1, // 5: x_timer.v1.XTimer.SayHello:output_type -> x_timer.v1.HelloReply
	4, // 6: x_timer.v1.XTimer.CreateTimer:output_type -> x_timer.v1.CreateTimerReply
	6, // 7: x_timer.v1.XTimer.ActiveTimer:output_type -> x_timer.v1.ActiveTimerReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_x_timer_v1_xtimer_proto_init() }
func file_x_timer_v1_xtimer_proto_init() {
	if File_x_timer_v1_xtimer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_x_timer_v1_xtimer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyHTTPParam); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTimerRequest); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTimerReply); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveTimerRequest); i {
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
		file_x_timer_v1_xtimer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveTimerReply); i {
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
			RawDescriptor: file_x_timer_v1_xtimer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_x_timer_v1_xtimer_proto_goTypes,
		DependencyIndexes: file_x_timer_v1_xtimer_proto_depIdxs,
		MessageInfos:      file_x_timer_v1_xtimer_proto_msgTypes,
	}.Build()
	File_x_timer_v1_xtimer_proto = out.File
	file_x_timer_v1_xtimer_proto_rawDesc = nil
	file_x_timer_v1_xtimer_proto_goTypes = nil
	file_x_timer_v1_xtimer_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: synergy.proto

package synergy

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RunnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Task string `protobuf:"bytes,2,opt,name=Task,proto3" json:"Task,omitempty"`
}

func (x *RunnerRequest) Reset() {
	*x = RunnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerRequest) ProtoMessage() {}

func (x *RunnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerRequest.ProtoReflect.Descriptor instead.
func (*RunnerRequest) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{0}
}

func (x *RunnerRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RunnerRequest) GetTask() string {
	if x != nil {
		return x.Task
	}
	return ""
}

type RunnerUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Factor int64  `protobuf:"varint,2,opt,name=Factor,proto3" json:"Factor,omitempty"`
}

func (x *RunnerUpdate) Reset() {
	*x = RunnerUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerUpdate) ProtoMessage() {}

func (x *RunnerUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerUpdate.ProtoReflect.Descriptor instead.
func (*RunnerUpdate) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{1}
}

func (x *RunnerUpdate) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RunnerUpdate) GetFactor() int64 {
	if x != nil {
		return x.Factor
	}
	return 0
}

type RunnerOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *RunnerOutput) Reset() {
	*x = RunnerOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunnerOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunnerOutput) ProtoMessage() {}

func (x *RunnerOutput) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunnerOutput.ProtoReflect.Descriptor instead.
func (*RunnerOutput) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{2}
}

func (x *RunnerOutput) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SvcStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Healtz bool   `protobuf:"varint,2,opt,name=Healtz,proto3" json:"Healtz,omitempty"`
	State  string `protobuf:"bytes,3,opt,name=State,proto3" json:"State,omitempty"`
	Error  string `protobuf:"bytes,4,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *SvcStat) Reset() {
	*x = SvcStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvcStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvcStat) ProtoMessage() {}

func (x *SvcStat) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvcStat.ProtoReflect.Descriptor instead.
func (*SvcStat) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{3}
}

func (x *SvcStat) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *SvcStat) GetHealtz() bool {
	if x != nil {
		return x.Healtz
	}
	return false
}

func (x *SvcStat) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *SvcStat) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SvrStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string     `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Load      int64      `protobuf:"varint,2,opt,name=Load,proto3" json:"Load,omitempty"`
	Release   string     `protobuf:"bytes,9,opt,name=Release,proto3" json:"Release,omitempty"`
	Instances []*SvcStat `protobuf:"bytes,3,rep,name=Instances,proto3" json:"Instances,omitempty"`
}

func (x *SvrStat) Reset() {
	*x = SvrStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SvrStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SvrStat) ProtoMessage() {}

func (x *SvrStat) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SvrStat.ProtoReflect.Descriptor instead.
func (*SvrStat) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{4}
}

func (x *SvrStat) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SvrStat) GetLoad() int64 {
	if x != nil {
		return x.Load
	}
	return 0
}

func (x *SvrStat) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *SvrStat) GetInstances() []*SvcStat {
	if x != nil {
		return x.Instances
	}
	return nil
}

type OprStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *OprStat) Reset() {
	*x = OprStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OprStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OprStat) ProtoMessage() {}

func (x *OprStat) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OprStat.ProtoReflect.Descriptor instead.
func (*OprStat) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{5}
}

func (x *OprStat) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_synergy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_synergy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_synergy_proto_rawDescGZIP(), []int{6}
}

func (x *Request) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

var File_synergy_proto protoreflect.FileDescriptor

var file_synergy_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x22, 0x33, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x36, 0x0a,
	0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x46,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x5d, 0x0a, 0x07, 0x53, 0x76, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x7a, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x7a, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x7b, 0x0a, 0x07, 0x53, 0x76, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x2e, 0x53, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x52, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x07, 0x4f, 0x70, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0xdf, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x7a, 0x12, 0x10,
	0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x53, 0x76, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x53, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72,
	0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65,
	0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70, 0x72, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x0a, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x79,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e,
	0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70, 0x72, 0x53, 0x74, 0x61, 0x74, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x53, 0x76, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x32, 0xa1, 0x03, 0x0a, 0x06, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x12, 0x2e, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x7a, 0x12, 0x10, 0x2e, 0x73,
	0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x53, 0x76, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x7a, 0x12, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x2e, 0x53, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79,
	0x2e, 0x53, 0x76, 0x63, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x79,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x1a, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e,
	0x6e, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x38, 0x0a,
	0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73,
	0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e,
	0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x73,
	0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x4f, 0x70, 0x72, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72,
	0x12, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67,
	0x79, 0x2e, 0x4f, 0x70, 0x72, 0x53, 0x74, 0x61, 0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_synergy_proto_rawDescOnce sync.Once
	file_synergy_proto_rawDescData = file_synergy_proto_rawDesc
)

func file_synergy_proto_rawDescGZIP() []byte {
	file_synergy_proto_rawDescOnce.Do(func() {
		file_synergy_proto_rawDescData = protoimpl.X.CompressGZIP(file_synergy_proto_rawDescData)
	})
	return file_synergy_proto_rawDescData
}

var file_synergy_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_synergy_proto_goTypes = []interface{}{
	(*RunnerRequest)(nil), // 0: synergy.RunnerRequest
	(*RunnerUpdate)(nil),  // 1: synergy.RunnerUpdate
	(*RunnerOutput)(nil),  // 2: synergy.RunnerOutput
	(*SvcStat)(nil),       // 3: synergy.SvcStat
	(*SvrStat)(nil),       // 4: synergy.SvrStat
	(*OprStat)(nil),       // 5: synergy.OprStat
	(*Request)(nil),       // 6: synergy.Request
}
var file_synergy_proto_depIdxs = []int32{
	3,  // 0: synergy.SvrStat.Instances:type_name -> synergy.SvcStat
	6,  // 1: synergy.Controller.Healtz:input_type -> synergy.Request
	0,  // 2: synergy.Controller.CreateRunner:input_type -> synergy.RunnerRequest
	1,  // 3: synergy.Controller.ConnectRunner:input_type -> synergy.RunnerUpdate
	1,  // 4: synergy.Controller.JoinRunner:input_type -> synergy.RunnerUpdate
	1,  // 5: synergy.Controller.StartRunner:input_type -> synergy.RunnerUpdate
	1,  // 6: synergy.Controller.StopRunner:input_type -> synergy.RunnerUpdate
	1,  // 7: synergy.Controller.UpdateRunner:input_type -> synergy.RunnerUpdate
	4,  // 8: synergy.Controller.RegistWorker:input_type -> synergy.SvrStat
	6,  // 9: synergy.Worker.Healtz:input_type -> synergy.Request
	0,  // 10: synergy.Worker.RunnerHealtz:input_type -> synergy.RunnerRequest
	0,  // 11: synergy.Worker.CreateRunner:input_type -> synergy.RunnerRequest
	1,  // 12: synergy.Worker.ConnectRunner:input_type -> synergy.RunnerUpdate
	1,  // 13: synergy.Worker.StartRunner:input_type -> synergy.RunnerUpdate
	1,  // 14: synergy.Worker.StopRunner:input_type -> synergy.RunnerUpdate
	1,  // 15: synergy.Worker.UpdateRunner:input_type -> synergy.RunnerUpdate
	4,  // 16: synergy.Controller.Healtz:output_type -> synergy.SvrStat
	3,  // 17: synergy.Controller.CreateRunner:output_type -> synergy.SvcStat
	2,  // 18: synergy.Controller.ConnectRunner:output_type -> synergy.RunnerOutput
	2,  // 19: synergy.Controller.JoinRunner:output_type -> synergy.RunnerOutput
	5,  // 20: synergy.Controller.StartRunner:output_type -> synergy.OprStat
	5,  // 21: synergy.Controller.StopRunner:output_type -> synergy.OprStat
	5,  // 22: synergy.Controller.UpdateRunner:output_type -> synergy.OprStat
	5,  // 23: synergy.Controller.RegistWorker:output_type -> synergy.OprStat
	4,  // 24: synergy.Worker.Healtz:output_type -> synergy.SvrStat
	3,  // 25: synergy.Worker.RunnerHealtz:output_type -> synergy.SvcStat
	3,  // 26: synergy.Worker.CreateRunner:output_type -> synergy.SvcStat
	2,  // 27: synergy.Worker.ConnectRunner:output_type -> synergy.RunnerOutput
	5,  // 28: synergy.Worker.StartRunner:output_type -> synergy.OprStat
	5,  // 29: synergy.Worker.StopRunner:output_type -> synergy.OprStat
	5,  // 30: synergy.Worker.UpdateRunner:output_type -> synergy.OprStat
	16, // [16:31] is the sub-list for method output_type
	1,  // [1:16] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_synergy_proto_init() }
func file_synergy_proto_init() {
	if File_synergy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_synergy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnerRequest); i {
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
		file_synergy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnerUpdate); i {
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
		file_synergy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunnerOutput); i {
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
		file_synergy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvcStat); i {
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
		file_synergy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SvrStat); i {
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
		file_synergy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OprStat); i {
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
		file_synergy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_synergy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_synergy_proto_goTypes,
		DependencyIndexes: file_synergy_proto_depIdxs,
		MessageInfos:      file_synergy_proto_msgTypes,
	}.Build()
	File_synergy_proto = out.File
	file_synergy_proto_rawDesc = nil
	file_synergy_proto_goTypes = nil
	file_synergy_proto_depIdxs = nil
}

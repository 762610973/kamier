//! 定义节点间grpc通信的接口

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: node.proto

package node

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

type Pid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Serial   int64  `protobuf:"varint,2,opt,name=serial,proto3" json:"serial,omitempty"`
}

func (x *Pid) Reset() {
	*x = Pid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pid) ProtoMessage() {}

func (x *Pid) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pid.ProtoReflect.Descriptor instead.
func (*Pid) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *Pid) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Pid) GetSerial() int64 {
	if x != nil {
		return x.Serial
	}
	return 0
}

// 节点间通信请求参数
type IpcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Serial   int64  `protobuf:"varint,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Arg      []byte `protobuf:"bytes,3,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *IpcReq) Reset() {
	*x = IpcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpcReq) ProtoMessage() {}

func (x *IpcReq) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpcReq.ProtoReflect.Descriptor instead.
func (*IpcReq) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *IpcReq) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *IpcReq) GetSerial() int64 {
	if x != nil {
		return x.Serial
	}
	return 0
}

func (x *IpcReq) GetArg() []byte {
	if x != nil {
		return x.Arg
	}
	return nil
}

// 节点间通信相应
type IpcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IpcRes) Reset() {
	*x = IpcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpcRes) ProtoMessage() {}

func (x *IpcRes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpcRes.ProtoReflect.Descriptor instead.
func (*IpcRes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

// 准备请求
type PrepareReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 准备时，校验一下本节点是否在参与计算的列表里面
	Members []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *PrepareReq) Reset() {
	*x = PrepareReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareReq) ProtoMessage() {}

func (x *PrepareReq) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareReq.ProtoReflect.Descriptor instead.
func (*PrepareReq) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *PrepareReq) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

// 准备响应
type PrepareRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PrepareRes) Reset() {
	*x = PrepareRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrepareRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrepareRes) ProtoMessage() {}

func (x *PrepareRes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrepareRes.ProtoReflect.Descriptor instead.
func (*PrepareRes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

// 启动请求参数
type StartReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FuncId  string   `protobuf:"bytes,1,opt,name=funcId,proto3" json:"funcId,omitempty"`
	Pid     *Pid     `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Members []string `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *StartReq) Reset() {
	*x = StartReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReq) ProtoMessage() {}

func (x *StartReq) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReq.ProtoReflect.Descriptor instead.
func (*StartReq) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *StartReq) GetFuncId() string {
	if x != nil {
		return x.FuncId
	}
	return ""
}

func (x *StartReq) GetPid() *Pid {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *StartReq) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

// 启动响应
type StartRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartRes) Reset() {
	*x = StartRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRes) ProtoMessage() {}

func (x *StartRes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRes.ProtoReflect.Descriptor instead.
func (*StartRes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

// 获取值的请求参数
type FetchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid        *Pid   `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	TargetName string `protobuf:"bytes,2,opt,name=targetName,proto3" json:"targetName,omitempty"`
	SourceName string `protobuf:"bytes,3,opt,name=sourceName,proto3" json:"sourceName,omitempty"`
	Step       int64  `protobuf:"varint,4,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *FetchReq) Reset() {
	*x = FetchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchReq) ProtoMessage() {}

func (x *FetchReq) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchReq.ProtoReflect.Descriptor instead.
func (*FetchReq) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *FetchReq) GetPid() *Pid {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *FetchReq) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *FetchReq) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *FetchReq) GetStep() int64 {
	if x != nil {
		return x.Step
	}
	return 0
}

// 获取值得响应
type FetchRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res []byte `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *FetchRes) Reset() {
	*x = FetchRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRes) ProtoMessage() {}

func (x *FetchRes) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRes.ProtoReflect.Descriptor instead.
func (*FetchRes) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{8}
}

func (x *FetchRes) GetRes() []byte {
	if x != nil {
		return x.Res
	}
	return nil
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x22, 0x39, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x4e, 0x0a,
	0x06, 0x49, 0x70, 0x63, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x72, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x61, 0x72, 0x67, 0x22, 0x08, 0x0a,
	0x06, 0x49, 0x70, 0x63, 0x52, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22,
	0x0c, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x22, 0x59, 0x0a,
	0x08, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x75, 0x6e,
	0x63, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x75, 0x6e, 0x63, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x69, 0x64, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x0a, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x7b, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x12, 0x1b, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x50, 0x69, 0x64, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x74, 0x65,
	0x70, 0x22, 0x1c, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x65, 0x73, 0x32,
	0xb1, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2d, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12, 0x10, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x12, 0x27,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x03, 0x49, 0x70, 0x63, 0x12, 0x0c,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x49, 0x70, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x49, 0x70, 0x63, 0x52, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x12, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_node_proto_goTypes = []interface{}{
	(*Pid)(nil),        // 0: node.Pid
	(*IpcReq)(nil),     // 1: node.IpcReq
	(*IpcRes)(nil),     // 2: node.IpcRes
	(*PrepareReq)(nil), // 3: node.PrepareReq
	(*PrepareRes)(nil), // 4: node.PrepareRes
	(*StartReq)(nil),   // 5: node.StartReq
	(*StartRes)(nil),   // 6: node.StartRes
	(*FetchReq)(nil),   // 7: node.FetchReq
	(*FetchRes)(nil),   // 8: node.FetchRes
}
var file_node_proto_depIdxs = []int32{
	0, // 0: node.StartReq.pid:type_name -> node.Pid
	0, // 1: node.FetchReq.pid:type_name -> node.Pid
	3, // 2: node.NodeService.Prepare:input_type -> node.PrepareReq
	5, // 3: node.NodeService.Start:input_type -> node.StartReq
	1, // 4: node.NodeService.Ipc:input_type -> node.IpcReq
	7, // 5: node.NodeService.Fetch:input_type -> node.FetchReq
	4, // 6: node.NodeService.Prepare:output_type -> node.PrepareRes
	6, // 7: node.NodeService.Start:output_type -> node.StartRes
	2, // 8: node.NodeService.Ipc:output_type -> node.IpcRes
	8, // 9: node.NodeService.Fetch:output_type -> node.FetchRes
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pid); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpcReq); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpcRes); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareReq); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrepareRes); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReq); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRes); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchReq); i {
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
		file_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRes); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: cb.proto

package hzcb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Deps     string `protobuf:"bytes,2,opt,name=deps,proto3" json:"deps,omitempty"`
	Upstream int32  `protobuf:"varint,3,opt,name=upstream,proto3" json:"upstream,omitempty"`
}

func (x *ClientReadRequest) Reset() {
	*x = ClientReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReadRequest) ProtoMessage() {}

func (x *ClientReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReadRequest.ProtoReflect.Descriptor instead.
func (*ClientReadRequest) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{0}
}

func (x *ClientReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ClientReadRequest) GetDeps() string {
	if x != nil {
		return x.Deps
	}
	return ""
}

func (x *ClientReadRequest) GetUpstream() int32 {
	if x != nil {
		return x.Upstream
	}
	return 0
}

type ClientReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Vc    string `protobuf:"bytes,2,opt,name=vc,proto3" json:"vc,omitempty"`
}

func (x *ClientReadResponse) Reset() {
	*x = ClientReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReadResponse) ProtoMessage() {}

func (x *ClientReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReadResponse.ProtoReflect.Descriptor instead.
func (*ClientReadResponse) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{1}
}

func (x *ClientReadResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ClientReadResponse) GetVc() string {
	if x != nil {
		return x.Vc
	}
	return ""
}

type ServerReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ServerReadRequest) Reset() {
	*x = ServerReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReadRequest) ProtoMessage() {}

func (x *ServerReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReadRequest.ProtoReflect.Descriptor instead.
func (*ServerReadRequest) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{2}
}

func (x *ServerReadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ServerReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Vc    string `protobuf:"bytes,2,opt,name=vc,proto3" json:"vc,omitempty"`
}

func (x *ServerReadResponse) Reset() {
	*x = ServerReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReadResponse) ProtoMessage() {}

func (x *ServerReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReadResponse.ProtoReflect.Descriptor instead.
func (*ServerReadResponse) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{3}
}

func (x *ServerReadResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ServerReadResponse) GetVc() string {
	if x != nil {
		return x.Vc
	}
	return ""
}

type ClientWriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Deps  string `protobuf:"bytes,3,opt,name=deps,proto3" json:"deps,omitempty"`
}

func (x *ClientWriteRequest) Reset() {
	*x = ClientWriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientWriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientWriteRequest) ProtoMessage() {}

func (x *ClientWriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientWriteRequest.ProtoReflect.Descriptor instead.
func (*ClientWriteRequest) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{4}
}

func (x *ClientWriteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ClientWriteRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ClientWriteRequest) GetDeps() string {
	if x != nil {
		return x.Deps
	}
	return ""
}

type ClientWriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vc string `protobuf:"bytes,1,opt,name=vc,proto3" json:"vc,omitempty"`
}

func (x *ClientWriteResponse) Reset() {
	*x = ClientWriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientWriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientWriteResponse) ProtoMessage() {}

func (x *ClientWriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientWriteResponse.ProtoReflect.Descriptor instead.
func (*ClientWriteResponse) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{5}
}

func (x *ClientWriteResponse) GetVc() string {
	if x != nil {
		return x.Vc
	}
	return ""
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{6}
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_cb_proto_rawDescGZIP(), []int{7}
}

func (x *HealthCheckResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_cb_proto protoreflect.FileDescriptor

var file_cb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x62, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x11, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x70, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x3a, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x76, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x63, 0x22, 0x25,
	0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76,
	0x63, 0x22, 0x50, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x70, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x63, 0x22, 0x14, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2d, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0x88, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x73, 0x68, 0x12, 0x40, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x63, 0x62, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x63, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x63, 0x62, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x63, 0x62, 0x2e, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x62, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x69, 0x61, 0x63, 0x2f, 0x68,
	0x7a, 0x63, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cb_proto_rawDescOnce sync.Once
	file_cb_proto_rawDescData = file_cb_proto_rawDesc
)

func file_cb_proto_rawDescGZIP() []byte {
	file_cb_proto_rawDescOnce.Do(func() {
		file_cb_proto_rawDescData = protoimpl.X.CompressGZIP(file_cb_proto_rawDescData)
	})
	return file_cb_proto_rawDescData
}

var file_cb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cb_proto_goTypes = []interface{}{
	(*ClientReadRequest)(nil),   // 0: cb.ClientReadRequest
	(*ClientReadResponse)(nil),  // 1: cb.ClientReadResponse
	(*ServerReadRequest)(nil),   // 2: cb.ServerReadRequest
	(*ServerReadResponse)(nil),  // 3: cb.ServerReadResponse
	(*ClientWriteRequest)(nil),  // 4: cb.ClientWriteRequest
	(*ClientWriteResponse)(nil), // 5: cb.ClientWriteResponse
	(*HealthCheckRequest)(nil),  // 6: cb.HealthCheckRequest
	(*HealthCheckResponse)(nil), // 7: cb.HealthCheckResponse
}
var file_cb_proto_depIdxs = []int32{
	6, // 0: cb.Mesh.HealthCheck:input_type -> cb.HealthCheckRequest
	2, // 1: cb.Mesh.ServerRead:input_type -> cb.ServerReadRequest
	0, // 2: cb.Mesh.ClientRead:input_type -> cb.ClientReadRequest
	4, // 3: cb.Mesh.ClientWrite:input_type -> cb.ClientWriteRequest
	7, // 4: cb.Mesh.HealthCheck:output_type -> cb.HealthCheckResponse
	3, // 5: cb.Mesh.ServerRead:output_type -> cb.ServerReadResponse
	1, // 6: cb.Mesh.ClientRead:output_type -> cb.ClientReadResponse
	5, // 7: cb.Mesh.ClientWrite:output_type -> cb.ClientWriteResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cb_proto_init() }
func file_cb_proto_init() {
	if File_cb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReadRequest); i {
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
		file_cb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReadResponse); i {
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
		file_cb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReadRequest); i {
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
		file_cb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReadResponse); i {
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
		file_cb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientWriteRequest); i {
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
		file_cb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientWriteResponse); i {
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
		file_cb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_cb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
			RawDescriptor: file_cb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cb_proto_goTypes,
		DependencyIndexes: file_cb_proto_depIdxs,
		MessageInfos:      file_cb_proto_msgTypes,
	}.Build()
	File_cb_proto = out.File
	file_cb_proto_rawDesc = nil
	file_cb_proto_goTypes = nil
	file_cb_proto_depIdxs = nil
}

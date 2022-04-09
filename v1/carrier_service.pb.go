// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: carrier_service.proto

package v1

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"hi,omitempty"
	Hi string `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetHi() string {
	if x != nil {
		return x.Hi
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"ids,omitempty"
	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{3}
}

func (x *Ids) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_carrier_service_proto protoreflect.FileDescriptor

var file_carrier_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x62,
	0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x69,
	0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x14, 0x0a, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x92, 0x07, 0x0a,
	0x0e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x2d, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x09, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x73, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0d,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x08, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f,
	0x6b, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12,
	0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x09, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x69, 0x64, 0x73, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x12, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x23,
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f,
	0x6b, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x64, 0x73, 0x1a, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x1a, 0x0d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x2b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x1a, 0x14, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x69, 0x64, 0x1a, 0x08,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x69, 0x64, 0x73, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22,
	0x00, 0x42, 0x71, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x13, 0x43,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carrier_service_proto_rawDescOnce sync.Once
	file_carrier_service_proto_rawDescData = file_carrier_service_proto_rawDesc
)

func file_carrier_service_proto_rawDescGZIP() []byte {
	file_carrier_service_proto_rawDescOnce.Do(func() {
		file_carrier_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_carrier_service_proto_rawDescData)
	})
	return file_carrier_service_proto_rawDescData
}

var file_carrier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_carrier_service_proto_goTypes = []interface{}{
	(*Ping)(nil),          // 0: user.Ping
	(*Token)(nil),         // 1: user.Token
	(*Id)(nil),            // 2: user.id
	(*Ids)(nil),           // 3: user.ids
	(*Business)(nil),      // 4: user.business
	(*Empty)(nil),         // 5: user.empty
	(*Location)(nil),      // 6: user.location
	(*Bid)(nil),           // 7: user.bid
	(*Bids)(nil),          // 8: user.bids
	(*Quote)(nil),         // 9: user.quote
	(*Ok)(nil),            // 10: user.ok
	(*Businesses)(nil),    // 11: user.businesses
	(*Locations)(nil),     // 12: user.locations
	(*Booking)(nil),       // 13: user.booking
	(*Bookings)(nil),      // 14: user.bookings
	(*Quotes)(nil),        // 15: user.quotes
	(*QuoteResponse)(nil), // 16: user.quote_response
}
var file_carrier_service_proto_depIdxs = []int32{
	4,  // 0: user.CarrierService.CreateBusiness:input_type -> user.business
	2,  // 1: user.CarrierService.GetBusinessById:input_type -> user.id
	3,  // 2: user.CarrierService.GetBusinesses:input_type -> user.ids
	2,  // 3: user.CarrierService.CloseBusiness:input_type -> user.id
	5,  // 4: user.CarrierService.GetLocations:input_type -> user.empty
	5,  // 5: user.CarrierService.GetLocation:input_type -> user.empty
	6,  // 6: user.CarrierService.CreateLocation:input_type -> user.location
	6,  // 7: user.CarrierService.UpdateLocation:input_type -> user.location
	2,  // 8: user.CarrierService.DeleteLocation:input_type -> user.id
	3,  // 9: user.CarrierService.DeleteLocations:input_type -> user.ids
	5,  // 10: user.CarrierService.PingCarrierService:input_type -> user.empty
	7,  // 11: user.CarrierService.BookAQuote:input_type -> user.bid
	8,  // 12: user.CarrierService.BookQuotes:input_type -> user.bids
	5,  // 13: user.CarrierService.GetBookingHistory:input_type -> user.empty
	2,  // 14: user.CarrierService.GetBookingById:input_type -> user.id
	5,  // 15: user.CarrierService.GetAllQuotes:input_type -> user.empty
	9,  // 16: user.CarrierService.GetNewQuotes:input_type -> user.quote
	9,  // 17: user.CarrierService.UpdateQuote:input_type -> user.quote
	2,  // 18: user.CarrierService.DeleteQuote:input_type -> user.id
	3,  // 19: user.CarrierService.DeleteQuotes:input_type -> user.ids
	10, // 20: user.CarrierService.CreateBusiness:output_type -> user.ok
	4,  // 21: user.CarrierService.GetBusinessById:output_type -> user.business
	11, // 22: user.CarrierService.GetBusinesses:output_type -> user.businesses
	10, // 23: user.CarrierService.CloseBusiness:output_type -> user.ok
	6,  // 24: user.CarrierService.GetLocations:output_type -> user.location
	12, // 25: user.CarrierService.GetLocation:output_type -> user.locations
	10, // 26: user.CarrierService.CreateLocation:output_type -> user.ok
	10, // 27: user.CarrierService.UpdateLocation:output_type -> user.ok
	10, // 28: user.CarrierService.DeleteLocation:output_type -> user.ok
	10, // 29: user.CarrierService.DeleteLocations:output_type -> user.ok
	10, // 30: user.CarrierService.PingCarrierService:output_type -> user.ok
	10, // 31: user.CarrierService.BookAQuote:output_type -> user.ok
	13, // 32: user.CarrierService.BookQuotes:output_type -> user.booking
	14, // 33: user.CarrierService.GetBookingHistory:output_type -> user.bookings
	13, // 34: user.CarrierService.GetBookingById:output_type -> user.booking
	15, // 35: user.CarrierService.GetAllQuotes:output_type -> user.quotes
	16, // 36: user.CarrierService.GetNewQuotes:output_type -> user.quote_response
	16, // 37: user.CarrierService.UpdateQuote:output_type -> user.quote_response
	10, // 38: user.CarrierService.DeleteQuote:output_type -> user.ok
	10, // 39: user.CarrierService.DeleteQuotes:output_type -> user.ok
	20, // [20:40] is the sub-list for method output_type
	0,  // [0:20] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_carrier_service_proto_init() }
func file_carrier_service_proto_init() {
	if File_carrier_service_proto != nil {
		return
	}
	file_quote_proto_init()
	file_location_proto_init()
	file_book_proto_init()
	file_ok_proto_init()
	file_business_proto_init()
	file_empty_proto_init()
	file_bid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_carrier_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_carrier_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_carrier_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_carrier_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
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
			RawDescriptor: file_carrier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carrier_service_proto_goTypes,
		DependencyIndexes: file_carrier_service_proto_depIdxs,
		MessageInfos:      file_carrier_service_proto_msgTypes,
	}.Build()
	File_carrier_service_proto = out.File
	file_carrier_service_proto_rawDesc = nil
	file_carrier_service_proto_goTypes = nil
	file_carrier_service_proto_depIdxs = nil
}

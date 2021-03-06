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
	Hi string `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty" dynamodbav:"hi,omitempty"`
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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{1}
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
	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty" dynamodbav:"ids,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{2}
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
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x0b, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x68,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x68, 0x69, 0x22, 0x14, 0x0a, 0x02, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x15, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xcf, 0x07, 0x0a, 0x0e, 0x43, 0x61, 0x72,
	0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x0c, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x06, 0x2e, 0x76, 0x31,
	0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64,
	0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x12, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x73, 0x1a, 0x0e, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0d,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x06, 0x2e,
	0x76, 0x31, 0x2e, 0x69, 0x64, 0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x41, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a,
	0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x1a, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x06, 0x2e, 0x76,
	0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x28, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x06, 0x2e,
	0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69,
	0x64, 0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x07,
	0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x73, 0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x12, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a,
	0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x26,
	0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x69, 0x64, 0x73, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x09, 0x2e, 0x76, 0x31,
	0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64,
	0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x76, 0x31, 0x2e,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x77, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x00,
	0x12, 0x24, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x69, 0x64, 0x73, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x1a, 0x06, 0x2e,
	0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x64, 0x73,
	0x1a, 0x06, 0x2e, 0x76, 0x31, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x42, 0x67, 0x0a, 0x06, 0x63, 0x6f,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65,
	0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_carrier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_carrier_service_proto_goTypes = []interface{}{
	(*Ping)(nil),           // 0: v1.Ping
	(*Id)(nil),             // 1: v1.id
	(*Ids)(nil),            // 2: v1.ids
	(*Business)(nil),       // 3: v1.business
	(*ShipmentStatus)(nil), // 4: v1.shipment_status
	(*Empty)(nil),          // 5: v1.empty
	(*Location)(nil),       // 6: v1.location
	(*BookingRequest)(nil), // 7: v1.booking_request
	(*Bids)(nil),           // 8: v1.bids
	(*Quote)(nil),          // 9: v1.quote
	(*Ok)(nil),             // 10: v1.ok
	(*Businesses)(nil),     // 11: v1.businesses
	(*Locations)(nil),      // 12: v1.locations
	(*Booking)(nil),        // 13: v1.booking
	(*Bookings)(nil),       // 14: v1.bookings
	(*Quotes)(nil),         // 15: v1.quotes
}
var file_carrier_service_proto_depIdxs = []int32{
	3,  // 0: v1.CarrierService.CreateBusiness:input_type -> v1.business
	1,  // 1: v1.CarrierService.GetBusinessById:input_type -> v1.id
	2,  // 2: v1.CarrierService.GetBusinesses:input_type -> v1.ids
	1,  // 3: v1.CarrierService.CloseBusiness:input_type -> v1.id
	1,  // 4: v1.CarrierService.TrackAShipment:input_type -> v1.id
	4,  // 5: v1.CarrierService.UpdateShipmentTracking:input_type -> v1.shipment_status
	5,  // 6: v1.CarrierService.GetLocations:input_type -> v1.empty
	1,  // 7: v1.CarrierService.GetLocation:input_type -> v1.id
	6,  // 8: v1.CarrierService.CreateLocation:input_type -> v1.location
	6,  // 9: v1.CarrierService.AddLocation:input_type -> v1.location
	6,  // 10: v1.CarrierService.UpdateLocation:input_type -> v1.location
	1,  // 11: v1.CarrierService.DeleteLocation:input_type -> v1.id
	2,  // 12: v1.CarrierService.DeleteLocations:input_type -> v1.ids
	5,  // 13: v1.CarrierService.PingCarrierService:input_type -> v1.empty
	7,  // 14: v1.CarrierService.BookAQuote:input_type -> v1.booking_request
	8,  // 15: v1.CarrierService.BookQuotes:input_type -> v1.bids
	5,  // 16: v1.CarrierService.GetBookingHistory:input_type -> v1.empty
	1,  // 17: v1.CarrierService.GetBookingById:input_type -> v1.id
	5,  // 18: v1.CarrierService.GetAllQuotes:input_type -> v1.empty
	9,  // 19: v1.CarrierService.GetNewQuotes:input_type -> v1.quote
	9,  // 20: v1.CarrierService.UpdateQuote:input_type -> v1.quote
	1,  // 21: v1.CarrierService.DeleteQuote:input_type -> v1.id
	2,  // 22: v1.CarrierService.DeleteQuotes:input_type -> v1.ids
	10, // 23: v1.CarrierService.CreateBusiness:output_type -> v1.ok
	3,  // 24: v1.CarrierService.GetBusinessById:output_type -> v1.business
	11, // 25: v1.CarrierService.GetBusinesses:output_type -> v1.businesses
	10, // 26: v1.CarrierService.CloseBusiness:output_type -> v1.ok
	4,  // 27: v1.CarrierService.TrackAShipment:output_type -> v1.shipment_status
	10, // 28: v1.CarrierService.UpdateShipmentTracking:output_type -> v1.ok
	12, // 29: v1.CarrierService.GetLocations:output_type -> v1.locations
	6,  // 30: v1.CarrierService.GetLocation:output_type -> v1.location
	10, // 31: v1.CarrierService.CreateLocation:output_type -> v1.ok
	10, // 32: v1.CarrierService.AddLocation:output_type -> v1.ok
	10, // 33: v1.CarrierService.UpdateLocation:output_type -> v1.ok
	10, // 34: v1.CarrierService.DeleteLocation:output_type -> v1.ok
	10, // 35: v1.CarrierService.DeleteLocations:output_type -> v1.ok
	10, // 36: v1.CarrierService.PingCarrierService:output_type -> v1.ok
	13, // 37: v1.CarrierService.BookAQuote:output_type -> v1.booking
	14, // 38: v1.CarrierService.BookQuotes:output_type -> v1.bookings
	14, // 39: v1.CarrierService.GetBookingHistory:output_type -> v1.bookings
	13, // 40: v1.CarrierService.GetBookingById:output_type -> v1.booking
	15, // 41: v1.CarrierService.GetAllQuotes:output_type -> v1.quotes
	9,  // 42: v1.CarrierService.GetNewQuotes:output_type -> v1.quote
	8,  // 43: v1.CarrierService.UpdateQuote:output_type -> v1.bids
	10, // 44: v1.CarrierService.DeleteQuote:output_type -> v1.ok
	10, // 45: v1.CarrierService.DeleteQuotes:output_type -> v1.ok
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
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
	file_book_proto_init()
	file_validate_token_proto_init()
	file_business_proto_init()
	file_empty_proto_init()
	file_bid_proto_init()
	file_location_proto_init()
	file_ok_proto_init()
	file_shipment_status_proto_init()
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
		file_carrier_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   3,
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

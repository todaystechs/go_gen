// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: carrier_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CarrierServicePing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"hi,omitempty"
	Hi *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty" dynamodbav:"hi,omitempty"`
}

func (x *CarrierServicePing) Reset() {
	*x = CarrierServicePing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarrierServicePing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarrierServicePing) ProtoMessage() {}

func (x *CarrierServicePing) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CarrierServicePing.ProtoReflect.Descriptor instead.
func (*CarrierServicePing) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{0}
}

func (x *CarrierServicePing) GetHi() *wrapperspb.StringValue {
	if x != nil {
		return x.Hi
	}
	return nil
}

type BusinessId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
}

func (x *BusinessId) Reset() {
	*x = BusinessId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessId) ProtoMessage() {}

func (x *BusinessId) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BusinessId.ProtoReflect.Descriptor instead.
func (*BusinessId) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{1}
}

func (x *BusinessId) GetBusinessId() *wrapperspb.StringValue {
	if x != nil {
		return x.BusinessId
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{2}
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

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{3}
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

	Id []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[4]
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
	return file_carrier_service_proto_rawDescGZIP(), []int{4}
}

func (x *Ids) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_carrier_service_proto_rawDescGZIP(), []int{5}
}

var File_carrier_service_proto protoreflect.FileDescriptor

var file_carrier_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0b, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x42, 0x0a, 0x12, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x02, 0x68, 0x69, 0x22, 0x4b, 0x0a, 0x0a, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x9d, 0x08, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x0d, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x08, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x4f, 0x66, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4f, 0x6b, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b,
	0x22, 0x00, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x08,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x09, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x73, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4f, 0x6b, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x41, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x69, 0x64, 0x73, 0x1a, 0x19,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x42,
	0x6f, 0x6f, 0x6b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x42, 0x69, 0x64, 0x73, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x08, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x69,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x69, 0x64, 0x44, 0x61, 0x74, 0x61, 0x73, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x09, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x73, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x42, 0x71, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x42, 0x13, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68,
	0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58,
	0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2,
	0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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

var file_carrier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_carrier_service_proto_goTypes = []interface{}{
	(*CarrierServicePing)(nil),     // 0: user.CarrierServicePing
	(*BusinessId)(nil),             // 1: user.BusinessId
	(*Token)(nil),                  // 2: user.Token
	(*Id)(nil),                     // 3: user.Id
	(*Ids)(nil),                    // 4: user.Ids
	(*Empty)(nil),                  // 5: user.Empty
	(*wrapperspb.StringValue)(nil), // 6: google.protobuf.StringValue
	(*BusinessData)(nil),           // 7: user.BusinessData
	(*GetBusinessRequest)(nil),     // 8: user.GetBusinessRequest
	(*Location)(nil),               // 9: user.Location
	(*BookBids)(nil),               // 10: user.BookBids
	(*QuoteData)(nil),              // 11: user.QuoteData
	(*Ok)(nil),                     // 12: user.Ok
	(*GetBusinessesResponse)(nil),  // 13: user.GetBusinessesResponse
	(*ListsOfLocation)(nil),        // 14: user.ListsOfLocation
	(*BookingResponseData)(nil),    // 15: user.BookingResponseData
	(*ListOfBookingResponse)(nil),  // 16: user.ListOfBookingResponse
	(*QuoteDatas)(nil),             // 17: user.QuoteDatas
	(*BidDatas)(nil),               // 18: user.BidDatas
}
var file_carrier_service_proto_depIdxs = []int32{
	6,  // 0: user.CarrierServicePing.hi:type_name -> google.protobuf.StringValue
	6,  // 1: user.BusinessId.business_id:type_name -> google.protobuf.StringValue
	7,  // 2: user.CarrierService.CreateBusiness:input_type -> user.BusinessData
	8,  // 3: user.CarrierService.GetBusinessById:input_type -> user.GetBusinessRequest
	5,  // 4: user.CarrierService.GetBusinesses:input_type -> user.Empty
	3,  // 5: user.CarrierService.CloseBusiness:input_type -> user.Id
	5,  // 6: user.CarrierService.GetLocations:input_type -> user.Empty
	5,  // 7: user.CarrierService.GetLocation:input_type -> user.Empty
	9,  // 8: user.CarrierService.CreateLocation:input_type -> user.Location
	9,  // 9: user.CarrierService.UpdateLocation:input_type -> user.Location
	3,  // 10: user.CarrierService.DeleteLocation:input_type -> user.Id
	4,  // 11: user.CarrierService.DeleteLocations:input_type -> user.Ids
	0,  // 12: user.CarrierService.PingCarrierService:input_type -> user.CarrierServicePing
	10, // 13: user.CarrierService.BookAQuote:input_type -> user.BookBids
	10, // 14: user.CarrierService.BookQuotes:input_type -> user.BookBids
	5,  // 15: user.CarrierService.GetBookingHistory:input_type -> user.Empty
	3,  // 16: user.CarrierService.GetBookingById:input_type -> user.Id
	5,  // 17: user.CarrierService.GetAllQuotes:input_type -> user.Empty
	11, // 18: user.CarrierService.GetNewQuotes:input_type -> user.QuoteData
	11, // 19: user.CarrierService.UpdateQuote:input_type -> user.QuoteData
	3,  // 20: user.CarrierService.DeleteQuote:input_type -> user.Id
	4,  // 21: user.CarrierService.DeleteQuotes:input_type -> user.Ids
	12, // 22: user.CarrierService.CreateBusiness:output_type -> user.Ok
	7,  // 23: user.CarrierService.GetBusinessById:output_type -> user.BusinessData
	13, // 24: user.CarrierService.GetBusinesses:output_type -> user.GetBusinessesResponse
	12, // 25: user.CarrierService.CloseBusiness:output_type -> user.Ok
	14, // 26: user.CarrierService.GetLocations:output_type -> user.ListsOfLocation
	9,  // 27: user.CarrierService.GetLocation:output_type -> user.Location
	12, // 28: user.CarrierService.CreateLocation:output_type -> user.Ok
	12, // 29: user.CarrierService.UpdateLocation:output_type -> user.Ok
	12, // 30: user.CarrierService.DeleteLocation:output_type -> user.Ok
	12, // 31: user.CarrierService.DeleteLocations:output_type -> user.Ok
	0,  // 32: user.CarrierService.PingCarrierService:output_type -> user.CarrierServicePing
	15, // 33: user.CarrierService.BookAQuote:output_type -> user.BookingResponseData
	16, // 34: user.CarrierService.BookQuotes:output_type -> user.ListOfBookingResponse
	16, // 35: user.CarrierService.GetBookingHistory:output_type -> user.ListOfBookingResponse
	16, // 36: user.CarrierService.GetBookingById:output_type -> user.ListOfBookingResponse
	17, // 37: user.CarrierService.GetAllQuotes:output_type -> user.QuoteDatas
	18, // 38: user.CarrierService.GetNewQuotes:output_type -> user.BidDatas
	18, // 39: user.CarrierService.UpdateQuote:output_type -> user.BidDatas
	12, // 40: user.CarrierService.DeleteQuote:output_type -> user.Ok
	12, // 41: user.CarrierService.DeleteQuotes:output_type -> user.Ok
	22, // [22:42] is the sub-list for method output_type
	2,  // [2:22] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
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
	if !protoimpl.UnsafeEnabled {
		file_carrier_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarrierServicePing); i {
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
			switch v := v.(*BusinessId); i {
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
		file_carrier_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_carrier_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_carrier_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			NumMessages:   6,
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

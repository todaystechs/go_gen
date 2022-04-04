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
	0x64, 0x32, 0xc8, 0x06, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x6c, 0x6c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x57, 0x69,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x4f, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f,
	0x6b, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x12, 0x50, 0x69, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x38, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12,
	0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4e, 0x65, 0x77, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x69, 0x64, 0x73, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x1a, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x69,
	0x64, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x66, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x71, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x13, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02,
	0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_carrier_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_carrier_service_proto_goTypes = []interface{}{
	(*CarrierServicePing)(nil),     // 0: user.CarrierServicePing
	(*BusinessId)(nil),             // 1: user.BusinessId
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*BusinessRequest)(nil),        // 3: user.BusinessRequest
	(*AllBusinessRequest)(nil),     // 4: user.AllBusinessRequest
	(*Location)(nil),               // 5: user.Location
	(*FetchQuotes)(nil),            // 6: user.FetchQuotes
	(*QuoteEntity)(nil),            // 7: user.QuoteEntity
	(*DeleteQuoteData)(nil),        // 8: user.DeleteQuoteData
	(*BookingData)(nil),            // 9: user.BookingData
	(*FetchBookingsRequest)(nil),   // 10: user.FetchBookingsRequest
	(*DynamoBusiness)(nil),         // 11: user.DynamoBusiness
	(*OkWithData)(nil),             // 12: user.OkWithData
	(*ListsOfLocation)(nil),        // 13: user.ListsOfLocation
	(*Ok)(nil),                     // 14: user.Ok
	(*QuoteEntities)(nil),          // 15: user.QuoteEntities
	(*QuoteBids)(nil),              // 16: user.QuoteBids
	(*BookingResponse)(nil),        // 17: user.BookingResponse
	(*ListOfBooking)(nil),          // 18: user.ListOfBooking
}
var file_carrier_service_proto_depIdxs = []int32{
	2,  // 0: user.CarrierServicePing.hi:type_name -> google.protobuf.StringValue
	2,  // 1: user.BusinessId.business_id:type_name -> google.protobuf.StringValue
	3,  // 2: user.CarrierService.GetBusinessById:input_type -> user.BusinessRequest
	4,  // 3: user.CarrierService.GetAllBusinesses:input_type -> user.AllBusinessRequest
	1,  // 4: user.CarrierService.GetLocations:input_type -> user.BusinessId
	5,  // 5: user.CarrierService.CreateLocation:input_type -> user.Location
	5,  // 6: user.CarrierService.UpdateLocation:input_type -> user.Location
	5,  // 7: user.CarrierService.DeleteLocation:input_type -> user.Location
	0,  // 8: user.CarrierService.PingCarrierService:input_type -> user.CarrierServicePing
	6,  // 9: user.CarrierService.GetAllQuotes:input_type -> user.FetchQuotes
	7,  // 10: user.CarrierService.GetNewQuotes:input_type -> user.QuoteEntity
	7,  // 11: user.CarrierService.UpdateQuote:input_type -> user.QuoteEntity
	8,  // 12: user.CarrierService.DeleteQuote:input_type -> user.DeleteQuoteData
	9,  // 13: user.CarrierService.BookQuote:input_type -> user.BookingData
	10, // 14: user.CarrierService.GetBookingHistory:input_type -> user.FetchBookingsRequest
	10, // 15: user.CarrierService.GetBookingById:input_type -> user.FetchBookingsRequest
	11, // 16: user.CarrierService.GetBusinessById:output_type -> user.DynamoBusiness
	12, // 17: user.CarrierService.GetAllBusinesses:output_type -> user.OkWithData
	13, // 18: user.CarrierService.GetLocations:output_type -> user.ListsOfLocation
	5,  // 19: user.CarrierService.CreateLocation:output_type -> user.Location
	14, // 20: user.CarrierService.UpdateLocation:output_type -> user.Ok
	14, // 21: user.CarrierService.DeleteLocation:output_type -> user.Ok
	0,  // 22: user.CarrierService.PingCarrierService:output_type -> user.CarrierServicePing
	15, // 23: user.CarrierService.GetAllQuotes:output_type -> user.QuoteEntities
	16, // 24: user.CarrierService.GetNewQuotes:output_type -> user.QuoteBids
	16, // 25: user.CarrierService.UpdateQuote:output_type -> user.QuoteBids
	14, // 26: user.CarrierService.DeleteQuote:output_type -> user.Ok
	17, // 27: user.CarrierService.BookQuote:output_type -> user.BookingResponse
	18, // 28: user.CarrierService.GetBookingHistory:output_type -> user.ListOfBooking
	18, // 29: user.CarrierService.GetBookingById:output_type -> user.ListOfBooking
	16, // [16:30] is the sub-list for method output_type
	2,  // [2:16] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_carrier_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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

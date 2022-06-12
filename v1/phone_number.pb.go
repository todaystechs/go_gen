// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: phone_number.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PhoneNumberType int32

const (
	PhoneNumberType_Home      PhoneNumberType = 0
	PhoneNumberType_Office    PhoneNumberType = 1
	PhoneNumberType_CellPhone PhoneNumberType = 2
)

// Enum value maps for PhoneNumberType.
var (
	PhoneNumberType_name = map[int32]string{
		0: "Home",
		1: "Office",
		2: "CellPhone",
	}
	PhoneNumberType_value = map[string]int32{
		"Home":      0,
		"Office":    1,
		"CellPhone": 2,
	}
)

func (x PhoneNumberType) Enum() *PhoneNumberType {
	p := new(PhoneNumberType)
	*p = x
	return p
}

func (x PhoneNumberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneNumberType) Descriptor() protoreflect.EnumDescriptor {
	return file_phone_number_proto_enumTypes[0].Descriptor()
}

func (PhoneNumberType) Type() protoreflect.EnumType {
	return &file_phone_number_proto_enumTypes[0]
}

func (x PhoneNumberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneNumberType.Descriptor instead.
func (PhoneNumberType) EnumDescriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{0}
}

type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"phone_number,omitempty"
	PhoneNumber string `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	// @gotags: dynamodbav:"phone_number_type,omitempty"
	Type PhoneNumberType `protobuf:"varint,2,opt,name=type,proto3,enum=v1.PhoneNumberType" json:"type,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phone_number_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_phone_number_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_phone_number_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneNumber) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PhoneNumber) GetType() PhoneNumberType {
	if x != nil {
		return x.Type
	}
	return PhoneNumberType_Home
}

var File_phone_number_proto protoreflect.FileDescriptor

var file_phone_number_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5c, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x2a,
	0x38, 0x0a, 0x11, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x65,
	0x6c, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x10, 0x02, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_phone_number_proto_rawDescOnce sync.Once
	file_phone_number_proto_rawDescData = file_phone_number_proto_rawDesc
)

func file_phone_number_proto_rawDescGZIP() []byte {
	file_phone_number_proto_rawDescOnce.Do(func() {
		file_phone_number_proto_rawDescData = protoimpl.X.CompressGZIP(file_phone_number_proto_rawDescData)
	})
	return file_phone_number_proto_rawDescData
}

var file_phone_number_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_phone_number_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_phone_number_proto_goTypes = []interface{}{
	(PhoneNumberType)(0), // 0: v1.phone_number_type
	(*PhoneNumber)(nil),  // 1: v1.phone_number
}
var file_phone_number_proto_depIdxs = []int32{
	0, // 0: v1.phone_number.type:type_name -> v1.phone_number_type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_phone_number_proto_init() }
func file_phone_number_proto_init() {
	if File_phone_number_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phone_number_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
			RawDescriptor: file_phone_number_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_phone_number_proto_goTypes,
		DependencyIndexes: file_phone_number_proto_depIdxs,
		EnumInfos:         file_phone_number_proto_enumTypes,
		MessageInfos:      file_phone_number_proto_msgTypes,
	}.Build()
	File_phone_number_proto = out.File
	file_phone_number_proto_rawDesc = nil
	file_phone_number_proto_goTypes = nil
	file_phone_number_proto_depIdxs = nil
}

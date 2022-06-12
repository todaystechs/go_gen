// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: between.proto

package user

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

type Between struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"starts_from,omitempty"
	StartsFrom string `protobuf:"bytes,1,opt,name=starts_from,json=startsFrom,proto3" json:"starts_from,omitempty" dynamodbav:"starts_from,omitempty"`
	// @gotags: dynamodbav:"ends_on,omitempty"
	EndOn string `protobuf:"bytes,2,opt,name=end_on,json=endOn,proto3" json:"end_on,omitempty" dynamodbav:"ends_on,omitempty"`
}

func (x *Between) Reset() {
	*x = Between{}
	if protoimpl.UnsafeEnabled {
		mi := &file_between_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Between) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Between) ProtoMessage() {}

func (x *Between) ProtoReflect() protoreflect.Message {
	mi := &file_between_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Between.ProtoReflect.Descriptor instead.
func (*Between) Descriptor() ([]byte, []int) {
	return file_between_proto_rawDescGZIP(), []int{0}
}

func (x *Between) GetStartsFrom() string {
	if x != nil {
		return x.StartsFrom
	}
	return ""
}

func (x *Between) GetEndOn() string {
	if x != nil {
		return x.EndOn
	}
	return ""
}

var File_between_proto protoreflect.FileDescriptor

var file_between_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x07, 0x62, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x46, 0x72, 0x6f,
	0x6d, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x75,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_between_proto_rawDescOnce sync.Once
	file_between_proto_rawDescData = file_between_proto_rawDesc
)

func file_between_proto_rawDescGZIP() []byte {
	file_between_proto_rawDescOnce.Do(func() {
		file_between_proto_rawDescData = protoimpl.X.CompressGZIP(file_between_proto_rawDescData)
	})
	return file_between_proto_rawDescData
}

var file_between_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_between_proto_goTypes = []interface{}{
	(*Between)(nil), // 0: user.between
}
var file_between_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_between_proto_init() }
func file_between_proto_init() {
	if File_between_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_between_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Between); i {
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
			RawDescriptor: file_between_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_between_proto_goTypes,
		DependencyIndexes: file_between_proto_depIdxs,
		MessageInfos:      file_between_proto_msgTypes,
	}.Build()
	File_between_proto = out.File
	file_between_proto_rawDesc = nil
	file_between_proto_goTypes = nil
	file_between_proto_depIdxs = nil
}

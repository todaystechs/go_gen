// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: locked.proto

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"is_locked,omitempty"
	IsLocked bool `protobuf:"varint,1,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty" dynamodbav:"is_locked,omitempty"`
	// @gotags: dynamodbav:"since,omitempty"
	Since string `protobuf:"bytes,2,opt,name=since,proto3" json:"since,omitempty" dynamodbav:"since,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locked_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_locked_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_locked_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetIsLocked() bool {
	if x != nil {
		return x.IsLocked
	}
	return false
}

func (x *Status) GetSince() string {
	if x != nil {
		return x.Since
	}
	return ""
}

var File_locked_proto protoreflect.FileDescriptor

var file_locked_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x22, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_locked_proto_rawDescOnce sync.Once
	file_locked_proto_rawDescData = file_locked_proto_rawDesc
)

func file_locked_proto_rawDescGZIP() []byte {
	file_locked_proto_rawDescOnce.Do(func() {
		file_locked_proto_rawDescData = protoimpl.X.CompressGZIP(file_locked_proto_rawDescData)
	})
	return file_locked_proto_rawDescData
}

var file_locked_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_locked_proto_goTypes = []interface{}{
	(*Status)(nil), // 0: v1.status
}
var file_locked_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_locked_proto_init() }
func file_locked_proto_init() {
	if File_locked_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_locked_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_locked_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_locked_proto_goTypes,
		DependencyIndexes: file_locked_proto_depIdxs,
		MessageInfos:      file_locked_proto_msgTypes,
	}.Build()
	File_locked_proto = out.File
	file_locked_proto_rawDesc = nil
	file_locked_proto_goTypes = nil
	file_locked_proto_depIdxs = nil
}

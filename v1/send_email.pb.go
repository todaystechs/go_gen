// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: send_email.proto

package v1

import (
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

type SendEmailData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"email_subject"
	EmailSubject string `protobuf:"bytes,1,opt,name=email_subject,json=emailSubject,proto3" json:"email_subject,omitempty" dynamodbav:"email_subject"`
	// @gotags: dynamodbav:"receiver_email"
	ReceiverEmailAddress string `protobuf:"bytes,2,opt,name=receiver_email_address,json=receiverEmailAddress,proto3" json:"receiver_email_address,omitempty" dynamodbav:"receiver_email"`
	// @gotags: dynamodbav:"receiver_name"
	ReceiverName string `protobuf:"bytes,3,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty" dynamodbav:"receiver_name"`
	// @gotags: dynamodbav:"email_purpose"
	EmailPurpose string `protobuf:"bytes,4,opt,name=email_purpose,json=emailPurpose,proto3" json:"email_purpose,omitempty" dynamodbav:"email_purpose"`
	// @gotags: dynamodbav:"token"
	Token string `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token"`
	// @gotags: dynamodbav:"success"
	Success bool `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty" dynamodbav:"success"`
}

func (x *SendEmailData) Reset() {
	*x = SendEmailData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_send_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailData) ProtoMessage() {}

func (x *SendEmailData) ProtoReflect() protoreflect.Message {
	mi := &file_send_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailData.ProtoReflect.Descriptor instead.
func (*SendEmailData) Descriptor() ([]byte, []int) {
	return file_send_email_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailData) GetEmailSubject() string {
	if x != nil {
		return x.EmailSubject
	}
	return ""
}

func (x *SendEmailData) GetReceiverEmailAddress() string {
	if x != nil {
		return x.ReceiverEmailAddress
	}
	return ""
}

func (x *SendEmailData) GetReceiverName() string {
	if x != nil {
		return x.ReceiverName
	}
	return ""
}

func (x *SendEmailData) GetEmailPurpose() string {
	if x != nil {
		return x.EmailPurpose
	}
	return ""
}

func (x *SendEmailData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendEmailData) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_send_email_proto protoreflect.FileDescriptor

var file_send_email_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x6c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_send_email_proto_rawDescOnce sync.Once
	file_send_email_proto_rawDescData = file_send_email_proto_rawDesc
)

func file_send_email_proto_rawDescGZIP() []byte {
	file_send_email_proto_rawDescOnce.Do(func() {
		file_send_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_send_email_proto_rawDescData)
	})
	return file_send_email_proto_rawDescData
}

var file_send_email_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_send_email_proto_goTypes = []interface{}{
	(*SendEmailData)(nil), // 0: user.SendEmailData
}
var file_send_email_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_send_email_proto_init() }
func file_send_email_proto_init() {
	if File_send_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_send_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailData); i {
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
			RawDescriptor: file_send_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_send_email_proto_goTypes,
		DependencyIndexes: file_send_email_proto_depIdxs,
		MessageInfos:      file_send_email_proto_msgTypes,
	}.Build()
	File_send_email_proto = out.File
	file_send_email_proto_rawDesc = nil
	file_send_email_proto_goTypes = nil
	file_send_email_proto_depIdxs = nil
}
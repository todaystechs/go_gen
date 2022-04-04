// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: sign_up.proto

package v1

import (
	_ "github.com/todaystechs/go_gen/v1/validate"
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

type SignUpData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"first_name,omitempty"
	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty" dynamodbav:"first_name,omitempty"`
	// @gotags: dynamodbav:"middle_name,omitempty"
	MiddleName string `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty" dynamodbav:"middle_name,omitempty"`
	// @gotags: dynamodbav:"last_name,omitempty"
	LastName string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty" dynamodbav:"last_name,omitempty"`
	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email,omitempty"`
	// @gotags: dynamodbav:"password,omitempty"
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty" dynamodbav:"password,omitempty"`
	// @gotags: dynamodbav:"phone_number,omitempty"
	PhoneNumber string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty" dynamodbav:"phone_number,omitempty"`
	// @gotags: dynamodbav:"role,omitempty"
	Role []Role `protobuf:"varint,7,rep,packed,name=role,proto3,enum=user.Role" json:"role,omitempty" dynamodbav:"role,omitempty"`
	// @gotags: dynamodbav:"company_name,omitempty"
	CompanyName string `protobuf:"bytes,8,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty" dynamodbav:"company_name,omitempty"`
}

func (x *SignUpData) Reset() {
	*x = SignUpData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sign_up_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignUpData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpData) ProtoMessage() {}

func (x *SignUpData) ProtoReflect() protoreflect.Message {
	mi := &file_sign_up_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpData.ProtoReflect.Descriptor instead.
func (*SignUpData) Descriptor() ([]byte, []int) {
	return file_sign_up_proto_rawDescGZIP(), []int{0}
}

func (x *SignUpData) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignUpData) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *SignUpData) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpData) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *SignUpData) GetRole() []Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *SignUpData) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

var File_sign_up_proto protoreflect.FileDescriptor

var file_sign_up_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x0a, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04,
	0x72, 0x02, 0x10, 0x08, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x10, 0x0a, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x69, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0b, 0x53, 0x69, 0x67,
	0x6e, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63,
	0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55,
	0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72,
	0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sign_up_proto_rawDescOnce sync.Once
	file_sign_up_proto_rawDescData = file_sign_up_proto_rawDesc
)

func file_sign_up_proto_rawDescGZIP() []byte {
	file_sign_up_proto_rawDescOnce.Do(func() {
		file_sign_up_proto_rawDescData = protoimpl.X.CompressGZIP(file_sign_up_proto_rawDescData)
	})
	return file_sign_up_proto_rawDescData
}

var file_sign_up_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sign_up_proto_goTypes = []interface{}{
	(*SignUpData)(nil), // 0: user.SignUpData
	(Role)(0),          // 1: user.Role
}
var file_sign_up_proto_depIdxs = []int32{
	1, // 0: user.SignUpData.role:type_name -> user.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sign_up_proto_init() }
func file_sign_up_proto_init() {
	if File_sign_up_proto != nil {
		return
	}
	file_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sign_up_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignUpData); i {
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
			RawDescriptor: file_sign_up_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sign_up_proto_goTypes,
		DependencyIndexes: file_sign_up_proto_depIdxs,
		MessageInfos:      file_sign_up_proto_msgTypes,
	}.Build()
	File_sign_up_proto = out.File
	file_sign_up_proto_rawDesc = nil
	file_sign_up_proto_goTypes = nil
	file_sign_up_proto_depIdxs = nil
}

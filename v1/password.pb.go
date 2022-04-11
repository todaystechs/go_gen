// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: password.proto

package v1

import (
	_ "github.com/todaystechs/go_gen/v1/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type ForgotPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"email"
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email"`
}

func (x *ForgotPassword) Reset() {
	*x = ForgotPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForgotPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPassword) ProtoMessage() {}

func (x *ForgotPassword) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPassword.ProtoReflect.Descriptor instead.
func (*ForgotPassword) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{0}
}

func (x *ForgotPassword) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ResetPasswordToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"issued_to,omitempty"
	IssuedTo string `protobuf:"bytes,1,opt,name=issued_to,json=issuedTo,proto3" json:"issued_to,omitempty" dynamodbav:"issued_to,omitempty"`
	// @gotags: dynamodbav:"issued_on,omitempty"
	IssuedOn string `protobuf:"bytes,2,opt,name=issued_on,json=issuedOn,proto3" json:"issued_on,omitempty" dynamodbav:"issued_on,omitempty"`
	// @gotags: dynamodbav:"expires_on,omitempty"
	ExpiresOn string `protobuf:"bytes,3,opt,name=expires_on,json=expiresOn,proto3" json:"expires_on,omitempty" dynamodbav:"expires_on,omitempty"`
	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,5,opt,name=businessId,proto3" json:"businessId,omitempty" dynamodbav:"business_id,omitempty"`
}

func (x *ResetPasswordToken) Reset() {
	*x = ResetPasswordToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordToken) ProtoMessage() {}

func (x *ResetPasswordToken) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordToken.ProtoReflect.Descriptor instead.
func (*ResetPasswordToken) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{1}
}

func (x *ResetPasswordToken) GetIssuedTo() string {
	if x != nil {
		return x.IssuedTo
	}
	return ""
}

func (x *ResetPasswordToken) GetIssuedOn() string {
	if x != nil {
		return x.IssuedOn
	}
	return ""
}

func (x *ResetPasswordToken) GetExpiresOn() string {
	if x != nil {
		return x.ExpiresOn
	}
	return ""
}

func (x *ResetPasswordToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResetPasswordToken) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

type ResetPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"new_password,omitempty"
	NewPassword string `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty" dynamodbav:"new_password,omitempty"`
	// @gotags: dynamodbav:"confirm_password,omitempty"
	ConfirmPassword string `protobuf:"bytes,3,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty" dynamodbav:"confirm_password,omitempty"`
	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email,omitempty"`
}

func (x *ResetPassword) Reset() {
	*x = ResetPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_password_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPassword) ProtoMessage() {}

func (x *ResetPassword) ProtoReflect() protoreflect.Message {
	mi := &file_password_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPassword.ProtoReflect.Descriptor instead.
func (*ResetPassword) Descriptor() ([]byte, []int) {
	return file_password_proto_rawDescGZIP(), []int{2}
}

func (x *ResetPassword) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResetPassword) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ResetPassword) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *ResetPassword) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_password_proto protoreflect.FileDescriptor

var file_password_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x32, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0xa5, 0x01, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x54, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x4f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0xb7, 0x01, 0x0a,
	0x0e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x2c, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02,
	0x10, 0x08, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x34, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04,
	0x72, 0x02, 0x10, 0x08, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x10, 0x05, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x6b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x42, 0x0d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73,
	0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_password_proto_rawDescOnce sync.Once
	file_password_proto_rawDescData = file_password_proto_rawDesc
)

func file_password_proto_rawDescGZIP() []byte {
	file_password_proto_rawDescOnce.Do(func() {
		file_password_proto_rawDescData = protoimpl.X.CompressGZIP(file_password_proto_rawDescData)
	})
	return file_password_proto_rawDescData
}

var file_password_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_password_proto_goTypes = []interface{}{
	(*ForgotPassword)(nil),     // 0: user.forgot_password
	(*ResetPasswordToken)(nil), // 1: user.reset_password_token
	(*ResetPassword)(nil),      // 2: user.reset_password
}
var file_password_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_password_proto_init() }
func file_password_proto_init() {
	if File_password_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_password_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForgotPassword); i {
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
		file_password_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordToken); i {
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
		file_password_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPassword); i {
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
			RawDescriptor: file_password_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_password_proto_goTypes,
		DependencyIndexes: file_password_proto_depIdxs,
		MessageInfos:      file_password_proto_msgTypes,
	}.Build()
	File_password_proto = out.File
	file_password_proto_rawDesc = nil
	file_password_proto_goTypes = nil
	file_password_proto_depIdxs = nil
}

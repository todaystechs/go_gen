// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: user_service.proto

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

type UserServicePingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"hi,omitempty"
	Hi string `protobuf:"bytes,1,opt,name=hi,proto3" json:"hi,omitempty" dynamodbav:"hi,omitempty"`
}

func (x *UserServicePingData) Reset() {
	*x = UserServicePingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserServicePingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserServicePingData) ProtoMessage() {}

func (x *UserServicePingData) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserServicePingData.ProtoReflect.Descriptor instead.
func (*UserServicePingData) Descriptor() ([]byte, []int) {
	return file_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserServicePingData) GetHi() string {
	if x != nil {
		return x.Hi
	}
	return ""
}

var File_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x69, 0x67, 0x6e,
	0x5f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6c, 0x6f, 0x67, 0x5f, 0x6f, 0x75, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x6f, 0x6b, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x25, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x68, 0x69, 0x32, 0x8c, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x00, 0x12, 0x26, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f,
	0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x4c,
	0x6f, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f,
	0x6b, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x72,
	0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x1a,
	0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x4f, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x10, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00,
	0x12, 0x27, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x6b, 0x57, 0x69, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x42, 0x6e, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04,
	0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_rawDescData = file_user_service_proto_rawDesc
)

func file_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_rawDescData)
	})
	return file_user_service_proto_rawDescData
}

var file_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_service_proto_goTypes = []interface{}{
	(*UserServicePingData)(nil), // 0: user.UserServicePingData
	(*SignUpData)(nil),          // 1: user.SignUpData
	(*LoginUserData)(nil),       // 2: user.LoginUserData
	(*LogOutData)(nil),          // 3: user.LogOutData
	(*ForgotPasswordData)(nil),  // 4: user.ForgotPasswordData
	(*ResetPasswordData)(nil),   // 5: user.ResetPasswordData
	(*UpdateUserRoleData)(nil),  // 6: user.UpdateUserRoleData
	(*AddStaffData)(nil),        // 7: user.AddStaffData
	(*UserData)(nil),            // 8: user.UserData
	(*MeData)(nil),              // 9: user.MeData
	(*ConfirmEmailData)(nil),    // 10: user.ConfirmEmailData
	(*UserHomeData)(nil),        // 11: user.UserHomeData
	(*Ok)(nil),                  // 12: user.Ok
	(*OkWithData)(nil),          // 13: user.OkWithData
}
var file_user_service_proto_depIdxs = []int32{
	0,  // 0: user.UserService.PingUserService:input_type -> user.UserServicePingData
	1,  // 1: user.UserService.SignUp:input_type -> user.SignUpData
	2,  // 2: user.UserService.LogIn:input_type -> user.LoginUserData
	3,  // 3: user.UserService.LogOut:input_type -> user.LogOutData
	4,  // 4: user.UserService.ForgotPassword:input_type -> user.ForgotPasswordData
	5,  // 5: user.UserService.ResetPassword:input_type -> user.ResetPasswordData
	6,  // 6: user.UserService.UpdateStaffRole:input_type -> user.UpdateUserRoleData
	7,  // 7: user.UserService.AddStaff:input_type -> user.AddStaffData
	8,  // 8: user.UserService.UpdateUser:input_type -> user.UserData
	9,  // 9: user.UserService.GetMe:input_type -> user.MeData
	10, // 10: user.UserService.ConfirmEmail:input_type -> user.ConfirmEmailData
	11, // 11: user.UserService.Home:input_type -> user.UserHomeData
	0,  // 12: user.UserService.PingUserService:output_type -> user.UserServicePingData
	12, // 13: user.UserService.SignUp:output_type -> user.Ok
	13, // 14: user.UserService.LogIn:output_type -> user.OkWithData
	12, // 15: user.UserService.LogOut:output_type -> user.Ok
	12, // 16: user.UserService.ForgotPassword:output_type -> user.Ok
	13, // 17: user.UserService.ResetPassword:output_type -> user.OkWithData
	13, // 18: user.UserService.UpdateStaffRole:output_type -> user.OkWithData
	13, // 19: user.UserService.AddStaff:output_type -> user.OkWithData
	13, // 20: user.UserService.UpdateUser:output_type -> user.OkWithData
	8,  // 21: user.UserService.GetMe:output_type -> user.UserData
	13, // 22: user.UserService.ConfirmEmail:output_type -> user.OkWithData
	13, // 23: user.UserService.Home:output_type -> user.OkWithData
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_init() }
func file_user_service_proto_init() {
	if File_user_service_proto != nil {
		return
	}
	file_sign_up_proto_init()
	file_login_proto_init()
	file_password_proto_init()
	file_log_out_proto_init()
	file_role_proto_init()
	file_add_staff_proto_init()
	file_user_proto_init()
	file_confirm_email_proto_init()
	file_home_proto_init()
	file_ok_with_data_proto_init()
	file_ok_proto_init()
	file_send_email_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserServicePingData); i {
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
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}

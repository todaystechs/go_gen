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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

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
	0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x84, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70,
	0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x1a,
	0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x20, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x49, 0x6e, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x23, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x67, 0x5f, 0x6f, 0x75, 0x74, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x66, 0x6f, 0x72, 0x67,
	0x6f, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x08, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x08,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22,
	0x00, 0x12, 0x27, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x08,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6f, 0x6b, 0x22, 0x00,
	0x12, 0x22, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6f, 0x6b, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x0f, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x1a, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x6e, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74,
	0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73,
	0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_user_service_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: user.empty
	(*SignUp)(nil),         // 1: user.sign_up
	(*Login)(nil),          // 2: user.login
	(*LogOut)(nil),         // 3: user.log_out
	(*ForgotPassword)(nil), // 4: user.forgot_password
	(*ResetPassword)(nil),  // 5: user.reset_password
	(*UpdateUserRole)(nil), // 6: user.update_user_role
	(*AddStaff)(nil),       // 7: user.add_staff
	(*User)(nil),           // 8: user.user
	(*Token)(nil),          // 9: user.token
	(*ConfirmEmail)(nil),   // 10: user.confirm_email
	(*UserHome)(nil),       // 11: user.user_home
	(*Ok)(nil),             // 12: user.ok
}
var file_user_service_proto_depIdxs = []int32{
	0,  // 0: user.UserService.Ping:input_type -> user.empty
	1,  // 1: user.UserService.SignUp:input_type -> user.sign_up
	2,  // 2: user.UserService.LogIn:input_type -> user.login
	3,  // 3: user.UserService.LogOut:input_type -> user.log_out
	4,  // 4: user.UserService.ForgotPassword:input_type -> user.forgot_password
	5,  // 5: user.UserService.ResetPassword:input_type -> user.reset_password
	6,  // 6: user.UserService.UpdateStaffRole:input_type -> user.update_user_role
	7,  // 7: user.UserService.AddStaff:input_type -> user.add_staff
	8,  // 8: user.UserService.UpdateUser:input_type -> user.user
	9,  // 9: user.UserService.GetMe:input_type -> user.token
	10, // 10: user.UserService.ConfirmEmail:input_type -> user.confirm_email
	11, // 11: user.UserService.Home:input_type -> user.user_home
	12, // 12: user.UserService.Ping:output_type -> user.ok
	12, // 13: user.UserService.SignUp:output_type -> user.ok
	12, // 14: user.UserService.LogIn:output_type -> user.ok
	12, // 15: user.UserService.LogOut:output_type -> user.ok
	12, // 16: user.UserService.ForgotPassword:output_type -> user.ok
	12, // 17: user.UserService.ResetPassword:output_type -> user.ok
	12, // 18: user.UserService.UpdateStaffRole:output_type -> user.ok
	12, // 19: user.UserService.AddStaff:output_type -> user.ok
	12, // 20: user.UserService.UpdateUser:output_type -> user.ok
	8,  // 21: user.UserService.GetMe:output_type -> user.user
	12, // 22: user.UserService.ConfirmEmail:output_type -> user.ok
	8,  // 23: user.UserService.Home:output_type -> user.user
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
	file_ok_proto_init()
	file_send_email_proto_init()
	file_empty_proto_init()
	file_token_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_depIdxs,
	}.Build()
	File_user_service_proto = out.File
	file_user_service_proto_rawDesc = nil
	file_user_service_proto_goTypes = nil
	file_user_service_proto_depIdxs = nil
}

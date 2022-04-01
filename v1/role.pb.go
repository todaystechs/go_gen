// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: role.proto

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

type Role int32

const (
	Role_SystemOwner              Role = 0
	Role_SystemManager            Role = 1
	Role_SystemStaff              Role = 2
	Role_SystemDeveloper          Role = 3
	Role_SystemWarehouseManager   Role = 4
	Role_SystemWarehouseStaff     Role = 5
	Role_BusinessOwner            Role = 6
	Role_BusinessManager          Role = 7
	Role_BusinessStaff            Role = 8
	Role_BusinessWarehouseManager Role = 9
	Role_BusinessWarehouseStaff   Role = 10
	Role_ThreePLOwner             Role = 11
	Role_ThreePLStaff             Role = 12
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0:  "SystemOwner",
		1:  "SystemManager",
		2:  "SystemStaff",
		3:  "SystemDeveloper",
		4:  "SystemWarehouseManager",
		5:  "SystemWarehouseStaff",
		6:  "BusinessOwner",
		7:  "BusinessManager",
		8:  "BusinessStaff",
		9:  "BusinessWarehouseManager",
		10: "BusinessWarehouseStaff",
		11: "ThreePLOwner",
		12: "ThreePLStaff",
	}
	Role_value = map[string]int32{
		"SystemOwner":              0,
		"SystemManager":            1,
		"SystemStaff":              2,
		"SystemDeveloper":          3,
		"SystemWarehouseManager":   4,
		"SystemWarehouseStaff":     5,
		"BusinessOwner":            6,
		"BusinessManager":          7,
		"BusinessStaff":            8,
		"BusinessWarehouseManager": 9,
		"BusinessWarehouseStaff":   10,
		"ThreePLOwner":             11,
		"ThreePLStaff":             12,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_role_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_role_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

type UpdateUserRoleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"new_role,omitempty"
	NewRole []Role `protobuf:"varint,2,rep,packed,name=new_role,json=newRole,proto3,enum=user.Role" json:"new_role,omitempty" dynamodbav:"new_role,omitempty"`
}

func (x *UpdateUserRoleData) Reset() {
	*x = UpdateUserRoleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRoleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRoleData) ProtoMessage() {}

func (x *UpdateUserRoleData) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRoleData.ProtoReflect.Descriptor instead.
func (*UpdateUserRoleData) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRoleData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateUserRoleData) GetNewRole() []Role {
	if x != nil {
		return x.NewRole
	}
	return nil
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x6e,
	0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x2a, 0x9f, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x57,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61, 0x66, 0x66, 0x10, 0x05, 0x12,
	0x11, 0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x66, 0x66, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x66, 0x66, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x50, 0x4c, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x68, 0x72, 0x65, 0x65, 0x50,
	0x4c, 0x53, 0x74, 0x61, 0x66, 0x66, 0x10, 0x0c, 0x42, 0x67, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f,
	0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e,
	0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72,
	0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_role_proto_goTypes = []interface{}{
	(Role)(0),                  // 0: user.Role
	(*UpdateUserRoleData)(nil), // 1: user.UpdateUserRoleData
}
var file_role_proto_depIdxs = []int32{
	0, // 0: user.UpdateUserRoleData.new_role:type_name -> user.Role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRoleData); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		EnumInfos:         file_role_proto_enumTypes,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}

// type Contact struct {
// 	FirstName              string `dynamodbav:"firstName" json:"firstName"`
// 	LastName               string `dynamodbav:"lastName" json:"lastName"`
// 	PhoneNumber            string `dynamodbav:"phoneNumber" json:"phoneNumber"`
// 	PhoneNumberExtension   string `dynamodbav:"phoneNumberExtension" json:"phoneNumberExtension"`
// 	EmailAddress           string `dynamodbav:"emailAddress" json:"emailAddress"`
// 	PhoneNumberDisplay     string `dynamodbav:"phoneNumberDisplay" json:"phoneNumberDisplay"`
// 	Role                   []Role `dynamodbav:"role,string,omitemtpy" json:"role,omitempty"`
// 	PrefferedContactMethod string `dynamodbav:"preferred_contact_method,string,omitemtpy" json:"preferred_contact_method,omitempty"`
// }

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: contact.proto

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

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName              string        `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName               string        `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	PhoneNumber            string        `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	PhoneNumberExtension   string        `protobuf:"bytes,4,opt,name=phoneNumberExtension,proto3" json:"phoneNumberExtension,omitempty"`
	EmailAddress           string        `protobuf:"bytes,5,opt,name=emailAddress,proto3" json:"emailAddress,omitempty"`
	PhoneNumberDisplay     string        `protobuf:"bytes,6,opt,name=phoneNumberDisplay,proto3" json:"phoneNumberDisplay,omitempty"`
	Roles                  []CarrierRole `protobuf:"varint,7,rep,packed,name=roles,proto3,enum=carrier.CarrierRole" json:"roles,omitempty"`
	PrefferedContactMethod string        `protobuf:"bytes,8,opt,name=prefferedContactMethod,proto3" json:"prefferedContactMethod,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Contact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Contact) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *Contact) GetPhoneNumberExtension() string {
	if x != nil {
		return x.PhoneNumberExtension
	}
	return ""
}

func (x *Contact) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Contact) GetPhoneNumberDisplay() string {
	if x != nil {
		return x.PhoneNumberDisplay
	}
	return ""
}

func (x *Contact) GetRoles() []CarrierRole {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *Contact) GetPrefferedContactMethod() string {
	if x != nil {
		return x.PrefferedContactMethod
	}
	return ""
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x1a, 0x12, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x12,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x72, 0x65, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x70, 0x72, 0x65, 0x66, 0x66, 0x65,
	0x72, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x42, 0x79, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x42,
	0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0xca, 0x02, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0xe2, 0x02, 0x13, 0x43, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),  // 0: carrier.Contact
	(CarrierRole)(0), // 1: carrier.CarrierRole
}
var file_contact_proto_depIdxs = []int32{
	1, // 0: carrier.Contact.roles:type_name -> carrier.CarrierRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	file_carrier_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}

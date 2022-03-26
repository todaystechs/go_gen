// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: address.proto

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

// type Address struct {
// 	StreetNumber                string `dynamodbav:"street_number" json:"street_number"`
// 	StreetName                  string `dynamodbav:"street_name" json:"street_name"`
// 	Municipality                string `dynamodbav:"municipality" json:"municipality"`
// 	CountrySecondarySubdivision string `dynamodbav:"country_secondary_subdivision"  json:"country_secondary_subdivision"`
// 	CountrySubdivision          string `dynamodbav:"country_subdivision" json:"country_subdivision"`
// 	CountrySubdivisionName      string `dynamodbav:"country_subdivision_name" json:"country_subdivision_name"`
// 	PostalCode                  string `dynamodbav:"postal_code" json:"postal_code"`
// 	State                       string `dynamodbav:"state" json:"state"`
// 	ExtendedPostalCode          string `dynamodbav:"extended_postal_code" json:"extended_postal_code"`
// 	CountryCode                 string `dynamodbav:"country_code" json:"country_code"`
// 	Country                     string `dynamodbav:"country" json:"country"`
// 	CountryCodeISO3             string `dynamodbav:"country_codeISO3" json:"country_codeISO3"`
// 	FreeformAddress             string `dynamodbav:"freeform_address" json:"freeform_address"`
// 	LocalName                   string `dynamodbav:"local_name" json:"local_name"`
// 	ID                          string `dynamodbav:"id" json:"id"`
// 	Type                        string `dynamodbav:"type" json:"type"`
// 	AddressLine2                string `dynamodbav:"address_line2" json:"address_line2"`
// 	AddressLine1                string `dynamodbav:"address_line1" json:"address_line1"`
// }
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreetNumber                string `protobuf:"bytes,1,opt,name=streetNumber,proto3" json:"streetNumber,omitempty"`
	StreetName                  string `protobuf:"bytes,2,opt,name=streetName,proto3" json:"streetName,omitempty"`
	Municipality                string `protobuf:"bytes,3,opt,name=municipality,proto3" json:"municipality,omitempty"`
	CountrySecondarySubdivision string `protobuf:"bytes,4,opt,name=CountrySecondarySubdivision,proto3" json:"CountrySecondarySubdivision,omitempty"`
	CountrySubdivision          string `protobuf:"bytes,5,opt,name=CountrySubdivision,proto3" json:"CountrySubdivision,omitempty"`
	CountrySubdivisionName      string `protobuf:"bytes,6,opt,name=countrySubdivisionName,proto3" json:"countrySubdivisionName,omitempty"`
	PostalCode                  string `protobuf:"bytes,7,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	State                       string `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	ExtendedPostalCode          string `protobuf:"bytes,9,opt,name=extendedPostalCode,proto3" json:"extendedPostalCode,omitempty"`
	CountryCode                 string `protobuf:"bytes,10,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	Country                     string `protobuf:"bytes,11,opt,name=country,proto3" json:"country,omitempty"`
	CountryCodeISO3             string `protobuf:"bytes,12,opt,name=countryCodeISO3,proto3" json:"countryCodeISO3,omitempty"`
	FreeFormAddress             string `protobuf:"bytes,13,opt,name=freeFormAddress,proto3" json:"freeFormAddress,omitempty"`
	LocalName                   string `protobuf:"bytes,14,opt,name=localName,proto3" json:"localName,omitempty"`
	Id                          string `protobuf:"bytes,15,opt,name=id,proto3" json:"id,omitempty"`
	Type                        string `protobuf:"bytes,16,opt,name=type,proto3" json:"type,omitempty"`
	AddressLine1                string `protobuf:"bytes,17,opt,name=addressLine1,proto3" json:"addressLine1,omitempty"`
	AddressLine2                string `protobuf:"bytes,18,opt,name=addressLine2,proto3" json:"addressLine2,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetStreetNumber() string {
	if x != nil {
		return x.StreetNumber
	}
	return ""
}

func (x *Address) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *Address) GetMunicipality() string {
	if x != nil {
		return x.Municipality
	}
	return ""
}

func (x *Address) GetCountrySecondarySubdivision() string {
	if x != nil {
		return x.CountrySecondarySubdivision
	}
	return ""
}

func (x *Address) GetCountrySubdivision() string {
	if x != nil {
		return x.CountrySubdivision
	}
	return ""
}

func (x *Address) GetCountrySubdivisionName() string {
	if x != nil {
		return x.CountrySubdivisionName
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetExtendedPostalCode() string {
	if x != nil {
		return x.ExtendedPostalCode
	}
	return ""
}

func (x *Address) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCountryCodeISO3() string {
	if x != nil {
		return x.CountryCodeISO3
	}
	return ""
}

func (x *Address) GetFreeFormAddress() string {
	if x != nil {
		return x.FreeFormAddress
	}
	return ""
}

func (x *Address) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Address) GetAddressLine1() string {
	if x != nil {
		return x.AddressLine1
	}
	return ""
}

func (x *Address) GetAddressLine2() string {
	if x != nil {
		return x.AddressLine2
	}
	return ""
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x22, 0x9b, 0x05, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x75, 0x6e, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x1b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79,
	0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x1b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x61, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x16, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x53, 0x4f, 0x33, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x53,
	0x4f, 0x33, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x65, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72, 0x65,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e,
	0x65, 0x32, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x42, 0x79, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x07,
	0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0xca, 0x02, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0xe2, 0x02, 0x13, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: carrier.Address
}
var file_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}

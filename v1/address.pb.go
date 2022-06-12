// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"address_line_1,omitempty"
	AddressLine_1 string `protobuf:"bytes,1,opt,name=address_line_1,json=addressLine1,proto3" json:"address_line_1,omitempty" dynamodbav:"address_line_1,omitempty"`
	// @gotags: dynamodbav:"address_line_2,omitempty"
	AddressLine_2 string `protobuf:"bytes,2,opt,name=address_line_2,json=addressLine2,proto3" json:"address_line_2,omitempty" dynamodbav:"address_line_2,omitempty"`
	// @gotags: dynamodbav:"street_name,omitempty"
	StreetName string `protobuf:"bytes,3,opt,name=street_name,json=streetName,proto3" json:"street_name,omitempty" dynamodbav:"street_name,omitempty"`
	// @gotags: dynamodbav:"city,omitempty"
	City string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty" dynamodbav:"city,omitempty"`
	// @gotags: dynamodbav:"county,omitempty"
	County string `protobuf:"bytes,5,opt,name=county,proto3" json:"county,omitempty" dynamodbav:"county,omitempty"`
	// @gotags: dynamodbav:"zip_code,omitempty"
	ZipCode string `protobuf:"bytes,6,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty" dynamodbav:"zip_code,omitempty"`
	// @gotags: dynamodbav:"state,omitempty"
	State string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty" dynamodbav:"state,omitempty"`
	// @gotags: dynamodbav:"extended_zip_code,omitempty"
	ExtendedZipCode string `protobuf:"bytes,8,opt,name=extended_zip_code,json=extendedZipCode,proto3" json:"extended_zip_code,omitempty" dynamodbav:"extended_zip_code,omitempty"`
	// @gotags: dynamodbav:"country,omitempty"
	Country string `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty" dynamodbav:"country,omitempty"`
	// @gotags: dynamodbav:"free_form_address,omitempty"
	FreeFormAddress string `protobuf:"bytes,10,opt,name=free_form_address,json=freeFormAddress,proto3" json:"free_form_address,omitempty" dynamodbav:"free_form_address,omitempty"`
	// @gotags: dynamodbav:"local_name,omitempty"
	LocalName string `protobuf:"bytes,11,opt,name=local_name,json=localName,proto3" json:"local_name,omitempty" dynamodbav:"local_name,omitempty"`
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

func (x *Address) GetAddressLine_1() string {
	if x != nil {
		return x.AddressLine_1
	}
	return ""
}

func (x *Address) GetAddressLine_2() string {
	if x != nil {
		return x.AddressLine_2
	}
	return ""
}

func (x *Address) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetExtendedZipCode() string {
	if x != nil {
		return x.ExtendedZipCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
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

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0xe4, 0x02, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x24, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x2a, 0x0a, 0x11, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72, 0x65, 0x65,
	0x46, 0x6f, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*Address)(nil), // 0: v1.address
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: carrier.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Carrier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,json=name,omitempty,proto3" json:"name,omitempty"`
	// @gotags: dynamodbav:"auth_url,omitempty"
	AuthUrl string `protobuf:"bytes,2,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	// @gotags: dynamodbav:"rate_url,omitempty"
	RateUrl string `protobuf:"bytes,3,opt,name=rate_url,json=rateUrl,proto3" json:"rate_url,omitempty"`
	// @gotags: dynamodbav:"add_address_url,omitempty"
	AddAddressUrl string `protobuf:"bytes,4,opt,name=add_address_url,json=addAddressUrl,proto3" json:"add_address_url,omitempty"`
	// @gotags: dynamodbav:"get_address_url,omitempty"
	GetAddressUrl string `protobuf:"bytes,5,opt,name=get_address_url,json=getAddressUrl,proto3" json:"get_address_url,omitempty"`
	// @gotags: dynamodbav:"quote_history_url,omitempty"
	QuoteHistoryUrl string `protobuf:"bytes,6,opt,name=quote_history_url,json=quoteHistoryUrl,proto3" json:"quote_history_url,omitempty"`
}

func (x *Carrier) Reset() {
	*x = Carrier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carrier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Carrier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Carrier) ProtoMessage() {}

func (x *Carrier) ProtoReflect() protoreflect.Message {
	mi := &file_carrier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Carrier.ProtoReflect.Descriptor instead.
func (*Carrier) Descriptor() ([]byte, []int) {
	return file_carrier_proto_rawDescGZIP(), []int{0}
}

func (x *Carrier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Carrier) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *Carrier) GetRateUrl() string {
	if x != nil {
		return x.RateUrl
	}
	return ""
}

func (x *Carrier) GetAddAddressUrl() string {
	if x != nil {
		return x.AddAddressUrl
	}
	return ""
}

func (x *Carrier) GetGetAddressUrl() string {
	if x != nil {
		return x.GetAddressUrl
	}
	return ""
}

func (x *Carrier) GetQuoteHistoryUrl() string {
	if x != nil {
		return x.QuoteHistoryUrl
	}
	return ""
}

var File_carrier_proto protoreflect.FileDescriptor

var file_carrier_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a, 0x07,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08,
	0x01, 0x52, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x61, 0x64, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x12,
	0x26, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x55, 0x72, 0x6c, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carrier_proto_rawDescOnce sync.Once
	file_carrier_proto_rawDescData = file_carrier_proto_rawDesc
)

func file_carrier_proto_rawDescGZIP() []byte {
	file_carrier_proto_rawDescOnce.Do(func() {
		file_carrier_proto_rawDescData = protoimpl.X.CompressGZIP(file_carrier_proto_rawDescData)
	})
	return file_carrier_proto_rawDescData
}

var file_carrier_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_carrier_proto_goTypes = []interface{}{
	(*Carrier)(nil), // 0: v1.carrier
}
var file_carrier_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carrier_proto_init() }
func file_carrier_proto_init() {
	if File_carrier_proto != nil {
		return
	}
	file_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_carrier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Carrier); i {
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
			RawDescriptor: file_carrier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_carrier_proto_goTypes,
		DependencyIndexes: file_carrier_proto_depIdxs,
		MessageInfos:      file_carrier_proto_msgTypes,
	}.Build()
	File_carrier_proto = out.File
	file_carrier_proto_rawDesc = nil
	file_carrier_proto_goTypes = nil
	file_carrier_proto_depIdxs = nil
}

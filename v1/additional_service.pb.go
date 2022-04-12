// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: additional_service.proto

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

type AdditionalService int32

const (
	AdditionalService_InsidePickup             AdditionalService = 0
	AdditionalService_InsideDelivery           AdditionalService = 1
	AdditionalService_PickupNotification       AdditionalService = 2
	AdditionalService_DeliveryNotification     AdditionalService = 3
	AdditionalService_LiftgatePickup           AdditionalService = 4
	AdditionalService_LiftGateDelivery         AdditionalService = 5
	AdditionalService_ProtectFromFreeze        AdditionalService = 6
	AdditionalService_SortAndSegregateDelivery AdditionalService = 7
)

// Enum value maps for AdditionalService.
var (
	AdditionalService_name = map[int32]string{
		0: "InsidePickup",
		1: "InsideDelivery",
		2: "PickupNotification",
		3: "DeliveryNotification",
		4: "LiftgatePickup",
		5: "LiftGateDelivery",
		6: "ProtectFromFreeze",
		7: "SortAndSegregateDelivery",
	}
	AdditionalService_value = map[string]int32{
		"InsidePickup":             0,
		"InsideDelivery":           1,
		"PickupNotification":       2,
		"DeliveryNotification":     3,
		"LiftgatePickup":           4,
		"LiftGateDelivery":         5,
		"ProtectFromFreeze":        6,
		"SortAndSegregateDelivery": 7,
	}
)

func (x AdditionalService) Enum() *AdditionalService {
	p := new(AdditionalService)
	*p = x
	return p
}

func (x AdditionalService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdditionalService) Descriptor() protoreflect.EnumDescriptor {
	return file_additional_service_proto_enumTypes[0].Descriptor()
}

func (AdditionalService) Type() protoreflect.EnumType {
	return &file_additional_service_proto_enumTypes[0]
}

func (x AdditionalService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdditionalService.Descriptor instead.
func (AdditionalService) EnumDescriptor() ([]byte, []int) {
	return file_additional_service_proto_rawDescGZIP(), []int{0}
}

var File_additional_service_proto protoreflect.FileDescriptor

var file_additional_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x2a, 0xcb, 0x01, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x69, 0x64,
	0x65, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x73,
	0x69, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x10, 0x01, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12,
	0x12, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x69, 0x66, 0x74, 0x47, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x10, 0x06,
	0x12, 0x1c, 0x0a, 0x18, 0x53, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x10, 0x07, 0x42, 0x74,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x16, 0x41, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_additional_service_proto_rawDescOnce sync.Once
	file_additional_service_proto_rawDescData = file_additional_service_proto_rawDesc
)

func file_additional_service_proto_rawDescGZIP() []byte {
	file_additional_service_proto_rawDescOnce.Do(func() {
		file_additional_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_additional_service_proto_rawDescData)
	})
	return file_additional_service_proto_rawDescData
}

var file_additional_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_additional_service_proto_goTypes = []interface{}{
	(AdditionalService)(0), // 0: user.additional_service
}
var file_additional_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_additional_service_proto_init() }
func file_additional_service_proto_init() {
	if File_additional_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_additional_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_additional_service_proto_goTypes,
		DependencyIndexes: file_additional_service_proto_depIdxs,
		EnumInfos:         file_additional_service_proto_enumTypes,
	}.Build()
	File_additional_service_proto = out.File
	file_additional_service_proto_rawDesc = nil
	file_additional_service_proto_goTypes = nil
	file_additional_service_proto_depIdxs = nil
}

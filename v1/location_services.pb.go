// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: location_services.proto

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

type LocationServices int32

const (
	// @gotags: dynamodbav:"pickup_location_with_dock,omitempty"
	LocationServices_pickup_location_with_dock LocationServices = 0
	// @gotags: dynamodbav:"liftgate_pickup,omitempty"
	LocationServices_liftgate_pickup LocationServices = 1
	// @gotags: dynamodbav:"pickup_appointment,omitempty"
	LocationServices_pickup_appointment LocationServices = 2
	// @gotags: dynamodbav:"inside_pickup,omitempty"
	LocationServices_inside_pickup LocationServices = 3
	// @gotags: dynamodbav:"pickup_notification,omitempty"
	LocationServices_pickup_notification LocationServices = 4
	// @gotags: dynamodbav:"delivery_location_with_dock,omitempty"
	LocationServices_delivery_location_with_dock LocationServices = 5
	// @gotags: dynamodbav:"liftgate_delivery,omitempty"
	LocationServices_liftgate_delivery LocationServices = 6
	// @gotags: dynamodbav:"delivery_appointment,omitempty"
	LocationServices_delivery_appointment LocationServices = 7
	// @gotags: dynamodbav:"inside_delivery,omitempty"
	LocationServices_inside_delivery LocationServices = 8
	// @gotags: dynamodbav:"delivery_notification,omitempty"
	LocationServices_delivery_notification LocationServices = 9
)

// Enum value maps for LocationServices.
var (
	LocationServices_name = map[int32]string{
		0: "pickup_location_with_dock",
		1: "liftgate_pickup",
		2: "pickup_appointment",
		3: "inside_pickup",
		4: "pickup_notification",
		5: "delivery_location_with_dock",
		6: "liftgate_delivery",
		7: "delivery_appointment",
		8: "inside_delivery",
		9: "delivery_notification",
	}
	LocationServices_value = map[string]int32{
		"pickup_location_with_dock":   0,
		"liftgate_pickup":             1,
		"pickup_appointment":          2,
		"inside_pickup":               3,
		"pickup_notification":         4,
		"delivery_location_with_dock": 5,
		"liftgate_delivery":           6,
		"delivery_appointment":        7,
		"inside_delivery":             8,
		"delivery_notification":       9,
	}
)

func (x LocationServices) Enum() *LocationServices {
	p := new(LocationServices)
	*p = x
	return p
}

func (x LocationServices) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LocationServices) Descriptor() protoreflect.EnumDescriptor {
	return file_location_services_proto_enumTypes[0].Descriptor()
}

func (LocationServices) Type() protoreflect.EnumType {
	return &file_location_services_proto_enumTypes[0]
}

func (x LocationServices) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LocationServices.Descriptor instead.
func (LocationServices) EnumDescriptor() ([]byte, []int) {
	return file_location_services_proto_rawDescGZIP(), []int{0}
}

var File_location_services_proto protoreflect.FileDescriptor

var file_location_services_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x2a, 0x8d, 0x02,
	0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x63, 0x6b,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x6c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x02, 0x12,
	0x11, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x63, 0x6b, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11,
	0x6c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x10, 0x06, 0x12, 0x18, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x07, 0x12, 0x13, 0x0a,
	0x0f, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x09, 0x42, 0x69, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64,
	0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02,
	0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_services_proto_rawDescOnce sync.Once
	file_location_services_proto_rawDescData = file_location_services_proto_rawDesc
)

func file_location_services_proto_rawDescGZIP() []byte {
	file_location_services_proto_rawDescOnce.Do(func() {
		file_location_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_services_proto_rawDescData)
	})
	return file_location_services_proto_rawDescData
}

var file_location_services_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_location_services_proto_goTypes = []interface{}{
	(LocationServices)(0), // 0: v1.location_services
}
var file_location_services_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_location_services_proto_init() }
func file_location_services_proto_init() {
	if File_location_services_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_location_services_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_location_services_proto_goTypes,
		DependencyIndexes: file_location_services_proto_depIdxs,
		EnumInfos:         file_location_services_proto_enumTypes,
	}.Build()
	File_location_services_proto = out.File
	file_location_services_proto_rawDesc = nil
	file_location_services_proto_goTypes = nil
	file_location_services_proto_depIdxs = nil
}

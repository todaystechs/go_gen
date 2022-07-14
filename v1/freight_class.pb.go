// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: freight_class.proto

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

type FreightClass int32

const (
	FreightClass_Class50  FreightClass = 0
	FreightClass_Class55  FreightClass = 1
	FreightClass_Class60  FreightClass = 2
	FreightClass_Class65  FreightClass = 3
	FreightClass_Class70  FreightClass = 4
	FreightClass_Class775 FreightClass = 5
	FreightClass_Class85  FreightClass = 6
	FreightClass_Class925 FreightClass = 7
	FreightClass_Class100 FreightClass = 8
	FreightClass_Class110 FreightClass = 9
	FreightClass_Class125 FreightClass = 10
	FreightClass_Class150 FreightClass = 11
	FreightClass_Class175 FreightClass = 12
	FreightClass_Class200 FreightClass = 13
	FreightClass_Class250 FreightClass = 14
	FreightClass_Class300 FreightClass = 15
	FreightClass_Class400 FreightClass = 16
	FreightClass_Class500 FreightClass = 17
	FreightClass_Class18  FreightClass = 18
)

// Enum value maps for FreightClass.
var (
	FreightClass_name = map[int32]string{
		0:  "Class50",
		1:  "Class55",
		2:  "Class60",
		3:  "Class65",
		4:  "Class70",
		5:  "Class775",
		6:  "Class85",
		7:  "Class925",
		8:  "Class100",
		9:  "Class110",
		10: "Class125",
		11: "Class150",
		12: "Class175",
		13: "Class200",
		14: "Class250",
		15: "Class300",
		16: "Class400",
		17: "Class500",
		18: "Class18",
	}
	FreightClass_value = map[string]int32{
		"Class50":  0,
		"Class55":  1,
		"Class60":  2,
		"Class65":  3,
		"Class70":  4,
		"Class775": 5,
		"Class85":  6,
		"Class925": 7,
		"Class100": 8,
		"Class110": 9,
		"Class125": 10,
		"Class150": 11,
		"Class175": 12,
		"Class200": 13,
		"Class250": 14,
		"Class300": 15,
		"Class400": 16,
		"Class500": 17,
		"Class18":  18,
	}
)

func (x FreightClass) Enum() *FreightClass {
	p := new(FreightClass)
	*p = x
	return p
}

func (x FreightClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FreightClass) Descriptor() protoreflect.EnumDescriptor {
	return file_freight_class_proto_enumTypes[0].Descriptor()
}

func (FreightClass) Type() protoreflect.EnumType {
	return &file_freight_class_proto_enumTypes[0]
}

func (x FreightClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FreightClass.Descriptor instead.
func (FreightClass) EnumDescriptor() ([]byte, []int) {
	return file_freight_class_proto_rawDescGZIP(), []int{0}
}

var File_freight_class_proto protoreflect.FileDescriptor

var file_freight_class_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x2a, 0x92, 0x02, 0x0a, 0x0d, 0x66, 0x72,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x35, 0x30, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x35, 0x35, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x36, 0x30,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x36, 0x35, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x37, 0x30, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x37, 0x37, 0x35, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x38, 0x35, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x39, 0x32, 0x35, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x30,
	0x30, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x31, 0x30, 0x10,
	0x09, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x32, 0x35, 0x10, 0x0a, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x35, 0x30, 0x10, 0x0b, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x37, 0x35, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x32, 0x30, 0x30, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x32, 0x35, 0x30, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x33, 0x30, 0x30, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x34, 0x30,
	0x30, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x35, 0x30, 0x30, 0x10,
	0x11, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x38, 0x10, 0x12, 0x42, 0x65,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2,
	0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_freight_class_proto_rawDescOnce sync.Once
	file_freight_class_proto_rawDescData = file_freight_class_proto_rawDesc
)

func file_freight_class_proto_rawDescGZIP() []byte {
	file_freight_class_proto_rawDescOnce.Do(func() {
		file_freight_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_freight_class_proto_rawDescData)
	})
	return file_freight_class_proto_rawDescData
}

var file_freight_class_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_freight_class_proto_goTypes = []interface{}{
	(FreightClass)(0), // 0: v1.freight_class
}
var file_freight_class_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_freight_class_proto_init() }
func file_freight_class_proto_init() {
	if File_freight_class_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_freight_class_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_freight_class_proto_goTypes,
		DependencyIndexes: file_freight_class_proto_depIdxs,
		EnumInfos:         file_freight_class_proto_enumTypes,
	}.Build()
	File_freight_class_proto = out.File
	file_freight_class_proto_rawDesc = nil
	file_freight_class_proto_goTypes = nil
	file_freight_class_proto_depIdxs = nil
}

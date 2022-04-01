// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: commodity.proto

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

type DimensionUOM int32

const (
	DimensionUOM_INCH DimensionUOM = 0
	DimensionUOM_CM   DimensionUOM = 1
)

// Enum value maps for DimensionUOM.
var (
	DimensionUOM_name = map[int32]string{
		0: "INCH",
		1: "CM",
	}
	DimensionUOM_value = map[string]int32{
		"INCH": 0,
		"CM":   1,
	}
)

func (x DimensionUOM) Enum() *DimensionUOM {
	p := new(DimensionUOM)
	*p = x
	return p
}

func (x DimensionUOM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DimensionUOM) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[0].Descriptor()
}

func (DimensionUOM) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[0]
}

func (x DimensionUOM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DimensionUOM.Descriptor instead.
func (DimensionUOM) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

type WeightUOM int32

const (
	WeightUOM_LB WeightUOM = 0
	WeightUOM_KG WeightUOM = 1
)

// Enum value maps for WeightUOM.
var (
	WeightUOM_name = map[int32]string{
		0: "LB",
		1: "KG",
	}
	WeightUOM_value = map[string]int32{
		"LB": 0,
		"KG": 1,
	}
)

func (x WeightUOM) Enum() *WeightUOM {
	p := new(WeightUOM)
	*p = x
	return p
}

func (x WeightUOM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeightUOM) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[1].Descriptor()
}

func (WeightUOM) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[1]
}

func (x WeightUOM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeightUOM.Descriptor instead.
func (WeightUOM) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

type Commodity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"density,omitempty"
	Density float32 `protobuf:"fixed32,1,opt,name=density,proto3" json:"density,omitempty" dynamodbav:"density,omitempty"`
	// @gotags: dynamodbav:"length,omitempty"
	Length float32 `protobuf:"fixed32,2,opt,name=length,proto3" json:"length,omitempty" dynamodbav:"length,omitempty"`
	// @gotags: dynamodbav:"width,omitempty"
	Width float32 `protobuf:"fixed32,3,opt,name=width,proto3" json:"width,omitempty" dynamodbav:"width,omitempty"`
	// @gotags: dynamodbav:"height,omitempty"
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty" dynamodbav:"height,omitempty"`
	// @gotags: dynamodbav:"weight,omitempty"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty" dynamodbav:"weight,omitempty"`
	// @gotags: dynamodbav:"dimension_uom,omitempty"
	Dimension_UOM DimensionUOM `protobuf:"varint,6,opt,name=dimension_UOM,json=dimensionUOM,proto3,enum=user.DimensionUOM" json:"dimension_UOM,omitempty" dynamodbav:"dimension_uom,omitempty"`
	// @gotags: dynamodbav:"weight_uom,omitempty"
	Weight_UOM WeightUOM `protobuf:"varint,7,opt,name=weight_UOM,json=weightUOM,proto3,enum=user.WeightUOM" json:"weight_UOM,omitempty" dynamodbav:"weight_uom,omitempty"`
	// @gotags: dynamodbav:"dimension_display,omitempty"
	DimensionDisplay string `protobuf:"bytes,8,opt,name=dimension_display,json=dimensionDisplay,proto3" json:"dimension_display,omitempty" dynamodbav:"dimension_display,omitempty"`
	// @gotags: dynamodbav:"package_type,omitempty"
	PackageType PackageType `protobuf:"varint,9,opt,name=package_type,json=packageType,proto3,enum=user.PackageType" json:"package_type,omitempty" dynamodbav:"package_type,omitempty"`
	// @gotags: dynamodbav:"quantity,omitempty"
	Quantity int32 `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" dynamodbav:"quantity,omitempty"`
	// @gotags: dynamodbav:"freight_class,omitempty"
	FreightClass FreightClass `protobuf:"varint,11,opt,name=freight_class,json=freightClass,proto3,enum=user.FreightClass" json:"freight_class,omitempty" dynamodbav:"freight_class,omitempty"`
}

func (x *Commodity) Reset() {
	*x = Commodity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commodity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commodity) ProtoMessage() {}

func (x *Commodity) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commodity.ProtoReflect.Descriptor instead.
func (*Commodity) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

func (x *Commodity) GetDensity() float32 {
	if x != nil {
		return x.Density
	}
	return 0
}

func (x *Commodity) GetLength() float32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Commodity) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Commodity) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Commodity) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Commodity) GetDimension_UOM() DimensionUOM {
	if x != nil {
		return x.Dimension_UOM
	}
	return DimensionUOM_INCH
}

func (x *Commodity) GetWeight_UOM() WeightUOM {
	if x != nil {
		return x.Weight_UOM
	}
	return WeightUOM_LB
}

func (x *Commodity) GetDimensionDisplay() string {
	if x != nil {
		return x.DimensionDisplay
	}
	return ""
}

func (x *Commodity) GetPackageType() PackageType {
	if x != nil {
		return x.PackageType
	}
	return PackageType_Pallet
}

func (x *Commodity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Commodity) GetFreightClass() FreightClass {
	if x != nil {
		return x.FreightClass
	}
	return FreightClass_Class50
}

var File_commodity_proto protoreflect.FileDescriptor

var file_commodity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x03, 0x0a, 0x09,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x64, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x37, 0x0a, 0x0d, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x55,
	0x4f, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x52, 0x0c, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x12, 0x2e, 0x0a, 0x0a, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x55, 0x4f, 0x4d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f, 0x4d, 0x52,
	0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f, 0x4d, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69,
	0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x34, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x37, 0x0a, 0x0d, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x2a, 0x20, 0x0a, 0x0c, 0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55,
	0x4f, 0x4d, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x43, 0x48, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x43, 0x4d, 0x10, 0x01, 0x2a, 0x1b, 0x0a, 0x09, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f,
	0x4d, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x42, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x47, 0x10,
	0x01, 0x42, 0x6c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02,
	0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commodity_proto_rawDescOnce sync.Once
	file_commodity_proto_rawDescData = file_commodity_proto_rawDesc
)

func file_commodity_proto_rawDescGZIP() []byte {
	file_commodity_proto_rawDescOnce.Do(func() {
		file_commodity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commodity_proto_rawDescData)
	})
	return file_commodity_proto_rawDescData
}

var file_commodity_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_commodity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commodity_proto_goTypes = []interface{}{
	(DimensionUOM)(0), // 0: user.DimensionUOM
	(WeightUOM)(0),    // 1: user.WeightUOM
	(*Commodity)(nil), // 2: user.Commodity
	(PackageType)(0),  // 3: user.PackageType
	(FreightClass)(0), // 4: user.FreightClass
}
var file_commodity_proto_depIdxs = []int32{
	0, // 0: user.Commodity.dimension_UOM:type_name -> user.DimensionUOM
	1, // 1: user.Commodity.weight_UOM:type_name -> user.WeightUOM
	3, // 2: user.Commodity.package_type:type_name -> user.PackageType
	4, // 3: user.Commodity.freight_class:type_name -> user.FreightClass
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_commodity_proto_init() }
func file_commodity_proto_init() {
	if File_commodity_proto != nil {
		return
	}
	file_package_proto_init()
	file_freight_class_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_commodity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commodity); i {
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
			RawDescriptor: file_commodity_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commodity_proto_goTypes,
		DependencyIndexes: file_commodity_proto_depIdxs,
		EnumInfos:         file_commodity_proto_enumTypes,
		MessageInfos:      file_commodity_proto_msgTypes,
	}.Build()
	File_commodity_proto = out.File
	file_commodity_proto_rawDesc = nil
	file_commodity_proto_goTypes = nil
	file_commodity_proto_depIdxs = nil
}

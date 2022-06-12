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

type DimensionUom int32

const (
	DimensionUom_INCH DimensionUom = 0
	DimensionUom_CM   DimensionUom = 1
)

// Enum value maps for DimensionUom.
var (
	DimensionUom_name = map[int32]string{
		0: "INCH",
		1: "CM",
	}
	DimensionUom_value = map[string]int32{
		"INCH": 0,
		"CM":   1,
	}
)

func (x DimensionUom) Enum() *DimensionUom {
	p := new(DimensionUom)
	*p = x
	return p
}

func (x DimensionUom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DimensionUom) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[0].Descriptor()
}

func (DimensionUom) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[0]
}

func (x DimensionUom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DimensionUom.Descriptor instead.
func (DimensionUom) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

type WeightUom int32

const (
	WeightUom_LB WeightUom = 0
	WeightUom_KG WeightUom = 1
)

// Enum value maps for WeightUom.
var (
	WeightUom_name = map[int32]string{
		0: "LB",
		1: "KG",
	}
	WeightUom_value = map[string]int32{
		"LB": 0,
		"KG": 1,
	}
)

func (x WeightUom) Enum() *WeightUom {
	p := new(WeightUom)
	*p = x
	return p
}

func (x WeightUom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeightUom) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[1].Descriptor()
}

func (WeightUom) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[1]
}

func (x WeightUom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeightUom.Descriptor instead.
func (WeightUom) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

type CommodityServices int32

const (
	// @gotags: dynamodbav:"stackable,omitempty"
	CommodityServices_stackable CommodityServices = 0
	// @gotags: dynamodbav:"protect_from_freeze,omitempty"
	CommodityServices_protect_from_freeze CommodityServices = 1
	// @gotags: dynamodbav:"sort_and_segregate,omitempty"
	CommodityServices_sort_and_segregate CommodityServices = 2
	// @gotags: dynamodbav:"guaranteed,omitempty"
	CommodityServices_guaranteed CommodityServices = 3
	// @gotags: dynamodbav:"hazardous,omitempty"
	CommodityServices_hazardous CommodityServices = 4
	// @gotags: dynamodbav:"commodity_instructions,omitempty"
	CommodityServices_commodity_instructions CommodityServices = 5
)

// Enum value maps for CommodityServices.
var (
	CommodityServices_name = map[int32]string{
		0: "stackable",
		1: "protect_from_freeze",
		2: "sort_and_segregate",
		3: "guaranteed",
		4: "hazardous",
		5: "commodity_instructions",
	}
	CommodityServices_value = map[string]int32{
		"stackable":              0,
		"protect_from_freeze":    1,
		"sort_and_segregate":     2,
		"guaranteed":             3,
		"hazardous":              4,
		"commodity_instructions": 5,
	}
)

func (x CommodityServices) Enum() *CommodityServices {
	p := new(CommodityServices)
	*p = x
	return p
}

func (x CommodityServices) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommodityServices) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[2].Descriptor()
}

func (CommodityServices) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[2]
}

func (x CommodityServices) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommodityServices.Descriptor instead.
func (CommodityServices) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{2}
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
	DimensionUom DimensionUom `protobuf:"varint,6,opt,name=dimension_uom,json=dimensionUom,proto3,enum=user.DimensionUom" json:"dimension_uom,omitempty" dynamodbav:"dimension_uom,omitempty"`
	// @gotags: dynamodbav:"weight_uom,omitempty"
	WeightUom WeightUom `protobuf:"varint,7,opt,name=weight_uom,json=weightUom,proto3,enum=user.WeightUom" json:"weight_uom,omitempty" dynamodbav:"weight_uom,omitempty"`
	// @gotags: dynamodbav:"dimension_display,omitempty"
	DimensionDisplay string `protobuf:"bytes,8,opt,name=dimension_display,json=dimensionDisplay,proto3" json:"dimension_display,omitempty" dynamodbav:"dimension_display,omitempty"`
	// @gotags: dynamodbav:"package_type,omitempty"
	PackageType PackageType `protobuf:"varint,9,opt,name=package_type,json=packageType,proto3,enum=user.PackageType" json:"package_type,omitempty" dynamodbav:"package_type,omitempty"`
	// @gotags: dynamodbav:"quantity,omitempty"
	Quantity int32 `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" dynamodbav:"quantity,omitempty"`
	// @gotags: dynamodbav:"freight_class,omitempty"
	FreightClass FreightClass `protobuf:"varint,11,opt,name=freight_class,json=freightClass,proto3,enum=user.FreightClass" json:"freight_class,omitempty" dynamodbav:"freight_class,omitempty"`
	// @gotags: dynamodbav:"stackable,omitempty"
	Stackable bool `protobuf:"varint,12,opt,name=stackable,proto3" json:"stackable,omitempty" dynamodbav:"stackable,omitempty"`
	// @gotags: dynamodbav:"protect_from_freeze,omitempty"
	ProtectFromFreeze bool `protobuf:"varint,13,opt,name=protect_from_freeze,json=protectFromFreeze,proto3" json:"protect_from_freeze,omitempty" dynamodbav:"protect_from_freeze,omitempty"`
	// @gotags: dynamodbav:"sort_and_segregate,omitempty"
	SortAndSegregate bool `protobuf:"varint,14,opt,name=sort_and_segregate,json=sortAndSegregate,proto3" json:"sort_and_segregate,omitempty" dynamodbav:"sort_and_segregate,omitempty"`
	// @gotags: dynamodbav:"guaranteed,omitempty"
	Guaranteed bool `protobuf:"varint,15,opt,name=guaranteed,proto3" json:"guaranteed,omitempty" dynamodbav:"guaranteed,omitempty"`
	// @gotags: dynamodbav:"hazardous,omitempty"
	Hazardous bool `protobuf:"varint,16,opt,name=hazardous,proto3" json:"hazardous,omitempty" dynamodbav:"hazardous,omitempty"`
	// @gotags: dynamodbav:"commodity_instructions,omitempty"
	CommodityInstructions string `protobuf:"bytes,17,opt,name=commodity_instructions,json=commodityInstructions,proto3" json:"commodity_instructions,omitempty" dynamodbav:"commodity_instructions,omitempty"`
	// @gotags: dynamodbav:"commodity_services,omitempty"
	CommodityServices []CommodityServices `protobuf:"varint,18,rep,packed,name=commodity_services,json=commodityServices,proto3,enum=user.CommodityServices" json:"commodity_services,omitempty" dynamodbav:"commodity_services,omitempty"`
	// @gotags: dynamodbav:"index,omitempty"
	Index string `protobuf:"bytes,19,opt,name=index,proto3" json:"index,omitempty" dynamodbav:"index,omitempty"`
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

func (x *Commodity) GetDimensionUom() DimensionUom {
	if x != nil {
		return x.DimensionUom
	}
	return DimensionUom_INCH
}

func (x *Commodity) GetWeightUom() WeightUom {
	if x != nil {
		return x.WeightUom
	}
	return WeightUom_LB
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
	return PackageType_pallet
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

func (x *Commodity) GetStackable() bool {
	if x != nil {
		return x.Stackable
	}
	return false
}

func (x *Commodity) GetProtectFromFreeze() bool {
	if x != nil {
		return x.ProtectFromFreeze
	}
	return false
}

func (x *Commodity) GetSortAndSegregate() bool {
	if x != nil {
		return x.SortAndSegregate
	}
	return false
}

func (x *Commodity) GetGuaranteed() bool {
	if x != nil {
		return x.Guaranteed
	}
	return false
}

func (x *Commodity) GetHazardous() bool {
	if x != nil {
		return x.Hazardous
	}
	return false
}

func (x *Commodity) GetCommodityInstructions() string {
	if x != nil {
		return x.CommodityInstructions
	}
	return ""
}

func (x *Commodity) GetCommodityServices() []CommodityServices {
	if x != nil {
		return x.CommodityServices
	}
	return nil
}

func (x *Commodity) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

var File_commodity_proto protoreflect.FileDescriptor

var file_commodity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x05, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x64, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x38, 0x0a, 0x0d, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x6f, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6f, 0x6d, 0x52, 0x0c, 0x64,
	0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x6f, 0x6d, 0x12, 0x2f, 0x0a, 0x0a, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x6f, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x75, 0x6f,
	0x6d, 0x52, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x6f, 0x6d, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x35, 0x0a, 0x0c, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0d,
	0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x66, 0x72, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x72,
	0x65, 0x65, 0x7a, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x6e, 0x64,
	0x5f, 0x73, 0x65, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x10, 0x73, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x67, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x7a, 0x61, 0x72, 0x64, 0x6f, 0x75, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x7a, 0x61, 0x72, 0x64, 0x6f, 0x75, 0x73,
	0x12, 0x35, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x47, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2a, 0x21, 0x0a, 0x0d, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x6f, 0x6d, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x43, 0x48, 0x10,
	0x00, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x4d, 0x10, 0x01, 0x2a, 0x1c, 0x0a, 0x0a, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x75, 0x6f, 0x6d, 0x12, 0x06, 0x0a, 0x02, 0x4c, 0x42, 0x10, 0x00, 0x12,
	0x06, 0x0a, 0x02, 0x4b, 0x47, 0x10, 0x01, 0x2a, 0x8f, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x0d,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x66, 0x72,
	0x65, 0x65, 0x7a, 0x65, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x61,
	0x6e, 0x64, 0x5f, 0x73, 0x65, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x67, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x68, 0x61, 0x7a, 0x61, 0x72, 0x64, 0x6f, 0x75, 0x73, 0x10, 0x04, 0x12, 0x1a, 0x0a,
	0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x10, 0x05, 0x42, 0x6c, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f,
	0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa,
	0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10,
	0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_commodity_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_commodity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_commodity_proto_goTypes = []interface{}{
	(DimensionUom)(0),      // 0: user.dimension_uom
	(WeightUom)(0),         // 1: user.weight_uom
	(CommodityServices)(0), // 2: user.commodity_services
	(*Commodity)(nil),      // 3: user.commodity
	(PackageType)(0),       // 4: user.package_type
	(FreightClass)(0),      // 5: user.freight_class
}
var file_commodity_proto_depIdxs = []int32{
	0, // 0: user.commodity.dimension_uom:type_name -> user.dimension_uom
	1, // 1: user.commodity.weight_uom:type_name -> user.weight_uom
	4, // 2: user.commodity.package_type:type_name -> user.package_type
	5, // 3: user.commodity.freight_class:type_name -> user.freight_class
	2, // 4: user.commodity.commodity_services:type_name -> user.commodity_services
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
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
			NumEnums:      3,
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

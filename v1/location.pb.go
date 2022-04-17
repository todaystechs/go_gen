// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: location.proto

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"id,omitempty"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id,omitempty"`
	// @gotags: dynamodbav:"name,omitempty"
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" dynamodbav:"name,omitempty"`
	// @gotags: dynamodbav:"address,omitempty"
	Address *Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" dynamodbav:"address,omitempty"`
	// @gotags: dynamodbav:"contact,omitempty"
	Contact *Contact `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty" dynamodbav:"contact,omitempty"`
	// @gotags: dynamodbav:"location_type,omitempty"
	LocationType LocationType `protobuf:"varint,5,opt,name=location_type,json=locationType,proto3,enum=user.LocationType" json:"location_type,omitempty" dynamodbav:"location_type,omitempty"`
	// @gotags: dynamodbav:"additional_service,omitempty"
	AdditionalService []AdditionalService `protobuf:"varint,6,rep,packed,name=additional_service,json=additionalService,proto3,enum=user.AdditionalService" json:"additional_service,omitempty" dynamodbav:"additional_service,omitempty"`
	// @gotags: dynamodbav:"business_hours,omitempty",
	BusinessHours *BusinessHours `protobuf:"bytes,7,opt,name=business_hours,json=businessHours,proto3" json:"business_hours,omitempty" dynamodbav:"business_hours,omitempty"`
	// @gotags: dynamodbav:"is_default_pick_up,omitempty"
	IsDefaultPickUp bool `protobuf:"varint,8,opt,name=is_default_pick_up,json=isDefaultPickUp,proto3" json:"is_default_pick_up,omitempty" dynamodbav:"is_default_pick_up,omitempty"`
	// @gotags: dynamodbav:"is_default_delivery,omitempty"
	IsDefaultDelivery bool `protobuf:"varint,9,opt,name=is_default_delivery,json=isDefaultDelivery,proto3" json:"is_default_delivery,omitempty" dynamodbav:"is_default_delivery,omitempty"`
	// @gotags: dynamodbav:"business_pk,omitempty"
	BusinessPk string `protobuf:"bytes,10,opt,name=business_pk,json=businessPk,proto3" json:"business_pk,omitempty" dynamodbav:"business_pk,omitempty"`
	// @gotags: dynamodbav:"business_sk,omitempty"
	BusinessSk string `protobuf:"bytes,11,opt,name=business_sk,json=businessSk,proto3" json:"business_sk,omitempty" dynamodbav:"business_sk,omitempty"`
	// @gotags: dynamodbav:"location_pk,omitempty"
	LocationPk string `protobuf:"bytes,12,opt,name=location_pk,json=locationPk,proto3" json:"location_pk,omitempty" dynamodbav:"location_pk,omitempty"`
	// @gotags: dynamodbav:"location_sk,omitempty"
	LocationSk string `protobuf:"bytes,13,opt,name=location_sk,json=locationSk,proto3" json:"location_sk,omitempty" dynamodbav:"location_sk,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Location) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Location) GetLocationType() LocationType {
	if x != nil {
		return x.LocationType
	}
	return LocationType_BusinessWithDock
}

func (x *Location) GetAdditionalService() []AdditionalService {
	if x != nil {
		return x.AdditionalService
	}
	return nil
}

func (x *Location) GetBusinessHours() *BusinessHours {
	if x != nil {
		return x.BusinessHours
	}
	return nil
}

func (x *Location) GetIsDefaultPickUp() bool {
	if x != nil {
		return x.IsDefaultPickUp
	}
	return false
}

func (x *Location) GetIsDefaultDelivery() bool {
	if x != nil {
		return x.IsDefaultDelivery
	}
	return false
}

func (x *Location) GetBusinessPk() string {
	if x != nil {
		return x.BusinessPk
	}
	return ""
}

func (x *Location) GetBusinessSk() string {
	if x != nil {
		return x.BusinessSk
	}
	return ""
}

func (x *Location) GetLocationPk() string {
	if x != nil {
		return x.LocationPk
	}
	return ""
}

func (x *Location) GetLocationSk() string {
	if x != nil {
		return x.LocationSk
	}
	return ""
}

type Locations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:locations,omitempty"
	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *Locations) Reset() {
	*x = Locations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locations) ProtoMessage() {}

func (x *Locations) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locations.ProtoReflect.Descriptor instead.
func (*Locations) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{1}
}

func (x *Locations) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

var File_location_proto protoreflect.FileDescriptor

var file_location_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x13, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x08, 0x6f, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x38, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x47, 0x0a, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x11, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x69, 0x63,
	0x6b, 0x55, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x70, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x50, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x73, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6b, 0x22, 0x39, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x6b, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0d,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61,
	0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02,
	0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_proto_rawDescOnce sync.Once
	file_location_proto_rawDescData = file_location_proto_rawDesc
)

func file_location_proto_rawDescGZIP() []byte {
	file_location_proto_rawDescOnce.Do(func() {
		file_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_proto_rawDescData)
	})
	return file_location_proto_rawDescData
}

var file_location_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_location_proto_goTypes = []interface{}{
	(*Location)(nil),       // 0: user.location
	(*Locations)(nil),      // 1: user.locations
	(*Address)(nil),        // 2: user.address
	(*Contact)(nil),        // 3: user.contact
	(LocationType)(0),      // 4: user.location_type
	(AdditionalService)(0), // 5: user.additional_service
	(*BusinessHours)(nil),  // 6: user.business_hours
}
var file_location_proto_depIdxs = []int32{
	2, // 0: user.location.address:type_name -> user.address
	3, // 1: user.location.contact:type_name -> user.contact
	4, // 2: user.location.location_type:type_name -> user.location_type
	5, // 3: user.location.additional_service:type_name -> user.additional_service
	6, // 4: user.location.business_hours:type_name -> user.business_hours
	0, // 5: user.locations.locations:type_name -> user.location
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_location_proto_init() }
func file_location_proto_init() {
	if File_location_proto != nil {
		return
	}
	file_location_type_proto_init()
	file_additional_service_proto_init()
	file_address_proto_init()
	file_contact_proto_init()
	file_business_hours_proto_init()
	file_ok_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Locations); i {
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
			RawDescriptor: file_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_location_proto_goTypes,
		DependencyIndexes: file_location_proto_depIdxs,
		MessageInfos:      file_location_proto_msgTypes,
	}.Build()
	File_location_proto = out.File
	file_location_proto_rawDesc = nil
	file_location_proto_goTypes = nil
	file_location_proto_depIdxs = nil
}

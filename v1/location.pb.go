// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: location.proto

package v1

import (
	_ "github.com/todaystechs/go_gen/v1/validate"
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
	// @gotags: dynamodbav:"company_name,omitempty"
	CompanyName string `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty" dynamodbav:"company_name,omitempty"`
	// @gotags: dynamodbav:"address,omitempty"
	Address *Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty" dynamodbav:"address,omitempty"`
	// @gotags: dynamodbav:"contact,omitempty"
	Contact *Contact `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty" dynamodbav:"contact,omitempty"`
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
	// @gotags: dynamodbav:"pickup_location_with_dock,omitempty"
	PickupLocationWithDock bool `protobuf:"varint,14,opt,name=pickup_location_with_dock,json=pickupLocationWithDock,proto3" json:"pickup_location_with_dock,omitempty" dynamodbav:"pickup_location_with_dock,omitempty"`
	// @gotags: dynamodbav:"liftgate_pickup,omitempty"
	LiftgatePickup bool `protobuf:"varint,15,opt,name=liftgate_pickup,json=liftgatePickup,proto3" json:"liftgate_pickup,omitempty" dynamodbav:"liftgate_pickup,omitempty"`
	// @gotags: dynamodbav:"pickup_appointment,omitempty"
	PickupAppointment bool `protobuf:"varint,16,opt,name=pickup_appointment,json=pickupAppointment,proto3" json:"pickup_appointment,omitempty" dynamodbav:"pickup_appointment,omitempty"`
	// @gotags: dynamodbav:"inside_pickup,omitempty"
	InsidePickup bool `protobuf:"varint,17,opt,name=inside_pickup,json=insidePickup,proto3" json:"inside_pickup,omitempty" dynamodbav:"inside_pickup,omitempty"`
	// @gotags: dynamodbav:"pickup_notification,omitempty"
	PickupNotification bool `protobuf:"varint,18,opt,name=pickup_notification,json=pickupNotification,proto3" json:"pickup_notification,omitempty" dynamodbav:"pickup_notification,omitempty"`
	// @gotags: dynamodbav:"delivery_location_with_dock,omitempty"
	DeliveryLocationWithDock bool `protobuf:"varint,19,opt,name=delivery_location_with_dock,json=deliveryLocationWithDock,proto3" json:"delivery_location_with_dock,omitempty" dynamodbav:"delivery_location_with_dock,omitempty"`
	// @gotags: dynamodbav:"liftgate_delivery,omitempty"
	LiftgateDelivery bool `protobuf:"varint,20,opt,name=liftgate_delivery,json=liftgateDelivery,proto3" json:"liftgate_delivery,omitempty" dynamodbav:"liftgate_delivery,omitempty"`
	// @gotags: dynamodbav:"delivery_appointment,omitempty"
	DeliveryAppointment bool `protobuf:"varint,21,opt,name=delivery_appointment,json=deliveryAppointment,proto3" json:"delivery_appointment,omitempty" dynamodbav:"delivery_appointment,omitempty"`
	// @gotags: dynamodbav:"inside_delivery,omitempty"
	InsideDelivery bool `protobuf:"varint,22,opt,name=inside_delivery,json=insideDelivery,proto3" json:"inside_delivery,omitempty" dynamodbav:"inside_delivery,omitempty"`
	// @gotags: dynamodbav:"delivery_notification,omitempty"
	DeliveryNotification bool `protobuf:"varint,23,opt,name=delivery_notification,json=deliveryNotification,proto3" json:"delivery_notification,omitempty" dynamodbav:"delivery_notification,omitempty"`
	// @gotags: dynamodbav:"pickup_ready_by,omitempty"
	PickupReadyBy string `protobuf:"bytes,24,opt,name=pickup_ready_by,json=pickupReadyBy,proto3" json:"pickup_ready_by,omitempty" dynamodbav:"pickup_ready_by,omitempty"`
	// @gotags: dynamodbav:"pickup_instructions,omitempty"
	PickupInstructions string `protobuf:"bytes,25,opt,name=pickup_instructions,json=pickupInstructions,proto3" json:"pickup_instructions,omitempty" dynamodbav:"pickup_instructions,omitempty"`
	// @gotags: dynamodbav:"delivery_instructions,omitempty"
	DeliveryInstructions string `protobuf:"bytes,26,opt,name=delivery_instructions,json=deliveryInstructions,proto3" json:"delivery_instructions,omitempty" dynamodbav:"delivery_instructions,omitempty"`
	// @gotags: dynamodbav:"location_services,omitempty"
	LocationServices []LocationServices `protobuf:"varint,27,rep,packed,name=location_services,json=locationServices,proto3,enum=v1.LocationServices" json:"location_services,omitempty" dynamodbav:"location_services,omitempty"`
	// @gotags: dynamodbav:"default_pickup_addresss,omitempty"
	DefaultPickupAddresss *Address `protobuf:"bytes,28,opt,name=default_pickup_addresss,json=defaultPickupAddresss,proto3" json:"default_pickup_addresss,omitempty" dynamodbav:"default_pickup_addresss,omitempty"`
	// @gotags: dynamodbav:"default_delivery_address,omitempty"
	DefaultDeliveryAddress *Address `protobuf:"bytes,29,opt,name=default_delivery_address,json=defaultDeliveryAddress,proto3" json:"default_delivery_address,omitempty" dynamodbav:"default_delivery_address,omitempty"`
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

func (x *Location) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
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

func (x *Location) GetPickupLocationWithDock() bool {
	if x != nil {
		return x.PickupLocationWithDock
	}
	return false
}

func (x *Location) GetLiftgatePickup() bool {
	if x != nil {
		return x.LiftgatePickup
	}
	return false
}

func (x *Location) GetPickupAppointment() bool {
	if x != nil {
		return x.PickupAppointment
	}
	return false
}

func (x *Location) GetInsidePickup() bool {
	if x != nil {
		return x.InsidePickup
	}
	return false
}

func (x *Location) GetPickupNotification() bool {
	if x != nil {
		return x.PickupNotification
	}
	return false
}

func (x *Location) GetDeliveryLocationWithDock() bool {
	if x != nil {
		return x.DeliveryLocationWithDock
	}
	return false
}

func (x *Location) GetLiftgateDelivery() bool {
	if x != nil {
		return x.LiftgateDelivery
	}
	return false
}

func (x *Location) GetDeliveryAppointment() bool {
	if x != nil {
		return x.DeliveryAppointment
	}
	return false
}

func (x *Location) GetInsideDelivery() bool {
	if x != nil {
		return x.InsideDelivery
	}
	return false
}

func (x *Location) GetDeliveryNotification() bool {
	if x != nil {
		return x.DeliveryNotification
	}
	return false
}

func (x *Location) GetPickupReadyBy() string {
	if x != nil {
		return x.PickupReadyBy
	}
	return ""
}

func (x *Location) GetPickupInstructions() string {
	if x != nil {
		return x.PickupInstructions
	}
	return ""
}

func (x *Location) GetDeliveryInstructions() string {
	if x != nil {
		return x.DeliveryInstructions
	}
	return ""
}

func (x *Location) GetLocationServices() []LocationServices {
	if x != nil {
		return x.LocationServices
	}
	return nil
}

func (x *Location) GetDefaultPickupAddresss() *Address {
	if x != nil {
		return x.DefaultPickupAddresss
	}
	return nil
}

func (x *Location) GetDefaultDeliveryAddress() *Address {
	if x != nil {
		return x.DefaultDeliveryAddress
	}
	return nil
}

type Locations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"locations,omitempty"
	Locations []*Location `protobuf:"bytes,1,rep,name=Locations,proto3" json:"Locations,omitempty" dynamodbav:"locations,omitempty"`
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
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x09, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x68,
	0x6f, 0x75, 0x72, 0x73, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x69, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x55, 0x70,
	0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69,
	0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6b, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50,
	0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x6b,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x6b, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6b, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x6b, 0x12, 0x39, 0x0a, 0x19, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x63,
	0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x63, 0x6b, 0x12,
	0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x69, 0x66, 0x74, 0x67, 0x61,
	0x74, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x69, 0x64,
	0x65, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c,
	0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x2f, 0x0a, 0x13,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a,
	0x1b, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x6f, 0x63, 0x6b, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x18, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x11,
	0x6c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6c, 0x69, 0x66, 0x74, 0x67, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x14, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x69,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x62, 0x79, 0x18, 0x18, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x61, 0x64, 0x79,
	0x42, 0x79, 0x12, 0x2f, 0x0a, 0x13, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x1b, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x10, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x17,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x15, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x73, 0x12, 0x45, 0x0a, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x1d, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x37, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x61, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74,
	0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02,
	0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*Location)(nil),      // 0: v1.location
	(*Locations)(nil),     // 1: v1.locations
	(*Address)(nil),       // 2: v1.address
	(*Contact)(nil),       // 3: v1.contact
	(*BusinessHours)(nil), // 4: v1.business_hours
	(LocationServices)(0), // 5: v1.location_services
}
var file_location_proto_depIdxs = []int32{
	2, // 0: v1.location.address:type_name -> v1.address
	3, // 1: v1.location.contact:type_name -> v1.contact
	4, // 2: v1.location.business_hours:type_name -> v1.business_hours
	5, // 3: v1.location.location_services:type_name -> v1.location_services
	2, // 4: v1.location.default_pickup_addresss:type_name -> v1.address
	2, // 5: v1.location.default_delivery_address:type_name -> v1.address
	0, // 6: v1.locations.Locations:type_name -> v1.location
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_location_proto_init() }
func file_location_proto_init() {
	if File_location_proto != nil {
		return
	}
	file_address_proto_init()
	file_contact_proto_init()
	file_business_hours_proto_init()
	file_location_services_proto_init()
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: book.proto

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

type FreightChargeTerm int32

const (
	FreightChargeTerm_prepaid         FreightChargeTerm = 0
	FreightChargeTerm_third_party     FreightChargeTerm = 1
	FreightChargeTerm_inbound_collect FreightChargeTerm = 2
)

// Enum value maps for FreightChargeTerm.
var (
	FreightChargeTerm_name = map[int32]string{
		0: "prepaid",
		1: "third_party",
		2: "inbound_collect",
	}
	FreightChargeTerm_value = map[string]int32{
		"prepaid":         0,
		"third_party":     1,
		"inbound_collect": 2,
	}
)

func (x FreightChargeTerm) Enum() *FreightChargeTerm {
	p := new(FreightChargeTerm)
	*p = x
	return p
}

func (x FreightChargeTerm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FreightChargeTerm) Descriptor() protoreflect.EnumDescriptor {
	return file_book_proto_enumTypes[0].Descriptor()
}

func (FreightChargeTerm) Type() protoreflect.EnumType {
	return &file_book_proto_enumTypes[0]
}

func (x FreightChargeTerm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FreightChargeTerm.Descriptor instead.
func (FreightChargeTerm) EnumDescriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"capacity_provider_bolURL,omitempty"
	CapacityProviderBolUrl string `protobuf:"bytes,1,opt,name=capacity_provider_bol_url,json=capacityProviderBolUrl,proto3" json:"capacity_provider_bol_url,omitempty" dynamodbav:"capacity_provider_bolURL,omitempty"`
	// @gotags: dynamodbav:"shipment_identifier,omitempty"
	ShipmentIdentifier string `protobuf:"bytes,2,opt,name=shipment_identifier,json=shipmentIdentifier,proto3" json:"shipment_identifier,omitempty" dynamodbav:"shipment_identifier,omitempty"`
	// @gotags: dynamodbav:"pickup_note,omitempty"
	PickupNote string `protobuf:"bytes,3,opt,name=pickup_note,json=pickupNote,proto3" json:"pickup_note,omitempty" dynamodbav:"pickup_note,omitempty"`
	// @gotags: dynamodbav:"pickup_date_time,omitempty"
	PickupDateTime string `protobuf:"bytes,4,opt,name=pickup_date_time,json=pickupDateTime,proto3" json:"pickup_date_time,omitempty" dynamodbav:"pickup_date_time,omitempty"`
	// @gotags: dynamodbav:"errors,omitempty"
	Errors []string `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty" dynamodbav:"errors,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *Result) GetCapacityProviderBolUrl() string {
	if x != nil {
		return x.CapacityProviderBolUrl
	}
	return ""
}

func (x *Result) GetShipmentIdentifier() string {
	if x != nil {
		return x.ShipmentIdentifier
	}
	return ""
}

func (x *Result) GetPickupNote() string {
	if x != nil {
		return x.PickupNote
	}
	return ""
}

func (x *Result) GetPickupDateTime() string {
	if x != nil {
		return x.PickupDateTime
	}
	return ""
}

func (x *Result) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BolId              string            `protobuf:"bytes,2,opt,name=bol_id,json=bolId,proto3" json:"bol_id,omitempty"`
	BolUrl             string            `protobuf:"bytes,3,opt,name=bol_url,json=bolUrl,proto3" json:"bol_url,omitempty"`
	PickupDate         string            `protobuf:"bytes,4,opt,name=pickup_date,json=pickupDate,proto3" json:"pickup_date,omitempty"`
	DiscountPercentage float64           `protobuf:"fixed64,5,opt,name=discount_percentage,json=discountPercentage,proto3" json:"discount_percentage,omitempty"`
	FreightChargeTerm  FreightChargeTerm `protobuf:"varint,8,opt,name=freight_charge_term,json=freightChargeTerm,proto3,enum=v1.FreightChargeTerm" json:"freight_charge_term,omitempty"`
	Quote              *Quote            `protobuf:"bytes,9,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *Booking) GetBolId() string {
	if x != nil {
		return x.BolId
	}
	return ""
}

func (x *Booking) GetBolUrl() string {
	if x != nil {
		return x.BolUrl
	}
	return ""
}

func (x *Booking) GetPickupDate() string {
	if x != nil {
		return x.PickupDate
	}
	return ""
}

func (x *Booking) GetDiscountPercentage() float64 {
	if x != nil {
		return x.DiscountPercentage
	}
	return 0
}

func (x *Booking) GetFreightChargeTerm() FreightChargeTerm {
	if x != nil {
		return x.FreightChargeTerm
	}
	return FreightChargeTerm_prepaid
}

func (x *Booking) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bid,omitempty"
	Bid *Bid `protobuf:"bytes,1,opt,name=bid,proto3" json:"bid,omitempty" dynamodbav:"bid,omitempty"`
	// @gotags: dynamodbav:"quote,omitempty"
	Quote *Quote `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty" dynamodbav:"quote,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *BookingRequest) GetBid() *Bid {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *BookingRequest) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

type Bookings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bookings,omitempty"
	Bookings []*Booking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty" dynamodbav:"bookings,omitempty"`
}

func (x *Bookings) Reset() {
	*x = Bookings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookings) ProtoMessage() {}

func (x *Bookings) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookings.ProtoReflect.Descriptor instead.
func (*Bookings) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *Bookings) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x72, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x62, 0x69, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2f, 0x0a, 0x13,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0xf5, 0x01, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x0a, 0x06,
	0x62, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x12, 0x47,
	0x0a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x52, 0x11, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68, 0x61,
	0x72, 0x67, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1f, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x62,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69,
	0x64, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x52, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x33, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x2a, 0x48, 0x0a, 0x13,
	0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x10, 0x02, 0x42, 0x5d, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x42, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2,
	0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_book_proto_goTypes = []interface{}{
	(FreightChargeTerm)(0), // 0: v1.freight_charge_term
	(*Result)(nil),         // 1: v1.Result
	(*Booking)(nil),        // 2: v1.booking
	(*BookingRequest)(nil), // 3: v1.booking_request
	(*Bookings)(nil),       // 4: v1.bookings
	(*Quote)(nil),          // 5: v1.quote
	(*Bid)(nil),            // 6: v1.bid
}
var file_book_proto_depIdxs = []int32{
	0, // 0: v1.booking.freight_charge_term:type_name -> v1.freight_charge_term
	5, // 1: v1.booking.quote:type_name -> v1.quote
	6, // 2: v1.booking_request.bid:type_name -> v1.bid
	5, // 3: v1.booking_request.quote:type_name -> v1.quote
	2, // 4: v1.bookings.bookings:type_name -> v1.booking
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	file_role_proto_init()
	file_amount_proto_init()
	file_ref_proto_init()
	file_quote_proto_init()
	file_bid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookings); i {
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
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		EnumInfos:         file_book_proto_enumTypes,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}

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

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"capacity_provider_bolURL,omitempty"
	CapacityProviderBolURL string `protobuf:"bytes,1,opt,name=capacity_provider_bolURL,json=capacityProviderBolURL,proto3" json:"capacity_provider_bolURL,omitempty" dynamodbav:"capacity_provider_bolURL,omitempty"`
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

func (x *Result) GetCapacityProviderBolURL() string {
	if x != nil {
		return x.CapacityProviderBolURL
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

	// @gotags: dynamodbav:"booking_id,omitempty"
	BookingId string `protobuf:"bytes,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty" dynamodbav:"booking_id,omitempty"`
	// @gotags: dynamodbav:"bol_url,omitempty"
	BolUrl string `protobuf:"bytes,2,opt,name=bol_url,json=bolUrl,proto3" json:"bol_url,omitempty" dynamodbav:"bol_url,omitempty"`
	// @gotags: dynamodbav:"invoice_url,omitempty"
	InvoiceUrl string `protobuf:"bytes,3,opt,name=invoice_url,json=invoiceUrl,proto3" json:"invoice_url,omitempty" dynamodbav:"invoice_url,omitempty"`
	// @gotags: dynamodbav:"invoice_due_date,omitempty"
	InvoiceDueDate string `protobuf:"bytes,4,opt,name=invoice_due_date,json=invoiceDueDate,proto3" json:"invoice_due_date,omitempty" dynamodbav:"invoice_due_date,omitempty"`
	// @gotags: dynamodbav:"pick_up_start,omitempty"
	PickUpStart string `protobuf:"bytes,5,opt,name=pick_up_start,json=pickUpStart,proto3" json:"pick_up_start,omitempty" dynamodbav:"pick_up_start,omitempty"`
	// @gotags: dynamodbav:"pick_up_end,omitempty"
	PickUpEnd string `protobuf:"bytes,6,opt,name=pick_up_end,json=pickUpEnd,proto3" json:"pick_up_end,omitempty" dynamodbav:"pick_up_end,omitempty"`
	// @gotags: dynamodbav:"booked_quote,omitempty"
	Quote *Quote `protobuf:"bytes,7,opt,name=quote,json=quote,omitempty,proto3" json:"quote,omitempty" dynamodbav:"booked_quote,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,8,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
	// @gotags: dynamodbav:"bid,omitempty"
	Bid *Bid `protobuf:"bytes,9,opt,name=bid,proto3" json:"bid,omitempty" dynamodbav:"bid,omitempty"`
	// @gotags: dynamodbav:"shipment_id,omitempty"
	ShipmentId string `protobuf:"bytes,10,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id,omitempty" dynamodbav:"shipment_id,omitempty"`
	// @gotags: dynamodbav:"secuirity_key,omitempty"
	SecuirityKey string `protobuf:"bytes,11,opt,name=secuirity_key,json=secuirityKey,proto3" json:"secuirity_key,omitempty" dynamodbav:"secuirity_key,omitempty"`
	// @gotags: dynamodbav:"pickup_number,omitempty"
	PickupNumber string `protobuf:"bytes,12,opt,name=pickup_number,json=pickupNumber,proto3" json:"pickup_number,omitempty" dynamodbav:"pickup_number,omitempty"`
	// @gotags: dynamodbav:"carrier_name,omitempty"
	CarrierName string `protobuf:"bytes,13,opt,name=carrier_name,json=carrierName,proto3" json:"carrier_name,omitempty" dynamodbav:"carrier_name,omitempty"`
	// @gotags: dynamodbav:"carrier_pro_number,omitempty"
	CarrierProNumber string `protobuf:"bytes,14,opt,name=carrier_pro_number,json=carrierProNumber,proto3" json:"carrier_pro_number,omitempty" dynamodbav:"carrier_pro_number,omitempty"`
	// @gotags: dynamodbav:"handeling_unit_total,omitempty"
	HandelingUnitTotal string `protobuf:"bytes,15,opt,name=handeling_unit_total,json=handelingUnitTotal,proto3" json:"handeling_unit_total,omitempty" dynamodbav:"handeling_unit_total,omitempty"`
	// @gotags: dynamodbav:"is_shipment_edit,omitempty"
	IsShipmentEdit bool `protobuf:"varint,16,opt,name=is_shipment_edit,json=isShipmentEdit,proto3" json:"is_shipment_edit,omitempty" dynamodbav:"is_shipment_edit,omitempty"`
	// @gotags: dynamodbav:"is_shipment_manual,omitempty"
	IsShipmentManual bool `protobuf:"varint,17,opt,name=is_shipment_manual,json=isShipmentManual,proto3" json:"is_shipment_manual,omitempty" dynamodbav:"is_shipment_manual,omitempty"`
	// @gotags: dynamodbav:"service_type,omitempty"
	ServiceType bool `protobuf:"varint,18,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty" dynamodbav:"service_type,omitempty"`
	// @gotags: dynamodbav:"is_tracking_email_send,omitempty"
	IsTrackingEmailSend bool `protobuf:"varint,19,opt,name=is_tracking_email_send,json=isTrackingEmailSend,proto3" json:"is_tracking_email_send,omitempty" dynamodbav:"is_tracking_email_send,omitempty"`
	// @gotags: dynamodbav:"is_tracking_API_enabled,omitempty"
	IsTracking_APIEnabled bool `protobuf:"varint,20,opt,name=is_tracking_API_enabled,json=isTrackingAPIEnabled,proto3" json:"is_tracking_API_enabled,omitempty" dynamodbav:"is_tracking_API_enabled,omitempty"`
	// @gotags: dynamodbav:"customer_BOL_number,omitempty"
	Customer_BOLNumber string `protobuf:"bytes,21,opt,name=customer_BOL_number,json=customerBOLNumber,proto3" json:"customer_BOL_number,omitempty" dynamodbav:"customer_BOL_number,omitempty"`
	// @gotags: dynamodbav:"shipper_email,omitempty"
	ShipperEmail string `protobuf:"bytes,22,opt,name=shipper_email,json=shipperEmail,proto3" json:"shipper_email,omitempty" dynamodbav:"shipper_email,omitempty"`
	// @gotags: dynamodbav:"consignee_email,omitempty"
	ConsigneeEmail string `protobuf:"bytes,23,opt,name=consignee_email,json=consigneeEmail,proto3" json:"consignee_email,omitempty" dynamodbav:"consignee_email,omitempty"`
	// @gotags: dynamodbav:"result,omitempty"
	Result *Result `protobuf:"bytes,24,opt,name=result,proto3" json:"result,omitempty" dynamodbav:"result,omitempty"`
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

func (x *Booking) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *Booking) GetBolUrl() string {
	if x != nil {
		return x.BolUrl
	}
	return ""
}

func (x *Booking) GetInvoiceUrl() string {
	if x != nil {
		return x.InvoiceUrl
	}
	return ""
}

func (x *Booking) GetInvoiceDueDate() string {
	if x != nil {
		return x.InvoiceDueDate
	}
	return ""
}

func (x *Booking) GetPickUpStart() string {
	if x != nil {
		return x.PickUpStart
	}
	return ""
}

func (x *Booking) GetPickUpEnd() string {
	if x != nil {
		return x.PickUpEnd
	}
	return ""
}

func (x *Booking) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *Booking) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Booking) GetBid() *Bid {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *Booking) GetShipmentId() string {
	if x != nil {
		return x.ShipmentId
	}
	return ""
}

func (x *Booking) GetSecuirityKey() string {
	if x != nil {
		return x.SecuirityKey
	}
	return ""
}

func (x *Booking) GetPickupNumber() string {
	if x != nil {
		return x.PickupNumber
	}
	return ""
}

func (x *Booking) GetCarrierName() string {
	if x != nil {
		return x.CarrierName
	}
	return ""
}

func (x *Booking) GetCarrierProNumber() string {
	if x != nil {
		return x.CarrierProNumber
	}
	return ""
}

func (x *Booking) GetHandelingUnitTotal() string {
	if x != nil {
		return x.HandelingUnitTotal
	}
	return ""
}

func (x *Booking) GetIsShipmentEdit() bool {
	if x != nil {
		return x.IsShipmentEdit
	}
	return false
}

func (x *Booking) GetIsShipmentManual() bool {
	if x != nil {
		return x.IsShipmentManual
	}
	return false
}

func (x *Booking) GetServiceType() bool {
	if x != nil {
		return x.ServiceType
	}
	return false
}

func (x *Booking) GetIsTrackingEmailSend() bool {
	if x != nil {
		return x.IsTrackingEmailSend
	}
	return false
}

func (x *Booking) GetIsTracking_APIEnabled() bool {
	if x != nil {
		return x.IsTracking_APIEnabled
	}
	return false
}

func (x *Booking) GetCustomer_BOLNumber() string {
	if x != nil {
		return x.Customer_BOLNumber
	}
	return ""
}

func (x *Booking) GetShipperEmail() string {
	if x != nil {
		return x.ShipperEmail
	}
	return ""
}

func (x *Booking) GetConsigneeEmail() string {
	if x != nil {
		return x.ConsigneeEmail
	}
	return ""
}

func (x *Booking) GetResult() *Result {
	if x != nil {
		return x.Result
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
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookings) ProtoMessage() {}

func (x *Bookings) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Bookings.ProtoReflect.Descriptor instead.
func (*Bookings) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
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
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x38, 0x0a, 0x18, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x2f, 0x0a, 0x13, 0x73,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x22,
	0xba, 0x07, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f,
	0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f,
	0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x65, 0x6e,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x45,
	0x6e, 0x64, 0x12, 0x35, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x0a, 0xba, 0xe9,
	0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2c,
	0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x62, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x64,
	0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x69, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x63, 0x75, 0x69, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x30, 0x0a, 0x14, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x75,
	0x6e, 0x69, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x68, 0x61, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69,
	0x73, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x69, 0x73, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33,
	0x0a, 0x16, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13,
	0x69, 0x73, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x69, 0x73, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x41, 0x50, 0x49, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x41, 0x50, 0x49, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x42, 0x4f, 0x4c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x42, 0x4f, 0x4c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x73, 0x42, 0x5d, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x6f, 0x6f,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73,
	0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58,
	0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_book_proto_goTypes = []interface{}{
	(*Result)(nil),   // 0: v1.Result
	(*Booking)(nil),  // 1: v1.booking
	(*Bookings)(nil), // 2: v1.bookings
	(*Quote)(nil),    // 3: v1.quote
	(*Bid)(nil),      // 4: v1.bid
}
var file_book_proto_depIdxs = []int32{
	3, // 0: v1.booking.quote:type_name -> v1.quote
	4, // 1: v1.booking.bid:type_name -> v1.bid
	0, // 2: v1.booking.result:type_name -> v1.Result
	1, // 3: v1.bookings.bookings:type_name -> v1.booking
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	file_role_proto_init()
	file_amount_proto_init()
	file_quote_proto_init()
	file_ref_proto_init()
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
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}

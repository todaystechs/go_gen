// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: book.proto

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
	Quote *Quote `protobuf:"bytes,7,opt,name=quote,proto3" json:"quote,omitempty" dynamodbav:"booked_quote,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,8,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
	// @gotags: dynamodbav:"bid_data,omitempty"
	Bid *Bid `protobuf:"bytes,9,opt,name=bid,proto3" json:"bid,omitempty" dynamodbav:"bid_data,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
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
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookings) ProtoMessage() {}

func (x *Bookings) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Bookings.ProtoReflect.Descriptor instead.
func (*Bookings) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
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
	0x74, 0x6f, 0x1a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x02,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6c, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6c, 0x55, 0x72,
	0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1e, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x45, 0x6e, 0x64,
	0x12, 0x1f, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x03, 0x62, 0x69, 0x64, 0x22, 0x33, 0x0a,
	0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x27, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_book_proto_goTypes = []interface{}{
	(*Booking)(nil),  // 0: v1.booking
	(*Bookings)(nil), // 1: v1.bookings
	(*Quote)(nil),    // 2: v1.quote
	(*Bid)(nil),      // 3: v1.bid
}
var file_book_proto_depIdxs = []int32{
	2, // 0: v1.booking.quote:type_name -> v1.quote
	3, // 1: v1.booking.bid:type_name -> v1.bid
	0, // 2: v1.bookings.bookings:type_name -> v1.booking
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   2,
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
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

type BookingData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
}

func (x *BookingData) Reset() {
	*x = BookingData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingData) ProtoMessage() {}

func (x *BookingData) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BookingData.ProtoReflect.Descriptor instead.
func (*BookingData) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookingData) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *BookingData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type BookingResponse struct {
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
	BookedQuote *QuoteEntity `protobuf:"bytes,7,opt,name=booked_quote,json=bookedQuote,proto3" json:"booked_quote,omitempty" dynamodbav:"booked_quote,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookingResponse) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *BookingResponse) GetBolUrl() string {
	if x != nil {
		return x.BolUrl
	}
	return ""
}

func (x *BookingResponse) GetInvoiceUrl() string {
	if x != nil {
		return x.InvoiceUrl
	}
	return ""
}

func (x *BookingResponse) GetInvoiceDueDate() string {
	if x != nil {
		return x.InvoiceDueDate
	}
	return ""
}

func (x *BookingResponse) GetPickUpStart() string {
	if x != nil {
		return x.PickUpStart
	}
	return ""
}

func (x *BookingResponse) GetPickUpEnd() string {
	if x != nil {
		return x.PickUpEnd
	}
	return ""
}

func (x *BookingResponse) GetBookedQuote() *QuoteEntity {
	if x != nil {
		return x.BookedQuote
	}
	return nil
}

type ListOfBooking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"booking_response,omitempty"
	BookingHistory []*BookingResponse `protobuf:"bytes,1,rep,name=booking_history,json=bookingHistory,proto3" json:"booking_history,omitempty" dynamodbav:"booking_response,omitempty"`
}

func (x *ListOfBooking) Reset() {
	*x = ListOfBooking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfBooking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfBooking) ProtoMessage() {}

func (x *ListOfBooking) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ListOfBooking.ProtoReflect.Descriptor instead.
func (*ListOfBooking) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfBooking) GetBookingHistory() []*BookingResponse {
	if x != nil {
		return x.BookingHistory
	}
	return nil
}

type FetchBookingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"booking_id,omitempty"
	BookingId []string `protobuf:"bytes,1,rep,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty" dynamodbav:"booking_id,omitempty"`
	// @gotags: dynamodbav:"pk,omitempty"
	Pk string `protobuf:"bytes,2,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk,omitempty"`
	// @gotags: dynamodbav:"start_from,omitempty"
	StartFrom string `protobuf:"bytes,3,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty" dynamodbav:"start_from,omitempty"`
	// @gotags: dynamodbav:"end_on,omitempty"
	EndOn string `protobuf:"bytes,4,opt,name=end_on,json=endOn,proto3" json:"end_on,omitempty" dynamodbav:"end_on,omitempty"`
}

func (x *FetchBookingsRequest) Reset() {
	*x = FetchBookingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBookingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBookingsRequest) ProtoMessage() {}

func (x *FetchBookingsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FetchBookingsRequest.ProtoReflect.Descriptor instead.
func (*FetchBookingsRequest) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *FetchBookingsRequest) GetBookingId() []string {
	if x != nil {
		return x.BookingId
	}
	return nil
}

func (x *FetchBookingsRequest) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *FetchBookingsRequest) GetStartFrom() string {
	if x != nil {
		return x.StartFrom
	}
	return ""
}

func (x *FetchBookingsRequest) GetEndOn() string {
	if x != nil {
		return x.EndOn
	}
	return ""
}

type BookEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_entity,omitempty"
	QuoteEntity *QuoteEntity `protobuf:"bytes,1,opt,name=quote_entity,json=quoteEntity,proto3" json:"quote_entity,omitempty" dynamodbav:"quote_entity,omitempty"`
	// @gotags: dynamodbav:"ref,omitempty"
	Ref *Refereses `protobuf:"bytes,2,opt,name=ref,proto3" json:"ref,omitempty" dynamodbav:"ref,omitempty"`
}

func (x *BookEntity) Reset() {
	*x = BookEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookEntity) ProtoMessage() {}

func (x *BookEntity) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookEntity.ProtoReflect.Descriptor instead.
func (*BookEntity) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{4}
}

func (x *BookEntity) GetQuoteEntity() *QuoteEntity {
	if x != nil {
		return x.QuoteEntity
	}
	return nil
}

func (x *BookEntity) GetRef() *Refereses {
	if x != nil {
		return x.Ref
	}
	return nil
}

type DynamoBookEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"pk,omitempty"
	Pk string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk,omitempty"`
	// @gotags: dynamodbav:"sk,omitempty"
	Sk string `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk,omitempty"`
	// @gotags: dynamodbav:"quote_pk,omitempty"
	BookPk string `protobuf:"bytes,3,opt,name=book_pk,json=bookPk,proto3" json:"book_pk,omitempty" dynamodbav:"quote_pk,omitempty"`
	// @gotags: dynamodbav:"quote_sk,omitempty"
	BookSk string `protobuf:"bytes,4,opt,name=book_sk,json=bookSk,proto3" json:"book_sk,omitempty" dynamodbav:"quote_sk,omitempty"`
	// @gotags: dynamodbav:"book_entity,omitempty"
	BookEntity *BookEntity `protobuf:"bytes,5,opt,name=book_entity,json=bookEntity,proto3" json:"book_entity,omitempty" dynamodbav:"book_entity,omitempty"`
}

func (x *DynamoBookEntity) Reset() {
	*x = DynamoBookEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamoBookEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamoBookEntity) ProtoMessage() {}

func (x *DynamoBookEntity) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamoBookEntity.ProtoReflect.Descriptor instead.
func (*DynamoBookEntity) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{5}
}

func (x *DynamoBookEntity) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *DynamoBookEntity) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *DynamoBookEntity) GetBookPk() string {
	if x != nil {
		return x.BookPk
	}
	return ""
}

func (x *DynamoBookEntity) GetBookSk() string {
	if x != nil {
		return x.BookSk
	}
	return ""
}

func (x *DynamoBookEntity) GetBookEntity() *BookEntity {
	if x != nil {
		return x.BookEntity
	}
	return nil
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x72, 0x65, 0x66, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8e, 0x02, 0x0a, 0x0f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6c, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6c, 0x55, 0x72, 0x6c,
	0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x75, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70,
	0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x1e, 0x0a, 0x0b, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x75, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x55, 0x70, 0x45, 0x6e, 0x64, 0x12,
	0x34, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x65, 0x64,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x7b, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x70, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x4f, 0x6e, 0x22, 0x65, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x34, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x73, 0x65, 0x73, 0x52, 0x03, 0x72, 0x65, 0x66, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x44,
	0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x42, 0x6f, 0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6b, 0x12,
	0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x50, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x6b,
	0x5f, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6f, 0x6b, 0x53,
	0x6b, 0x12, 0x31, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x42, 0x67, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_book_proto_goTypes = []interface{}{
	(*BookingData)(nil),          // 0: user.BookingData
	(*BookingResponse)(nil),      // 1: user.BookingResponse
	(*ListOfBooking)(nil),        // 2: user.ListOfBooking
	(*FetchBookingsRequest)(nil), // 3: user.FetchBookingsRequest
	(*BookEntity)(nil),           // 4: user.BookEntity
	(*DynamoBookEntity)(nil),     // 5: user.DynamoBookEntity
	(*QuoteEntity)(nil),          // 6: user.QuoteEntity
	(*Refereses)(nil),            // 7: user.refereses
}
var file_book_proto_depIdxs = []int32{
	6, // 0: user.BookingResponse.booked_quote:type_name -> user.QuoteEntity
	1, // 1: user.ListOfBooking.booking_history:type_name -> user.BookingResponse
	6, // 2: user.BookEntity.quote_entity:type_name -> user.QuoteEntity
	7, // 3: user.BookEntity.ref:type_name -> user.refereses
	4, // 4: user.DynamoBookEntity.book_entity:type_name -> user.BookEntity
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
	file_quote_proto_init()
	file_ref_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingData); i {
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
			switch v := v.(*BookingResponse); i {
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
			switch v := v.(*ListOfBooking); i {
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
			switch v := v.(*FetchBookingsRequest); i {
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
		file_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookEntity); i {
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
		file_book_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamoBookEntity); i {
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
			NumMessages:   6,
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

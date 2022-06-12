// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: quote.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty"`
	// @gotags: dynamodbav:"requester_id,omitempty"
	RequesterId string `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	// @gotags: dynamodbav:"business_pk,omitempty"
	BusinessPk string `protobuf:"bytes,3,opt,name=business_pk,json=businessPk,proto3" json:"business_pk,omitempty"`
	// @gotags: dynamodbav:"business_sk,omitempty"
	BusinessSk string `protobuf:"bytes,4,opt,name=business_sk,json=businessSk,proto3" json:"business_sk,omitempty"`
	// @gotags: dynamodbav:"mode,omitempty"
	Mode string `protobuf:"bytes,5,opt,name=mode,proto3" json:"mode,omitempty"`
	// @gotags: dynamodbav:"liable_party_id,omitempty"
	LiablePartyId string `protobuf:"bytes,6,opt,name=liable_party_id,json=liablePartyId,proto3" json:"liable_party_id,omitempty"`
	// @gotags: dynamodbav:"pickup,omitempty"
	Pickup *Location `protobuf:"bytes,7,opt,name=pickup,proto3" json:"pickup,omitempty"`
	// @gotags: dynamodbav:"delivery,omitempty"
	Delivery *Location `protobuf:"bytes,8,opt,name=delivery,proto3" json:"delivery,omitempty"`
	// @gotags: dynamodbav:"commodities,omitempty"
	Commodities []*Commodity `protobuf:"bytes,9,rep,name=commodities,proto3" json:"commodities,omitempty"`
	// @gotags: dynamodbav:"pickup_date,omitempty"
	PickupDate string `protobuf:"bytes,10,opt,name=pickup_date,json=pickupDate,proto3" json:"pickup_date,omitempty"`
	// @gotags: dynamodbav:"display_date,omitempty"
	DisplayDate string `protobuf:"bytes,11,opt,name=display_date,json=displayDate,proto3" json:"display_date,omitempty"`
	// @gotags: dynamodbav:"delivery_date,omitempty"
	DeliveryDate string `protobuf:"bytes,12,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	// @gotags: dynamodbav:"bids,omitempty"
	Bids []*Bid `protobuf:"bytes,13,rep,name=bids,proto3" json:"bids,omitempty"`
	// @gotags: dynamodbav:"vendor_bids,omitempty"
	VendorBids map[string]*anypb.Any `protobuf:"bytes,14,rep,name=vendor_bids,json=vendorBids,proto3" json:"vendor_bids,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{0}
}

func (x *Quote) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *Quote) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *Quote) GetBusinessPk() string {
	if x != nil {
		return x.BusinessPk
	}
	return ""
}

func (x *Quote) GetBusinessSk() string {
	if x != nil {
		return x.BusinessSk
	}
	return ""
}

func (x *Quote) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Quote) GetLiablePartyId() string {
	if x != nil {
		return x.LiablePartyId
	}
	return ""
}

func (x *Quote) GetPickup() *Location {
	if x != nil {
		return x.Pickup
	}
	return nil
}

func (x *Quote) GetDelivery() *Location {
	if x != nil {
		return x.Delivery
	}
	return nil
}

func (x *Quote) GetCommodities() []*Commodity {
	if x != nil {
		return x.Commodities
	}
	return nil
}

func (x *Quote) GetPickupDate() string {
	if x != nil {
		return x.PickupDate
	}
	return ""
}

func (x *Quote) GetDisplayDate() string {
	if x != nil {
		return x.DisplayDate
	}
	return ""
}

func (x *Quote) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *Quote) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

func (x *Quote) GetVendorBids() map[string]*anypb.Any {
	if x != nil {
		return x.VendorBids
	}
	return nil
}

type Quotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quotes,omitempty"
	Quotes []*Quote `protobuf:"bytes,1,rep,name=quotes,proto3" json:"quotes,omitempty"`
}

func (x *Quotes) Reset() {
	*x = Quotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quotes) ProtoMessage() {}

func (x *Quotes) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quotes.ProtoReflect.Descriptor instead.
func (*Quotes) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{1}
}

func (x *Quotes) GetQuotes() []*Quote {
	if x != nil {
		return x.Quotes
	}
	return nil
}

type QuoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote,omitempty"
	Quote *Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	// @gotags: dynamodbav:"bids,omitempty"
	Bids []*Bid `protobuf:"bytes,2,rep,name=bids,proto3" json:"bids,omitempty"`
}

func (x *QuoteResponse) Reset() {
	*x = QuoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteResponse) ProtoMessage() {}

func (x *QuoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteResponse.ProtoReflect.Descriptor instead.
func (*QuoteResponse) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{2}
}

func (x *QuoteResponse) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *QuoteResponse) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_quote_proto protoreflect.FileDescriptor

var file_quote_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76,
	0x31, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x04, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x6b, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x70,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x12, 0x28, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x2f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x04, 0x62, 0x69,
	0x64, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x62, 0x69, 0x64,
	0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x69, 0x64, 0x73, 0x1a, 0x53,
	0x0a, 0x0f, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x69, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x21, 0x0a,
	0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x06, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x76, 0x31, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_quote_proto_rawDescOnce sync.Once
	file_quote_proto_rawDescData = file_quote_proto_rawDesc
)

func file_quote_proto_rawDescGZIP() []byte {
	file_quote_proto_rawDescOnce.Do(func() {
		file_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_quote_proto_rawDescData)
	})
	return file_quote_proto_rawDescData
}

var file_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_quote_proto_goTypes = []interface{}{
	(*Quote)(nil),         // 0: v1.quote
	(*Quotes)(nil),        // 1: v1.quotes
	(*QuoteResponse)(nil), // 2: v1.quote_response
	nil,                   // 3: v1.quote.VendorBidsEntry
	(*Location)(nil),      // 4: v1.location
	(*Commodity)(nil),     // 5: v1.commodity
	(*Bid)(nil),           // 6: v1.bid
	(*anypb.Any)(nil),     // 7: google.protobuf.Any
}
var file_quote_proto_depIdxs = []int32{
	4, // 0: v1.quote.pickup:type_name -> v1.location
	4, // 1: v1.quote.delivery:type_name -> v1.location
	5, // 2: v1.quote.commodities:type_name -> v1.commodity
	6, // 3: v1.quote.bids:type_name -> v1.bid
	3, // 4: v1.quote.vendor_bids:type_name -> v1.quote.VendorBidsEntry
	0, // 5: v1.quotes.quotes:type_name -> v1.quote
	0, // 6: v1.quote_response.quote:type_name -> v1.quote
	6, // 7: v1.quote_response.bids:type_name -> v1.bid
	7, // 8: v1.quote.VendorBidsEntry.value:type_name -> google.protobuf.Any
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_quote_proto_init() }
func file_quote_proto_init() {
	if File_quote_proto != nil {
		return
	}
	file_commodity_proto_init()
	file_carrier_proto_init()
	file_amount_proto_init()
	file_bid_proto_init()
	file_location_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
		file_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quotes); i {
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
		file_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteResponse); i {
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
			RawDescriptor: file_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_quote_proto_goTypes,
		DependencyIndexes: file_quote_proto_depIdxs,
		MessageInfos:      file_quote_proto_msgTypes,
	}.Build()
	File_quote_proto = out.File
	file_quote_proto_rawDesc = nil
	file_quote_proto_goTypes = nil
	file_quote_proto_depIdxs = nil
}

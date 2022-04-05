// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: quote.proto

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

// @gotags: dynamodbav:"quote_entity,omitempty" json:"quote_entity,omitempty""
type QuoteEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
	// @gotags: dynamodbav:"requester_id,omitempty"
	RequesterId string `protobuf:"bytes,2,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty" dynamodbav:"requester_id,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,3,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
	// @gotags: dynamodbav:"mode,omitempty"
	Mode string `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty" dynamodbav:"mode,omitempty"`
	// @gotags: dynamodbav:"liable_party_id,omitempty"
	LiablePartyId string `protobuf:"bytes,5,opt,name=liable_party_id,json=liablePartyId,proto3" json:"liable_party_id,omitempty" dynamodbav:"liable_party_id,omitempty"`
	// @gotags: dynamodbav:"shipping_detail,omitempty"
	ShippingDetail *ShippingDetail `protobuf:"bytes,6,opt,name=shipping_detail,json=shippingDetail,proto3" json:"shipping_detail,omitempty" dynamodbav:"shipping_detail,omitempty"`
	// @gotags: dynamodbav:"additional_services,omitempty"
	AdditionalServices []AdditionalService `protobuf:"varint,7,rep,packed,name=additional_services,json=additionalServices,proto3,enum=user.AdditionalService" json:"additional_services,omitempty" dynamodbav:"additional_services,omitempty"`
	// @gotags: dynamodbav:"commodities,omitempty"
	Commodities []*Commodity `protobuf:"bytes,8,rep,name=commodities,proto3" json:"commodities,omitempty" dynamodbav:"commodities,omitempty"`
}

func (x *QuoteEntity) Reset() {
	*x = QuoteEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteEntity) ProtoMessage() {}

func (x *QuoteEntity) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QuoteEntity.ProtoReflect.Descriptor instead.
func (*QuoteEntity) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{0}
}

func (x *QuoteEntity) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *QuoteEntity) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *QuoteEntity) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *QuoteEntity) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *QuoteEntity) GetLiablePartyId() string {
	if x != nil {
		return x.LiablePartyId
	}
	return ""
}

func (x *QuoteEntity) GetShippingDetail() *ShippingDetail {
	if x != nil {
		return x.ShippingDetail
	}
	return nil
}

func (x *QuoteEntity) GetAdditionalServices() []AdditionalService {
	if x != nil {
		return x.AdditionalServices
	}
	return nil
}

func (x *QuoteEntity) GetCommodities() []*Commodity {
	if x != nil {
		return x.Commodities
	}
	return nil
}

type DynamoQuoteEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"pk,omitempty"
	Pk string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk,omitempty"`
	// @gotags: dynamodbav:"sk,omitempty"
	Sk string `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk,omitempty"`
	// @gotags: dynamodbav:"quote_pk,omitempty"
	QuotePk string `protobuf:"bytes,3,opt,name=quote_pk,json=quotePk,proto3" json:"quote_pk,omitempty" dynamodbav:"quote_pk,omitempty"`
	// @gotags: dynamodbav:"quote_sk,omitempty"
	QuoteSk string `protobuf:"bytes,4,opt,name=quote_sk,json=quoteSk,proto3" json:"quote_sk,omitempty" dynamodbav:"quote_sk,omitempty"`
	// @gotags: dynamodbav:"quote_entity,omitempty"
	QuoteEntity *QuoteEntity `protobuf:"bytes,5,opt,name=quote_entity,json=quoteEntity,proto3" json:"quote_entity,omitempty" dynamodbav:"quote_entity,omitempty"`
}

func (x *DynamoQuoteEntity) Reset() {
	*x = DynamoQuoteEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamoQuoteEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamoQuoteEntity) ProtoMessage() {}

func (x *DynamoQuoteEntity) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DynamoQuoteEntity.ProtoReflect.Descriptor instead.
func (*DynamoQuoteEntity) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{1}
}

func (x *DynamoQuoteEntity) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *DynamoQuoteEntity) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *DynamoQuoteEntity) GetQuotePk() string {
	if x != nil {
		return x.QuotePk
	}
	return ""
}

func (x *DynamoQuoteEntity) GetQuoteSk() string {
	if x != nil {
		return x.QuoteSk
	}
	return ""
}

func (x *DynamoQuoteEntity) GetQuoteEntity() *QuoteEntity {
	if x != nil {
		return x.QuoteEntity
	}
	return nil
}

// @gotags: dynamodbav:"quote_bid_entity,omitempty" json:"quote_bid_entity,omitempty"
type BidEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
	// @gotags: dynamodbav:"sort_id,omitempty"
	Sk string `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sort_id,omitempty"`
	// @gotags: dynamodbav:"carrier,omitempty"
	Carrier *Carrier `protobuf:"bytes,3,opt,name=carrier,proto3" json:"carrier,omitempty" dynamodbav:"carrier,omitempty"`
	// @gotags: dynamodbav:"amount,omitempty"
	Amount *Amount `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty" dynamodbav:"amount,omitempty"`
	// @gotags: dynamodbav:"transit_time,omitempty"
	TransitTime string `protobuf:"bytes,5,opt,name=transit_time,json=transitTime,proto3" json:"transit_time,omitempty" dynamodbav:"transit_time,omitempty"`
	// @gotags: dynamodbav:"guranteed,omitempty"
	Guranteed bool `protobuf:"varint,6,opt,name=guranteed,proto3" json:"guranteed,omitempty" dynamodbav:"guranteed,omitempty"`
}

func (x *BidEntity) Reset() {
	*x = BidEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidEntity) ProtoMessage() {}

func (x *BidEntity) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BidEntity.ProtoReflect.Descriptor instead.
func (*BidEntity) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{2}
}

func (x *BidEntity) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *BidEntity) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *BidEntity) GetCarrier() *Carrier {
	if x != nil {
		return x.Carrier
	}
	return nil
}

func (x *BidEntity) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *BidEntity) GetTransitTime() string {
	if x != nil {
		return x.TransitTime
	}
	return ""
}

func (x *BidEntity) GetGuranteed() bool {
	if x != nil {
		return x.Guranteed
	}
	return false
}

type DynamoQuoteBidEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"pk,omitempty"
	Pk string `protobuf:"bytes,1,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk,omitempty"`
	// @gotags: dynamodbav:"sk,omitempty"
	Sk string `protobuf:"bytes,2,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk,omitempty"`
	// @gotags: dynamodbav:"bid_pk,omitempty"
	BidPk string `protobuf:"bytes,3,opt,name=bid_pk,json=bidPk,proto3" json:"bid_pk,omitempty" dynamodbav:"bid_pk,omitempty"`
	// @gotags: dynamodbav:"bid_sk,omitempty"
	BidSk string `protobuf:"bytes,4,opt,name=bid_sk,json=bidSk,proto3" json:"bid_sk,omitempty" dynamodbav:"bid_sk,omitempty"`
	// @gotags: dynamodbav:"bid_entity,omitempty"
	BidEntity *BidEntity `protobuf:"bytes,5,opt,name=bid_entity,json=bidEntity,proto3" json:"bid_entity,omitempty" dynamodbav:"bid_entity,omitempty"`
}

func (x *DynamoQuoteBidEntity) Reset() {
	*x = DynamoQuoteBidEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DynamoQuoteBidEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DynamoQuoteBidEntity) ProtoMessage() {}

func (x *DynamoQuoteBidEntity) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DynamoQuoteBidEntity.ProtoReflect.Descriptor instead.
func (*DynamoQuoteBidEntity) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{3}
}

func (x *DynamoQuoteBidEntity) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *DynamoQuoteBidEntity) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *DynamoQuoteBidEntity) GetBidPk() string {
	if x != nil {
		return x.BidPk
	}
	return ""
}

func (x *DynamoQuoteBidEntity) GetBidSk() string {
	if x != nil {
		return x.BidSk
	}
	return ""
}

func (x *DynamoQuoteBidEntity) GetBidEntity() *BidEntity {
	if x != nil {
		return x.BidEntity
	}
	return nil
}

type QuoteBids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_entity,omitempty"
	QuoteEntity *QuoteEntity `protobuf:"bytes,1,opt,name=quote_entity,json=quoteEntity,proto3" json:"quote_entity,omitempty" dynamodbav:"quote_entity,omitempty"`
	// @gotags: dynamodbav:"quote_bids,omitempty"
	QuoteBids []*BidEntity `protobuf:"bytes,2,rep,name=quote_bids,json=quoteBids,proto3" json:"quote_bids,omitempty" dynamodbav:"quote_bids,omitempty"`
}

func (x *QuoteBids) Reset() {
	*x = QuoteBids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteBids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteBids) ProtoMessage() {}

func (x *QuoteBids) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteBids.ProtoReflect.Descriptor instead.
func (*QuoteBids) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{4}
}

func (x *QuoteBids) GetQuoteEntity() *QuoteEntity {
	if x != nil {
		return x.QuoteEntity
	}
	return nil
}

func (x *QuoteBids) GetQuoteBids() []*BidEntity {
	if x != nil {
		return x.QuoteBids
	}
	return nil
}

type FetchQuotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"start_from,omitempty"
	StartFrom string `protobuf:"bytes,1,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty" dynamodbav:"start_from,omitempty"`
	// @gotags: dynamodbav:"end_on,omitempty"
	EndOn string `protobuf:"bytes,2,opt,name=end_on,json=endOn,proto3" json:"end_on,omitempty" dynamodbav:"end_on,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,3,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
}

func (x *FetchQuotes) Reset() {
	*x = FetchQuotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchQuotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchQuotes) ProtoMessage() {}

func (x *FetchQuotes) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchQuotes.ProtoReflect.Descriptor instead.
func (*FetchQuotes) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{5}
}

func (x *FetchQuotes) GetStartFrom() string {
	if x != nil {
		return x.StartFrom
	}
	return ""
}

func (x *FetchQuotes) GetEndOn() string {
	if x != nil {
		return x.EndOn
	}
	return ""
}

func (x *FetchQuotes) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

// @gotags: dynamodbav:"quote_entities,omitempty" json:"quote_entities,omitempty"
type QuoteEntities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_entities,omitempty"
	QuoteEntities []*QuoteEntity `protobuf:"bytes,1,rep,name=quote_entities,json=quoteEntities,proto3" json:"quote_entities,omitempty" dynamodbav:"quote_entities,omitempty"`
}

func (x *QuoteEntities) Reset() {
	*x = QuoteEntities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteEntities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteEntities) ProtoMessage() {}

func (x *QuoteEntities) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteEntities.ProtoReflect.Descriptor instead.
func (*QuoteEntities) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{6}
}

func (x *QuoteEntities) GetQuoteEntities() []*QuoteEntity {
	if x != nil {
		return x.QuoteEntities
	}
	return nil
}

type DeleteQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
}

func (x *DeleteQuote) Reset() {
	*x = DeleteQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_quote_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQuote) ProtoMessage() {}

func (x *DeleteQuote) ProtoReflect() protoreflect.Message {
	mi := &file_quote_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQuote.ProtoReflect.Descriptor instead.
func (*DeleteQuote) Descriptor() ([]byte, []int) {
	return file_quote_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteQuote) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

var File_quote_proto protoreflect.FileDescriptor

var file_quote_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x15, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xe4, 0x02, 0x0a, 0x0b, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x3d, 0x0a,
	0x0f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0e, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x48, 0x0a, 0x13,
	0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x12, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x11, 0x44, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6b, 0x12,
	0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x50, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x5f, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x53, 0x6b, 0x12, 0x34, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xc6, 0x01, 0x0a, 0x09,
	0x42, 0x69, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x73, 0x6b, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x72,
	0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x65, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x42, 0x69, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x70, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6b, 0x12, 0x15, 0x0a,
	0x06, 0x62, 0x69, 0x64, 0x5f, 0x70, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x69, 0x64, 0x50, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x5f, 0x73, 0x6b, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x64, 0x53, 0x6b, 0x12, 0x2e, 0x0a, 0x0a, 0x62,
	0x69, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x09, 0x62, 0x69, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x71, 0x0a, 0x09, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x42, 0x69, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x52, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e,
	0x0a, 0x0a, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x62, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x69, 0x64, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x09, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x42, 0x69, 0x64, 0x73, 0x22, 0x64,
	0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x15, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e,
	0x64, 0x4f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x28, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x42, 0x68, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73,
	0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_quote_proto_goTypes = []interface{}{
	(*QuoteEntity)(nil),          // 0: user.QuoteEntity
	(*DynamoQuoteEntity)(nil),    // 1: user.DynamoQuoteEntity
	(*BidEntity)(nil),            // 2: user.BidEntity
	(*DynamoQuoteBidEntity)(nil), // 3: user.DynamoQuoteBidEntity
	(*QuoteBids)(nil),            // 4: user.QuoteBids
	(*FetchQuotes)(nil),          // 5: user.FetchQuotes
	(*QuoteEntities)(nil),        // 6: user.QuoteEntities
	(*DeleteQuote)(nil),          // 7: user.deleteQuote
	(*ShippingDetail)(nil),       // 8: user.ShippingDetail
	(AdditionalService)(0),       // 9: user.AdditionalService
	(*Commodity)(nil),            // 10: user.Commodity
	(*Carrier)(nil),              // 11: user.Carrier
	(*Amount)(nil),               // 12: user.Amount
}
var file_quote_proto_depIdxs = []int32{
	8,  // 0: user.QuoteEntity.shipping_detail:type_name -> user.ShippingDetail
	9,  // 1: user.QuoteEntity.additional_services:type_name -> user.AdditionalService
	10, // 2: user.QuoteEntity.commodities:type_name -> user.Commodity
	0,  // 3: user.DynamoQuoteEntity.quote_entity:type_name -> user.QuoteEntity
	11, // 4: user.BidEntity.carrier:type_name -> user.Carrier
	12, // 5: user.BidEntity.amount:type_name -> user.Amount
	2,  // 6: user.DynamoQuoteBidEntity.bid_entity:type_name -> user.BidEntity
	0,  // 7: user.QuoteBids.quote_entity:type_name -> user.QuoteEntity
	2,  // 8: user.QuoteBids.quote_bids:type_name -> user.BidEntity
	0,  // 9: user.QuoteEntities.quote_entities:type_name -> user.QuoteEntity
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_quote_proto_init() }
func file_quote_proto_init() {
	if File_quote_proto != nil {
		return
	}
	file_shipping_detail_proto_init()
	file_additional_service_proto_init()
	file_commodity_proto_init()
	file_carrier_proto_init()
	file_amount_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteEntity); i {
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
			switch v := v.(*DynamoQuoteEntity); i {
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
			switch v := v.(*BidEntity); i {
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
		file_quote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DynamoQuoteBidEntity); i {
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
		file_quote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteBids); i {
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
		file_quote_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchQuotes); i {
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
		file_quote_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteEntities); i {
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
		file_quote_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQuote); i {
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
			NumMessages:   8,
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

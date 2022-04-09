// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: bid.proto

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

type Bid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
	// @gotags: dynamodbav:"bid_id,omitempty"
	BidId string `protobuf:"bytes,2,opt,name=bid_id,json=bidId,proto3" json:"bid_id,omitempty" dynamodbav:"bid_id,omitempty"`
	// @gotags: dynamodbav:"carrier,omitempty"
	Carrier *Carrier `protobuf:"bytes,3,opt,name=carrier,proto3" json:"carrier,omitempty" dynamodbav:"carrier,omitempty"`
	// @gotags: dynamodbav:"amount,omitempty"
	Amount *Amount `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty" dynamodbav:"amount,omitempty"`
	// @gotags: dynamodbav:"transit_time,omitempty"
	TransitTime string `protobuf:"bytes,5,opt,name=transit_time,json=transitTime,proto3" json:"transit_time,omitempty" dynamodbav:"transit_time,omitempty"`
	// @gotags: dynamodbav:"guranteed,omitempty"
	Guranteed bool `protobuf:"varint,6,opt,name=guranteed,proto3" json:"guranteed,omitempty" dynamodbav:"guranteed,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,7,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
}

func (x *Bid) Reset() {
	*x = Bid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bid) ProtoMessage() {}

func (x *Bid) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bid.ProtoReflect.Descriptor instead.
func (*Bid) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{0}
}

func (x *Bid) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *Bid) GetBidId() string {
	if x != nil {
		return x.BidId
	}
	return ""
}

func (x *Bid) GetCarrier() *Carrier {
	if x != nil {
		return x.Carrier
	}
	return nil
}

func (x *Bid) GetAmount() *Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Bid) GetTransitTime() string {
	if x != nil {
		return x.TransitTime
	}
	return ""
}

func (x *Bid) GetGuranteed() bool {
	if x != nil {
		return x.Guranteed
	}
	return false
}

func (x *Bid) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

type Bids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bids,omitempty"
	Bids []*Bid `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids,omitempty" dynamodbav:"bids,omitempty"`
}

func (x *Bids) Reset() {
	*x = Bids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bids) ProtoMessage() {}

func (x *Bids) ProtoReflect() protoreflect.Message {
	mi := &file_bid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bids.ProtoReflect.Descriptor instead.
func (*Bids) Descriptor() ([]byte, []int) {
	return file_bid_proto_rawDescGZIP(), []int{1}
}

func (x *Bids) GetBids() []*Bid {
	if x != nil {
		return x.Bids
	}
	return nil
}

var File_bid_proto protoreflect.FileDescriptor

var file_bid_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x1a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8,
	0x01, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x69, 0x64, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x24, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x67,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x04, 0x62, 0x69, 0x64,
	0x73, 0x12, 0x1d, 0x0a, 0x04, 0x62, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x04, 0x62, 0x69, 0x64, 0x73,
	0x42, 0x66, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x42, 0x08, 0x42, 0x69,
	0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73,
	0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58,
	0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xe2, 0x02,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bid_proto_rawDescOnce sync.Once
	file_bid_proto_rawDescData = file_bid_proto_rawDesc
)

func file_bid_proto_rawDescGZIP() []byte {
	file_bid_proto_rawDescOnce.Do(func() {
		file_bid_proto_rawDescData = protoimpl.X.CompressGZIP(file_bid_proto_rawDescData)
	})
	return file_bid_proto_rawDescData
}

var file_bid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bid_proto_goTypes = []interface{}{
	(*Bid)(nil),     // 0: user.bid
	(*Bids)(nil),    // 1: user.bids
	(*Carrier)(nil), // 2: user.carrier
	(*Amount)(nil),  // 3: user.amount
}
var file_bid_proto_depIdxs = []int32{
	2, // 0: user.bid.carrier:type_name -> user.carrier
	3, // 1: user.bid.amount:type_name -> user.amount
	0, // 2: user.bids.bids:type_name -> user.bid
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bid_proto_init() }
func file_bid_proto_init() {
	if File_bid_proto != nil {
		return
	}
	file_amount_proto_init()
	file_carrier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bid); i {
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
		file_bid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bids); i {
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
			RawDescriptor: file_bid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bid_proto_goTypes,
		DependencyIndexes: file_bid_proto_depIdxs,
		MessageInfos:      file_bid_proto_msgTypes,
	}.Build()
	File_bid_proto = out.File
	file_bid_proto_rawDesc = nil
	file_bid_proto_goTypes = nil
	file_bid_proto_depIdxs = nil
}
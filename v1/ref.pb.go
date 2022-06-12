// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: ref.proto

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

type Refereses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"vendor_bol,omitempty"
	VendorBol string `protobuf:"bytes,1,opt,name=vendor_bol,json=vendorBol,proto3" json:"vendor_bol,omitempty"`
	// @gotags: dynamodbav:"bol,omitempty"
	Bol string `protobuf:"bytes,2,opt,name=bol,proto3" json:"bol,omitempty"`
	// @gotags: dynamodbav:"vendor_reference_id,omitempty"
	VendorReferenceId string `protobuf:"bytes,3,opt,name=vendor_reference_id,json=vendorReferenceId,proto3" json:"vendor_reference_id,omitempty"`
	// @gotags: dynamodbav:"pickup_no,omitempty"
	PickupNo string `protobuf:"bytes,4,opt,name=pickup_no,json=pickupNo,proto3" json:"pickup_no,omitempty"`
	// @gotags: dynamodbav:"invoice_no,omitempty"
	InvoiceNo string `protobuf:"bytes,5,opt,name=invoice_no,json=invoiceNo,proto3" json:"invoice_no,omitempty"`
	// @gotags: dynamodbav:"vendor_invoice_no,omitempty"
	VendorInvoiceNo string `protobuf:"bytes,6,opt,name=vendor_invoice_no,json=vendorInvoiceNo,proto3" json:"vendor_invoice_no,omitempty"`
}

func (x *Refereses) Reset() {
	*x = Refereses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ref_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Refereses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Refereses) ProtoMessage() {}

func (x *Refereses) ProtoReflect() protoreflect.Message {
	mi := &file_ref_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Refereses.ProtoReflect.Descriptor instead.
func (*Refereses) Descriptor() ([]byte, []int) {
	return file_ref_proto_rawDescGZIP(), []int{0}
}

func (x *Refereses) GetVendorBol() string {
	if x != nil {
		return x.VendorBol
	}
	return ""
}

func (x *Refereses) GetBol() string {
	if x != nil {
		return x.Bol
	}
	return ""
}

func (x *Refereses) GetVendorReferenceId() string {
	if x != nil {
		return x.VendorReferenceId
	}
	return ""
}

func (x *Refereses) GetPickupNo() string {
	if x != nil {
		return x.PickupNo
	}
	return ""
}

func (x *Refereses) GetInvoiceNo() string {
	if x != nil {
		return x.InvoiceNo
	}
	return ""
}

func (x *Refereses) GetVendorInvoiceNo() string {
	if x != nil {
		return x.VendorInvoiceNo
	}
	return ""
}

var File_ref_proto protoreflect.FileDescriptor

var file_ref_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22,
	0xd4, 0x01, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x73, 0x65, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x6f, 0x6c, 0x12, 0x2e,
	0x0a, 0x13, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x42, 0x5c, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31,
	0x42, 0x08, 0x52, 0x65, 0x66, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74,
	0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02,
	0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ref_proto_rawDescOnce sync.Once
	file_ref_proto_rawDescData = file_ref_proto_rawDesc
)

func file_ref_proto_rawDescGZIP() []byte {
	file_ref_proto_rawDescOnce.Do(func() {
		file_ref_proto_rawDescData = protoimpl.X.CompressGZIP(file_ref_proto_rawDescData)
	})
	return file_ref_proto_rawDescData
}

var file_ref_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ref_proto_goTypes = []interface{}{
	(*Refereses)(nil), // 0: v1.refereses
}
var file_ref_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ref_proto_init() }
func file_ref_proto_init() {
	if File_ref_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ref_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Refereses); i {
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
			RawDescriptor: file_ref_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ref_proto_goTypes,
		DependencyIndexes: file_ref_proto_depIdxs,
		MessageInfos:      file_ref_proto_msgTypes,
	}.Build()
	File_ref_proto = out.File
	file_ref_proto_rawDesc = nil
	file_ref_proto_goTypes = nil
	file_ref_proto_depIdxs = nil
}

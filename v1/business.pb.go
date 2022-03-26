// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: business.proto

package v1

import (
	_ "github.com/todaystechs/go_gen/v1/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BusinessData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type                 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	BusinessName         *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=business_name,json=businessName,proto3" json:"business_name,omitempty"`
	BusinessId           *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	BusinessEmail        *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=business_email,json=businessEmail,proto3" json:"business_email,omitempty"`
	AccountingEmail      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=accounting_email,json=accountingEmail,proto3" json:"accounting_email,omitempty"`
	CustomerServiceEmail *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=customer_service_email,json=customerServiceEmail,proto3" json:"customer_service_email,omitempty"`
	HighPriorityEmail    *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=high_priority_email,json=highPriorityEmail,proto3" json:"high_priority_email,omitempty"`
	AvatarUrl            *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	CreatedOn            *timestamppb.Timestamp  `protobuf:"bytes,9,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	UpdatedOn            *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
	DeletedOn            *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=deleted_on,json=deletedOn,proto3" json:"deleted_on,omitempty"`
	PhoneNumbers         []*PhoneNumberData      `protobuf:"bytes,12,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	SK                   *wrapperspb.StringValue `protobuf:"bytes,13,opt,name=SK,proto3" json:"SK,omitempty"`
	PK                   *wrapperspb.StringValue `protobuf:"bytes,14,opt,name=PK,proto3" json:"PK,omitempty"`
}

func (x *BusinessData) Reset() {
	*x = BusinessData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessData) ProtoMessage() {}

func (x *BusinessData) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessData.ProtoReflect.Descriptor instead.
func (*BusinessData) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessData) GetType() *wrapperspb.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *BusinessData) GetBusinessName() *wrapperspb.StringValue {
	if x != nil {
		return x.BusinessName
	}
	return nil
}

func (x *BusinessData) GetBusinessId() *wrapperspb.StringValue {
	if x != nil {
		return x.BusinessId
	}
	return nil
}

func (x *BusinessData) GetBusinessEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.BusinessEmail
	}
	return nil
}

func (x *BusinessData) GetAccountingEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.AccountingEmail
	}
	return nil
}

func (x *BusinessData) GetCustomerServiceEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.CustomerServiceEmail
	}
	return nil
}

func (x *BusinessData) GetHighPriorityEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.HighPriorityEmail
	}
	return nil
}

func (x *BusinessData) GetAvatarUrl() *wrapperspb.StringValue {
	if x != nil {
		return x.AvatarUrl
	}
	return nil
}

func (x *BusinessData) GetCreatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *BusinessData) GetUpdatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedOn
	}
	return nil
}

func (x *BusinessData) GetDeletedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedOn
	}
	return nil
}

func (x *BusinessData) GetPhoneNumbers() []*PhoneNumberData {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *BusinessData) GetSK() *wrapperspb.StringValue {
	if x != nil {
		return x.SK
	}
	return nil
}

func (x *BusinessData) GetPK() *wrapperspb.StringValue {
	if x != nil {
		return x.PK
	}
	return nil
}

type BusinessDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []*wrapperspb.StringValue `protobuf:"bytes,1,rep,name=token,proto3" json:"token,omitempty"`
}

func (x *BusinessDataRequest) Reset() {
	*x = BusinessDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessDataRequest) ProtoMessage() {}

func (x *BusinessDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessDataRequest.ProtoReflect.Descriptor instead.
func (*BusinessDataRequest) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{1}
}

func (x *BusinessDataRequest) GetToken() []*wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_business_proto protoreflect.FileDescriptor

var file_business_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x06, 0x0a, 0x0c, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d,
	0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x43, 0x0a,
	0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x47, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x52, 0x0a, 0x16, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x4c, 0x0a, 0x13, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x68, 0x69, 0x67, 0x68,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3b, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x3e, 0x0a, 0x0d, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x53,
	0x4b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x53, 0x4b, 0x12, 0x2c, 0x0a, 0x02, 0x50, 0x4b, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x02, 0x50, 0x4b, 0x22, 0x49, 0x0a, 0x13, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x7f, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x42, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67,
	0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x42, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xca, 0x02, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0xe2, 0x02, 0x14, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_business_proto_rawDescOnce sync.Once
	file_business_proto_rawDescData = file_business_proto_rawDesc
)

func file_business_proto_rawDescGZIP() []byte {
	file_business_proto_rawDescOnce.Do(func() {
		file_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_business_proto_rawDescData)
	})
	return file_business_proto_rawDescData
}

var file_business_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_business_proto_goTypes = []interface{}{
	(*BusinessData)(nil),           // 0: business.BusinessData
	(*BusinessDataRequest)(nil),    // 1: business.BusinessDataRequest
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*PhoneNumberData)(nil),        // 4: business.PhoneNumberData
}
var file_business_proto_depIdxs = []int32{
	2,  // 0: business.BusinessData.type:type_name -> google.protobuf.StringValue
	2,  // 1: business.BusinessData.business_name:type_name -> google.protobuf.StringValue
	2,  // 2: business.BusinessData.business_id:type_name -> google.protobuf.StringValue
	2,  // 3: business.BusinessData.business_email:type_name -> google.protobuf.StringValue
	2,  // 4: business.BusinessData.accounting_email:type_name -> google.protobuf.StringValue
	2,  // 5: business.BusinessData.customer_service_email:type_name -> google.protobuf.StringValue
	2,  // 6: business.BusinessData.high_priority_email:type_name -> google.protobuf.StringValue
	2,  // 7: business.BusinessData.avatar_url:type_name -> google.protobuf.StringValue
	3,  // 8: business.BusinessData.created_on:type_name -> google.protobuf.Timestamp
	3,  // 9: business.BusinessData.updated_on:type_name -> google.protobuf.Timestamp
	3,  // 10: business.BusinessData.deleted_on:type_name -> google.protobuf.Timestamp
	4,  // 11: business.BusinessData.phone_numbers:type_name -> business.PhoneNumberData
	2,  // 12: business.BusinessData.SK:type_name -> google.protobuf.StringValue
	2,  // 13: business.BusinessData.PK:type_name -> google.protobuf.StringValue
	2,  // 14: business.BusinessDataRequest.token:type_name -> google.protobuf.StringValue
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_business_proto_init() }
func file_business_proto_init() {
	if File_business_proto != nil {
		return
	}
	file_phone_number_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessData); i {
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
		file_business_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessDataRequest); i {
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
			RawDescriptor: file_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_business_proto_goTypes,
		DependencyIndexes: file_business_proto_depIdxs,
		MessageInfos:      file_business_proto_msgTypes,
	}.Build()
	File_business_proto = out.File
	file_business_proto_rawDesc = nil
	file_business_proto_goTypes = nil
	file_business_proto_depIdxs = nil
}

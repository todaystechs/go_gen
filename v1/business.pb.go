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
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type,omitempty"`
	// @gotags: dynamodbav:"business_name,omitempty"
	BusinessName string `protobuf:"bytes,2,opt,name=business_name,json=business_name,omitempty,proto3" json:"business_name,omitempty" dynamodbav:"business_name,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,3,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
	// @gotags: dynamodbav:"business_email,omitempty"
	BusinessEmail string `protobuf:"bytes,4,opt,name=business_email,json=businessEmail,proto3" json:"business_email,omitempty" dynamodbav:"business_email,omitempty"`
	// @gotags: dynamodbav:"accounting_email,omitempty"
	AccountingEmail string `protobuf:"bytes,5,opt,name=accounting_email,json=accountingEmail,proto3" json:"accounting_email,omitempty" dynamodbav:"accounting_email,omitempty"`
	// @gotags: dynamodbav:"customer_service_email,omitempty"
	CustomerServiceEmail string `protobuf:"bytes,6,opt,name=customer_service_email,json=customerServiceEmail,proto3" json:"customer_service_email,omitempty" dynamodbav:"customer_service_email,omitempty"`
	// @gotags: dynamodbav:"high_priority_email,omitempty"
	HighPriorityEmail string `protobuf:"bytes,7,opt,name=high_priority_email,json=highPriorityEmail,proto3" json:"high_priority_email,omitempty" dynamodbav:"high_priority_email,omitempty"`
	// @gotags: dynamodbav:"avatar_url,omitempty"
	AvatarUrl string `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty" dynamodbav:"avatar_url,omitempty"`
	// @gotags: dynamodbav:"admin_email,omitempty"
	AdminEmail string `protobuf:"bytes,9,opt,name=admin_email,json=adminEmail,proto3" json:"admin_email,omitempty" dynamodbav:"admin_email,omitempty"`
	// @gotags: dynamodbav:"created_on,omitempty"
	CreatedOn string `protobuf:"bytes,10,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty" dynamodbav:"created_on,omitempty"`
	// @gotags: dynamodbav:"updated_on,omitempty"
	UpdatedOn string `protobuf:"bytes,11,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty" dynamodbav:"updated_on,omitempty"`
	// @gotags: dynamodbav:"deleted_on,omitempty"
	DeletedOn string `protobuf:"bytes,12,opt,name=deleted_on,json=deletedOn,proto3" json:"deleted_on,omitempty" dynamodbav:"deleted_on,omitempty"`
	// @gotags: dynamodbav:"phone_numbers,omitempty"
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,13,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty" dynamodbav:"phone_numbers,omitempty"`
	// @gotags: dynamodbav:"sk,omitempty"
	Sk string `protobuf:"bytes,14,opt,name=sk,proto3" json:"sk,omitempty" dynamodbav:"sk,omitempty"`
	// @gotags: dynamodbav:"pk,omitempty"
	Pk string `protobuf:"bytes,15,opt,name=pk,proto3" json:"pk,omitempty" dynamodbav:"pk,omitempty"`
	// @gotags: dynamodbav:"address,omitempty"
	Address *Address `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty" dynamodbav:"address,omitempty"`
	// @gotags: dynamodbav:"address_updated,omitempty"
	AddressUpdated bool `protobuf:"varint,17,opt,name=address_updated,json=addressUpdated,proto3" json:"address_updated,omitempty" dynamodbav:"address_updated,omitempty"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{0}
}

func (x *Business) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Business) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *Business) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Business) GetBusinessEmail() string {
	if x != nil {
		return x.BusinessEmail
	}
	return ""
}

func (x *Business) GetAccountingEmail() string {
	if x != nil {
		return x.AccountingEmail
	}
	return ""
}

func (x *Business) GetCustomerServiceEmail() string {
	if x != nil {
		return x.CustomerServiceEmail
	}
	return ""
}

func (x *Business) GetHighPriorityEmail() string {
	if x != nil {
		return x.HighPriorityEmail
	}
	return ""
}

func (x *Business) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Business) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

func (x *Business) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *Business) GetUpdatedOn() string {
	if x != nil {
		return x.UpdatedOn
	}
	return ""
}

func (x *Business) GetDeletedOn() string {
	if x != nil {
		return x.DeletedOn
	}
	return ""
}

func (x *Business) GetPhoneNumbers() []*PhoneNumber {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *Business) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *Business) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *Business) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Business) GetAddressUpdated() bool {
	if x != nil {
		return x.AddressUpdated
	}
	return false
}

type Businesses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"businesses,omitempty"
	Businesses []*Business `protobuf:"bytes,1,rep,name=businesses,proto3" json:"businesses,omitempty" dynamodbav:"businesses,omitempty"`
}

func (x *Businesses) Reset() {
	*x = Businesses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Businesses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Businesses) ProtoMessage() {}

func (x *Businesses) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Businesses.ProtoReflect.Descriptor instead.
func (*Businesses) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{1}
}

func (x *Businesses) GetBusinesses() []*Business {
	if x != nil {
		return x.Businesses
	}
	return nil
}

var File_business_proto protoreflect.FileDescriptor

var file_business_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf7, 0x04, 0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05,
	0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x17, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x69, 0x67, 0x68, 0x5f,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x4f, 0x6e, 0x12, 0x35, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0c, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x73,
	0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x70,
	0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x25, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76,
	0x31, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x65, 0x73, 0x42, 0x61, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x31, 0x42, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x6f, 0x64, 0x61, 0x79, 0x73, 0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca,
	0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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
	(*Business)(nil),    // 0: v1.business
	(*Businesses)(nil),  // 1: v1.businesses
	(*PhoneNumber)(nil), // 2: v1.phone_number
	(*Address)(nil),     // 3: v1.address
}
var file_business_proto_depIdxs = []int32{
	2, // 0: v1.business.phone_numbers:type_name -> v1.phone_number
	3, // 1: v1.business.address:type_name -> v1.address
	0, // 2: v1.businesses.businesses:type_name -> v1.business
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_business_proto_init() }
func file_business_proto_init() {
	if File_business_proto != nil {
		return
	}
	file_phone_number_proto_init()
	file_address_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
			switch v := v.(*Businesses); i {
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

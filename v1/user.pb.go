// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: user.proto

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

// @gotags: dynamodbav:"user_data,omitempty"
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// @gotags: dynamodbav:"cognito_id,omitempty"
	CognitoId string `protobuf:"bytes,2,opt,name=cognito_id,json=cognitoId,proto3" json:"cognito_id,omitempty"`
	// @gotags: dynamodbav:"user_name,omitempty"
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// @gotags: dynamodbav:"user_id,omitempty"
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// @gotags: dynamodbav:"first_name,omitempty"
	FirstName string `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// @gotags: dynamodbav:"middle_name,omitempty"
	MiddleName string `protobuf:"bytes,6,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	// @gotags: dynamodbav:"last_name,omitempty"
	LastName string `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	// @gotags: dynamodbav:"hashed_password,omitempty"
	HashedPassword string `protobuf:"bytes,9,opt,name=hashed_password,json=hashedPassword,proto3" json:"hashed_password,omitempty"`
	// @gotags: dynamodbav:"avatar_url,omitempty"
	AvatarUrl string `protobuf:"bytes,10,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	// @gotags: dynamodbav:"roles,omitempty"
	Roles []Role `protobuf:"varint,11,rep,packed,name=roles,proto3,enum=user.Role" json:"roles,omitempty"`
	// @gotags: dynamodbav:"new_password_required,omitempty"
	NewPasswordRequired bool `protobuf:"varint,12,opt,name=new_password_required,json=newPasswordRequired,proto3" json:"new_password_required,omitempty"`
	// @gotags: dynamodbav:"password_changed_at,omitempty"
	PasswordChangedAt string `protobuf:"bytes,13,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	// @gotags: dynamodbav:"created_on,omitempty"
	CreatedOn string `protobuf:"bytes,14,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	// @gotags: dynamodbav:"updated_on,omitempty"
	UpdatedOn string `protobuf:"bytes,15,opt,name=updated_on,json=updatedOn,proto3" json:"updated_on,omitempty"`
	// @gotags: dynamodbav:"deleted_on,omitempty"
	DeletedOn string `protobuf:"bytes,16,opt,name=deleted_on,json=deletedOn,proto3" json:"deleted_on,omitempty"`
	// @gotags: dynamodbav:"phone_numbers,omitempty"
	PhoneNumbers []*PhoneNumber `protobuf:"bytes,17,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	// @gotags: dynamodbav:"email_verified,omitempty"
	EmailVerified bool `protobuf:"varint,18,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	// @gotags: dynamodbav:"reset_password_tokens,omitempty"
	ResetPasswordTokens []string `protobuf:"bytes,19,rep,name=reset_password_tokens,json=resetPasswordTokens,proto3" json:"reset_password_tokens,omitempty"`
	// @gotags: dynamodbav:"reset_password_token,omitempty"
	ResetPasswordToken map[string]float64 `protobuf:"bytes,20,rep,name=reset_password_token,json=resetPasswordToken,proto3" json:"reset_password_token,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// @gotags: dynamodbav:"sessions,omitempty"
	Sessions []string `protobuf:"bytes,21,rep,name=sessions,proto3" json:"sessions,omitempty"`
	// @gotags: dynamodbav:"sk,omitempty"
	Pk string `protobuf:"bytes,22,opt,name=pk,proto3" json:"pk,omitempty"`
	// @gotags: dynamodbav:"pk,omitempty"
	Sk string `protobuf:"bytes,23,opt,name=sk,proto3" json:"sk,omitempty"`
	// @gotags: dynamodbav:"business_ids,omitempty"
	BusinessIds []string `protobuf:"bytes,24,rep,name=business_ids,json=businessIds,proto3" json:"business_ids,omitempty"`
	// @gotags: dynamodbav:"unsuscribed_to_marketing_email,omitempty"
	UnsuscribedToMarketingEmail bool `protobuf:"varint,25,opt,name=unsuscribed_to_marketing_email,json=unsuscribedToMarketingEmail,proto3" json:"unsuscribed_to_marketing_email,omitempty"`
	// @gotags: dynamodbav:"status,omitempty"
	Status *Status `protobuf:"bytes,26,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *User) GetCognitoId() string {
	if x != nil {
		return x.CognitoId
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetHashedPassword() string {
	if x != nil {
		return x.HashedPassword
	}
	return ""
}

func (x *User) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *User) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *User) GetNewPasswordRequired() bool {
	if x != nil {
		return x.NewPasswordRequired
	}
	return false
}

func (x *User) GetPasswordChangedAt() string {
	if x != nil {
		return x.PasswordChangedAt
	}
	return ""
}

func (x *User) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

func (x *User) GetUpdatedOn() string {
	if x != nil {
		return x.UpdatedOn
	}
	return ""
}

func (x *User) GetDeletedOn() string {
	if x != nil {
		return x.DeletedOn
	}
	return ""
}

func (x *User) GetPhoneNumbers() []*PhoneNumber {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetResetPasswordTokens() []string {
	if x != nil {
		return x.ResetPasswordTokens
	}
	return nil
}

func (x *User) GetResetPasswordToken() map[string]float64 {
	if x != nil {
		return x.ResetPasswordToken
	}
	return nil
}

func (x *User) GetSessions() []string {
	if x != nil {
		return x.Sessions
	}
	return nil
}

func (x *User) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *User) GetSk() string {
	if x != nil {
		return x.Sk
	}
	return ""
}

func (x *User) GetBusinessIds() []string {
	if x != nil {
		return x.BusinessIds
	}
	return nil
}

func (x *User) GetUnsuscribedToMarketingEmail() bool {
	if x != nil {
		return x.UnsuscribedToMarketingEmail
	}
	return false
}

func (x *User) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x08, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x27, 0x0a, 0x0f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x44, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x22, 0xba, 0xe9, 0xc0, 0x03, 0x1d, 0x82, 0x01, 0x1a, 0x18, 0x00, 0x18, 0x01,
	0x18, 0x02, 0x18, 0x03, 0x18, 0x04, 0x18, 0x05, 0x18, 0x06, 0x18, 0x07, 0x18, 0x08, 0x18, 0x09,
	0x18, 0x0a, 0x18, 0x0b, 0x18, 0x0c, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x6e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x37,
	0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x32,
	0x0a, 0x15, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x54, 0x0a, 0x14, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x12, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x70, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6b, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03,
	0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x73, 0x12, 0x43, 0x0a, 0x1e, 0x75, 0x6e, 0x73, 0x75, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x75, 0x6e, 0x73, 0x75,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x54, 0x6f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x45, 0x0a,
	0x17, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x67, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x64, 0x61, 0x79, 0x73,
	0x74, 0x65, 0x63, 0x68, 0x73, 0x2f, 0x67, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0xca, 0x02, 0x04, 0x55,
	0x73, 0x65, 0x72, 0xe2, 0x02, 0x10, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x55, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil),        // 0: user.user
	nil,                 // 1: user.user.ResetPasswordTokenEntry
	(Role)(0),           // 2: user.role
	(*PhoneNumber)(nil), // 3: user.phone_number
	(*Status)(nil),      // 4: user.status
}
var file_user_proto_depIdxs = []int32{
	2, // 0: user.user.roles:type_name -> user.role
	3, // 1: user.user.phone_numbers:type_name -> user.phone_number
	1, // 2: user.user.reset_password_token:type_name -> user.user.ResetPasswordTokenEntry
	4, // 3: user.user.status:type_name -> user.status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_role_proto_init()
	file_phone_number_proto_init()
	file_password_proto_init()
	file_locked_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on UserData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserDataMultiError, or nil
// if none found.
func (m *UserData) ValidateAll() error {
	return m.validate(true)
}

func (m *UserData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetType()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Type",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "Type",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCognitoId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "CognitoId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "CognitoId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCognitoId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "CognitoId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UserName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UserName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "UserName",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFirstName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "FirstName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "FirstName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFirstName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "FirstName",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMiddleName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "MiddleName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "MiddleName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMiddleName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "MiddleName",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLastName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "LastName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "LastName",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLastName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "LastName",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEmail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "Email",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHashedPassword()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "HashedPassword",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "HashedPassword",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHashedPassword()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "HashedPassword",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAvatarUrl()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "AvatarUrl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "AvatarUrl",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAvatarUrl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "AvatarUrl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNewPasswordRequired()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "NewPasswordRequired",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "NewPasswordRequired",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNewPasswordRequired()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "NewPasswordRequired",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPasswordChangedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "PasswordChangedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "PasswordChangedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPasswordChangedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "PasswordChangedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedOn()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "CreatedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "CreatedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "CreatedOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedOn()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UpdatedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UpdatedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "UpdatedOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDeletedOn()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "DeletedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "DeletedOn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeletedOn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "DeletedOn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPhoneNumbers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserDataValidationError{
					field:  fmt.Sprintf("PhoneNumbers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetEmailVarified()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "EmailVarified",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "EmailVarified",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmailVarified()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "EmailVarified",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetForgotPasswordTokens() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("ForgotPasswordTokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("ForgotPasswordTokens[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserDataValidationError{
					field:  fmt.Sprintf("ForgotPasswordTokens[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSessions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("Sessions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("Sessions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserDataValidationError{
					field:  fmt.Sprintf("Sessions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetSK()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "SK",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "SK",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSK()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "SK",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPK()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "PK",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "PK",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPK()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "PK",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetBusinessIds() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("BusinessIds[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserDataValidationError{
						field:  fmt.Sprintf("BusinessIds[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserDataValidationError{
					field:  fmt.Sprintf("BusinessIds[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetUnsuscribedToMarketingEmail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UnsuscribedToMarketingEmail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "UnsuscribedToMarketingEmail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUnsuscribedToMarketingEmail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "UnsuscribedToMarketingEmail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserDataMultiError(errors)
	}
	return nil
}

// UserDataMultiError is an error wrapping multiple validation errors returned
// by UserData.ValidateAll() if the designated constraints aren't met.
type UserDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserDataMultiError) AllErrors() []error { return m }

// UserDataValidationError is the validation error returned by
// UserData.Validate if the designated constraints aren't met.
type UserDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserDataValidationError) ErrorName() string { return "UserDataValidationError" }

// Error satisfies the builtin error interface
func (e UserDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserDataValidationError{}

// Validate checks the field values on MeData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MeData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MeData with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MeDataMultiError, or nil if none found.
func (m *MeData) ValidateAll() error {
	return m.validate(true)
}

func (m *MeData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetToken() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MeDataValidationError{
						field:  fmt.Sprintf("Token[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MeDataValidationError{
						field:  fmt.Sprintf("Token[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MeDataValidationError{
					field:  fmt.Sprintf("Token[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return MeDataMultiError(errors)
	}
	return nil
}

// MeDataMultiError is an error wrapping multiple validation errors returned by
// MeData.ValidateAll() if the designated constraints aren't met.
type MeDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MeDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MeDataMultiError) AllErrors() []error { return m }

// MeDataValidationError is the validation error returned by MeData.Validate if
// the designated constraints aren't met.
type MeDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeDataValidationError) ErrorName() string { return "MeDataValidationError" }

// Error satisfies the builtin error interface
func (e MeDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMeData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeDataValidationError{}

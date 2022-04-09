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

	// no validation rules for Type

	// no validation rules for CognitoId

	// no validation rules for UserName

	// no validation rules for UserId

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Email

	// no validation rules for HashedPassword

	// no validation rules for AvatarUrl

	// no validation rules for NewPasswordRequired

	// no validation rules for PasswordChangedAt

	// no validation rules for CreatedOn

	// no validation rules for UpdatedOn

	// no validation rules for DeletedOn

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

	// no validation rules for EmailVerified

	// no validation rules for ResetPasswordToken

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for UnsuscribedToMarketingEmail

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserDataValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserDataValidationError{
				field:  "Status",
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

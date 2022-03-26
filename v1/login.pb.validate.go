// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: login.proto

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

// Validate checks the field values on LoginUserData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginUserData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUserData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginUserDataMultiError, or
// nil if none found.
func (m *LoginUserData) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUserData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEmail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserDataValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserDataValidationError{
					field:  "Email",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserDataValidationError{
				field:  "Email",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPassword()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserDataValidationError{
					field:  "Password",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserDataValidationError{
					field:  "Password",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPassword()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserDataValidationError{
				field:  "Password",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoginUserDataMultiError(errors)
	}
	return nil
}

// LoginUserDataMultiError is an error wrapping multiple validation errors
// returned by LoginUserData.ValidateAll() if the designated constraints
// aren't met.
type LoginUserDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserDataMultiError) AllErrors() []error { return m }

// LoginUserDataValidationError is the validation error returned by
// LoginUserData.Validate if the designated constraints aren't met.
type LoginUserDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserDataValidationError) ErrorName() string { return "LoginUserDataValidationError" }

// Error satisfies the builtin error interface
func (e LoginUserDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserDataValidationError{}

// Validate checks the field values on LoginUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginUserResponseMultiError, or nil if none found.
func (m *LoginUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSuccess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "Success",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "Success",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSuccess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserResponseValidationError{
				field:  "Success",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStatusCode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "StatusCode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "StatusCode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatusCode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserResponseValidationError{
				field:  "StatusCode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "Token",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserResponseValidationError{
				field:  "Token",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUserResponseValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUserResponseValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoginUserResponseMultiError(errors)
	}
	return nil
}

// LoginUserResponseMultiError is an error wrapping multiple validation errors
// returned by LoginUserResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserResponseMultiError) AllErrors() []error { return m }

// LoginUserResponseValidationError is the validation error returned by
// LoginUserResponse.Validate if the designated constraints aren't met.
type LoginUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserResponseValidationError) ErrorName() string {
	return "LoginUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LoginUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserResponseValidationError{}

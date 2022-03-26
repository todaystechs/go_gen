// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: token.proto

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

// Validate checks the field values on RefreshTokenData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenDataMultiError, or nil if none found.
func (m *RefreshTokenData) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRefreshToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RefreshTokenDataValidationError{
					field:  "RefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RefreshTokenDataValidationError{
					field:  "RefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RefreshTokenDataValidationError{
				field:  "RefreshToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RefreshTokenDataMultiError(errors)
	}
	return nil
}

// RefreshTokenDataMultiError is an error wrapping multiple validation errors
// returned by RefreshTokenData.ValidateAll() if the designated constraints
// aren't met.
type RefreshTokenDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenDataMultiError) AllErrors() []error { return m }

// RefreshTokenDataValidationError is the validation error returned by
// RefreshTokenData.Validate if the designated constraints aren't met.
type RefreshTokenDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenDataValidationError) ErrorName() string { return "RefreshTokenDataValidationError" }

// Error satisfies the builtin error interface
func (e RefreshTokenDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenDataValidationError{}

// Validate checks the field values on RefreshTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenResponseMultiError, or nil if none found.
func (m *RefreshTokenResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAccessToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RefreshTokenResponseValidationError{
					field:  "AccessToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RefreshTokenResponseValidationError{
					field:  "AccessToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAccessToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RefreshTokenResponseValidationError{
				field:  "AccessToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRefreshToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RefreshTokenResponseValidationError{
					field:  "RefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RefreshTokenResponseValidationError{
					field:  "RefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RefreshTokenResponseValidationError{
				field:  "RefreshToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RefreshTokenResponseMultiError(errors)
	}
	return nil
}

// RefreshTokenResponseMultiError is an error wrapping multiple validation
// errors returned by RefreshTokenResponse.ValidateAll() if the designated
// constraints aren't met.
type RefreshTokenResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenResponseMultiError) AllErrors() []error { return m }

// RefreshTokenResponseValidationError is the validation error returned by
// RefreshTokenResponse.Validate if the designated constraints aren't met.
type RefreshTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenResponseValidationError) ErrorName() string {
	return "RefreshTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenResponseValidationError{}

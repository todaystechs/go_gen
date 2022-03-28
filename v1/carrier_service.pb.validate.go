// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: carrier_service.proto

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

// Validate checks the field values on CarrierServicePing with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CarrierServicePing) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CarrierServicePing with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CarrierServicePingMultiError, or nil if none found.
func (m *CarrierServicePing) ValidateAll() error {
	return m.validate(true)
}

func (m *CarrierServicePing) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHi()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CarrierServicePingValidationError{
					field:  "Hi",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CarrierServicePingValidationError{
					field:  "Hi",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHi()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CarrierServicePingValidationError{
				field:  "Hi",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CarrierServicePingMultiError(errors)
	}
	return nil
}

// CarrierServicePingMultiError is an error wrapping multiple validation errors
// returned by CarrierServicePing.ValidateAll() if the designated constraints
// aren't met.
type CarrierServicePingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CarrierServicePingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CarrierServicePingMultiError) AllErrors() []error { return m }

// CarrierServicePingValidationError is the validation error returned by
// CarrierServicePing.Validate if the designated constraints aren't met.
type CarrierServicePingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CarrierServicePingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CarrierServicePingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CarrierServicePingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CarrierServicePingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CarrierServicePingValidationError) ErrorName() string {
	return "CarrierServicePingValidationError"
}

// Error satisfies the builtin error interface
func (e CarrierServicePingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCarrierServicePing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CarrierServicePingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CarrierServicePingValidationError{}

// Validate checks the field values on BusinessId with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BusinessId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BusinessId with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BusinessIdMultiError, or
// nil if none found.
func (m *BusinessId) ValidateAll() error {
	return m.validate(true)
}

func (m *BusinessId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBusinessId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessIdValidationError{
					field:  "BusinessId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessIdValidationError{
					field:  "BusinessId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBusinessId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessIdValidationError{
				field:  "BusinessId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BusinessIdMultiError(errors)
	}
	return nil
}

// BusinessIdMultiError is an error wrapping multiple validation errors
// returned by BusinessId.ValidateAll() if the designated constraints aren't met.
type BusinessIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BusinessIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BusinessIdMultiError) AllErrors() []error { return m }

// BusinessIdValidationError is the validation error returned by
// BusinessId.Validate if the designated constraints aren't met.
type BusinessIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BusinessIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BusinessIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BusinessIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BusinessIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BusinessIdValidationError) ErrorName() string { return "BusinessIdValidationError" }

// Error satisfies the builtin error interface
func (e BusinessIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBusinessId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BusinessIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BusinessIdValidationError{}
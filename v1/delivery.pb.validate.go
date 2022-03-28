// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: delivery.proto

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

// Validate checks the field values on Delivery with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Delivery) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Delivery with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeliveryMultiError, or nil
// if none found.
func (m *Delivery) ValidateAll() error {
	return m.validate(true)
}

func (m *Delivery) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeliveryValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeliveryValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeliveryValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LocationType

	// no validation rules for Date

	if len(errors) > 0 {
		return DeliveryMultiError(errors)
	}
	return nil
}

// DeliveryMultiError is an error wrapping multiple validation errors returned
// by Delivery.ValidateAll() if the designated constraints aren't met.
type DeliveryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeliveryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeliveryMultiError) AllErrors() []error { return m }

// DeliveryValidationError is the validation error returned by
// Delivery.Validate if the designated constraints aren't met.
type DeliveryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeliveryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeliveryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeliveryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeliveryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeliveryValidationError) ErrorName() string { return "DeliveryValidationError" }

// Error satisfies the builtin error interface
func (e DeliveryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDelivery.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeliveryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeliveryValidationError{}
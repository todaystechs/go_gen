// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pick_up.proto

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

// Validate checks the field values on PickUp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PickUp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PickUp with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PickUpMultiError, or nil if none found.
func (m *PickUp) ValidateAll() error {
	return m.validate(true)
}

func (m *PickUp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PickUpValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PickUpValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PickUpValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LocationType

	// no validation rules for Date

	if len(errors) > 0 {
		return PickUpMultiError(errors)
	}

	return nil
}

// PickUpMultiError is an error wrapping multiple validation errors returned by
// PickUp.ValidateAll() if the designated constraints aren't met.
type PickUpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PickUpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PickUpMultiError) AllErrors() []error { return m }

// PickUpValidationError is the validation error returned by PickUp.Validate if
// the designated constraints aren't met.
type PickUpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PickUpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PickUpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PickUpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PickUpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PickUpValidationError) ErrorName() string { return "PickUpValidationError" }

// Error satisfies the builtin error interface
func (e PickUpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPickUp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PickUpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PickUpValidationError{}

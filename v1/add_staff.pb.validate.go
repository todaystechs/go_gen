// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: add_staff.proto

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

// Validate checks the field values on AddStaff with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddStaff) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddStaff with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddStaffMultiError, or nil
// if none found.
func (m *AddStaff) ValidateAll() error {
	return m.validate(true)
}

func (m *AddStaff) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for NewStaffEmail

	// no validation rules for Password

	// no validation rules for FirstName

	// no validation rules for LastName

	if all {
		switch v := interface{}(m.GetPhoneNumber()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddStaffValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddStaffValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddStaffValidationError{
				field:  "PhoneNumber",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return AddStaffMultiError(errors)
	}

	return nil
}

// AddStaffMultiError is an error wrapping multiple validation errors returned
// by AddStaff.ValidateAll() if the designated constraints aren't met.
type AddStaffMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddStaffMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddStaffMultiError) AllErrors() []error { return m }

// AddStaffValidationError is the validation error returned by
// AddStaff.Validate if the designated constraints aren't met.
type AddStaffValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddStaffValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddStaffValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddStaffValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddStaffValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddStaffValidationError) ErrorName() string { return "AddStaffValidationError" }

// Error satisfies the builtin error interface
func (e AddStaffValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddStaff.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddStaffValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddStaffValidationError{}

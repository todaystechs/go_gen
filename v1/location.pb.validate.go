// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: location.proto

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

// Validate checks the field values on Location with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Location) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Location with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocationMultiError, or nil
// if none found.
func (m *Location) ValidateAll() error {
	return m.validate(true)
}

func (m *Location) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocationValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetContact()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "Contact",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "Contact",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContact()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocationValidationError{
				field:  "Contact",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LocationType

	if all {
		switch v := interface{}(m.GetBusinessHours()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "BusinessHours",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocationValidationError{
					field:  "BusinessHours",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBusinessHours()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocationValidationError{
				field:  "BusinessHours",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsDefaultPickUp

	// no validation rules for IsDefaultDelivery

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for LocationPk

	// no validation rules for LocationSk

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return LocationMultiError(errors)
	}
	return nil
}

// LocationMultiError is an error wrapping multiple validation errors returned
// by Location.ValidateAll() if the designated constraints aren't met.
type LocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocationMultiError) AllErrors() []error { return m }

// LocationValidationError is the validation error returned by
// Location.Validate if the designated constraints aren't met.
type LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationValidationError) ErrorName() string { return "LocationValidationError" }

// Error satisfies the builtin error interface
func (e LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationValidationError{}

// Validate checks the field values on ListsOfLocation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListsOfLocation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListsOfLocation with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListsOfLocationMultiError, or nil if none found.
func (m *ListsOfLocation) ValidateAll() error {
	return m.validate(true)
}

func (m *ListsOfLocation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetListsOfLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListsOfLocationValidationError{
					field:  "ListsOfLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListsOfLocationValidationError{
					field:  "ListsOfLocation",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetListsOfLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListsOfLocationValidationError{
				field:  "ListsOfLocation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListsOfLocationMultiError(errors)
	}
	return nil
}

// ListsOfLocationMultiError is an error wrapping multiple validation errors
// returned by ListsOfLocation.ValidateAll() if the designated constraints
// aren't met.
type ListsOfLocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListsOfLocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListsOfLocationMultiError) AllErrors() []error { return m }

// ListsOfLocationValidationError is the validation error returned by
// ListsOfLocation.Validate if the designated constraints aren't met.
type ListsOfLocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListsOfLocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListsOfLocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListsOfLocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListsOfLocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListsOfLocationValidationError) ErrorName() string { return "ListsOfLocationValidationError" }

// Error satisfies the builtin error interface
func (e ListsOfLocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListsOfLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListsOfLocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListsOfLocationValidationError{}

// Validate checks the field values on DynamoLocation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DynamoLocation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamoLocation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DynamoLocationMultiError,
// or nil if none found.
func (m *DynamoLocation) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamoLocation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for LocationPk

	// no validation rules for LocationSk

	if all {
		switch v := interface{}(m.GetLocationEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DynamoLocationValidationError{
					field:  "LocationEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DynamoLocationValidationError{
					field:  "LocationEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocationEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamoLocationValidationError{
				field:  "LocationEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DynamoLocationMultiError(errors)
	}
	return nil
}

// DynamoLocationMultiError is an error wrapping multiple validation errors
// returned by DynamoLocation.ValidateAll() if the designated constraints
// aren't met.
type DynamoLocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamoLocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamoLocationMultiError) AllErrors() []error { return m }

// DynamoLocationValidationError is the validation error returned by
// DynamoLocation.Validate if the designated constraints aren't met.
type DynamoLocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamoLocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamoLocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamoLocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamoLocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamoLocationValidationError) ErrorName() string { return "DynamoLocationValidationError" }

// Error satisfies the builtin error interface
func (e DynamoLocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamoLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamoLocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamoLocationValidationError{}

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

	// no validation rules for BusinessPk

	// no validation rules for BusinessSk

	// no validation rules for LocationPk

	// no validation rules for LocationSk

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

// Validate checks the field values on Locations with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Locations) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Locations with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocationsMultiError, or nil
// if none found.
func (m *Locations) ValidateAll() error {
	return m.validate(true)
}

func (m *Locations) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLocations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocationsValidationError{
						field:  fmt.Sprintf("Locations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocationsValidationError{
						field:  fmt.Sprintf("Locations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocationsValidationError{
					field:  fmt.Sprintf("Locations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LocationsMultiError(errors)
	}

	return nil
}

// LocationsMultiError is an error wrapping multiple validation errors returned
// by Locations.ValidateAll() if the designated constraints aren't met.
type LocationsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocationsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocationsMultiError) AllErrors() []error { return m }

// LocationsValidationError is the validation error returned by
// Locations.Validate if the designated constraints aren't met.
type LocationsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationsValidationError) ErrorName() string { return "LocationsValidationError" }

// Error satisfies the builtin error interface
func (e LocationsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocations.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationsValidationError{}

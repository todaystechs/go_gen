// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: book.proto

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

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultMultiError, or nil if none found.
func (m *Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CapacityProviderBolUrl

	// no validation rules for ShipmentIdentifier

	// no validation rules for PickupNote

	// no validation rules for PickupDateTime

	if len(errors) > 0 {
		return ResultMultiError(errors)
	}

	return nil
}

// ResultMultiError is an error wrapping multiple validation errors returned by
// Result.ValidateAll() if the designated constraints aren't met.
type ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultMultiError) AllErrors() []error { return m }

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on Booking with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Booking) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Booking with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BookingMultiError, or nil if none found.
func (m *Booking) ValidateAll() error {
	return m.validate(true)
}

func (m *Booking) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BookingId

	// no validation rules for BolUrl

	// no validation rules for InvoiceUrl

	// no validation rules for InvoiceDueDate

	// no validation rules for BusinessId

	// no validation rules for BidId

	// no validation rules for ShipmentId

	// no validation rules for SecuirityKey

	// no validation rules for PickupNumber

	// no validation rules for CarrierName

	// no validation rules for CarrierProNumber

	// no validation rules for HandelingUnitTotal

	// no validation rules for IsShipmentEdit

	// no validation rules for IsShipmentManual

	// no validation rules for ServiceType

	// no validation rules for IsTrackingEmailSend

	// no validation rules for IsTrackingApiEnabled

	// no validation rules for CustomerBolNumber

	// no validation rules for ShipperEmail

	// no validation rules for ConsigneeEmail

	// no validation rules for LogoIcon

	// no validation rules for CustomerRefNumber

	// no validation rules for CustomerPoNumber

	// no validation rules for QuoteId

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for VendorName

	if len(errors) > 0 {
		return BookingMultiError(errors)
	}

	return nil
}

// BookingMultiError is an error wrapping multiple validation errors returned
// by Booking.ValidateAll() if the designated constraints aren't met.
type BookingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingMultiError) AllErrors() []error { return m }

// BookingValidationError is the validation error returned by Booking.Validate
// if the designated constraints aren't met.
type BookingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingValidationError) ErrorName() string { return "BookingValidationError" }

// Error satisfies the builtin error interface
func (e BookingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBooking.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingValidationError{}

// Validate checks the field values on BookingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BookingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookingRequestMultiError,
// or nil if none found.
func (m *BookingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BookingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBid()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingRequestValidationError{
					field:  "Bid",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingRequestValidationError{
					field:  "Bid",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBid()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingRequestValidationError{
				field:  "Bid",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetQuote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingRequestValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingRequestValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingRequestValidationError{
				field:  "Quote",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BookingRequestMultiError(errors)
	}

	return nil
}

// BookingRequestMultiError is an error wrapping multiple validation errors
// returned by BookingRequest.ValidateAll() if the designated constraints
// aren't met.
type BookingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingRequestMultiError) AllErrors() []error { return m }

// BookingRequestValidationError is the validation error returned by
// BookingRequest.Validate if the designated constraints aren't met.
type BookingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingRequestValidationError) ErrorName() string { return "BookingRequestValidationError" }

// Error satisfies the builtin error interface
func (e BookingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingRequestValidationError{}

// Validate checks the field values on Bookings with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bookings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bookings with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookingsMultiError, or nil
// if none found.
func (m *Bookings) ValidateAll() error {
	return m.validate(true)
}

func (m *Bookings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBookings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BookingsValidationError{
						field:  fmt.Sprintf("Bookings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BookingsValidationError{
						field:  fmt.Sprintf("Bookings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BookingsValidationError{
					field:  fmt.Sprintf("Bookings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BookingsMultiError(errors)
	}

	return nil
}

// BookingsMultiError is an error wrapping multiple validation errors returned
// by Bookings.ValidateAll() if the designated constraints aren't met.
type BookingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingsMultiError) AllErrors() []error { return m }

// BookingsValidationError is the validation error returned by
// Bookings.Validate if the designated constraints aren't met.
type BookingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingsValidationError) ErrorName() string { return "BookingsValidationError" }

// Error satisfies the builtin error interface
func (e BookingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingsValidationError{}

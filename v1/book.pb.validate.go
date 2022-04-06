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

// Validate checks the field values on BookingData with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BookingData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookingData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookingDataMultiError, or
// nil if none found.
func (m *BookingData) ValidateAll() error {
	return m.validate(true)
}

func (m *BookingData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for Token

	if len(errors) > 0 {
		return BookingDataMultiError(errors)
	}
	return nil
}

// BookingDataMultiError is an error wrapping multiple validation errors
// returned by BookingData.ValidateAll() if the designated constraints aren't met.
type BookingDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingDataMultiError) AllErrors() []error { return m }

// BookingDataValidationError is the validation error returned by
// BookingData.Validate if the designated constraints aren't met.
type BookingDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingDataValidationError) ErrorName() string { return "BookingDataValidationError" }

// Error satisfies the builtin error interface
func (e BookingDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookingData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingDataValidationError{}

// Validate checks the field values on BookingResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BookingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BookingResponseMultiError, or nil if none found.
func (m *BookingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BookingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for BookingId

	// no validation rules for BolUrl

	// no validation rules for InvoiceUrl

	// no validation rules for InvoiceDueDate

	// no validation rules for PickUpStart

	// no validation rules for PickUpEnd

	if all {
		switch v := interface{}(m.GetBookedQuote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "BookedQuote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "BookedQuote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBookedQuote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "BookedQuote",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BookingResponseMultiError(errors)
	}
	return nil
}

// BookingResponseMultiError is an error wrapping multiple validation errors
// returned by BookingResponse.ValidateAll() if the designated constraints
// aren't met.
type BookingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingResponseMultiError) AllErrors() []error { return m }

// BookingResponseValidationError is the validation error returned by
// BookingResponse.Validate if the designated constraints aren't met.
type BookingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingResponseValidationError) ErrorName() string { return "BookingResponseValidationError" }

// Error satisfies the builtin error interface
func (e BookingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingResponseValidationError{}

// Validate checks the field values on ListOfBooking with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListOfBooking) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListOfBooking with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListOfBookingMultiError, or
// nil if none found.
func (m *ListOfBooking) ValidateAll() error {
	return m.validate(true)
}

func (m *ListOfBooking) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBookingHistory() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListOfBookingValidationError{
						field:  fmt.Sprintf("BookingHistory[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListOfBookingValidationError{
						field:  fmt.Sprintf("BookingHistory[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListOfBookingValidationError{
					field:  fmt.Sprintf("BookingHistory[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListOfBookingMultiError(errors)
	}
	return nil
}

// ListOfBookingMultiError is an error wrapping multiple validation errors
// returned by ListOfBooking.ValidateAll() if the designated constraints
// aren't met.
type ListOfBookingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListOfBookingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListOfBookingMultiError) AllErrors() []error { return m }

// ListOfBookingValidationError is the validation error returned by
// ListOfBooking.Validate if the designated constraints aren't met.
type ListOfBookingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListOfBookingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListOfBookingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListOfBookingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListOfBookingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListOfBookingValidationError) ErrorName() string { return "ListOfBookingValidationError" }

// Error satisfies the builtin error interface
func (e ListOfBookingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListOfBooking.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListOfBookingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListOfBookingValidationError{}

// Validate checks the field values on FetchBookingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FetchBookingsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchBookingsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FetchBookingsRequestMultiError, or nil if none found.
func (m *FetchBookingsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchBookingsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pk

	// no validation rules for StartFrom

	// no validation rules for EndOn

	if len(errors) > 0 {
		return FetchBookingsRequestMultiError(errors)
	}
	return nil
}

// FetchBookingsRequestMultiError is an error wrapping multiple validation
// errors returned by FetchBookingsRequest.ValidateAll() if the designated
// constraints aren't met.
type FetchBookingsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchBookingsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchBookingsRequestMultiError) AllErrors() []error { return m }

// FetchBookingsRequestValidationError is the validation error returned by
// FetchBookingsRequest.Validate if the designated constraints aren't met.
type FetchBookingsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchBookingsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchBookingsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchBookingsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchBookingsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchBookingsRequestValidationError) ErrorName() string {
	return "FetchBookingsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FetchBookingsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchBookingsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchBookingsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchBookingsRequestValidationError{}

// Validate checks the field values on BookEntity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BookEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookEntity with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookEntityMultiError, or
// nil if none found.
func (m *BookEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *BookEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuoteEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookEntityValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookEntityValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuoteEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookEntityValidationError{
				field:  "QuoteEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRef()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookEntityValidationError{
					field:  "Ref",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookEntityValidationError{
					field:  "Ref",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRef()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookEntityValidationError{
				field:  "Ref",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BookEntityMultiError(errors)
	}
	return nil
}

// BookEntityMultiError is an error wrapping multiple validation errors
// returned by BookEntity.ValidateAll() if the designated constraints aren't met.
type BookEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookEntityMultiError) AllErrors() []error { return m }

// BookEntityValidationError is the validation error returned by
// BookEntity.Validate if the designated constraints aren't met.
type BookEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookEntityValidationError) ErrorName() string { return "BookEntityValidationError" }

// Error satisfies the builtin error interface
func (e BookEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookEntityValidationError{}

// Validate checks the field values on DynamoBookEntity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DynamoBookEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamoBookEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DynamoBookEntityMultiError, or nil if none found.
func (m *DynamoBookEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamoBookEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for BookPk

	// no validation rules for BookSk

	if all {
		switch v := interface{}(m.GetBookEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DynamoBookEntityValidationError{
					field:  "BookEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DynamoBookEntityValidationError{
					field:  "BookEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBookEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamoBookEntityValidationError{
				field:  "BookEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DynamoBookEntityMultiError(errors)
	}
	return nil
}

// DynamoBookEntityMultiError is an error wrapping multiple validation errors
// returned by DynamoBookEntity.ValidateAll() if the designated constraints
// aren't met.
type DynamoBookEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamoBookEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamoBookEntityMultiError) AllErrors() []error { return m }

// DynamoBookEntityValidationError is the validation error returned by
// DynamoBookEntity.Validate if the designated constraints aren't met.
type DynamoBookEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamoBookEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamoBookEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamoBookEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamoBookEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamoBookEntityValidationError) ErrorName() string { return "DynamoBookEntityValidationError" }

// Error satisfies the builtin error interface
func (e DynamoBookEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamoBookEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamoBookEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamoBookEntityValidationError{}

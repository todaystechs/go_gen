// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: quote.proto

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

// Validate checks the field values on Quote with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Quote) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Quote with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in QuoteMultiError, or nil if none found.
func (m *Quote) ValidateAll() error {
	return m.validate(true)
}

func (m *Quote) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for RequesterId

	// no validation rules for BusinessPk

	// no validation rules for BusinessSk

	// no validation rules for Mode

	// no validation rules for LiablePartyId

	if all {
		switch v := interface{}(m.GetPickup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPickup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteValidationError{
				field:  "Pickup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDelivery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDelivery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteValidationError{
				field:  "Delivery",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCommodities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteValidationError{
					field:  fmt.Sprintf("Commodities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for PickupDate

	// no validation rules for DisplayDate

	// no validation rules for DeliveryDate

	for idx, item := range m.GetBids() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteValidationError{
					field:  fmt.Sprintf("Bids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	{
		sorted_keys := make([]string, len(m.GetVendorBids()))
		i := 0
		for key := range m.GetVendorBids() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetVendorBids()[key]
			_ = val

			// no validation rules for VendorBids[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, QuoteValidationError{
							field:  fmt.Sprintf("VendorBids[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, QuoteValidationError{
							field:  fmt.Sprintf("VendorBids[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return QuoteValidationError{
						field:  fmt.Sprintf("VendorBids[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	// no validation rules for TotalItems

	// no validation rules for TotalWeight

	if len(errors) > 0 {
		return QuoteMultiError(errors)
	}

	return nil
}

// QuoteMultiError is an error wrapping multiple validation errors returned by
// Quote.ValidateAll() if the designated constraints aren't met.
type QuoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteMultiError) AllErrors() []error { return m }

// QuoteValidationError is the validation error returned by Quote.Validate if
// the designated constraints aren't met.
type QuoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteValidationError) ErrorName() string { return "QuoteValidationError" }

// Error satisfies the builtin error interface
func (e QuoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteValidationError{}

// Validate checks the field values on Quotes with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Quotes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Quotes with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in QuotesMultiError, or nil if none found.
func (m *Quotes) ValidateAll() error {
	return m.validate(true)
}

func (m *Quotes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuotesValidationError{
						field:  fmt.Sprintf("Quotes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuotesValidationError{
						field:  fmt.Sprintf("Quotes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuotesValidationError{
					field:  fmt.Sprintf("Quotes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuotesMultiError(errors)
	}

	return nil
}

// QuotesMultiError is an error wrapping multiple validation errors returned by
// Quotes.ValidateAll() if the designated constraints aren't met.
type QuotesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesMultiError) AllErrors() []error { return m }

// QuotesValidationError is the validation error returned by Quotes.Validate if
// the designated constraints aren't met.
type QuotesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesValidationError) ErrorName() string { return "QuotesValidationError" }

// Error satisfies the builtin error interface
func (e QuotesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesValidationError{}

// Validate checks the field values on QuoteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteResponseMultiError, or
// nil if none found.
func (m *QuoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteResponseValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteResponseValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteResponseValidationError{
				field:  "Quote",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetBids() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteResponseValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteResponseValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteResponseValidationError{
					field:  fmt.Sprintf("Bids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteResponseMultiError(errors)
	}

	return nil
}

// QuoteResponseMultiError is an error wrapping multiple validation errors
// returned by QuoteResponse.ValidateAll() if the designated constraints
// aren't met.
type QuoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteResponseMultiError) AllErrors() []error { return m }

// QuoteResponseValidationError is the validation error returned by
// QuoteResponse.Validate if the designated constraints aren't met.
type QuoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteResponseValidationError) ErrorName() string { return "QuoteResponseValidationError" }

// Error satisfies the builtin error interface
func (e QuoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteResponseValidationError{}

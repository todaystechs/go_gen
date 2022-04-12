// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bid.proto

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

// Validate checks the field values on Bid with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Bid) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bid with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BidMultiError, or nil if none found.
func (m *Bid) ValidateAll() error {
	return m.validate(true)
}

func (m *Bid) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for BidId

	// no validation rules for Carrier

	if all {
		switch v := interface{}(m.GetAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BidValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BidValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BidValidationError{
				field:  "Amount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TransitTime

	// no validation rules for Guranteed

	// no validation rules for BusinessId

	// no validation rules for CompanyImage

	if len(errors) > 0 {
		return BidMultiError(errors)
	}

	return nil
}

// BidMultiError is an error wrapping multiple validation errors returned by
// Bid.ValidateAll() if the designated constraints aren't met.
type BidMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BidMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BidMultiError) AllErrors() []error { return m }

// BidValidationError is the validation error returned by Bid.Validate if the
// designated constraints aren't met.
type BidValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BidValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BidValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BidValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BidValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BidValidationError) ErrorName() string { return "BidValidationError" }

// Error satisfies the builtin error interface
func (e BidValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBid.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BidValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BidValidationError{}

// Validate checks the field values on Bids with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Bids) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bids with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BidsMultiError, or nil if none found.
func (m *Bids) ValidateAll() error {
	return m.validate(true)
}

func (m *Bids) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBids() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BidsValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BidsValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BidsValidationError{
					field:  fmt.Sprintf("Bids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BidsMultiError(errors)
	}

	return nil
}

// BidsMultiError is an error wrapping multiple validation errors returned by
// Bids.ValidateAll() if the designated constraints aren't met.
type BidsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BidsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BidsMultiError) AllErrors() []error { return m }

// BidsValidationError is the validation error returned by Bids.Validate if the
// designated constraints aren't met.
type BidsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BidsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BidsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BidsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BidsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BidsValidationError) ErrorName() string { return "BidsValidationError" }

// Error satisfies the builtin error interface
func (e BidsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBids.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BidsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BidsValidationError{}

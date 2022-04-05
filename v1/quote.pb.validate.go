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

// Validate checks the field values on QuoteEntity with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteEntity with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteEntityMultiError, or
// nil if none found.
func (m *QuoteEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for RequesterId

	// no validation rules for BusinessId

	// no validation rules for Mode

	// no validation rules for LiablePartyId

	if all {
		switch v := interface{}(m.GetShippingDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteEntityValidationError{
					field:  "ShippingDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteEntityValidationError{
					field:  "ShippingDetail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShippingDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteEntityValidationError{
				field:  "ShippingDetail",
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
					errors = append(errors, QuoteEntityValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteEntityValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteEntityValidationError{
					field:  fmt.Sprintf("Commodities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteEntityMultiError(errors)
	}
	return nil
}

// QuoteEntityMultiError is an error wrapping multiple validation errors
// returned by QuoteEntity.ValidateAll() if the designated constraints aren't met.
type QuoteEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteEntityMultiError) AllErrors() []error { return m }

// QuoteEntityValidationError is the validation error returned by
// QuoteEntity.Validate if the designated constraints aren't met.
type QuoteEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteEntityValidationError) ErrorName() string { return "QuoteEntityValidationError" }

// Error satisfies the builtin error interface
func (e QuoteEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteEntityValidationError{}

// Validate checks the field values on DynamoQuoteEntity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DynamoQuoteEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamoQuoteEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DynamoQuoteEntityMultiError, or nil if none found.
func (m *DynamoQuoteEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamoQuoteEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for QuotePk

	// no validation rules for QuoteSk

	if all {
		switch v := interface{}(m.GetQuoteEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DynamoQuoteEntityValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DynamoQuoteEntityValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuoteEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamoQuoteEntityValidationError{
				field:  "QuoteEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DynamoQuoteEntityMultiError(errors)
	}
	return nil
}

// DynamoQuoteEntityMultiError is an error wrapping multiple validation errors
// returned by DynamoQuoteEntity.ValidateAll() if the designated constraints
// aren't met.
type DynamoQuoteEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamoQuoteEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamoQuoteEntityMultiError) AllErrors() []error { return m }

// DynamoQuoteEntityValidationError is the validation error returned by
// DynamoQuoteEntity.Validate if the designated constraints aren't met.
type DynamoQuoteEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamoQuoteEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamoQuoteEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamoQuoteEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamoQuoteEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamoQuoteEntityValidationError) ErrorName() string {
	return "DynamoQuoteEntityValidationError"
}

// Error satisfies the builtin error interface
func (e DynamoQuoteEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamoQuoteEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamoQuoteEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamoQuoteEntityValidationError{}

// Validate checks the field values on BidEntity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BidEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BidEntity with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BidEntityMultiError, or nil
// if none found.
func (m *BidEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *BidEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for Sk

	if all {
		switch v := interface{}(m.GetCarrier()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BidEntityValidationError{
					field:  "Carrier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BidEntityValidationError{
					field:  "Carrier",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCarrier()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BidEntityValidationError{
				field:  "Carrier",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAmount()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BidEntityValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BidEntityValidationError{
					field:  "Amount",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAmount()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BidEntityValidationError{
				field:  "Amount",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TransitTime

	// no validation rules for Guranteed

	if len(errors) > 0 {
		return BidEntityMultiError(errors)
	}
	return nil
}

// BidEntityMultiError is an error wrapping multiple validation errors returned
// by BidEntity.ValidateAll() if the designated constraints aren't met.
type BidEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BidEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BidEntityMultiError) AllErrors() []error { return m }

// BidEntityValidationError is the validation error returned by
// BidEntity.Validate if the designated constraints aren't met.
type BidEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BidEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BidEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BidEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BidEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BidEntityValidationError) ErrorName() string { return "BidEntityValidationError" }

// Error satisfies the builtin error interface
func (e BidEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBidEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BidEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BidEntityValidationError{}

// Validate checks the field values on DynamoQuoteBidEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DynamoQuoteBidEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DynamoQuoteBidEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DynamoQuoteBidEntityMultiError, or nil if none found.
func (m *DynamoQuoteBidEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *DynamoQuoteBidEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pk

	// no validation rules for Sk

	// no validation rules for BidPk

	// no validation rules for BidSk

	if all {
		switch v := interface{}(m.GetBidEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DynamoQuoteBidEntityValidationError{
					field:  "BidEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DynamoQuoteBidEntityValidationError{
					field:  "BidEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBidEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DynamoQuoteBidEntityValidationError{
				field:  "BidEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DynamoQuoteBidEntityMultiError(errors)
	}
	return nil
}

// DynamoQuoteBidEntityMultiError is an error wrapping multiple validation
// errors returned by DynamoQuoteBidEntity.ValidateAll() if the designated
// constraints aren't met.
type DynamoQuoteBidEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DynamoQuoteBidEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DynamoQuoteBidEntityMultiError) AllErrors() []error { return m }

// DynamoQuoteBidEntityValidationError is the validation error returned by
// DynamoQuoteBidEntity.Validate if the designated constraints aren't met.
type DynamoQuoteBidEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamoQuoteBidEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamoQuoteBidEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamoQuoteBidEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamoQuoteBidEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamoQuoteBidEntityValidationError) ErrorName() string {
	return "DynamoQuoteBidEntityValidationError"
}

// Error satisfies the builtin error interface
func (e DynamoQuoteBidEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamoQuoteBidEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamoQuoteBidEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamoQuoteBidEntityValidationError{}

// Validate checks the field values on QuoteBids with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteBids) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteBids with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteBidsMultiError, or nil
// if none found.
func (m *QuoteBids) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteBids) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuoteEntity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteBidsValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteBidsValidationError{
					field:  "QuoteEntity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuoteEntity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteBidsValidationError{
				field:  "QuoteEntity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetQuoteBids() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteBidsValidationError{
						field:  fmt.Sprintf("QuoteBids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteBidsValidationError{
						field:  fmt.Sprintf("QuoteBids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteBidsValidationError{
					field:  fmt.Sprintf("QuoteBids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteBidsMultiError(errors)
	}
	return nil
}

// QuoteBidsMultiError is an error wrapping multiple validation errors returned
// by QuoteBids.ValidateAll() if the designated constraints aren't met.
type QuoteBidsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteBidsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteBidsMultiError) AllErrors() []error { return m }

// QuoteBidsValidationError is the validation error returned by
// QuoteBids.Validate if the designated constraints aren't met.
type QuoteBidsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteBidsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteBidsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteBidsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteBidsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteBidsValidationError) ErrorName() string { return "QuoteBidsValidationError" }

// Error satisfies the builtin error interface
func (e QuoteBidsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteBids.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteBidsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteBidsValidationError{}

// Validate checks the field values on FetchQuotes with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FetchQuotes) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchQuotes with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FetchQuotesMultiError, or
// nil if none found.
func (m *FetchQuotes) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchQuotes) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StartFrom

	// no validation rules for EndOn

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return FetchQuotesMultiError(errors)
	}
	return nil
}

// FetchQuotesMultiError is an error wrapping multiple validation errors
// returned by FetchQuotes.ValidateAll() if the designated constraints aren't met.
type FetchQuotesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchQuotesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchQuotesMultiError) AllErrors() []error { return m }

// FetchQuotesValidationError is the validation error returned by
// FetchQuotes.Validate if the designated constraints aren't met.
type FetchQuotesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchQuotesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchQuotesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchQuotesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchQuotesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchQuotesValidationError) ErrorName() string { return "FetchQuotesValidationError" }

// Error satisfies the builtin error interface
func (e FetchQuotesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchQuotes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchQuotesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchQuotesValidationError{}

// Validate checks the field values on QuoteEntities with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteEntities) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteEntities with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteEntitiesMultiError, or
// nil if none found.
func (m *QuoteEntities) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteEntities) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuoteEntities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteEntitiesValidationError{
						field:  fmt.Sprintf("QuoteEntities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteEntitiesValidationError{
						field:  fmt.Sprintf("QuoteEntities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteEntitiesValidationError{
					field:  fmt.Sprintf("QuoteEntities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteEntitiesMultiError(errors)
	}
	return nil
}

// QuoteEntitiesMultiError is an error wrapping multiple validation errors
// returned by QuoteEntities.ValidateAll() if the designated constraints
// aren't met.
type QuoteEntitiesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteEntitiesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteEntitiesMultiError) AllErrors() []error { return m }

// QuoteEntitiesValidationError is the validation error returned by
// QuoteEntities.Validate if the designated constraints aren't met.
type QuoteEntitiesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteEntitiesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteEntitiesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteEntitiesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteEntitiesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteEntitiesValidationError) ErrorName() string { return "QuoteEntitiesValidationError" }

// Error satisfies the builtin error interface
func (e QuoteEntitiesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteEntities.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteEntitiesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteEntitiesValidationError{}

// Validate checks the field values on DeleteQuote with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteQuote) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteQuote with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteQuoteMultiError, or
// nil if none found.
func (m *DeleteQuote) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteQuote) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	if len(errors) > 0 {
		return DeleteQuoteMultiError(errors)
	}
	return nil
}

// DeleteQuoteMultiError is an error wrapping multiple validation errors
// returned by DeleteQuote.ValidateAll() if the designated constraints aren't met.
type DeleteQuoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteQuoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteQuoteMultiError) AllErrors() []error { return m }

// DeleteQuoteValidationError is the validation error returned by
// DeleteQuote.Validate if the designated constraints aren't met.
type DeleteQuoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteQuoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteQuoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteQuoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteQuoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteQuoteValidationError) ErrorName() string { return "DeleteQuoteValidationError" }

// Error satisfies the builtin error interface
func (e DeleteQuoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteQuote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteQuoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteQuoteValidationError{}

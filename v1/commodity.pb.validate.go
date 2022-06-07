// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: commodity.proto

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

// Validate checks the field values on Commodity with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Commodity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Commodity with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommodityMultiError, or nil
// if none found.
func (m *Commodity) ValidateAll() error {
	return m.validate(true)
}

func (m *Commodity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Density

	// no validation rules for Length

	// no validation rules for Width

	// no validation rules for Height

	// no validation rules for Weight

	// no validation rules for DimensionUom

	// no validation rules for WeightUom

	// no validation rules for DimensionDisplay

	// no validation rules for PackageType

	// no validation rules for Quantity

	// no validation rules for FreightClass

	// no validation rules for Stackable

	// no validation rules for ProtectFromFreeze

	// no validation rules for SortAndSegregate

	// no validation rules for Guaranteed

	// no validation rules for Hazardous

	// no validation rules for Instructions

	if len(errors) > 0 {
		return CommodityMultiError(errors)
	}

	return nil
}

// CommodityMultiError is an error wrapping multiple validation errors returned
// by Commodity.ValidateAll() if the designated constraints aren't met.
type CommodityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommodityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommodityMultiError) AllErrors() []error { return m }

// CommodityValidationError is the validation error returned by
// Commodity.Validate if the designated constraints aren't met.
type CommodityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommodityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommodityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommodityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommodityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommodityValidationError) ErrorName() string { return "CommodityValidationError" }

// Error satisfies the builtin error interface
func (e CommodityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommodity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommodityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommodityValidationError{}

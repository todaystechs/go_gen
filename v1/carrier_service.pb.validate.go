// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: carrier_service.proto

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

// Validate checks the field values on Ping with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Ping) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ping with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PingMultiError, or nil if none found.
func (m *Ping) ValidateAll() error {
	return m.validate(true)
}

func (m *Ping) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Hi

	if len(errors) > 0 {
		return PingMultiError(errors)
	}

	return nil
}

// PingMultiError is an error wrapping multiple validation errors returned by
// Ping.ValidateAll() if the designated constraints aren't met.
type PingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingMultiError) AllErrors() []error { return m }

// PingValidationError is the validation error returned by Ping.Validate if the
// designated constraints aren't met.
type PingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingValidationError) ErrorName() string { return "PingValidationError" }

// Error satisfies the builtin error interface
func (e PingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingValidationError{}

// Validate checks the field values on Id with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Id) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Id with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdMultiError, or nil if none found.
func (m *Id) ValidateAll() error {
	return m.validate(true)
}

func (m *Id) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return IdMultiError(errors)
	}

	return nil
}

// IdMultiError is an error wrapping multiple validation errors returned by
// Id.ValidateAll() if the designated constraints aren't met.
type IdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdMultiError) AllErrors() []error { return m }

// IdValidationError is the validation error returned by Id.Validate if the
// designated constraints aren't met.
type IdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdValidationError) ErrorName() string { return "IdValidationError" }

// Error satisfies the builtin error interface
func (e IdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdValidationError{}

// Validate checks the field values on Ids with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Ids) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ids with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IdsMultiError, or nil if none found.
func (m *Ids) ValidateAll() error {
	return m.validate(true)
}

func (m *Ids) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return IdsMultiError(errors)
	}

	return nil
}

// IdsMultiError is an error wrapping multiple validation errors returned by
// Ids.ValidateAll() if the designated constraints aren't met.
type IdsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdsMultiError) AllErrors() []error { return m }

// IdsValidationError is the validation error returned by Ids.Validate if the
// designated constraints aren't met.
type IdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdsValidationError) ErrorName() string { return "IdsValidationError" }

// Error satisfies the builtin error interface
func (e IdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdsValidationError{}

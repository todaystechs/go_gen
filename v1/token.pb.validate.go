// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: token.proto

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

// Validate checks the field values on RefreshToken with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RefreshToken) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshToken with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RefreshTokenMultiError, or
// nil if none found.
func (m *RefreshToken) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshToken) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return RefreshTokenMultiError(errors)
	}

	return nil
}

// RefreshTokenMultiError is an error wrapping multiple validation errors
// returned by RefreshToken.ValidateAll() if the designated constraints aren't met.
type RefreshTokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenMultiError) AllErrors() []error { return m }

// RefreshTokenValidationError is the validation error returned by
// RefreshToken.Validate if the designated constraints aren't met.
type RefreshTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenValidationError) ErrorName() string { return "RefreshTokenValidationError" }

// Error satisfies the builtin error interface
func (e RefreshTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenValidationError{}

// Validate checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Token) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Token with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TokenMultiError, or nil if none found.
func (m *Token) ValidateAll() error {
	return m.validate(true)
}

func (m *Token) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return TokenMultiError(errors)
	}

	return nil
}

// TokenMultiError is an error wrapping multiple validation errors returned by
// Token.ValidateAll() if the designated constraints aren't met.
type TokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TokenMultiError) AllErrors() []error { return m }

// TokenValidationError is the validation error returned by Token.Validate if
// the designated constraints aren't met.
type TokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenValidationError) ErrorName() string { return "TokenValidationError" }

// Error satisfies the builtin error interface
func (e TokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenValidationError{}
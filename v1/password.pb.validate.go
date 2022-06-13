// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: password.proto

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

// Validate checks the field values on ForgotPassword with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForgotPassword) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForgotPassword with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForgotPasswordMultiError,
// or nil if none found.
func (m *ForgotPassword) ValidateAll() error {
	return m.validate(true)
}

func (m *ForgotPassword) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	if len(errors) > 0 {
		return ForgotPasswordMultiError(errors)
	}

	return nil
}

// ForgotPasswordMultiError is an error wrapping multiple validation errors
// returned by ForgotPassword.ValidateAll() if the designated constraints
// aren't met.
type ForgotPasswordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForgotPasswordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForgotPasswordMultiError) AllErrors() []error { return m }

// ForgotPasswordValidationError is the validation error returned by
// ForgotPassword.Validate if the designated constraints aren't met.
type ForgotPasswordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForgotPasswordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForgotPasswordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForgotPasswordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForgotPasswordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForgotPasswordValidationError) ErrorName() string { return "ForgotPasswordValidationError" }

// Error satisfies the builtin error interface
func (e ForgotPasswordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForgotPassword.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForgotPasswordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForgotPasswordValidationError{}

// Validate checks the field values on ResetPasswordToken with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ResetPasswordToken) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetPasswordToken with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ResetPasswordTokenMultiError, or nil if none found.
func (m *ResetPasswordToken) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetPasswordToken) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IssuedTo

	// no validation rules for IssuedOn

	// no validation rules for ExpiresOn

	// no validation rules for Token

	// no validation rules for BusinessId

	if len(errors) > 0 {
		return ResetPasswordTokenMultiError(errors)
	}

	return nil
}

// ResetPasswordTokenMultiError is an error wrapping multiple validation errors
// returned by ResetPasswordToken.ValidateAll() if the designated constraints
// aren't met.
type ResetPasswordTokenMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetPasswordTokenMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetPasswordTokenMultiError) AllErrors() []error { return m }

// ResetPasswordTokenValidationError is the validation error returned by
// ResetPasswordToken.Validate if the designated constraints aren't met.
type ResetPasswordTokenValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordTokenValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordTokenValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordTokenValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordTokenValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordTokenValidationError) ErrorName() string {
	return "ResetPasswordTokenValidationError"
}

// Error satisfies the builtin error interface
func (e ResetPasswordTokenValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPasswordToken.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordTokenValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordTokenValidationError{}

// Validate checks the field values on ResetPassword with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResetPassword) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResetPassword with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResetPasswordMultiError, or
// nil if none found.
func (m *ResetPassword) ValidateAll() error {
	return m.validate(true)
}

func (m *ResetPassword) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for NewPassword

	// no validation rules for ConfirmPassword

	// no validation rules for Email

	if len(errors) > 0 {
		return ResetPasswordMultiError(errors)
	}

	return nil
}

// ResetPasswordMultiError is an error wrapping multiple validation errors
// returned by ResetPassword.ValidateAll() if the designated constraints
// aren't met.
type ResetPasswordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResetPasswordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResetPasswordMultiError) AllErrors() []error { return m }

// ResetPasswordValidationError is the validation error returned by
// ResetPassword.Validate if the designated constraints aren't met.
type ResetPasswordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResetPasswordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResetPasswordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResetPasswordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResetPasswordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResetPasswordValidationError) ErrorName() string { return "ResetPasswordValidationError" }

// Error satisfies the builtin error interface
func (e ResetPasswordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResetPassword.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResetPasswordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResetPasswordValidationError{}

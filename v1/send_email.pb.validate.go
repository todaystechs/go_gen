// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: send_email.proto

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

// Validate checks the field values on SendEmail with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendEmail) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendEmail with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendEmailMultiError, or nil
// if none found.
func (m *SendEmail) ValidateAll() error {
	return m.validate(true)
}

func (m *SendEmail) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EmailSubject

	// no validation rules for ReceiverEmailAddress

	// no validation rules for ReceiverName

	// no validation rules for EmailPurpose

	// no validation rules for Token

	// no validation rules for Success

	if len(errors) > 0 {
		return SendEmailMultiError(errors)
	}
	return nil
}

// SendEmailMultiError is an error wrapping multiple validation errors returned
// by SendEmail.ValidateAll() if the designated constraints aren't met.
type SendEmailMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendEmailMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendEmailMultiError) AllErrors() []error { return m }

// SendEmailValidationError is the validation error returned by
// SendEmail.Validate if the designated constraints aren't met.
type SendEmailValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendEmailValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendEmailValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendEmailValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendEmailValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendEmailValidationError) ErrorName() string { return "SendEmailValidationError" }

// Error satisfies the builtin error interface
func (e SendEmailValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendEmail.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendEmailValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendEmailValidationError{}

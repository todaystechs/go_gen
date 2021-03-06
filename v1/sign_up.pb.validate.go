// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: sign_up.proto

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

// Validate checks the field values on SignUp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignUp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignUp with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SignUpMultiError, or nil if none found.
func (m *SignUp) ValidateAll() error {
	return m.validate(true)
}

func (m *SignUp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FirstName

	// no validation rules for MiddleName

	// no validation rules for LastName

	// no validation rules for Email

	// no validation rules for Password

	// no validation rules for PhoneNumber

	// no validation rules for TermsAggrement

	// no validation rules for CompanyName

	if len(errors) > 0 {
		return SignUpMultiError(errors)
	}

	return nil
}

// SignUpMultiError is an error wrapping multiple validation errors returned by
// SignUp.ValidateAll() if the designated constraints aren't met.
type SignUpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignUpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignUpMultiError) AllErrors() []error { return m }

// SignUpValidationError is the validation error returned by SignUp.Validate if
// the designated constraints aren't met.
type SignUpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignUpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignUpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignUpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignUpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignUpValidationError) ErrorName() string { return "SignUpValidationError" }

// Error satisfies the builtin error interface
func (e SignUpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignUp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignUpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignUpValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: console/v1/example.proto

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

// Validate checks the field values on GetExampleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetExampleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExampleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExampleRequestMultiError, or nil if none found.
func (m *GetExampleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExampleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetExampleRequestMultiError(errors)
	}

	return nil
}

// GetExampleRequestMultiError is an error wrapping multiple validation errors
// returned by GetExampleRequest.ValidateAll() if the designated constraints
// aren't met.
type GetExampleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExampleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExampleRequestMultiError) AllErrors() []error { return m }

// GetExampleRequestValidationError is the validation error returned by
// GetExampleRequest.Validate if the designated constraints aren't met.
type GetExampleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExampleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExampleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExampleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExampleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExampleRequestValidationError) ErrorName() string {
	return "GetExampleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExampleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExampleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExampleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExampleRequestValidationError{}

// Validate checks the field values on GetExampleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExampleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExampleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExampleResponseMultiError, or nil if none found.
func (m *GetExampleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExampleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return GetExampleResponseMultiError(errors)
	}

	return nil
}

// GetExampleResponseMultiError is an error wrapping multiple validation errors
// returned by GetExampleResponse.ValidateAll() if the designated constraints
// aren't met.
type GetExampleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExampleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExampleResponseMultiError) AllErrors() []error { return m }

// GetExampleResponseValidationError is the validation error returned by
// GetExampleResponse.Validate if the designated constraints aren't met.
type GetExampleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExampleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExampleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExampleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExampleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExampleResponseValidationError) ErrorName() string {
	return "GetExampleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExampleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExampleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExampleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExampleResponseValidationError{}

// Validate checks the field values on PostExampleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PostExampleRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PostExampleRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PostExampleRequestMultiError, or nil if none found.
func (m *PostExampleRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PostExampleRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return PostExampleRequestMultiError(errors)
	}

	return nil
}

// PostExampleRequestMultiError is an error wrapping multiple validation errors
// returned by PostExampleRequest.ValidateAll() if the designated constraints
// aren't met.
type PostExampleRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PostExampleRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PostExampleRequestMultiError) AllErrors() []error { return m }

// PostExampleRequestValidationError is the validation error returned by
// PostExampleRequest.Validate if the designated constraints aren't met.
type PostExampleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostExampleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostExampleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostExampleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostExampleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostExampleRequestValidationError) ErrorName() string {
	return "PostExampleRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PostExampleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostExampleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostExampleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostExampleRequestValidationError{}

// Validate checks the field values on PostExampleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PostExampleResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PostExampleResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PostExampleResponseMultiError, or nil if none found.
func (m *PostExampleResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PostExampleResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return PostExampleResponseMultiError(errors)
	}

	return nil
}

// PostExampleResponseMultiError is an error wrapping multiple validation
// errors returned by PostExampleResponse.ValidateAll() if the designated
// constraints aren't met.
type PostExampleResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PostExampleResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PostExampleResponseMultiError) AllErrors() []error { return m }

// PostExampleResponseValidationError is the validation error returned by
// PostExampleResponse.Validate if the designated constraints aren't met.
type PostExampleResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostExampleResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostExampleResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostExampleResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostExampleResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostExampleResponseValidationError) ErrorName() string {
	return "PostExampleResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PostExampleResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostExampleResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostExampleResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostExampleResponseValidationError{}

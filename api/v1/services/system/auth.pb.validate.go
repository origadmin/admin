// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/auth.proto

package system

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

// Validate checks the field values on ListResourcesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourcesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourcesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourcesRequestMultiError, or nil if none found.
func (m *ListResourcesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourcesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageSize

	// no validation rules for PageToken

	// no validation rules for Current

	// no validation rules for NoPaging

	if len(errors) > 0 {
		return ListResourcesRequestMultiError(errors)
	}

	return nil
}

// ListResourcesRequestMultiError is an error wrapping multiple validation
// errors returned by ListResourcesRequest.ValidateAll() if the designated
// constraints aren't met.
type ListResourcesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourcesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourcesRequestMultiError) AllErrors() []error { return m }

// ListResourcesRequestValidationError is the validation error returned by
// ListResourcesRequest.Validate if the designated constraints aren't met.
type ListResourcesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourcesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourcesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourcesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourcesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourcesRequestValidationError) ErrorName() string {
	return "ListResourcesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourcesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourcesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourcesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourcesRequestValidationError{}

// Validate checks the field values on ListResourcesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourcesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourcesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourcesResponseMultiError, or nil if none found.
func (m *ListResourcesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourcesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResourcesResponseValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResourcesResponseValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResourcesResponseValidationError{
					field:  fmt.Sprintf("Resources[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalSize

	if len(errors) > 0 {
		return ListResourcesResponseMultiError(errors)
	}

	return nil
}

// ListResourcesResponseMultiError is an error wrapping multiple validation
// errors returned by ListResourcesResponse.ValidateAll() if the designated
// constraints aren't met.
type ListResourcesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourcesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourcesResponseMultiError) AllErrors() []error { return m }

// ListResourcesResponseValidationError is the validation error returned by
// ListResourcesResponse.Validate if the designated constraints aren't met.
type ListResourcesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourcesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourcesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourcesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourcesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourcesResponseValidationError) ErrorName() string {
	return "ListResourcesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourcesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourcesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourcesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourcesResponseValidationError{}
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/resource.proto

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

	// no validation rules for Id

	// no validation rules for Current

	// no validation rules for PageSize

	// no validation rules for PageToken

	// no validation rules for NoPaging

	// no validation rules for OnlyCount

	// no validation rules for Type

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

	// no validation rules for TotalSize

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

	// no validation rules for Current

	// no validation rules for PageSize

	// no validation rules for NextPageToken

	if m.Extra != nil {

		if all {
			switch v := interface{}(m.GetExtra()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResourcesResponseValidationError{
						field:  "Extra",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResourcesResponseValidationError{
						field:  "Extra",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetExtra()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResourcesResponseValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

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

// Validate checks the field values on GetResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceRequestMultiError, or nil if none found.
func (m *GetResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetResourceRequestMultiError(errors)
	}

	return nil
}

// GetResourceRequestMultiError is an error wrapping multiple validation errors
// returned by GetResourceRequest.ValidateAll() if the designated constraints
// aren't met.
type GetResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceRequestMultiError) AllErrors() []error { return m }

// GetResourceRequestValidationError is the validation error returned by
// GetResourceRequest.Validate if the designated constraints aren't met.
type GetResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceRequestValidationError) ErrorName() string {
	return "GetResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceRequestValidationError{}

// Validate checks the field values on GetResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetResourceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetResourceResponseMultiError, or nil if none found.
func (m *GetResourceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResourceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResourceResponseValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResourceResponseMultiError(errors)
	}

	return nil
}

// GetResourceResponseMultiError is an error wrapping multiple validation
// errors returned by GetResourceResponse.ValidateAll() if the designated
// constraints aren't met.
type GetResourceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResourceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResourceResponseMultiError) AllErrors() []error { return m }

// GetResourceResponseValidationError is the validation error returned by
// GetResourceResponse.Validate if the designated constraints aren't met.
type GetResourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResourceResponseValidationError) ErrorName() string {
	return "GetResourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetResourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResourceResponseValidationError{}

// Validate checks the field values on CreateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateResourceRequestMultiError, or nil if none found.
func (m *CreateResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Parent

	// no validation rules for ResourceId

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResourceRequestValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateResourceRequestMultiError(errors)
	}

	return nil
}

// CreateResourceRequestMultiError is an error wrapping multiple validation
// errors returned by CreateResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceRequestMultiError) AllErrors() []error { return m }

// CreateResourceRequestValidationError is the validation error returned by
// CreateResourceRequest.Validate if the designated constraints aren't met.
type CreateResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceRequestValidationError) ErrorName() string {
	return "CreateResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceRequestValidationError{}

// Validate checks the field values on CreateResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateResourceResponseMultiError, or nil if none found.
func (m *CreateResourceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateResourceResponseValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateResourceResponseMultiError(errors)
	}

	return nil
}

// CreateResourceResponseMultiError is an error wrapping multiple validation
// errors returned by CreateResourceResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateResourceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceResponseMultiError) AllErrors() []error { return m }

// CreateResourceResponseValidationError is the validation error returned by
// CreateResourceResponse.Validate if the designated constraints aren't met.
type CreateResourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceResponseValidationError) ErrorName() string {
	return "CreateResourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceResponseValidationError{}

// Validate checks the field values on UpdateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceRequestMultiError, or nil if none found.
func (m *UpdateResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateResourceRequestValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateResourceRequestValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateResourceRequestMultiError(errors)
	}

	return nil
}

// UpdateResourceRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceRequestMultiError) AllErrors() []error { return m }

// UpdateResourceRequestValidationError is the validation error returned by
// UpdateResourceRequest.Validate if the designated constraints aren't met.
type UpdateResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceRequestValidationError) ErrorName() string {
	return "UpdateResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceRequestValidationError{}

// Validate checks the field values on UpdateResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceResponseMultiError, or nil if none found.
func (m *UpdateResourceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResource()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateResourceResponseValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateResourceResponseValidationError{
				field:  "Resource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateResourceResponseMultiError(errors)
	}

	return nil
}

// UpdateResourceResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateResourceResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateResourceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceResponseMultiError) AllErrors() []error { return m }

// UpdateResourceResponseValidationError is the validation error returned by
// UpdateResourceResponse.Validate if the designated constraints aren't met.
type UpdateResourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceResponseValidationError) ErrorName() string {
	return "UpdateResourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceResponseValidationError{}

// Validate checks the field values on DeleteResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteResourceRequestMultiError, or nil if none found.
func (m *DeleteResourceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteResourceRequestMultiError(errors)
	}

	return nil
}

// DeleteResourceRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteResourceRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteResourceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceRequestMultiError) AllErrors() []error { return m }

// DeleteResourceRequestValidationError is the validation error returned by
// DeleteResourceRequest.Validate if the designated constraints aren't met.
type DeleteResourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceRequestValidationError) ErrorName() string {
	return "DeleteResourceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceRequestValidationError{}

// Validate checks the field values on DeleteResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteResourceResponseMultiError, or nil if none found.
func (m *DeleteResourceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEmpty()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteResourceResponseValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteResourceResponseValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmpty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteResourceResponseValidationError{
				field:  "Empty",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeleteResourceResponseMultiError(errors)
	}

	return nil
}

// DeleteResourceResponseMultiError is an error wrapping multiple validation
// errors returned by DeleteResourceResponse.ValidateAll() if the designated
// constraints aren't met.
type DeleteResourceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceResponseMultiError) AllErrors() []error { return m }

// DeleteResourceResponseValidationError is the validation error returned by
// DeleteResourceResponse.Validate if the designated constraints aren't met.
type DeleteResourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceResponseValidationError) ErrorName() string {
	return "DeleteResourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceResponseValidationError{}

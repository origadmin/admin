// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: system/position.proto

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

// Validate checks the field values on ListPositionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPositionsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPositionsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPositionsRequestMultiError, or nil if none found.
func (m *ListPositionsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPositionsRequest) validate(all bool) error {
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

	if len(errors) > 0 {
		return ListPositionsRequestMultiError(errors)
	}

	return nil
}

// ListPositionsRequestMultiError is an error wrapping multiple validation
// errors returned by ListPositionsRequest.ValidateAll() if the designated
// constraints aren't met.
type ListPositionsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPositionsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPositionsRequestMultiError) AllErrors() []error { return m }

// ListPositionsRequestValidationError is the validation error returned by
// ListPositionsRequest.Validate if the designated constraints aren't met.
type ListPositionsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPositionsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPositionsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPositionsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPositionsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPositionsRequestValidationError) ErrorName() string {
	return "ListPositionsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPositionsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPositionsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPositionsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPositionsRequestValidationError{}

// Validate checks the field values on ListPositionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPositionsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPositionsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPositionsResponseMultiError, or nil if none found.
func (m *ListPositionsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPositionsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalSize

	for idx, item := range m.GetPositions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPositionsResponseValidationError{
						field:  fmt.Sprintf("Positions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPositionsResponseValidationError{
						field:  fmt.Sprintf("Positions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPositionsResponseValidationError{
					field:  fmt.Sprintf("Positions[%v]", idx),
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
					errors = append(errors, ListPositionsResponseValidationError{
						field:  "Extra",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPositionsResponseValidationError{
						field:  "Extra",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetExtra()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPositionsResponseValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListPositionsResponseMultiError(errors)
	}

	return nil
}

// ListPositionsResponseMultiError is an error wrapping multiple validation
// errors returned by ListPositionsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListPositionsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPositionsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPositionsResponseMultiError) AllErrors() []error { return m }

// ListPositionsResponseValidationError is the validation error returned by
// ListPositionsResponse.Validate if the designated constraints aren't met.
type ListPositionsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPositionsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPositionsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPositionsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPositionsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPositionsResponseValidationError) ErrorName() string {
	return "ListPositionsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPositionsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPositionsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPositionsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPositionsResponseValidationError{}

// Validate checks the field values on GetPositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPositionRequestMultiError, or nil if none found.
func (m *GetPositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetPositionRequestMultiError(errors)
	}

	return nil
}

// GetPositionRequestMultiError is an error wrapping multiple validation errors
// returned by GetPositionRequest.ValidateAll() if the designated constraints
// aren't met.
type GetPositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPositionRequestMultiError) AllErrors() []error { return m }

// GetPositionRequestValidationError is the validation error returned by
// GetPositionRequest.Validate if the designated constraints aren't met.
type GetPositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPositionRequestValidationError) ErrorName() string {
	return "GetPositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPositionRequestValidationError{}

// Validate checks the field values on GetPositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetPositionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetPositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetPositionResponseMultiError, or nil if none found.
func (m *GetPositionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetPositionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetPositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetPositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPositionResponseValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetPositionResponseMultiError(errors)
	}

	return nil
}

// GetPositionResponseMultiError is an error wrapping multiple validation
// errors returned by GetPositionResponse.ValidateAll() if the designated
// constraints aren't met.
type GetPositionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetPositionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetPositionResponseMultiError) AllErrors() []error { return m }

// GetPositionResponseValidationError is the validation error returned by
// GetPositionResponse.Validate if the designated constraints aren't met.
type GetPositionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPositionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPositionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPositionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPositionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPositionResponseValidationError) ErrorName() string {
	return "GetPositionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPositionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPositionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPositionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPositionResponseValidationError{}

// Validate checks the field values on CreatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePositionRequestMultiError, or nil if none found.
func (m *CreatePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Parent

	// no validation rules for PositionId

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePositionRequestValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePositionRequestMultiError(errors)
	}

	return nil
}

// CreatePositionRequestMultiError is an error wrapping multiple validation
// errors returned by CreatePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type CreatePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePositionRequestMultiError) AllErrors() []error { return m }

// CreatePositionRequestValidationError is the validation error returned by
// CreatePositionRequest.Validate if the designated constraints aren't met.
type CreatePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePositionRequestValidationError) ErrorName() string {
	return "CreatePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePositionRequestValidationError{}

// Validate checks the field values on CreatePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePositionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePositionResponseMultiError, or nil if none found.
func (m *CreatePositionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePositionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePositionResponseValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePositionResponseMultiError(errors)
	}

	return nil
}

// CreatePositionResponseMultiError is an error wrapping multiple validation
// errors returned by CreatePositionResponse.ValidateAll() if the designated
// constraints aren't met.
type CreatePositionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePositionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePositionResponseMultiError) AllErrors() []error { return m }

// CreatePositionResponseValidationError is the validation error returned by
// CreatePositionResponse.Validate if the designated constraints aren't met.
type CreatePositionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePositionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePositionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePositionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePositionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePositionResponseValidationError) ErrorName() string {
	return "CreatePositionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePositionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePositionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePositionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePositionResponseValidationError{}

// Validate checks the field values on UpdatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePositionRequestMultiError, or nil if none found.
func (m *UpdatePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePositionRequestValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePositionRequestValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePositionRequestMultiError(errors)
	}

	return nil
}

// UpdatePositionRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePositionRequestMultiError) AllErrors() []error { return m }

// UpdatePositionRequestValidationError is the validation error returned by
// UpdatePositionRequest.Validate if the designated constraints aren't met.
type UpdatePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePositionRequestValidationError) ErrorName() string {
	return "UpdatePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePositionRequestValidationError{}

// Validate checks the field values on UpdatePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePositionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePositionResponseMultiError, or nil if none found.
func (m *UpdatePositionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePositionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPosition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePositionResponseValidationError{
					field:  "Position",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPosition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePositionResponseValidationError{
				field:  "Position",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePositionResponseMultiError(errors)
	}

	return nil
}

// UpdatePositionResponseMultiError is an error wrapping multiple validation
// errors returned by UpdatePositionResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdatePositionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePositionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePositionResponseMultiError) AllErrors() []error { return m }

// UpdatePositionResponseValidationError is the validation error returned by
// UpdatePositionResponse.Validate if the designated constraints aren't met.
type UpdatePositionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePositionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePositionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePositionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePositionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePositionResponseValidationError) ErrorName() string {
	return "UpdatePositionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePositionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePositionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePositionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePositionResponseValidationError{}

// Validate checks the field values on DeletePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePositionRequestMultiError, or nil if none found.
func (m *DeletePositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeletePositionRequestMultiError(errors)
	}

	return nil
}

// DeletePositionRequestMultiError is an error wrapping multiple validation
// errors returned by DeletePositionRequest.ValidateAll() if the designated
// constraints aren't met.
type DeletePositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePositionRequestMultiError) AllErrors() []error { return m }

// DeletePositionRequestValidationError is the validation error returned by
// DeletePositionRequest.Validate if the designated constraints aren't met.
type DeletePositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePositionRequestValidationError) ErrorName() string {
	return "DeletePositionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePositionRequestValidationError{}

// Validate checks the field values on DeletePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePositionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePositionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePositionResponseMultiError, or nil if none found.
func (m *DeletePositionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePositionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEmpty()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeletePositionResponseValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeletePositionResponseValidationError{
					field:  "Empty",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEmpty()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeletePositionResponseValidationError{
				field:  "Empty",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DeletePositionResponseMultiError(errors)
	}

	return nil
}

// DeletePositionResponseMultiError is an error wrapping multiple validation
// errors returned by DeletePositionResponse.ValidateAll() if the designated
// constraints aren't met.
type DeletePositionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePositionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePositionResponseMultiError) AllErrors() []error { return m }

// DeletePositionResponseValidationError is the validation error returned by
// DeletePositionResponse.Validate if the designated constraints aren't met.
type DeletePositionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePositionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePositionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePositionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePositionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePositionResponseValidationError) ErrorName() string {
	return "DeletePositionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePositionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePositionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePositionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePositionResponseValidationError{}

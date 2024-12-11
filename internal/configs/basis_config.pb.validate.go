// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configs/basis_config.proto

package configs

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

// Validate checks the field values on BasisConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BasisConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BasisConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BasisConfigMultiError, or
// nil if none found.
func (m *BasisConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *BasisConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRootUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BasisConfigValidationError{
					field:  "RootUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BasisConfigValidationError{
					field:  "RootUser",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRootUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BasisConfigValidationError{
				field:  "RootUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCaptcha()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BasisConfigValidationError{
					field:  "Captcha",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BasisConfigValidationError{
					field:  "Captcha",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCaptcha()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BasisConfigValidationError{
				field:  "Captcha",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BasisConfigMultiError(errors)
	}

	return nil
}

// BasisConfigMultiError is an error wrapping multiple validation errors
// returned by BasisConfig.ValidateAll() if the designated constraints aren't met.
type BasisConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BasisConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BasisConfigMultiError) AllErrors() []error { return m }

// BasisConfigValidationError is the validation error returned by
// BasisConfig.Validate if the designated constraints aren't met.
type BasisConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BasisConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BasisConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BasisConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BasisConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BasisConfigValidationError) ErrorName() string { return "BasisConfigValidationError" }

// Error satisfies the builtin error interface
func (e BasisConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBasisConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BasisConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BasisConfigValidationError{}

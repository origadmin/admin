// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: configs/captcha.proto

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

// Validate checks the field values on Captcha with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Captcha) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Captcha with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CaptchaMultiError, or nil if none found.
func (m *Captcha) ValidateAll() error {
	return m.validate(true)
}

func (m *Captcha) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Length

	// no validation rules for Width

	// no validation rules for Height

	if _, ok := _Captcha_CacheType_InLookup[m.GetCacheType()]; !ok {
		err := CaptchaValidationError{
			field:  "CacheType",
			reason: "value must be in list [memory redis]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRedis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CaptchaValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CaptchaValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRedis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CaptchaValidationError{
				field:  "Redis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CaptchaMultiError(errors)
	}

	return nil
}

// CaptchaMultiError is an error wrapping multiple validation errors returned
// by Captcha.ValidateAll() if the designated constraints aren't met.
type CaptchaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CaptchaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CaptchaMultiError) AllErrors() []error { return m }

// CaptchaValidationError is the validation error returned by Captcha.Validate
// if the designated constraints aren't met.
type CaptchaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CaptchaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CaptchaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CaptchaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CaptchaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CaptchaValidationError) ErrorName() string { return "CaptchaValidationError" }

// Error satisfies the builtin error interface
func (e CaptchaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaptcha.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CaptchaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CaptchaValidationError{}

var _Captcha_CacheType_InLookup = map[string]struct{}{
	"memory": {},
	"redis":  {},
}

// Validate checks the field values on Captcha_Redis with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Captcha_Redis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Captcha_Redis with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Captcha_RedisMultiError, or
// nil if none found.
func (m *Captcha_Redis) ValidateAll() error {
	return m.validate(true)
}

func (m *Captcha_Redis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Db

	// no validation rules for KeyPrefix

	if len(errors) > 0 {
		return Captcha_RedisMultiError(errors)
	}

	return nil
}

// Captcha_RedisMultiError is an error wrapping multiple validation errors
// returned by Captcha_Redis.ValidateAll() if the designated constraints
// aren't met.
type Captcha_RedisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Captcha_RedisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Captcha_RedisMultiError) AllErrors() []error { return m }

// Captcha_RedisValidationError is the validation error returned by
// Captcha_Redis.Validate if the designated constraints aren't met.
type Captcha_RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Captcha_RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Captcha_RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Captcha_RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Captcha_RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Captcha_RedisValidationError) ErrorName() string { return "Captcha_RedisValidationError" }

// Error satisfies the builtin error interface
func (e Captcha_RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCaptcha_Redis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Captcha_RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Captcha_RedisValidationError{}

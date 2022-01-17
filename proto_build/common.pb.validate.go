// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package proto_build

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

// Validate checks the field values on ImageInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ImageInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImageInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ImageInfoMultiError, or nil
// if none found.
func (m *ImageInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *ImageInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ImageInfoValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Type

	// no validation rules for Url

	if m.GetWidth() <= 0 {
		err := ImageInfoValidationError{
			field:  "Width",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetHeight() <= 0 {
		err := ImageInfoValidationError{
			field:  "Height",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ImageInfoMultiError(errors)
	}

	return nil
}

// ImageInfoMultiError is an error wrapping multiple validation errors returned
// by ImageInfo.ValidateAll() if the designated constraints aren't met.
type ImageInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageInfoMultiError) AllErrors() []error { return m }

// ImageInfoValidationError is the validation error returned by
// ImageInfo.Validate if the designated constraints aren't met.
type ImageInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageInfoValidationError) ErrorName() string { return "ImageInfoValidationError" }

// Error satisfies the builtin error interface
func (e ImageInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageInfoValidationError{}

// Validate checks the field values on Pager with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pager) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pager with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PagerMultiError, or nil if none found.
func (m *Pager) ValidateAll() error {
	return m.validate(true)
}

func (m *Pager) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageZie

	// no validation rules for TotalRows

	// no validation rules for TotalPages

	if len(errors) > 0 {
		return PagerMultiError(errors)
	}

	return nil
}

// PagerMultiError is an error wrapping multiple validation errors returned by
// Pager.ValidateAll() if the designated constraints aren't met.
type PagerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PagerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PagerMultiError) AllErrors() []error { return m }

// PagerValidationError is the validation error returned by Pager.Validate if
// the designated constraints aren't met.
type PagerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagerValidationError) ErrorName() string { return "PagerValidationError" }

// Error satisfies the builtin error interface
func (e PagerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPager.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagerValidationError{}

// Validate checks the field values on CommonResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommonResponseMultiError,
// or nil if none found.
func (m *CommonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ErrorCode

	// no validation rules for BusinessCode

	// no validation rules for Message

	if len(errors) > 0 {
		return CommonResponseMultiError(errors)
	}

	return nil
}

// CommonResponseMultiError is an error wrapping multiple validation errors
// returned by CommonResponse.ValidateAll() if the designated constraints
// aren't met.
type CommonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonResponseMultiError) AllErrors() []error { return m }

// CommonResponseValidationError is the validation error returned by
// CommonResponse.Validate if the designated constraints aren't met.
type CommonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonResponseValidationError) ErrorName() string { return "CommonResponseValidationError" }

// Error satisfies the builtin error interface
func (e CommonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonResponseValidationError{}

// Validate checks the field values on UploadImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadImageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadImageRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadImageRequestMultiError, or nil if none found.
func (m *UploadImageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadImageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetContent()) < 1 {
		err := UploadImageRequestValidationError{
			field:  "Content",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UploadImageRequestMultiError(errors)
	}

	return nil
}

// UploadImageRequestMultiError is an error wrapping multiple validation errors
// returned by UploadImageRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadImageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadImageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadImageRequestMultiError) AllErrors() []error { return m }

// UploadImageRequestValidationError is the validation error returned by
// UploadImageRequest.Validate if the designated constraints aren't met.
type UploadImageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadImageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadImageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadImageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadImageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadImageRequestValidationError) ErrorName() string {
	return "UploadImageRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadImageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadImageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadImageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadImageRequestValidationError{}

// Validate checks the field values on UploadImageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UploadImageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadImageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadImageResponseMultiError, or nil if none found.
func (m *UploadImageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadImageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return UploadImageResponseMultiError(errors)
	}

	return nil
}

// UploadImageResponseMultiError is an error wrapping multiple validation
// errors returned by UploadImageResponse.ValidateAll() if the designated
// constraints aren't met.
type UploadImageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadImageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadImageResponseMultiError) AllErrors() []error { return m }

// UploadImageResponseValidationError is the validation error returned by
// UploadImageResponse.Validate if the designated constraints aren't met.
type UploadImageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadImageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadImageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadImageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadImageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadImageResponseValidationError) ErrorName() string {
	return "UploadImageResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UploadImageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadImageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadImageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadImageResponseValidationError{}

// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// ValidateFields checks the field values on GetGatewayConfigurationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GetGatewayConfigurationRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetGatewayConfigurationRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "gateway_ids":

			if m.GetGatewayIds() == nil {
				return GetGatewayConfigurationRequestValidationError{
					field:  "gateway_ids",
					reason: "value is required",
				}
			}

			if v, ok := interface{}(m.GetGatewayIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetGatewayConfigurationRequestValidationError{
						field:  "gateway_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "format":

			if utf8.RuneCountInString(m.GetFormat()) > 36 {
				return GetGatewayConfigurationRequestValidationError{
					field:  "format",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GetGatewayConfigurationRequest_Format_Pattern.MatchString(m.GetFormat()) {
				return GetGatewayConfigurationRequestValidationError{
					field:  "format",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "type":

			if utf8.RuneCountInString(m.GetType()) > 36 {
				return GetGatewayConfigurationRequestValidationError{
					field:  "type",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GetGatewayConfigurationRequest_Type_Pattern.MatchString(m.GetType()) {
				return GetGatewayConfigurationRequestValidationError{
					field:  "type",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "filename":

			if utf8.RuneCountInString(m.GetFilename()) > 36 {
				return GetGatewayConfigurationRequestValidationError{
					field:  "filename",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GetGatewayConfigurationRequest_Filename_Pattern.MatchString(m.GetFilename()) {
				return GetGatewayConfigurationRequestValidationError{
					field:  "filename",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-._]?[a-z0-9]){2,}$|^$\"",
				}
			}

		default:
			return GetGatewayConfigurationRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetGatewayConfigurationRequestValidationError is the validation error
// returned by GetGatewayConfigurationRequest.ValidateFields if the designated
// constraints aren't met.
type GetGatewayConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGatewayConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGatewayConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGatewayConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGatewayConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGatewayConfigurationRequestValidationError) ErrorName() string {
	return "GetGatewayConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetGatewayConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGatewayConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGatewayConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGatewayConfigurationRequestValidationError{}

var _GetGatewayConfigurationRequest_Format_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

var _GetGatewayConfigurationRequest_Type_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

var _GetGatewayConfigurationRequest_Filename_Pattern = regexp.MustCompile("^[a-z0-9](?:[-._]?[a-z0-9]){2,}$|^$")

// ValidateFields checks the field values on GetGatewayConfigurationResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *GetGatewayConfigurationResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetGatewayConfigurationResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "contents":
			// no validation rules for Contents
		default:
			return GetGatewayConfigurationResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetGatewayConfigurationResponseValidationError is the validation error
// returned by GetGatewayConfigurationResponse.ValidateFields if the
// designated constraints aren't met.
type GetGatewayConfigurationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGatewayConfigurationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGatewayConfigurationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGatewayConfigurationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGatewayConfigurationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGatewayConfigurationResponseValidationError) ErrorName() string {
	return "GetGatewayConfigurationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetGatewayConfigurationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGatewayConfigurationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGatewayConfigurationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGatewayConfigurationResponseValidationError{}

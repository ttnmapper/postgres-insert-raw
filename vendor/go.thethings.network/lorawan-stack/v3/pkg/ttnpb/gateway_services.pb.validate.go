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

// ValidateFields checks the field values on PullGatewayConfigurationRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *PullGatewayConfigurationRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = PullGatewayConfigurationRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "gateway_ids":

			if v, ok := interface{}(m.GetGatewayIds()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return PullGatewayConfigurationRequestValidationError{
						field:  "gateway_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "field_mask":

			if v, ok := interface{}(m.GetFieldMask()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return PullGatewayConfigurationRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return PullGatewayConfigurationRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// PullGatewayConfigurationRequestValidationError is the validation error
// returned by PullGatewayConfigurationRequest.ValidateFields if the
// designated constraints aren't met.
type PullGatewayConfigurationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PullGatewayConfigurationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PullGatewayConfigurationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PullGatewayConfigurationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PullGatewayConfigurationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PullGatewayConfigurationRequestValidationError) ErrorName() string {
	return "PullGatewayConfigurationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PullGatewayConfigurationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPullGatewayConfigurationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PullGatewayConfigurationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PullGatewayConfigurationRequestValidationError{}

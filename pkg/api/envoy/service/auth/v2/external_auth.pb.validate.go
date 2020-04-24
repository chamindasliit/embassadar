// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/auth/v2/external_auth.proto

package envoy_service_auth_v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _external_auth_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CheckRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAttributes()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CheckRequestValidationError{
					field:  "Attributes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CheckRequestValidationError is the validation error returned by
// CheckRequest.Validate if the designated constraints aren't met.
type CheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckRequestValidationError) ErrorName() string { return "CheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckRequestValidationError{}

// Validate checks the field values on DeniedHttpResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeniedHttpResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetStatus() == nil {
		return DeniedHttpResponseValidationError{
			field:  "Status",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetStatus()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DeniedHttpResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return DeniedHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for Body

	return nil
}

// DeniedHttpResponseValidationError is the validation error returned by
// DeniedHttpResponse.Validate if the designated constraints aren't met.
type DeniedHttpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeniedHttpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeniedHttpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeniedHttpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeniedHttpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeniedHttpResponseValidationError) ErrorName() string {
	return "DeniedHttpResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeniedHttpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeniedHttpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeniedHttpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeniedHttpResponseValidationError{}

// Validate checks the field values on OkHttpResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OkHttpResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return OkHttpResponseValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// OkHttpResponseValidationError is the validation error returned by
// OkHttpResponse.Validate if the designated constraints aren't met.
type OkHttpResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OkHttpResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OkHttpResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OkHttpResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OkHttpResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OkHttpResponseValidationError) ErrorName() string { return "OkHttpResponseValidationError" }

// Error satisfies the builtin error interface
func (e OkHttpResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOkHttpResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OkHttpResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OkHttpResponseValidationError{}

// Validate checks the field values on CheckResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetStatus()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CheckResponseValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	switch m.HttpResponse.(type) {

	case *CheckResponse_DeniedResponse:

		{
			tmp := m.GetDeniedResponse()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CheckResponseValidationError{
						field:  "DeniedResponse",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *CheckResponse_OkResponse:

		{
			tmp := m.GetOkResponse()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return CheckResponseValidationError{
						field:  "OkResponse",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// CheckResponseValidationError is the validation error returned by
// CheckResponse.Validate if the designated constraints aren't met.
type CheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckResponseValidationError) ErrorName() string { return "CheckResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckResponseValidationError{}

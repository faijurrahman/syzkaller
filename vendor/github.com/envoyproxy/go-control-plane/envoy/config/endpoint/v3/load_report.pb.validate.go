//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/endpoint/v3/load_report.proto

package endpointv3

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

// Validate checks the field values on UpstreamLocalityStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamLocalityStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamLocalityStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamLocalityStatsMultiError, or nil if none found.
func (m *UpstreamLocalityStats) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamLocalityStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocality()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamLocalityStatsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamLocalityStatsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamLocalityStatsValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TotalSuccessfulRequests

	// no validation rules for TotalRequestsInProgress

	// no validation rules for TotalErrorRequests

	// no validation rules for TotalIssuedRequests

	for idx, item := range m.GetLoadMetricStats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamLocalityStatsValidationError{
					field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetUpstreamEndpointStats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("UpstreamEndpointStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("UpstreamEndpointStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamLocalityStatsValidationError{
					field:  fmt.Sprintf("UpstreamEndpointStats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Priority

	if len(errors) > 0 {
		return UpstreamLocalityStatsMultiError(errors)
	}

	return nil
}

// UpstreamLocalityStatsMultiError is an error wrapping multiple validation
// errors returned by UpstreamLocalityStats.ValidateAll() if the designated
// constraints aren't met.
type UpstreamLocalityStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamLocalityStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamLocalityStatsMultiError) AllErrors() []error { return m }

// UpstreamLocalityStatsValidationError is the validation error returned by
// UpstreamLocalityStats.Validate if the designated constraints aren't met.
type UpstreamLocalityStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamLocalityStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamLocalityStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamLocalityStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamLocalityStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamLocalityStatsValidationError) ErrorName() string {
	return "UpstreamLocalityStatsValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamLocalityStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamLocalityStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamLocalityStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamLocalityStatsValidationError{}

// Validate checks the field values on UpstreamEndpointStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpstreamEndpointStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpstreamEndpointStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpstreamEndpointStatsMultiError, or nil if none found.
func (m *UpstreamEndpointStats) ValidateAll() error {
	return m.validate(true)
}

func (m *UpstreamEndpointStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamEndpointStatsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamEndpointStatsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamEndpointStatsValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpstreamEndpointStatsValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpstreamEndpointStatsValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpstreamEndpointStatsValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TotalSuccessfulRequests

	// no validation rules for TotalRequestsInProgress

	// no validation rules for TotalErrorRequests

	// no validation rules for TotalIssuedRequests

	for idx, item := range m.GetLoadMetricStats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UpstreamEndpointStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UpstreamEndpointStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UpstreamEndpointStatsValidationError{
					field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UpstreamEndpointStatsMultiError(errors)
	}

	return nil
}

// UpstreamEndpointStatsMultiError is an error wrapping multiple validation
// errors returned by UpstreamEndpointStats.ValidateAll() if the designated
// constraints aren't met.
type UpstreamEndpointStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpstreamEndpointStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpstreamEndpointStatsMultiError) AllErrors() []error { return m }

// UpstreamEndpointStatsValidationError is the validation error returned by
// UpstreamEndpointStats.Validate if the designated constraints aren't met.
type UpstreamEndpointStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamEndpointStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamEndpointStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamEndpointStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamEndpointStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamEndpointStatsValidationError) ErrorName() string {
	return "UpstreamEndpointStatsValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamEndpointStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamEndpointStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamEndpointStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamEndpointStatsValidationError{}

// Validate checks the field values on EndpointLoadMetricStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EndpointLoadMetricStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndpointLoadMetricStats with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EndpointLoadMetricStatsMultiError, or nil if none found.
func (m *EndpointLoadMetricStats) ValidateAll() error {
	return m.validate(true)
}

func (m *EndpointLoadMetricStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MetricName

	// no validation rules for NumRequestsFinishedWithMetric

	// no validation rules for TotalMetricValue

	if len(errors) > 0 {
		return EndpointLoadMetricStatsMultiError(errors)
	}

	return nil
}

// EndpointLoadMetricStatsMultiError is an error wrapping multiple validation
// errors returned by EndpointLoadMetricStats.ValidateAll() if the designated
// constraints aren't met.
type EndpointLoadMetricStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointLoadMetricStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointLoadMetricStatsMultiError) AllErrors() []error { return m }

// EndpointLoadMetricStatsValidationError is the validation error returned by
// EndpointLoadMetricStats.Validate if the designated constraints aren't met.
type EndpointLoadMetricStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointLoadMetricStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointLoadMetricStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointLoadMetricStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointLoadMetricStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointLoadMetricStatsValidationError) ErrorName() string {
	return "EndpointLoadMetricStatsValidationError"
}

// Error satisfies the builtin error interface
func (e EndpointLoadMetricStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointLoadMetricStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointLoadMetricStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointLoadMetricStatsValidationError{}

// Validate checks the field values on ClusterStats with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClusterStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClusterStats with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClusterStatsMultiError, or
// nil if none found.
func (m *ClusterStats) ValidateAll() error {
	return m.validate(true)
}

func (m *ClusterStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetClusterName()) < 1 {
		err := ClusterStatsValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ClusterServiceName

	if len(m.GetUpstreamLocalityStats()) < 1 {
		err := ClusterStatsValidationError{
			field:  "UpstreamLocalityStats",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetUpstreamLocalityStats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterStatsValidationError{
						field:  fmt.Sprintf("UpstreamLocalityStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterStatsValidationError{
						field:  fmt.Sprintf("UpstreamLocalityStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterStatsValidationError{
					field:  fmt.Sprintf("UpstreamLocalityStats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalDroppedRequests

	for idx, item := range m.GetDroppedRequests() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterStatsValidationError{
						field:  fmt.Sprintf("DroppedRequests[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterStatsValidationError{
						field:  fmt.Sprintf("DroppedRequests[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterStatsValidationError{
					field:  fmt.Sprintf("DroppedRequests[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetLoadReportInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClusterStatsValidationError{
					field:  "LoadReportInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClusterStatsValidationError{
					field:  "LoadReportInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLoadReportInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterStatsValidationError{
				field:  "LoadReportInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ClusterStatsMultiError(errors)
	}

	return nil
}

// ClusterStatsMultiError is an error wrapping multiple validation errors
// returned by ClusterStats.ValidateAll() if the designated constraints aren't met.
type ClusterStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterStatsMultiError) AllErrors() []error { return m }

// ClusterStatsValidationError is the validation error returned by
// ClusterStats.Validate if the designated constraints aren't met.
type ClusterStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStatsValidationError) ErrorName() string { return "ClusterStatsValidationError" }

// Error satisfies the builtin error interface
func (e ClusterStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStatsValidationError{}

// Validate checks the field values on ClusterStats_DroppedRequests with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClusterStats_DroppedRequests) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClusterStats_DroppedRequests with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClusterStats_DroppedRequestsMultiError, or nil if none found.
func (m *ClusterStats_DroppedRequests) ValidateAll() error {
	return m.validate(true)
}

func (m *ClusterStats_DroppedRequests) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCategory()) < 1 {
		err := ClusterStats_DroppedRequestsValidationError{
			field:  "Category",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for DroppedCount

	if len(errors) > 0 {
		return ClusterStats_DroppedRequestsMultiError(errors)
	}

	return nil
}

// ClusterStats_DroppedRequestsMultiError is an error wrapping multiple
// validation errors returned by ClusterStats_DroppedRequests.ValidateAll() if
// the designated constraints aren't met.
type ClusterStats_DroppedRequestsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterStats_DroppedRequestsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterStats_DroppedRequestsMultiError) AllErrors() []error { return m }

// ClusterStats_DroppedRequestsValidationError is the validation error returned
// by ClusterStats_DroppedRequests.Validate if the designated constraints
// aren't met.
type ClusterStats_DroppedRequestsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStats_DroppedRequestsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStats_DroppedRequestsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStats_DroppedRequestsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStats_DroppedRequestsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStats_DroppedRequestsValidationError) ErrorName() string {
	return "ClusterStats_DroppedRequestsValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterStats_DroppedRequestsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStats_DroppedRequests.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStats_DroppedRequestsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStats_DroppedRequestsValidationError{}

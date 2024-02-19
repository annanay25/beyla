// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
)

// Status is an optional final status for this span. Semantically, when Status was not
// set, that means the span ended without errors and to assume Status.Ok (code = 0).
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewStatus function to create new instances.
// Important: zero-initialized instance is not valid for use.
type Status struct {
	orig  *otlptrace.Status
	state *internal.State
}

func newStatus(orig *otlptrace.Status, state *internal.State) Status {
	return Status{orig: orig, state: state}
}

// NewStatus creates a new empty Status.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewStatus() Status {
	state := internal.StateMutable
	return newStatus(&otlptrace.Status{}, &state)
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms Status) MoveTo(dest Status) {
	ms.state.AssertMutable()
	dest.state.AssertMutable()
	*dest.orig = *ms.orig
	*ms.orig = otlptrace.Status{}
}

// Code returns the code associated with this Status.
func (ms Status) Code() StatusCode {
	return StatusCode(ms.orig.Code)
}

// SetCode replaces the code associated with this Status.
func (ms Status) SetCode(v StatusCode) {
	ms.state.AssertMutable()
	ms.orig.Code = otlptrace.Status_StatusCode(v)
}

// Message returns the message associated with this Status.
func (ms Status) Message() string {
	return ms.orig.Message
}

// SetMessage replaces the message associated with this Status.
func (ms Status) SetMessage(v string) {
	ms.state.AssertMutable()
	ms.orig.Message = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms Status) CopyTo(dest Status) {
	dest.state.AssertMutable()
	dest.SetCode(ms.Code())
	dest.SetMessage(ms.Message())
}
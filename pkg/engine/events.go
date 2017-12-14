// Copyright 2017, Pulumi Corporation.  All rights reserved.

package engine

import (
	"fmt"

	"github.com/pulumi/pulumi/pkg/diag"
)

// Event represents an event generated by the engine during an operation. The underlying
// type for the `Payload` field will differ depending on the value of the `Type` field
type Event struct {
	Type    EventType
	Payload interface{}
}

// EventType is the kind of event being emitted.
type EventType string

const (
	CancelEvent      EventType = "cancel"
	StdoutColorEvent EventType = "stdoutcolor"
	DiagEvent        EventType = "diag"
)

func cancelEvent() Event {
	return Event{Type: CancelEvent}
}

// DiagEventPayload is the payload for an event with type `diag`
type DiagEventPayload struct {
	Severity diag.Severity
	Color    diag.Color
	Message  string
}

type StdoutEventPayload struct {
	Color   diag.Color
	Message string
}

func stdOutEventWithColor(s fmt.Stringer, color diag.Color) Event {
	return Event{
		Type: StdoutColorEvent,
		Payload: StdoutEventPayload{
			Color:   color,
			Message: s.String(),
		},
	}
}

func diagDebugEvent(color diag.Color, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Debug,
			Color:    color,
			Message:  msg,
		},
	}
}

func diagInfoEvent(color diag.Color, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Info,
			Color:    color,
			Message:  msg,
		},
	}
}

func diagInfoerrEvent(color diag.Color, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Infoerr,
			Color:    color,
			Message:  msg,
		},
	}
}

func diagErrorEvent(color diag.Color, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Error,
			Color:    color,
			Message:  msg,
		},
	}
}

func diagWarningEvent(color diag.Color, msg string) Event {
	return Event{
		Type: DiagEvent,
		Payload: DiagEventPayload{
			Severity: diag.Warning,
			Color:    color,
			Message:  msg,
		},
	}
}

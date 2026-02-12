/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package duration implements the functions, types, and interfaces for the module.
package duration

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

type Duration = time.Duration

func IsZero(duration *durationpb.Duration) bool {
	return duration.AsDuration() == time.Duration(0)
}

// ParseDuration converts a string representation of a duration to a time.ParseDuration.
// It uses time.ParseDuration to parse the string, and returns a zero duration
// if the parsing fails.
func ParseDuration(duration string) time.Duration {
	// Attempt to parse the duration string into a time.ParseDuration.
	t, err := time.ParseDuration(duration)
	if err != nil {
		// If parsing fails, return a zero duration.
		return time.Duration(0)
	}
	return t
}

// String converts a time.ParseDuration to a string representation.
// It uses the String method of time.ParseDuration to generate the string.
func String(duration time.Duration) string {
	return duration.String()
}

// PBString converts a protobuf duration to a string representation.
// It uses the String method of durationpb.ParseDuration to generate the string.
func PBString(duration *durationpb.Duration) string {
	return duration.String()
}

// ToPB converts a string representation of a duration to a protobuf duration.
// It first converts the string to a time.ParseDuration using the ParseDuration function,
// and then uses durationpb.New to create a protobuf duration from it.
func ToPB(duration string) *durationpb.Duration {
	return durationpb.New(ParseDuration(duration))
}

// FromPB converts a protobuf duration to a time.ParseDuration.
// It uses the AsDuration method of durationpb.ParseDuration to generate the time.ParseDuration.
func FromPB(duration *durationpb.Duration) time.Duration {
	return duration.AsDuration()
}

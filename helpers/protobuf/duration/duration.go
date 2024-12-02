/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package duration implements the functions, types, and interfaces for the module.
package duration

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

// Duration converts a string representation of a duration to a time.Duration.
// It uses time.ParseDuration to parse the string, and returns a zero duration
// if the parsing fails.
func Duration(duration string) time.Duration {
	// Attempt to parse the duration string into a time.Duration.
	t, err := time.ParseDuration(duration)
	if err != nil {
		// If parsing fails, return a zero duration.
		return time.Duration(0)
	}
	return t
}

func PB2Duration(duration *durationpb.Duration) time.Duration {
	return duration.AsDuration()
}

// String converts a time.Duration to a string representation.
// It uses the String method of time.Duration to generate the string.
func String(duration time.Duration) string {
	return duration.String()
}

// Str2PB converts a string representation of a duration to a protobuf duration.
// It first converts the string to a time.Duration using the Duration function,
// and then uses durationpb.New to create a protobuf duration from it.
func Str2PB(duration string) *durationpb.Duration {
	return durationpb.New(Duration(duration))
}

// PB2Str converts a protobuf duration to a string representation.
// It uses the String method of durationpb.Duration to generate the string.
func PB2Str(duration *durationpb.Duration) string {
	return duration.String()
}

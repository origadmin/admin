/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package time provides utility functions for handling time operations and conversions
// between standard Go time types and protobuf timestamp formats with timezone support.
package time

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/origadmin/toolkits/i18n/tz"
)

// Time is an alias for the standard time.Time type for convenience
type Time = time.Time

var (
	// ZeroTime represents the zero value of time.Time (January 1, year 1, 00:00:00 UTC)
	ZeroTime = time.Time{}
)

// IsZero checks if the given time is zero time in the local timezone.
// It converts the time to the local timezone before checking.
// Returns true if the time is zero time, false otherwise.
func IsZero(t time.Time) bool {
	location, err := tz.GetLocation()
	if err != nil {
		return false
	}
	return t.In(location).IsZero()
}

// ParseTime parses a string in RFC3339 format to time.Time using the local timezone.
// Returns ZeroTime if parsing fails or timezone cannot be determined.
func ParseTime(t string) time.Time {
	location, err := tz.GetLocation()
	if err != nil {
		return ZeroTime
	}
	locTime, err := time.ParseInLocation(time.RFC3339, t, location)
	if err != nil {
		return ZeroTime
	}
	return locTime
}

// String converts a time.Time to RFC3339 formatted string using the local timezone.
// Returns an empty string if timezone cannot be determined.
func String(t time.Time) string {
	location, err := tz.GetLocation()
	if err != nil {
		return ""
	}
	return t.In(location).Format(time.RFC3339)
}

// PBString converts a protobuf Timestamp to RFC3339 formatted string.
// Uses UTC time for conversion regardless of local timezone.
func PBString(t *timestamppb.Timestamp) string {
	return t.AsTime().Format(time.RFC3339)
}

// ToPB converts a time.Time to protobuf Timestamp.
// This is a convenience wrapper around timestamppb.New().
func ToPB(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// FromPB converts a protobuf Timestamp to time.Time.
// This is a convenience wrapper around t.AsTime().
func FromPB(t *timestamppb.Timestamp) time.Time {
	return t.AsTime()
}

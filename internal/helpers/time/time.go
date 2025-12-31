/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package time implements the functions, types, and interfaces for the module.
package time

import (
	"time"

	"github.com/origadmin/toolkits/i18n/tz"
)

type Time = time.Time

func IsZero(t time.Time) bool {
	location, err := tz.GetLocation()
	if err != nil {
		return false
	}
	return t.In(location).IsZero()
}

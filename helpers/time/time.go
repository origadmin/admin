/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package time implements the functions, types, and interfaces for the module.
package time

import (
	"time"

	"github.com/origadmin/contrib/i18n/tz"
)

type Time = time.Time

func IsZero(t time.Time) bool {
	location, err := time.LoadLocation(tz.Location())
	if err != nil {
		return false
	}
	return t.In(location).IsZero()
}

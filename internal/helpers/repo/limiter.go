/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package repo

// PageLimiter is a limiter for pagination.
type PageLimiter struct {
	DefaultPageSize int
	MaxPageSize     int
	HardLimit       int
}

// DefaultLimiter returns a default PageLimiter.
func DefaultLimiter() PageLimiter {
	return PageLimiter{
		DefaultPageSize: DefaultPageSize,
		MaxPageSize:     MaxPageSize,
		HardLimit:       HardLimit,
	}
}

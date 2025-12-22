/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package biz

import (
	"origadmin/application/admin/internal/helpers/pagination"
)

var (
	defaultLimiter = pagination.DefaultLimiter()
)

type UpdateHooker interface {
	UpdateRules()
}

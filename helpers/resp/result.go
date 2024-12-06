/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/net/pagination"
)

type Result struct {
	pagination.UnimplementedResponder
	Success bool           `json:"success"`
	Total   int32          `json:"total,omitempty"`
	Data    any            `json:"data,omitempty"`
	Extra   any            `json:"extra,omitempty"`
	Error   *httperr.Error `json:"error,omitempty"`
}

func (r Result) GetSuccess() bool {
	return r.Success
}

func (r Result) GetTotal() int32 {
	return r.Total
}

func (r Result) GetData() any {
	return r.Data
}

func (r Result) GetError() error {
	return r.Error
}

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"
	"net/http"

	"github.com/origadmin/toolkits/errors/httperr"
	"github.com/origadmin/toolkits/net/pagination"
	"google.golang.org/protobuf/proto"
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
func ResultError(rw http.ResponseWriter, status int, err error) {
	code, herr := decodeError(false, status, err)
	_ = resultResult(rw, code, &Result{
		Success: false,
		Error:   herr,
	})
}

func ResultAny(rw http.ResponseWriter, status int, data any, err error) {
	if err != nil {
		ResultError(rw, status, err)
		return
	}
	err = resultResult(rw, http.StatusOK, data)
	if err != nil {
		ResultError(rw, status, err)
		return
	}
}

func resultResult(rw http.ResponseWriter, status int, data any) error {
	if msg, ok := data.(proto.Message); ok {
		buf, err := mo.Marshal(msg)
		if err != nil {
			//context.Error(errors.Wrap(err, "marshal proto message error"))
			return err
		}
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(status)
		_, _ = rw.Write(buf)
		return nil
	}
	return resultJSON(rw, status, data)
}

func resultJSON(rw http.ResponseWriter, status int, data any) error {
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	_, _ = rw.Write(v)
	return nil
}

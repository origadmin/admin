/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"

	commonv1 "github.com/origadmin/runtime/api/gen/go/config/common/v1"
	"github.com/origadmin/runtime/errors"
	datav1 "origadmin/application/admin/internal/helpers/resp/data/v1"
)

type Message interface {
	Message() proto.Message
}

type Error datav1.Error

type Data datav1.Data
type SourceData = datav1.SourceData
type DataArray = datav1.DataArray
type SourceDataArray = datav1.SourceDataArray
type Page = datav1.Page
type SourcePage = datav1.SourcePage
type Token = datav1.Token
type StringResult = datav1.StringData

type Result struct {
	Success       bool            `json:"success,omitempty"`
	Total         int32           `json:"total,omitempty"`
	NextPageToken *string         `json:"next_page_token,omitempty"`
	Data          json.RawMessage `json:"data,omitempty"`
	Extra         json.RawMessage `json:"extra,omitempty"`
	Error         *errors.Error   `json:"error,omitempty"`
}

type ResultBytes struct {
	Success       bool    `json:"success,omitempty"`
	Total         int32   `json:"total,omitempty"`
	NextPageToken *string `json:"next_page_token,omitempty"`
	Data          []byte  `json:"data,omitempty"`
	Error         []byte  `json:"error,omitempty"`
	Extra         []byte  `json:"extra,omitempty"`
}

type PageResponse struct {
	commonv1.Pagination
	Data  []json.RawMessage          `json:"data"`
	Extra map[string]json.RawMessage `json:"extra,omitempty"`
}

func (x *PageResponse) MarshalJSON() ([]byte, error) {
	type Alias PageResponse

	aux := &struct {
		Data  []json.RawMessage          `json:"data"`
		Extra map[string]json.RawMessage `json:"extra,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(x),
	}

	for _, d := range x.Data {
		aux.Data = append(aux.Data, json.RawMessage(d))
	}

	// 转换 extra 字段
	aux.Extra = make(map[string]json.RawMessage)
	for k, v := range x.Extra {
		aux.Extra[k] = json.RawMessage(v)
	}

	return json.Marshal(aux)
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s, code: %d, message: %s, detail: %s", e.Message, e.Code, e.Message, e.Detail)
}

func (r *Data) GetSuccess() bool {
	return r.Success
}

func (r *Data) GetData() any {
	return r.Data
}

func (r *Data) GetError() error {
	return (*Error)(r.Error)
}

func (r *Data) Message() proto.Message {
	return (*datav1.Data)(r)
}

func ResultError(rw http.ResponseWriter, status int, err error) {
	code, herr := decodeError(false, status, err)
	_ = resultResult(rw, code, &Data{
		Success: false,
		Error:   (*datav1.Error)(herr),
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

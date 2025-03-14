/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"
	"net/http"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	datav1 "origadmin/application/admin/helpers/resp/data/v1"
)

const (
	HTTPErrorPrefix = "http.response.status."
)

var (
	Empty *anypb.Any
	mo    protojson.MarshalOptions
	resp  Response
)

func init() {
	Empty = &anypb.Any{}
}

type Response struct {
	alwaysSucceed bool
}

func (r Response) result(context transhttp.Context, status int, data any) {
	if v, ok := data.(Message); ok {
		if err := r.resultProtoJSON(context, status, v.Message()); err != nil {
			r.Error(context, status, err)
		}
		return
	}

	if msg, ok := data.(proto.Message); ok {
		if err := r.resultProtoJSON(context, status, msg); err != nil {
			r.Error(context, status, err)
		}
		return
	}

	if err := r.resultJSON(context, status, data); err != nil {
		r.Error(context, status, err)
	}
	return
}

func (r Response) resultJSON(context transhttp.Context, status int, data any) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	writeBytesWithJSON(context.Response(), status, bytes)
	return nil
}

func (r Response) resultProtoJSON(context transhttp.Context, status int, msg proto.Message) error {
	buf, err := mo.Marshal(msg)
	if err != nil {
		return err
	}
	writeBytesWithJSON(context.Response(), status, buf)
	return nil
}

func (r Response) Error(context transhttp.Context, status int, err error) {
	code, herr := decodeError(r.alwaysSucceed, status, err)
	r.result(context, code, &Data{
		Success: false,
		Error:   (*datav1.Error)(herr),
	})
}

func (r Response) JSON(context transhttp.Context, status int, data any) {
	r.result(context, status, data)
}

func (r Response) Bytes(context transhttp.Context, status int, data []byte) {
	writeBytesWithJSON(context.Response(), status, data)
}

func (r Response) Any(context transhttp.Context, status int, data any, err error) {
	if err != nil {
		r.Error(context, status, err)
		return
	}
	r.JSON(context, http.StatusOK, data)
}

func ResponseErrorEncoder(writer http.ResponseWriter, request *http.Request, err error) {
	ResultError(writer, http.StatusInternalServerError, err)
	return
}

func writeBytesWithJSON(rw http.ResponseWriter, status int, data []byte) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	_, _ = rw.Write(data)
}

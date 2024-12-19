/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"
	"fmt"
	"net/http"

	kerr "github.com/go-kratos/kratos/v2/errors"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	HTTPErrorPrefix = "http.response.status."
)

var Empty *anypb.Any

// Any converts the given arguments into a protobuf Any type.
func Any(args ...any) *anypb.Any {
	switch len(args) {
	case 0:
		return Empty
	case 1:
		if message, ok := args[0].(proto.Message); ok {
			return marshalProto(message)
		}
		return marshalAny(args[0])
	default:
		return marshalAny(args)
	}
}

func marshalProto(message proto.Message) *anypb.Any {
	var extra anypb.Any
	if err := extra.MarshalFrom(message); err != nil {
		return Empty
	}
	return &extra
}

func marshalAny(message any) *anypb.Any {
	var extra anypb.Any
	marshal, err := json.Marshal(message)
	if err != nil {
		return Empty
	}
	if err := protojson.Unmarshal(marshal, &extra); err != nil {
		return Empty
	}
	return &extra
}

type Response struct {
	alwaysSucceed bool
}

var mo protojson.MarshalOptions

func (r Response) result(context transhttp.Context, status int, data any) {
	if msg, ok := data.(proto.Message); ok {
		buf, err := mo.Marshal(msg)
		if err != nil {
			//context.Error(errors.Wrap(err, "marshal proto message error"))
			return
		}
		context.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		context.Response().WriteHeader(status)
		context.Response().Write(buf)
		//context.Set(gins.ContextResponseBodBytesKey, buf)
		//context.Data(status, "application/json; charset=utf-8", buf)
		//context.Abort()
		return
	}
	context.JSON(status, data)
	return
}

func (r Response) resultProtoJSON(context transhttp.Context, status int, msg proto.Message) {
	buf, err := mo.Marshal(msg)
	if err != nil {
		r.Error(context, http.StatusInternalServerError, errors.Wrap(err, "marshal proto message error"))
		return
	}
	//context.Set(gins.ContextResponseBodBytesKey, buf)
	context.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
	context.Response().WriteHeader(status)
	context.Response().Write(buf)
	return
}

func (r Response) Error(context transhttp.Context, status int, err error) {
	code, herr := r.decodeError(status, err)
	r.result(context, code, &Result{
		Success: false,
		Error:   herr,
	})
}

func (r Response) decodeError(code int, err error) (int, *httperr.Error) {
	var ierr *httperr.Error
	var status int
	if ok := errors.As(err, &ierr); ok {
		status = int(ierr.Code)
	} else if ke := kerr.FromError(err); ke != nil {
		status = int(ke.Code)
		if ke.Reason == "" {
			ke.Reason = "UNKNOWN_REASON"
		}
		ierr = &httperr.Error{
			ID:     HTTPErrorPrefix + ke.Reason,
			Code:   ke.Code,
			Detail: ke.Message,
		}
	} else {
		status = http.StatusInternalServerError
		ierr = &httperr.Error{
			ID:     HTTPErrorPrefix + "UNKNOWN",
			Code:   http.StatusInternalServerError,
			Detail: fmt.Sprintf("unhandled error: %v", err),
		}
	}
	if code >= 500 {
		log.Debugf("Error: %v, 5xx status: %d", ierr, code)
	}
	if r.alwaysSucceed {
		status = http.StatusOK
	}
	if !r.alwaysSucceed && code != status {
		ierr.Code = int32(status)
	}
	if code == 0 {
		code = status
	}
	return code, ierr
}

func (r Response) JSON(context transhttp.Context, status int, data any) {
	if v, ok := data.(*Result); ok {
		context.JSON(status, v)
		return
	}
	if v, ok := data.(proto.Message); ok {
		r.resultProtoJSON(context, status, v)
		return
	}
	r.result(context, status, &Result{
		Success: true,
		Data:    data,
	})
}

func (r Response) Any(context transhttp.Context, status int, data any, err error) {
	if err != nil {
		r.Error(context, status, err)
		return
	}
	r.JSON(context, http.StatusOK, data)
}

var Default = Response{}

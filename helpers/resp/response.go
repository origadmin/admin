/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	kerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/origadmin/contrib/transport/gins"
	"github.com/origadmin/runtime/log"
	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/errors/httperr"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
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

func (r Response) result(context *gins.Context, status int, data any) {
	if msg, ok := data.(proto.Message); ok {
		buf, err := mo.Marshal(msg)
		if err != nil {
			context.Error(errors.Wrap(err, "marshal proto message error"))
			return
		}
		context.Set(gins.ContextResponseBodBytesKey, buf)
		context.Data(status, "application/json; charset=utf-8", buf)
		context.Abort()
		return
	}
	context.JSON(status, data)
	return
}

func (r Response) Error(context *gins.Context, status int, err error) {
	code, herr := r.decodeError(status, err)
	r.result(context, code, &Result{
		Success: false,
		Error:   herr,
	})
	_ = context.Error(gin.Error{
		Err:  herr,
		Type: gin.ErrorTypeAny,
		Meta: context.Request.URL.RawQuery,
	})
}

func (r Response) decodeError(code int, err error) (int, *httperr.Error) {
	var ierr *httperr.Error
	var status int
	if ok := errors.As(err, &ierr); ok {
		status = int(ierr.Code)
	} else if ke := kerr.FromError(err); ke != nil {
		status = int(ke.Code)
		ierr = &httperr.Error{
			ID:     "http.Response.status." + ke.Reason,
			Code:   ke.Code,
			Detail: ke.Message,
		}
	} else {
		status = http.StatusInternalServerError
		ierr = &httperr.Error{
			ID:     "http.Response.status.unknown",
			Code:   http.StatusInternalServerError,
			Detail: fmt.Sprintf("unhandled error: %v", err),
		}
	}
	if code >= 500 {
		log.Debugf("Error: %v, 5xx status: %d", ierr, code)
	}
	if r.alwaysSucceed {
		code = http.StatusOK
	}
	if !r.alwaysSucceed && code != status {
		ierr.Code = int32(status)
	}

	return code, ierr
}

func (r Response) JSON(context *gins.Context, status int, data any) {
	if v, ok := data.(*Result); ok {
		r.result(context, status, v)
		return
	}
	r.result(context, status, &Result{
		Success: true,
		Data:    data,
	})
}

func (r Response) Any(context *gins.Context, status int, data any, err error) {
	if err != nil {
		r.Error(context, status, err)
		return
	}
	r.JSON(context, 200, data)
}

var Default = Response{}

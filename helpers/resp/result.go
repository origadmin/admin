/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/origadmin/contrib/transport/gins"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Result(ctx *gin.Context, result any, err error) {
	if err != nil {
		gins.ResultError(ctx, err)
		return
	}
	gins.ResultJSON(ctx, 200, result)
}

// Any converts the given arguments into a protobuf Any type.
func Any(args ...any) *anypb.Any {
	switch len(args) {
	case 0:
		return nil
	case 1:
		var extra anypb.Any
		if message, ok := args[0].(proto.Message); ok {
			if err := extra.MarshalFrom(message); err != nil {
				return nil
			}
			return &extra
		}
		marshal, err := json.Marshal(args[0])
		if err != nil {
			return nil
		}
		if err := protojson.Unmarshal(marshal, &extra); err != nil {
			return nil
		}
		return &extra
	default:
		var extra anypb.Any
		marshal, err := json.Marshal(args)
		if err != nil {
			return nil
		}
		if err := protojson.Unmarshal(marshal, &extra); err != nil {
			return nil
		}
		return &extra
	}
}

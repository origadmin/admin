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

//var Empty = &anypb.Any{
//	TypeUrl: "type.googleapis.com/google.protobuf.Empty",
//	Value:   []byte{},
//}

var Empty *anypb.Any

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

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package resp implements the functions, types, and interfaces for the module.
package resp

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Any converts the given arguments into a protobuf Any type.
func Any(args ...any) *anypb.Any {
	if len(args) == 0 {
		return Empty
	}
	if len(args) == 1 {
		if message, ok := args[0].(proto.Message); ok {
			return marshalProto(message)
		}
		return marshalAny(args[0])
	}
	return marshalAny(args)
}

func marshalProto(message proto.Message) *anypb.Any {
	extra := &anypb.Any{}
	if err := extra.MarshalFrom(message); err != nil {
		return Empty
	}
	return extra
}

func marshalAny(message any) *anypb.Any {
	extra := &anypb.Any{}
	marshal, err := json.Marshal(message)
	if err != nil {
		return Empty
	}
	if err := protojson.Unmarshal(marshal, extra); err != nil {
		return Empty
	}
	return extra
}

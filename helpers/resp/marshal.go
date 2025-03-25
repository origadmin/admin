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
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	datav1 "origadmin/application/admin/helpers/resp/data/v1"
)

// Any converts the given arguments into a protobuf Any type.
func Any(args ...any) *anypb.Any {
	if len(args) == 0 {
		return Empty
	}
	if len(args) == 1 {
		if message, ok := args[0].(proto.Message); ok {
			return Proto2Any(message)
		}
		return Any2AnyPB(args[0])
	}
	var protos []proto.Message
	for i := range args {
		message, ok := args[i].(proto.Message)
		if !ok {
			return Any2AnyPB(args)
		}
		protos = append(protos, message)
	}
	if len(protos) == len(args) {
		return ProtoArray2AnyPB(protos...)
	}
	return Any2AnyPB(args)
}

func ProtoArray2AnyPB[T proto.Message](args ...T) *anypb.Any {
	if len(args) == 0 {
		return Empty
	}
	vals := &datav1.AnyData{Data: Proto2AnyPBArray(args...)}
	val, err := anypb.New(vals)
	if err != nil {
		return Empty
	}
	return val
}

// Proto2AnyPBArray converts the given arguments into a protobuf Any type.
func Proto2AnyPBArray[T proto.Message](args ...T) []*anypb.Any {
	if len(args) == 0 {
		return nil
	}

	var result []*anypb.Any
	for _, arg := range args {
		result = append(result, Proto2Any(arg))
	}
	return result
}

func Proto2Struct[T proto.Message](args ...T) *structpb.Value {
	if len(args) == 0 {
		return nil
	}

	var result []any
	for _, arg := range args {
		result = append(result, arg)
	}
	list, err := structpb.NewList(result)
	if err != nil {
		return nil
	}
	return structpb.NewListValue(list)
}

func Proto2Any(message proto.Message) *anypb.Any {
	val, err := anypb.New(message)
	if err != nil {
		return Empty
	}
	return val
}

func Any2AnyPB(message any) *anypb.Any {
	if message == nil {
		return Empty
	}
	if v, ok := message.(proto.Message); ok {
		return Proto2Any(v)
	}

	marshal, err := json.Marshal(message)
	if err != nil {
		return Empty
	}
	bytes := wrapperspb.Bytes(marshal)
	val, err := anypb.New(bytes)
	if err != nil {
		return Empty
	}
	return val
}

func Proto2JSONArray[T proto.Message](msgs ...T) ([]json.RawMessage, error) {
	var arr []json.RawMessage
	for _, msg := range msgs {
		b, err := protojson.Marshal(msg)
		if err != nil {
			return nil, err
		}
		arr = append(arr, b)
	}
	return arr, nil
}

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package id implements the functions, types, and interfaces for the module.
package id

import (
	"github.com/sony/sonyflake"
)

var (
	generator = New()
)

func New() *sonyflake.Sonyflake {
	s := sonyflake.Settings{}
	generator, err := sonyflake.New(s)
	if err != nil {
		panic(err)
	}
	return generator
}

func Gen() int64 {
	id, err := generator.NextID()
	if err != nil {
		return 0
	}
	return int64(id)
}

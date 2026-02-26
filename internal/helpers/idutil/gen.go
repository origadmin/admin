/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package idutil implements the functions, types, and contracts for the module.
package idutil

import (
	"github.com/origadmin/toolkits/identifier"
	_ "github.com/origadmin/toolkits/identifier/snowflake"
	_ "github.com/origadmin/toolkits/identifier/uuid"
)

var (
	generator     = identifier.Get("snowflake")
	uuidGenerator = identifier.Get("uuid")
)

func Gen() int64 {
	v, _ := generator.GenerateNumber()
	return v
}

func GenUUID() string {
	v, _ := uuidGenerator.GenerateString()
	return v
}

func GenStringUUID() (string, error) {
	v, err := uuidGenerator.GenerateString()
	return v, err
}

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package policy embedding the policy files for Casbin.
package policy

import (
	"embed"
	_ "embed"
)

//go:embed *.csv
var policies embed.FS

// Policies returns the embedded file system containing all policies.
func Policies() embed.FS {
	return policies
}

// Policy reads and returns the contents of a policy file from the embedded file system.
//
// Args:
//
//	name (string): The name of the policy file to read.
//
// Returns:
//
//	([]byte, error): The contents of the policy file as a byte slice, or an error if the file does not exist.
func Policy(name string) ([]byte, error) {
	return policies.ReadFile(name)
}

// MustPolicy reads and returns the contents of a policy file from the embedded file system.
//
// Note: This function is identical to Policy and may be removed in the future.
//
// Args:
//
//	name (string): The name of the policy file to read.
//
// Returns:
//
//	([]byte, error): The contents of the policy file as a byte slice, or an error if the file does not exist.
func MustPolicy(name string) []byte {
	bytes, err := policies.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return bytes
}

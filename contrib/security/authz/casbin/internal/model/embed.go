/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package model embedding the model files for Casbin.
package model

import (
	"embed"
)

//go:embed rbac_model.conf
var DefaultRbacModel string

//go:embed rbac_with_domains.conf
var DefaultRbacWithDomainModel string

//go:embed abac_rule_model.conf
var DefaultAbacModel string

//go:embed basic_model.conf
var DefaultAclModel string

//go:embed keymatch_model.conf
var DefaultRestfullModel string

//go:embed restfull_with_role.conf
var DefaultRestfullWithRoleModel string

//go:embed *.conf
var models embed.FS

// Models returns the embedded file system containing all models.
func Models() embed.FS {
	// Return the embedded file system.
	return models
}

// Model reads a model from the embedded file system by name.
//
// Args:
//
//	name (string): The name of the model to read.
//
// Returns:
//
//	string: The contents of the model file.
//	error: Any error that occurred while reading the model file.
func Model(name string) (string, error) {
	// Read the model file from the embedded file system.
	bytes, err := models.ReadFile(name)
	if err != nil {
		// If an error occurred, return an empty string and the error.
		return "", err
	}
	// Convert the file contents to a string and return.
	return string(bytes), nil
}

// MustModel reads a model from the embedded file system by name, panicking on error.
//
// Args:
//
//	name (string): The name of the model to read.
//
// Returns:
//
//	string: The contents of the model file.
func MustModel(name string) string {
	// Read the model file, panicking on error.
	model, err := Model(name)
	if err != nil {
		// If an error occurred, panic with the error.
		panic(err)
	}
	// Return the model file contents.
	return model
}

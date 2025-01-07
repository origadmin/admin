/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ent implements the functions, types, and interfaces for the module.
package ent

// uniform field length constants
const (
	LenTiny   = 32  // Small fields, such as: code
	LenSmall  = 64  // Smaller fields, such as: name
	LenMedium = 128 // Medium field, such as: title
	LenLarge  = 256 // Large field, such as: description
	LenHuge   = 512 // Huge field, such as: content
)

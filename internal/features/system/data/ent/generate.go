/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package entity is the data access object for SYS.
package ent

//go:generate go run entgo.io/ent/cmd/ent generate --template ./template --feature intercept --feature sql/versioned-migration --feature sql/lock --feature sql/modifier ./schema

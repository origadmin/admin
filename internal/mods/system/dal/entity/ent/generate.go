/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ent is the data access object for SYS.
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --template ./template --feature sql/versioned-migration --feature sql/lock --feature sql/modifier ./schema

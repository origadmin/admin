// Copyright (c) 2024 KasaAdmin. All rights reserved.

// Package ent is the data access object for SYS.
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/lock --feature sql/modifier ./schema
//go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema

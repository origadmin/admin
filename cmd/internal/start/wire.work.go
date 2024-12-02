//go:build !wireinject && GOWORK
// +build !wireinject,GOWORK

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// The build tag makes sure the stub is not built in the final build.
//
//go:generate go run github.com/google/wire/cmd/wire
package start

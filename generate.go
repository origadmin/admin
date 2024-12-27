/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package main is the main package
package main

// for linux or macos you can use make generate
// for windows you can use make this generate

//generate helloworld proto file
// if you want to generate the client code to the same directory, please use the following command
////go:generate kratos proto client -p=../third_party .

//=paths=source_relative:. outputs to the same directory with the proto file
////go:generate protoc -I. -I../third_party --go_out=. --go-http_out=. --go-grpc_out=. --validate_out=lang=go:. --go-gins_out=. ./v1/proto/helloworld/*.proto

//// generate *.pb.go
////go:generate protoc -I. -I./third_party --go_out=. ./api/v1/proto/helloworld/*.proto
//// generate *_http.pb.go
////go:generate protoc -I. -I./third_party --go-http_out=. ./api/v1/proto/helloworld/*.proto
//// generate *_grpc.pb.go
////go:generate protoc -I. -I./third_party --go-grpc_out=. ./api/v1/proto/helloworld/*.proto
//// generate *_gins.pb.go
////go:generate protoc -I. -I./third_party --go-gins_out=. ./api/v1/proto/helloworld/*.proto
//// generate *_errors.pb.go
////go:generate protoc -I. -I./third_party --go-errors_out=. ./api/v1/proto/helloworld/*.proto
//// generate *.pb.validate.go

//// generate openapi.yaml
////go:generate protoc -I. -I./third_party --openapi_out=naming=proto,fq_schema_naming=true,default_response=false:api/v1/services ./api/v1/proto/helloworld/*.proto

// generate a greeter server template
// kratos proto server -t ./internal/agent api/v1/proto/system/menu.proto
// kratos proto server -t ./internal/agent api/v1/proto/system/role.proto
// kratos proto server -t ./internal/agent api/v1/proto/system/user.proto

//go:generate buf dep update
//go:generate buf build
//go:generate buf generate

//go:generate protoc -I. -I./third_party --go_out=paths=source_relative:. ./helpers/resp/data/v1/*.proto

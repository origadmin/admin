//go:build tools

package tools

import (
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking"
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"
	_ "github.com/envoyproxy/protoc-gen-validate"
	_ "github.com/go-kratos/kratos/cmd/kratos/v2"
	_ "github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2"
	_ "github.com/google/gnostic/cmd/protoc-gen-openapi"
	_ "github.com/google/wire/cmd/wire"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

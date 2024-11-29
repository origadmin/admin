GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

PROJECT_ORG=OrigAdmin
THIRD_PARTY_PATH=third_party

PROTO_INTERNAL_PATH=internal
PROTO_TOOLKITS_PATH=toolkits
PROTO_API_PATH=api
OPENAPI_DOCS_PATH=resources/docs/openapi

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	#Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	VERSION=$(shell git describe --tags --always)
	BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
	HEAD_TAG=$(shell git tag --points-at '${gitHash}')
	# gitHash Current commit id, same as gitCommit result
	gitHash = $(shell git rev-parse HEAD)

	# Use PowerShell to find .proto files, convert to relative paths, and replace \ with /
	INTERNAL_PROTO_FILES := $(shell powershell -Command "Get-ChildItem -Recurse ${PROTO_INTERNAL_PATH} -Filter *.proto | Resolve-Path -Relative")
	TOOLKITS_PROTO_FILES := $(shell powershell -Command "Get-ChildItem -Recurse ${PROTO_TOOLKITS_PATH} -Filter *.proto | Resolve-Path -Relative")
	API_PROTO_FILES := $(shell powershell -Command "Get-ChildItem -Recurse ${PROTO_API_PATH} -Filter *.proto | Resolve-Path -Relative")

	# Replace \ with /
	INTERNAL_PROTO_FILES := $(subst \,/, $(INTERNAL_PROTO_FILES))
	TOOLKITS_PROTO_FILES := $(subst \,/, $(TOOLKITS_PROTO_FILES))
	API_PROTO_FILES := $(subst \,/, $(API_PROTO_FILES))

	BUILT_DATE = $(shell powershell -Command "Get-Date -Format 'yyyy-MM-ddTHH:mm:ssK'")
	TREE_STATE = $(shell powershell -Command "if ((git status) -match 'clean') { 'clean' } else { 'dirty' }")
	TAG = $(shell powershell -Command "if ((git tag --points-at '${gitHash}') -match '^v') { '$(HEAD_TAG)' } else { '${gitHash}' }")
	# buildDate = $(shell TZ=Asia/Shanghai date +%F\ %T%z | tr 'T' ' ')
	# same as gitHash previously
	COMMIT = $(shell git log --pretty=format:'%h' -n 1)
else
	VERSION=$(shell git describe --tags --always)
	BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
	HEAD_TAG=$(shell git tag --points-at '${gitHash}')
	# gitHash Current commit id, same as gitCommit result
	gitHash = $(shell git rev-parse HEAD)

	INTERNAL_PROTO_FILES=$(shell find ${PROTO_INTERNAL_PATH} -name *.proto)
	TOOLKITS_PROTO_FILES=$(shell find ${PROTO_TOOLKITS_PATH} -name *.proto)
	API_PROTO_FILES=$(shell find ${PROTO_API_PATH} -name *.proto)

	BUILT_DATE = $(shell TZ=Asia/Shanghai date +%FT%T%z)
	TREE_STATE = $(shell if git status | grep -q 'clean'; then echo clean; else echo dirty; fi)
	TAG = $(shell if git tag --points-at "${gitHash}" | grep -q '^v'; then echo $(HEAD_TAG); else echo ${gitHash}; fi)
	# buildDate = $(shell TZ=Asia/Shanghai date +%F\ %T%z | tr 'T' ' ')
	# same as gitHash previously
	COMMIT = $(shell git log --pretty=format:'%h' -n 1)
endif

BUILT_BY = $(PROJECT_ORG)

ifeq ($(ENV), release)
    LDFLAGS = -s -w
endif
MODULE_PATH=github.com/origadmin/toolkits/version
LDFLAGS := -X $(MODULE_PATH).gitTag=$(TAG) \
           -X $(MODULE_PATH).buildDate=$(BUILT_DATE) \
           -X $(MODULE_PATH).gitCommit=$(COMMIT) \
           -X $(MODULE_PATH).gitTreeState=$(TREE_STATE) \
           -X $(MODULE_PATH).gitBranch=$(BRANCH) \
           -X $(MODULE_PATH).version=$(VERSION)

PROTO_PATH := --proto_path=. --proto_path=./third_party

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest

	go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: deps
# update third_party proto
deps:
	buf export buf.build/bufbuild/protovalidate -o $(THIRD_PARTY_PATH)
	buf export buf.build/protocolbuffers/wellknowntypes -o $(THIRD_PARTY_PATH)
	buf export buf.build/googleapis/googleapis -o $(THIRD_PARTY_PATH)
	buf export buf.build/envoyproxy/protoc-gen-validate -o $(THIRD_PARTY_PATH)
	buf export buf.build/gnostic/gnostic -o $(THIRD_PARTY_PATH)
	buf export buf.build/kratos/apis -o $(THIRD_PARTY_PATH)
	buf export buf.build/origadmin/rpcerr -o $(THIRD_PARTY_PATH)
	buf export buf.build/origadmin/runtime -o $(THIRD_PARTY_PATH)
	buf export buf.build/origadmin/entgen -o $(THIRD_PARTY_PATH)

.PHONY: config
# generate internal proto or use ./internal/generate.go
config: 
	protoc ${PROTO_PATH} \
	--go_out=paths=source_relative:./internal \
	--validate_out=lang=go:. \
	$(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto or use ./api/generate.go
api:
#	protoc --proto_path=./api \
#	       --proto_path=$(THIRD_PARTY_PATH) \
# 	       --go_out=paths=source_relative:./api \
# 	       --go-http_out=paths=source_relative:./api \
# 	       --go-grpc_out=paths=source_relative:./api \
#	       --openapi_out=fq_schema_naming=true,default_response=false:. \
#	       $(API_PROTO_FILES)
	protoc ${PROTO_PATH} \
		--go_out=. \
		--go-http_out=. \
		--go-grpc_out=. \
		--go-gins_out=. \
		--go-errors_out=. \
		--validate_out=lang=go:. \
		$(API_PROTO_FILES)

.PHONY: openapi
# generate the openapi spec file
openapi:
	protoc ${PROTO_PATH} \
	--openapi_out=output_mode=merge,naming=proto,fq_schema_naming=true,default_response=false:${OPENAPI_DOCS_PATH} \
	$(API_PROTO_FILES)

.PHONY: ent
# generate ent proto or use ./toolkits/generate.go
ent:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--ent_out=./database/ent/schema \
		api/v1/proto/secondworld/greeter.proto

.PHONY: pre
# pre
pre:
	goreleaser build --single-target --clean --snapshot

.PHONY: build
# build
build:
	go build -ldflags "$(LDFLAGS)" -gcflags=all="-N -l" -o ./dist/ ./...

.PHONY: release
# release
release:
	goreleaser release --clean

#.PHONY: server
## server used generate a service at first
#server:
#	kratos proto server -t ./internal/mods/helloworld/service ./api/v1/protos/helloworld/greeter.proto
#
#.PHONY: client
## client used when proto file is in the same directory
#client:
#	kratos proto client ./api

.PHONY: generate
# generate
generate:
	go generate ./...
	go mod tidy

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;
	make openapi;

.PHONY: http
# run http request
http:
#	docker run --rm -i -t -v $PWD:/workdir jetbrains/intellij-http-client run.http
#   docker run -v %CD%:/local swaggerapi/swagger-codegen-cli generate -l csharp -o /output/csharp -i https://petstore.swagger.io/v2/swagger.json

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

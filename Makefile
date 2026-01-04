GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

PROJECT_ORG=OrigAdmin
THIRD_PARTY_PATH=third_party

PROTO_INTERNAL_PATH=internal
PROTO_TOOLKITS_PATH=toolkits
PROTO_API_PATH=api
OPENAPI_DOCS_PATH=resources/api-docs/openapi

# Path to the web UI submodule, relative to this Makefile
WEBUI_PATH=./webui

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
	TAG = $(shell powershell -Command "if ((git tag --points-at '${gitHash}') -match '^v') { '$(HEAD_HEAD_TAG)' } else { '${gitHash}' }")
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
#	TREE_STATE = $(shell if git status | grep -q 'clean'; then echo clean; else echo dirty; fi)
   	TREE_STATE = $(shell [ -z "$$(git status --porcelain)" ] && echo clean || echo dirty)
#	TAG = $(shell #if git tag --points-at "${gitHash}" | grep -q '^v'; then echo $(HEAD_TAG); else echo ${gitHash}; fi)
	TAG = $(shell git tag --points-at "${gitHash}" | grep -q '^v' && echo $(HEAD_TAG) || echo ${gitHash})
	# buildDate = $(shell TZ=Asia/Shanghai date +%F\ %T%z | tr 'T' ' ')
	# same as gitHash previously
	COMMIT = $(shell git log --pretty=format:'%h' -n 1)
endif

BUILT_BY = $(PROJECT_ORG)

# LDFLAGS are now primarily managed by .goreleaser.yaml, but can be kept for direct go build commands if any.
# We will let goreleaser handle the version injection to ensure consistency.
LDFLAGS := -X github.com/origadmin/toolkits/version.gitTag=$(TAG) \
           -X github.com/origadmin/toolkits/version.buildDate=$(BUILT_DATE) \
           -X github.com/origadmin/toolkits/version.gitCommit=$(COMMIT) \
           -X github.com/origadmin/toolkits/version.gitTreeState=$(TREE_STATE) \
           -X github.com/origadmin/toolkits/version.gitBranch=$(BRANCH) \
           -X github.com/origadmin/toolkits/version.gitVersion=$(VERSION)

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
	go install entgo.io/ent/cmd/ent@latest

.PHONY: deps
# update third_party proto
deps:
	buf export buf.build/bufbuild/protovalidate -o $(THIRD_PARTY_PATH)
	buf export buf.build/protocolbuffers/wellknowntypes -o $(THIRD_PARTY_PATH)
	buf export buf.build/googleapis/googleapis -o $(THIRD_PARTY_PATH)
	buf export buf.build/envoyproxy/protoc-gen-validate -o $(THIRD_PARTY_PATH)
	buf export buf.build/gnostic/gnostic -o $(THIRD_PARTY_PATH)
	buf export buf.build/kratos/apis -o $(THIRD_PARTY_PATH)
	buf export buf.build/origadmin/runtime -o $(THIRD_PARTY_PATH)
	buf export buf.build/origadmin/contrib -o $(THIRD_PARTY_PATH)

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

.PHONY: build-ui
# build the web UI from the submodule
build-ui:
	@echo "Building Web UI from submodule..."
	@cd $(WEBUI_PATH) && npm install && npm run build

.PHONY: build
# build a standard backend-only snapshot binary
build:
	@echo "Building standard backend-only snapshot..."
	goreleaser build --single-target --clean --snapshot

.PHONY: build-all-in-one
# build an all-in-one snapshot binary with embedded UI
build-all-in-one: build-ui
	@echo "Building all-in-one snapshot with embedded UI..."
	goreleaser build --single-target --clean --snapshot --config .goreleaser.all-in-one.yaml

.PHONY: release
# create a full release (backend-only)
release:
	goreleaser release --clean

.PHONY: release-all-in-one
# create a full all-in-one release with embedded UI
release-all-in-one: build-ui
	goreleaser release --config .goreleaser.all-in-one.yaml --clean

#.PHONY: server
## server used generate a service at first
#server:
#	kratos proto server -t ./internal/features/helloworld/service ./api/v1/protos/helloworld/greeter.proto
#
#.PHONY: client
## client used when proto file is in the same directory
#client:
#	kratos proto client ./api

.PHONY: gen
#gen
#go mod tidy
#buf dep update
#buf build
#buf generate # generate proto files
#go generate ./internal/features/system/dal/entity/generate.go #generate dal entity
#go generate ./cmd/system #generate system module
#go generate ./cmd/internal/start #generate main module start
gen:
	go mod tidy

	buf dep update
	buf build
	buf generate

	@#echo "Generating Protobuf code for helpers/resp/data/v1..."
	@#protoc -I. -I./third_party --go_out=paths=source_relative:. ./helpers/resp/data/v1/*.proto

	@echo "Generating Protobuf code for conf/pb..."
	@protoc -I. -I./third_party --go_out=paths=source_relative:. --validate_out=paths=source_relative,lang=go:. ./internal/conf/pb/*.proto

	go generate ./internal/data/entity/ent/generate.go
	go generate ./cmd/system
	go generate ./cmd/auth

.PHONY: all
# generate all
all:
	$(MAKE) gen;
	$(MAKE) openapi;

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

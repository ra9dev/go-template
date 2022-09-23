GOLANG_CI_LINT_VERSION ?= v1.46.2
lint-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/golangci-lint)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANG_CI_LINT_VERSION)
endif

IMPORTS_REVISER_VERSION ?= v2.5.1
imports-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/goimports-reviser)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/incu6us/goimports-reviser/v2@$(IMPORTS_REVISER_VERSION)
endif

MIGRATE_VERSION ?= v4.15.2
migrate-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/migrate)","")
	GOBIN=$(LOCAL_BIN) go install -tags '$(DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
endif

SWAG_GO_VERSION ?= v1.8.4
swagger-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/swag)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/swaggo/swag/cmd/swag@$(SWAG_GO_VERSION)
endif

PROTOC_GEN_GO_VERSION ?= v1.28
PROTOC_GEN_GO_GRPC_VERSION ?= v1.2
grpc-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/protoc-gen-go)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)
endif
ifeq ("$(wildcard $(LOCAL_BIN)/protoc-gen-go-grpc)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)
endif

deps:
	$(call describe_job,"Installing dependencies")
	$(MAKE) lint-deps
	$(MAKE) imports-deps
	$(MAKE) migrate-deps
	$(MAKE) swagger-deps
	$(MAKE) grpc-deps
	go mod tidy

git-hooks:
	$(call describe_job,"Setting up git hooks")
	/bin/sh ./scripts/hooks.sh

environment:
	$(call describe_job,"Local development setup")
	$(MAKE) git-hooks
	$(MAKE) deps
	docker-compose up --force-recreate --remove-orphans -d
	sleep 5
	$(MAKE) migrate

swagger:
	$(call describe_job,"Generating swagger docs")
	$(MAKE) swagger-deps
	$(LOCAL_BIN)/swag init -g internal/api/admin/doc.go -p snakecase -o docs --parseInternal

PROTOC := $(shell command -v protoc 2> /dev/null)
PROTO_PATH ?= ./pb/*.proto
grpc:
ifndef PROTOC
	$(error "protoc is not installed. Visit https://grpc.io/docs/protoc-installation")
endif
	$(call describe_job,"Generating grpc code")
	protoc \
    	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=. --go_opt=paths=source_relative --go_opt=M$(PROTO_PATH)=./pb \
    	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=M$(PROTO_PATH)=./pb \
    	-I /usr/local/include:$(THIRD_PARTY_PROTO_PATH):. \
    	$(PROTO_PATH)

run-api:
	$(call describe_job,"Starting API server")
	go run cmd/*.go api

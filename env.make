GOLANG_CI_LINT_VERSION ?= latest
lint-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/golangci-lint)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANG_CI_LINT_VERSION)
endif

IMPORTS_REVISER_VERSION ?= latest
imports-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/goimports-reviser)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/incu6us/goimports-reviser/v2@$(IMPORTS_REVISER_VERSION)
endif

MIGRATE_VERSION ?= latest
migrate-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/migrate)","")
	GOBIN=$(LOCAL_BIN) go install -tags '$(DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
endif

SWAG_GO_VERSION ?= latest
swagger-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/swag)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/swaggo/swag/cmd/swag@$(SWAG_GO_VERSION)
endif

PROTOC_GEN_GO_VERSION ?= latest
PROTOC_GEN_GO_GRPC_VERSION ?= latest
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
	#sleep 5 TODO: uncomment when migrations are added
	#$(MAKE) migrate TODO: uncomment when migrations are added

swagger:
	$(call describe_job,"Generating swagger docs")
	$(MAKE) swagger-deps
	$(LOCAL_BIN)/swag init -g internal/api/http/doc.go -p snakecase -o docs --parseInternal

RUN_CMD ?= api
run:
	$(call describe_job,"Running command $(RUN_CMD)")
	go run cmd/*.go $(RUN_CMD)

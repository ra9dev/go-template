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
	GOBIN=$(LOCAL_BIN) go install -tags '$(DB_DRIVER),$(ADS_DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
endif

SWAG_GO_VERSION ?= v1.8.4
swagger-deps:
ifeq ("$(wildcard $(LOCAL_BIN)/swag)","")
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/swaggo/swag/cmd/swag@$(SWAG_GO_VERSION)
endif

deps:
	$(call describe_job,"Installing dependencies")
	$(MAKE) lint-deps
	$(MAKE) imports-deps
	$(MAKE) migrate-deps
	$(MAKE) swagger-deps
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

run-api:
	$(call describe_job,"Starting API server")
	go run cmd/*.go api
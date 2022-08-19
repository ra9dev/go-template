include func.make

DIR:=$(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
LOCAL_BIN:=$(DIR)/bin

GOLANG_CI_LINT_VERSION ?= v1.46.2
lint-deps:
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANG_CI_LINT_VERSION)

IMPORTS_REVISER_VERSION ?= v2.5.1
imports-deps:
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/incu6us/goimports-reviser/v2@$(IMPORTS_REVISER_VERSION)

git-hooks:
	$(call describe_job,"Setting up git hooks")
	/bin/sh ./scripts/hooks.sh

deps:
	$(call describe_job,"Installing dependencies")
	$(MAKE) lint-deps
	$(MAKE) imports-deps
	go mod tidy

environment:
	$(call describe_job,"Local development setup")
	$(MAKE) git-hooks
	$(MAKE) deps
	docker-compose up --force-recreate --remove-orphans -d

run-api:
	$(call describe_job,"Starting API server")
	go run cmd/*.go api
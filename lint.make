include func.make

# TODO remove integration when will refactor testing
GO_IMPORTS := $(LOCAL_BIN)/goimports-reviser
imports:
	$(call describe_job,"Running imports")
	find . -name \*.go -not -path "./vendor/*" -not -path "*/pb/*" -not -path "./integration/*" -exec $(GO_IMPORTS) -file-path {} -rm-unused -set-alias -format \;

GO_LINT := $(LOCAL_BIN)/golangci-lint
lint-fast:
	$(call describe_job,"Running linter")
	$(GO_LINT) run --fast --fix

lint-full:
	$(call describe_job,"Running linter")
	$(GO_LINT) run --fix
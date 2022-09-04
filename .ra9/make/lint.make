imports:
	$(call describe_job,"Running imports")
	$(MAKE) lint-deps
	find . -name \*.go -not -path "./vendor/*" -not -path "*/pb/*" -not -path "./integration/*" -exec $(LOCAL_BIN)/goimports-reviser -file-path {} -rm-unused -set-alias -format \;

lint:
	$(call describe_job,"Running linter")
	$(MAKE) lint-deps
	$(LOCAL_BIN)/golangci-lint run
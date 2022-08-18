MIGRATION_NAME ?= migration
MIGRATE_VERSION ?= v4.15.2
MIGRATE_ARGS ?= up
DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= go-template
DB_DRIVER ?= postgres
DB_URL ?= "$(DB_DRIVER)://$(DB_USER)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migration:
	GOBIN=$(LOCAL_BIN) go install -tags '$(DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
	$(LOCAL_BIN)/migrate create -ext sql -dir migrations -seq $(MIGRATION_NAME)

migrate:
	GOBIN=$(LOCAL_BIN) go install -tags '$(DB_DRIVER)' -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)
	$(LOCAL_BIN)/migrate -path ./migrations -verbose -database $(DB_URL) $(MIGRATE_ARGS)
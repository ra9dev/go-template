DB_PORT ?= 5432
DB_USER ?= postgres
DB_NAME ?= go-template
DB_DRIVER := postgres
DB_URL ?= "$(DB_DRIVER)://$(DB_USER)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable"

MIGRATION_NAME ?= migration
MIGRATIONS_DIR ?= migrations
MIGRATIONS_EXT ?= sql
migration:
	$(call describe_job,"Generating migration '$(MIGRATION_NAME)'")
	$(MAKE) migrate-deps
	$(LOCAL_BIN)/migrate create -ext $(MIGRATIONS_EXT) -dir $(MIGRATIONS_DIR) -seq $(MIGRATION_NAME)

MIGRATE_ARGS ?= up
MIGRATE_URL ?= $(DB_URL)
migrate:
	$(call describe_job,"Migrating $(MIGRATE_ARGS)")
	$(MAKE) migrate-deps
	$(LOCAL_BIN)/migrate -path $(MIGRATIONS_DIR) -verbose -database $(MIGRATE_URL) $(MIGRATE_ARGS)
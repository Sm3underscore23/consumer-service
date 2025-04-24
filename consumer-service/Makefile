LOCAL_BIN := "$(CURDIR)"/iternal/bin
LOCAL_MIGRATION_DIR := "$(CURDIR)"/migrations
LOCAL_MIGRATION_DSN := "host=localhost port=8082 dbname=order_db user=admin password=1234"

install-goose:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.24.2

local-migration-create:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) create create_tables sql

local-migration-status:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir $(LOCAL_MIGRATION_DIR) postgres $(LOCAL_MIGRATION_DSN) down -v
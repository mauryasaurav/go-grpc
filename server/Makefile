BINARY=bin/engine

DB := go-grpc
DB_DIR := db
POSTGRES_USER := postgres
POSTGRES_HOST := localhost
POSTGRES_PASSWORD := testing123
POSTGRES_PORT := 5432


.PHONY: migrate-prepare
migrate-prepare:
	@echo "Installing golang-migrate"
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: db-migrateup
db-migrateup:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ up

.PHONY: db-migratedown
db-migratedown:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ down

.PHONY: db-migrateforce
db-migrateforce:
	migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ force 1

# Application settings
BINARY_NAME=app
BINARY_DIR=bin
MAIN_FILE=cmd/app/main.go
MIGRATIONS_DIR=db/migrations

# Include .env file
-include .env
export

# Database connection string (using env vars)
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)

# Build commands
build:
	@echo "Building application..."
	@mkdir -p $(BINARY_DIR)
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build complete: $(BINARY_DIR)/$(BINARY_NAME)"

run: build
	@echo "Starting application..."
	@./$(BINARY_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build files..."
	@rm -rf $(BINARY_DIR)
	@echo "Clean complete"

# Migration commands
migration:
	@if [ -z "$(name)" ]; then \
        echo "Error: Migration name is required. Usage: make migration name=<migration_name>"; \
        exit 1; \
    else \
        migrate create -ext sql -dir "$(MIGRATIONS_DIR)" "$(name)"; \
        echo "Migration created: $(name)"; \
    fi

migrate-up:
	@echo "====== MIGRATE UP ======"
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up


migrate-down:
	@echo "====== MIGRATE DOWN ======"
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" down

migrate-version:
	@echo "====== MIGRATE VERSION ======"
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" version

migrate-force:
	@echo "====== MIGRATE FORCE ======"
	migrate -path "$(MIGRATIONS_DIR)" -database "$(DATABASE_URL)" force $(VERSION)

migrate-fresh:
	@echo "====== MIGRATE FRESH ======"

	@echo "====== MIGRATE DOWN ======"
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down

	@echo "====== MIGRATE UP ======"
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up


db-test:
	@echo "Testing database connection..."
	@echo "Database URL: $(DATABASE_URL)"
	@if command -v psql >/dev/null; then \
		psql "$(DATABASE_URL)" -c "\conninfo"; \
	else \
		echo "psql not found. Please install PostgreSQL client tools."; \
	fi

.PHONY: build run clean migration migrate-up migrate-down migrate-force migrate-version db-test

sqlc:
	@echo "Generate sqlc..." && \
	sqlc generate

	@if command -v psql >/dev/null; then \
		echo "successfully generate sqlc..."; \
	fi

list:
	@echo "Available make commands:"
	@echo "  build              - Build the application"
	@echo "  run                - Build and run the application"
	@echo "  clean              - Remove build files"
	@echo "  sqlc               - Generate sqlc"
	@echo "  migration          - Create new migration (usage: make migration name=<name>)"
	@echo "  migrate-up         - Run all pending migrations"
	@echo "  migrate-up-one     - Run next pending migration"
	@echo "  migrate-down       - Rollback last migration"
	@echo "  migrate-down-all   - Rollback all migrations"
	@echo "  migrate-force      - Force migration version (usage: make migrate-force version=<version>)"
	@echo "  migrate-version    - Show current migration version"
	@echo "  db-test            - Test database connection"
	@echo "  list               - Show this help message"
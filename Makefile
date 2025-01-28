# THANKS TO CHATGPT FOR MAKE THIS HELPER

# Application settings
BINARY_NAME=app
BINARY_DIR=bin
MAIN_FILE=cmd/app/main.go

# Database settings
MIGRATIONS_DIR=db/migrations
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=meedeley
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Build commands
.PHONY: build run clean migration migrate-up migrate-down migrate-force migrate-version db-test

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
		migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq "$(name)"; \
		echo "Migration file created for: $(name)"; \
	fi

migrate-up:
	@echo "Running all pending migrations..." && \
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up

migrate-up-one:
	@echo "Running next pending migration..." && \
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up 1

migrate-down:
	@echo "Rolling back last migration..." && \
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down 1

migrate-down-all:
	@echo "Rolling back all migrations..." && \
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Error: Version is required. Usage: make migrate-force version=<version_number>"; \
		exit 1; \
	else \
		echo "Forcing migration version to $(version)..." && \
		migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" force $(version); \
	fi

migrate-version:
	@echo "Current migration version:" && \
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" version

# Helper command to test database connection
db-test:
	@echo "Testing database connection..."
	@echo "Database URL: $(DATABASE_URL)"
	@if command -v psql >/dev/null; then \
		psql "$(DATABASE_URL)" -c "\conninfo"; \
	else \
		echo "psql not found. Please install PostgreSQL client tools."; \
	fi

run:
	go build -o bin/service cmd/app/main.go
	./bin/app

build :
	go build -o bin/app cmd/app/main.go




migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: Migration name is required. Usage: make migration arg=migration_name"; \
	else \
		migrate create -ext sql -dir db/migrations "create_$(name)_table"; \
		echo "Migration file created for: create_$(name)_table"; \
	fi


migrate-up:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && \
		echo "Running all pending migrations..." && \
		migrate -path db/migrations -database "$$DATABASE_URL" up; \
	else \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	fi

migrate-up-one:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && \
		echo "Running next pending migration..." && \
		migrate -path db/migrations -database "$$DATABASE_URL" up 1; \
	else \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	fi

migrate-down:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && \
		echo "Rolling back last migration..." && \
		migrate -path db/migrations -database "$$DATABASE_URL" down 1; \
	else \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	fi

migrate-down-all:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && \
		echo "Rolling back all migrations..." && \
		migrate -path db/migrations -database "$$DATABASE_URL" down; \
	else \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	fi

migrate-force:
	@if [ -z "$(version)" ]; then \
		echo "Error: Version is required. Usage: make migrate-force version=<version_number>"; \
	elif [ ! -f .env ]; then \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	else \
		export $$(cat .env | xargs) && \
		echo "Forcing migration version to $(version)..." && \
		migrate -path db/migrations -database "$$DATABASE_URL" force $(version); \
	fi

migrate-version:
	@if [ -f .env ]; then \
		export $$(cat .env | xargs) && \
		echo "Current migration version:" && \
		migrate -path db/migrations -database "$$DATABASE_URL" version; \
	else \
		echo "Error: .env file not found. Please create .env file with DATABASE_URL"; \
		exit 1; \
	fi
.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD TIDY..."
	@go mod tidy

	@echo "RUNNING GO MOD VENDOR..."
	@go mod vendor

postgres:
	@echo "Starting Postgres docker container"
	@docker start probi-postgres || (echo "Container probi-postgres not found. Run docker-compose up -d or create the container first." && exit 1)
	@echo "Postgres is running."
	@echo "============================"

redis:
	@echo "Starting Redis docker container"
	@docker start probi-redis || (echo "Container probi-redis not found. Run docker-compose up -d or create the container first." && exit 1)
	@echo "Redis is running."
	@echo "============================"

run: postgres redis
	@if [ "$$ENV" = "production" ]; then \
		fluent-bit -c fluent-bit.conf; \
	fi
	@go run cmd/main.go

clean:
	@docker stop probi-postgres probi-redis || true

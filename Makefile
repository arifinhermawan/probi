.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD TIDY..."
	@go mod tidy

	@echo "RUNNING GO MOD VENDOR..."
	@go mod vendor

# Start Docker Compose services
docker-compose-up:
	@echo "Starting Docker containers using docker-compose"
	@docker-compose up -d

# Stop Docker Compose services
docker-compose-down:
	@echo "Stopping Docker containers using docker-compose"
	@docker-compose down

# Run HTTP service
run-http: docker-compose-up
	@go run cmd/http/main.go

# Run Cron service
run-cron: docker-compose-up
	@go run cmd/cron/main.go

# Run NSQ service
run-nsq: docker-compose-up
	@go run cmd/nsq/main.go

# Fluent Bit service
fluent:
	fluent-bit -c fluent-bit.conf

# Clean up - stop services if running
clean: docker-compose-down
	@docker stop probi-postgres probi-redis || true
	@docker stop probi-nsq || true

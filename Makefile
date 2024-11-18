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

rabbit:
	@echo "Starting RabbitMQ docker container"
	@docker start probi-rabbitmq || (echo "Container probi-rabbitmq not found. Run docker-compose up -d or create the container first." && exit 1)
	@echo "rabbitmq is running."
	@echo "============================"

run-http: postgres redis rabbit
	@go run cmd/http/main.go

run-cron: postgres redis rabbit
	@go run cmd/cron/main.go

run-mq: postgres redis rabbit
	@go run cmd/mq/main.go

fluent:
	fluent-bit -c fluent-bit.conf

clean:
	@docker stop probi-postgres probi-redis probi-rabbitmq|| true

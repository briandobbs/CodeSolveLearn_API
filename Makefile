SHELL := /bin/bash
.NOTPARALLEL:
.PHONY: default stack-up stack-down stack-down-volumes run

# Include environment variables from a .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Default command (runs the application)
default: run

# Start the Docker containers in detached mode using docker-compose
stack-up:
	docker-compose --env-file .env up -d

# Stop and remove Docker containers using docker-compose
stack-down:
	docker-compose --env-file .env down

# Stop and remove Docker containers and volumes using docker-compose
stack-down-volumes:
	docker-compose --env-file .env down -v --remove-orphans

# Run the Go application (locally)
run:
	go run -tags=jsoniter cmd/codesolvelearn_api/main.go

# creates a new database migration file in db/migrations
# usage: make migrate-new NAME=[name of the migration file]
migrate-new:
	dbmate new ${NAME}

migrate-up-dbmate:
	dbmate --url "${DB_PROTOCOL}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}${DB_OPTIONS}" up

# create the database (if it does not already exist) and run any pending migrations
migrate-up: migrate-up-dbmate

# roll back the most recent migration
migrate-down:
	dbmate --url "${DB_PROTOCOL}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}${DB_OPTIONS}" down

test: unit-test

unit-test:
	go test -tags=jsoniter -v ./...

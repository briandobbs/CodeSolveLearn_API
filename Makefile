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

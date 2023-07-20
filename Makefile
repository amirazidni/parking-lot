PWD := $(shell pwd)

DOCKER_SERVICES := "$(shell cd docker/parking-lot && docker-compose ps --services)"
DOCKER_RUNNINGS := "$(shell cd docker/parking-lot && docker-compose ps --services --filter 'status=running')"

all: vendor build build-docker run-docker

vendor:
	@echo "Update modules (go mod tidy & go mod vendor)"
	@go mod tidy
	@go mod vendor

build:
	@echo "Building parking-lot service"
	@go build .

build-docker:
	@echo "Building parking-lot docker image"
	@mv ./parking-lot ./docker/parking-lot/
	@cd docker/parking-lot && docker-compose build

run-docker:
	@if [ $(DOCKER_SERVICES) != $(DOCKER_RUNNINGS) ]; then \
		echo "Run docker containers"; \
		cd docker/parking-lot && docker-compose up -d; \
	else \
		echo "Restart docker containers"; \
		cd docker/parking-lot && docker-compose restart; \
	fi
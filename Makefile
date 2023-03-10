.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t kokifourier/gotodo:${DOCKER_TAG} --target deploy ./

up: ## Do docker compose up
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Do docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	cd src && go test -race -shuffle=on ./...

migrate: ## Execute migration
	mysqldef --user=todo --password=todo --port=33306 --host=localhost todo < ./_tools/mysql/schema.sql

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

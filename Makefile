include .env
export

export LINTER=$(shell go env GOPATH)/bin/golangci-lint-v2

export PROJECT_ROOT=$(CURDIR)

env/up:
	docker compose up -d git-diff-app-postgres

env/down:
	docker compose down git-diff-app-postgres

env/reset:
	@echo -n " > Confirm env-reset? [y/N] " && read ans && [ $${ans:-N} = y ]
	docker compose down git-diff-app-postgres
	rm -rf ./pgdata

migration/create:
	@[ "$(name)" ] || { echo "Example usage: make migration/create name=migration_name"; exit 1; }
	docker compose run --rm git-diff-app-migrate create -ext sql -dir /migrations -seq $(name)

migration/upgrade:
	@make migration/action action=up

migration/downgrade:
	@make migration/action action=down

migration/action:
	@[ "$(action)" ] || { echo "Example usage: make migration/action action=up 3"; exit 1; }
	docker compose run --rm git-diff-app-migrate -path /migrations \
    -database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@git-diff-app-postgres:5432/${POSTGRES_DB}?sslmode=disable \
    "$(action)"

run:
	${LINTER} run
	go mod tidy && \
	go run cmd/git-diff-app/main.go

run/linter:
	${LINTER} run

run/git-diff-app:
	go mod tidy && \
	go run cmd/git-diff-app/main.go
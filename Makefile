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
	rm -rf ${PGDATA_HOME}

test/env/up:
	docker compose up -d git-diff-app-test-postgres

test/env/reset:
	@make test/env/down

test/env/down:
	docker compose down git-diff-app-test-postgres

test/migrate/up:
	@make test/migrate/action action=up

test/migrate/down:
	@make test/migrate/action action=down

test/migrate/action:
	docker compose run --rm git-diff-app-migrate -path /migrations \
	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@git-diff-app-test-postgres:5432/${POSTGRES_DB}?sslmode=disable \
	"$(action)"

test:
	go run cmd/tests/integration_tests.go

migrate/create:
	@[ "$(name)" ] || { echo "Example usage: make migrate/create name=migration_name"; exit 1; }
	docker compose run --rm git-diff-app-migrate create -ext sql -dir /migrations -seq $(name)

migrate/up:
	@make migrate/action action=up

migrate/down:
	@make migrate/action action=down

migrate/action:
	@[ "$(action)" ] || { echo "Example usage: make migrate/action action=up 3"; exit 1; }
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

ssl/keygen:
	mkdir -p keys && \
	openssl req -x509 -newkey rsa:4096 \
	-keyout keys/server.key \
	-out keys/server.crt \
	-days 365 -nodes -subj "/CN=localhost"
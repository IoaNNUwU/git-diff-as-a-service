SHELL := /usr/bin/bash

include .env
export

export PROJECT_ROOT=$(CURDIR)

env-up:
	docker compose up -d git-diff-app-postgres

env-down:
	docker compose down git-diff-app-postgres

env-reset:

all:
    @set /p name="Enter your name: " & call echo Hello %%name%%


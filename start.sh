#!/bin/bash

set -o allexport
source .env
set +o allexport

docker compose --env-file .env -f docker-compose-local.yml up -d

tern migrate --migrations internal/pgstore/migrations --config internal/pgstore/migrations/tern.conf
printf " \033[0;32m✔\033[0m tern migrate \n"

sqlc generate -f internal/pgstore/sqlc.yml
printf " \033[0;32m✔\033[0m sqlc generate \n"

go run cmd/journey/journey.go

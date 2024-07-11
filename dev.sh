#!/bin/bash

set -o allexport
source .env
set +o allexport

docker compose --env-file .env -f docker-compose-local.yml up -d

goapi-gen --package=spec --out internal/api/spec/journey.gen.spec.go internal/api/spec/journey.spec.json
printf " \033[0;32m✔\033[0m API Specs \n"

tern migrate --migrations internal/pgstore/migrations --config internal/pgstore/migrations/tern.conf
printf " \033[0;32m✔\033[0m Migrations \n"

sqlc generate -f internal/pgstore/sqlc.yml
printf " \033[0;32m✔\033[0m SQL generated \n"

go run cmd/journey/journey.go
